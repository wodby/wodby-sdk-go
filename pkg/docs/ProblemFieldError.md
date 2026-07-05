# ProblemFieldError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 

## Methods

### NewProblemFieldError

`func NewProblemFieldError(detail string, ) *ProblemFieldError`

NewProblemFieldError instantiates a new ProblemFieldError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemFieldErrorWithDefaults

`func NewProblemFieldErrorWithDefaults() *ProblemFieldError`

NewProblemFieldErrorWithDefaults instantiates a new ProblemFieldError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *ProblemFieldError) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ProblemFieldError) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ProblemFieldError) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *ProblemFieldError) HasField() bool`

HasField returns a boolean if a field has been set.

### GetCode

`func (o *ProblemFieldError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProblemFieldError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProblemFieldError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ProblemFieldError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetail

`func (o *ProblemFieldError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ProblemFieldError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ProblemFieldError) SetDetail(v string)`

SetDetail sets Detail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


