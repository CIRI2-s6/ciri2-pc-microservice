package controllers

import (
	"ciri2-pc-microservice/models"
	"ciri2-pc-microservice/responses"
	"ciri2-pc-microservice/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func BatchCreateComponent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var components []models.Component
		//validate the request body
		if err := c.BindJSON(&components); err != nil {
			c.JSON(http.StatusBadRequest, responses.ComponentResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		for _, component := range components {
			if validationErr := validate.Struct(&component); validationErr != nil {
				c.JSON(http.StatusBadRequest, responses.ComponentResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
				return
			}
		}

		var newComponents []interface{}
		for _, component := range components {
			newComponent := models.Component{
				Id:         primitive.NewObjectID(),
				Name:       component.Name,
				Type:       component.Type,
				Properties: component.Properties,
			}
			newComponents = append(newComponents, newComponent)
		}

		result, err := services.BatchCreateComponent(newComponents)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ComponentResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.ComponentResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetComponent() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		result, err := services.GetComponent(id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ComponentResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.ComponentResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})

	}
}
