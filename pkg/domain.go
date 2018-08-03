/*
 * Wodby API (beta)
 *
 * Documentation https://docs.wodby.com/dev GitHub https://github.com/wodby/wodby-api 
 *
 * API version: 3.0.x
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package wodby-api

type Domain struct {

	Id string `json:"id"`

	Name string `json:"name"`

	OrgId string `json:"org_id"`

	InstanceId string `json:"instance_id"`

	Enabled bool `json:"enabled,omitempty"`

	Primary bool `json:"primary,omitempty"`

	Protected bool `json:"protected,omitempty"`

	Indexed bool `json:"indexed,omitempty"`

	Ssl bool `json:"ssl,omitempty"`

	SslRequired bool `json:"ssl_required,omitempty"`

	SslCustom bool `json:"ssl_custom,omitempty"`

	Hsts bool `json:"hsts,omitempty"`

	HstsSubdomains bool `json:"hsts_subdomains,omitempty"`

	RedirectToWww bool `json:"redirect_to_www,omitempty"`

	RedirectToNonWww bool `json:"redirect_to_non_www,omitempty"`

	Status string `json:"status"`

	Type_ string `json:"type"`

	Created int32 `json:"created"`

	Updated int32 `json:"updated"`
}
