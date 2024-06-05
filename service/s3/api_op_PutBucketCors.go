// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalChecksum "github.com/aws/aws-sdk-go-v2/service/internal/checksum"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation is not supported by directory buckets.
//
// Sets the cors configuration for your bucket. If the configuration exists,
// Amazon S3 replaces it.
//
// To use this operation, you must be allowed to perform the s3:PutBucketCORS
// action. By default, the bucket owner has this permission and can grant it to
// others.
//
// You set this configuration on a bucket so that the bucket can service
// cross-origin requests. For example, you might want to enable a request whose
// origin is http://www.example.com to access your Amazon S3 bucket at
// my.example.bucket.com by using the browser's XMLHttpRequest capability.
//
// To enable cross-origin resource sharing (CORS) on a bucket, you add the cors
// subresource to the bucket. The cors subresource is an XML document in which you
// configure rules that identify origins and the HTTP methods that can be executed
// on your bucket. The document is limited to 64 KB in size.
//
// When Amazon S3 receives a cross-origin request (or a pre-flight OPTIONS
// request) against a bucket, it evaluates the cors configuration on the bucket
// and uses the first CORSRule rule that matches the incoming browser request to
// enable a cross-origin request. For a rule to match, the following conditions
// must be met:
//
//   - The request's Origin header must match AllowedOrigin elements.
//
//   - The request method (for example, GET, PUT, HEAD, and so on) or the
//     Access-Control-Request-Method header in case of a pre-flight OPTIONS request
//     must be one of the AllowedMethod elements.
//
//   - Every header specified in the Access-Control-Request-Headers request header
//     of a pre-flight request must match an AllowedHeader element.
//
// For more information about CORS, go to [Enabling Cross-Origin Resource Sharing] in the Amazon S3 User Guide.
//
// The following operations are related to PutBucketCors :
//
// [GetBucketCors]
//
// [DeleteBucketCors]
//
// [RESTOPTIONSobject]
//
// [GetBucketCors]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketCors.html
// [Enabling Cross-Origin Resource Sharing]: https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html
// [RESTOPTIONSobject]: https://docs.aws.amazon.com/AmazonS3/latest/API/RESTOPTIONSobject.html
// [DeleteBucketCors]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketCors.html
func (c *Client) PutBucketCors(ctx context.Context, params *PutBucketCorsInput, optFns ...func(*Options)) (*PutBucketCorsOutput, error) {
	if params == nil {
		params = &PutBucketCorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketCors", params, optFns, c.addOperationPutBucketCorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketCorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketCorsInput struct {

	// Specifies the bucket impacted by the cors configuration.
	//
	// This member is required.
	Bucket *string

	// Describes the cross-origin access configuration for objects in an Amazon S3
	// bucket. For more information, see [Enabling Cross-Origin Resource Sharing]in the Amazon S3 User Guide.
	//
	// [Enabling Cross-Origin Resource Sharing]: https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html
	//
	// This member is required.
	CORSConfiguration *types.CORSConfiguration

	// Indicates the algorithm used to create the checksum for the object when you use
	// the SDK. This header will not provide any additional functionality if you don't
	// use the SDK. When you send this header, there must be a corresponding
	// x-amz-checksum or x-amz-trailer header sent. Otherwise, Amazon S3 fails the
	// request with the HTTP status code 400 Bad Request . For more information, see [Checking object integrity]
	// in the Amazon S3 User Guide.
	//
	// If you provide an individual checksum, Amazon S3 ignores any provided
	// ChecksumAlgorithm parameter.
	//
	// [Checking object integrity]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html
	ChecksumAlgorithm types.ChecksumAlgorithm

	// The base64-encoded 128-bit MD5 digest of the data. This header must be used as
	// a message integrity check to verify that the request body was not corrupted in
	// transit. For more information, go to [RFC 1864.]
	//
	// For requests made using the Amazon Web Services Command Line Interface (CLI) or
	// Amazon Web Services SDKs, this field is calculated automatically.
	//
	// [RFC 1864.]: http://www.ietf.org/rfc/rfc1864.txt
	ContentMD5 *string

	// The account ID of the expected bucket owner. If the account ID that you provide
	// does not match the actual owner of the bucket, the request fails with the HTTP
	// status code 403 Forbidden (access denied).
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

func (in *PutBucketCorsInput) bindEndpointParams(p *EndpointParameters) {
	p.Bucket = in.Bucket
	p.UseS3ExpressControlEndpoint = ptr.Bool(true)
}

type PutBucketCorsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBucketCorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketCors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketCors{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutBucketCors"); err != nil {
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
	if err = addPutBucketContextMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addTimeOffsetDeserializer(stack, c); err != nil {
		return err
	}
	if err = addOpPutBucketCorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketCors(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addPutBucketCorsInputChecksumMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addPutBucketCorsUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
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
	if err = s3cust.AddExpressDefaultChecksumMiddleware(stack); err != nil {
		return err
	}
	return nil
}

func (v *PutBucketCorsInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opPutBucketCors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutBucketCors",
	}
}

// getPutBucketCorsRequestAlgorithmMember gets the request checksum algorithm
// value provided as input.
func getPutBucketCorsRequestAlgorithmMember(input interface{}) (string, bool) {
	in := input.(*PutBucketCorsInput)
	if len(in.ChecksumAlgorithm) == 0 {
		return "", false
	}
	return string(in.ChecksumAlgorithm), true
}

func addPutBucketCorsInputChecksumMiddlewares(stack *middleware.Stack, options Options) error {
	return internalChecksum.AddInputMiddleware(stack, internalChecksum.InputMiddlewareOptions{
		GetAlgorithm:                     getPutBucketCorsRequestAlgorithmMember,
		RequireChecksum:                  true,
		EnableTrailingChecksum:           false,
		EnableComputeSHA256PayloadHash:   true,
		EnableDecodedContentLengthHeader: true,
	})
}

// getPutBucketCorsBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getPutBucketCorsBucketMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketCorsInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addPutBucketCorsUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getPutBucketCorsBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
