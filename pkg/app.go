/*
 * Wodby API Client
 *
 * Wodby Developer Documentation https://wodby.com/docs/dev
 *
 * API version: 3.0.13
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type App struct {

	Created int32 `json:"created"`

	GitRepoId string `json:"git_repo_id,omitempty"`

	Id string `json:"id"`

	Name string `json:"name"`

	OrgId string `json:"org_id"`

	Status string `json:"status"`

	Title string `json:"title"`

	Updated int32 `json:"updated"`
}
