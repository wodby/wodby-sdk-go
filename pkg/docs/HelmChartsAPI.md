# \HelmChartsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InspectHelmChart**](HelmChartsAPI.md#InspectHelmChart) | **Post** /helm-charts/actions/inspect | Inspect Helm chart



## InspectHelmChart

> HelmChartAnalysis InspectHelmChart(ctx).HelmChartInput(helmChartInput).Execute()

Inspect Helm chart



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wodby/wodby-sdk-go/v4/pkg"
)

func main() {
	helmChartInput := *openapiclient.NewHelmChartInput("Chart_example") // HelmChartInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmChartsAPI.InspectHelmChart(context.Background()).HelmChartInput(helmChartInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmChartsAPI.InspectHelmChart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InspectHelmChart`: HelmChartAnalysis
	fmt.Fprintf(os.Stdout, "Response from `HelmChartsAPI.InspectHelmChart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInspectHelmChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **helmChartInput** | [**HelmChartInput**](HelmChartInput.md) |  | 

### Return type

[**HelmChartAnalysis**](HelmChartAnalysis.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

