/*
 * Wodby API (beta)
 *
 * Documentation https://docs.wodby.com/dev GitHub https://github.com/wodby/wodby-api 
 *
 * API version: 3.0.x
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package wodby-api

type RequestAppCreateGit struct {

	RepoId *Uuid `json:"repo_id"`

	// Commit, branch or tag
	TreeIsh string `json:"tree_ish,omitempty"`
}
