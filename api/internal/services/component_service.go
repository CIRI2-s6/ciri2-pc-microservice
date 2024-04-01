package services

import (
	"ciri2-pc-microservice/internal/models"
	"ciri2-pc-microservice/internal/repositories"
)

type ComponentService struct{}

var componentRepository repositories.ComponentRepository

func (c ComponentService) BatchCreateComponent(components []models.Component) (interface{}, error) {
	interfaceSlice := make([]interface{}, len(components))
	for i, v := range components {
		interfaceSlice[i] = v
	}

	result, err := componentRepository.BatchInsert(interfaceSlice)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c ComponentService) GetComponent(id string) (interface{}, error) {

	result, err := componentRepository.FindOne(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c ComponentService) FindPaginatedComponent(pagination models.Pagination) (interface{}, int, error) {
	result, total, err := componentRepository.FindPaginated(pagination)
	if err != nil {
		return nil, 0, err
	}
	return result, total, nil
}

func (c ComponentService) CheckAlreadyExistingComponents(componentNames []string) ([]string, error) {
	return componentRepository.CheckAlreadyExistingComponents(componentNames)
}
