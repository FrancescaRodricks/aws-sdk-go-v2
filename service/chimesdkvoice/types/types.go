// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A validated address.
type Address struct {

	// The city of an address.
	City *string

	// The country of an address.
	Country *string

	// An address suffix location, such as the S. Unit A in Central Park S. Unit A .
	PostDirectional *string

	// The postal code of an address.
	PostalCode *string

	// The zip + 4 or postal code + 4 of an address.
	PostalCodePlus4 *string

	// An address prefix location, such as the N in N. Third St.
	PreDirectional *string

	// The state of an address.
	State *string

	// The address street, such as 8th Avenue .
	StreetName *string

	// The numeric portion of an address.
	StreetNumber *string

	// The address suffix, such as the N in 8th Avenue N .
	StreetSuffix *string

	noSmithyDocumentSerde
}

// The details of an Amazon Chime SDK Voice Connector call.
type CallDetails struct {

	// Identifies a person as the caller or the callee.
	IsCaller *bool

	// The transaction ID of a Voice Connector call.
	TransactionId *string

	// The Voice Connector ID.
	VoiceConnectorId *string

	noSmithyDocumentSerde
}

// A suggested address.
type CandidateAddress struct {

	// The city of the candidate address.
	City *string

	// The country of the candidate address.
	Country *string

	// The postal code of the candidate address.
	PostalCode *string

	// The zip + 4 or postal code +4 of the candidate address.
	PostalCodePlus4 *string

	// The state of the candidate address.
	State *string

	// The street information of the candidate address.
	StreetInfo *string

	// The numeric portion of the candidate address.
	StreetNumber *string

	noSmithyDocumentSerde
}

// The SIP credentials used to authenticate requests to an Amazon Chime SDK Voice
// Connector.
type Credential struct {

	// The RFC2617 compliant password associated with the SIP credentials, in US-ASCII
	// format.
	Password *string

	// The RFC2617 compliant user name associated with the SIP credentials, in
	// US-ASCII format.
	Username *string

	noSmithyDocumentSerde
}

// The Dialed Number Identification Service (DNIS) emergency calling configuration
// details associated with an Amazon Chime SDK Voice Connector's emergency calling
// configuration.
type DNISEmergencyCallingConfiguration struct {

	// The country from which emergency calls are allowed, in ISO 3166-1 alpha-2
	// format.
	//
	// This member is required.
	CallingCountry *string

	// The DNIS phone number that you route emergency calls to, in E.164 format.
	//
	// This member is required.
	EmergencyPhoneNumber *string

	// The DNIS phone number for routing test emergency calls to, in E.164 format.
	TestPhoneNumber *string

	noSmithyDocumentSerde
}

// The emergency calling configuration details associated with an Amazon Chime SDK
// Voice Connector.
type EmergencyCallingConfiguration struct {

	// The Dialed Number Identification Service (DNIS) emergency calling configuration
	// details.
	DNIS []DNISEmergencyCallingConfiguration

	noSmithyDocumentSerde
}

// The country and area code for a proxy phone number in a proxy phone session.
type GeoMatchParams struct {

	// The area code.
	//
	// This member is required.
	AreaCode *string

	// The country.
	//
	// This member is required.
	Country *string

	noSmithyDocumentSerde
}

// The logging configuration associated with an Amazon Chime SDK Voice Connector.
// Specifies whether SIP message logs can be sent to Amazon CloudWatch Logs.
type LoggingConfiguration struct {

	// Enables or disables media metrics logging.
	EnableMediaMetricLogs *bool

	// Boolean that enables sending SIP message logs to Amazon CloudWatch.
	EnableSIPLogs *bool

	noSmithyDocumentSerde
}

// The configuration for a call analytics task.
type MediaInsightsConfiguration struct {

	// The configuration's ARN.
	ConfigurationArn *string

	// Denotes the configration as enabled or disabled.
	Disabled *bool

	noSmithyDocumentSerde
}

// A phone number for which an order has been placed.
type OrderedPhoneNumber struct {

	// The phone number, in E.164 format.
	E164PhoneNumber *string

	// The phone number status.
	Status OrderedPhoneNumberStatus

	noSmithyDocumentSerde
}

