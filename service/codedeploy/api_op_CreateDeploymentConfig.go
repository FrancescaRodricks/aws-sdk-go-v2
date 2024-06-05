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

// Creates a deployment configuration.
func (c *Client) CreateDeploymentConfig(ctx context.Context, params *CreateDeploymentConfigInput, optFns ...func(*Options)) (*CreateDeploymentConfigOutput, error) {
	if params == nil {
		params = &CreateDeploymentConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDeploymentConfig", params, optFns, c.addOperationCreateDeploymentConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDeploymentConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateDeploymentConfig operation.
type CreateDeploymentConfigInput struct {

	// The name of the deployment configuration to create.
	//
	// This member is required.
	DeploymentConfigName *string

	// The destination platform type for the deployment ( Lambda , Server , or ECS ).
	ComputePlatform types.ComputePlatform

	// The minimum number of healthy instances that should be available at any time
	// during the deployment. There are two parameters expected in the input: type and
	// value.
	//
	// The type parameter takes either of the following values:
	//
	//   - HOST_COUNT: The value parameter represents the minimum number of healthy
	//   instances as an absolute value.
	//
	//   - FLEET_PERCENT: The value parameter represents the minimum number of healthy
	//   instances as a percentage of the total number of instances in the deployment. If
	//   you specify FLEET_PERCENT, at the start of the deployment, CodeDeploy converts
	//   the percentage to the equivalent number of instances and rounds up fractional
	//   instances.
	//
	// The value parameter takes an integer.
	//
	// For example, to set a minimum of 95% healthy instance, specify a type of
	// FLEET_PERCENT and a value of 95.
	MinimumHealthyHosts *types.MinimumHealthyHosts

	// The configuration that specifies how the deployment traffic is routed.
	TrafficRoutingConfig *types.TrafficRoutingConfig

	// Configure the ZonalConfig object if you want CodeDeploy to deploy your
	// application to one [Availability Zone]at a time, within an Amazon Web Services Region.
	//
	// For more information about the zonal configuration feature, see [zonal configuration] in the
	// CodeDeploy User Guide.
	//
	// [zonal configuration]: https://docs.aws.amazon.com/codedeploy/latest/userguide/deployment-configurations-create.html#zonal-config
	// [Availability Zone]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-availability-zones
	ZonalConfig *types.ZonalConfig

	noSmithyDocumentSerde
}

// Represents the output of a CreateDeploymentConfig operation.
type CreateDeploymentConfigOutput struct {

	// A unique deployment configuration ID.
	DeploymentConfigId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDeploymentConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDeploymentConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDeploymentConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDeploymentConfig"); err != nil {
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
	if err = addOpCreateDeploymentConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDeploymentConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDeploymentConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDeploymentConfig",
	}
}
