// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sends a message to an Amazon SNS topic, a text message (SMS message) directly
// to a phone number, or a message to a mobile platform endpoint (when you specify
// the TargetArn ).
//
// If you send a message to a topic, Amazon SNS delivers the message to each
// endpoint that is subscribed to the topic. The format of the message depends on
// the notification protocol for each subscribed endpoint.
//
// When a messageId is returned, the message is saved and Amazon SNS immediately
// delivers it to subscribers.
//
// To use the Publish action for publishing a message to a mobile endpoint, such
// as an app on a Kindle device or mobile phone, you must specify the EndpointArn
// for the TargetArn parameter. The EndpointArn is returned when making a call with
// the CreatePlatformEndpoint action.
//
// For more information about formatting messages, see [Send Custom Platform-Specific Payloads in Messages to Mobile Devices].
//
// You can publish messages only to topics and endpoints in the same Amazon Web
// Services Region.
//
// [Send Custom Platform-Specific Payloads in Messages to Mobile Devices]: https://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-custommessage.html
func (c *Client) Publish(ctx context.Context, params *PublishInput, optFns ...func(*Options)) (*PublishOutput, error) {
	if params == nil {
		params = &PublishInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Publish", params, optFns, c.addOperationPublishMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PublishOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for Publish action.
type PublishInput struct {

	// The message you want to send.
	//
	// If you are publishing to a topic and you want to send the same message to all
	// transport protocols, include the text of the message as a String value. If you
	// want to send different messages for each transport protocol, set the value of
	// the MessageStructure parameter to json and use a JSON object for the Message
	// parameter.
	//
	// Constraints:
	//
	//   - With the exception of SMS, messages must be UTF-8 encoded strings and at
	//   most 256 KB in size (262,144 bytes, not 262,144 characters).
	//
	//   - For SMS, each message can contain up to 140 characters. This character
	//   limit depends on the encoding schema. For example, an SMS message can contain
	//   160 GSM characters, 140 ASCII characters, or 70 UCS-2 characters.
	//
	// If you publish a message that exceeds this size limit, Amazon SNS sends the
	//   message as multiple messages, each fitting within the size limit. Messages
	//   aren't truncated mid-word but are cut off at whole-word boundaries.
	//
	// The total size limit for a single SMS Publish action is 1,600 characters.
	//
	// JSON-specific constraints:
	//
	//   - Keys in the JSON object that correspond to supported transport protocols
	//   must have simple JSON string values.
	//
	//   - The values will be parsed (unescaped) before they are used in outgoing
	//   messages.
	//
	//   - Outbound notifications are JSON encoded (meaning that the characters will
	//   be reescaped for sending).
	//
	//   - Values have a minimum length of 0 (the empty string, "", is allowed).
	//
	//   - Values have a maximum length bounded by the overall message size (so,
	//   including multiple protocols may limit message sizes).
	//
	//   - Non-string values will cause the key to be ignored.
	//
	//   - Keys that do not correspond to supported transport protocols are ignored.
	//
	//   - Duplicate keys are not allowed.
	//
	//   - Failure to parse or validate any key or value in the message will cause the
	//   Publish call to return an error (no partial delivery).
	//
	// This member is required.
	Message *string

	// Message attributes for Publish action.
	MessageAttributes map[string]types.MessageAttributeValue

	// This parameter applies only to FIFO (first-in-first-out) topics. The
	// MessageDeduplicationId can contain up to 128 alphanumeric characters (a-z, A-Z,
	// 0-9) and punctuation (!"#$%&'()*+,-./:;<=>?@[\]^_`{|}~) .
	//
	// Every message must have a unique MessageDeduplicationId , which is a token used
	// for deduplication of sent messages. If a message with a particular
	// MessageDeduplicationId is sent successfully, any message sent with the same
	// MessageDeduplicationId during the 5-minute deduplication interval is treated as
	// a duplicate.
	//
	// If the topic has ContentBasedDeduplication set, the system generates a
	// MessageDeduplicationId based on the contents of the message. Your
	// MessageDeduplicationId overrides the generated one.
	MessageDeduplicationId *string

	// This parameter applies only to FIFO (first-in-first-out) topics. The
	// MessageGroupId can contain up to 128 alphanumeric characters (a-z, A-Z, 0-9)
	// and punctuation (!"#$%&'()*+,-./:;<=>?@[\]^_`{|}~) .
	//
	// The MessageGroupId is a tag that specifies that a message belongs to a specific
	// message group. Messages that belong to the same message group are processed in a
	// FIFO manner (however, messages in different message groups might be processed
	// out of order). Every message must include a MessageGroupId .
	MessageGroupId *string

	// Set MessageStructure to json if you want to send a different message for each
	// protocol. For example, using one publish action, you can send a short message to
	// your SMS subscribers and a longer message to your email subscribers. If you set
	// MessageStructure to json , the value of the Message parameter must:
	//
	//   - be a syntactically valid JSON object; and
	//
	//   - contain at least a top-level JSON key of "default" with a value that is a
	//   string.
	//
	// You can define other top-level keys that define the message you want to send to
	// a specific transport protocol (e.g., "http").
	//
	// Valid value: json
	MessageStructure *string

	// The phone number to which you want to deliver an SMS message. Use E.164 format.
	//
	// If you don't specify a value for the PhoneNumber parameter, you must specify a
	// value for the TargetArn or TopicArn parameters.
	PhoneNumber *string

	// Optional parameter to be used as the "Subject" line when the message is
	// delivered to email endpoints. This field will also be included, if present, in
	// the standard JSON messages delivered to other endpoints.
	//
	// Constraints: Subjects must be ASCII text that begins with a letter, number, or
	// punctuation mark; must not include line breaks or control characters; and must
	// be less than 100 characters long.
	Subject *string

	// If you don't specify a value for the TargetArn parameter, you must specify a
	// value for the PhoneNumber or TopicArn parameters.
	TargetArn *string

	// The topic you want to publish to.
	//
	// If you don't specify a value for the TopicArn parameter, you must specify a
	// value for the PhoneNumber or TargetArn parameters.
	TopicArn *string

	noSmithyDocumentSerde
}

// Response for Publish action.
type PublishOutput struct {

	// Unique identifier assigned to the published message.
	//
	// Length Constraint: Maximum 100 characters
	MessageId *string

	// This response element applies only to FIFO (first-in-first-out) topics.
	//
	// The sequence number is a large, non-consecutive number that Amazon SNS assigns
	// to each message. The length of SequenceNumber is 128 bits. SequenceNumber
	// continues to increase for each MessageGroupId .
	SequenceNumber *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPublishMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPublish{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPublish{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "Publish"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addTimeOffsetDeserializer(stack, c); err != nil {
		return err
	}
	if err = addOpPublishValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPublish(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPublish(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "Publish",
	}
}
