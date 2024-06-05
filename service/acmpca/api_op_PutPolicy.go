// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attaches a resource-based policy to a private CA.
//
// A policy can also be applied by sharing a private CA through Amazon Web
// Services Resource Access Manager (RAM). For more information, see [Attach a Policy for Cross-Account Access].
//
// The policy can be displayed with [GetPolicy] and removed with [DeletePolicy].
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
// [Using a Resource Based Policy with Amazon Web Services Private CA]: https://docs.aws.amazon.com/privateca/latest/userguide/pca-rbp.html
// [Using a Service Linked Role with ACM]: https://docs.aws.amazon.com/acm/latest/userguide/acm-slr.html
// [Attach a Policy for Cross-Account Access]: https://docs.aws.amazon.com/privateca/latest/userguide/pca-ram.html
// [DeletePolicy]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_DeletePolicy.html
// [GetPolicy]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_GetPolicy.html
func (c *Client) PutPolicy(ctx context.Context, params *PutPolicyInput, optFns ...func(*Options)) (*PutPolicyOutput, error) {
	if params == nil {
		params = &PutPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutPolicy", params, optFns, c.addOperationPutPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutPolicyInput struct {

	// The path and file name of a JSON-formatted IAM policy to attach to the
	// specified private CA resource. If this policy does not contain all required
	// statements or if it includes any statement that is not allowed, the PutPolicy
	// action returns an InvalidPolicyException . For information about IAM policy and
	// statement structure, see [Overview of JSON Policies].
	//
	// [Overview of JSON Policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policies-json
	//
	// This member is required.
	Policy *string

	// The Amazon Resource Number (ARN) of the private CA to associate with the
	// policy. The ARN of the CA can be found by calling the [ListCertificateAuthorities]action.
	//
	// [ListCertificateAuthorities]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListCertificateAuthorities.html
	//
	// This member is required.
	ResourceArn *string

	noSmithyDocumentSerde
}

type PutPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutPolicy"); err != nil {
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
	if err = addOpPutPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutPolicy",
	}
}
