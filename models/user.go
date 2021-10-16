package models

import "gopkg.in/mgo.v2/bson"

type (
	// represents mongo model structure
	User struct {
		Id     bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name   string        `json:"name" bson:"name"`
		Gender string        `json:"gender" bson:"gender"`
		Age    int           `json:"age" bson:"age"`
	}
)
