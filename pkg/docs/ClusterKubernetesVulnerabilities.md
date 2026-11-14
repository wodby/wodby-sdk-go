# ClusterKubernetesVulnerabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | The upstream patch version the vulnerabilities apply to, such as 1.33.5. | 
**Vulnerabilities** | [**[]ClusterKubernetesVulnerability**](ClusterKubernetesVulnerability.md) |  | 

## Methods

### NewClusterKubernetesVulnerabilities

`func NewClusterKubernetesVulnerabilities(version string, vulnerabilities []ClusterKubernetesVulnerability, ) *ClusterKubernetesVulnerabilities`

NewClusterKubernetesVulnerabilities instantiates a new ClusterKubernetesVulnerabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterKubernetesVulnerabilitiesWithDefaults

`func NewClusterKubernetesVulnerabilitiesWithDefaults() *ClusterKubernetesVulnerabilities`

NewClusterKubernetesVulnerabilitiesWithDefaults instantiates a new ClusterKubernetesVulnerabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ClusterKubernetesVulnerabilities) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClusterKubernetesVulnerabilities) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClusterKubernetesVulnerabilities) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVulnerabilities

`func (o *ClusterKubernetesVulnerabilities) GetVulnerabilities() []ClusterKubernetesVulnerability`

GetVulnerabilities returns the Vulnerabilities field if non-nil, zero value otherwise.

### GetVulnerabilitiesOk

`func (o *ClusterKubernetesVulnerabilities) GetVulnerabilitiesOk() (*[]ClusterKubernetesVulnerability, bool)`

GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilities

`func (o *ClusterKubernetesVulnerabilities) SetVulnerabilities(v []ClusterKubernetesVulnerability)`

SetVulnerabilities sets Vulnerabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


