package controllers

import (
	"ciri2-pc-microservice/internal/models"
	"ciri2-pc-microservice/internal/responses"
	"ciri2-pc-microservice/internal/services"
	"ciri2-pc-microservice/internal/utils"
	"encoding/json"
	"strconv"

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

		// Get pagination parameters from URL
		pagination.Skip, _ = strconv.Atoi(c.Query("skip"))
		pagination.Limit, _ = strconv.Atoi(c.Query("limit"))

		// Parse the sort parameter as JSON if provided
		sortJSON := c.Query("sort")
		if sortJSON != "" {
			var sort map[string]interface{}
			if err := json.Unmarshal([]byte(sortJSON), &sort); err != nil {
				c.JSON(http.StatusBadRequest, responses.ComponentResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			pagination.Sort = sort
		}

		// Parse the filter parameter as JSON if provided
		filterJSON := c.Query("filter")
		if filterJSON != "" {
			var filter map[string]interface{}
			if err := json.Unmarshal([]byte(filterJSON), &filter); err != nil {
				c.JSON(http.StatusBadRequest, responses.ComponentResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			pagination.Filter = filter
		}

		if validationErr := validate.Struct(&pagination); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ComponentResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		result, total, err := componentService.FindPaginatedComponent(pagination)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ComponentResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.ComponentResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result, "total": total}})
	}
}

func (c ComponentController) CheckAlreadyExistingComponents() gin.HandlerFunc {
	return func(c *gin.Context) {
		var componentNames []string

		if !utils.BindJSONAndValidate(c, &componentNames) {
			return
		}
		result, err := componentService.CheckAlreadyExistingComponents(componentNames)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ComponentResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.ComponentResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}