// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a continuous deployment for a target, which is a Greengrass core device
// or group of core devices. When you add a new core device to a group of core
// devices that has a deployment, IoT Greengrass deploys that group's deployment to
// the new device.
//
// You can define one deployment for each target. When you create a new deployment
// for a target that has an existing deployment, you replace the previous
// deployment. IoT Greengrass applies the new deployment to the target devices.
//
// Every deployment has a revision number that indicates how many deployment
// revisions you define for a target. Use this operation to create a new revision
// of an existing deployment.
//
// For more information, see the [Create deployments] in the IoT Greengrass V2 Developer Guide.
//
// [Create deployments]: https://docs.aws.amazon.com/greengrass/v2/developerguide/create-deployments.html
func (c *Client) CreateDeployment(ctx context.Context, params *CreateDeploymentInput, optFns ...func(*Options)) (*CreateDeploymentOutput, error) {
	if params == nil {
		params = &CreateDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDeployment", params, optFns, c.addOperationCreateDeploymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDeploymentInput struct {

	// The [ARN] of the target IoT thing or thing group. When creating a subdeployment, the
	// targetARN can only be a thing group.
	//
	// [ARN]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	//
	// This member is required.
	TargetArn *string

	// A unique, case-sensitive identifier that you can provide to ensure that the
	// request is idempotent. Idempotency means that the request is successfully
	// processed only once, even if you send the request multiple times. When a request
	// succeeds, and you specify the same client token for subsequent successful
	// requests, the IoT Greengrass V2 service returns the successful response that it
	// caches from the previous request. IoT Greengrass V2 caches successful responses
	// for idempotent requests for up to 8 hours.
	ClientToken *string

	// The components to deploy. This is a dictionary, where each key is the name of a
	// component, and each key's value is the version and configuration to deploy for
	// that component.
	Components map[string]types.ComponentDeploymentSpecification

	// The name of the deployment.
	DeploymentName *string

	// The deployment policies for the deployment. These policies define how the
	// deployment updates components and handles failure.
	DeploymentPolicies *types.DeploymentPolicies

	// The job configuration for the deployment configuration. The job configuration
	// specifies the rollout, timeout, and stop configurations for the deployment
	// configuration.
	IotJobConfiguration *types.DeploymentIoTJobConfiguration

	// The parent deployment's target [ARN] within a subdeployment.
	//
	// [ARN]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	ParentTargetArn *string

	// A list of key-value pairs that contain metadata for the resource. For more
	// information, see [Tag your resources]in the IoT Greengrass V2 Developer Guide.
	//
	// [Tag your resources]: https://docs.aws.amazon.com/greengrass/v2/developerguide/tag-resources.html
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateDeploymentOutput struct {

	// The ID of the deployment.
	DeploymentId *string

	// The [ARN] of the IoT job that applies the deployment to target devices.
	//
	// [ARN]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	IotJobArn *string

	// The ID of the IoT job that applies the deployment to target devices.
	IotJobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDeploymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDeployment"); err != nil {
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
	if err = addIdempotencyToken_opCreateDeploymentMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateDeploymentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDeployment(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateDeployment struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDeployment) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDeployment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDeploymentInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDeploymentInput ")
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
func addIdempotencyToken_opCreateDeploymentMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateDeployment{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDeployment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDeployment",
	}
}
