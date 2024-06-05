// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2query

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Automatically adds idempotency tokens.
func (c *Client) QueryIdempotencyTokenAutoFill(ctx context.Context, params *QueryIdempotencyTokenAutoFillInput, optFns ...func(*Options)) (*QueryIdempotencyTokenAutoFillOutput, error) {
	if params == nil {
		params = &QueryIdempotencyTokenAutoFillInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "QueryIdempotencyTokenAutoFill", params, optFns, c.addOperationQueryIdempotencyTokenAutoFillMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*QueryIdempotencyTokenAutoFillOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type QueryIdempotencyTokenAutoFillInput struct {
	Token *string

	noSmithyDocumentSerde
}

type QueryIdempotencyTokenAutoFillOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationQueryIdempotencyTokenAutoFillMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpQueryIdempotencyTokenAutoFill{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpQueryIdempotencyTokenAutoFill{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "QueryIdempotencyTokenAutoFill"); err != nil {
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
	if err = addIdempotencyToken_opQueryIdempotencyTokenAutoFillMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opQueryIdempotencyTokenAutoFill(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpQueryIdempotencyTokenAutoFill struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpQueryIdempotencyTokenAutoFill) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpQueryIdempotencyTokenAutoFill) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*QueryIdempotencyTokenAutoFillInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *QueryIdempotencyTokenAutoFillInput ")
	}

	if input.Token == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.Token = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opQueryIdempotencyTokenAutoFillMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpQueryIdempotencyTokenAutoFill{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opQueryIdempotencyTokenAutoFill(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "QueryIdempotencyTokenAutoFill",
	}
}
