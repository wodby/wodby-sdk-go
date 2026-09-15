package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	client "github.com/wodby/wodby-sdk-go/v3/pkg"
	"golang.org/x/oauth2"
)

// Exercise real request serialization and decoding against a local HTTP server.
func TestClientDependencies(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-API-KEY") != "test-key" {
			t.Error("API key missing")
		}
		switch r.URL.Path {
		case "/orgs":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`[{"id":"org-1","title":"Example"}]`))
		case "/instances/instance-1/deploy":
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Error(err)
			}
			if body["post_deployment"] != true {
				t.Error("JSON body changed")
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"task":{"id":"task-1"}}`))
		default:
			w.WriteHeader(http.StatusForbidden)
		}
	}))
	defer server.Close()
	cfg := client.NewConfiguration()
	cfg.BasePath = server.URL
	api := client.NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), client.ContextAPIKey, client.APIKey{Key: "test-key"})
	orgs, _, err := api.OrganizationApi.GetOrgs(ctx, nil)
	if err != nil || len(orgs) != 1 || orgs[0].Id != "org-1" {
		t.Fatalf("Decode: %v %v", orgs, err)
	}
	result, _, err := api.InstanceApi.DeployInstance(ctx, "instance-1", map[string]interface{}{"data": client.RequestInstanceDeploy{PostDeployment: true}})
	if err != nil || result.Task.Id != "task-1" {
		t.Fatalf("Deploy: %v %v", result, err)
	}
	_, response, err := api.OrganizationApi.GetOrg(ctx, "missing")
	if err == nil || response.StatusCode != http.StatusForbidden {
		t.Fatal("Error status lost")
	}
}

func TestOAuthTransport(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer token" {
			t.Error("OAuth token missing")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[]`))
	}))
	defer server.Close()
	cfg := client.NewConfiguration()
	cfg.BasePath = server.URL
	ctx := context.WithValue(context.Background(), client.ContextOAuth2, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"}))
	if _, _, err := client.NewAPIClient(cfg).OrganizationApi.GetOrgs(ctx, nil); err != nil {
		t.Fatal(err)
	}
}
