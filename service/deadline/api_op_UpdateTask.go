// Code generated by smithy-go-codegen DO NOT EDIT.

package deadline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/deadline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a task.
func (c *Client) UpdateTask(ctx context.Context, params *UpdateTaskInput, optFns ...func(*Options)) (*UpdateTaskOutput, error) {
	if params == nil {
		params = &UpdateTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTask", params, optFns, c.addOperationUpdateTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTaskInput struct {

	// The farm ID to update.
	//
	// This member is required.
	FarmId *string

	// The job ID to update.
	//
	// This member is required.
	JobId *string

	// The queue ID to update.
	//
	// This member is required.
	QueueId *string

	// The step ID to update.
	//
	// This member is required.
	StepId *string

	// The run status with which to start the task.
	//
	// This member is required.
	TargetRunStatus types.TaskTargetRunStatus

	// The task ID to update.
	//
	// This member is required.
	TaskId *string

	// The unique token which the server uses to recognize retries of the same request.
	ClientToken *string

	noSmithyDocumentSerde
}

type UpdateTaskOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateTask"); err != nil {
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
	if err = addEndpointPrefix_opUpdateTaskMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opUpdateTaskMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTask(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opUpdateTaskMiddleware struct {
}

func (*endpointPrefix_opUpdateTaskMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opUpdateTaskMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "management." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opUpdateTaskMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opUpdateTaskMiddleware{}, "ResolveEndpointV2", middleware.After)
}

type idempotencyToken_initializeOpUpdateTask struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateTask) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateTaskInput ")
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
func addIdempotencyToken_opUpdateTaskMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateTask{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateTask",
	}
}
