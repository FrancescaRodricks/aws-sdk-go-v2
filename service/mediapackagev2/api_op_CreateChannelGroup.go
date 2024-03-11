// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackagev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Create a channel group to group your channels and origin endpoints. A channel
// group is the top-level resource that consists of channels and origin endpoints
// that are associated with it and that provides predictable URLs for stream
// delivery. All channels and origin endpoints within the channel group are
// guaranteed to share the DNS. You can create only one channel group with each
// request.
func (c *Client) CreateChannelGroup(ctx context.Context, params *CreateChannelGroupInput, optFns ...func(*Options)) (*CreateChannelGroupOutput, error) {
	if params == nil {
		params = &CreateChannelGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateChannelGroup", params, optFns, c.addOperationCreateChannelGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateChannelGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateChannelGroupInput struct {

	// The name that describes the channel group. The name is the primary identifier
	// for the channel group, and must be unique for your account in the AWS Region.
	// You can't use spaces in the name. You can't change the name after you create the
	// channel group.
	//
	// This member is required.
	ChannelGroupName *string

	// A unique, case-sensitive token that you provide to ensure the idempotency of
	// the request.
	ClientToken *string

	// Enter any descriptive text that helps you to identify the channel group.
	Description *string

	// A comma-separated list of tag key:value pairs that you define. For example:
	// "Key1": "Value1",
	//     "Key2": "Value2"
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateChannelGroupOutput struct {

	// The Amazon Resource Name (ARN) associated with the resource.
	//
	// This member is required.
	Arn *string

	// The name that describes the channel group. The name is the primary identifier
	// for the channel group, and must be unique for your account in the AWS Region.
	//
	// This member is required.
	ChannelGroupName *string

	// The date and time the channel group was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The output domain where the source stream should be sent. Integrate the egress
	// domain with a downstream CDN (such as Amazon CloudFront) or playback device.
	//
	// This member is required.
	EgressDomain *string

	// The date and time the channel group was modified.
	//
	// This member is required.
	ModifiedAt *time.Time

	// The description for your channel group.
	Description *string

	// The current Entity Tag (ETag) associated with this resource. The entity tag can
	// be used to safely make concurrent updates to the resource.
	ETag *string

	// The comma-separated list of tag key:value pairs assigned to the channel group.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateChannelGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateChannelGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateChannelGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateChannelGroup"); err != nil {
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
	if err = addIdempotencyToken_opCreateChannelGroupMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateChannelGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateChannelGroup(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateChannelGroup struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateChannelGroup) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateChannelGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateChannelGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateChannelGroupInput ")
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
func addIdempotencyToken_opCreateChannelGroupMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateChannelGroup{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateChannelGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateChannelGroup",
	}
}
