/*
 * Wodby API Client
 *
 * Wodby Developer Documentation https://wodby.com/docs/dev
 *
 * API version: 3.0.16
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type User struct {

	Created int32 `json:"created"`

	Email string `json:"email"`

	FirstName string `json:"first_name,omitempty"`

	Id string `json:"id"`

	LastName string `json:"last_name,omitempty"`

	Name string `json:"name"`

	Updated int32 `json:"updated"`
}
