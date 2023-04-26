/*
 * Wodby API Client
 *
 * Wodby Developer Documentation https://wodby.com/docs/dev
 *
 * API version: 3.0.18
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type StackService struct {

	Default_ bool `json:"default,omitempty"`

	Id string `json:"id,omitempty"`

	Implementations []StackServiceImplementation `json:"implementations,omitempty"`

	Name string `json:"name,omitempty"`

	Required bool `json:"required,omitempty"`

	Title string `json:"title,omitempty"`
}
