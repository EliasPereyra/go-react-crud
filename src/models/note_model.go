package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	_id         primitive.ObjectID `bson:"_id"`
	Title       string             `json:"name" form:"name"`
	Description string             `json:"description" form:"description"`
	Img_url     string             `json:"img_url" form:"img_url"`
	Created_at  primitive.DateTime `json:"created_at" form:"created_at"`
	Updated_at  primitive.DateTime `json:"updated_at" form:"updated_at"`
	User_id     string             `json:"user_id" form:"user_id"`
	Category_id int                `json:"category_id" form:"category_id"`
}
