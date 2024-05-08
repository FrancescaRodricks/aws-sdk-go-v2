// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The TagLogGroup operation is on the path to deprecation. We recommend that you
// use [TagResource]instead.
//
// Adds or updates the specified tags for the specified log group.
//
// To list the tags for a log group, use [ListTagsForResource]. To remove tags, use [UntagResource].
//
// For more information about tags, see [Tag Log Groups in Amazon CloudWatch Logs] in the Amazon CloudWatch Logs User Guide.
//
// CloudWatch Logs doesn’t support IAM policies that prevent users from assigning
// specified tags to log groups using the aws:Resource/key-name  or aws:TagKeys
// condition keys. For more information about using tags to control access, see [Controlling access to Amazon Web Services resources using tags].
//
// Deprecated: Please use the generic tagging API TagResource
//
// [TagResource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_UntagResource.html
// [Tag Log Groups in Amazon CloudWatch Logs]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Working-with-log-groups-and-streams.html#log-group-tagging
// [Controlling access to Amazon Web Services resources using tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html
// [ListTagsForResource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ListTagsForResource.html
func (c *Client) TagLogGroup(ctx context.Context, params *TagLogGroupInput, optFns ...func(*Options)) (*TagLogGroupOutput, error) {
	if params == nil {
		params = &TagLogGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TagLogGroup", params, optFns, c.addOperationTagLogGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TagLogGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagLogGroupInput struct {

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string

	// The key-value pairs to use for the tags.
	//
	// This member is required.
	Tags map[string]string

	noSmithyDocumentSerde
}

type TagLogGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTagLogGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTagLogGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTagLogGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "TagLogGroup"); err != nil {
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
	if err = addOpTagLogGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTagLogGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTagLogGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "TagLogGroup",
	}
}
