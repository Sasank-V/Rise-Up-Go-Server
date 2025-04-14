package user

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Sasank-V/Rise-Up-Go-Server/internal/database"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/lib"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/types"
	"github.com/Sasank-V/Rise-Up-Go-Server/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserColl *mongo.Collection
var userConnect sync.Once

func ConnectUserCollection() {
	userConnect.Do(func() {
		db := database.InitDB()
		CreateUserCollection(db)
		UserColl = db.Collection(lib.UserCollectionName)
	})
}

func UserExists(googleID string) (bool, error) {
	ctx, cancel := database.GetContext()
	defer cancel()

	filter := bson.M{
		"_id": googleID,
	}

	res := UserColl.FindOne(ctx, filter)
	if res.Err() == mongo.ErrNoDocuments {
		return false, nil
	} else if res.Err() != nil {
		return false, res.Err()
	} else {
		var user User
		res.Decode(&user)
		return true, nil

	}
}

func CheckUserRole(userID string, role string) (bool, error) {
	ctx, cancel := database.GetContext()
	defer cancel()

	var user User
	res := UserColl.FindOne(ctx, bson.M{
		"_id": userID,
	})
	if res.Err() == mongo.ErrNoDocuments {
		return false, fmt.Errorf("no User found with the given ID")
	} else if res.Err() != nil {
		return false, res.Err()
	}
	res.Decode(&user)
	if strings.Compare(user.Role, role) != 0 {
		return false, nil
	}
	return true, nil
}

func GetBasicUserInfo(userID string) (User, error) {
	ctx, cancel := database.GetContext()
	defer cancel()
	filter := bson.M{"_id": userID}

	var user User
	err := UserColl.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func AddUser(info types.SigninRequest) error {
	user := User{
		ID:           info.GoogleID,
		Name:         info.Name,
		Email:        info.Email,
		Picture:      info.Picture,
		Bio:          "",
		Location:     "",
		Role:         info.Role,
		RoleID:       "",
		AccessToken:  info.AccessToken,
		RefreshToken: info.RefreshToken,
		ExpiresAt:    info.ExpiresAt,
	}
	fmt.Println(info.Role)

	ctx, cancel := database.GetContext()
	defer cancel()

	//Create Generic User
	res, err := UserColl.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	insertedUserID, err := utils.GetInsertedIDAsString(res.InsertedID)
	if err != nil {
		return err
	}

	//Create Specific User
	var roleID string
	if strings.Compare(user.Role, string(LearnerRole)) == 0 {
		roleID, err = AddLearner(insertedUserID)
	} else if strings.Compare(user.Role, string(MentorRole)) == 0 {
		roleID, err = AddMentor(insertedUserID)
	} else {
		roleID, err = AddOrganisation(insertedUserID)
	}

	if err != nil {
		return err
	}

	//Link Generic and Specific User
	filter := bson.M{
		"_id": user.ID,
	}
	values := bson.M{
		"$set": bson.M{
			"role_id": roleID,
		},
	}
	res2, err := UserColl.UpdateOne(ctx, filter, values)
	if err != nil {
		return err
	}
	if res2.ModifiedCount == 0 {
		return fmt.Errorf("no user found with the user id")
	}

	fmt.Println("Creating specific user for role:", user.Role)
	fmt.Println("Generated role ID:", roleID, "Err:", err)

	return nil

}
