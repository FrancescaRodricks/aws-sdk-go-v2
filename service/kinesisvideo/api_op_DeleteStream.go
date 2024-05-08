// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a Kinesis video stream and the data contained in the stream.
//
// This method marks the stream for deletion, and makes the data in the stream
// inaccessible immediately.
//
// To ensure that you have the latest version of the stream before deleting it,
// you can specify the stream version. Kinesis Video Streams assigns a version to
// each stream. When you update a stream, Kinesis Video Streams assigns a new
// version number. To get the latest stream version, use the DescribeStream API.
//
// This operation requires permission for the KinesisVideo:DeleteStream action.
func (c *Client) DeleteStream(ctx context.Context, params *DeleteStreamInput, optFns ...func(*Options)) (*DeleteStreamOutput, error) {
	if params == nil {
		params = &DeleteStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteStream", params, optFns, c.addOperationDeleteStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteStreamInput struct {

	// The Amazon Resource Name (ARN) of the stream that you want to delete.
	//
	// This member is required.
	StreamARN *string

	// Optional: The version of the stream that you want to delete.
	//
	// Specify the version as a safeguard to ensure that your are deleting the correct
	// stream. To get the stream version, use the DescribeStream API.
	//
	// If not specified, only the CreationTime is checked before deleting the stream.
	CurrentVersion *string

	noSmithyDocumentSerde
}

type DeleteStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteStream{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteStream"); err != nil {
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
	if err = addOpDeleteStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteStream",
	}
}
