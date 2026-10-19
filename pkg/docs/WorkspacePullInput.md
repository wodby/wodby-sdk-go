# WorkspacePullInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scripts** | Pointer to **bool** | Run the environment&#39;s post-deployment scripts after the pull. | [optional] [default to true]

## Methods

### NewWorkspacePullInput

`func NewWorkspacePullInput() *WorkspacePullInput`

NewWorkspacePullInput instantiates a new WorkspacePullInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspacePullInputWithDefaults

`func NewWorkspacePullInputWithDefaults() *WorkspacePullInput`

NewWorkspacePullInputWithDefaults instantiates a new WorkspacePullInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScripts

`func (o *WorkspacePullInput) GetScripts() bool`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *WorkspacePullInput) GetScriptsOk() (*bool, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *WorkspacePullInput) SetScripts(v bool)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *WorkspacePullInput) HasScripts() bool`

HasScripts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


