// Code generated by smithy-go-codegen DO NOT EDIT.

package mediastoredata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"time"
)

// Downloads the object at the specified path. If the object’s upload availability
// is set to streaming , AWS Elemental MediaStore downloads the object even if it’s
// still uploading the object.
func (c *Client) GetObject(ctx context.Context, params *GetObjectInput, optFns ...func(*Options)) (*GetObjectOutput, error) {
	if params == nil {
		params = &GetObjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetObject", params, optFns, c.addOperationGetObjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetObjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetObjectInput struct {

	// The path (including the file name) where the object is stored in the container.
	// Format: //
	//
	// For example, to upload the file mlaw.avi to the folder path premium\canada in
	// the container movies , enter the path premium/canada/mlaw.avi .
	//
	// Do not include the container name in this path.
	//
	// If the path includes any folders that don't exist yet, the service creates
	// them. For example, suppose you have an existing premium/usa subfolder. If you
	// specify premium/canada , the service creates a canada subfolder in the premium
	// folder. You then have two subfolders, usa and canada , in the premium folder.
	//
	// There is no correlation between the path to the source and the path (folders)
	// in the container in AWS Elemental MediaStore.
	//
	// For more information about folders and how they exist in a container, see the [AWS Elemental MediaStore User Guide].
	//
	// The file name is the name that is assigned to the file that you upload. The
	// file can have the same name inside and outside of AWS Elemental MediaStore, or
	// it can have the same name. The file name can include or omit an extension.
	//
	// [AWS Elemental MediaStore User Guide]: http://docs.aws.amazon.com/mediastore/latest/ug/
	//
	// This member is required.
	Path *string

	// The range bytes of an object to retrieve. For more information about the Range
	// header, see [http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35]. AWS Elemental MediaStore ignores this header for partially
	// uploaded objects that have streaming upload availability.
	//
	// [http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35]: http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35
	Range *string

	noSmithyDocumentSerde
}

type GetObjectOutput struct {

	// The HTML status code of the request. Status codes ranging from 200 to 299
	// indicate success. All other status codes indicate the type of error that
	// occurred.
	//
	// This member is required.
	StatusCode int32

	// The bytes of the object.
	Body io.ReadCloser

	// An optional CacheControl header that allows the caller to control the object's
	// cache behavior. Headers can be passed in as specified in the HTTP spec at [https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9].
	//
	// Headers with a custom user-defined value are also accepted.
	//
	// [https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9]: https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9
	CacheControl *string

	// The length of the object in bytes.
	ContentLength *int64

	// The range of bytes to retrieve.
	ContentRange *string

	// The content type of the object.
	ContentType *string

	// The ETag that represents a unique instance of the object.
	ETag *string

	// The date and time that the object was last modified.
	LastModified *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetObjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetObject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetObject{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetObject"); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addTimeOffsetDeserializer(stack, c); err != nil {
		return err
	}
	if err = addOpGetObjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetObject(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetObject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetObject",
	}
}
