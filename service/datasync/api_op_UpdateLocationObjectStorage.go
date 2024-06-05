// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates some parameters of an existing DataSync location for an object storage
// system.
func (c *Client) UpdateLocationObjectStorage(ctx context.Context, params *UpdateLocationObjectStorageInput, optFns ...func(*Options)) (*UpdateLocationObjectStorageOutput, error) {
	if params == nil {
		params = &UpdateLocationObjectStorageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLocationObjectStorage", params, optFns, c.addOperationUpdateLocationObjectStorageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLocationObjectStorageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLocationObjectStorageInput struct {

	// Specifies the ARN of the object storage system location that you're updating.
	//
	// This member is required.
	LocationArn *string

	// Specifies the access key (for example, a user name) if credentials are required
	// to authenticate with the object storage server.
	AccessKey *string

	// Specifies the Amazon Resource Names (ARNs) of the DataSync agents that can
	// securely connect with your location.
	AgentArns []string

	// Specifies the secret key (for example, a password) if credentials are required
	// to authenticate with the object storage server.
	SecretKey *string

	// Specifies a certificate chain for DataSync to authenticate with your object
	// storage system if the system uses a private or self-signed certificate authority
	// (CA). You must specify a single .pem file with a full certificate chain (for
	// example, file:///home/user/.ssh/object_storage_certificates.pem ).
	//
	// The certificate chain might include:
	//
	//   - The object storage system's certificate
	//
	//   - All intermediate certificates (if there are any)
	//
	//   - The root certificate of the signing CA
	//
	// You can concatenate your certificates into a .pem file (which can be up to
	// 32768 bytes before base64 encoding). The following example cat command creates
	// an object_storage_certificates.pem file that includes three certificates:
	//
	//     cat object_server_certificate.pem intermediate_certificate.pem
	//     ca_root_certificate.pem > object_storage_certificates.pem
	//
	// To use this parameter, configure ServerProtocol to HTTPS .
	//
	// Updating this parameter doesn't interfere with tasks that you have in progress.
	ServerCertificate []byte

	// Specifies the port that your object storage server accepts inbound network
	// traffic on (for example, port 443).
	ServerPort *int32

	// Specifies the protocol that your object storage server uses to communicate.
	ServerProtocol types.ObjectStorageServerProtocol

	// Specifies the object prefix for your object storage server. If this is a source
	// location, DataSync only copies objects with this prefix. If this is a
	// destination location, DataSync writes all objects with this prefix.
	Subdirectory *string

	noSmithyDocumentSerde
}

type UpdateLocationObjectStorageOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLocationObjectStorageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLocationObjectStorage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLocationObjectStorage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateLocationObjectStorage"); err != nil {
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
	if err = addOpUpdateLocationObjectStorageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLocationObjectStorage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLocationObjectStorage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateLocationObjectStorage",
	}
}
