# KubernetesEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Warning or Normal. | 
**Reason** | **string** |  | 
**Message** | **string** | The event&#39;s message with known secret values masked. | 
**ObjectKind** | **string** | Kind of the object the event is about, such as Pod or PersistentVolumeClaim. | 
**ObjectName** | **string** |  | 
**AppServiceId** | **NullableInt32** | The app service the object belongs to. Null for objects no app service defines, or ones Wodby can&#39;t attribute. | 
**Source** | **NullableString** | The component that reported the event, such as kubelet or default-scheduler. | 
**Count** | **int32** | How many times the event occurred. | 
**FirstSeen** | **time.Time** |  | 
**LastSeen** | **time.Time** |  | 

## Methods

### NewKubernetesEvent

`func NewKubernetesEvent(type_ string, reason string, message string, objectKind string, objectName string, appServiceId NullableInt32, source NullableString, count int32, firstSeen time.Time, lastSeen time.Time, ) *KubernetesEvent`

NewKubernetesEvent instantiates a new KubernetesEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEventWithDefaults

`func NewKubernetesEventWithDefaults() *KubernetesEvent`

NewKubernetesEventWithDefaults instantiates a new KubernetesEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *KubernetesEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubernetesEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KubernetesEvent) SetType(v string)`

SetType sets Type field to given value.


### GetReason

`func (o *KubernetesEvent) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *KubernetesEvent) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *KubernetesEvent) SetReason(v string)`

SetReason sets Reason field to given value.


### GetMessage

`func (o *KubernetesEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KubernetesEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KubernetesEvent) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetObjectKind

`func (o *KubernetesEvent) GetObjectKind() string`

GetObjectKind returns the ObjectKind field if non-nil, zero value otherwise.

### GetObjectKindOk

`func (o *KubernetesEvent) GetObjectKindOk() (*string, bool)`

GetObjectKindOk returns a tuple with the ObjectKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectKind

`func (o *KubernetesEvent) SetObjectKind(v string)`

SetObjectKind sets ObjectKind field to given value.


### GetObjectName

`func (o *KubernetesEvent) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *KubernetesEvent) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *KubernetesEvent) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.


### GetAppServiceId

`func (o *KubernetesEvent) GetAppServiceId() int32`

GetAppServiceId returns the AppServiceId field if non-nil, zero value otherwise.

### GetAppServiceIdOk

`func (o *KubernetesEvent) GetAppServiceIdOk() (*int32, bool)`

GetAppServiceIdOk returns a tuple with the AppServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceId

`func (o *KubernetesEvent) SetAppServiceId(v int32)`

SetAppServiceId sets AppServiceId field to given value.


### SetAppServiceIdNil

`func (o *KubernetesEvent) SetAppServiceIdNil(b bool)`

 SetAppServiceIdNil sets the value for AppServiceId to be an explicit nil

### UnsetAppServiceId
`func (o *KubernetesEvent) UnsetAppServiceId()`

UnsetAppServiceId ensures that no value is present for AppServiceId, not even an explicit nil
### GetSource

`func (o *KubernetesEvent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *KubernetesEvent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *KubernetesEvent) SetSource(v string)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *KubernetesEvent) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *KubernetesEvent) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetCount

`func (o *KubernetesEvent) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *KubernetesEvent) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *KubernetesEvent) SetCount(v int32)`

SetCount sets Count field to given value.


### GetFirstSeen

`func (o *KubernetesEvent) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *KubernetesEvent) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *KubernetesEvent) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.


### GetLastSeen

`func (o *KubernetesEvent) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *KubernetesEvent) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *KubernetesEvent) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