// Origination settings enable your SIP hosts to receive inbound calls using your
// Amazon Chime SDK Voice Connector.
//
// The parameters listed below are not required, but you must use at least one.
type Origination struct {

	// When origination settings are disabled, inbound calls are not enabled for your
	// Amazon Chime SDK Voice Connector. This parameter is not required, but you must
	// specify this parameter or Routes .
	Disabled *bool

	// The call distribution properties defined for your SIP hosts. Valid range:
	// Minimum value of 1. Maximum value of 20. This parameter is not required, but you
	// must specify this parameter or Disabled .
	Routes []OriginationRoute

	noSmithyDocumentSerde
}

// Origination routes define call distribution properties for your SIP hosts to
// receive inbound calls using an Amazon Chime SDK Voice Connector. Limit: Ten
// origination routes for each Voice Connector.
//
// The parameters listed below are not required, but you must use at least one.
type OriginationRoute struct {

	// The FQDN or IP address to contact for origination traffic.
	Host *string

	// The designated origination route port. Defaults to 5060.
	Port *int32

	// The priority associated with the host, with 1 being the highest priority.
	// Higher priority hosts are attempted first.
	Priority *int32

	// The protocol to use for the origination route. Encryption-enabled Amazon Chime
	// SDK Voice Connectors use TCP protocol by default.
	Protocol OriginationRouteProtocol

	// The weight assigned to an origination route. When hosts have equal priority,
	// calls are distributed between them based on their relative weights.
	Weight *int32

	noSmithyDocumentSerde
}

// The phone number and proxy phone number for a participant in an Amazon Chime
// SDK Voice Connector proxy session.
type Participant struct {

	// The participant's phone number.
	PhoneNumber *string

	// The participant's proxy phone number.
	ProxyPhoneNumber *string

	noSmithyDocumentSerde
}

// A phone number used to call an Amazon Chime SDK Voice Connector.
type PhoneNumber struct {

	// The phone number's associations.
	Associations []PhoneNumberAssociation

	// The outbound calling name associated with the phone number.
	CallingName *string

	// The outbound calling name status.
	CallingNameStatus CallingNameStatus

	// The phone number's capabilities.
	Capabilities *PhoneNumberCapabilities

	// The phone number's country. Format: ISO 3166-1 alpha-2.
	Country *string

	// The phone number creation timestamp, in ISO 8601 format.
	CreatedTimestamp *time.Time

	// The deleted phone number timestamp, in ISO 8601 format.
	DeletionTimestamp *time.Time

	// The phone number, in E.164 format.
	E164PhoneNumber *string

	// The name of the phone number.
	Name *string

	// The phone number's order ID.
	OrderId *string

	// The phone number's ID.
	PhoneNumberId *string

	// The phone number's product type.
	ProductType PhoneNumberProductType

	// The phone number's status.
	Status PhoneNumberStatus

	// The phone number's type.
	Type PhoneNumberType

	// The updated phone number timestamp, in ISO 8601 format.
	UpdatedTimestamp *time.Time

	noSmithyDocumentSerde
}

// The phone number associations, such as an Amazon Chime SDK account ID, user ID,
// Voice Connector ID, or Voice Connector group ID.
type PhoneNumberAssociation struct {

	// The timestamp of the phone number association, in ISO 8601 format.
	AssociatedTimestamp *time.Time

	// Defines the association with an Amazon Chime SDK account ID, user ID, Voice
	// Connector ID, or Voice Connector group ID.
	Name PhoneNumberAssociationName

	// Contains the ID for the entity specified in Name.
	Value *string

	noSmithyDocumentSerde
}

// The phone number capabilities for Amazon Chime SDK phone numbers, such as
// enabled inbound and outbound calling, and text messaging.
type PhoneNumberCapabilities struct {

	// Allows or denies inbound calling for the specified phone number.
	InboundCall *bool

	// Allows or denies inbound MMS messaging for the specified phone number.
	InboundMMS *bool

	// Allows or denies inbound SMS messaging for the specified phone number.
	InboundSMS *bool

	// Allows or denies outbound calling for the specified phone number.
	OutboundCall *bool

	// Allows or denies inbound MMS messaging for the specified phone number.
	OutboundMMS *bool

	// Allows or denies outbound SMS messaging for the specified phone number.
	OutboundSMS *bool

	noSmithyDocumentSerde
}

