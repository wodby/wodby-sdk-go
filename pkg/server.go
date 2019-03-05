/*
 * Wodby API Client
 *
 * Wodby Developer Documentation https://wodby.com/docs/dev
 *
 * API version: 3.0.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type Server struct {

	Created int32 `json:"created"`

	Id string `json:"id"`

	IpAddress string `json:"ip_address,omitempty"`

	OrgId string `json:"org_id"`

	Status string `json:"status"`

	Title string `json:"title"`

	Updated int32 `json:"updated"`

	Version string `json:"version"`
}
