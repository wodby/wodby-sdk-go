# ServiceManifest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Raw** | **string** |  | 
**Scalable** | **bool** |  | 
**BasicAuth** | **bool** | Whether the service asks for its public HTTP endpoints to be protected with basic auth by default. | 
**Labels** | Pointer to **[]string** |  | [optional] 
**CrossApp** | **bool** | Whether links of services in other app environments, of the same app or another one, may point at this service. | 
**Exports** | [**[]ServiceManifestExport**](ServiceManifestExport.md) | Values the services linking to this one read as {{link.exports.&lt;name&gt;}}. | 
**Integrations** | Pointer to [**[]ServiceIntegrationRequirement**](ServiceIntegrationRequirement.md) |  | [optional] 
**Links** | Pointer to [**[]ServiceManifestLink**](ServiceManifestLink.md) |  | [optional] 
**LinkGroups** | Pointer to [**[]ServiceManifestLinkGroup**](ServiceManifestLinkGroup.md) | Links that are alternatives to one another. An enabled app service needs at least min of them configured. | [optional] 

## Methods

### NewServiceManifest

`func NewServiceManifest(raw string, scalable bool, basicAuth bool, crossApp bool, exports []ServiceManifestExport, ) *ServiceManifest`

NewServiceManifest instantiates a new ServiceManifest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceManifestWithDefaults

`func NewServiceManifestWithDefaults() *ServiceManifest`

NewServiceManifestWithDefaults instantiates a new ServiceManifest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRaw

`func (o *ServiceManifest) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *ServiceManifest) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *ServiceManifest) SetRaw(v string)`

SetRaw sets Raw field to given value.


### GetScalable

`func (o *ServiceManifest) GetScalable() bool`

GetScalable returns the Scalable field if non-nil, zero value otherwise.

### GetScalableOk

`func (o *ServiceManifest) GetScalableOk() (*bool, bool)`

GetScalableOk returns a tuple with the Scalable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalable

`func (o *ServiceManifest) SetScalable(v bool)`

SetScalable sets Scalable field to given value.


### GetBasicAuth

`func (o *ServiceManifest) GetBasicAuth() bool`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *ServiceManifest) GetBasicAuthOk() (*bool, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *ServiceManifest) SetBasicAuth(v bool)`

SetBasicAuth sets BasicAuth field to given value.


### GetLabels

`func (o *ServiceManifest) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ServiceManifest) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ServiceManifest) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ServiceManifest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCrossApp

`func (o *ServiceManifest) GetCrossApp() bool`

GetCrossApp returns the CrossApp field if non-nil, zero value otherwise.

### GetCrossAppOk

`func (o *ServiceManifest) GetCrossAppOk() (*bool, bool)`

GetCrossAppOk returns a tuple with the CrossApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossApp

`func (o *ServiceManifest) SetCrossApp(v bool)`

SetCrossApp sets CrossApp field to given value.


### GetExports

`func (o *ServiceManifest) GetExports() []ServiceManifestExport`

GetExports returns the Exports field if non-nil, zero value otherwise.

### GetExportsOk

`func (o *ServiceManifest) GetExportsOk() (*[]ServiceManifestExport, bool)`

GetExportsOk returns a tuple with the Exports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExports

`func (o *ServiceManifest) SetExports(v []ServiceManifestExport)`

SetExports sets Exports field to given value.


### GetIntegrations

`func (o *ServiceManifest) GetIntegrations() []ServiceIntegrationRequirement`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *ServiceManifest) GetIntegrationsOk() (*[]ServiceIntegrationRequirement, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *ServiceManifest) SetIntegrations(v []ServiceIntegrationRequirement)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *ServiceManifest) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetLinks

`func (o *ServiceManifest) GetLinks() []ServiceManifestLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServiceManifest) GetLinksOk() (*[]ServiceManifestLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServiceManifest) SetLinks(v []ServiceManifestLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServiceManifest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetLinkGroups

`func (o *ServiceManifest) GetLinkGroups() []ServiceManifestLinkGroup`

GetLinkGroups returns the LinkGroups field if non-nil, zero value otherwise.

### GetLinkGroupsOk

`func (o *ServiceManifest) GetLinkGroupsOk() (*[]ServiceManifestLinkGroup, bool)`

GetLinkGroupsOk returns a tuple with the LinkGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkGroups

`func (o *ServiceManifest) SetLinkGroups(v []ServiceManifestLinkGroup)`

SetLinkGroups sets LinkGroups field to given value.

### HasLinkGroups

`func (o *ServiceManifest) HasLinkGroups() bool`

HasLinkGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


