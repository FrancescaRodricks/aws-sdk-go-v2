// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the SAML provider metadocument that was uploaded when the IAM SAML
// provider resource object was created or updated.
//
// This operation requires [Signature Version 4].
//
// [Signature Version 4]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
func (c *Client) GetSAMLProvider(ctx context.Context, params *GetSAMLProviderInput, optFns ...func(*Options)) (*GetSAMLProviderOutput, error) {
	if params == nil {
		params = &GetSAMLProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSAMLProvider", params, optFns, c.addOperationGetSAMLProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSAMLProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSAMLProviderInput struct {

	// The Amazon Resource Name (ARN) of the SAML provider resource object in IAM to
	// get information about.
	//
	// For more information about ARNs, see [Amazon Resource Names (ARNs)] in the Amazon Web Services General
	// Reference.
	//
	// [Amazon Resource Names (ARNs)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	//
	// This member is required.
	SAMLProviderArn *string

	noSmithyDocumentSerde
}

// Contains the response to a successful GetSAMLProvider request.
type GetSAMLProviderOutput struct {

	// The date and time when the SAML provider was created.
	CreateDate *time.Time

	// The XML metadata document that includes information about an identity provider.
	SAMLMetadataDocument *string

	// A list of tags that are attached to the specified IAM SAML provider. The
	// returned list of tags is sorted by tag key. For more information about tagging,
	// see [Tagging IAM resources]in the IAM User Guide.
	//
	// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
	Tags []types.Tag

	// The expiration date and time for the SAML provider.
	ValidUntil *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSAMLProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetSAMLProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetSAMLProvider{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSAMLProvider"); err != nil {
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
	if err = addOpGetSAMLProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSAMLProvider(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSAMLProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSAMLProvider",
	}
}
