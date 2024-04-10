// Package services is the business layer between the controllers and the repositories, and it contains the business logic for the application
package services

import (
	"ciri2-pc-microservice/internal/models"
	"ciri2-pc-microservice/internal/repositories"
)

type ComponentService struct{}

var componentRepository repositories.ComponentRepository

// BatchCreateComponent creates multiple components
func (c ComponentService) BatchCreateComponent(components []models.ComponentInput) (interface{}, error) {
	interfaceSlice := make([]interface{}, len(components))
	for i, v := range components {
		interfaceSlice[i] = models.Component{
			Name:       v.Name,
			Type:       v.Type,
			Properties: v.Properties,
		}
	}

	result, err := componentRepository.BatchInsert(interfaceSlice)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetComponent gets a component by id
func (c ComponentService) GetComponent(id string) (interface{}, error) {

	result, err := componentRepository.FindOne(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindPaginatedComponent finds components paginated
func (c ComponentService) FindPaginatedComponent(pagination models.Pagination) (interface{}, int, error) {
	result, total, err := componentRepository.FindPaginated(pagination)
	if err != nil {
		return nil, 0, err
	}
	return result, total, nil
}

// CheckAlreadyExistingComponents checks if components already exist
func (c ComponentService) CheckAlreadyExistingComponents(componentNames []string) ([]string, error) {
	return componentRepository.CheckAlreadyExistingComponents(componentNames)
}
