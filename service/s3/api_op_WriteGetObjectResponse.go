// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"strings"
	"time"
)

// This operation is not supported by directory buckets.
//
// Passes transformed objects to a GetObject operation when using Object Lambda
// access points. For information about Object Lambda access points, see [Transforming objects with Object Lambda access points]in the
// Amazon S3 User Guide.
//
// This operation supports metadata that can be returned by [GetObject], in addition to
// RequestRoute , RequestToken , StatusCode , ErrorCode , and ErrorMessage . The
// GetObject response metadata is supported so that the WriteGetObjectResponse
// caller, typically an Lambda function, can provide the same metadata when it
// internally invokes GetObject . When WriteGetObjectResponse is called by a
// customer-owned Lambda function, the metadata returned to the end user GetObject
// call might differ from what Amazon S3 would normally return.
//
// You can include any number of metadata headers. When including a metadata
// header, it should be prefaced with x-amz-meta . For example,
// x-amz-meta-my-custom-header: MyCustomValue . The primary use case for this is to
// forward GetObject metadata.
//
// Amazon Web Services provides some prebuilt Lambda functions that you can use
// with S3 Object Lambda to detect and redact personally identifiable information
// (PII) and decompress S3 objects. These Lambda functions are available in the
// Amazon Web Services Serverless Application Repository, and can be selected
// through the Amazon Web Services Management Console when you create your Object
// Lambda access point.
//
// Example 1: PII Access Control - This Lambda function uses Amazon Comprehend, a
// natural language processing (NLP) service using machine learning to find
// insights and relationships in text. It automatically detects personally
// identifiable information (PII) such as names, addresses, dates, credit card
// numbers, and social security numbers from documents in your Amazon S3 bucket.
//
// Example 2: PII Redaction - This Lambda function uses Amazon Comprehend, a
// natural language processing (NLP) service using machine learning to find
// insights and relationships in text. It automatically redacts personally
// identifiable information (PII) such as names, addresses, dates, credit card
// numbers, and social security numbers from documents in your Amazon S3 bucket.
//
// Example 3: Decompression - The Lambda function S3ObjectLambdaDecompression, is
// equipped to decompress objects stored in S3 in one of six compressed file
// formats including bzip2, gzip, snappy, zlib, zstandard and ZIP.
//
// For information on how to view and use these functions, see [Using Amazon Web Services built Lambda functions] in the Amazon S3
// User Guide.
//
// [Transforming objects with Object Lambda access points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/transforming-objects.html
// [Using Amazon Web Services built Lambda functions]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/olap-examples.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
func (c *Client) WriteGetObjectResponse(ctx context.Context, params *WriteGetObjectResponseInput, optFns ...func(*Options)) (*WriteGetObjectResponseOutput, error) {
	if params == nil {
		params = &WriteGetObjectResponseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "WriteGetObjectResponse", params, optFns, c.addOperationWriteGetObjectResponseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*WriteGetObjectResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type WriteGetObjectResponseInput struct {

	// Route prefix to the HTTP URL generated.
	//
	// This member is required.
	RequestRoute *string

	// A single use encrypted token that maps WriteGetObjectResponse to the end user
	// GetObject request.
	//
	// This member is required.
	RequestToken *string

	// Indicates that a range of bytes was specified.
	AcceptRanges *string

	// The object data.
	Body io.Reader

	//  Indicates whether the object stored in Amazon S3 uses an S3 bucket key for
	// server-side encryption with Amazon Web Services KMS (SSE-KMS).
	BucketKeyEnabled *bool

	// Specifies caching behavior along the request/reply chain.
	CacheControl *string

	// This header can be used as a data integrity check to verify that the data
	// received is the same data that was originally sent. This specifies the
	// base64-encoded, 32-bit CRC32 checksum of the object returned by the Object
	// Lambda function. This may not match the checksum for the object stored in Amazon
	// S3. Amazon S3 will perform validation of the checksum values only when the
	// original GetObject request required checksum validation. For more information
	// about checksums, see [Checking object integrity]in the Amazon S3 User Guide.
	//
	// Only one checksum header can be specified at a time. If you supply multiple
	// checksum headers, this request will fail.
	//
	// [Checking object integrity]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html
	ChecksumCRC32 *string

	// This header can be used as a data integrity check to verify that the data
	// received is the same data that was originally sent. This specifies the
	// base64-encoded, 32-bit CRC32C checksum of the object returned by the Object
	// Lambda function. This may not match the checksum for the object stored in Amazon
	// S3. Amazon S3 will perform validation of the checksum values only when the
	// original GetObject request required checksum validation. For more information
	// about checksums, see [Checking object integrity]in the Amazon S3 User Guide.
	//
	// Only one checksum header can be specified at a time. If you supply multiple
	// checksum headers, this request will fail.
	//
	// [Checking object integrity]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html
	ChecksumCRC32C *string

	// This header can be used as a data integrity check to verify that the data
	// received is the same data that was originally sent. This specifies the
	// base64-encoded, 160-bit SHA-1 digest of the object returned by the Object Lambda
	// function. This may not match the checksum for the object stored in Amazon S3.
	// Amazon S3 will perform validation of the checksum values only when the original
	// GetObject request required checksum validation. For more information about
	// checksums, see [Checking object integrity]in the Amazon S3 User Guide.
	//
	// Only one checksum header can be specified at a time. If you supply multiple
	// checksum headers, this request will fail.
	//
	// [Checking object integrity]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html
	ChecksumSHA1 *string

	// This header can be used as a data integrity check to verify that the data
	// received is the same data that was originally sent. This specifies the
	// base64-encoded, 256-bit SHA-256 digest of the object returned by the Object
	// Lambda function. This may not match the checksum for the object stored in Amazon
	// S3. Amazon S3 will perform validation of the checksum values only when the
	// original GetObject request required checksum validation. For more information
	// about checksums, see [Checking object integrity]in the Amazon S3 User Guide.
	//
	// Only one checksum header can be specified at a time. If you supply multiple
	// checksum headers, this request will fail.
	//
	// [Checking object integrity]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html
	ChecksumSHA256 *string

	// Specifies presentational information for the object.
	ContentDisposition *string

	// Specifies what content encodings have been applied to the object and thus what
	// decoding mechanisms must be applied to obtain the media-type referenced by the
	// Content-Type header field.
	ContentEncoding *string

	// The language the content is in.
	ContentLanguage *string

	// The size of the content body in bytes.
	ContentLength *int64

	// The portion of the object returned in the response.
	ContentRange *string

	// A standard MIME type describing the format of the object data.
	ContentType *string

	// Specifies whether an object stored in Amazon S3 is ( true ) or is not ( false )
	// a delete marker.
	DeleteMarker *bool

	// An opaque identifier assigned by a web server to a specific version of a
	// resource found at a URL.
	ETag *string

	// A string that uniquely identifies an error condition. Returned in the  tag of
	// the error XML response for a corresponding GetObject call. Cannot be used with
	// a successful StatusCode header or when the transformed object is provided in
	// the body. All error codes from S3 are sentence-cased. The regular expression
	// (regex) value is "^[A-Z][a-zA-Z]+$" .
	ErrorCode *string

	// Contains a generic description of the error condition. Returned in the tag of
	// the error XML response for a corresponding GetObject call. Cannot be used with
	// a successful StatusCode header or when the transformed object is provided in
	// body.
	ErrorMessage *string

	// If the object expiration is configured (see PUT Bucket lifecycle), the response
	// includes this header. It includes the expiry-date and rule-id key-value pairs
	// that provide the object expiration information. The value of the rule-id is
	// URL-encoded.
	Expiration *string

	// The date and time at which the object is no longer cacheable.
	Expires *time.Time

	// The date and time that the object was last modified.
	LastModified *time.Time

	// A map of metadata to store with the object in S3.
	Metadata map[string]string

	// Set to the number of metadata entries not returned in x-amz-meta headers. This
	// can happen if you create metadata using an API like SOAP that supports more
	// flexible metadata than the REST API. For example, using SOAP, you can create
	// metadata whose values are not legal HTTP headers.
	MissingMeta *int32

	// Indicates whether an object stored in Amazon S3 has an active legal hold.
	ObjectLockLegalHoldStatus types.ObjectLockLegalHoldStatus

	// Indicates whether an object stored in Amazon S3 has Object Lock enabled. For
	// more information about S3 Object Lock, see [Object Lock].
	//
	// [Object Lock]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock.html
	ObjectLockMode types.ObjectLockMode

	// The date and time when Object Lock is configured to expire.
	ObjectLockRetainUntilDate *time.Time

	// The count of parts this object has.
	PartsCount *int32

	// Indicates if request involves bucket that is either a source or destination in
	// a Replication rule. For more information about S3 Replication, see [Replication].
	//
	// [Replication]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/replication.html
	ReplicationStatus types.ReplicationStatus

	// If present, indicates that the requester was successfully charged for the
	// request.
	//
	// This functionality is not supported for directory buckets.
	RequestCharged types.RequestCharged

	// Provides information about object restoration operation and expiration time of
	// the restored object copy.
	Restore *string

	// Encryption algorithm used if server-side encryption with a customer-provided
	// encryption key was specified for object stored in Amazon S3.
	SSECustomerAlgorithm *string

	//  128-bit MD5 digest of customer-provided encryption key used in Amazon S3 to
	// encrypt data stored in S3. For more information, see [Protecting data using server-side encryption with customer-provided encryption keys (SSE-C)].
	//
	// [Protecting data using server-side encryption with customer-provided encryption keys (SSE-C)]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/ServerSideEncryptionCustomerKeys.html
	SSECustomerKeyMD5 *string

	//  If present, specifies the ID (Key ID, Key ARN, or Key Alias) of the Amazon Web
	// Services Key Management Service (Amazon Web Services KMS) symmetric encryption
	// customer managed key that was used for stored in Amazon S3 object.
	SSEKMSKeyId *string

	//  The server-side encryption algorithm used when storing requested object in
	// Amazon S3 (for example, AES256, aws:kms ).
	ServerSideEncryption types.ServerSideEncryption

	// The integer status code for an HTTP response of a corresponding GetObject
	// request. The following is a list of status codes.
	//
	//   - 200 - OK
	//
	//   - 206 - Partial Content
	//
	//   - 304 - Not Modified
	//
	//   - 400 - Bad Request
	//
	//   - 401 - Unauthorized
	//
	//   - 403 - Forbidden
	//
	//   - 404 - Not Found
	//
	//   - 405 - Method Not Allowed
	//
	//   - 409 - Conflict
	//
	//   - 411 - Length Required
	//
	//   - 412 - Precondition Failed
	//
	//   - 416 - Range Not Satisfiable
	//
	//   - 500 - Internal Server Error
	//
	//   - 503 - Service Unavailable
	StatusCode *int32

	// Provides storage class information of the object. Amazon S3 returns this header
	// for all objects except for S3 Standard storage class objects.
	//
	// For more information, see [Storage Classes].
	//
	// [Storage Classes]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html
	StorageClass types.StorageClass

	// The number of tags, if any, on the object.
	TagCount *int32

	// An ID used to reference a specific version of the object.
	VersionId *string

	noSmithyDocumentSerde
}

func (in *WriteGetObjectResponseInput) bindEndpointParams(p *EndpointParameters) {

	p.UseObjectLambdaEndpoint = ptr.Bool(true)
}

type WriteGetObjectResponseOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationWriteGetObjectResponseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpWriteGetObjectResponse{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpWriteGetObjectResponse{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "WriteGetObjectResponse"); err != nil {
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
	if err = addUnsignedPayload(stack); err != nil {
		return err
	}
	if err = addContentSHA256Header(stack); err != nil {
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
	if err = addPutBucketContextMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addTimeOffsetDeserializer(stack, c); err != nil {
		return err
	}
	if err = addEndpointPrefix_opWriteGetObjectResponseMiddleware(stack); err != nil {
		return err
	}
	if err = addOpWriteGetObjectResponseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opWriteGetObjectResponse(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addWriteGetObjectResponseUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.UseDynamicPayloadSigningMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSerializeImmutableHostnameBucketMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opWriteGetObjectResponseMiddleware struct {
}

func (*endpointPrefix_opWriteGetObjectResponseMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opWriteGetObjectResponseMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	opaqueInput := getOperationInput(ctx)
	input, ok := opaqueInput.(*WriteGetObjectResponseInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", opaqueInput)
	}

	var prefix strings.Builder
	if input.RequestRoute == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("RequestRoute forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.RequestRoute) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("RequestRoute forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.RequestRoute)}
	} else {
		prefix.WriteString(*input.RequestRoute)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opWriteGetObjectResponseMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opWriteGetObjectResponseMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opWriteGetObjectResponse(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "WriteGetObjectResponse",
	}
}

func addWriteGetObjectResponseUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: nopGetBucketAccessor,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           true,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
