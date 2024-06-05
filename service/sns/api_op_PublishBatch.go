// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Publishes up to ten messages to the specified topic. This is a batch version of
// Publish . For FIFO topics, multiple messages within a single batch are published
// in the order they are sent, and messages are deduplicated within the batch and
// across batches for 5 minutes.
//
// The result of publishing each message is reported individually in the response.
// Because the batch request can result in a combination of successful and
// unsuccessful actions, you should check for batch errors even when the call
// returns an HTTP status code of 200 .
//
// The maximum allowed individual message size and the maximum total payload size
// (the sum of the individual lengths of all of the batched messages) are both 256
// KB (262,144 bytes).
//
// Some actions take lists of parameters. These lists are specified using the
// param.n notation. Values of n are integers starting from 1. For example, a
// parameter list with two elements looks like this:
//
// &AttributeName.1=first
//
// &AttributeName.2=second
//
// If you send a batch message to a topic, Amazon SNS publishes the batch message
// to each endpoint that is subscribed to the topic. The format of the batch
// message depends on the notification protocol for each subscribed endpoint.
//
// When a messageId is returned, the batch message is saved and Amazon SNS
// immediately delivers the message to subscribers.
func (c *Client) PublishBatch(ctx context.Context, params *PublishBatchInput, optFns ...func(*Options)) (*PublishBatchOutput, error) {
	if params == nil {
		params = &PublishBatchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PublishBatch", params, optFns, c.addOperationPublishBatchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PublishBatchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PublishBatchInput struct {

	// A list of PublishBatch request entries to be sent to the SNS topic.
	//
	// This member is required.
	PublishBatchRequestEntries []types.PublishBatchRequestEntry

	// The Amazon resource name (ARN) of the topic you want to batch publish to.
	//
	// This member is required.
	TopicArn *string

	noSmithyDocumentSerde
}

type PublishBatchOutput struct {

	// A list of failed PublishBatch responses.
	Failed []types.BatchResultErrorEntry

	// A list of successful PublishBatch responses.
	Successful []types.PublishBatchResultEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPublishBatchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPublishBatch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPublishBatch{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PublishBatch"); err != nil {
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
	if err = addOpPublishBatchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPublishBatch(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPublishBatch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PublishBatch",
	}
}
