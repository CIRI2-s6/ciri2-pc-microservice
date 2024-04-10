// Package models holds the data models for the application
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

type ComponentInput struct {
	Name       string                 `json:"name,omitempty" validate:"required"`
	Type       ComponentType          `json:"type,omitempty" validate:"required"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}

type ComponentResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
