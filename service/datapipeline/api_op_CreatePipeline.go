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

// Creates a new, empty pipeline. Use PutPipelineDefinition to populate the pipeline.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.CreatePipeline Content-Length: 91 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"name": "myPipeline", "uniqueId": "123456789", "description": "This is my
// first pipeline"}
//
// HTTP/1.1 200 x-amzn-RequestId: b16911ce-0774-11e2-af6f-6bc7a6be60d9
// Content-Type: application/x-amz-json-1.1 Content-Length: 40 Date: Mon, 12 Nov
// 2012 17:50:53 GMT
//
// {"pipelineId": "df-06372391ZG65EXAMPLE"}
func (c *Client) CreatePipeline(ctx context.Context, params *CreatePipelineInput, optFns ...func(*Options)) (*CreatePipelineOutput, error) {
	if params == nil {
		params = &CreatePipelineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePipeline", params, optFns, c.addOperationCreatePipelineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreatePipeline.
type CreatePipelineInput struct {

	// The name for the pipeline. You can use the same name for multiple pipelines
	// associated with your AWS account, because AWS Data Pipeline assigns each
	// pipeline a unique pipeline identifier.
	//
	// This member is required.
	Name *string

	// A unique identifier. This identifier is not the same as the pipeline identifier
	// assigned by AWS Data Pipeline. You are responsible for defining the format and
	// ensuring the uniqueness of this identifier. You use this parameter to ensure
	// idempotency during repeated calls to CreatePipeline . For example, if the first
	// call to CreatePipeline does not succeed, you can pass in the same unique
	// identifier and pipeline name combination on a subsequent call to CreatePipeline
	// . CreatePipeline ensures that if a pipeline already exists with the same name
	// and unique identifier, a new pipeline is not created. Instead, you'll receive
	// the pipeline identifier from the previous attempt. The uniqueness of the name
	// and unique identifier combination is scoped to the AWS account or IAM user
	// credentials.
	//
	// This member is required.
	UniqueId *string

	// The description for the pipeline.
	Description *string

	// A list of tags to associate with the pipeline at creation. Tags let you control
	// access to pipelines. For more information, see [Controlling User Access to Pipelines]in the AWS Data Pipeline
	// Developer Guide.
	//
	// [Controlling User Access to Pipelines]: http://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-control-access.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Contains the output of CreatePipeline.
type CreatePipelineOutput struct {

	// The ID that AWS Data Pipeline assigns the newly created pipeline. For example,
	// df-06372391ZG65EXAMPLE .
	//
	// This member is required.
	PipelineId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePipelineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePipeline"); err != nil {
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
	if err = addOpCreatePipelineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePipeline(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePipeline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePipeline",
	}
}
