// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Validates the specified pipeline definition to ensure that it is well formed
// and can be run without error.
//
// Example 1 This example sets an valid pipeline configuration and returns
// success.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.ValidatePipelineDefinition Content-Length: 936 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"pipelineId": "df-06372391ZG65EXAMPLE", "pipelineObjects": [ {"id": "Default",
// "name": "Default", "fields": [ {"key": "workerGroup", "stringValue":
// "MyworkerGroup"} ] }, {"id": "Schedule", "name": "Schedule", "fields": [ {"key":
// "startDateTime", "stringValue": "2012-09-25T17:00:00"}, {"key": "type",
// "stringValue": "Schedule"}, {"key": "period", "stringValue": "1 hour"}, {"key":
// "endDateTime", "stringValue": "2012-09-25T18:00:00"} ] }, {"id": "SayHello",
// "name": "SayHello", "fields": [ {"key": "type", "stringValue":
// "ShellCommandActivity"}, {"key": "command", "stringValue": "echo hello"},
// {"key": "parent", "refValue": "Default"}, {"key": "schedule", "refValue":
// "Schedule"}
//
// ] } ] }
//
// x-amzn-RequestId: 92c9f347-0776-11e2-8a14-21bb8a1f50ef Content-Type:
// application/x-amz-json-1.1 Content-Length: 18 Date: Mon, 12 Nov 2012 17:50:53
// GMT
//
// {"errored": false}
//
// Example 2 This example sets an invalid pipeline configuration and returns the
// associated set of validation errors.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.ValidatePipelineDefinition Content-Length: 903 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"pipelineId": "df-06372391ZG65EXAMPLE", "pipelineObjects": [ {"id": "Default",
// "name": "Default", "fields": [ {"key": "workerGroup", "stringValue":
// "MyworkerGroup"} ] }, {"id": "Schedule", "name": "Schedule", "fields": [ {"key":
// "startDateTime", "stringValue": "bad-time"}, {"key": "type", "stringValue":
// "Schedule"}, {"key": "period", "stringValue": "1 hour"}, {"key": "endDateTime",
// "stringValue": "2012-09-25T18:00:00"} ] }, {"id": "SayHello", "name":
// "SayHello", "fields": [ {"key": "type", "stringValue": "ShellCommandActivity"},
// {"key": "command", "stringValue": "echo hello"}, {"key": "parent", "refValue":
// "Default"}, {"key": "schedule", "refValue": "Schedule"}
//
// ] } ] }
//
// x-amzn-RequestId: 496a1f5a-0e6a-11e2-a61c-bd6312c92ddd Content-Type:
// application/x-amz-json-1.1 Content-Length: 278 Date: Mon, 12 Nov 2012 17:50:53
// GMT
//
// {"errored": true, "validationErrors": [ {"errors": ["INVALID_FIELD_VALUE:
// 'startDateTime' value must be a literal datetime value."], "id": "Schedule"} ] }
func (c *Client) ValidatePipelineDefinition(ctx context.Context, params *ValidatePipelineDefinitionInput, optFns ...func(*Options)) (*ValidatePipelineDefinitionOutput, error) {
	if params == nil {
		params = &ValidatePipelineDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ValidatePipelineDefinition", params, optFns, c.addOperationValidatePipelineDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ValidatePipelineDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ValidatePipelineDefinition.
type ValidatePipelineDefinitionInput struct {

	// The ID of the pipeline.
	//
	// This member is required.
	PipelineId *string

	// The objects that define the pipeline changes to validate against the pipeline.
	//
	// This member is required.
	PipelineObjects []types.PipelineObject

	// The parameter objects used with the pipeline.
	ParameterObjects []types.ParameterObject

	// The parameter values used with the pipeline.
	ParameterValues []types.ParameterValue

	noSmithyDocumentSerde
}

// Contains the output of ValidatePipelineDefinition.
type ValidatePipelineDefinitionOutput struct {

	// Indicates whether there were validation errors.
	//
	// This member is required.
	Errored bool

	// Any validation errors that were found.
	ValidationErrors []types.ValidationError

	// Any validation warnings that were found.
	ValidationWarnings []types.ValidationWarning

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationValidatePipelineDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpValidatePipelineDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpValidatePipelineDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ValidatePipelineDefinition"); err != nil {
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
	if err = addOpValidatePipelineDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opValidatePipelineDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opValidatePipelineDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ValidatePipelineDefinition",
	}
}
