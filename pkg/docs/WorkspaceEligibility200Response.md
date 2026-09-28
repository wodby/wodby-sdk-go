# WorkspaceEligibility200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eligible** | Pointer to **bool** |  | [optional] 
**Reasons** | Pointer to **[]string** |  | [optional] 
**SourceStackServiceId** | Pointer to **NullableInt32** |  | [optional] 
**ConsumerStackServiceIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewWorkspaceEligibility200Response

`func NewWorkspaceEligibility200Response() *WorkspaceEligibility200Response`

NewWorkspaceEligibility200Response instantiates a new WorkspaceEligibility200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceEligibility200ResponseWithDefaults

`func NewWorkspaceEligibility200ResponseWithDefaults() *WorkspaceEligibility200Response`

NewWorkspaceEligibility200ResponseWithDefaults instantiates a new WorkspaceEligibility200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEligible

`func (o *WorkspaceEligibility200Response) GetEligible() bool`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *WorkspaceEligibility200Response) GetEligibleOk() (*bool, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *WorkspaceEligibility200Response) SetEligible(v bool)`

SetEligible sets Eligible field to given value.

### HasEligible

`func (o *WorkspaceEligibility200Response) HasEligible() bool`

HasEligible returns a boolean if a field has been set.

### GetReasons

`func (o *WorkspaceEligibility200Response) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *WorkspaceEligibility200Response) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *WorkspaceEligibility200Response) SetReasons(v []string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *WorkspaceEligibility200Response) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetSourceStackServiceId

`func (o *WorkspaceEligibility200Response) GetSourceStackServiceId() int32`

GetSourceStackServiceId returns the SourceStackServiceId field if non-nil, zero value otherwise.

### GetSourceStackServiceIdOk

`func (o *WorkspaceEligibility200Response) GetSourceStackServiceIdOk() (*int32, bool)`

GetSourceStackServiceIdOk returns a tuple with the SourceStackServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStackServiceId

`func (o *WorkspaceEligibility200Response) SetSourceStackServiceId(v int32)`

SetSourceStackServiceId sets SourceStackServiceId field to given value.

### HasSourceStackServiceId

`func (o *WorkspaceEligibility200Response) HasSourceStackServiceId() bool`

HasSourceStackServiceId returns a boolean if a field has been set.

### SetSourceStackServiceIdNil

`func (o *WorkspaceEligibility200Response) SetSourceStackServiceIdNil(b bool)`

 SetSourceStackServiceIdNil sets the value for SourceStackServiceId to be an explicit nil

### UnsetSourceStackServiceId
`func (o *WorkspaceEligibility200Response) UnsetSourceStackServiceId()`

UnsetSourceStackServiceId ensures that no value is present for SourceStackServiceId, not even an explicit nil
### GetConsumerStackServiceIds

`func (o *WorkspaceEligibility200Response) GetConsumerStackServiceIds() []int32`

GetConsumerStackServiceIds returns the ConsumerStackServiceIds field if non-nil, zero value otherwise.

### GetConsumerStackServiceIdsOk

`func (o *WorkspaceEligibility200Response) GetConsumerStackServiceIdsOk() (*[]int32, bool)`

GetConsumerStackServiceIdsOk returns a tuple with the ConsumerStackServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerStackServiceIds

`func (o *WorkspaceEligibility200Response) SetConsumerStackServiceIds(v []int32)`

SetConsumerStackServiceIds sets ConsumerStackServiceIds field to given value.

### HasConsumerStackServiceIds

`func (o *WorkspaceEligibility200Response) HasConsumerStackServiceIds() bool`

HasConsumerStackServiceIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


