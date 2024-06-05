// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a processing job.
func (c *Client) CreateProcessingJob(ctx context.Context, params *CreateProcessingJobInput, optFns ...func(*Options)) (*CreateProcessingJobOutput, error) {
	if params == nil {
		params = &CreateProcessingJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProcessingJob", params, optFns, c.addOperationCreateProcessingJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProcessingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProcessingJobInput struct {

	// Configures the processing job to run a specified Docker container image.
	//
	// This member is required.
	AppSpecification *types.AppSpecification

	//  The name of the processing job. The name must be unique within an Amazon Web
	// Services Region in the Amazon Web Services account.
	//
	// This member is required.
	ProcessingJobName *string

	// Identifies the resources, ML compute instances, and ML storage volumes to
	// deploy for a processing job. In distributed training, you specify more than one
	// instance.
	//
	// This member is required.
	ProcessingResources *types.ProcessingResources

	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume
	// to perform tasks on your behalf.
	//
	// This member is required.
	RoleArn *string

	// The environment variables to set in the Docker container. Up to 100 key and
	// values entries in the map are supported.
	Environment map[string]string

	// Associates a SageMaker job as a trial component with an experiment and trial.
	// Specified when you call the following APIs:
	//
	// [CreateProcessingJob]
	//
	// [CreateTrainingJob]
	//
	// [CreateTransformJob]
	//
	// [CreateTransformJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTransformJob.html
	// [CreateTrainingJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html
	// [CreateProcessingJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateProcessingJob.html
	ExperimentConfig *types.ExperimentConfig

	// Networking options for a processing job, such as whether to allow inbound and
	// outbound network calls to and from processing containers, and the VPC subnets
	// and security groups to use for VPC-enabled processing jobs.
	NetworkConfig *types.NetworkConfig

	// An array of inputs configuring the data to download into the processing
	// container.
	ProcessingInputs []types.ProcessingInput

	// Output configuration for the processing job.
	ProcessingOutputConfig *types.ProcessingOutputConfig

	// The time limit for how long the processing job is allowed to run.
	StoppingCondition *types.ProcessingStoppingCondition

	// (Optional) An array of key-value pairs. For more information, see [Using Cost Allocation Tags] in the
	// Amazon Web Services Billing and Cost Management User Guide.
	//
	// [Using Cost Allocation Tags]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateProcessingJobOutput struct {

	// The Amazon Resource Name (ARN) of the processing job.
	//
	// This member is required.
	ProcessingJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProcessingJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateProcessingJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateProcessingJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateProcessingJob"); err != nil {
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
	if err = addOpCreateProcessingJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProcessingJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateProcessingJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateProcessingJob",
	}
}
