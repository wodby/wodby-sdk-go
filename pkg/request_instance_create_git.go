/*
 * Wodby API
 *
 * API documentation https://docs.wodby.com/dev.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type RequestInstanceCreateGit struct {

	// Commit, branch or tag
	TreeIsh string `json:"tree_ish,omitempty"`
}
