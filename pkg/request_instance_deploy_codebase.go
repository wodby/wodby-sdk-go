/*
 * Wodby API Client
 *
 * Wodby Developer Documentation https://wodby.com/docs/dev
 *
 * API version: 3.0.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type RequestInstanceDeployCodebase struct {

	Git *RequestInstanceCreateGit `json:"git,omitempty"`

	PostDeployment bool `json:"post_deployment,omitempty"`
}
