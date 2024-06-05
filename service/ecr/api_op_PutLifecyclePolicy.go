// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates the lifecycle policy for the specified repository. For more
// information, see [Lifecycle policy template].
//
// [Lifecycle policy template]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html
func (c *Client) PutLifecyclePolicy(ctx context.Context, params *PutLifecyclePolicyInput, optFns ...func(*Options)) (*PutLifecyclePolicyOutput, error) {
	if params == nil {
		params = &PutLifecyclePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutLifecyclePolicy", params, optFns, c.addOperationPutLifecyclePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutLifecyclePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLifecyclePolicyInput struct {

	// The JSON repository policy text to apply to the repository.
	//
	// This member is required.
	LifecyclePolicyText *string

	// The name of the repository to receive the policy.
	//
	// This member is required.
	RepositoryName *string

	// The Amazon Web Services account ID associated with the registry that contains
	// the repository. If you do  not specify a registry, the default registry is
	// assumed.
	RegistryId *string

	noSmithyDocumentSerde
}

type PutLifecyclePolicyOutput struct {

	// The JSON repository policy text.
	LifecyclePolicyText *string

	// The registry ID associated with the request.
	RegistryId *string

	// The repository name associated with the request.
	RepositoryName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutLifecyclePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutLifecyclePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutLifecyclePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutLifecyclePolicy"); err != nil {
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
	if err = addOpPutLifecyclePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutLifecyclePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutLifecyclePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutLifecyclePolicy",
	}
}
