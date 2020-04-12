/*
 * Wodby API Client
 *
 * Wodby Developer Documentation https://wodby.com/docs/dev
 *
 * API version: 3.0.13
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type Task struct {

	Created int32 `json:"created"`

	End int32 `json:"end,omitempty"`

	Id string `json:"id"`

	OrgId string `json:"org_id"`

	Start int32 `json:"start,omitempty"`

	Status string `json:"status"`

	Title string `json:"title"`

	Ttl int32 `json:"ttl"`

	Updated int32 `json:"updated"`
}
