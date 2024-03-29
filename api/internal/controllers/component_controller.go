package controllers

import (
	"ciri2-pc-microservice/internal/models"
	"ciri2-pc-microservice/internal/responses"
	"ciri2-pc-microservice/internal/services"
	"ciri2-pc-microservice/internal/utils"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ComponentController struct{}

var componentService services.ComponentService

var validate = validator.New()

func (c ComponentController) BatchCreateComponent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var components []models.Component
		//validate the request body
		if !utils.BindJSONAndValidate(c, &components) {
			return
		}

		//use the validator library to validate required fields
		for _, component := range components {
			if validationErr := validate.Struct(&component); validationErr != nil {
				c.JSON(http.StatusBadRequest, responses.ComponentResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
				return
			}
		}

		result, err := componentService.BatchCreateComponent(components)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ComponentResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.ComponentResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func (c ComponentController) GetComponent() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		result, err := componentService.GetComponent(id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ComponentResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.ComponentResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})

	}
}

func (c ComponentController) FindPaginatedComponent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pagination models.Pagination

		if !utils.BindJSONAndValidate(c, &pagination) {
			return
		}

		if validationErr := validate.Struct(&pagination); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ComponentResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		result, err := componentService.FindPaginatedComponent(pagination)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ComponentResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.ComponentResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}
