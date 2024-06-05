// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This example uses a @mediaType trait on the payload to force a custom
// content-type to be serialized.
func (c *Client) HttpPayloadTraitsWithMediaType(ctx context.Context, params *HttpPayloadTraitsWithMediaTypeInput, optFns ...func(*Options)) (*HttpPayloadTraitsWithMediaTypeOutput, error) {
	if params == nil {
		params = &HttpPayloadTraitsWithMediaTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "HttpPayloadTraitsWithMediaType", params, optFns, c.addOperationHttpPayloadTraitsWithMediaTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*HttpPayloadTraitsWithMediaTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpPayloadTraitsWithMediaTypeInput struct {

	// This value conforms to the media type: text/plain
	Blob []byte

	Foo *string

	noSmithyDocumentSerde
}

type HttpPayloadTraitsWithMediaTypeOutput struct {

	// This value conforms to the media type: text/plain
	Blob []byte

	Foo *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationHttpPayloadTraitsWithMediaTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpHttpPayloadTraitsWithMediaType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpHttpPayloadTraitsWithMediaType{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "HttpPayloadTraitsWithMediaType"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opHttpPayloadTraitsWithMediaType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opHttpPayloadTraitsWithMediaType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "HttpPayloadTraitsWithMediaType",
	}
}
