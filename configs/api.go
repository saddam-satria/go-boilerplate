package configs

import (
	"time"

	"github.com/gin-gonic/gin"
)

type HttpRoute func(httpRoute *gin.RouterGroup)

type ApiResponse struct {
	Message    interface{} `json:"message"`
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
	Time       string      `json:"timestamp"`
}

func NewApiResponse(ctx *ApiResponse) ApiResponse {
	ctx.Time = time.Now().Format("2006-01-02 15:04:05")
	return *ctx
}

type ResponseHandler struct {
	Engine *gin.Context
}

type ErrorMessage struct {
	Message   string `json:"message"`
	ErrorType string `json:"errorType"`
}

func (ctx *ResponseHandler) Success(message string, statusCode int, data interface{}) {
	response := NewApiResponse(&ApiResponse{
		Message:    message,
		StatusCode: statusCode,
		Data:       data,
	})

	ctx.Engine.JSON(statusCode, response)
}

func (ctx *ResponseHandler) Error(message []ErrorMessage, statusCode int, data interface{}) {
	response := NewApiResponse(&ApiResponse{
		Message:    message,
		StatusCode: statusCode,
		Data:       data,
	})

	ctx.Engine.JSON(statusCode, response)
}
