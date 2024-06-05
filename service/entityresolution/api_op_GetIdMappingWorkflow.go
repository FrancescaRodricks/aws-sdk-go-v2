// Code generated by smithy-go-codegen DO NOT EDIT.

package entityresolution

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/entityresolution/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the IdMappingWorkflow with a given name, if it exists.
func (c *Client) GetIdMappingWorkflow(ctx context.Context, params *GetIdMappingWorkflowInput, optFns ...func(*Options)) (*GetIdMappingWorkflowOutput, error) {
	if params == nil {
		params = &GetIdMappingWorkflowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIdMappingWorkflow", params, optFns, c.addOperationGetIdMappingWorkflowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIdMappingWorkflowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIdMappingWorkflowInput struct {

	// The name of the workflow.
	//
	// This member is required.
	WorkflowName *string

	noSmithyDocumentSerde
}

type GetIdMappingWorkflowOutput struct {

	// The timestamp of when the workflow was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// An object which defines the idMappingType and the providerProperties .
	//
	// This member is required.
	IdMappingTechniques *types.IdMappingTechniques

	// A list of InputSource objects, which have the fields InputSourceARN and
	// SchemaName .
	//
	// This member is required.
	InputSourceConfig []types.IdMappingWorkflowInputSource

	// The Amazon Resource Name (ARN) of the IAM role. Entity Resolution assumes this
	// role to access Amazon Web Services resources on your behalf.
	//
	// This member is required.
	RoleArn *string

	// The timestamp of when the workflow was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The ARN (Amazon Resource Name) that Entity Resolution generated for the
	// IdMappingWorkflow .
	//
	// This member is required.
	WorkflowArn *string

	// The name of the workflow.
	//
	// This member is required.
	WorkflowName *string

	// A description of the workflow.
	Description *string

	// A list of OutputSource objects, each of which contains fields OutputS3Path and
	// KMSArn .
	OutputSourceConfig []types.IdMappingWorkflowOutputSource

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetIdMappingWorkflowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetIdMappingWorkflow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetIdMappingWorkflow{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetIdMappingWorkflow"); err != nil {
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
	if err = addOpGetIdMappingWorkflowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIdMappingWorkflow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetIdMappingWorkflow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetIdMappingWorkflow",
	}
}
