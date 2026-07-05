# ProblemDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Title** | **string** |  | 
**Status** | **int32** |  | 
**Detail** | **string** |  | 
**Instance** | Pointer to **string** |  | [optional] 
**Code** | **string** |  | 
**Message** | **string** |  | 
**Errors** | Pointer to [**[]ProblemFieldError**](ProblemFieldError.md) |  | [optional] 

## Methods

### NewProblemDetails

`func NewProblemDetails(type_ string, title string, status int32, detail string, code string, message string, ) *ProblemDetails`

NewProblemDetails instantiates a new ProblemDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemDetailsWithDefaults

`func NewProblemDetailsWithDefaults() *ProblemDetails`

NewProblemDetailsWithDefaults instantiates a new ProblemDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProblemDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProblemDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProblemDetails) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *ProblemDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProblemDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProblemDetails) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *ProblemDetails) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProblemDetails) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProblemDetails) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *ProblemDetails) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ProblemDetails) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ProblemDetails) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetInstance

`func (o *ProblemDetails) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ProblemDetails) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ProblemDetails) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ProblemDetails) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetCode

`func (o *ProblemDetails) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProblemDetails) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProblemDetails) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ProblemDetails) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProblemDetails) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProblemDetails) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrors

`func (o *ProblemDetails) GetErrors() []ProblemFieldError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ProblemDetails) GetErrorsOk() (*[]ProblemFieldError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ProblemDetails) SetErrors(v []ProblemFieldError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ProblemDetails) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


