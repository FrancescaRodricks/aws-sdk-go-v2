// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Applies a policy to an image. We recommend that you call the RAM API [CreateResourceShare] to share
// resources. If you call the Image Builder API PutImagePolicy , you must also call
// the RAM API [PromoteResourceShareCreatedFromPolicy]in order for the resource to be visible to all principals with whom
// the resource is shared.
//
// [PromoteResourceShareCreatedFromPolicy]: https://docs.aws.amazon.com/ram/latest/APIReference/API_PromoteResourceShareCreatedFromPolicy.html
// [CreateResourceShare]: https://docs.aws.amazon.com/ram/latest/APIReference/API_CreateResourceShare.html
func (c *Client) PutImagePolicy(ctx context.Context, params *PutImagePolicyInput, optFns ...func(*Options)) (*PutImagePolicyOutput, error) {
	if params == nil {
		params = &PutImagePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutImagePolicy", params, optFns, c.addOperationPutImagePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutImagePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutImagePolicyInput struct {

	// The Amazon Resource Name (ARN) of the image that this policy should be applied
	// to.
	//
	// This member is required.
	ImageArn *string

	// The policy to apply.
	//
	// This member is required.
	Policy *string

	noSmithyDocumentSerde
}

type PutImagePolicyOutput struct {

	// The Amazon Resource Name (ARN) of the image that this policy was applied to.
	ImageArn *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutImagePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutImagePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutImagePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutImagePolicy"); err != nil {
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
	if err = addOpPutImagePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutImagePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutImagePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutImagePolicy",
	}
}
