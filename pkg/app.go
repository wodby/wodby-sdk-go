/*
 * Wodby API (beta)
 *
 * API documentation https://docs.wodby.com/dev.
 *
 * API version: 3.0.x
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type App struct {

	Id string `json:"id"`

	Title string `json:"title"`

	Name string `json:"name"`

	OrgId string `json:"org_id"`

	GitRepoId string `json:"git_repo_id,omitempty"`

	Status string `json:"status"`

	Created int32 `json:"created"`

	Updated int32 `json:"updated"`
}
