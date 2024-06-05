// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the resource-based policy attached to a private CA. Deletion will
// remove any access that the policy has granted. If there is no policy attached to
// the private CA, this action will return successful.
//
// If you delete a policy that was applied through Amazon Web Services Resource
// Access Manager (RAM), the CA will be removed from all shares in which it was
// included.
//
// The Certificate Manager Service Linked Role that the policy supports is not
// affected when you delete the policy.
//
// The current policy can be shown with [GetPolicy] and updated with [PutPolicy].
//
// About Policies
//
//   - A policy grants access on a private CA to an Amazon Web Services customer
//     account, to Amazon Web Services Organizations, or to an Amazon Web Services
//     Organizations unit. Policies are under the control of a CA administrator. For
//     more information, see [Using a Resource Based Policy with Amazon Web Services Private CA].
//
//   - A policy permits a user of Certificate Manager (ACM) to issue ACM
//     certificates signed by a CA in another account.
//
//   - For ACM to manage automatic renewal of these certificates, the ACM user
//     must configure a Service Linked Role (SLR). The SLR allows the ACM service to
//     assume the identity of the user, subject to confirmation against the Amazon Web
//     Services Private CA policy. For more information, see [Using a Service Linked Role with ACM].
//
//   - Updates made in Amazon Web Services Resource Manager (RAM) are reflected in
//     policies. For more information, see [Attach a Policy for Cross-Account Access].
//
// [PutPolicy]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_PutPolicy.html
// [Using a Resource Based Policy with Amazon Web Services Private CA]: https://docs.aws.amazon.com/privateca/latest/userguide/pca-rbp.html
// [Using a Service Linked Role with ACM]: https://docs.aws.amazon.com/acm/latest/userguide/acm-slr.html
// [Attach a Policy for Cross-Account Access]: https://docs.aws.amazon.com/privateca/latest/userguide/pca-ram.html
// [GetPolicy]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_GetPolicy.html
func (c *Client) DeletePolicy(ctx context.Context, params *DeletePolicyInput, optFns ...func(*Options)) (*DeletePolicyOutput, error) {
	if params == nil {
		params = &DeletePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePolicy", params, optFns, c.addOperationDeletePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePolicyInput struct {

	// The Amazon Resource Number (ARN) of the private CA that will have its policy
	// deleted. You can find the CA's ARN by calling the [ListCertificateAuthorities]action. The ARN value must
	// have the form
	// arn:aws:acm-pca:region:account:certificate-authority/01234567-89ab-cdef-0123-0123456789ab
	// .
	//
	// [ListCertificateAuthorities]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListCertificateAuthorities.html
	//
	// This member is required.
	ResourceArn *string

	noSmithyDocumentSerde
}

type DeletePolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeletePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeletePolicy"); err != nil {
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
	if err = addOpDeletePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeletePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeletePolicy",
	}
}
