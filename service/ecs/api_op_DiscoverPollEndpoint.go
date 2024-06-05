// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action is only used by the Amazon ECS agent, and it is not intended for
// use outside of the agent.
//
// Returns an endpoint for the Amazon ECS agent to poll for updates.
func (c *Client) DiscoverPollEndpoint(ctx context.Context, params *DiscoverPollEndpointInput, optFns ...func(*Options)) (*DiscoverPollEndpointOutput, error) {
	if params == nil {
		params = &DiscoverPollEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DiscoverPollEndpoint", params, optFns, c.addOperationDiscoverPollEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DiscoverPollEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DiscoverPollEndpointInput struct {

	// The short name or full Amazon Resource Name (ARN) of the cluster that the
	// container instance belongs to.
	Cluster *string

	// The container instance ID or full ARN of the container instance. For more
	// information about the ARN format, see [Amazon Resource Name (ARN)]in the Amazon ECS Developer Guide.
	//
	// [Amazon Resource Name (ARN)]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids
	ContainerInstance *string

	noSmithyDocumentSerde
}

type DiscoverPollEndpointOutput struct {

	// The endpoint for the Amazon ECS agent to poll.
	Endpoint *string

	// The endpoint for the Amazon ECS agent to poll for Service Connect
	// configuration. For more information, see [Service Connect]in the Amazon Elastic Container
	// Service Developer Guide.
	//
	// [Service Connect]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html
	ServiceConnectEndpoint *string

	// The telemetry endpoint for the Amazon ECS agent.
	TelemetryEndpoint *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDiscoverPollEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDiscoverPollEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDiscoverPollEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DiscoverPollEndpoint"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDiscoverPollEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDiscoverPollEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DiscoverPollEndpoint",
	}
}
