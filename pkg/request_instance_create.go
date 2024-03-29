/*
 * Wodby API Client
 *
 * Wodby Developer Documentation https://wodby.com/docs/dev
 *
 * API version: 3.0.18
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type RequestInstanceCreate struct {

	AppId string `json:"app_id"`

	Git *RequestInstanceCreateGit `json:"git,omitempty"`

	Name string `json:"name"`

	PostDeployment bool `json:"post_deployment,omitempty"`

	ServerId string `json:"server_id"`

	Title string `json:"title,omitempty"`

	Type_ *InstanceType `json:"type"`
}
