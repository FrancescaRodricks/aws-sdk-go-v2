// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Revokes permissions on a private CA granted to the Certificate Manager (ACM)
// service principal (acm.amazonaws.com).
//
// These permissions allow ACM to issue and renew ACM certificates that reside in
// the same Amazon Web Services account as the CA. If you revoke these permissions,
// ACM will no longer renew the affected certificates automatically.
//
// Permissions can be granted with the [CreatePermission] action and listed with the [ListPermissions] action.
//
// About Permissions
//
//   - If the private CA and the certificates it issues reside in the same
//     account, you can use CreatePermission to grant permissions for ACM to carry
//     out automatic certificate renewals.
//
//   - For automatic certificate renewal to succeed, the ACM service principal
//     needs permissions to create, retrieve, and list certificates.
//
//   - If the private CA and the ACM certificates reside in different accounts,
//     then permissions cannot be used to enable automatic renewals. Instead, the ACM
//     certificate owner must set up a resource-based policy to enable cross-account
//     issuance and renewals. For more information, see [Using a Resource Based Policy with Amazon Web Services Private CA].
//
// [Using a Resource Based Policy with Amazon Web Services Private CA]: https://docs.aws.amazon.com/privateca/latest/userguide/pca-rbp.html
// [CreatePermission]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreatePermission.html
// [ListPermissions]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListPermissions.html
func (c *Client) DeletePermission(ctx context.Context, params *DeletePermissionInput, optFns ...func(*Options)) (*DeletePermissionOutput, error) {
	if params == nil {
		params = &DeletePermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePermission", params, optFns, c.addOperationDeletePermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePermissionInput struct {

	// The Amazon Resource Number (ARN) of the private CA that issued the permissions.
	// You can find the CA's ARN by calling the [ListCertificateAuthorities]action. This must have the following
	// form:
	//
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	// .
	//
	// [ListCertificateAuthorities]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListCertificateAuthorities.html
	//
	// This member is required.
	CertificateAuthorityArn *string

	// The Amazon Web Services service or identity that will have its CA permissions
	// revoked. At this time, the only valid service principal is acm.amazonaws.com
	//
	// This member is required.
	Principal *string

	// The Amazon Web Services account that calls this action.
	SourceAccount *string

	noSmithyDocumentSerde
}

type DeletePermissionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeletePermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePermission{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeletePermission"); err != nil {
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
	if err = addOpDeletePermissionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePermission(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeletePermission(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeletePermission",
	}
}
