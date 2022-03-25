package services

import (
	"github.com/go-kafka-microservice/InventoryService/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryServices interface {
	RegisterInventory(*models.Inventory) error
	AddProduct(primitive.ObjectID, *models.Product) error
	GetProduct(primitive.ObjectID, primitive.ObjectID) (*models.Product, error)
}
