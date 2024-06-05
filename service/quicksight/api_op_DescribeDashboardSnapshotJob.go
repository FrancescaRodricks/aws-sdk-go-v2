// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes an existing snapshot job.
//
// Poll job descriptions after a job starts to know the status of the job. For
// information on available status codes, see JobStatus .
func (c *Client) DescribeDashboardSnapshotJob(ctx context.Context, params *DescribeDashboardSnapshotJobInput, optFns ...func(*Options)) (*DescribeDashboardSnapshotJobOutput, error) {
	if params == nil {
		params = &DescribeDashboardSnapshotJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDashboardSnapshotJob", params, optFns, c.addOperationDescribeDashboardSnapshotJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDashboardSnapshotJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDashboardSnapshotJobInput struct {

	// The ID of the Amazon Web Services account that the dashboard snapshot job is
	// executed in.
	//
	// This member is required.
	AwsAccountId *string

	// The ID of the dashboard that you have started a snapshot job for.
	//
	// This member is required.
	DashboardId *string

	// The ID of the job to be described. The job ID is set when you start a new job
	// with a StartDashboardSnapshotJob API call.
	//
	// This member is required.
	SnapshotJobId *string

	noSmithyDocumentSerde
}

type DescribeDashboardSnapshotJobOutput struct {

	// The Amazon Resource Name (ARN) for the snapshot job. The job ARN is generated
	// when you start a new job with a StartDashboardSnapshotJob API call.
	Arn *string

	//  The ID of the Amazon Web Services account that the dashboard snapshot job is
	// executed in.
	AwsAccountId *string

	//  The time that the snapshot job was created.
	CreatedTime *time.Time

	// The ID of the dashboard that you have started a snapshot job for.
	DashboardId *string

	// Indicates the status of a job. The status updates as the job executes. This
	// shows one of the following values.
	//
	//   - COMPLETED - The job was completed successfully.
	//
	//   - FAILED - The job failed to execute.
	//
	//   - QUEUED - The job is queued and hasn't started yet.
	//
	//   - RUNNING - The job is still running.
	JobStatus types.SnapshotJobStatus

	//  The time that the snapshot job status was last updated.
	LastUpdatedTime *time.Time

	//  The Amazon Web Services request ID for this operation.
	RequestId *string

	// The snapshot configuration of the job. This information is provided when you
	// make a StartDashboardSnapshotJob API call.
	SnapshotConfiguration *types.SnapshotConfiguration

	// The ID of the job to be described. The job ID is set when you start a new job
	// with a StartDashboardSnapshotJob API call.
	SnapshotJobId *string

	// The HTTP status of the request
	Status int32

	// The user configuration for the snapshot job. This information is provided when
	// you make a StartDashboardSnapshotJob API call.
	UserConfiguration *types.SnapshotUserConfigurationRedacted

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDashboardSnapshotJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDashboardSnapshotJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDashboardSnapshotJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDashboardSnapshotJob"); err != nil {
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
	if err = addOpDescribeDashboardSnapshotJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDashboardSnapshotJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDashboardSnapshotJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDashboardSnapshotJob",
	}
}
