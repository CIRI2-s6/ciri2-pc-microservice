package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ComponentType string

const (
	CPU ComponentType = "CPU"
	GPU ComponentType = "GPU"
	RAM ComponentType = "RAM"
)

type Component struct {
	Id         primitive.ObjectID     `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string                 `json:"name,omitempty" validate:"required"`
	Type       ComponentType          `json:"type,omitempty" validate:"required"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}
