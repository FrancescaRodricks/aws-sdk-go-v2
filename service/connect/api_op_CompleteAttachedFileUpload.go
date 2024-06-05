// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows you to confirm that the attached file has been uploaded using the
// pre-signed URL provided in the StartAttachedFileUpload API.
func (c *Client) CompleteAttachedFileUpload(ctx context.Context, params *CompleteAttachedFileUploadInput, optFns ...func(*Options)) (*CompleteAttachedFileUploadOutput, error) {
	if params == nil {
		params = &CompleteAttachedFileUploadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CompleteAttachedFileUpload", params, optFns, c.addOperationCompleteAttachedFileUploadMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CompleteAttachedFileUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to CompleteAttachedFileUpload API
type CompleteAttachedFileUploadInput struct {

	// The resource to which the attached file is (being) uploaded to. [Cases] are the only
	// current supported resource.
	//
	// This value must be a valid ARN.
	//
	// [Cases]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_CreateCase.html
	//
	// This member is required.
	AssociatedResourceArn *string

	// The unique identifier of the attached file resource.
	//
	// This member is required.
	FileId *string

	// The unique identifier of the Connect instance.
	//
	// This member is required.
	InstanceId *string

	noSmithyDocumentSerde
}

// Response from CompleteAttachedFileUpload API
type CompleteAttachedFileUploadOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCompleteAttachedFileUploadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCompleteAttachedFileUpload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCompleteAttachedFileUpload{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CompleteAttachedFileUpload"); err != nil {
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
	if err = addOpCompleteAttachedFileUploadValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCompleteAttachedFileUpload(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCompleteAttachedFileUpload(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CompleteAttachedFileUpload",
	}
}
