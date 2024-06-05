// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Request an origination phone number for use in your account. For more
// information on phone number request see [Requesting a number]in the Amazon Pinpoint User Guide.
//
// [Requesting a number]: https://docs.aws.amazon.com/pinpoint/latest/userguide/settings-sms-request-number.html
func (c *Client) RequestPhoneNumber(ctx context.Context, params *RequestPhoneNumberInput, optFns ...func(*Options)) (*RequestPhoneNumberOutput, error) {
	if params == nil {
		params = &RequestPhoneNumberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RequestPhoneNumber", params, optFns, c.addOperationRequestPhoneNumberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RequestPhoneNumberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RequestPhoneNumberInput struct {

	// The two-character code, in ISO 3166-1 alpha-2 format, for the country or
	// region.
	//
	// This member is required.
	IsoCountryCode *string

	// The type of message. Valid values are TRANSACTIONAL for messages that are
	// critical or time-sensitive and PROMOTIONAL for messages that aren't critical or
	// time-sensitive.
	//
	// This member is required.
	MessageType types.MessageType

	// Indicates if the phone number will be used for text messages, voice messages,
	// or both.
	//
	// This member is required.
	NumberCapabilities []types.NumberCapability

	// The type of phone number to request.
	//
	// This member is required.
	NumberType types.RequestableNumberType

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. If you don't specify a client token, a randomly generated token is
	// used for the request to ensure idempotency.
	ClientToken *string

	// By default this is set to false. When set to true the phone number can't be
	// deleted.
	DeletionProtectionEnabled *bool

	// The name of the OptOutList to associate with the phone number. You can use the
	// OptOutListName or OptOutListArn.
	OptOutListName *string

	// The pool to associated with the phone number. You can use the PoolId or
	// PoolArn.
	PoolId *string

	// Use this field to attach your phone number for an external registration process.
	RegistrationId *string

	// An array of tags (key and value pairs) associate with the requested phone
	// number.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type RequestPhoneNumberOutput struct {

	// The time when the phone number was created, in [UNIX epoch time] format.
	//
	// [UNIX epoch time]: https://www.epochconverter.com/
	CreatedTimestamp *time.Time

	// By default this is set to false. When set to true the phone number can't be
	// deleted.
	DeletionProtectionEnabled bool

	// The two-character code, in ISO 3166-1 alpha-2 format, for the country or
	// region.
	IsoCountryCode *string

	// The type of message. Valid values are TRANSACTIONAL for messages that are
	// critical or time-sensitive and PROMOTIONAL for messages that aren't critical or
	// time-sensitive.
	MessageType types.MessageType

	// The monthly price, in US dollars, to lease the phone number.
	MonthlyLeasingPrice *string

	// Indicates if the phone number will be used for text messages, voice messages or
	// both.
	NumberCapabilities []types.NumberCapability

	// The type of number that was released.
	NumberType types.RequestableNumberType

	// The name of the OptOutList that is associated with the requested phone number.
	OptOutListName *string

	// The new phone number that was requested.
	PhoneNumber *string

	// The Amazon Resource Name (ARN) of the requested phone number.
	PhoneNumberArn *string

	// The unique identifier of the new phone number.
	PhoneNumberId *string

	// The unique identifier of the pool associated with the phone number
	PoolId *string

	// The unique identifier for the registration.
	RegistrationId *string

	// By default this is set to false. When an end recipient sends a message that
	// begins with HELP or STOP to one of your dedicated numbers, Amazon Pinpoint
	// automatically replies with a customizable message and adds the end recipient to
	// the OptOutList. When set to true you're responsible for responding to HELP and
	// STOP requests. You're also responsible for tracking and honoring opt-out
	// requests.
	SelfManagedOptOutsEnabled bool

	// The current status of the request.
	Status types.NumberStatus

	// An array of key and value pair tags that are associated with the phone number.
	Tags []types.Tag

	// The ARN used to identify the two way channel.
	TwoWayChannelArn *string

	// An optional IAM Role Arn for a service to assume, to be able to post inbound
	// SMS messages.
	TwoWayChannelRole *string

	// By default this is set to false. When set to true you can receive incoming text
	// messages from your end recipients.
	TwoWayEnabled bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRequestPhoneNumberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpRequestPhoneNumber{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpRequestPhoneNumber{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RequestPhoneNumber"); err != nil {
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
	if err = addIdempotencyToken_opRequestPhoneNumberMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpRequestPhoneNumberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRequestPhoneNumber(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpRequestPhoneNumber struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpRequestPhoneNumber) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpRequestPhoneNumber) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*RequestPhoneNumberInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *RequestPhoneNumberInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opRequestPhoneNumberMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpRequestPhoneNumber{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opRequestPhoneNumber(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RequestPhoneNumber",
	}
}
