/*
 * Wodby API
 *
 * Developers knowledge base https://wodby.com/docs/dev.
 *
 * API version: 3.x
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type RequestAppCreateGit struct {

	RepoId *Uuid `json:"repo_id"`

	// Commit, branch or tag
	TreeIsh string `json:"tree_ish,omitempty"`
}
