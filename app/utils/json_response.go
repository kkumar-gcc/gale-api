package utils

import (
	"github.com/goravel/framework/contracts/http"

	"goravel/app/constants"
)

type JsonResponseBuilder struct {
	status  int
	message string
	errors  any
	data    any
}

func NewJsonResponse() *JsonResponseBuilder {
	return &JsonResponseBuilder{}
}

func (builder *JsonResponseBuilder) SetStatus(status int) *JsonResponseBuilder {
	builder.status = status
	return builder
}

func (builder *JsonResponseBuilder) SetMessage(message constants.Message) *JsonResponseBuilder {
	builder.status = message.Status
	builder.message = message.Message
	return builder
}

func (builder *JsonResponseBuilder) SetErrors(errors any) *JsonResponseBuilder {
	builder.errors = errors
	return builder
}

func (builder *JsonResponseBuilder) SetData(data any) *JsonResponseBuilder {
	builder.data = data
	return builder
}

func (builder *JsonResponseBuilder) Build(ctx http.Context) http.Response {
	response := make(map[string]any)

	if builder.status != 0 {
		response["status"] = builder.status
	}
	if builder.message != "" {
		response["message"] = builder.message
	}
	if builder.errors != nil {
		response["errors"] = builder.errors
	}
	if builder.data != nil {
		response["data"] = builder.data
	}

	return ctx.Response().Status(builder.status).Json(response)
}
