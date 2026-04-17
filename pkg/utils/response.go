package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{Success: true, Data: data})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, Response{Success: true, Data: data})
}

func Message(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{Success: true, Message: message})
}

func BadRequest(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, ErrorResponse{Success: false, Error: err})
}

func Unauthorized(c *gin.Context, err string) {
	c.JSON(http.StatusUnauthorized, ErrorResponse{Success: false, Error: err})
}

func Forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, ErrorResponse{Success: false, Error: "insufficient permissions"})
}

func NotFound(c *gin.Context, err string) {
	c.JSON(http.StatusNotFound, ErrorResponse{Success: false, Error: err})
}

func InternalError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, ErrorResponse{Success: false, Error: "internal server error"})
}
