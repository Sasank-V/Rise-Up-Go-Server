package models

type Review struct {
	From   string `bson:"from" json:"from"`
	To     string `bson:"to" json:"to"`
	Rating int    `bson:"rating" json:"rating"`
	Body   string `bson:"body" json:"body"`
}
