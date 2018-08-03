/*
 * Wodby API (beta)
 *
 * Documentation https://docs.wodby.com/dev GitHub https://github.com/wodby/wodby-api 
 *
 * API version: 3.0.x
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package wodby-api

type Instance struct {

	Id string `json:"id"`

	Name string `json:"name"`

	Title string `json:"title"`

	Type_ *InstanceType `json:"type,omitempty"`

	OrgId string `json:"org_id"`

	AppId string `json:"app_id"`

	GitRepoId string `json:"git_repo_id,omitempty"`

	Status string `json:"status"`

	Created int32 `json:"created"`

	Updated int32 `json:"updated"`
}
