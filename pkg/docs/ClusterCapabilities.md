# ClusterCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvoyGateway** | **bool** | Whether this cluster uses Envoy Gateway for application routing. | 
**RedirectRoutes** | **bool** | Whether this cluster supports routes with the REDIRECT action. | 

## Methods

### NewClusterCapabilities

`func NewClusterCapabilities(envoyGateway bool, redirectRoutes bool, ) *ClusterCapabilities`

NewClusterCapabilities instantiates a new ClusterCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCapabilitiesWithDefaults

`func NewClusterCapabilitiesWithDefaults() *ClusterCapabilities`

NewClusterCapabilitiesWithDefaults instantiates a new ClusterCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvoyGateway

`func (o *ClusterCapabilities) GetEnvoyGateway() bool`

GetEnvoyGateway returns the EnvoyGateway field if non-nil, zero value otherwise.

### GetEnvoyGatewayOk

`func (o *ClusterCapabilities) GetEnvoyGatewayOk() (*bool, bool)`

GetEnvoyGatewayOk returns a tuple with the EnvoyGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvoyGateway

`func (o *ClusterCapabilities) SetEnvoyGateway(v bool)`

SetEnvoyGateway sets EnvoyGateway field to given value.


### GetRedirectRoutes

`func (o *ClusterCapabilities) GetRedirectRoutes() bool`

GetRedirectRoutes returns the RedirectRoutes field if non-nil, zero value otherwise.

### GetRedirectRoutesOk

`func (o *ClusterCapabilities) GetRedirectRoutesOk() (*bool, bool)`

GetRedirectRoutesOk returns a tuple with the RedirectRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectRoutes

`func (o *ClusterCapabilities) SetRedirectRoutes(v bool)`

SetRedirectRoutes sets RedirectRoutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


