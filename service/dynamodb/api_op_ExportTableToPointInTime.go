// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Exports table data to an S3 bucket. The table must have point in time recovery
// enabled, and you can export data from any time within the point in time recovery
// window.
func (c *Client) ExportTableToPointInTime(ctx context.Context, params *ExportTableToPointInTimeInput, optFns ...func(*Options)) (*ExportTableToPointInTimeOutput, error) {
	if params == nil {
		params = &ExportTableToPointInTimeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportTableToPointInTime", params, optFns, c.addOperationExportTableToPointInTimeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportTableToPointInTimeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportTableToPointInTimeInput struct {

	// The name of the Amazon S3 bucket to export the snapshot to.
	//
	// This member is required.
	S3Bucket *string

	// The Amazon Resource Name (ARN) associated with the table to export.
	//
	// This member is required.
	TableArn *string

	// Providing a ClientToken makes the call to ExportTableToPointInTimeInput
	// idempotent, meaning that multiple identical calls have the same effect as one
	// single call.
	//
	// A client token is valid for 8 hours after the first request that uses it is
	// completed. After 8 hours, any request with the same client token is treated as a
	// new request. Do not resubmit the same request with the same client token for
	// more than 8 hours, or the result might not be idempotent.
	//
	// If you submit a request with the same client token but a change in other
	// parameters within the 8-hour idempotency window, DynamoDB returns an
	// ImportConflictException .
	ClientToken *string

	// The format for the exported data. Valid values for ExportFormat are
	// DYNAMODB_JSON or ION .
	ExportFormat types.ExportFormat

	// Time in the past from which to export table data, counted in seconds from the
	// start of the Unix epoch. The table export will be a snapshot of the table's
	// state at this point in time.
	ExportTime *time.Time

	// Choice of whether to execute as a full export or incremental export. Valid
	// values are FULL_EXPORT or INCREMENTAL_EXPORT. The default value is FULL_EXPORT.
	// If INCREMENTAL_EXPORT is provided, the IncrementalExportSpecification must also
	// be used.
	ExportType types.ExportType

	// Optional object containing the parameters specific to an incremental export.
	IncrementalExportSpecification *types.IncrementalExportSpecification

	// The ID of the Amazon Web Services account that owns the bucket the export will
	// be stored in.
	//
	// S3BucketOwner is a required parameter when exporting to a S3 bucket in another
	// account.
	S3BucketOwner *string

	// The Amazon S3 bucket prefix to use as the file name and path of the exported
	// snapshot.
	S3Prefix *string

	// Type of encryption used on the bucket where export data will be stored. Valid
	// values for S3SseAlgorithm are:
	//
	//   - AES256 - server-side encryption with Amazon S3 managed keys
	//
	//   - KMS - server-side encryption with KMS managed keys
	S3SseAlgorithm types.S3SseAlgorithm

	// The ID of the KMS managed key used to encrypt the S3 bucket where export data
	// will be stored (if applicable).
	S3SseKmsKeyId *string

	noSmithyDocumentSerde
}

type ExportTableToPointInTimeOutput struct {

	// Contains a description of the table export.
	ExportDescription *types.ExportDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportTableToPointInTimeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpExportTableToPointInTime{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpExportTableToPointInTime{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ExportTableToPointInTime"); err != nil {
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
	if err = addIdempotencyToken_opExportTableToPointInTimeMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpExportTableToPointInTimeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportTableToPointInTime(options.Region), middleware.Before); err != nil {
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
	if err = addValidateResponseChecksum(stack, options); err != nil {
		return err
	}
	if err = addAcceptEncodingGzip(stack, options); err != nil {
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

type idempotencyToken_initializeOpExportTableToPointInTime struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpExportTableToPointInTime) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpExportTableToPointInTime) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ExportTableToPointInTimeInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ExportTableToPointInTimeInput ")
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
func addIdempotencyToken_opExportTableToPointInTimeMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpExportTableToPointInTime{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opExportTableToPointInTime(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ExportTableToPointInTime",
	}
}
