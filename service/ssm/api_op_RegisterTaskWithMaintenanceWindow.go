// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a new task to a maintenance window.
func (c *Client) RegisterTaskWithMaintenanceWindow(ctx context.Context, params *RegisterTaskWithMaintenanceWindowInput, optFns ...func(*Options)) (*RegisterTaskWithMaintenanceWindowOutput, error) {
	if params == nil {
		params = &RegisterTaskWithMaintenanceWindowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterTaskWithMaintenanceWindow", params, optFns, c.addOperationRegisterTaskWithMaintenanceWindowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterTaskWithMaintenanceWindowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterTaskWithMaintenanceWindowInput struct {

	// The ARN of the task to run.
	//
	// This member is required.
	TaskArn *string

	// The type of task being registered.
	//
	// This member is required.
	TaskType types.MaintenanceWindowTaskType

	// The ID of the maintenance window the task should be added to.
	//
	// This member is required.
	WindowId *string

	// The CloudWatch alarm you want to apply to your maintenance window task.
	AlarmConfiguration *types.AlarmConfiguration

	// User-provided idempotency token.
	ClientToken *string

	// Indicates whether tasks should continue to run after the cutoff time specified
	// in the maintenance windows is reached.
	//
	//   - CONTINUE_TASK : When the cutoff time is reached, any tasks that are running
	//   continue. The default value.
	//
	//   - CANCEL_TASK :
	//
	//   - For Automation, Lambda, Step Functions tasks: When the cutoff time is
	//   reached, any task invocations that are already running continue, but no new task
	//   invocations are started.
	//
	//   - For Run Command tasks: When the cutoff time is reached, the system sends a CancelCommand
	//   operation that attempts to cancel the command associated with the task. However,
	//   there is no guarantee that the command will be terminated and the underlying
	//   process stopped.
	//
	// The status for tasks that are not completed is TIMED_OUT .
	CutoffBehavior types.MaintenanceWindowTaskCutoffBehavior

	// An optional description for the task.
	Description *string

	// A structure containing information about an Amazon Simple Storage Service
	// (Amazon S3) bucket to write managed node-level logs to.
	//
	// LoggingInfo has been deprecated. To specify an Amazon Simple Storage Service
	// (Amazon S3) bucket to contain logs, instead use the OutputS3BucketName and
	// OutputS3KeyPrefix options in the TaskInvocationParameters structure. For
	// information about how Amazon Web Services Systems Manager handles these options
	// for the supported maintenance window task types, see MaintenanceWindowTaskInvocationParameters.
	LoggingInfo *types.LoggingInfo

	// The maximum number of targets this task can be run for, in parallel.
	//
	// Although this element is listed as "Required: No", a value can be omitted only
	// when you are registering or updating a [targetless task]You must provide a value in all other
	// cases.
	//
	// For maintenance window tasks without a target specified, you can't supply a
	// value for this option. Instead, the system inserts a placeholder value of 1 .
	// This value doesn't affect the running of your task.
	//
	// [targetless task]: https://docs.aws.amazon.com/systems-manager/latest/userguide/maintenance-windows-targetless-tasks.html
	MaxConcurrency *string

	// The maximum number of errors allowed before this task stops being scheduled.
	//
	// Although this element is listed as "Required: No", a value can be omitted only
	// when you are registering or updating a [targetless task]You must provide a value in all other
	// cases.
	//
	// For maintenance window tasks without a target specified, you can't supply a
	// value for this option. Instead, the system inserts a placeholder value of 1 .
	// This value doesn't affect the running of your task.
	//
	// [targetless task]: https://docs.aws.amazon.com/systems-manager/latest/userguide/maintenance-windows-targetless-tasks.html
	MaxErrors *string

	// An optional name for the task.
	Name *string

	// The priority of the task in the maintenance window, the lower the number the
	// higher the priority. Tasks in a maintenance window are scheduled in priority
	// order with tasks that have the same priority scheduled in parallel.
	Priority *int32

	// The Amazon Resource Name (ARN) of the IAM service role for Amazon Web Services
	// Systems Manager to assume when running a maintenance window task. If you do not
	// specify a service role ARN, Systems Manager uses a service-linked role in your
	// account. If no appropriate service-linked role for Systems Manager exists in
	// your account, it is created when you run RegisterTaskWithMaintenanceWindow .
	//
	// However, for an improved security posture, we strongly recommend creating a
	// custom policy and custom service role for running your maintenance window tasks.
	// The policy can be crafted to provide only the permissions needed for your
	// particular maintenance window tasks. For more information, see [Setting up maintenance windows]in the in the
	// Amazon Web Services Systems Manager User Guide.
	//
	// [Setting up maintenance windows]: https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-maintenance-permissions.html
	ServiceRoleArn *string

	// The targets (either managed nodes or maintenance window targets).
	//
	// One or more targets must be specified for maintenance window Run Command-type
	// tasks. Depending on the task, targets are optional for other maintenance window
	// task types (Automation, Lambda, and Step Functions). For more information about
	// running tasks that don't specify targets, see [Registering maintenance window tasks without targets]in the Amazon Web Services
	// Systems Manager User Guide.
	//
	// Specify managed nodes using the following format:
	//
	//     Key=InstanceIds,Values=,
	//
	// Specify maintenance window targets using the following format:
	//
	//     Key=WindowTargetIds,Values=,
	//
	// [Registering maintenance window tasks without targets]: https://docs.aws.amazon.com/systems-manager/latest/userguide/maintenance-windows-targetless-tasks.html
	Targets []types.Target

	// The parameters that the task should use during execution. Populate only the
	// fields that match the task type. All other fields should be empty.
	TaskInvocationParameters *types.MaintenanceWindowTaskInvocationParameters

	// The parameters that should be passed to the task when it is run.
	//
	// TaskParameters has been deprecated. To specify parameters to pass to a task
	// when it runs, instead use the Parameters option in the TaskInvocationParameters
	// structure. For information about how Systems Manager handles these options for
	// the supported maintenance window task types, see MaintenanceWindowTaskInvocationParameters.
	TaskParameters map[string]types.MaintenanceWindowTaskParameterValueExpression

	noSmithyDocumentSerde
}

type RegisterTaskWithMaintenanceWindowOutput struct {

	// The ID of the task in the maintenance window.
	WindowTaskId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterTaskWithMaintenanceWindowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterTaskWithMaintenanceWindow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterTaskWithMaintenanceWindow{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterTaskWithMaintenanceWindow"); err != nil {
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
	if err = addIdempotencyToken_opRegisterTaskWithMaintenanceWindowMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpRegisterTaskWithMaintenanceWindowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterTaskWithMaintenanceWindow(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpRegisterTaskWithMaintenanceWindow struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpRegisterTaskWithMaintenanceWindow) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpRegisterTaskWithMaintenanceWindow) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*RegisterTaskWithMaintenanceWindowInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *RegisterTaskWithMaintenanceWindowInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opRegisterTaskWithMaintenanceWindowMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpRegisterTaskWithMaintenanceWindow{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opRegisterTaskWithMaintenanceWindow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterTaskWithMaintenanceWindow",
	}
}
