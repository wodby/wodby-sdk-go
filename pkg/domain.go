/*
 * Wodby API
 *
 * The Wodby API is organized around REST. Our API has predictable, resource-oriented URLs, and uses HTTP response codes to indicate API errors. We use built-in HTTP features, like HTTP authentication and HTTP verbs, which are understood by off-the-shelf HTTP clients. JSON is returned by all API responses, including errors. 
 *
 * API version: 3.x
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type Domain struct {

	Created int32 `json:"created"`

	Enabled bool `json:"enabled,omitempty"`

	Hsts bool `json:"hsts,omitempty"`

	HstsSubdomains bool `json:"hsts_subdomains,omitempty"`

	Id string `json:"id"`

	Indexed bool `json:"indexed,omitempty"`

	InstanceId string `json:"instance_id"`

	Name string `json:"name"`

	OrgId string `json:"org_id"`

	Primary bool `json:"primary,omitempty"`

	Protected bool `json:"protected,omitempty"`

	RedirectToNonWww bool `json:"redirect_to_non_www,omitempty"`

	RedirectToWww bool `json:"redirect_to_www,omitempty"`

	Ssl bool `json:"ssl,omitempty"`

	SslCustom bool `json:"ssl_custom,omitempty"`

	SslRequired bool `json:"ssl_required,omitempty"`

	Status string `json:"status"`

	Type_ string `json:"type"`

	Updated int32 `json:"updated"`
}
