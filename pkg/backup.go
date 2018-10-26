/*
 * Wodby API Client
 *
 * Wodby Developer Documentation https://wodby.com/docs/dev
 *
 * API version: 3.0.7
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type Backup struct {

	Created int32 `json:"created"`

	Eta int32 `json:"eta,omitempty"`

	Id string `json:"id"`

	InstanceId string `json:"instance_id"`

	OrgId string `json:"org_id"`

	Spent int32 `json:"spent,omitempty"`

	Status string `json:"status"`

	Type_ string `json:"type"`

	Updated int32 `json:"updated"`
}
