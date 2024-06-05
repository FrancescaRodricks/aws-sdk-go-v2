// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// For a blue/green deployment, starts the process of rerouting traffic from
// instances in the original environment to instances in the replacement
// environment without waiting for a specified wait time to elapse. (Traffic
// rerouting, which is achieved by registering instances in the replacement
// environment with the load balancer, can start as soon as all instances have a
// status of Ready.)
func (c *Client) ContinueDeployment(ctx context.Context, params *ContinueDeploymentInput, optFns ...func(*Options)) (*ContinueDeploymentOutput, error) {
	if params == nil {
		params = &ContinueDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ContinueDeployment", params, optFns, c.addOperationContinueDeploymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ContinueDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ContinueDeploymentInput struct {

	//  The unique ID of a blue/green deployment for which you want to start rerouting
	// traffic to the replacement environment.
	DeploymentId *string

	//  The status of the deployment's waiting period. READY_WAIT indicates that the
	// deployment is ready to start shifting traffic. TERMINATION_WAIT indicates that
	// the traffic is shifted, but the original target is not terminated.
	DeploymentWaitType types.DeploymentWaitType

	noSmithyDocumentSerde
}

type ContinueDeploymentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationContinueDeploymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpContinueDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpContinueDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ContinueDeployment"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opContinueDeployment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opContinueDeployment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ContinueDeployment",
	}
}
