// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a composite model and all of the assets that were created from the
// model. Each asset created from the model inherits the updated asset model's
// property and hierarchy definitions. For more information, see [Updating assets and models]in the IoT
// SiteWise User Guide.
//
// If you remove a property from a composite asset model, IoT SiteWise deletes all
// previous data for that property. You can’t change the type or data type of an
// existing property.
//
// To replace an existing composite asset model property with a new one with the
// same name , do the following:
//
//   - Submit an UpdateAssetModelCompositeModel request with the entire existing
//     property removed.
//
//   - Submit a second UpdateAssetModelCompositeModel request that includes the new
//     property. The new asset property will have the same name as the previous one
//     and IoT SiteWise will generate a new unique id .
//
// [Updating assets and models]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/update-assets-and-models.html
func (c *Client) UpdateAssetModelCompositeModel(ctx context.Context, params *UpdateAssetModelCompositeModelInput, optFns ...func(*Options)) (*UpdateAssetModelCompositeModelOutput, error) {
	if params == nil {
		params = &UpdateAssetModelCompositeModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAssetModelCompositeModel", params, optFns, c.addOperationUpdateAssetModelCompositeModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAssetModelCompositeModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAssetModelCompositeModelInput struct {

	// The ID of a composite model on this asset model.
	//
	// This member is required.
	AssetModelCompositeModelId *string

	// A unique, friendly name for the composite model.
	//
	// This member is required.
	AssetModelCompositeModelName *string

	// The ID of the asset model, in UUID format.
	//
	// This member is required.
	AssetModelId *string

	// A description for the composite model.
	AssetModelCompositeModelDescription *string

	// An external ID to assign to the asset model. You can only set the external ID
	// of the asset model if it wasn't set when it was created, or you're setting it to
	// the exact same thing as when it was created.
	AssetModelCompositeModelExternalId *string

	// The property definitions of the composite model. For more information, see .
	//
	// You can specify up to 200 properties per composite model. For more information,
	// see [Quotas]in the IoT SiteWise User Guide.
	//
	// [Quotas]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html
	AssetModelCompositeModelProperties []types.AssetModelProperty

	// A unique case-sensitive identifier that you can provide to ensure the
	// idempotency of the request. Don't reuse this client token if a new idempotent
	// request is required.
	ClientToken *string

	noSmithyDocumentSerde
}

type UpdateAssetModelCompositeModelOutput struct {

	// The path to the composite model listing the parent composite models.
	//
	// This member is required.
	AssetModelCompositeModelPath []types.AssetModelCompositeModelPathSegment

	// Contains current status information for an asset model. For more information,
	// see [Asset and model states]in the IoT SiteWise User Guide.
	//
	// [Asset and model states]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-and-model-states.html
	//
	// This member is required.
	AssetModelStatus *types.AssetModelStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAssetModelCompositeModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAssetModelCompositeModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAssetModelCompositeModel{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateAssetModelCompositeModel"); err != nil {
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
	if err = addEndpointPrefix_opUpdateAssetModelCompositeModelMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opUpdateAssetModelCompositeModelMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateAssetModelCompositeModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAssetModelCompositeModel(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opUpdateAssetModelCompositeModelMiddleware struct {
}

func (*endpointPrefix_opUpdateAssetModelCompositeModelMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opUpdateAssetModelCompositeModelMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opUpdateAssetModelCompositeModelMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opUpdateAssetModelCompositeModelMiddleware{}, "ResolveEndpointV2", middleware.After)
}

type idempotencyToken_initializeOpUpdateAssetModelCompositeModel struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateAssetModelCompositeModel) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateAssetModelCompositeModel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateAssetModelCompositeModelInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateAssetModelCompositeModelInput ")
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
func addIdempotencyToken_opUpdateAssetModelCompositeModelMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateAssetModelCompositeModel{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateAssetModelCompositeModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateAssetModelCompositeModel",
	}
}
