// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Simulates the results of running a pipeline activity on a message payload.
func (c *Client) RunPipelineActivity(ctx context.Context, params *RunPipelineActivityInput, optFns ...func(*Options)) (*RunPipelineActivityOutput, error) {
	if params == nil {
		params = &RunPipelineActivityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RunPipelineActivity", params, optFns, c.addOperationRunPipelineActivityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RunPipelineActivityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RunPipelineActivityInput struct {

	// The sample message payloads on which the pipeline activity is run.
	//
	// This member is required.
	Payloads [][]byte

	// The pipeline activity that is run. This must not be a channel activity or a
	// data store activity because these activities are used in a pipeline only to load
	// the original message and to store the (possibly) transformed message. If a
	// Lambda activity is specified, only short-running Lambda functions (those with a
	// timeout of less than 30 seconds or less) can be used.
	//
	// This member is required.
	PipelineActivity *types.PipelineActivity

	noSmithyDocumentSerde
}

type RunPipelineActivityOutput struct {

	// In case the pipeline activity fails, the log message that is generated.
	LogResult *string

	// The enriched or transformed sample message payloads as base64-encoded strings.
	// (The results of running the pipeline activity on each input sample message
	// payload, encoded in base64.)
	Payloads [][]byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRunPipelineActivityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRunPipelineActivity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRunPipelineActivity{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RunPipelineActivity"); err != nil {
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
	if err = addOpRunPipelineActivityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRunPipelineActivity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRunPipelineActivity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RunPipelineActivity",
	}
}
