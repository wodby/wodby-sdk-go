/*
 * Wodby API Reference
 *
 * # Introduction  The Wodby API is organized around REST. Our API has predictable, resource-oriented URLs, and uses HTTP response codes to indicate API errors. We use built-in HTTP features, like HTTP authentication and HTTP verbs, which are understood by off-the-shelf HTTP clients.  JSON is returned by all API responses, including errors.  # Authentication  Authenticate your account by including your secret key in API requests. You can manage your API keys in the Dashboard. Your API keys carry many privileges, so be sure to keep them secure! Do not share your secret API keys in publicly accessible areas such as GitHub, client-side code, and so forth.  All API requests must be made over HTTPS. Calls made over plain HTTP will fail. API requests without authentication will also fail.  Example of authenticated request: ``` curl https://api.wodby.com/api/v3/user -H 'X-API-KEY: YOUR_API_KEY' ``` 
 *
 * API version: 3.0.2
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
