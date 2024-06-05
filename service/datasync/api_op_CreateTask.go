// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Configures a task, which defines where and how DataSync transfers your data.
//
// A task includes a source location, destination location, and transfer options
// (such as bandwidth limits, scheduling, and more).
//
// If you're planning to transfer data to or from an Amazon S3 location, review [how DataSync can affect your S3 request charges]
// and the [DataSync pricing page]before you begin.
//
// [how DataSync can affect your S3 request charges]: https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#create-s3-location-s3-requests
// [DataSync pricing page]: http://aws.amazon.com/datasync/pricing/
func (c *Client) CreateTask(ctx context.Context, params *CreateTaskInput, optFns ...func(*Options)) (*CreateTaskOutput, error) {
	if params == nil {
		params = &CreateTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTask", params, optFns, c.addOperationCreateTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateTaskRequest
type CreateTaskInput struct {

	// Specifies the ARN of your transfer's destination location.
	//
	// This member is required.
	DestinationLocationArn *string

	// Specifies the ARN of your transfer's source location.
	//
	// This member is required.
	SourceLocationArn *string

	// Specifies the Amazon Resource Name (ARN) of an Amazon CloudWatch log group for
	// monitoring your task.
	CloudWatchLogGroupArn *string

	// Specifies exclude filters that define the files, objects, and folders in your
	// source location that you don't want DataSync to transfer. For more information
	// and examples, see [Specifying what DataSync transfers by using filters].
	//
	// [Specifying what DataSync transfers by using filters]: https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html
	Excludes []types.FilterRule

	// Specifies include filters define the files, objects, and folders in your source
	// location that you want DataSync to transfer. For more information and examples,
	// see [Specifying what DataSync transfers by using filters].
	//
	// [Specifying what DataSync transfers by using filters]: https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html
	Includes []types.FilterRule

	// Configures a manifest, which is a list of files or objects that you want
	// DataSync to transfer. For more information and configuration examples, see [Specifying what DataSync transfers by using a manifest].
	//
	// When using this parameter, your caller identity (the role that you're using
	// DataSync with) must have the iam:PassRole permission. The [AWSDataSyncFullAccess] policy includes this
	// permission.
	//
	// [AWSDataSyncFullAccess]: https://docs.aws.amazon.com/datasync/latest/userguide/security-iam-awsmanpol.html#security-iam-awsmanpol-awsdatasyncfullaccess
	// [Specifying what DataSync transfers by using a manifest]: https://docs.aws.amazon.com/datasync/latest/userguide/transferring-with-manifest.html
	ManifestConfig *types.ManifestConfig

	// Specifies the name of your task.
	Name *string

	// Specifies your task's settings, such as preserving file metadata, verifying
	// data integrity, among other options.
	Options *types.Options

	// Specifies a schedule for when you want your task to run. For more information,
	// see [Scheduling your task].
	//
	// [Scheduling your task]: https://docs.aws.amazon.com/datasync/latest/userguide/task-scheduling.html
	Schedule *types.TaskSchedule

	// Specifies the tags that you want to apply to your task.
	//
	// Tags are key-value pairs that help you manage, filter, and search for your
	// DataSync resources.
	Tags []types.TagListEntry

	// Specifies how you want to configure a task report, which provides detailed
	// information about your DataSync transfer. For more information, see [Monitoring your DataSync transfers with task reports].
	//
	// When using this parameter, your caller identity (the role that you're using
	// DataSync with) must have the iam:PassRole permission. The [AWSDataSyncFullAccess] policy includes this
	// permission.
	//
	// [AWSDataSyncFullAccess]: https://docs.aws.amazon.com/datasync/latest/userguide/security-iam-awsmanpol.html#security-iam-awsmanpol-awsdatasyncfullaccess
	// [Monitoring your DataSync transfers with task reports]: https://docs.aws.amazon.com/datasync/latest/userguide/task-reports.html
	TaskReportConfig *types.TaskReportConfig

	noSmithyDocumentSerde
}

// CreateTaskResponse
type CreateTaskOutput struct {

	// The Amazon Resource Name (ARN) of the task.
	TaskArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateTask"); err != nil {
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
	if err = addOpCreateTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateTask",
	}
}
