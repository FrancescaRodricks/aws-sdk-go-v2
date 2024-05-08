// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a platform application object for one of the supported push
// notification services, such as APNS and GCM (Firebase Cloud Messaging), to which
// devices and mobile apps may register. You must specify PlatformPrincipal and
// PlatformCredential attributes when using the CreatePlatformApplication action.
//
// PlatformPrincipal and PlatformCredential are received from the notification
// service.
//
//   - For ADM , PlatformPrincipal is client id and PlatformCredential is client
//     secret .
//
//   - For Baidu , PlatformPrincipal is API key and PlatformCredential is secret
//     key .
//
//   - For APNS and APNS_SANDBOX using certificate credentials, PlatformPrincipal
//     is SSL certificate and PlatformCredential is private key .
//
//   - For APNS and APNS_SANDBOX using token credentials, PlatformPrincipal is
//     signing key ID and PlatformCredential is signing key .
//
//   - For GCM (Firebase Cloud Messaging) using key credentials, there is no
//     PlatformPrincipal . The PlatformCredential is API key .
//
//   - For GCM (Firebase Cloud Messaging) using token credentials, there is no
//     PlatformPrincipal . The PlatformCredential is a JSON formatted private key
//     file. When using the Amazon Web Services CLI, the file must be in string format
//     and special characters must be ignored. To format the file correctly, Amazon SNS
//     recommends using the following command: SERVICE_JSON=`jq @json <<< cat
//     service.json` .
//
//   - For MPNS , PlatformPrincipal is TLS certificate and PlatformCredential is
//     private key .
//
//   - For WNS , PlatformPrincipal is Package Security Identifier and
//     PlatformCredential is secret key .
//
// You can use the returned PlatformApplicationArn as an attribute for the
// CreatePlatformEndpoint action.
func (c *Client) CreatePlatformApplication(ctx context.Context, params *CreatePlatformApplicationInput, optFns ...func(*Options)) (*CreatePlatformApplicationOutput, error) {
	if params == nil {
		params = &CreatePlatformApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePlatformApplication", params, optFns, c.addOperationCreatePlatformApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePlatformApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for CreatePlatformApplication action.
type CreatePlatformApplicationInput struct {

	// For a list of attributes, see [SetPlatformApplicationAttributes]SetPlatformApplicationAttributes .
	//
	// [SetPlatformApplicationAttributes]: https://docs.aws.amazon.com/sns/latest/api/API_SetPlatformApplicationAttributes.html
	//
	// This member is required.
	Attributes map[string]string

	// Application names must be made up of only uppercase and lowercase ASCII
	// letters, numbers, underscores, hyphens, and periods, and must be between 1 and
	// 256 characters long.
	//
	// This member is required.
	Name *string

	// The following platforms are supported: ADM (Amazon Device Messaging), APNS
	// (Apple Push Notification Service), APNS_SANDBOX, and GCM (Firebase Cloud
	// Messaging).
	//
	// This member is required.
	Platform *string

	noSmithyDocumentSerde
}

// Response from CreatePlatformApplication action.
type CreatePlatformApplicationOutput struct {

	// PlatformApplicationArn is returned.
	PlatformApplicationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePlatformApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreatePlatformApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreatePlatformApplication{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePlatformApplication"); err != nil {
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
	if err = addOpCreatePlatformApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePlatformApplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePlatformApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePlatformApplication",
	}
}
