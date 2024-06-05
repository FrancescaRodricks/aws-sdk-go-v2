// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an export task so that you can efficiently export data from a log group
// to an Amazon S3 bucket. When you perform a CreateExportTask operation, you must
// use credentials that have permission to write to the S3 bucket that you specify
// as the destination.
//
// Exporting log data to S3 buckets that are encrypted by KMS is supported.
// Exporting log data to Amazon S3 buckets that have S3 Object Lock enabled with a
// retention period is also supported.
//
// Exporting to S3 buckets that are encrypted with AES-256 is supported.
//
// This is an asynchronous call. If all the required information is provided, this
// operation initiates an export task and responds with the ID of the task. After
// the task has started, you can use [DescribeExportTasks]to get the status of the export task. Each
// account can only have one active ( RUNNING or PENDING ) export task at a time.
// To cancel an export task, use [CancelExportTask].
//
// You can export logs from multiple log groups or multiple time ranges to the
// same S3 bucket. To separate log data for each export task, specify a prefix to
// be used as the Amazon S3 key prefix for all exported objects.
//
// Time-based sorting on chunks of log data inside an exported file is not
// guaranteed. You can sort the exported log field data by using Linux utilities.
//
// [CancelExportTask]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CancelExportTask.html
// [DescribeExportTasks]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeExportTasks.html
func (c *Client) CreateExportTask(ctx context.Context, params *CreateExportTaskInput, optFns ...func(*Options)) (*CreateExportTaskOutput, error) {
	if params == nil {
		params = &CreateExportTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateExportTask", params, optFns, c.addOperationCreateExportTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateExportTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateExportTaskInput struct {

	// The name of S3 bucket for the exported log data. The bucket must be in the same
	// Amazon Web Services Region.
	//
	// This member is required.
	Destination *string

	// The start time of the range for the request, expressed as the number of
	// milliseconds after Jan 1, 1970 00:00:00 UTC . Events with a timestamp earlier
	// than this time are not exported.
	//
	// This member is required.
	From *int64

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string

	// The end time of the range for the request, expressed as the number of
	// milliseconds after Jan 1, 1970 00:00:00 UTC . Events with a timestamp later than
	// this time are not exported.
	//
	// You must specify a time that is not earlier than when this log group was
	// created.
	//
	// This member is required.
	To *int64

	// The prefix used as the start of the key for every object exported. If you don't
	// specify a value, the default is exportedlogs .
	DestinationPrefix *string

	// Export only log streams that match the provided prefix. If you don't specify a
	// value, no prefix filter is applied.
	LogStreamNamePrefix *string

	// The name of the export task.
	TaskName *string

	noSmithyDocumentSerde
}

type CreateExportTaskOutput struct {

	// The ID of the export task.
	TaskId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateExportTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateExportTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateExportTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateExportTask"); err != nil {
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
	if err = addOpCreateExportTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateExportTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateExportTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateExportTask",
	}
}
