/*
 * Wodby API (beta)
 *
 * API documentation https://docs.wodby.com/dev.
 *
 * API version: 3.0.x
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type Backup struct {

	Id string `json:"id"`

	Type_ string `json:"type"`

	OrgId string `json:"org_id"`

	InstanceId string `json:"instance_id"`

	Eta int32 `json:"eta,omitempty"`

	Spent int32 `json:"spent,omitempty"`

	Status string `json:"status"`

	Created int32 `json:"created"`

	Updated int32 `json:"updated"`
}
