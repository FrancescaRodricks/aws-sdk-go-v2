// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Installs an application to the device in a remote access session. For Android
// applications, the file must be in .apk format. For iOS applications, the file
// must be in .ipa format.
func (c *Client) InstallToRemoteAccessSession(ctx context.Context, params *InstallToRemoteAccessSessionInput, optFns ...func(*Options)) (*InstallToRemoteAccessSessionOutput, error) {
	if params == nil {
		params = &InstallToRemoteAccessSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InstallToRemoteAccessSession", params, optFns, c.addOperationInstallToRemoteAccessSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InstallToRemoteAccessSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to install an Android application (in .apk format) or an
// iOS application (in .ipa format) as part of a remote access session.
type InstallToRemoteAccessSessionInput struct {

	// The ARN of the app about which you are requesting information.
	//
	// This member is required.
	AppArn *string

	// The Amazon Resource Name (ARN) of the remote access session about which you are
	// requesting information.
	//
	// This member is required.
	RemoteAccessSessionArn *string

	noSmithyDocumentSerde
}

// Represents the response from the server after AWS Device Farm makes a request
// to install to a remote access session.
type InstallToRemoteAccessSessionOutput struct {

	// An app to upload or that has been uploaded.
	AppUpload *types.Upload

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationInstallToRemoteAccessSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpInstallToRemoteAccessSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpInstallToRemoteAccessSession{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "InstallToRemoteAccessSession"); err != nil {
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
	if err = addOpInstallToRemoteAccessSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInstallToRemoteAccessSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInstallToRemoteAccessSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InstallToRemoteAccessSession",
	}
}
