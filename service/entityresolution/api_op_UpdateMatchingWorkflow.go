// Code generated by smithy-go-codegen DO NOT EDIT.

package entityresolution

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/entityresolution/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing MatchingWorkflow . This method is identical to
// CreateMatchingWorkflow , except it uses an HTTP PUT request instead of a POST
// request, and the MatchingWorkflow must already exist for the method to succeed.
func (c *Client) UpdateMatchingWorkflow(ctx context.Context, params *UpdateMatchingWorkflowInput, optFns ...func(*Options)) (*UpdateMatchingWorkflowOutput, error) {
	if params == nil {
		params = &UpdateMatchingWorkflowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMatchingWorkflow", params, optFns, c.addOperationUpdateMatchingWorkflowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMatchingWorkflowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMatchingWorkflowInput struct {

	// A list of InputSource objects, which have the fields InputSourceARN and
	// SchemaName .
	//
	// This member is required.
	InputSourceConfig []types.InputSource

	// A list of OutputSource objects, each of which contains fields OutputS3Path ,
	// ApplyNormalization , and Output .
	//
	// This member is required.
	OutputSourceConfig []types.OutputSource

	// An object which defines the resolutionType and the ruleBasedProperties .
	//
	// This member is required.
	ResolutionTechniques *types.ResolutionTechniques

	// The Amazon Resource Name (ARN) of the IAM role. Entity Resolution assumes this
	// role to create resources on your behalf as part of workflow execution.
	//
	// This member is required.
	RoleArn *string

	// The name of the workflow to be retrieved.
	//
	// This member is required.
	WorkflowName *string

	// A description of the workflow.
	Description *string

	// An object which defines an incremental run type and has only incrementalRunType
	// as a field.
	IncrementalRunConfig *types.IncrementalRunConfig

	noSmithyDocumentSerde
}

type UpdateMatchingWorkflowOutput struct {

	// A list of InputSource objects, which have the fields InputSourceARN and
	// SchemaName .
	//
	// This member is required.
	InputSourceConfig []types.InputSource

	// A list of OutputSource objects, each of which contains fields OutputS3Path ,
	// ApplyNormalization , and Output .
	//
	// This member is required.
	OutputSourceConfig []types.OutputSource

	// An object which defines the resolutionType and the ruleBasedProperties
	//
	// This member is required.
	ResolutionTechniques *types.ResolutionTechniques

	// The Amazon Resource Name (ARN) of the IAM role. Entity Resolution assumes this
	// role to create resources on your behalf as part of workflow execution.
	//
	// This member is required.
	RoleArn *string

	// The name of the workflow.
	//
	// This member is required.
	WorkflowName *string

	// A description of the workflow.
	Description *string

	// An object which defines an incremental run type and has only incrementalRunType
	// as a field.
	IncrementalRunConfig *types.IncrementalRunConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMatchingWorkflowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateMatchingWorkflow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateMatchingWorkflow{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateMatchingWorkflow"); err != nil {
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
	if err = addOpUpdateMatchingWorkflowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMatchingWorkflow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateMatchingWorkflow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateMatchingWorkflow",
	}
}
