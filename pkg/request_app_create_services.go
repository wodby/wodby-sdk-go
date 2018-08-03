/*
 * Wodby API (beta)
 *
 * Documentation https://docs.wodby.com/dev GitHub https://github.com/wodby/wodby-api 
 *
 * API version: 3.0.x
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package wodby-api

type RequestAppCreateServices struct {

	Name string `json:"name"`

	Implementation string `json:"implementation,omitempty"`

	Enable bool `json:"enable,omitempty"`
}