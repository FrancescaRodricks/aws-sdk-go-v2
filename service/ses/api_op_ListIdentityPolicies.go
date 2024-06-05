// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of sending authorization policies that are attached to the given
// identity (an email address or a domain). This operation returns only a list. To
// get the actual policy content, use GetIdentityPolicies .
//
// This operation is for the identity owner only. If you have not verified the
// identity, it returns an error.
//
// Sending authorization is a feature that enables an identity owner to authorize
// other senders to use its identities. For information about using sending
// authorization, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/sending-authorization.html
func (c *Client) ListIdentityPolicies(ctx context.Context, params *ListIdentityPoliciesInput, optFns ...func(*Options)) (*ListIdentityPoliciesOutput, error) {
	if params == nil {
		params = &ListIdentityPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIdentityPolicies", params, optFns, c.addOperationListIdentityPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIdentityPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to return a list of sending authorization policies that
// are attached to an identity. Sending authorization is an Amazon SES feature that
// enables you to authorize other senders to use your identities. For information,
// see the [Amazon SES Developer Guide].
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/sending-authorization.html
type ListIdentityPoliciesInput struct {

	// The identity that is associated with the policy for which the policies are
	// listed. You can specify an identity by using its name or by using its Amazon
	// Resource Name (ARN). Examples: user@example.com , example.com ,
	// arn:aws:ses:us-east-1:123456789012:identity/example.com .
	//
	// To successfully call this operation, you must own the identity.
	//
	// This member is required.
	Identity *string

	noSmithyDocumentSerde
}

// A list of names of sending authorization policies that apply to an identity.
type ListIdentityPoliciesOutput struct {

	// A list of names of policies that apply to the specified identity.
	//
	// This member is required.
	PolicyNames []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListIdentityPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListIdentityPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListIdentityPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListIdentityPolicies"); err != nil {
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
	if err = addOpListIdentityPoliciesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListIdentityPolicies(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListIdentityPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListIdentityPolicies",
	}
}
