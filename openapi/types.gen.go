// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package openapi

// SigninRequest defines model for SigninRequest.
type SigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignupRequest defines model for SignupRequest.
type SignupRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// Task defines model for Task.
type Task struct {
	Title string `json:"title"`
}

// PostSigninJSONBody defines parameters for PostSignin.
type PostSigninJSONBody SigninRequest

// PostSignupJSONBody defines parameters for PostSignup.
type PostSignupJSONBody SignupRequest

// PostTasksJSONBody defines parameters for PostTasks.
type PostTasksJSONBody Task

// PostSigninJSONRequestBody defines body for PostSignin for application/json ContentType.
type PostSigninJSONRequestBody PostSigninJSONBody

// PostSignupJSONRequestBody defines body for PostSignup for application/json ContentType.
type PostSignupJSONRequestBody PostSignupJSONBody

// PostTasksJSONRequestBody defines body for PostTasks for application/json ContentType.
type PostTasksJSONRequestBody PostTasksJSONBody