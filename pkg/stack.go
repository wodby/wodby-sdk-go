/*
 * Wodby API
 *
 * API documentation https://docs.wodby.com/dev.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type Stack struct {

	Created int32 `json:"created"`

	Id string `json:"id"`

	NewVersion string `json:"new_version,omitempty"`

	OrgId string `json:"org_id"`

	RevisionNumber int32 `json:"revision_number,omitempty"`

	Services []StackService `json:"services"`

	Title string `json:"title"`

	Updated int32 `json:"updated"`

	Version string `json:"version,omitempty"`
}
