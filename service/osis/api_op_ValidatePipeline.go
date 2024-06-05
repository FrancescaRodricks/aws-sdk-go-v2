// Code generated by smithy-go-codegen DO NOT EDIT.

package osis

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/osis/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Checks whether an OpenSearch Ingestion pipeline configuration is valid prior to
// creation. For more information, see [Creating Amazon OpenSearch Ingestion pipelines].
//
// [Creating Amazon OpenSearch Ingestion pipelines]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/creating-pipeline.html
func (c *Client) ValidatePipeline(ctx context.Context, params *ValidatePipelineInput, optFns ...func(*Options)) (*ValidatePipelineOutput, error) {
	if params == nil {
		params = &ValidatePipelineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ValidatePipeline", params, optFns, c.addOperationValidatePipelineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ValidatePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ValidatePipelineInput struct {

	// The pipeline configuration in YAML format. The command accepts the pipeline
	// configuration as a string or within a .yaml file. If you provide the
	// configuration as a string, each new line must be escaped with \n .
	//
	// This member is required.
	PipelineConfigurationBody *string

	noSmithyDocumentSerde
}

type ValidatePipelineOutput struct {

	// A list of errors if the configuration is invalid.
	Errors []types.ValidationMessage

	// A boolean indicating whether or not the pipeline configuration is valid.
	IsValid *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationValidatePipelineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpValidatePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpValidatePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ValidatePipeline"); err != nil {
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
	if err = addOpValidatePipelineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opValidatePipeline(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opValidatePipeline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ValidatePipeline",
	}
}
