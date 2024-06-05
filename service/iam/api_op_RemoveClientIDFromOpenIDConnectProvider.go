// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified client ID (also known as audience) from the list of
// client IDs registered for the specified IAM OpenID Connect (OIDC) provider
// resource object.
//
// This operation is idempotent; it does not fail or return an error if you try to
// remove a client ID that does not exist.
func (c *Client) RemoveClientIDFromOpenIDConnectProvider(ctx context.Context, params *RemoveClientIDFromOpenIDConnectProviderInput, optFns ...func(*Options)) (*RemoveClientIDFromOpenIDConnectProviderOutput, error) {
	if params == nil {
		params = &RemoveClientIDFromOpenIDConnectProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveClientIDFromOpenIDConnectProvider", params, optFns, c.addOperationRemoveClientIDFromOpenIDConnectProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveClientIDFromOpenIDConnectProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveClientIDFromOpenIDConnectProviderInput struct {

	// The client ID (also known as audience) to remove from the IAM OIDC provider
	// resource. For more information about client IDs, see CreateOpenIDConnectProvider.
	//
	// This member is required.
	ClientID *string

	// The Amazon Resource Name (ARN) of the IAM OIDC provider resource to remove the
	// client ID from. You can get a list of OIDC provider ARNs by using the ListOpenIDConnectProvidersoperation.
	//
	// For more information about ARNs, see [Amazon Resource Names (ARNs)] in the Amazon Web Services General
	// Reference.
	//
	// [Amazon Resource Names (ARNs)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	//
	// This member is required.
	OpenIDConnectProviderArn *string

	noSmithyDocumentSerde
}

type RemoveClientIDFromOpenIDConnectProviderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRemoveClientIDFromOpenIDConnectProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRemoveClientIDFromOpenIDConnectProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRemoveClientIDFromOpenIDConnectProvider{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RemoveClientIDFromOpenIDConnectProvider"); err != nil {
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
	if err = addOpRemoveClientIDFromOpenIDConnectProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveClientIDFromOpenIDConnectProvider(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRemoveClientIDFromOpenIDConnectProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RemoveClientIDFromOpenIDConnectProvider",
	}
}
