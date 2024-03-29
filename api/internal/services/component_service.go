package services

import (
	"ciri2-pc-microservice/internal/models"
	"ciri2-pc-microservice/internal/repositories"
)

type ComponentService struct{}

var componentService repositories.ComponentRepository

func (c ComponentService) BatchCreateComponent(components []models.Component) (interface{}, error) {
	interfaceSlice := make([]interface{}, len(components))
	for i, v := range components {
		interfaceSlice[i] = v
	}

	result, err := componentService.BatchInsert(interfaceSlice)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c ComponentService) GetComponent(id string) (interface{}, error) {

	result, err := componentService.FindOne(id)
	if err != nil {
		return nil, err
	}
	print(result)
	return result, nil
}

func (c ComponentService) FindPaginatedComponent(pagination models.Pagination) (interface{}, error) {
	result, err := componentService.FindPaginated(pagination)
	if err != nil {
		return nil, err
	}
	return result, nil
}
