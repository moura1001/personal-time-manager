package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Http struct {
	Description string `json:"description,omitempty"`
	Metadata    string `json:"metadata,omitempty"`
	StatusCode  int    `json:"statusCode"`
}

func (e Http) Error() string {
	return fmt.Sprintf("description: %s,  metadata: %s", e.Description, e.Metadata)
}

func NewHttpError(description, metadata string, statusCode int) Http {
	return Http{
		Description: description,
		Metadata:    metadata,
		StatusCode:  statusCode,
	}
}

func NewDefaultHttpError(err error) Http {
	return Http{
		Description: err.Error(),
		StatusCode:  http.StatusInternalServerError,
	}
}

func DefaultErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case Http:
				c.AbortWithStatusJSON(e.StatusCode, e)
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
					"message":     "Service Unavailable",
					"description": e.Error(),
				})
			}
		}
	}
}
