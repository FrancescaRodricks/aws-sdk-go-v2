// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Imports the signing and encryption certificates that you need to create local
// (AS2) profiles and partner profiles.
func (c *Client) ImportCertificate(ctx context.Context, params *ImportCertificateInput, optFns ...func(*Options)) (*ImportCertificateOutput, error) {
	if params == nil {
		params = &ImportCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportCertificate", params, optFns, c.addOperationImportCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportCertificateInput struct {

	//   - For the CLI, provide a file path for a certificate in URI format. For
	//   example, --certificate file://encryption-cert.pem . Alternatively, you can
	//   provide the raw content.
	//
	//   - For the SDK, specify the raw content of a certificate file. For example,
	//   --certificate "`cat encryption-cert.pem`" .
	//
	// This member is required.
	Certificate *string

	// Specifies whether this certificate is used for signing or encryption.
	//
	// This member is required.
	Usage types.CertificateUsageType

	// An optional date that specifies when the certificate becomes active.
	ActiveDate *time.Time

	// An optional list of certificates that make up the chain for the certificate
	// that's being imported.
	CertificateChain *string

	// A short description that helps identify the certificate.
	Description *string

	// An optional date that specifies when the certificate becomes inactive.
	InactiveDate *time.Time

	//   - For the CLI, provide a file path for a private key in URI format.For
	//   example, --private-key file://encryption-key.pem . Alternatively, you can
	//   provide the raw content of the private key file.
	//
	//   - For the SDK, specify the raw content of a private key file. For example,
	//   --private-key "`cat encryption-key.pem`"
	PrivateKey *string

	// Key-value pairs that can be used to group and search for certificates.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type ImportCertificateOutput struct {

	// An array of identifiers for the imported certificates. You use this identifier
	// for working with profiles and partner profiles.
	//
	// This member is required.
	CertificateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpImportCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ImportCertificate"); err != nil {
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
	if err = addOpImportCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ImportCertificate",
	}
}
