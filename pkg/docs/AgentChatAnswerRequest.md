# AgentChatAnswerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **int32** | The ask_user message that waits for an answer. | 
**Answers** | **[]string** |  | 

## Methods

### NewAgentChatAnswerRequest

`func NewAgentChatAnswerRequest(messageId int32, answers []string, ) *AgentChatAnswerRequest`

NewAgentChatAnswerRequest instantiates a new AgentChatAnswerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentChatAnswerRequestWithDefaults

`func NewAgentChatAnswerRequestWithDefaults() *AgentChatAnswerRequest`

NewAgentChatAnswerRequestWithDefaults instantiates a new AgentChatAnswerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *AgentChatAnswerRequest) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *AgentChatAnswerRequest) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *AgentChatAnswerRequest) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.


### GetAnswers

`func (o *AgentChatAnswerRequest) GetAnswers() []string`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *AgentChatAnswerRequest) GetAnswersOk() (*[]string, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *AgentChatAnswerRequest) SetAnswers(v []string)`

SetAnswers sets Answers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


