/*
 * Wodby API
 *
 * API documentation https://docs.wodby.com/dev.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type RequestInstanceCreate struct {

	AppId *Uuid `json:"app_id"`

	Git *RequestInstanceCreateGit `json:"git,omitempty"`

	Name string `json:"name"`

	PostDeployment bool `json:"post_deployment,omitempty"`

	ServerId *Uuid `json:"server_id"`

	Title string `json:"title,omitempty"`

	Type_ *InstanceType `json:"type"`
}
