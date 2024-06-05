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
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation is not supported by directory buckets.
//
// Enables notifications of specified events for a bucket. For more information
// about event notifications, see [Configuring Event Notifications].
//
// Using this API, you can replace an existing notification configuration. The
// configuration is an XML file that defines the event types that you want Amazon
// S3 to publish and the destination where you want Amazon S3 to publish an event
// notification when it detects an event of the specified type.
//
// By default, your bucket has no event notifications configured. That is, the
// notification configuration will be an empty NotificationConfiguration .
//
// This action replaces the existing notification configuration with the
// configuration you include in the request body.
//
// After Amazon S3 receives this request, it first verifies that any Amazon Simple
// Notification Service (Amazon SNS) or Amazon Simple Queue Service (Amazon SQS)
// destination exists, and that the bucket owner has permission to publish to it by
// sending a test notification. In the case of Lambda destinations, Amazon S3
// verifies that the Lambda function permissions grant Amazon S3 permission to
// invoke the function from the Amazon S3 bucket. For more information, see [Configuring Notifications for Amazon S3 Events].
//
// You can disable notifications by adding the empty NotificationConfiguration
// element.
//
// For more information about the number of event notification configurations that
// you can create per bucket, see [Amazon S3 service quotas]in Amazon Web Services General Reference.
//
// By default, only the bucket owner can configure notifications on a bucket.
// However, bucket owners can use a bucket policy to grant permission to other
// users to set this configuration with the required s3:PutBucketNotification
// permission.
//
// The PUT notification is an atomic operation. For example, suppose your
// notification configuration includes SNS topic, SQS queue, and Lambda function
// configurations. When you send a PUT request with this configuration, Amazon S3
// sends test messages to your SNS topic. If the message fails, the entire PUT
// action will fail, and Amazon S3 will not add the configuration to your bucket.
//
// If the configuration in the request body includes only one TopicConfiguration
// specifying only the s3:ReducedRedundancyLostObject event type, the response
// will also include the x-amz-sns-test-message-id header containing the message
// ID of the test notification sent to the topic.
//
// The following action is related to PutBucketNotificationConfiguration :
//
// [GetBucketNotificationConfiguration]
//
// [Configuring Notifications for Amazon S3 Events]: https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html
// [Amazon S3 service quotas]: https://docs.aws.amazon.com/general/latest/gr/s3.html#limits_s3
// [GetBucketNotificationConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketNotificationConfiguration.html
// [Configuring Event Notifications]: https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html
func (c *Client) PutBucketNotificationConfiguration(ctx context.Context, params *PutBucketNotificationConfigurationInput, optFns ...func(*Options)) (*PutBucketNotificationConfigurationOutput, error) {
	if params == nil {
		params = &PutBucketNotificationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketNotificationConfiguration", params, optFns, c.addOperationPutBucketNotificationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketNotificationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketNotificationConfigurationInput struct {

	// The name of the bucket.
	//
	// This member is required.
	Bucket *string

	// A container for specifying the notification configuration of the bucket. If
	// this element is empty, notifications are turned off for the bucket.
	//
	// This member is required.
	NotificationConfiguration *types.NotificationConfiguration

	// The account ID of the expected bucket owner. If the account ID that you provide
	// does not match the actual owner of the bucket, the request fails with the HTTP
	// status code 403 Forbidden (access denied).
	ExpectedBucketOwner *string

	// Skips validation of Amazon SQS, Amazon SNS, and Lambda destinations. True or
	// false value.
	SkipDestinationValidation *bool

	noSmithyDocumentSerde
}

func (in *PutBucketNotificationConfigurationInput) bindEndpointParams(p *EndpointParameters) {
	p.Bucket = in.Bucket
	p.UseS3ExpressControlEndpoint = ptr.Bool(true)
}

type PutBucketNotificationConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBucketNotificationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketNotificationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketNotificationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutBucketNotificationConfiguration"); err != nil {
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
	if err = addOpPutBucketNotificationConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketNotificationConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addPutBucketNotificationConfigurationUpdateEndpoint(stack, options); err != nil {
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

func (v *PutBucketNotificationConfigurationInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opPutBucketNotificationConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutBucketNotificationConfiguration",
	}
}

// getPutBucketNotificationConfigurationBucketMember returns a pointer to string
// denoting a provided bucket member valueand a boolean indicating if the input has
// a modeled bucket name,
func getPutBucketNotificationConfigurationBucketMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketNotificationConfigurationInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addPutBucketNotificationConfigurationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getPutBucketNotificationConfigurationBucketMember,
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
