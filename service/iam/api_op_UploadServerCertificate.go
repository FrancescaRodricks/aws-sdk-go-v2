// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Uploads a server certificate entity for the Amazon Web Services account. The
// server certificate entity includes a public key certificate, a private key, and
// an optional certificate chain, which should all be PEM-encoded.
//
// We recommend that you use [Certificate Manager] to provision, manage, and deploy your server
// certificates. With ACM you can request a certificate, deploy it to Amazon Web
// Services resources, and let ACM handle certificate renewals for you.
// Certificates provided by ACM are free. For more information about using ACM, see
// the [Certificate Manager User Guide].
//
// For more information about working with server certificates, see [Working with server certificates] in the IAM
// User Guide. This topic includes a list of Amazon Web Services services that can
// use the server certificates that you manage with IAM.
//
// For information about the number of server certificates you can upload, see [IAM and STS quotas] in
// the IAM User Guide.
//
// Because the body of the public key certificate, private key, and the
// certificate chain can be large, you should use POST rather than GET when calling
// UploadServerCertificate . For information about setting up signatures and
// authorization through the API, see [Signing Amazon Web Services API requests]in the Amazon Web Services General
// Reference. For general information about using the Query API with IAM, see [Calling the API by making HTTP query requests]in
// the IAM User Guide.
//
// [Certificate Manager]: https://docs.aws.amazon.com/acm/
// [Certificate Manager User Guide]: https://docs.aws.amazon.com/acm/latest/userguide/
// [IAM and STS quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
// [Working with server certificates]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html
// [Signing Amazon Web Services API requests]: https://docs.aws.amazon.com/general/latest/gr/signing_aws_api_requests.html
// [Calling the API by making HTTP query requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/programming.html
func (c *Client) UploadServerCertificate(ctx context.Context, params *UploadServerCertificateInput, optFns ...func(*Options)) (*UploadServerCertificateOutput, error) {
	if params == nil {
		params = &UploadServerCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UploadServerCertificate", params, optFns, c.addOperationUploadServerCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UploadServerCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UploadServerCertificateInput struct {

	// The contents of the public key certificate in PEM-encoded format.
	//
	// The [regex pattern] used to validate this parameter is a string of characters consisting of
	// the following:
	//
	//   - Any printable ASCII character ranging from the space character ( \u0020 )
	//   through the end of the ASCII character range
	//
	//   - The printable characters in the Basic Latin and Latin-1 Supplement
	//   character set (through \u00FF )
	//
	//   - The special characters tab ( \u0009 ), line feed ( \u000A ), and carriage
	//   return ( \u000D )
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	CertificateBody *string

	// The contents of the private key in PEM-encoded format.
	//
	// The [regex pattern] used to validate this parameter is a string of characters consisting of
	// the following:
	//
	//   - Any printable ASCII character ranging from the space character ( \u0020 )
	//   through the end of the ASCII character range
	//
	//   - The printable characters in the Basic Latin and Latin-1 Supplement
	//   character set (through \u00FF )
	//
	//   - The special characters tab ( \u0009 ), line feed ( \u000A ), and carriage
	//   return ( \u000D )
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	PrivateKey *string

	// The name for the server certificate. Do not include the path in this value. The
	// name of the certificate cannot contain any spaces.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	ServerCertificateName *string

	// The contents of the certificate chain. This is typically a concatenation of the
	// PEM-encoded public key certificates of the chain.
	//
	// The [regex pattern] used to validate this parameter is a string of characters consisting of
	// the following:
	//
	//   - Any printable ASCII character ranging from the space character ( \u0020 )
	//   through the end of the ASCII character range
	//
	//   - The printable characters in the Basic Latin and Latin-1 Supplement
	//   character set (through \u00FF )
	//
	//   - The special characters tab ( \u0009 ), line feed ( \u000A ), and carriage
	//   return ( \u000D )
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	CertificateChain *string

	// The path for the server certificate. For more information about paths, see [IAM identifiers] in
	// the IAM User Guide.
	//
	// This parameter is optional. If it is not included, it defaults to a slash (/).
	// This parameter allows (through its [regex pattern]) a string of characters consisting of
	// either a forward slash (/) by itself or a string that must begin and end with
	// forward slashes. In addition, it can contain any ASCII character from the ! (
	// \u0021 ) through the DEL character ( \u007F ), including most punctuation
	// characters, digits, and upper and lowercased letters.
	//
	// If you are uploading a server certificate specifically for use with Amazon
	// CloudFront distributions, you must specify a path using the path parameter. The
	// path must begin with /cloudfront and must include a trailing slash (for
	// example, /cloudfront/test/ ).
	//
	// [IAM identifiers]: https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html
	// [regex pattern]: http://wikipedia.org/wiki/regex
	Path *string

	// A list of tags that you want to attach to the new IAM server certificate
	// resource. Each tag consists of a key name and an associated value. For more
	// information about tagging, see [Tagging IAM resources]in the IAM User Guide.
	//
	// If any one of the tags is invalid or if you exceed the allowed maximum number
	// of tags, then the entire request fails and the resource is not created.
	//
	// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Contains the response to a successful UploadServerCertificate request.
type UploadServerCertificateOutput struct {

	// The meta information of the uploaded server certificate without its certificate
	// body, certificate chain, and private key.
	ServerCertificateMetadata *types.ServerCertificateMetadata

	// A list of tags that are attached to the new IAM server certificate. The
	// returned list of tags is sorted by tag key. For more information about tagging,
	// see [Tagging IAM resources]in the IAM User Guide.
	//
	// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUploadServerCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUploadServerCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUploadServerCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UploadServerCertificate"); err != nil {
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
	if err = addOpUploadServerCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUploadServerCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUploadServerCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UploadServerCertificate",
	}
}
