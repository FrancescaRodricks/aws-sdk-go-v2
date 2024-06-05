// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Restart pipelines in one channel that is currently running.
func (c *Client) RestartChannelPipelines(ctx context.Context, params *RestartChannelPipelinesInput, optFns ...func(*Options)) (*RestartChannelPipelinesOutput, error) {
	if params == nil {
		params = &RestartChannelPipelinesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestartChannelPipelines", params, optFns, c.addOperationRestartChannelPipelinesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestartChannelPipelinesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Pipelines to restart.
type RestartChannelPipelinesInput struct {

	// ID of channel
	//
	// This member is required.
	ChannelId *string

	// An array of pipelines to restart in this channel. Format PIPELINE_0 or
	// PIPELINE_1.
	PipelineIds []types.ChannelPipelineIdToRestart

	noSmithyDocumentSerde
}

// Placeholder documentation for RestartChannelPipelinesResponse
type RestartChannelPipelinesOutput struct {

	// The unique arn of the channel.
	Arn *string

	// Specification of CDI inputs for this channel
	CdiInputSpecification *types.CdiInputSpecification

	// The class for this channel. STANDARD for a channel with two pipelines or
	// SINGLE_PIPELINE for a channel with one pipeline.
	ChannelClass types.ChannelClass

	// A list of destinations of the channel. For UDP outputs, there is one
	// destination per output. For other types (HLS, for example), there is one
	// destination per packager.
	Destinations []types.OutputDestination

	// The endpoints where outgoing connections initiate from
	EgressEndpoints []types.ChannelEgressEndpoint

	// Encoder Settings
	EncoderSettings *types.EncoderSettings

	// The unique id of the channel.
	Id *string

	// List of input attachments for channel.
	InputAttachments []types.InputAttachment

	// Specification of network and file inputs for this channel
	InputSpecification *types.InputSpecification

	// The log level being written to CloudWatch Logs.
	LogLevel types.LogLevel

	// Maintenance settings for this channel.
	Maintenance *types.MaintenanceStatus

	// The time in milliseconds by when the PVRE restart must occur.
	MaintenanceStatus *string

	// The name of the channel. (user-mutable)
	Name *string

	// Runtime details for the pipelines of a running channel.
	PipelineDetails []types.PipelineDetail

	// The number of currently healthy pipelines.
	PipelinesRunningCount *int32

	// The Amazon Resource Name (ARN) of the role assumed when running the Channel.
	RoleArn *string

	// Placeholder documentation for ChannelState
	State types.ChannelState

	// A collection of key-value pairs.
	Tags map[string]string

	// Settings for VPC output
	Vpc *types.VpcOutputSettingsDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestartChannelPipelinesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRestartChannelPipelines{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRestartChannelPipelines{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RestartChannelPipelines"); err != nil {
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
	if err = addOpRestartChannelPipelinesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestartChannelPipelines(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestartChannelPipelines(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RestartChannelPipelines",
	}
}
