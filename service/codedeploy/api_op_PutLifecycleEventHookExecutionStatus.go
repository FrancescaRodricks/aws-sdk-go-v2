// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the result of a Lambda validation function. The function validates
// lifecycle hooks during a deployment that uses the AWS Lambda or Amazon ECS
// compute platform. For AWS Lambda deployments, the available lifecycle hooks are
// BeforeAllowTraffic and AfterAllowTraffic. For Amazon ECS deployments, the
// available lifecycle hooks are BeforeInstall, AfterInstall,
// AfterAllowTestTraffic, BeforeAllowTraffic, and AfterAllowTraffic. Lambda
// validation functions return Succeeded or Failed. For more information, see
// AppSpec 'hooks' Section for an AWS Lambda Deployment
// (https://docs.aws.amazon.com/codedeploy/latest/userguide/reference-appspec-file-structure-hooks.html#appspec-hooks-lambda)
// and AppSpec 'hooks' Section for an Amazon ECS Deployment
// (https://docs.aws.amazon.com/codedeploy/latest/userguide/reference-appspec-file-structure-hooks.html#appspec-hooks-ecs).
func (c *Client) PutLifecycleEventHookExecutionStatus(ctx context.Context, params *PutLifecycleEventHookExecutionStatusInput, optFns ...func(*Options)) (*PutLifecycleEventHookExecutionStatusOutput, error) {
	if params == nil {
		params = &PutLifecycleEventHookExecutionStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutLifecycleEventHookExecutionStatus", params, optFns, c.addOperationPutLifecycleEventHookExecutionStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutLifecycleEventHookExecutionStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLifecycleEventHookExecutionStatusInput struct {

	// The unique ID of a deployment. Pass this ID to a Lambda function that validates
	// a deployment lifecycle event.
	DeploymentId *string

	// The execution ID of a deployment's lifecycle hook. A deployment lifecycle hook
	// is specified in the hooks section of the AppSpec file.
	LifecycleEventHookExecutionId *string

	// The result of a Lambda function that validates a deployment lifecycle event.
	// Succeeded and Failed are the only valid values for status.
	Status types.LifecycleEventStatus

	noSmithyDocumentSerde
}

type PutLifecycleEventHookExecutionStatusOutput struct {

	// The execution ID of the lifecycle event hook. A hook is specified in the hooks
	// section of the deployment's AppSpec file.
	LifecycleEventHookExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutLifecycleEventHookExecutionStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutLifecycleEventHookExecutionStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutLifecycleEventHookExecutionStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutLifecycleEventHookExecutionStatus(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opPutLifecycleEventHookExecutionStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "PutLifecycleEventHookExecutionStatus",
	}
}