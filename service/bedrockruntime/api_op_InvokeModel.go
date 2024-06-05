// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockruntime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Invokes the specified Amazon Bedrock model to run inference using the prompt
// and inference parameters provided in the request body. You use model inference
// to generate text, images, and embeddings.
//
// For example code, see Invoke model code examples in the Amazon Bedrock User
// Guide.
//
// This operation requires permission for the bedrock:InvokeModel action.
func (c *Client) InvokeModel(ctx context.Context, params *InvokeModelInput, optFns ...func(*Options)) (*InvokeModelOutput, error) {
	if params == nil {
		params = &InvokeModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InvokeModel", params, optFns, c.addOperationInvokeModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InvokeModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InvokeModelInput struct {

	// The prompt and inference parameters in the format specified in the contentType
	// in the header. To see the format and content of the request and response bodies
	// for different models, refer to [Inference parameters]. For more information, see [Run inference] in the Bedrock User
	// Guide.
	//
	// [Inference parameters]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-parameters.html
	// [Run inference]: https://docs.aws.amazon.com/bedrock/latest/userguide/api-methods-run.html
	//
	// This member is required.
	Body []byte

	// The unique identifier of the model to invoke to run inference.
	//
	// The modelId to provide depends on the type of model that you use:
	//
	//   - If you use a base model, specify the model ID or its ARN. For a list of
	//   model IDs for base models, see [Amazon Bedrock base model IDs (on-demand throughput)]in the Amazon Bedrock User Guide.
	//
	//   - If you use a provisioned model, specify the ARN of the Provisioned
	//   Throughput. For more information, see [Run inference using a Provisioned Throughput]in the Amazon Bedrock User Guide.
	//
	//   - If you use a custom model, first purchase Provisioned Throughput for it.
	//   Then specify the ARN of the resulting provisioned model. For more information,
	//   see [Use a custom model in Amazon Bedrock]in the Amazon Bedrock User Guide.
	//
	// [Run inference using a Provisioned Throughput]: https://docs.aws.amazon.com/bedrock/latest/userguide/prov-thru-use.html
	// [Use a custom model in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-customization-use.html
	// [Amazon Bedrock base model IDs (on-demand throughput)]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-ids.html#model-ids-arns
	//
	// This member is required.
	ModelId *string

	// The desired MIME type of the inference body in the response. The default value
	// is application/json .
	Accept *string

	// The MIME type of the input data in the request. The default value is
	// application/json .
	ContentType *string

	// The unique identifier of the guardrail that you want to use. If you don't
	// provide a value, no guardrail is applied to the invocation.
	//
	// An error will be thrown in the following situations.
	//
	//   - You don't provide a guardrail identifier but you specify the
	//   amazon-bedrock-guardrailConfig field in the request body.
	//
	//   - You enable the guardrail but the contentType isn't application/json .
	//
	//   - You provide a guardrail identifier, but guardrailVersion isn't specified.
	GuardrailIdentifier *string

	// The version number for the guardrail. The value can also be DRAFT .
	GuardrailVersion *string

	// Specifies whether to enable or disable the Bedrock trace. If enabled, you can
	// see the full Bedrock trace.
	Trace types.Trace

	noSmithyDocumentSerde
}

type InvokeModelOutput struct {

	// Inference response from the model in the format specified in the contentType
	// header. To see the format and content of the request and response bodies for
	// different models, refer to [Inference parameters].
	//
	// [Inference parameters]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-parameters.html
	//
	// This member is required.
	Body []byte

	// The MIME type of the inference result.
	//
	// This member is required.
	ContentType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationInvokeModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInvokeModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInvokeModel{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "InvokeModel"); err != nil {
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
	if err = addOpInvokeModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInvokeModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInvokeModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InvokeModel",
	}
}
