// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an instance refresh.
//
// This operation is part of the [instance refresh feature] in Amazon EC2 Auto Scaling, which helps you
// update instances in your Auto Scaling group. This feature is helpful, for
// example, when you have a new AMI or a new user data script. You just need to
// create a new launch template that specifies the new AMI or user data script.
// Then start an instance refresh to immediately begin the process of updating
// instances in the group.
//
// If successful, the request's response contains a unique ID that you can use to
// track the progress of the instance refresh. To query its status, call the DescribeInstanceRefreshesAPI.
// To describe the instance refreshes that have already run, call the DescribeInstanceRefreshesAPI. To
// cancel an instance refresh that is in progress, use the CancelInstanceRefreshAPI.
//
// An instance refresh might fail for several reasons, such as EC2 launch
// failures, misconfigured health checks, or not ignoring or allowing the
// termination of instances that are in Standby state or protected from scale in.
// You can monitor for failed EC2 launches using the scaling activities. To find
// the scaling activities, call the DescribeScalingActivitiesAPI.
//
// If you enable auto rollback, your Auto Scaling group will be rolled back
// automatically when the instance refresh fails. You can enable this feature
// before starting an instance refresh by specifying the AutoRollback property in
// the instance refresh preferences. Otherwise, to roll back an instance refresh
// before it finishes, use the RollbackInstanceRefreshAPI.
//
// [instance refresh feature]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-refresh.html
func (c *Client) StartInstanceRefresh(ctx context.Context, params *StartInstanceRefreshInput, optFns ...func(*Options)) (*StartInstanceRefreshOutput, error) {
	if params == nil {
		params = &StartInstanceRefreshInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartInstanceRefresh", params, optFns, c.addOperationStartInstanceRefreshMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartInstanceRefreshOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartInstanceRefreshInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The desired configuration. For example, the desired configuration can specify a
	// new launch template or a new version of the current launch template.
	//
	// Once the instance refresh succeeds, Amazon EC2 Auto Scaling updates the
	// settings of the Auto Scaling group to reflect the new desired configuration.
	//
	// When you specify a new launch template or a new version of the current launch
	// template for your desired configuration, consider enabling the SkipMatching
	// property in preferences. If it's enabled, Amazon EC2 Auto Scaling skips
	// replacing instances that already use the specified launch template and instance
	// types. This can help you reduce the number of replacements that are required to
	// apply updates.
	DesiredConfiguration *types.DesiredConfiguration

	// Sets your preferences for the instance refresh so that it performs as expected
	// when you start it. Includes the instance warmup time, the minimum and maximum
	// healthy percentages, and the behaviors that you want Amazon EC2 Auto Scaling to
	// use if instances that are in Standby state or protected from scale in are
	// found. You can also choose to enable additional features, such as the following:
	//
	//   - Auto rollback
	//
	//   - Checkpoints
	//
	//   - CloudWatch alarms
	//
	//   - Skip matching
	Preferences *types.RefreshPreferences

	// The strategy to use for the instance refresh. The only valid value is Rolling .
	Strategy types.RefreshStrategy

	noSmithyDocumentSerde
}

type StartInstanceRefreshOutput struct {

	// A unique ID for tracking the progress of the instance refresh.
	InstanceRefreshId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartInstanceRefreshMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpStartInstanceRefresh{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpStartInstanceRefresh{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartInstanceRefresh"); err != nil {
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
	if err = addOpStartInstanceRefreshValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartInstanceRefresh(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartInstanceRefresh(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartInstanceRefresh",
	}
}
