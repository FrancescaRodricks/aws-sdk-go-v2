// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified delivery stream and its status. For example, after your
// delivery stream is created, call DescribeDeliveryStream to see whether the
// delivery stream is ACTIVE and therefore ready for data to be sent to it.
//
// If the status of a delivery stream is CREATING_FAILED , this status doesn't
// change, and you can't invoke CreateDeliveryStreamagain on it. However, you can invoke the DeleteDeliveryStream
// operation to delete it. If the status is DELETING_FAILED , you can force
// deletion by invoking DeleteDeliveryStreamagain but with DeleteDeliveryStreamInput$AllowForceDelete set to true.
func (c *Client) DescribeDeliveryStream(ctx context.Context, params *DescribeDeliveryStreamInput, optFns ...func(*Options)) (*DescribeDeliveryStreamOutput, error) {
	if params == nil {
		params = &DescribeDeliveryStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDeliveryStream", params, optFns, c.addOperationDescribeDeliveryStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDeliveryStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDeliveryStreamInput struct {

	// The name of the delivery stream.
	//
	// This member is required.
	DeliveryStreamName *string

	// The ID of the destination to start returning the destination information.
	// Firehose supports one destination per delivery stream.
	ExclusiveStartDestinationId *string

	// The limit on the number of destinations to return. You can have one destination
	// per delivery stream.
	Limit *int32

	noSmithyDocumentSerde
}

type DescribeDeliveryStreamOutput struct {

	// Information about the delivery stream.
	//
	// This member is required.
	DeliveryStreamDescription *types.DeliveryStreamDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDeliveryStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDeliveryStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDeliveryStream{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDeliveryStream"); err != nil {
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
	if err = addOpDescribeDeliveryStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDeliveryStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDeliveryStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDeliveryStream",
	}
}
