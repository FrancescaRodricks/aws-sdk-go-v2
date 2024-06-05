// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies an account setting for all users on an account for whom no individual
// account setting has been specified. Account settings are set on a per-Region
// basis.
func (c *Client) PutAccountSettingDefault(ctx context.Context, params *PutAccountSettingDefaultInput, optFns ...func(*Options)) (*PutAccountSettingDefaultOutput, error) {
	if params == nil {
		params = &PutAccountSettingDefaultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAccountSettingDefault", params, optFns, c.addOperationPutAccountSettingDefaultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAccountSettingDefaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAccountSettingDefaultInput struct {

	// The resource name for which to modify the account setting.
	//
	// The following are the valid values for the account setting name.
	//
	//   - serviceLongArnFormat - When modified, the Amazon Resource Name (ARN) and
	//   resource ID format of the resource type for a specified user, role, or the root
	//   user for an account is affected. The opt-in and opt-out account setting must be
	//   set for each Amazon ECS resource separately. The ARN and resource ID format of a
	//   resource is defined by the opt-in status of the user or role that created the
	//   resource. You must turn on this setting to use Amazon ECS features such as
	//   resource tagging.
	//
	//   - taskLongArnFormat - When modified, the Amazon Resource Name (ARN) and
	//   resource ID format of the resource type for a specified user, role, or the root
	//   user for an account is affected. The opt-in and opt-out account setting must be
	//   set for each Amazon ECS resource separately. The ARN and resource ID format of a
	//   resource is defined by the opt-in status of the user or role that created the
	//   resource. You must turn on this setting to use Amazon ECS features such as
	//   resource tagging.
	//
	//   - containerInstanceLongArnFormat - When modified, the Amazon Resource Name
	//   (ARN) and resource ID format of the resource type for a specified user, role, or
	//   the root user for an account is affected. The opt-in and opt-out account setting
	//   must be set for each Amazon ECS resource separately. The ARN and resource ID
	//   format of a resource is defined by the opt-in status of the user or role that
	//   created the resource. You must turn on this setting to use Amazon ECS features
	//   such as resource tagging.
	//
	//   - awsvpcTrunking - When modified, the elastic network interface (ENI) limit
	//   for any new container instances that support the feature is changed. If
	//   awsvpcTrunking is turned on, any new container instances that support the
	//   feature are launched have the increased ENI limits available to them. For more
	//   information, see [Elastic Network Interface Trunking]in the Amazon Elastic Container Service Developer Guide.
	//
	//   - containerInsights - When modified, the default setting indicating whether
	//   Amazon Web Services CloudWatch Container Insights is turned on for your clusters
	//   is changed. If containerInsights is turned on, any new clusters that are
	//   created will have Container Insights turned on unless you disable it during
	//   cluster creation. For more information, see [CloudWatch Container Insights]in the Amazon Elastic Container
	//   Service Developer Guide.
	//
	//   - dualStackIPv6 - When turned on, when using a VPC in dual stack mode, your
	//   tasks using the awsvpc network mode can have an IPv6 address assigned. For
	//   more information on using IPv6 with tasks launched on Amazon EC2 instances, see [Using a VPC in dual-stack mode]
	//   . For more information on using IPv6 with tasks launched on Fargate, see [Using a VPC in dual-stack mode].
	//
	//   - fargateFIPSMode - If you specify fargateFIPSMode , Fargate FIPS 140
	//   compliance is affected.
	//
	//   - fargateTaskRetirementWaitPeriod - When Amazon Web Services determines that a
	//   security or infrastructure update is needed for an Amazon ECS task hosted on
	//   Fargate, the tasks need to be stopped and new tasks launched to replace them.
	//   Use fargateTaskRetirementWaitPeriod to configure the wait time to retire a
	//   Fargate task. For information about the Fargate tasks maintenance, see [Amazon Web Services Fargate task maintenance]in the
	//   Amazon ECS Developer Guide.
	//
	//   - tagResourceAuthorization - Amazon ECS is introducing tagging authorization
	//   for resource creation. Users must have permissions for actions that create the
	//   resource, such as ecsCreateCluster . If tags are specified when you create a
	//   resource, Amazon Web Services performs additional authorization to verify if
	//   users or roles have permissions to create tags. Therefore, you must grant
	//   explicit permissions to use the ecs:TagResource action. For more information,
	//   see [Grant permission to tag resources on creation]in the Amazon ECS Developer Guide.
	//
	//   - guardDutyActivate - The guardDutyActivate parameter is read-only in Amazon
	//   ECS and indicates whether Amazon ECS Runtime Monitoring is enabled or disabled
	//   by your security administrator in your Amazon ECS account. Amazon GuardDuty
	//   controls this account setting on your behalf. For more information, see [Protecting Amazon ECS workloads with Amazon ECS Runtime Monitoring].
	//
	// [Grant permission to tag resources on creation]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/supported-iam-actions-tagging.html
	// [Using a VPC in dual-stack mode]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/fargate-task-networking.html#fargate-task-networking-vpc-dual-stack
	// [Amazon Web Services Fargate task maintenance]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-maintenance.html
	// [Elastic Network Interface Trunking]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-instance-eni.html
	// [Protecting Amazon ECS workloads with Amazon ECS Runtime Monitoring]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-guard-duty-integration.html
	// [CloudWatch Container Insights]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cloudwatch-container-insights.html
	//
	// This member is required.
	Name types.SettingName

	// The account setting value for the specified principal ARN. Accepted values are
	// enabled , disabled , on , and off .
	//
	// When you specify fargateTaskRetirementWaitPeriod for the name , the following
	// are the valid values:
	//
	//   - 0 - Amazon Web Services sends the notification, and immediately retires the
	//   affected tasks.
	//
	//   - 7 - Amazon Web Services sends the notification, and waits 7 calendar days to
	//   retire the tasks.
	//
	//   - 14 - Amazon Web Services sends the notification, and waits 14 calendar days
	//   to retire the tasks.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type PutAccountSettingDefaultOutput struct {

	// The current setting for a resource.
	Setting *types.Setting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAccountSettingDefaultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutAccountSettingDefault{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAccountSettingDefault{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutAccountSettingDefault"); err != nil {
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
	if err = addOpPutAccountSettingDefaultValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAccountSettingDefault(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutAccountSettingDefault(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutAccountSettingDefault",
	}
}
