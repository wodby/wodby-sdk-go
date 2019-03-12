/*
 * Wodby API Client
 *
 * Wodby Developer Documentation https://wodby.com/docs/dev
 *
 * API version: 3.0.11
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type Instance struct {

	AppId string `json:"app_id"`

	Created int32 `json:"created"`

	GitRepoId string `json:"git_repo_id,omitempty"`

	GitRepoTarget string `json:"git_repo_target,omitempty"`

	HasNewVersion bool `json:"has_new_version,omitempty"`

	Id string `json:"id"`

	Name string `json:"name"`

	OrgId string `json:"org_id"`

	Status string `json:"status"`

	Title string `json:"title"`

	Type_ *InstanceType `json:"type,omitempty"`

	Updated int32 `json:"updated"`
}
