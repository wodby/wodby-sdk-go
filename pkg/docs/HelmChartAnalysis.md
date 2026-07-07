# HelmChartAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chart** | [**HelmChartMetadata**](HelmChartMetadata.md) |  | 
**Release** | **string** |  | 
**Namespace** | **string** |  | 
**ResourceCount** | **int32** |  | 
**Workloads** | Pointer to [**[]HelmChartWorkload**](HelmChartWorkload.md) |  | [optional] 
**Services** | Pointer to [**[]HelmChartService**](HelmChartService.md) |  | [optional] 
**VolumeClaims** | Pointer to [**[]HelmChartVolumeClaim**](HelmChartVolumeClaim.md) |  | [optional] 
**Crds** | Pointer to [**[]HelmChartResource**](HelmChartResource.md) |  | [optional] 
**ClusterResources** | Pointer to [**[]HelmChartResource**](HelmChartResource.md) |  | [optional] 
**Hooks** | Pointer to [**[]HelmChartResource**](HelmChartResource.md) |  | [optional] 
**UnsupportedKinds** | Pointer to **[]string** |  | [optional] 
**Warnings** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHelmChartAnalysis

`func NewHelmChartAnalysis(chart HelmChartMetadata, release string, namespace string, resourceCount int32, ) *HelmChartAnalysis`

NewHelmChartAnalysis instantiates a new HelmChartAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmChartAnalysisWithDefaults

`func NewHelmChartAnalysisWithDefaults() *HelmChartAnalysis`

NewHelmChartAnalysisWithDefaults instantiates a new HelmChartAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChart

`func (o *HelmChartAnalysis) GetChart() HelmChartMetadata`

GetChart returns the Chart field if non-nil, zero value otherwise.

### GetChartOk

`func (o *HelmChartAnalysis) GetChartOk() (*HelmChartMetadata, bool)`

GetChartOk returns a tuple with the Chart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChart

`func (o *HelmChartAnalysis) SetChart(v HelmChartMetadata)`

SetChart sets Chart field to given value.


### GetRelease

`func (o *HelmChartAnalysis) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *HelmChartAnalysis) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *HelmChartAnalysis) SetRelease(v string)`

SetRelease sets Release field to given value.


### GetNamespace

`func (o *HelmChartAnalysis) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *HelmChartAnalysis) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *HelmChartAnalysis) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetResourceCount

`func (o *HelmChartAnalysis) GetResourceCount() int32`

GetResourceCount returns the ResourceCount field if non-nil, zero value otherwise.

### GetResourceCountOk

`func (o *HelmChartAnalysis) GetResourceCountOk() (*int32, bool)`

GetResourceCountOk returns a tuple with the ResourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCount

`func (o *HelmChartAnalysis) SetResourceCount(v int32)`

SetResourceCount sets ResourceCount field to given value.


### GetWorkloads

`func (o *HelmChartAnalysis) GetWorkloads() []HelmChartWorkload`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *HelmChartAnalysis) GetWorkloadsOk() (*[]HelmChartWorkload, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *HelmChartAnalysis) SetWorkloads(v []HelmChartWorkload)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *HelmChartAnalysis) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.

### GetServices

`func (o *HelmChartAnalysis) GetServices() []HelmChartService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *HelmChartAnalysis) GetServicesOk() (*[]HelmChartService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *HelmChartAnalysis) SetServices(v []HelmChartService)`

SetServices sets Services field to given value.

### HasServices

`func (o *HelmChartAnalysis) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetVolumeClaims

`func (o *HelmChartAnalysis) GetVolumeClaims() []HelmChartVolumeClaim`

GetVolumeClaims returns the VolumeClaims field if non-nil, zero value otherwise.

### GetVolumeClaimsOk

`func (o *HelmChartAnalysis) GetVolumeClaimsOk() (*[]HelmChartVolumeClaim, bool)`

GetVolumeClaimsOk returns a tuple with the VolumeClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeClaims

`func (o *HelmChartAnalysis) SetVolumeClaims(v []HelmChartVolumeClaim)`

SetVolumeClaims sets VolumeClaims field to given value.

### HasVolumeClaims

`func (o *HelmChartAnalysis) HasVolumeClaims() bool`

HasVolumeClaims returns a boolean if a field has been set.

### GetCrds

`func (o *HelmChartAnalysis) GetCrds() []HelmChartResource`

GetCrds returns the Crds field if non-nil, zero value otherwise.

### GetCrdsOk

`func (o *HelmChartAnalysis) GetCrdsOk() (*[]HelmChartResource, bool)`

GetCrdsOk returns a tuple with the Crds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrds

`func (o *HelmChartAnalysis) SetCrds(v []HelmChartResource)`

SetCrds sets Crds field to given value.

### HasCrds

`func (o *HelmChartAnalysis) HasCrds() bool`

HasCrds returns a boolean if a field has been set.

### GetClusterResources

`func (o *HelmChartAnalysis) GetClusterResources() []HelmChartResource`

GetClusterResources returns the ClusterResources field if non-nil, zero value otherwise.

### GetClusterResourcesOk

`func (o *HelmChartAnalysis) GetClusterResourcesOk() (*[]HelmChartResource, bool)`

GetClusterResourcesOk returns a tuple with the ClusterResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterResources

`func (o *HelmChartAnalysis) SetClusterResources(v []HelmChartResource)`

SetClusterResources sets ClusterResources field to given value.

### HasClusterResources

`func (o *HelmChartAnalysis) HasClusterResources() bool`

HasClusterResources returns a boolean if a field has been set.

### GetHooks

`func (o *HelmChartAnalysis) GetHooks() []HelmChartResource`

GetHooks returns the Hooks field if non-nil, zero value otherwise.

### GetHooksOk

`func (o *HelmChartAnalysis) GetHooksOk() (*[]HelmChartResource, bool)`

GetHooksOk returns a tuple with the Hooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooks

`func (o *HelmChartAnalysis) SetHooks(v []HelmChartResource)`

SetHooks sets Hooks field to given value.

### HasHooks

`func (o *HelmChartAnalysis) HasHooks() bool`

HasHooks returns a boolean if a field has been set.

### GetUnsupportedKinds

`func (o *HelmChartAnalysis) GetUnsupportedKinds() []string`

GetUnsupportedKinds returns the UnsupportedKinds field if non-nil, zero value otherwise.

### GetUnsupportedKindsOk

`func (o *HelmChartAnalysis) GetUnsupportedKindsOk() (*[]string, bool)`

GetUnsupportedKindsOk returns a tuple with the UnsupportedKinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsupportedKinds

`func (o *HelmChartAnalysis) SetUnsupportedKinds(v []string)`

SetUnsupportedKinds sets UnsupportedKinds field to given value.

### HasUnsupportedKinds

`func (o *HelmChartAnalysis) HasUnsupportedKinds() bool`

HasUnsupportedKinds returns a boolean if a field has been set.

### GetWarnings

`func (o *HelmChartAnalysis) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *HelmChartAnalysis) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *HelmChartAnalysis) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *HelmChartAnalysis) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


