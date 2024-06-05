// Code generated by smithy-go-codegen DO NOT EDIT.

package ivschat

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ivschat/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates a specified logging configuration.
func (c *Client) UpdateLoggingConfiguration(ctx context.Context, params *UpdateLoggingConfigurationInput, optFns ...func(*Options)) (*UpdateLoggingConfigurationOutput, error) {
	if params == nil {
		params = &UpdateLoggingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLoggingConfiguration", params, optFns, c.addOperationUpdateLoggingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLoggingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLoggingConfigurationInput struct {

	// Identifier of the logging configuration to be updated.
	//
	// This member is required.
	Identifier *string

	// A complex type that contains a destination configuration for where chat content
	// will be logged. There can be only one type of destination ( cloudWatchLogs ,
	// firehose , or s3 ) in a destinationConfiguration .
	DestinationConfiguration types.DestinationConfiguration

	// Logging-configuration name. The value does not need to be unique.
	Name *string

	noSmithyDocumentSerde
}

type UpdateLoggingConfigurationOutput struct {

	// Logging-configuration ARN, from the request (if identifier was an ARN).
	Arn *string

	// Time when the logging configuration was created. This is an ISO 8601 timestamp;
	// note that this is returned as a string.
	CreateTime *time.Time

	// A complex type that contains a destination configuration for where chat content
	// will be logged, from the request. There is only one type of destination (
	// cloudWatchLogs , firehose , or s3 ) in a destinationConfiguration .
	DestinationConfiguration types.DestinationConfiguration

	// Logging-configuration ID, generated by the system. This is a relative
	// identifier, the part of the ARN that uniquely identifies the room.
	Id *string

	// Logging-configuration name, from the request (if specified).
	Name *string

	// The state of the logging configuration. When the state is ACTIVE , the
	// configuration is ready to log chat content.
	State types.UpdateLoggingConfigurationState

	// Tags attached to the resource. Array of maps, each of the form string:string
	// (key:value) .
	Tags map[string]string

	// Time of the logging configuration’s last update. This is an ISO 8601 timestamp;
	// note that this is returned as a string.
	UpdateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLoggingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateLoggingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateLoggingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateLoggingConfiguration"); err != nil {
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
	if err = addOpUpdateLoggingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLoggingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLoggingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateLoggingConfiguration",
	}
}
