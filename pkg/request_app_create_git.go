/*
 * Wodby API Client
 *
 * Wodby Developer Documentation https://wodby.com/docs/dev
 *
 * API version: 3.0.13
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type RequestAppCreateGit struct {

	RepoId string `json:"repo_id"`

	// Commit, branch or tag
	TreeIsh string `json:"tree_ish,omitempty"`
}
