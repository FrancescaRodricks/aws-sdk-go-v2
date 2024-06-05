// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Exports a read set to Amazon S3.
func (c *Client) StartReadSetExportJob(ctx context.Context, params *StartReadSetExportJobInput, optFns ...func(*Options)) (*StartReadSetExportJobOutput, error) {
	if params == nil {
		params = &StartReadSetExportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartReadSetExportJob", params, optFns, c.addOperationStartReadSetExportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartReadSetExportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartReadSetExportJobInput struct {

	// A location for exported files in Amazon S3.
	//
	// This member is required.
	Destination *string

	// A service role for the job.
	//
	// This member is required.
	RoleArn *string

	// The read set's sequence store ID.
	//
	// This member is required.
	SequenceStoreId *string

	// The job's source files.
	//
	// This member is required.
	Sources []types.ExportReadSet

	// To ensure that jobs don't run multiple times, specify a unique token for each
	// job.
	ClientToken *string

	noSmithyDocumentSerde
}

type StartReadSetExportJobOutput struct {

	// When the job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The job's output location.
	//
	// This member is required.
	Destination *string

	// The job's ID.
	//
	// This member is required.
	Id *string

	// The read set's sequence store ID.
	//
	// This member is required.
	SequenceStoreId *string

	// The job's status.
	//
	// This member is required.
	Status types.ReadSetExportJobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartReadSetExportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartReadSetExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartReadSetExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartReadSetExportJob"); err != nil {
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
	if err = addEndpointPrefix_opStartReadSetExportJobMiddleware(stack); err != nil {
		return err
	}
	if err = addOpStartReadSetExportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartReadSetExportJob(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opStartReadSetExportJobMiddleware struct {
}

func (*endpointPrefix_opStartReadSetExportJobMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opStartReadSetExportJobMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "control-storage-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opStartReadSetExportJobMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opStartReadSetExportJobMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opStartReadSetExportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartReadSetExportJob",
	}
}
