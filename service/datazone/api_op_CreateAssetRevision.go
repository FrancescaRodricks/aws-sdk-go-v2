// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a revision of the asset.
func (c *Client) CreateAssetRevision(ctx context.Context, params *CreateAssetRevisionInput, optFns ...func(*Options)) (*CreateAssetRevisionOutput, error) {
	if params == nil {
		params = &CreateAssetRevisionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAssetRevision", params, optFns, c.addOperationCreateAssetRevisionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAssetRevisionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAssetRevisionInput struct {

	// The unique identifier of the domain where the asset is being revised.
	//
	// This member is required.
	DomainIdentifier *string

	// The identifier of the asset.
	//
	// This member is required.
	Identifier *string

	// Te revised name of the asset.
	//
	// This member is required.
	Name *string

	// A unique, case-sensitive identifier that is provided to ensure the idempotency
	// of the request.
	ClientToken *string

	// The revised description of the asset.
	Description *string

	// The metadata forms to be attached to the asset as part of asset revision.
	FormsInput []types.FormInput

	// The glossary terms to be attached to the asset as part of asset revision.
	GlossaryTerms []string

	// The configuration of the automatically generated business-friendly metadata for
	// the asset.
	PredictionConfiguration *types.PredictionConfiguration

	// The revision type of the asset.
	TypeRevision *string

	noSmithyDocumentSerde
}

type CreateAssetRevisionOutput struct {

	// The unique identifier of the Amazon DataZone domain where the asset was revised.
	//
	// This member is required.
	DomainId *string

	// The metadata forms that were attached to the asset as part of the asset
	// revision.
	//
	// This member is required.
	FormsOutput []types.FormOutput

	// The unique identifier of the asset revision.
	//
	// This member is required.
	Id *string

	// The revised name of the asset.
	//
	// This member is required.
	Name *string

	// The unique identifier of the revised project that owns the asset.
	//
	// This member is required.
	OwningProjectId *string

	// The revision of the asset.
	//
	// This member is required.
	Revision *string

	// The identifier of the revision type.
	//
	// This member is required.
	TypeIdentifier *string

	// The revision type of the asset.
	//
	// This member is required.
	TypeRevision *string

	// The timestamp of when the asset revision occured.
	CreatedAt *time.Time

	// The Amazon DataZone user who performed the asset revision.
	CreatedBy *string

	// The revised asset description.
	Description *string

	// The external identifier of the asset.
	ExternalIdentifier *string

	// The timestamp of when the first asset revision occured.
	FirstRevisionCreatedAt *time.Time

	// The Amazon DataZone user who performed the first asset revision.
	FirstRevisionCreatedBy *string

	// The glossary terms that were attached to the asset as part of asset revision.
	GlossaryTerms []string

	// The latest data point that was imported into the time series form for the
	// asset.
	LatestTimeSeriesDataPointFormsOutput []types.TimeSeriesDataPointSummaryFormOutput

	// The details of an asset published in an Amazon DataZone catalog.
	Listing *types.AssetListingDetails

	// The configuration of the automatically generated business-friendly metadata for
	// the asset.
	PredictionConfiguration *types.PredictionConfiguration

	// The read-only metadata forms that were attached to the asset as part of the
	// asset revision.
	ReadOnlyFormsOutput []types.FormOutput

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAssetRevisionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAssetRevision{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAssetRevision{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAssetRevision"); err != nil {
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
	if err = addIdempotencyToken_opCreateAssetRevisionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAssetRevisionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAssetRevision(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateAssetRevision struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAssetRevision) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAssetRevision) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAssetRevisionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAssetRevisionInput ")
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
func addIdempotencyToken_opCreateAssetRevisionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateAssetRevision{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAssetRevision(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAssetRevision",
	}
}
