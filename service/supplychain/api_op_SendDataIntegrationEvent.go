// Code generated by smithy-go-codegen DO NOT EDIT.

package supplychain

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/supplychain/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Send transactional data events with real-time data for analysis or monitoring.
func (c *Client) SendDataIntegrationEvent(ctx context.Context, params *SendDataIntegrationEventInput, optFns ...func(*Options)) (*SendDataIntegrationEventOutput, error) {
	if params == nil {
		params = &SendDataIntegrationEventInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendDataIntegrationEvent", params, optFns, c.addOperationSendDataIntegrationEventMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendDataIntegrationEventOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request parameters for SendDataIntegrationEvent.
type SendDataIntegrationEventInput struct {

	// The data payload of the event.
	//
	// This member is required.
	Data *string

	// Event identifier (for example, orderId for InboundOrder) used for data sharing
	// or partitioning.
	//
	// This member is required.
	EventGroupId *string

	// The data event type.
	//
	// This member is required.
	EventType types.DataIntegrationEventType

	// The AWS Supply Chain instance identifier.
	//
	// This member is required.
	InstanceId *string

	// The idempotent client token.
	ClientToken *string

	// The event timestamp (in epoch seconds).
	EventTimestamp *time.Time

	noSmithyDocumentSerde
}

// The response parameters for SendDataIntegrationEvent.
type SendDataIntegrationEventOutput struct {

	// The unique event identifier.
	//
	// This member is required.
	EventId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendDataIntegrationEventMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSendDataIntegrationEvent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSendDataIntegrationEvent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendDataIntegrationEvent"); err != nil {
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
	if err = addIdempotencyToken_opSendDataIntegrationEventMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpSendDataIntegrationEventValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendDataIntegrationEvent(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpSendDataIntegrationEvent struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpSendDataIntegrationEvent) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpSendDataIntegrationEvent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*SendDataIntegrationEventInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *SendDataIntegrationEventInput ")
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
func addIdempotencyToken_opSendDataIntegrationEventMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpSendDataIntegrationEvent{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opSendDataIntegrationEvent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendDataIntegrationEvent",
	}
}
