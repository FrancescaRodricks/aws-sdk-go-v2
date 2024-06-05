// Code generated by smithy-go-codegen DO NOT EDIT.

package fis

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/fis/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an experiment template.
//
// An experiment template includes the following components:
//
//   - Targets: A target can be a specific resource in your Amazon Web Services
//     environment, or one or more resources that match criteria that you specify, for
//     example, resources that have specific tags.
//
//   - Actions: The actions to carry out on the target. You can specify multiple
//     actions, the duration of each action, and when to start each action during an
//     experiment.
//
//   - Stop conditions: If a stop condition is triggered while an experiment is
//     running, the experiment is automatically stopped. You can define a stop
//     condition as a CloudWatch alarm.
//
// For more information, see [experiment templates] in the Fault Injection Service User Guide.
//
// [experiment templates]: https://docs.aws.amazon.com/fis/latest/userguide/experiment-templates.html
func (c *Client) CreateExperimentTemplate(ctx context.Context, params *CreateExperimentTemplateInput, optFns ...func(*Options)) (*CreateExperimentTemplateOutput, error) {
	if params == nil {
		params = &CreateExperimentTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateExperimentTemplate", params, optFns, c.addOperationCreateExperimentTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateExperimentTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateExperimentTemplateInput struct {

	// The actions for the experiment.
	//
	// This member is required.
	Actions map[string]types.CreateExperimentTemplateActionInput

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	//
	// This member is required.
	ClientToken *string

	// A description for the experiment template.
	//
	// This member is required.
	Description *string

	// The Amazon Resource Name (ARN) of an IAM role that grants the FIS service
	// permission to perform service actions on your behalf.
	//
	// This member is required.
	RoleArn *string

	// The stop conditions.
	//
	// This member is required.
	StopConditions []types.CreateExperimentTemplateStopConditionInput

	// The experiment options for the experiment template.
	ExperimentOptions *types.CreateExperimentTemplateExperimentOptionsInput

	// The configuration for experiment logging.
	LogConfiguration *types.CreateExperimentTemplateLogConfigurationInput

	// The tags to apply to the experiment template.
	Tags map[string]string

	// The targets for the experiment.
	Targets map[string]types.CreateExperimentTemplateTargetInput

	noSmithyDocumentSerde
}

type CreateExperimentTemplateOutput struct {

	// Information about the experiment template.
	ExperimentTemplate *types.ExperimentTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateExperimentTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateExperimentTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateExperimentTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateExperimentTemplate"); err != nil {
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
	if err = addIdempotencyToken_opCreateExperimentTemplateMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateExperimentTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateExperimentTemplate(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateExperimentTemplate struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateExperimentTemplate) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateExperimentTemplate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateExperimentTemplateInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateExperimentTemplateInput ")
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
func addIdempotencyToken_opCreateExperimentTemplateMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateExperimentTemplate{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateExperimentTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateExperimentTemplate",
	}
}
