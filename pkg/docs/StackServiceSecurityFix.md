# StackServiceSecurityFix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VulnerabilityId** | **string** | The CVE ID when there is one, otherwise the advisory&#39;s own ID. | 
**Severity** | **string** |  | 
**PackageName** | **string** |  | 
**PackageVersion** | **string** | The vulnerable version in the image the stack service deploys now. | 
**Exploited** | **bool** | Whether the vulnerability is on CISA&#39;s catalog of known exploited vulnerabilities. | 
**Epss** | Pointer to **float32** | The EPSS probability that the vulnerability is exploited in the next 30 days, from 0 to 1. Absent when unknown. | [optional] 
**Tracked** | **bool** | Whether Wodby tracks the vulnerability itself because public databases don&#39;t list it yet. | 
**Url** | Pointer to **string** | The vulnerability&#39;s public record. Absent when there is none. | [optional] 
**OptionVersions** | **[]string** | The stack service&#39;s option versions whose images have the vulnerability, such as 8.3 and 8.4. | 

## Methods

### NewStackServiceSecurityFix

`func NewStackServiceSecurityFix(vulnerabilityId string, severity string, packageName string, packageVersion string, exploited bool, tracked bool, optionVersions []string, ) *StackServiceSecurityFix`

NewStackServiceSecurityFix instantiates a new StackServiceSecurityFix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackServiceSecurityFixWithDefaults

`func NewStackServiceSecurityFixWithDefaults() *StackServiceSecurityFix`

NewStackServiceSecurityFixWithDefaults instantiates a new StackServiceSecurityFix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVulnerabilityId

`func (o *StackServiceSecurityFix) GetVulnerabilityId() string`

GetVulnerabilityId returns the VulnerabilityId field if non-nil, zero value otherwise.

### GetVulnerabilityIdOk

`func (o *StackServiceSecurityFix) GetVulnerabilityIdOk() (*string, bool)`

GetVulnerabilityIdOk returns a tuple with the VulnerabilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityId

`func (o *StackServiceSecurityFix) SetVulnerabilityId(v string)`

SetVulnerabilityId sets VulnerabilityId field to given value.


### GetSeverity

`func (o *StackServiceSecurityFix) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *StackServiceSecurityFix) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *StackServiceSecurityFix) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetPackageName

`func (o *StackServiceSecurityFix) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *StackServiceSecurityFix) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *StackServiceSecurityFix) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.


### GetPackageVersion

`func (o *StackServiceSecurityFix) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *StackServiceSecurityFix) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *StackServiceSecurityFix) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.


### GetExploited

`func (o *StackServiceSecurityFix) GetExploited() bool`

GetExploited returns the Exploited field if non-nil, zero value otherwise.

### GetExploitedOk

`func (o *StackServiceSecurityFix) GetExploitedOk() (*bool, bool)`

GetExploitedOk returns a tuple with the Exploited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploited

`func (o *StackServiceSecurityFix) SetExploited(v bool)`

SetExploited sets Exploited field to given value.


### GetEpss

`func (o *StackServiceSecurityFix) GetEpss() float32`

GetEpss returns the Epss field if non-nil, zero value otherwise.

### GetEpssOk

`func (o *StackServiceSecurityFix) GetEpssOk() (*float32, bool)`

GetEpssOk returns a tuple with the Epss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpss

`func (o *StackServiceSecurityFix) SetEpss(v float32)`

SetEpss sets Epss field to given value.

### HasEpss

`func (o *StackServiceSecurityFix) HasEpss() bool`

HasEpss returns a boolean if a field has been set.

### GetTracked

`func (o *StackServiceSecurityFix) GetTracked() bool`

GetTracked returns the Tracked field if non-nil, zero value otherwise.

### GetTrackedOk

`func (o *StackServiceSecurityFix) GetTrackedOk() (*bool, bool)`

GetTrackedOk returns a tuple with the Tracked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracked

`func (o *StackServiceSecurityFix) SetTracked(v bool)`

SetTracked sets Tracked field to given value.


### GetUrl

`func (o *StackServiceSecurityFix) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StackServiceSecurityFix) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StackServiceSecurityFix) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StackServiceSecurityFix) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetOptionVersions

`func (o *StackServiceSecurityFix) GetOptionVersions() []string`

GetOptionVersions returns the OptionVersions field if non-nil, zero value otherwise.

### GetOptionVersionsOk

`func (o *StackServiceSecurityFix) GetOptionVersionsOk() (*[]string, bool)`

GetOptionVersionsOk returns a tuple with the OptionVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionVersions

`func (o *StackServiceSecurityFix) SetOptionVersions(v []string)`

SetOptionVersions sets OptionVersions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


