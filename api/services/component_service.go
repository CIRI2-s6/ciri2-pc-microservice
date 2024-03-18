package services

import (
	"ciri2-pc-microservice/repositories"
)

func BatchCreateComponent(components []interface{}) (interface{}, error) {
	result, err := repositories.BatchInsert(components)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetComponent(id string) (interface{}, error) {

	result, err := repositories.FindOne(id)
	if err != nil {
		return nil, err
	}
	print(result)
	return result, nil
}
