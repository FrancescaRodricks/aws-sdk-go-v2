// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an HSM client certificate that an Amazon Redshift cluster will use to
// connect to the client's HSM in order to store and retrieve the keys used to
// encrypt the cluster databases.
//
// The command returns a public key, which you must store in the HSM. In addition
// to creating the HSM certificate, you must create an Amazon Redshift HSM
// configuration that provides a cluster the information needed to store and use
// encryption keys in the HSM. For more information, go to [Hardware Security Modules]in the Amazon Redshift
// Cluster Management Guide.
//
// [Hardware Security Modules]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-db-encryption.html#working-with-HSM
func (c *Client) CreateHsmClientCertificate(ctx context.Context, params *CreateHsmClientCertificateInput, optFns ...func(*Options)) (*CreateHsmClientCertificateOutput, error) {
	if params == nil {
		params = &CreateHsmClientCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHsmClientCertificate", params, optFns, c.addOperationCreateHsmClientCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHsmClientCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHsmClientCertificateInput struct {

	// The identifier to be assigned to the new HSM client certificate that the
	// cluster will use to connect to the HSM to use the database encryption keys.
	//
	// This member is required.
	HsmClientCertificateIdentifier *string

	// A list of tag instances.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateHsmClientCertificateOutput struct {

	// Returns information about an HSM client certificate. The certificate is stored
	// in a secure Hardware Storage Module (HSM), and used by the Amazon Redshift
	// cluster to encrypt data files.
	HsmClientCertificate *types.HsmClientCertificate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateHsmClientCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateHsmClientCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateHsmClientCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateHsmClientCertificate"); err != nil {
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
	if err = addOpCreateHsmClientCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHsmClientCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateHsmClientCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateHsmClientCertificate",
	}
}
