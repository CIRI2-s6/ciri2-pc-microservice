// Package controllers is the entry point for the application
package controllers

import (
	"ciri2-pc-microservice/internal/models"
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

// BatchCreateComponent ... Batch create multiple components
// @Summary Batch create multiple components
// @Description Batch create multiple components
// @Tags Components
// @Success 200 {object} models.ComponentResponse
// @Failure 404 {object} object
// @Router /component [post]
// @Param components body []models.ComponentInput true "Components"
func (c ComponentController) BatchCreateComponent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var components []models.ComponentInput
		//validate the request body
		if !utils.BindJSONAndValidate(c, &components) {
			return
		}
		//use the validator library to validate required fields
		for _, component := range components {
			if validationErr := validate.Struct(&component); validationErr != nil {
				c.JSON(http.StatusBadRequest, models.ComponentResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
				return
			}
		}

		result, err := componentService.BatchCreateComponent(components)

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ComponentResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, models.ComponentResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

// GetComponent ... Get one component
// @Summary Get one component by id
// @Description Get one component by id
// @Tags Components
// @Success 200 {object} models.ComponentResponse
// @Failure 404 {object} object
// @Router /component/{id} [get]
// @Param id path string true "Component ID"
func (c ComponentController) GetComponent() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		result, err := componentService.GetComponent(id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ComponentResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, models.ComponentResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})

	}
}

// GetPaginatedComponent ... Get paginated components
// @Summary Get components with pagination
// @Description Endpoint will get the components that match the pagination parameters
// @Tags Components
// @Success 200 {object} models.ComponentResponse
// @Failure 404 {object} object
// @Router /component [get]
// @Param skip query int false "Skip"
// @Param limit query int false "Limit"
// @Param sort query string false "Sort"
// @Param filter query string false "Filter"
func (c ComponentController) GetPaginatedComponent() gin.HandlerFunc {
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
				c.JSON(http.StatusBadRequest, models.ComponentResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			pagination.Sort = sort
		}

		// Parse the filter parameter as JSON if provided
		filterJSON := c.Query("filter")
		if filterJSON != "" {
			var filter map[string]interface{}
			if err := json.Unmarshal([]byte(filterJSON), &filter); err != nil {
				c.JSON(http.StatusBadRequest, models.ComponentResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			pagination.Filter = filter
		}

		if validationErr := validate.Struct(&pagination); validationErr != nil {
			c.JSON(http.StatusBadRequest, models.ComponentResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		result, total, err := componentService.FindPaginatedComponent(pagination)

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ComponentResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, models.ComponentResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result, "total": total}})
	}
}

// CheckAlreadyExistingComponents ... Check if components already exist
// @Summary Check if components already exist
// @Description Check if components already exist
// @Tags Components
// @Success 200 {object} models.ComponentResponse
// @Failure 404 {object} object
// @Router /component/check [post]
// @Param components body []string true "Component Names"
func (c ComponentController) CheckAlreadyExistingComponents() gin.HandlerFunc {
	return func(c *gin.Context) {
		var componentNames []string

		if !utils.BindJSONAndValidate(c, &componentNames) {
			return
		}
		result, err := componentService.CheckAlreadyExistingComponents(componentNames)

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ComponentResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, models.ComponentResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}
