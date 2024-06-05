// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves details about a specific a maintenance window execution.
func (c *Client) GetMaintenanceWindowExecution(ctx context.Context, params *GetMaintenanceWindowExecutionInput, optFns ...func(*Options)) (*GetMaintenanceWindowExecutionOutput, error) {
	if params == nil {
		params = &GetMaintenanceWindowExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMaintenanceWindowExecution", params, optFns, c.addOperationGetMaintenanceWindowExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMaintenanceWindowExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMaintenanceWindowExecutionInput struct {

	// The ID of the maintenance window execution that includes the task.
	//
	// This member is required.
	WindowExecutionId *string

	noSmithyDocumentSerde
}

type GetMaintenanceWindowExecutionOutput struct {

	// The time the maintenance window finished running.
	EndTime *time.Time

	// The time the maintenance window started running.
	StartTime *time.Time

	// The status of the maintenance window execution.
	Status types.MaintenanceWindowExecutionStatus

	// The details explaining the status. Not available for all status values.
	StatusDetails *string

	// The ID of the task executions from the maintenance window execution.
	TaskIds []string

	// The ID of the maintenance window execution.
	WindowExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMaintenanceWindowExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMaintenanceWindowExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMaintenanceWindowExecution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMaintenanceWindowExecution"); err != nil {
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
	if err = addOpGetMaintenanceWindowExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMaintenanceWindowExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMaintenanceWindowExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMaintenanceWindowExecution",
	}
}
