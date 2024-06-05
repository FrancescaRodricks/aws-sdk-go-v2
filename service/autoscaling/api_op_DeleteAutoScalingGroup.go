// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified Auto Scaling group.
//
// If the group has instances or scaling activities in progress, you must specify
// the option to force the deletion in order for it to succeed. The force delete
// operation will also terminate the EC2 instances. If the group has a warm pool,
// the force delete option also deletes the warm pool.
//
// To remove instances from the Auto Scaling group before deleting it, call the DetachInstances
// API with the list of instances and the option to decrement the desired capacity.
// This ensures that Amazon EC2 Auto Scaling does not launch replacement instances.
//
// To terminate all instances before deleting the Auto Scaling group, call the UpdateAutoScalingGroup
// API and set the minimum size and desired capacity of the Auto Scaling group to
// zero.
//
// If the group has scaling policies, deleting the group deletes the policies, the
// underlying alarm actions, and any alarm that no longer has an associated action.
//
// For more information, see [Delete your Auto Scaling infrastructure] in the Amazon EC2 Auto Scaling User Guide.
//
// [Delete your Auto Scaling infrastructure]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-process-shutdown.html
func (c *Client) DeleteAutoScalingGroup(ctx context.Context, params *DeleteAutoScalingGroupInput, optFns ...func(*Options)) (*DeleteAutoScalingGroupOutput, error) {
	if params == nil {
		params = &DeleteAutoScalingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAutoScalingGroup", params, optFns, c.addOperationDeleteAutoScalingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAutoScalingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAutoScalingGroupInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// Specifies that the group is to be deleted along with all instances associated
	// with the group, without waiting for all instances to be terminated. This action
	// also deletes any outstanding lifecycle actions associated with the group.
	ForceDelete *bool

	noSmithyDocumentSerde
}

type DeleteAutoScalingGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteAutoScalingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteAutoScalingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteAutoScalingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteAutoScalingGroup"); err != nil {
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
	if err = addOpDeleteAutoScalingGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAutoScalingGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteAutoScalingGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteAutoScalingGroup",
	}
}
