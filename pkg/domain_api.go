/*
 * Wodby API Client
 *
 * Wodby Developer Documentation https://wodby.com/docs/dev
 *
 * API version: 3.0.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type DomainApiService service


/* DomainApiService 
 Retrieve domain
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id Domain ID
 @return Domain*/
func (a *DomainApiService) GetDomain(ctx context.Context, id string) (Domain,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Domain
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/domains/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-API-KEY"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* DomainApiService 
 Retrieve domains
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "orgId" (string) Organization ID
     @param "instanceId" (string) Instance ID
     @param "serverId" (string) Server ID
     @param "status" (string) Domain status
     @param "type_" (string) Domain type
     @param "name" (string) Domain name
 @return []Domain*/
func (a *DomainApiService) GetDomains(ctx context.Context, localVarOptionals map[string]interface{}) ([]Domain,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []Domain
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/domains"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["status"], "string", "status"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["type_"], "string", "type_"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["name"], "string", "name"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["orgId"].(string); localVarOk {
		localVarQueryParams.Add("org_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["instanceId"].(string); localVarOk {
		localVarQueryParams.Add("instance_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["serverId"].(string); localVarOk {
		localVarQueryParams.Add("server_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["status"].(string); localVarOk {
		localVarQueryParams.Add("status", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["type_"].(string); localVarOk {
		localVarQueryParams.Add("type", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["name"].(string); localVarOk {
		localVarQueryParams.Add("name", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-API-KEY"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

