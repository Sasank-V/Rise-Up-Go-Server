package job

import (
	"fmt"
	"sync"
	"time"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/core/user"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/lib"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/types"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var JobColl *mongo.Collection
var connectJob sync.Once

func ConnectJobCollection() {
	connectJob.Do(func() {
		db := database.InitDB()
		CreateJobCollection(db)
		JobColl = db.Collection(lib.JobCollectionName)
	})
}

func GetAllJobs(page int64) ([]types.FullJob, int64, error) {
	ctx, cancel := database.GetContext()
	defer cancel()

	const pageSize int64 = 20
	skip := (page - 1) * pageSize

	total, err := JobColl.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	opts := options.Find().SetSkip(skip).SetLimit(pageSize)

	cursor, err := JobColl.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var jobs []Job
	if err = cursor.All(ctx, &jobs); err != nil {
		return nil, 0, err
	}

	var (
		allJobs   = make([]types.FullJob, len(jobs))
		wg        sync.WaitGroup
		jobErr    error
		userInfos = make([]user.User, len(jobs))
	)

	for i, job := range jobs {
		wg.Add(1)

		go func(i int, job Job) {
			defer wg.Done()

			userInfo, err := user.GetBasicUserInfo(job.Owner)
			if err != nil {
				jobErr = fmt.Errorf("error fetching user info for job %s: %w", job.ID, err)
				return
			}

			userInfos[i] = userInfo
			allJobs[i] = types.FullJob{
				ID:                 job.ID,
				Owner:              job.Owner,
				OwnerName:          userInfo.Name,
				OwnerPicture:       userInfo.Picture,
				Title:              job.Title,
				Description:        job.Description,
				SkillTags:          job.SkillTags,
				WorkMode:           string(job.WorkMode),
				JobType:            string(job.JobType),
				Location:           job.Location,
				SalaryRangeStart:   job.SalaryRangeStart,
				SalaryRangeEnd:     job.SalaryRangeEnd,
				EvaluationCriteria: job.EvaluationCriteria,
				Contact:            job.Contact,
				PostedAt:           job.PostedAt,
				Active:             job.Active,
			}
		}(i, job)
	}

	wg.Wait()

	if jobErr != nil {
		return []types.FullJob{}, 0, jobErr
	}

	return allJobs, total, nil
}

func AddJob(info types.CreateJobRequest) error {
	ctx, cancel := database.GetContext()
	defer cancel()

	newJob := Job{
		Owner:              info.Owner,
		Title:              info.Title,
		Description:        info.Description,
		SkillTags:          info.SkillTags,
		WorkMode:           WorkMode(info.WorkMode),
		JobType:            JobType(info.JobType),
		Location:           info.Location,
		SalaryRangeStart:   info.SalaryRangeStart,
		SalaryRangeEnd:     info.SalaryRangeEnd,
		EvaluationCriteria: info.EvaluationCriteria,
		Contact:            info.Contact,
		Active:             true,
		PostedAt:           time.Now(),
	}

	res, err := JobColl.InsertOne(ctx, newJob)
	if err != nil {
		return err
	}
	id, err := utils.GetInsertedIDAsString(res.InsertedID)
	if err != nil {
		return err
	}

	err = user.AddJobToOrganisation(info.Owner, id)
	return err
}

func UpdateJob(updateInfo types.UpdateJobRequest) error {
	objID, err := primitive.ObjectIDFromHex(updateInfo.JobID)
	if err != nil {
		return fmt.Errorf("invalid jobID: %v", err)
	}

	updateDoc := bson.M{}
	if updateInfo.Title != nil {
		updateDoc["title"] = *updateInfo.Title
	}
	if updateInfo.Description != nil {
		updateDoc["description"] = *updateInfo.Description
	}
	if updateInfo.SkillTags != nil {
		updateDoc["skill_tags"] = *updateInfo.SkillTags
	}
	if updateInfo.WorkMode != nil {
		updateDoc["work_mode"] = *updateInfo.WorkMode
	}
	if updateInfo.JobType != nil {
		updateDoc["job_type"] = *updateInfo.JobType
	}
	if updateInfo.Location != nil {
		updateDoc["location"] = *updateInfo.Location
	}
	if updateInfo.SalaryRangeStart != nil {
		updateDoc["salary_range_start"] = *updateInfo.SalaryRangeStart
	}
	if updateInfo.SalaryRangeEnd != nil {
		updateDoc["salary_range_end"] = *updateInfo.SalaryRangeEnd
	}
	if updateInfo.EvaluationCriteria != nil {
		updateDoc["evaluation_criteria"] = *updateInfo.EvaluationCriteria
	}
	if updateInfo.Active != nil {
		updateDoc["active"] = *updateInfo.Active
	}
	if updateInfo.Contact != nil {
		updateDoc["contact"] = *updateInfo.Contact
	}

	if len(updateDoc) == 0 {
		return fmt.Errorf("no fields to update")
	}

	update := bson.M{
		"$set": updateDoc,
	}

	// Perform the update.
	ctx, cancel := database.GetContext()
	defer cancel()
	filter := bson.M{"_id": objID, "owner": updateInfo.UserId}
	res, err := JobColl.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return fmt.Errorf("no job found with id %s", updateInfo.JobID)
	}
	return nil
}
