// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Applies a policy to a component. We recommend that you call the RAM API [CreateResourceShare] to
// share resources. If you call the Image Builder API PutComponentPolicy , you must
// also call the RAM API [PromoteResourceShareCreatedFromPolicy]in order for the resource to be visible to all principals
// with whom the resource is shared.
//
// [PromoteResourceShareCreatedFromPolicy]: https://docs.aws.amazon.com/ram/latest/APIReference/API_PromoteResourceShareCreatedFromPolicy.html
// [CreateResourceShare]: https://docs.aws.amazon.com/ram/latest/APIReference/API_CreateResourceShare.html
func (c *Client) PutComponentPolicy(ctx context.Context, params *PutComponentPolicyInput, optFns ...func(*Options)) (*PutComponentPolicyOutput, error) {
	if params == nil {
		params = &PutComponentPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutComponentPolicy", params, optFns, c.addOperationPutComponentPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutComponentPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutComponentPolicyInput struct {

	// The Amazon Resource Name (ARN) of the component that this policy should be
	// applied to.
	//
	// This member is required.
	ComponentArn *string

	// The policy to apply.
	//
	// This member is required.
	Policy *string

	noSmithyDocumentSerde
}

type PutComponentPolicyOutput struct {

	// The Amazon Resource Name (ARN) of the component that this policy was applied to.
	ComponentArn *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutComponentPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutComponentPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutComponentPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutComponentPolicy"); err != nil {
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
	if err = addOpPutComponentPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutComponentPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutComponentPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutComponentPolicy",
	}
}
