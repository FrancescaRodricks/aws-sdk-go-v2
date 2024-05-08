// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an SSL/TLS certificate for an Amazon Lightsail content delivery network
// (CDN) distribution and a container service.
//
// After the certificate is valid, use the AttachCertificateToDistribution action
// to use the certificate and its domains with your distribution. Or use the
// UpdateContainerService action to use the certificate and its domains with your
// container service.
//
// Only certificates created in the us-east-1 Amazon Web Services Region can be
// attached to Lightsail distributions. Lightsail distributions are global
// resources that can reference an origin in any Amazon Web Services Region, and
// distribute its content globally. However, all distributions are located in the
// us-east-1 Region.
func (c *Client) CreateCertificate(ctx context.Context, params *CreateCertificateInput, optFns ...func(*Options)) (*CreateCertificateOutput, error) {
	if params == nil {
		params = &CreateCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCertificate", params, optFns, c.addOperationCreateCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCertificateInput struct {

	// The name for the certificate.
	//
	// This member is required.
	CertificateName *string

	// The domain name ( example.com ) for the certificate.
	//
	// This member is required.
	DomainName *string

	// An array of strings that specify the alternate domains ( example2.com ) and
	// subdomains ( blog.example.com ) for the certificate.
	//
	// You can specify a maximum of nine alternate domains (in addition to the primary
	// domain name).
	//
	// Wildcard domain entries ( *.example.com ) are not supported.
	SubjectAlternativeNames []string

	// The tag keys and optional values to add to the certificate during create.
	//
	// Use the TagResource action to tag a resource after it's created.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateCertificateOutput struct {

	// An object that describes the certificate created.
	Certificate *types.CertificateSummary

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCertificate"); err != nil {
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
	if err = addOpCreateCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCertificate",
	}
}
