// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation is not supported by directory buckets.
//
// Gets the Object Lock configuration for a bucket. The rule specified in the
// Object Lock configuration will be applied by default to every new object placed
// in the specified bucket. For more information, see [Locking Objects].
//
// The following action is related to GetObjectLockConfiguration :
//
// [GetObjectAttributes]
//
// [GetObjectAttributes]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributes.html
// [Locking Objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html
func (c *Client) GetObjectLockConfiguration(ctx context.Context, params *GetObjectLockConfigurationInput, optFns ...func(*Options)) (*GetObjectLockConfigurationOutput, error) {
	if params == nil {
		params = &GetObjectLockConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetObjectLockConfiguration", params, optFns, c.addOperationGetObjectLockConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetObjectLockConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetObjectLockConfigurationInput struct {

	// The bucket whose Object Lock configuration you want to retrieve.
	//
	// Access points - When you use this action with an access point, you must provide
	// the alias of the access point in place of the bucket name or specify the access
	// point ARN. When using the access point ARN, you must direct requests to the
	// access point hostname. The access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// action with an access point through the Amazon Web Services SDKs, you provide
	// the access point ARN in place of the bucket name. For more information about
	// access point ARNs, see [Using access points]in the Amazon S3 User Guide.
	//
	// [Using access points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-access-points.html
	//
	// This member is required.
	Bucket *string

	// The account ID of the expected bucket owner. If the account ID that you provide
	// does not match the actual owner of the bucket, the request fails with the HTTP
	// status code 403 Forbidden (access denied).
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

func (in *GetObjectLockConfigurationInput) bindEndpointParams(p *EndpointParameters) {
	p.Bucket = in.Bucket

}

type GetObjectLockConfigurationOutput struct {

	// The specified bucket's Object Lock configuration.
	ObjectLockConfiguration *types.ObjectLockConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetObjectLockConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetObjectLockConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetObjectLockConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetObjectLockConfiguration"); err != nil {
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
	if err = addOpGetObjectLockConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetObjectLockConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addGetObjectLockConfigurationUpdateEndpoint(stack, options); err != nil {
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
	return nil
}

func (v *GetObjectLockConfigurationInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opGetObjectLockConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetObjectLockConfiguration",
	}
}

// getGetObjectLockConfigurationBucketMember returns a pointer to string denoting
// a provided bucket member valueand a boolean indicating if the input has a
// modeled bucket name,
func getGetObjectLockConfigurationBucketMember(input interface{}) (*string, bool) {
	in := input.(*GetObjectLockConfigurationInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addGetObjectLockConfigurationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getGetObjectLockConfigurationBucketMember,
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
