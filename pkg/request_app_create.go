/*
 * Wodby API Client
 *
 * Wodby Developer Documentation https://wodby.com/docs/dev
 *
 * API version: 3.0.14
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type RequestAppCreate struct {

	DeploymentType string `json:"deployment_type,omitempty"`

	Docroot string `json:"docroot,omitempty"`

	Git *RequestAppCreateGit `json:"git,omitempty"`

	InstanceName string `json:"instance_name,omitempty"`

	InstanceTitle string `json:"instance_title,omitempty"`

	InstanceType *InstanceType `json:"instance_type,omitempty"`

	Name string `json:"name"`

	OrgId string `json:"org_id"`

	PostDeployment bool `json:"post_deployment,omitempty"`

	ServerId string `json:"server_id"`

	Services []RequestAppCreateServices `json:"services,omitempty"`

	Sitename string `json:"sitename,omitempty"`

	StackId string `json:"stack_id"`

	Title string `json:"title,omitempty"`
}
