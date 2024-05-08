// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Places an outbound call to a contact, and then initiates the flow. It performs
// the actions in the flow that's specified (in ContactFlowId ).
//
// Agents do not initiate the outbound API, which means that they do not dial the
// contact. If the flow places an outbound call to a contact, and then puts the
// contact in queue, the call is then routed to the agent, like any other inbound
// case.
//
// There is a 60-second dialing timeout for this operation. If the call is not
// connected after 60 seconds, it fails.
//
// UK numbers with a 447 prefix are not allowed by default. Before you can dial
// these UK mobile numbers, you must submit a service quota increase request. For
// more information, see [Amazon Connect Service Quotas]in the Amazon Connect Administrator Guide.
//
// Campaign calls are not allowed by default. Before you can make a call with
// TrafficType = CAMPAIGN , you must submit a service quota increase request to the
// quota [Amazon Connect campaigns].
//
// [Amazon Connect Service Quotas]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html
// [Amazon Connect campaigns]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html#outbound-communications-quotas
func (c *Client) StartOutboundVoiceContact(ctx context.Context, params *StartOutboundVoiceContactInput, optFns ...func(*Options)) (*StartOutboundVoiceContactOutput, error) {
	if params == nil {
		params = &StartOutboundVoiceContactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartOutboundVoiceContact", params, optFns, c.addOperationStartOutboundVoiceContactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartOutboundVoiceContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartOutboundVoiceContactInput struct {

	// The identifier of the flow for the outbound call. To see the ContactFlowId in
	// the Amazon Connect admin website, on the navigation menu go to Routing, Contact
	// Flows. Choose the flow. On the flow page, under the name of the flow, choose
	// Show additional flow information. The ContactFlowId is the last part of the ARN,
	// shown here in bold:
	//
	// arn:aws:connect:us-west-2:xxxxxxxxxxxx:instance/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/contact-flow/846ec553-a005-41c0-8341-xxxxxxxxxxxx
	//
	// This member is required.
	ContactFlowId *string

	// The phone number of the customer, in E.164 format.
	//
	// This member is required.
	DestinationPhoneNumber *string

	// The identifier of the Amazon Connect instance. You can [find the instance ID] in the Amazon Resource
	// Name (ARN) of the instance.
	//
	// [find the instance ID]: https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html
	//
	// This member is required.
	InstanceId *string

	// Configuration of the answering machine detection for this outbound call.
	AnswerMachineDetectionConfig *types.AnswerMachineDetectionConfig

	// A custom key-value pair using an attribute map. The attributes are standard
	// Amazon Connect attributes, and can be accessed in flows just like any other
	// contact attributes.
	//
	// There can be up to 32,768 UTF-8 bytes across all key-value pairs per contact.
	// Attribute keys can include only alphanumeric, dash, and underscore characters.
	Attributes map[string]string

	// The campaign identifier of the outbound communication.
	CampaignId *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If not provided, the Amazon Web Services SDK populates this
	// field. For more information about idempotency, see [Making retries safe with idempotent APIs]. The token is valid for 7
	// days after creation. If a contact is already started, the contact ID is
	// returned.
	//
	// [Making retries safe with idempotent APIs]: https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/
	ClientToken *string

	// A description of the voice contact that is shown to an agent in the Contact
	// Control Panel (CCP).
	Description *string

	// The name of a voice contact that is shown to an agent in the Contact Control
	// Panel (CCP).
	Name *string

	// The queue for the call. If you specify a queue, the phone displayed for caller
	// ID is the phone number specified in the queue. If you do not specify a queue,
	// the queue defined in the flow is used. If you do not specify a queue, you must
	// specify a source phone number.
	QueueId *string

	// A formatted URL that is shown to an agent in the Contact Control Panel (CCP).
	// Contacts can have the following reference types at the time of creation: URL |
	// NUMBER | STRING | DATE | EMAIL . ATTACHMENT is not a supported reference type
	// during voice contact creation.
	References map[string]types.Reference

	// The contactId that is related to this contact. Linking voice, task, or chat by
	// using RelatedContactID copies over contact attributes from the related contact
	// to the new contact. All updates to user-defined attributes in the new contact
	// are limited to the individual contact ID. There are no limits to the number of
	// contacts that can be linked by using RelatedContactId .
	RelatedContactId *string

	// The phone number associated with the Amazon Connect instance, in E.164 format.
	// If you do not specify a source phone number, you must specify a queue.
	SourcePhoneNumber *string

	// Denotes the class of traffic. Calls with different traffic types are handled
	// differently by Amazon Connect. The default value is GENERAL . Use CAMPAIGN if
	// EnableAnswerMachineDetection is set to true . For all other cases, use GENERAL .
	TrafficType types.TrafficType

	noSmithyDocumentSerde
}

type StartOutboundVoiceContactOutput struct {

	// The identifier of this contact within the Amazon Connect instance.
	ContactId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartOutboundVoiceContactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartOutboundVoiceContact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartOutboundVoiceContact{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartOutboundVoiceContact"); err != nil {
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
	if err = addIdempotencyToken_opStartOutboundVoiceContactMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartOutboundVoiceContactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartOutboundVoiceContact(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartOutboundVoiceContact struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartOutboundVoiceContact) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartOutboundVoiceContact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartOutboundVoiceContactInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartOutboundVoiceContactInput ")
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
func addIdempotencyToken_opStartOutboundVoiceContactMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartOutboundVoiceContact{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartOutboundVoiceContact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartOutboundVoiceContact",
	}
}
