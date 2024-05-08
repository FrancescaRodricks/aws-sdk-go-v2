// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Creates a decoder manifest using your existing CAN DBC file from your local
//
// device.
func (c *Client) ImportDecoderManifest(ctx context.Context, params *ImportDecoderManifestInput, optFns ...func(*Options)) (*ImportDecoderManifestOutput, error) {
	if params == nil {
		params = &ImportDecoderManifestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportDecoderManifest", params, optFns, c.addOperationImportDecoderManifestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportDecoderManifestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportDecoderManifestInput struct {

	//  The name of the decoder manifest to import.
	//
	// This member is required.
	Name *string

	//  The file to load into an Amazon Web Services account.
	//
	// This member is required.
	NetworkFileDefinitions []types.NetworkFileDefinition

	noSmithyDocumentSerde
}

type ImportDecoderManifestOutput struct {

	//  The Amazon Resource Name (ARN) of the decoder manifest that was imported.
	//
	// This member is required.
	Arn *string

	//  The name of the imported decoder manifest.
	//
	// This member is required.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportDecoderManifestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpImportDecoderManifest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpImportDecoderManifest{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ImportDecoderManifest"); err != nil {
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
	if err = addOpImportDecoderManifestValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportDecoderManifest(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportDecoderManifest(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ImportDecoderManifest",
	}
}
