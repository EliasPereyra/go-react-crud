package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	_id				primitive.ObjectID `bson:"_id"`
	Name			string		`json:"name" form:"name"`
	Lastname	string		`json:"lastName" form:"lastName"`
	Age				int 			`json:"age" form:"age"`
	Note_id		string		`json:"note_id" form:"note_id"`
} 
