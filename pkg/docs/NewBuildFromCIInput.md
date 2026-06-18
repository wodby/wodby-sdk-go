# NewBuildFromCIInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppServiceId** | **int32** |  | 
**GitCommitSHA** | **string** |  | 
**GitRef** | **string** |  | 
**GitRefType** | **string** |  | 
**BuildNum** | **int32** |  | 
**BuildId** | **string** |  | 
**Workflow** | Pointer to **NullableString** |  | [optional] 
**GitCommitAuthorName** | Pointer to **NullableString** |  | [optional] 
**GitCommitAuthorEmail** | Pointer to **NullableString** |  | [optional] 
**GitCommitMessage** | Pointer to **NullableString** |  | [optional] 
**Provider** | **string** |  | 
**PostDeployment** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNewBuildFromCIInput

`func NewNewBuildFromCIInput(appServiceId int32, gitCommitSHA string, gitRef string, gitRefType string, buildNum int32, buildId string, provider string, ) *NewBuildFromCIInput`

NewNewBuildFromCIInput instantiates a new NewBuildFromCIInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewBuildFromCIInputWithDefaults

`func NewNewBuildFromCIInputWithDefaults() *NewBuildFromCIInput`

NewNewBuildFromCIInputWithDefaults instantiates a new NewBuildFromCIInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppServiceId

`func (o *NewBuildFromCIInput) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *NewBuildFromCIInput) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *NewBuildFromCIInput) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### GetGitCommitSHA

`func (o *NewBuildFromCIInput) GetGitCommitSHA() string`

GetGitCommitSHA returns the GitCommitSHA field if non-nil, zero value otherwise.

### GetGitCommitSHAOk

`func (o *NewBuildFromCIInput) GetGitCommitSHAOk() (*string, bool)`

GetGitCommitSHAOk returns a tuple with the GitCommitSHA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitSHA

`func (o *NewBuildFromCIInput) SetGitCommitSHA(v string)`

SetGitCommitSHA sets GitCommitSHA field to given value.


### GetGitRef

`func (o *NewBuildFromCIInput) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *NewBuildFromCIInput) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *NewBuildFromCIInput) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.


### GetGitRefType

`func (o *NewBuildFromCIInput) GetGitRefType() string`

GetGitRefType returns the GitRefType field if non-nil, zero value otherwise.

### GetGitRefTypeOk

`func (o *NewBuildFromCIInput) GetGitRefTypeOk() (*string, bool)`

GetGitRefTypeOk returns a tuple with the GitRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRefType

`func (o *NewBuildFromCIInput) SetGitRefType(v string)`

SetGitRefType sets GitRefType field to given value.


### GetBuildNum

`func (o *NewBuildFromCIInput) GetBuildNum() int32`

GetBuildNum returns the BuildNum field if non-nil, zero value otherwise.

### GetBuildNumOk

`func (o *NewBuildFromCIInput) GetBuildNumOk() (*int32, bool)`

GetBuildNumOk returns a tuple with the BuildNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNum

`func (o *NewBuildFromCIInput) SetBuildNum(v int32)`

SetBuildNum sets BuildNum field to given value.


### GetBuildId

`func (o *NewBuildFromCIInput) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *NewBuildFromCIInput) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *NewBuildFromCIInput) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.


### GetWorkflow

`func (o *NewBuildFromCIInput) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *NewBuildFromCIInput) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *NewBuildFromCIInput) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *NewBuildFromCIInput) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### SetWorkflowNil

`func (o *NewBuildFromCIInput) SetWorkflowNil(b bool)`

 SetWorkflowNil sets the value for Workflow to be an explicit nil

### UnsetWorkflow
`func (o *NewBuildFromCIInput) UnsetWorkflow()`

UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil
### GetGitCommitAuthorName

`func (o *NewBuildFromCIInput) GetGitCommitAuthorName() string`

GetGitCommitAuthorName returns the GitCommitAuthorName field if non-nil, zero value otherwise.

### GetGitCommitAuthorNameOk

`func (o *NewBuildFromCIInput) GetGitCommitAuthorNameOk() (*string, bool)`

GetGitCommitAuthorNameOk returns a tuple with the GitCommitAuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitAuthorName

`func (o *NewBuildFromCIInput) SetGitCommitAuthorName(v string)`

SetGitCommitAuthorName sets GitCommitAuthorName field to given value.

