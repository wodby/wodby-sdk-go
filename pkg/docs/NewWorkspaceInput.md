# NewWorkspaceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** |  | 
**StorageClassName** | Pointer to **string** |  | [optional] 
**StorageServiceName** | Pointer to **string** |  | [optional] 
**CodeSize** | Pointer to **int32** |  | [optional] [default to 10]
**HomeSize** | Pointer to **int32** |  | [optional] [default to 5]

## Methods

### NewNewWorkspaceInput

`func NewNewWorkspaceInput(branch string, ) *NewWorkspaceInput`

NewNewWorkspaceInput instantiates a new NewWorkspaceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewWorkspaceInputWithDefaults

`func NewNewWorkspaceInputWithDefaults() *NewWorkspaceInput`

NewNewWorkspaceInputWithDefaults instantiates a new NewWorkspaceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *NewWorkspaceInput) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *NewWorkspaceInput) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *NewWorkspaceInput) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetStorageClassName

`func (o *NewWorkspaceInput) GetStorageClassName() string`

GetStorageClassName returns the StorageClassName field if non-nil, zero value otherwise.

### GetStorageClassNameOk

`func (o *NewWorkspaceInput) GetStorageClassNameOk() (*string, bool)`

GetStorageClassNameOk returns a tuple with the StorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassName

`func (o *NewWorkspaceInput) SetStorageClassName(v string)`

SetStorageClassName sets StorageClassName field to given value.

### HasStorageClassName

`func (o *NewWorkspaceInput) HasStorageClassName() bool`

HasStorageClassName returns a boolean if a field has been set.

### GetStorageServiceName

`func (o *NewWorkspaceInput) GetStorageServiceName() string`

GetStorageServiceName returns the StorageServiceName field if non-nil, zero value otherwise.

### GetStorageServiceNameOk

`func (o *NewWorkspaceInput) GetStorageServiceNameOk() (*string, bool)`

GetStorageServiceNameOk returns a tuple with the StorageServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServiceName

`func (o *NewWorkspaceInput) SetStorageServiceName(v string)`

SetStorageServiceName sets StorageServiceName field to given value.

### HasStorageServiceName

`func (o *NewWorkspaceInput) HasStorageServiceName() bool`

HasStorageServiceName returns a boolean if a field has been set.

### GetCodeSize

`func (o *NewWorkspaceInput) GetCodeSize() int32`

GetCodeSize returns the CodeSize field if non-nil, zero value otherwise.

### GetCodeSizeOk

`func (o *NewWorkspaceInput) GetCodeSizeOk() (*int32, bool)`

GetCodeSizeOk returns a tuple with the CodeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSize

`func (o *NewWorkspaceInput) SetCodeSize(v int32)`

SetCodeSize sets CodeSize field to given value.

### HasCodeSize

`func (o *NewWorkspaceInput) HasCodeSize() bool`

HasCodeSize returns a boolean if a field has been set.

### GetHomeSize

`func (o *NewWorkspaceInput) GetHomeSize() int32`

GetHomeSize returns the HomeSize field if non-nil, zero value otherwise.

### GetHomeSizeOk

`func (o *NewWorkspaceInput) GetHomeSizeOk() (*int32, bool)`

GetHomeSizeOk returns a tuple with the HomeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeSize

`func (o *NewWorkspaceInput) SetHomeSize(v int32)`

SetHomeSize sets HomeSize field to given value.

### HasHomeSize

`func (o *NewWorkspaceInput) HasHomeSize() bool`

HasHomeSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


