package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	ProductName string             `json:"product_name" bson:"product_name"`
	Description string             `json:"description"  bson:"description"`
	Quantity    string             `json:"quantity"     bson:"quantity"`
	Price       string             `json:"price"        bson:"price"`
	Ratings     uint               `json:"ratings"      bson:"ratings"`
	ImageUrl    string             `json:"image_url"    bson:"image_url"`
}