### HasGitCommitAuthorName

`func (o *NewBuildFromCIInput) HasGitCommitAuthorName() bool`

HasGitCommitAuthorName returns a boolean if a field has been set.

### SetGitCommitAuthorNameNil

`func (o *NewBuildFromCIInput) SetGitCommitAuthorNameNil(b bool)`

 SetGitCommitAuthorNameNil sets the value for GitCommitAuthorName to be an explicit nil

### UnsetGitCommitAuthorName
`func (o *NewBuildFromCIInput) UnsetGitCommitAuthorName()`

UnsetGitCommitAuthorName ensures that no value is present for GitCommitAuthorName, not even an explicit nil
### GetGitCommitAuthorEmail

`func (o *NewBuildFromCIInput) GetGitCommitAuthorEmail() string`

GetGitCommitAuthorEmail returns the GitCommitAuthorEmail field if non-nil, zero value otherwise.

### GetGitCommitAuthorEmailOk

`func (o *NewBuildFromCIInput) GetGitCommitAuthorEmailOk() (*string, bool)`

GetGitCommitAuthorEmailOk returns a tuple with the GitCommitAuthorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitAuthorEmail

`func (o *NewBuildFromCIInput) SetGitCommitAuthorEmail(v string)`

SetGitCommitAuthorEmail sets GitCommitAuthorEmail field to given value.

### HasGitCommitAuthorEmail

`func (o *NewBuildFromCIInput) HasGitCommitAuthorEmail() bool`

HasGitCommitAuthorEmail returns a boolean if a field has been set.

### SetGitCommitAuthorEmailNil

`func (o *NewBuildFromCIInput) SetGitCommitAuthorEmailNil(b bool)`

 SetGitCommitAuthorEmailNil sets the value for GitCommitAuthorEmail to be an explicit nil

### UnsetGitCommitAuthorEmail
`func (o *NewBuildFromCIInput) UnsetGitCommitAuthorEmail()`

UnsetGitCommitAuthorEmail ensures that no value is present for GitCommitAuthorEmail, not even an explicit nil
### GetGitCommitMessage

`func (o *NewBuildFromCIInput) GetGitCommitMessage() string`

GetGitCommitMessage returns the GitCommitMessage field if non-nil, zero value otherwise.

### GetGitCommitMessageOk

`func (o *NewBuildFromCIInput) GetGitCommitMessageOk() (*string, bool)`

GetGitCommitMessageOk returns a tuple with the GitCommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitMessage

`func (o *NewBuildFromCIInput) SetGitCommitMessage(v string)`

SetGitCommitMessage sets GitCommitMessage field to given value.

### HasGitCommitMessage

`func (o *NewBuildFromCIInput) HasGitCommitMessage() bool`

HasGitCommitMessage returns a boolean if a field has been set.

### SetGitCommitMessageNil

`func (o *NewBuildFromCIInput) SetGitCommitMessageNil(b bool)`

 SetGitCommitMessageNil sets the value for GitCommitMessage to be an explicit nil

### UnsetGitCommitMessage
`func (o *NewBuildFromCIInput) UnsetGitCommitMessage()`

UnsetGitCommitMessage ensures that no value is present for GitCommitMessage, not even an explicit nil
### GetProvider

`func (o *NewBuildFromCIInput) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NewBuildFromCIInput) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NewBuildFromCIInput) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetPostDeployment

`func (o *NewBuildFromCIInput) GetPostDeployment() string`

GetPostDeployment returns the PostDeployment field if non-nil, zero value otherwise.

### GetPostDeploymentOk

`func (o *NewBuildFromCIInput) GetPostDeploymentOk() (*string, bool)`

GetPostDeploymentOk returns a tuple with the PostDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostDeployment

`func (o *NewBuildFromCIInput) SetPostDeployment(v string)`

SetPostDeployment sets PostDeployment field to given value.

### HasPostDeployment

`func (o *NewBuildFromCIInput) HasPostDeployment() bool`

HasPostDeployment returns a boolean if a field has been set.

### SetPostDeploymentNil

`func (o *NewBuildFromCIInput) SetPostDeploymentNil(b bool)`

 SetPostDeploymentNil sets the value for PostDeployment to be an explicit nil

### UnsetPostDeployment
`func (o *NewBuildFromCIInput) UnsetPostDeployment()`

UnsetPostDeployment ensures that no value is present for PostDeployment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


