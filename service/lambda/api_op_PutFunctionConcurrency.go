// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the maximum number of simultaneous executions for a function, and reserves
// capacity for that concurrency level.
//
// Concurrency settings apply to the function as a whole, including all published
// versions and the unpublished version. Reserving concurrency both ensures that
// your function has capacity to process the specified number of events
// simultaneously, and prevents it from scaling beyond that level. Use GetFunctionto see the
// current setting for a function.
//
// Use GetAccountSettings to see your Regional concurrency limit. You can reserve concurrency for as
// many functions as you like, as long as you leave at least 100 simultaneous
// executions unreserved for functions that aren't configured with a per-function
// limit. For more information, see [Lambda function scaling].
//
// [Lambda function scaling]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-scaling.html
func (c *Client) PutFunctionConcurrency(ctx context.Context, params *PutFunctionConcurrencyInput, optFns ...func(*Options)) (*PutFunctionConcurrencyOutput, error) {
	if params == nil {
		params = &PutFunctionConcurrencyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutFunctionConcurrency", params, optFns, c.addOperationPutFunctionConcurrencyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutFunctionConcurrencyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutFunctionConcurrencyInput struct {

	// The name or ARN of the Lambda function.
	//
	// Name formats
	//
	//   - Function name – my-function .
	//
	//   - Function ARN – arn:aws:lambda:us-west-2:123456789012:function:my-function .
	//
	//   - Partial ARN – 123456789012:function:my-function .
	//
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// The number of simultaneous executions to reserve for the function.
	//
	// This member is required.
	ReservedConcurrentExecutions *int32

	noSmithyDocumentSerde
}

type PutFunctionConcurrencyOutput struct {

	// The number of concurrent executions that are reserved for this function. For
	// more information, see [Managing Lambda reserved concurrency].
	//
	// [Managing Lambda reserved concurrency]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-concurrency.html
	ReservedConcurrentExecutions *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutFunctionConcurrencyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutFunctionConcurrency{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutFunctionConcurrency{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutFunctionConcurrency"); err != nil {
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
	if err = addOpPutFunctionConcurrencyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutFunctionConcurrency(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutFunctionConcurrency(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutFunctionConcurrency",
	}
}
