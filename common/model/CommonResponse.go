package model

import (
	"AdminPro/common/enum"
)

// CommonResponse represents a generic response structure.
type CommonResponse[T any] struct {
	Code enum.ResponseCodeEnum `json:"code"`
	Msg  string                `json:"msg"`
	Data T                     `json:"data,omitempty"`
}

// IsSuccess checks if the response is successful.
func (r *CommonResponse[any]) IsSuccess() bool {
	return r.Code == enum.SUCCESS
}

// Success creates a successful response with no data.
func (r *CommonResponse[any]) Success() *CommonResponse[any] {
	r.Code = enum.SUCCESS
	return r
}

func (r *CommonResponse[any]) SuccessFrom(msg string, data any) *CommonResponse[any] {
	r.Code = enum.SUCCESS
	r.Msg = msg
	r.Data = data
	return r
}

// SuccessWithMsg creates a successful response with a custom message.
func (r *CommonResponse[any]) SuccessFromMsg(msg string) *CommonResponse[any] {
	r.Code = enum.SUCCESS
	r.Msg = msg
	return r

}

// SuccessWithData creates a successful response with data.
func (r *CommonResponse[any]) SuccessFromData(data any) *CommonResponse[any] {
	r.Code = enum.SUCCESS
	r.Data = data
	return r
}

// Failure creates a failure response with a predefined code.
func (r *CommonResponse[any]) Failure(code enum.ResponseCodeEnum) *CommonResponse[any] {
	r.Code = code
	r.Msg = enum.GetResponseMsg(code)
	return r
}

// FailureWithMsg creates a failure response with a predefined code and custom message.
func (r *CommonResponse[any]) FailureFromMsg(code enum.ResponseCodeEnum, msg string) *CommonResponse[any] {
	r.Code = code
	r.Msg = msg
	return r
}

// FailureWithMsg creates a failure response with a predefined code and custom message.
func (r *CommonResponse[any]) FailureFromError(msg string) *CommonResponse[any] {
	r.Code = enum.ERROR
	r.Msg = msg
	return r
}