// The phone number's country.
type PhoneNumberCountry struct {

	// The phone number country code. Format: ISO 3166-1 alpha-2.
	CountryCode *string

	// The supported phone number types.
	SupportedPhoneNumberTypes []PhoneNumberType

	noSmithyDocumentSerde
}

// If a phone number action fails for one or more of the phone numbers in a
// request, a list of the failed phone numbers is returned, along with error codes
// and error messages.
type PhoneNumberError struct {

	// The error code.
	ErrorCode ErrorCode

	// The error message.
	ErrorMessage *string

	// The phone number ID for which the action failed.
	PhoneNumberId *string

	noSmithyDocumentSerde
}

// The details of an Amazon Chime SDK phone number order.
type PhoneNumberOrder struct {

	// The phone number order creation time stamp, in ISO 8601 format.
	CreatedTimestamp *time.Time

	// The type of phone number being ordered, local or toll-free.
	OrderType PhoneNumberOrderType

	// The ordered phone number details, such as the phone number in E.164 format and
	// the phone number status.
	OrderedPhoneNumbers []OrderedPhoneNumber

	// The ID of the phone order.
	PhoneNumberOrderId *string

	// The phone number order product type.
	ProductType PhoneNumberProductType

	// The status of the phone number order.
	Status PhoneNumberOrderStatus

	// The updated phone number order time stamp, in ISO 8601 format.
	UpdatedTimestamp *time.Time

	noSmithyDocumentSerde
}

// The proxy configuration for an Amazon Chime SDK Voice Connector.
type Proxy struct {

	// The default number of minutes allowed for proxy sessions.
	DefaultSessionExpiryMinutes *int32

	// When true, stops proxy sessions from being created on the specified Amazon
	// Chime SDK Voice Connector.
	Disabled *bool

	// The phone number to route calls to after a proxy session expires.
	FallBackPhoneNumber *string

	// The countries for proxy phone numbers to be selected from.
	PhoneNumberCountries []string

	noSmithyDocumentSerde
}

// The proxy session for an Amazon Chime SDK Voice Connector.
type ProxySession struct {

	// The proxy session capabilities.
	Capabilities []Capability

	// The created time stamp, in ISO 8601 format.
	CreatedTimestamp *time.Time

	// The ended time stamp, in ISO 8601 format.
	EndedTimestamp *time.Time

	// The number of minutes allowed for the proxy session.
	ExpiryMinutes *int32

	// The preference for matching the country or area code of the proxy phone number
	// with that of the first participant.
	GeoMatchLevel GeoMatchLevel

	// The country and area code for the proxy phone number.
	GeoMatchParams *GeoMatchParams

	// The proxy session name.
	Name *string

	// The preference for proxy phone number reuse, or stickiness, between the same
	// participants across sessions.
	NumberSelectionBehavior NumberSelectionBehavior

	// The proxy session participants.
	Participants []Participant

	// The proxy session ID.
	ProxySessionId *string

	// The proxy session status.
	Status ProxySessionStatus

	// The updated time stamp, in ISO 8601 format.
	UpdatedTimestamp *time.Time

	// The Voice Connector ID.
	VoiceConnectorId *string

	noSmithyDocumentSerde
}

// A structure that contains the configuration settings for server-side encryption.
//
// We only support symmetric keys. Do not use asymmetric or HMAC keys, or KMS
// aliases.
type ServerSideEncryptionConfiguration struct {

	// The ARN of the KMS key used to encrypt the enrollment data in a voice profile
	// domain. Asymmetric customer managed keys are not supported.
	//
	// This member is required.
	KmsKeyArn *string

	noSmithyDocumentSerde
}

// The details of the SIP media application, including name and endpoints. An AWS
// account can have multiple SIP media applications.
type SipMediaApplication struct {

	// The AWS Region in which the SIP media application is created.
	AwsRegion *string

	// The SIP media application creation timestamp, in ISO 8601 format.
	CreatedTimestamp *time.Time

	// List of endpoints for a SIP media application. Currently, only one endpoint per
	// SIP media application is permitted.
	Endpoints []SipMediaApplicationEndpoint

	// The SIP media application's name.
	Name *string

	// The ARN of the SIP media application.
	SipMediaApplicationArn *string

	// A SIP media application's ID.
	SipMediaApplicationId *string

	// The time at which the SIP media application was updated.
	UpdatedTimestamp *time.Time

	noSmithyDocumentSerde
}

