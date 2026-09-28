# WorkspaceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ready** | **bool** |  | 
**Reason** | **string** |  | 
**Host** | **string** |  | 
**Port** | **int32** |  | 
**Username** | **string** |  | 
**WorkingDirectory** | **string** |  | 
**HostKeyFingerprint** | **string** |  | 

## Methods

### NewWorkspaceConnection

`func NewWorkspaceConnection(ready bool, reason string, host string, port int32, username string, workingDirectory string, hostKeyFingerprint string, ) *WorkspaceConnection`

NewWorkspaceConnection instantiates a new WorkspaceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceConnectionWithDefaults

`func NewWorkspaceConnectionWithDefaults() *WorkspaceConnection`

NewWorkspaceConnectionWithDefaults instantiates a new WorkspaceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReady

`func (o *WorkspaceConnection) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *WorkspaceConnection) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *WorkspaceConnection) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetReason

`func (o *WorkspaceConnection) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WorkspaceConnection) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WorkspaceConnection) SetReason(v string)`

SetReason sets Reason field to given value.


### GetHost

`func (o *WorkspaceConnection) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *WorkspaceConnection) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *WorkspaceConnection) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *WorkspaceConnection) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *WorkspaceConnection) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *WorkspaceConnection) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUsername

`func (o *WorkspaceConnection) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *WorkspaceConnection) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *WorkspaceConnection) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetWorkingDirectory

`func (o *WorkspaceConnection) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *WorkspaceConnection) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *WorkspaceConnection) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.


### GetHostKeyFingerprint

`func (o *WorkspaceConnection) GetHostKeyFingerprint() string`

GetHostKeyFingerprint returns the HostKeyFingerprint field if non-nil, zero value otherwise.

### GetHostKeyFingerprintOk

`func (o *WorkspaceConnection) GetHostKeyFingerprintOk() (*string, bool)`

GetHostKeyFingerprintOk returns a tuple with the HostKeyFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostKeyFingerprint

`func (o *WorkspaceConnection) SetHostKeyFingerprint(v string)`

SetHostKeyFingerprint sets HostKeyFingerprint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


