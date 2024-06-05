// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified workflow type.
//
// Note: Prior to deletion, workflow types must first be deprecated.
//
// After a workflow type has been deleted, you cannot create new executions of
// that type. Executions that started before the type was deleted will continue to
// run.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
//   - Use a Resource element with the domain name to limit the action to only
//     specified domains.
//
//   - Use an Action element to allow or deny permission to call this action.
//
//   - Constrain the following parameters by using a Condition element with the
//     appropriate keys.
//
//   - workflowType.name : String constraint. The key is swf:workflowType.name .
//
//   - workflowType.version : String constraint. The key is
//     swf:workflowType.version .
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func (c *Client) DeleteWorkflowType(ctx context.Context, params *DeleteWorkflowTypeInput, optFns ...func(*Options)) (*DeleteWorkflowTypeOutput, error) {
	if params == nil {
		params = &DeleteWorkflowTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteWorkflowType", params, optFns, c.addOperationDeleteWorkflowTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteWorkflowTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteWorkflowTypeInput struct {

	// The name of the domain in which the workflow type is registered.
	//
	// This member is required.
	Domain *string

	// The workflow type to delete.
	//
	// This member is required.
	WorkflowType *types.WorkflowType

	noSmithyDocumentSerde
}

type DeleteWorkflowTypeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteWorkflowTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteWorkflowType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteWorkflowType{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteWorkflowType"); err != nil {
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
	if err = addOpDeleteWorkflowTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteWorkflowType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteWorkflowType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteWorkflowType",
	}
}