// The Alexa Skill configuration of a SIP media application.
//
// Due to changes made by the Amazon Alexa service, this data type is no longer
// available for use. For more information, refer to the [Alexa Smart Properties]page.
//
// [Alexa Smart Properties]: https://developer.amazon.com/en-US/alexa/alexasmartproperties
type SipMediaApplicationAlexaSkillConfiguration struct {

	// The ID of the Alexa Skill configuration.
	//
	// This member is required.
	AlexaSkillIds []string

	// The status of the Alexa Skill configuration.
	//
	// This member is required.
	AlexaSkillStatus AlexaSkillStatus

	noSmithyDocumentSerde
}

// A Call instance for a SIP media application.
type SipMediaApplicationCall struct {

	// The call's transaction ID.
	TransactionId *string

	noSmithyDocumentSerde
}

// The endpoint assigned to a SIP media application.
type SipMediaApplicationEndpoint struct {

	// Valid Amazon Resource Name (ARN) of the Lambda function, version, or alias. The
	// function must be created in the same AWS Region as the SIP media application.
	LambdaArn *string

	noSmithyDocumentSerde
}

// The logging configuration of a SIP media application.
type SipMediaApplicationLoggingConfiguration struct {

	// Enables message logging for the specified SIP media application.
	EnableSipMediaApplicationMessageLogs *bool

	noSmithyDocumentSerde
}

// The details of a SIP rule, including name, triggers, and target applications.
// An AWS account can have multiple SIP rules.
type SipRule struct {

	// The time at which the SIP rule was created, in ISO 8601 format.
	CreatedTimestamp *time.Time

	// Indicates whether the SIP rule is enabled or disabled. You must disable a rule
	// before you can delete it.
	Disabled *bool

	// A SIP rule's name.
	Name *string

	// A SIP rule's ID.
	SipRuleId *string

	// The target SIP media application and other details, such as priority and AWS
	// Region, to be specified in the SIP rule. Only one SIP rule per AWS Region can be
	// provided.
	TargetApplications []SipRuleTargetApplication

	// The type of trigger set for a SIP rule, either a phone number or a URI request
	// host name.
	TriggerType SipRuleTriggerType

	// The value set for a SIP rule's trigger type. Either a phone number or a URI
	// hostname.
	TriggerValue *string

	// The time at which the SIP rule was updated, in ISO 8601 format.
	UpdatedTimestamp *time.Time

	noSmithyDocumentSerde
}

// A target SIP media application and other details, such as priority and AWS
// Region, to be specified in the SIP rule. Only one SIP rule per AWS Region can be
// provided.
type SipRuleTargetApplication struct {

	// The AWS Region of a rule's target SIP media application.
	AwsRegion *string

	// The priority setting of a rule's target SIP media application.
	Priority *int32

	// The ID of a rule's target SIP media application.
	SipMediaApplicationId *string

	noSmithyDocumentSerde
}

// The details of a speaker search task.
type SpeakerSearchDetails struct {

	// The result value in the speaker search details.
	Results []SpeakerSearchResult

	// The status of a voice print generation operation, VoiceprintGenerationSuccess
	// or VoiceprintGenerationFailure ..
	VoiceprintGenerationStatus *string

	noSmithyDocumentSerde
}

// The result of a speaker search analysis.
type SpeakerSearchResult struct {

	// The confidence score in the speaker search analysis.
	ConfidenceScore float32

	// The voice profile ID.
	VoiceProfileId *string

	noSmithyDocumentSerde
}

// A representation of an asynchronous request to perform speaker search analysis
// on a Voice Connector call.
type SpeakerSearchTask struct {

	// The call details of a speaker search task.
	CallDetails *CallDetails

	// The time at which a speaker search task was created.
	CreatedTimestamp *time.Time

	// The details of a speaker search task.
	SpeakerSearchDetails *SpeakerSearchDetails

	// The speaker search task ID.
	SpeakerSearchTaskId *string

	// The status of the speaker search task, IN_QUEUE , IN_PROGRESS , PARTIAL_SUCCESS
	// , SUCCEEDED , FAILED , or STOPPED .
	SpeakerSearchTaskStatus *string

	// The time at which the speaker search task began.
	StartedTimestamp *time.Time

	// A detailed message about the status of a speaker search.
	StatusMessage *string

	// The time at which a speaker search task was updated.
	UpdatedTimestamp *time.Time

	noSmithyDocumentSerde
}

