package utils

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetInsertedIDAsString(id interface{}) (string, error) {
	switch v := id.(type) {
	case primitive.ObjectID:
		return v.Hex(), nil
	case string:
		return v, nil
	default:
		return "", fmt.Errorf("unexpected InsertedID type: %T", id)
	}
}
