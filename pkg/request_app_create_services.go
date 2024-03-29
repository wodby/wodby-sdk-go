/*
 * Wodby API Client
 *
 * Wodby Developer Documentation https://wodby.com/docs/dev
 *
 * API version: 3.0.18
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type RequestAppCreateServices struct {

	Enable bool `json:"enable,omitempty"`

	Implementation string `json:"implementation,omitempty"`

	Name string `json:"name"`
}
