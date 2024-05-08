// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the name and/or the path of the specified server certificate stored in
// IAM.
//
// For more information about working with server certificates, see [Working with server certificates] in the IAM
// User Guide. This topic also includes a list of Amazon Web Services services that
// can use the server certificates that you manage with IAM.
//
// You should understand the implications of changing a server certificate's path
// or name. For more information, see [Renaming a server certificate]in the IAM User Guide.
//
// The person making the request (the principal), must have permission to change
// the server certificate with the old name and the new name. For example, to
// change the certificate named ProductionCert to ProdCert , the principal must
// have a policy that allows them to update both certificates. If the principal has
// permission to update the ProductionCert group, but not the ProdCert
// certificate, then the update fails. For more information about permissions, see [Access management]
// in the IAM User Guide.
//
// [Renaming a server certificate]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs_manage.html#RenamingServerCerts
// [Access management]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access.html
// [Working with server certificates]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html
func (c *Client) UpdateServerCertificate(ctx context.Context, params *UpdateServerCertificateInput, optFns ...func(*Options)) (*UpdateServerCertificateOutput, error) {
	if params == nil {
		params = &UpdateServerCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateServerCertificate", params, optFns, c.addOperationUpdateServerCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateServerCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServerCertificateInput struct {

	// The name of the server certificate that you want to update.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	ServerCertificateName *string

	// The new path for the server certificate. Include this only if you are updating
	// the server certificate's path.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of
	// either a forward slash (/) by itself or a string that must begin and end with
	// forward slashes. In addition, it can contain any ASCII character from the ! (
	// \u0021 ) through the DEL character ( \u007F ), including most punctuation
	// characters, digits, and upper and lowercased letters.
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	NewPath *string

	// The new name for the server certificate. Include this only if you are updating
	// the server certificate's name. The name of the certificate cannot contain any
	// spaces.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	NewServerCertificateName *string

	noSmithyDocumentSerde
}

type UpdateServerCertificateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateServerCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateServerCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateServerCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateServerCertificate"); err != nil {
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
	if err = addOpUpdateServerCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateServerCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateServerCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateServerCertificate",
	}
}
