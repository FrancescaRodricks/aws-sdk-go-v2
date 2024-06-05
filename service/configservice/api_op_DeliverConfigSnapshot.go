// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Schedules delivery of a configuration snapshot to the Amazon S3 bucket in the
// specified delivery channel. After the delivery has started, Config sends the
// following notifications using an Amazon SNS topic that you have specified.
//
//   - Notification of the start of the delivery.
//
//   - Notification of the completion of the delivery, if the delivery was
//     successfully completed.
//
//   - Notification of delivery failure, if the delivery failed.
func (c *Client) DeliverConfigSnapshot(ctx context.Context, params *DeliverConfigSnapshotInput, optFns ...func(*Options)) (*DeliverConfigSnapshotOutput, error) {
	if params == nil {
		params = &DeliverConfigSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeliverConfigSnapshot", params, optFns, c.addOperationDeliverConfigSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeliverConfigSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DeliverConfigSnapshot action.
type DeliverConfigSnapshotInput struct {

	// The name of the delivery channel through which the snapshot is delivered.
	//
	// This member is required.
	DeliveryChannelName *string

	noSmithyDocumentSerde
}

// The output for the DeliverConfigSnapshot action, in JSON format.
type DeliverConfigSnapshotOutput struct {

	// The ID of the snapshot that is being created.
	ConfigSnapshotId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeliverConfigSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeliverConfigSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeliverConfigSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeliverConfigSnapshot"); err != nil {
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
	if err = addOpDeliverConfigSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeliverConfigSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeliverConfigSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeliverConfigSnapshot",
	}
}
