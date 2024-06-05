// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides information about a state machine execution, such as the state machine
// associated with the execution, the execution input and output, and relevant
// execution metadata. If you've [redriven]an execution, you can use this API action to
// return information about the redrives of that execution. In addition, you can
// use this API action to return the Map Run Amazon Resource Name (ARN) if the
// execution was dispatched by a Map Run.
//
// If you specify a version or alias ARN when you call the StartExecution API action,
// DescribeExecution returns that ARN.
//
// This operation is eventually consistent. The results are best effort and may
// not reflect very recent updates and changes.
//
// Executions of an EXPRESS state machine aren't supported by DescribeExecution
// unless a Map Run dispatched them.
//
// [redriven]: https://docs.aws.amazon.com/step-functions/latest/dg/redrive-executions.html
func (c *Client) DescribeExecution(ctx context.Context, params *DescribeExecutionInput, optFns ...func(*Options)) (*DescribeExecutionOutput, error) {
	if params == nil {
		params = &DescribeExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeExecution", params, optFns, c.addOperationDescribeExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeExecutionInput struct {

	// The Amazon Resource Name (ARN) of the execution to describe.
	//
	// This member is required.
	ExecutionArn *string

	noSmithyDocumentSerde
}

type DescribeExecutionOutput struct {

	// The Amazon Resource Name (ARN) that identifies the execution.
	//
	// This member is required.
	ExecutionArn *string

	// The date the execution is started.
	//
	// This member is required.
	StartDate *time.Time

	// The Amazon Resource Name (ARN) of the executed stated machine.
	//
	// This member is required.
	StateMachineArn *string

	// The current status of the execution.
	//
	// This member is required.
	Status types.ExecutionStatus

	// The cause string if the state machine execution failed.
	Cause *string

	// The error string if the state machine execution failed.
	Error *string

	// The string that contains the JSON input data of the execution. Length
	// constraints apply to the payload size, and are expressed as bytes in UTF-8
	// encoding.
	Input *string

	// Provides details about execution input or output.
	InputDetails *types.CloudWatchEventsExecutionDataDetails

	// The Amazon Resource Name (ARN) that identifies a Map Run, which dispatched this
	// execution.
	MapRunArn *string

	// The name of the execution.
	//
	// A name must not contain:
	//
	//   - white space
	//
	//   - brackets < > { } [ ]
	//
	//   - wildcard characters ? *
	//
	//   - special characters " # % \ ^ | ~ ` $ & , ; : /
	//
	//   - control characters ( U+0000-001F , U+007F-009F )
	//
	// To enable logging with CloudWatch Logs, the name should only contain 0-9, A-Z,
	// a-z, - and _.
	Name *string

	// The JSON output data of the execution. Length constraints apply to the payload
	// size, and are expressed as bytes in UTF-8 encoding.
	//
	// This field is set only if the execution succeeds. If the execution fails, this
	// field is null.
	Output *string

	// Provides details about execution input or output.
	OutputDetails *types.CloudWatchEventsExecutionDataDetails

	// The number of times you've redriven an execution. If you have not yet redriven
	// an execution, the redriveCount is 0. This count is only updated if you
	// successfully redrive an execution.
	RedriveCount *int32

	// The date the execution was last redriven. If you have not yet redriven an
	// execution, the redriveDate is null.
	//
	// The redriveDate is unavailable if you redrive a Map Run that starts child
	// workflow executions of type EXPRESS .
	RedriveDate *time.Time

	// Indicates whether or not an execution can be redriven at a given point in time.
	//
	//   - For executions of type STANDARD , redriveStatus is NOT_REDRIVABLE if calling
	//   the RedriveExecutionAPI action would return the ExecutionNotRedrivable error.
	//
	//   - For a Distributed Map that includes child workflows of type STANDARD ,
	//   redriveStatus indicates whether or not the Map Run can redrive child workflow
	//   executions.
	//
	//   - For a Distributed Map that includes child workflows of type EXPRESS ,
	//   redriveStatus indicates whether or not the Map Run can redrive child workflow
	//   executions.
	//
	// You can redrive failed or timed out EXPRESS workflows only if they're a part of
	//   a Map Run. When you [redrive]the Map Run, these workflows are restarted using the StartExecutionAPI
	//   action.
	//
	// [redrive]: https://docs.aws.amazon.com/step-functions/latest/dg/redrive-map-run.html
	RedriveStatus types.ExecutionRedriveStatus

	// When redriveStatus is NOT_REDRIVABLE , redriveStatusReason specifies the reason
	// why an execution cannot be redriven.
	//
	//   - For executions of type STANDARD , or for a Distributed Map that includes
	//   child workflows of type STANDARD , redriveStatusReason can include one of the
	//   following reasons:
	//
	//   - State machine is in DELETING status .
	//
	//   - Execution is RUNNING and cannot be redriven .
	//
	//   - Execution is SUCCEEDED and cannot be redriven .
	//
	//   - Execution was started before the launch of RedriveExecution .
	//
	//   - Execution history event limit exceeded .
	//
	//   - Execution has exceeded the max execution time .
	//
	//   - Execution redrivable period exceeded .
	//
	//   - For a Distributed Map that includes child workflows of type EXPRESS ,
	//   redriveStatusReason is only returned if the child workflows are not
	//   redrivable. This happens when the child workflow executions have completed
	//   successfully.
	RedriveStatusReason *string

	// The Amazon Resource Name (ARN) of the state machine alias associated with the
	// execution. The alias ARN is a combination of state machine ARN and the alias
	// name separated by a colon (:). For example, stateMachineARN:PROD .
	//
	// If you start an execution from a StartExecution request with a state machine
	// version ARN, this field will be null.
	StateMachineAliasArn *string

	// The Amazon Resource Name (ARN) of the state machine version associated with the
	// execution. The version ARN is a combination of state machine ARN and the version
	// number separated by a colon (:). For example, stateMachineARN:1 .
	//
	// If you start an execution from a StartExecution request without specifying a
	// state machine version or alias ARN, Step Functions returns a null value.
	StateMachineVersionArn *string

	// If the execution ended, the date the execution stopped.
	StopDate *time.Time

	// The X-Ray trace header that was passed to the execution.
	TraceHeader *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeExecution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeExecution"); err != nil {
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
	if err = addOpDescribeExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeExecution",
	}
}
