/*
 * Wodby API (beta)
 *
 * API documentation https://docs.wodby.com/dev.
 *
 * API version: 3.0.x
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type Stack struct {

	Id string `json:"id"`

	Title string `json:"title"`

	Version string `json:"version,omitempty"`

	RevisionNumber int32 `json:"revision_number,omitempty"`

	Services []StackService `json:"services"`

	OrgId string `json:"org_id"`

	Created int32 `json:"created"`

	Updated int32 `json:"updated"`
}