// The streaming configuration associated with an Amazon Chime SDK Voice
// Connector. Specifies whether media streaming is enabled for sending to Amazon
// Kinesis, and shows the retention period for the Amazon Kinesis data, in hours.
type StreamingConfiguration struct {

	// The amount of time, in hours, to the Kinesis data.
	//
	// This member is required.
	DataRetentionInHours *int32

	// When true, streaming to Kinesis is off.
	//
	// This member is required.
	Disabled *bool

	// The call analytics configuration.
	MediaInsightsConfiguration *MediaInsightsConfiguration

	// The streaming notification targets.
	StreamingNotificationTargets []StreamingNotificationTarget

	noSmithyDocumentSerde
}

// The target recipient for a streaming configuration notification.
type StreamingNotificationTarget struct {

	// The streaming notification target.
	NotificationTarget NotificationTarget

	noSmithyDocumentSerde
}

// Describes a tag applied to a resource.
type Tag struct {

	// The tag's key.
	//
	// This member is required.
	Key *string

	// The tag's value.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// Termination settings enable SIP hosts to make outbound calls using an Amazon
// Chime SDK Voice Connector.
type Termination struct {

	// The countries to which calls are allowed, in ISO 3166-1 alpha-2 format.
	// Required.
	CallingRegions []string

	// The IP addresses allowed to make calls, in CIDR format.
	CidrAllowedList []string

	// The limit on calls per second. Max value based on account service quota.
	// Default value of 1.
	CpsLimit *int32

	// The default outbound calling number.
	DefaultPhoneNumber *string

	// When termination is disabled, outbound calls cannot be made.
	Disabled *bool

	noSmithyDocumentSerde
}

// The termination health details, including the source IP address and timestamp
// of the last successful SIP OPTIONS message from your SIP infrastructure.
type TerminationHealth struct {

	// The source IP address.
	Source *string

	// The timestamp, in ISO 8601 format.
	Timestamp *time.Time

	noSmithyDocumentSerde
}

// The phone number ID, product type, or calling name fields to update, used with
// the BatchUpdatePhoneNumberand UpdatePhoneNumber actions.
type UpdatePhoneNumberRequestItem struct {

	// The phone number ID to update.
	//
	// This member is required.
	PhoneNumberId *string

	// The outbound calling name to update.
	CallingName *string

	// The name of the phone number.
	Name *string

	// The product type to update.
	ProductType PhoneNumberProductType

	noSmithyDocumentSerde
}

// The Amazon Chime SDK Voice Connector configuration, including outbound host
// name and encryption settings.
type VoiceConnector struct {

	// The AWS Region in which the Voice Connector is created. Default: us-east-1.
	AwsRegion VoiceConnectorAwsRegion

	// The Voice Connector's creation timestamp, in ISO 8601 format.
	CreatedTimestamp *time.Time

	// The Voice Connector's name.
	Name *string

	// The outbound host name for the Voice Connector.
	OutboundHostName *string

	// Enables or disables encryption for the Voice Connector.
	RequireEncryption *bool

	// The Voice Connector's updated timestamp, in ISO 8601 format.
	UpdatedTimestamp *time.Time

	// The ARN of the Voice Connector.
	VoiceConnectorArn *string

	// The Voice Connector's ID.
	VoiceConnectorId *string

	noSmithyDocumentSerde
}

// The Amazon Chime SDK Voice Connector group configuration, including associated
// Voice Connectors. You can include Voice Connectors from different AWS Regions in
// a group. This creates a fault tolerant mechanism for fallback in case of
// availability events.
type VoiceConnectorGroup struct {

	// The Voice Connector group's creation time stamp, in ISO 8601 format.
	CreatedTimestamp *time.Time

	// The name of a Voice Connector group.
	Name *string

	// The Voice Connector group's creation time stamp, in ISO 8601 format.
	UpdatedTimestamp *time.Time

	// The ARN of the Voice Connector group.
	VoiceConnectorGroupArn *string

	// The ID of a Voice Connector group.
	VoiceConnectorGroupId *string

	// The Voice Connectors to which you route inbound calls.
	VoiceConnectorItems []VoiceConnectorItem

	noSmithyDocumentSerde
}

// For Amazon Chime SDK Voice Connector groups, the Amazon Chime SDK Voice
// Connectors to which you route inbound calls. Includes priority configuration
// settings. Limit: 3 VoiceConnectorItems per Voice Connector group.
type VoiceConnectorItem struct {

	// The priority setting of a Voice Connector item. Calls are routed to hosts in
	// priority order, with 1 as the highest priority. When hosts have equal priority,
	// the system distributes calls among them based on their relative weight.
	//
	// This member is required.
	Priority *int32

	// The Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string

	noSmithyDocumentSerde
}

// The Amazon Chime SDK Voice Connector settings. Includes any Amazon S3 buckets
// designated for storing call detail records.
type VoiceConnectorSettings struct {

	// The S3 bucket that stores the Voice Connector's call detail records.
	CdrBucket *string

	noSmithyDocumentSerde
}

// The combination of a voice print and caller ID.
type VoiceProfile struct {

	// The time at which the voice profile was created and enrolled.
	CreatedTimestamp *time.Time

	// The time at which a voice profile expires unless you re-enroll the caller via
	// the UpdateVoiceProfile API.
	ExpirationTimestamp *time.Time

	// The time at which the voice profile was last updated.
	UpdatedTimestamp *time.Time

	// The ARN of the voice profile.
	VoiceProfileArn *string

	// The ID of the domain that contains the voice profile.
	VoiceProfileDomainId *string

	// The ID of the voice profile.
	VoiceProfileId *string

	noSmithyDocumentSerde
}

// A collection of voice profiles.
type VoiceProfileDomain struct {

	// The time at which the voice profile domain was created.
	CreatedTimestamp *time.Time

	// The description of the voice profile domain.
	Description *string

	// The name of the voice profile domain.
	Name *string

	// A structure that contains the configuration settings for server-side encryption.
	ServerSideEncryptionConfiguration *ServerSideEncryptionConfiguration

	// The time at which the voice profile was last updated.
	UpdatedTimestamp *time.Time

	// The voice profile domain's Amazon Resource Number (ARN).
	VoiceProfileDomainArn *string

	// The ID of the voice profile domain.
	VoiceProfileDomainId *string

	noSmithyDocumentSerde
}

// A high-level overview of a voice profile domain.
type VoiceProfileDomainSummary struct {

	// The time at which the voice profile domain summary was created.
	CreatedTimestamp *time.Time

	// Describes the voice profile domain summary.
	Description *string

	// The name of the voice profile domain summary.
	Name *string

	// The time at which the voice profile domain summary was last updated.
	UpdatedTimestamp *time.Time

	// The ARN of a voice profile in a voice profile domain summary.
	VoiceProfileDomainArn *string

	// The ID of the voice profile domain summary.
	VoiceProfileDomainId *string

	noSmithyDocumentSerde
}

// A high-level summary of a voice profile.
type VoiceProfileSummary struct {

	// The time at which a voice profile summary was created.
	CreatedTimestamp *time.Time

	// Extends the life of the voice profile. You can use UpdateVoiceProfile to
	// refresh an existing voice profile's voice print and extend the life of the
	// summary.
	ExpirationTimestamp *time.Time

	// The time at which a voice profile summary was last updated.
	UpdatedTimestamp *time.Time

	// The ARN of the voice profile in a voice profile summary.
	VoiceProfileArn *string

	// The ID of the voice profile domain in a voice profile summary.
	VoiceProfileDomainId *string

	// The ID of the voice profile in a voice profile summary.
	VoiceProfileId *string

	noSmithyDocumentSerde
}

// A representation of an asynchronous request to perform voice tone analysis on a
// Voice Connector call.
type VoiceToneAnalysisTask struct {

	// The call details of a voice tone analysis task.
	CallDetails *CallDetails

	// The time at which a voice tone analysis task was created.
	CreatedTimestamp *time.Time

	// The time at which a voice tone analysis task started.
	StartedTimestamp *time.Time

	// The status of a voice tone analysis task.
	StatusMessage *string

	// The time at which a voice tone analysis task was updated.
	UpdatedTimestamp *time.Time

	// The ID of the voice tone analysis task.
	VoiceToneAnalysisTaskId *string

	// The status of a voice tone analysis task, IN_QUEUE , IN_PROGRESS ,
	// PARTIAL_SUCCESS , SUCCEEDED , FAILED , or STOPPED .
	VoiceToneAnalysisTaskStatus *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
