// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a description of a model explainability job definition.
func (c *Client) DescribeModelExplainabilityJobDefinition(ctx context.Context, params *DescribeModelExplainabilityJobDefinitionInput, optFns ...func(*Options)) (*DescribeModelExplainabilityJobDefinitionOutput, error) {
	if params == nil {
		params = &DescribeModelExplainabilityJobDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeModelExplainabilityJobDefinition", params, optFns, c.addOperationDescribeModelExplainabilityJobDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeModelExplainabilityJobDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeModelExplainabilityJobDefinitionInput struct {

	// The name of the model explainability job definition. The name must be unique
	// within an Amazon Web Services Region in the Amazon Web Services account.
	//
	// This member is required.
	JobDefinitionName *string

	noSmithyDocumentSerde
}

type DescribeModelExplainabilityJobDefinitionOutput struct {

	// The time at which the model explainability job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the model explainability job.
	//
	// This member is required.
	JobDefinitionArn *string

	// The name of the explainability job definition. The name must be unique within
	// an Amazon Web Services Region in the Amazon Web Services account.
	//
	// This member is required.
	JobDefinitionName *string

	// Identifies the resources to deploy for a monitoring job.
	//
	// This member is required.
	JobResources *types.MonitoringResources

	// Configures the model explainability job to run a specified Docker container
	// image.
	//
	// This member is required.
	ModelExplainabilityAppSpecification *types.ModelExplainabilityAppSpecification

	// Inputs for the model explainability job.
	//
	// This member is required.
	ModelExplainabilityJobInput *types.ModelExplainabilityJobInput

	// The output configuration for monitoring jobs.
	//
	// This member is required.
	ModelExplainabilityJobOutputConfig *types.MonitoringOutputConfig

	// The Amazon Resource Name (ARN) of the IAM role that has read permission to the
	// input data location and write permission to the output data location in Amazon
	// S3.
	//
	// This member is required.
	RoleArn *string

	// The baseline configuration for a model explainability job.
	ModelExplainabilityBaselineConfig *types.ModelExplainabilityBaselineConfig

	// Networking options for a model explainability job.
	NetworkConfig *types.MonitoringNetworkConfig

	// A time limit for how long the monitoring job is allowed to run before stopping.
	StoppingCondition *types.MonitoringStoppingCondition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeModelExplainabilityJobDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeModelExplainabilityJobDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeModelExplainabilityJobDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeModelExplainabilityJobDefinition"); err != nil {
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
	if err = addOpDescribeModelExplainabilityJobDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeModelExplainabilityJobDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeModelExplainabilityJobDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeModelExplainabilityJobDefinition",
	}
}
