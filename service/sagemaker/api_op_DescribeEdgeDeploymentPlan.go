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

// Describes an edge deployment plan with deployment status per stage.
func (c *Client) DescribeEdgeDeploymentPlan(ctx context.Context, params *DescribeEdgeDeploymentPlanInput, optFns ...func(*Options)) (*DescribeEdgeDeploymentPlanOutput, error) {
	if params == nil {
		params = &DescribeEdgeDeploymentPlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEdgeDeploymentPlan", params, optFns, c.addOperationDescribeEdgeDeploymentPlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEdgeDeploymentPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEdgeDeploymentPlanInput struct {

	// The name of the deployment plan to describe.
	//
	// This member is required.
	EdgeDeploymentPlanName *string

	// The maximum number of results to select (50 by default).
	MaxResults *int32

	// If the edge deployment plan has enough stages to require tokening, then this is
	// the response from the last list of stages returned.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeEdgeDeploymentPlanOutput struct {

	// The device fleet used for this edge deployment plan.
	//
	// This member is required.
	DeviceFleetName *string

	// The ARN of edge deployment plan.
	//
	// This member is required.
	EdgeDeploymentPlanArn *string

	// The name of the edge deployment plan.
	//
	// This member is required.
	EdgeDeploymentPlanName *string

	// List of models associated with the edge deployment plan.
	//
	// This member is required.
	ModelConfigs []types.EdgeDeploymentModelConfig

	// List of stages in the edge deployment plan.
	//
	// This member is required.
	Stages []types.DeploymentStageStatusSummary

	// The time when the edge deployment plan was created.
	CreationTime *time.Time

	// The number of edge devices that failed the deployment.
	EdgeDeploymentFailed *int32

	// The number of edge devices yet to pick up deployment, or in progress.
	EdgeDeploymentPending *int32

	// The number of edge devices with the successful deployment.
	EdgeDeploymentSuccess *int32

	// The time when the edge deployment plan was last updated.
	LastModifiedTime *time.Time

	// Token to use when calling the next set of stages in the edge deployment plan.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEdgeDeploymentPlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEdgeDeploymentPlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEdgeDeploymentPlan{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEdgeDeploymentPlan"); err != nil {
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
	if err = addOpDescribeEdgeDeploymentPlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEdgeDeploymentPlan(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEdgeDeploymentPlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEdgeDeploymentPlan",
	}
}
