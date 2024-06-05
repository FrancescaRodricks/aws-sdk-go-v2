// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the value of one or more queue attributes, like a policy. When you change
// a queue's attributes, the change can take up to 60 seconds for most of the
// attributes to propagate throughout the Amazon SQS system. Changes made to the
// MessageRetentionPeriod attribute can take up to 15 minutes and will impact
// existing messages in the queue potentially causing them to be expired and
// deleted if the MessageRetentionPeriod is reduced below the age of existing
// messages.
//
//   - In the future, new attributes might be added. If you write code that calls
//     this action, we recommend that you structure your code so that it can handle new
//     attributes gracefully.
//
//   - Cross-account permissions don't apply to this action. For more information,
//     see [Grant cross-account permissions to a role and a username]in the Amazon SQS Developer Guide.
//
//   - To remove the ability to change queue permissions, you must deny permission
//     to the AddPermission , RemovePermission , and SetQueueAttributes actions in
//     your IAM policy.
//
// [Grant cross-account permissions to a role and a username]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name
func (c *Client) SetQueueAttributes(ctx context.Context, params *SetQueueAttributesInput, optFns ...func(*Options)) (*SetQueueAttributesOutput, error) {
	if params == nil {
		params = &SetQueueAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetQueueAttributes", params, optFns, c.addOperationSetQueueAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetQueueAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetQueueAttributesInput struct {

	// A map of attributes to set.
	//
	// The following lists the names, descriptions, and values of the special request
	// parameters that the SetQueueAttributes action uses:
	//
	//   - DelaySeconds – The length of time, in seconds, for which the delivery of all
	//   messages in the queue is delayed. Valid values: An integer from 0 to 900 (15
	//   minutes). Default: 0.
	//
	//   - MaximumMessageSize – The limit of how many bytes a message can contain
	//   before Amazon SQS rejects it. Valid values: An integer from 1,024 bytes (1 KiB)
	//   up to 262,144 bytes (256 KiB). Default: 262,144 (256 KiB).
	//
	//   - MessageRetentionPeriod – The length of time, in seconds, for which Amazon
	//   SQS retains a message. Valid values: An integer representing seconds, from 60 (1
	//   minute) to 1,209,600 (14 days). Default: 345,600 (4 days). When you change a
	//   queue's attributes, the change can take up to 60 seconds for most of the
	//   attributes to propagate throughout the Amazon SQS system. Changes made to the
	//   MessageRetentionPeriod attribute can take up to 15 minutes and will impact
	//   existing messages in the queue potentially causing them to be expired and
	//   deleted if the MessageRetentionPeriod is reduced below the age of existing
	//   messages.
	//
	//   - Policy – The queue's policy. A valid Amazon Web Services policy. For more
	//   information about policy structure, see [Overview of Amazon Web Services IAM Policies]in the Identity and Access Management
	//   User Guide.
	//
	//   - ReceiveMessageWaitTimeSeconds – The length of time, in seconds, for which a ReceiveMessage
	//   action waits for a message to arrive. Valid values: An integer from 0 to 20
	//   (seconds). Default: 0.
	//
	//   - VisibilityTimeout – The visibility timeout for the queue, in seconds. Valid
	//   values: An integer from 0 to 43,200 (12 hours). Default: 30. For more
	//   information about the visibility timeout, see [Visibility Timeout]in the Amazon SQS Developer
	//   Guide.
	//
	// The following attributes apply only to [dead-letter queues:]
	//
	//   - RedrivePolicy – The string that includes the parameters for the dead-letter
	//   queue functionality of the source queue as a JSON object. The parameters are as
	//   follows:
	//
	//   - deadLetterTargetArn – The Amazon Resource Name (ARN) of the dead-letter
	//   queue to which Amazon SQS moves messages after the value of maxReceiveCount is
	//   exceeded.
	//
	//   - maxReceiveCount – The number of times a message is delivered to the source
	//   queue before being moved to the dead-letter queue. Default: 10. When the
	//   ReceiveCount for a message exceeds the maxReceiveCount for a queue, Amazon SQS
	//   moves the message to the dead-letter-queue.
	//
	//   - RedriveAllowPolicy – The string that includes the parameters for the
	//   permissions for the dead-letter queue redrive permission and which source queues
	//   can specify dead-letter queues as a JSON object. The parameters are as follows:
	//
	//   - redrivePermission – The permission type that defines which source queues can
	//   specify the current queue as the dead-letter queue. Valid values are:
	//
	//   - allowAll – (Default) Any source queues in this Amazon Web Services account
	//   in the same Region can specify this queue as the dead-letter queue.
	//
	//   - denyAll – No source queues can specify this queue as the dead-letter queue.
	//
	//   - byQueue – Only queues specified by the sourceQueueArns parameter can specify
	//   this queue as the dead-letter queue.
	//
	//   - sourceQueueArns – The Amazon Resource Names (ARN)s of the source queues that
	//   can specify this queue as the dead-letter queue and redrive messages. You can
	//   specify this parameter only when the redrivePermission parameter is set to
	//   byQueue . You can specify up to 10 source queue ARNs. To allow more than 10
	//   source queues to specify dead-letter queues, set the redrivePermission
	//   parameter to allowAll .
	//
	// The dead-letter queue of a FIFO queue must also be a FIFO queue. Similarly, the
	// dead-letter queue of a standard queue must also be a standard queue.
	//
	// The following attributes apply only to [server-side-encryption]:
	//
	//   - KmsMasterKeyId – The ID of an Amazon Web Services managed customer master
	//   key (CMK) for Amazon SQS or a custom CMK. For more information, see [Key Terms]. While
	//   the alias of the AWS-managed CMK for Amazon SQS is always alias/aws/sqs , the
	//   alias of a custom CMK can, for example, be alias/MyAlias . For more examples,
	//   see [KeyId]in the Key Management Service API Reference.
	//
	//   - KmsDataKeyReusePeriodSeconds – The length of time, in seconds, for which
	//   Amazon SQS can reuse a [data key]to encrypt or decrypt messages before calling KMS
	//   again. An integer representing seconds, between 60 seconds (1 minute) and 86,400
	//   seconds (24 hours). Default: 300 (5 minutes). A shorter time period provides
	//   better security but results in more calls to KMS which might incur charges after
	//   Free Tier. For more information, see [How Does the Data Key Reuse Period Work?].
	//
	//   - SqsManagedSseEnabled – Enables server-side queue encryption using SQS owned
	//   encryption keys. Only one server-side encryption option is supported per queue
	//   (for example, [SSE-KMS]or [SSE-SQS]).
	//
	// The following attribute applies only to [FIFO (first-in-first-out) queues]:
	//
	//   - ContentBasedDeduplication – Enables content-based deduplication. For more
	//   information, see [Exactly-once processing]in the Amazon SQS Developer Guide. Note the following:
	//
	//   - Every message must have a unique MessageDeduplicationId .
	//
	//   - You may provide a MessageDeduplicationId explicitly.
	//
	//   - If you aren't able to provide a MessageDeduplicationId and you enable
	//   ContentBasedDeduplication for your queue, Amazon SQS uses a SHA-256 hash to
	//   generate the MessageDeduplicationId using the body of the message (but not the
	//   attributes of the message).
	//
	//   - If you don't provide a MessageDeduplicationId and the queue doesn't have
	//   ContentBasedDeduplication set, the action fails with an error.
	//
	//   - If the queue has ContentBasedDeduplication set, your MessageDeduplicationId
	//   overrides the generated one.
	//
	//   - When ContentBasedDeduplication is in effect, messages with identical content
	//   sent within the deduplication interval are treated as duplicates and only one
	//   copy of the message is delivered.
	//
	//   - If you send one message with ContentBasedDeduplication enabled and then
	//   another message with a MessageDeduplicationId that is the same as the one
	//   generated for the first MessageDeduplicationId , the two messages are treated
	//   as duplicates and only one copy of the message is delivered.
	//
	// The following attributes apply only to [high throughput for FIFO queues]:
	//
	//   - DeduplicationScope – Specifies whether message deduplication occurs at the
	//   message group or queue level. Valid values are messageGroup and queue .
	//
	//   - FifoThroughputLimit – Specifies whether the FIFO queue throughput quota
	//   applies to the entire queue or per message group. Valid values are perQueue
	//   and perMessageGroupId . The perMessageGroupId value is allowed only when the
	//   value for DeduplicationScope is messageGroup .
	//
	// To enable high throughput for FIFO queues, do the following:
	//
	//   - Set DeduplicationScope to messageGroup .
	//
	//   - Set FifoThroughputLimit to perMessageGroupId .
	//
	// If you set these attributes to anything other than the values shown for
	// enabling high throughput, normal throughput is in effect and deduplication
	// occurs as specified.
	//
	// For information on throughput quotas, see [Quotas related to messages] in the Amazon SQS Developer Guide.
	//
	// [SSE-KMS]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sse-existing-queue.html
	// [data key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#data-keys
	// [How Does the Data Key Reuse Period Work?]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-how-does-the-data-key-reuse-period-work
	// [SSE-SQS]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sqs-sse-queue.html
	// [high throughput for FIFO queues]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/high-throughput-fifo.html
	// [Overview of Amazon Web Services IAM Policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/PoliciesOverview.html
	// [dead-letter queues:]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html
	// [Exactly-once processing]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues-exactly-once-processing.html
	// [KeyId]: https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters
	// [Quotas related to messages]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/quotas-messages.html
	// [Visibility Timeout]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html
	// [Key Terms]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms
	// [server-side-encryption]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html
	// [FIFO (first-in-first-out) queues]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html
	//
	// This member is required.
	Attributes map[string]string

	// The URL of the Amazon SQS queue whose attributes are set.
	//
	// Queue URLs and names are case-sensitive.
	//
	// This member is required.
	QueueUrl *string

	noSmithyDocumentSerde
}

type SetQueueAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetQueueAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpSetQueueAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpSetQueueAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetQueueAttributes"); err != nil {
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
	if err = addOpSetQueueAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetQueueAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetQueueAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetQueueAttributes",
	}
}
