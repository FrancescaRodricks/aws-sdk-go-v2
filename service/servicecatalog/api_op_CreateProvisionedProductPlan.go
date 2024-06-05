// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a plan.
//
// A plan includes the list of resources to be created (when provisioning a new
// product) or modified (when updating a provisioned product) when the plan is
// executed.
//
// You can create one plan for each provisioned product. To create a plan for an
// existing provisioned product, the product status must be AVAILABLE or TAINTED.
//
// To view the resource changes in the change set, use DescribeProvisionedProductPlan. To create or modify the
// provisioned product, use ExecuteProvisionedProductPlan.
func (c *Client) CreateProvisionedProductPlan(ctx context.Context, params *CreateProvisionedProductPlanInput, optFns ...func(*Options)) (*CreateProvisionedProductPlanOutput, error) {
	if params == nil {
		params = &CreateProvisionedProductPlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProvisionedProductPlan", params, optFns, c.addOperationCreateProvisionedProductPlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProvisionedProductPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProvisionedProductPlanInput struct {

	// A unique identifier that you provide to ensure idempotency. If multiple
	// requests differ only by the idempotency token, the same response is returned for
	// each repeated request.
	//
	// This member is required.
	IdempotencyToken *string

	// The name of the plan.
	//
	// This member is required.
	PlanName *string

	// The plan type.
	//
	// This member is required.
	PlanType types.ProvisionedProductPlanType

	// The product identifier.
	//
	// This member is required.
	ProductId *string

	// A user-friendly name for the provisioned product. This value must be unique for
	// the Amazon Web Services account and cannot be updated after the product is
	// provisioned.
	//
	// This member is required.
	ProvisionedProductName *string

	// The identifier of the provisioning artifact.
	//
	// This member is required.
	ProvisioningArtifactId *string

	// The language code.
	//
	//   - jp - Japanese
	//
	//   - zh - Chinese
	AcceptLanguage *string

	// Passed to CloudFormation. The SNS topic ARNs to which to publish stack-related
	// events.
	NotificationArns []string

	// The path identifier of the product. This value is optional if the product has a
	// default path, and required if the product has more than one path. To list the
	// paths for a product, use ListLaunchPaths.
	PathId *string

	// Parameters specified by the administrator that are required for provisioning
	// the product.
	ProvisioningParameters []types.UpdateProvisioningParameter

	// One or more tags.
	//
	// If the plan is for an existing provisioned product, the product must have a
	// RESOURCE_UPDATE constraint with TagUpdatesOnProvisionedProduct set to ALLOWED
	// to allow tag updates.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateProvisionedProductPlanOutput struct {

	// The plan identifier.
	PlanId *string

	// The name of the plan.
	PlanName *string

	// The product identifier.
	ProvisionProductId *string

	// The user-friendly name of the provisioned product.
	ProvisionedProductName *string

	// The identifier of the provisioning artifact.
	ProvisioningArtifactId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProvisionedProductPlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateProvisionedProductPlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateProvisionedProductPlan{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateProvisionedProductPlan"); err != nil {
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
	if err = addIdempotencyToken_opCreateProvisionedProductPlanMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateProvisionedProductPlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProvisionedProductPlan(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateProvisionedProductPlan struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateProvisionedProductPlan) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateProvisionedProductPlan) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateProvisionedProductPlanInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateProvisionedProductPlanInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateProvisionedProductPlanMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateProvisionedProductPlan{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateProvisionedProductPlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateProvisionedProductPlan",
	}
}
