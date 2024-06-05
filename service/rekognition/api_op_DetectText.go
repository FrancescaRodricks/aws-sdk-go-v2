// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Detects text in the input image and converts it into machine-readable text.
//
// Pass the input image as base64-encoded image bytes or as a reference to an
// image in an Amazon S3 bucket. If you use the AWS CLI to call Amazon Rekognition
// operations, you must pass it as a reference to an image in an Amazon S3 bucket.
// For the AWS CLI, passing image bytes is not supported. The image must be either
// a .png or .jpeg formatted file.
//
// The DetectText operation returns text in an array of TextDetection elements, TextDetections .
// Each TextDetection element provides information about a single word or line of
// text that was detected in the image.
//
// A word is one or more script characters that are not separated by spaces.
// DetectText can detect up to 100 words in an image.
//
// A line is a string of equally spaced words. A line isn't necessarily a complete
// sentence. For example, a driver's license number is detected as a line. A line
// ends when there is no aligned text after it. Also, a line ends when there is a
// large gap between words, relative to the length of the words. This means,
// depending on the gap between words, Amazon Rekognition may detect multiple lines
// in text aligned in the same direction. Periods don't represent the end of a
// line. If a sentence spans multiple lines, the DetectText operation returns
// multiple lines.
//
// To determine whether a TextDetection element is a line of text or a word, use
// the TextDetection object Type field.
//
// To be detected, text must be within +/- 90 degrees orientation of the
// horizontal axis.
//
// For more information, see Detecting text in the Amazon Rekognition Developer
// Guide.
func (c *Client) DetectText(ctx context.Context, params *DetectTextInput, optFns ...func(*Options)) (*DetectTextOutput, error) {
	if params == nil {
		params = &DetectTextInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectText", params, optFns, c.addOperationDetectTextMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectTextOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectTextInput struct {

	// The input image as base64-encoded bytes or an Amazon S3 object. If you use the
	// AWS CLI to call Amazon Rekognition operations, you can't pass image bytes.
	//
	// If you are using an AWS SDK to call Amazon Rekognition, you might not need to
	// base64-encode image bytes passed using the Bytes field. For more information,
	// see Images in the Amazon Rekognition developer guide.
	//
	// This member is required.
	Image *types.Image

	// Optional parameters that let you set the criteria that the text must meet to be
	// included in your response.
	Filters *types.DetectTextFilters

	noSmithyDocumentSerde
}

type DetectTextOutput struct {

	// An array of text that was detected in the input image.
	TextDetections []types.TextDetection

	// The model version used to detect text.
	TextModelVersion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetectTextMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetectText{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectText{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DetectText"); err != nil {
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
	if err = addOpDetectTextValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetectText(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetectText(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DetectText",
	}
}
