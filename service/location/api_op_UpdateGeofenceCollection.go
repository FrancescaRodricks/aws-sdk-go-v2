// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the specified properties of a given geofence collection.
func (c *Client) UpdateGeofenceCollection(ctx context.Context, params *UpdateGeofenceCollectionInput, optFns ...func(*Options)) (*UpdateGeofenceCollectionOutput, error) {
	if params == nil {
		params = &UpdateGeofenceCollectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGeofenceCollection", params, optFns, c.addOperationUpdateGeofenceCollectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGeofenceCollectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGeofenceCollectionInput struct {

	// The name of the geofence collection to update.
	//
	// This member is required.
	CollectionName *string

	// Updates the description for the geofence collection.
	Description *string

	// No longer used. If included, the only allowed value is RequestBasedUsage .
	//
	// Deprecated: Deprecated. If included, the only allowed value is
	// RequestBasedUsage.
	PricingPlan types.PricingPlan

	// This parameter is no longer used.
	//
	// Deprecated: Deprecated. No longer allowed.
	PricingPlanDataSource *string

	noSmithyDocumentSerde
}

type UpdateGeofenceCollectionOutput struct {

	// The Amazon Resource Name (ARN) of the updated geofence collection. Used to
	// specify a resource across Amazon Web Services.
	//
	//   - Format example:
	//   arn:aws:geo:region:account-id:geofence-collection/ExampleGeofenceCollection
	//
	// This member is required.
	CollectionArn *string

	// The name of the updated geofence collection.
	//
	// This member is required.
	CollectionName *string

	// The time when the geofence collection was last updated in [ISO 8601] format:
	// YYYY-MM-DDThh:mm:ss.sssZ
	//
	// [ISO 8601]: https://www.iso.org/iso-8601-date-and-time-format.html
	//
	// This member is required.
	UpdateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateGeofenceCollectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateGeofenceCollection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateGeofenceCollection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateGeofenceCollection"); err != nil {
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
	if err = addEndpointPrefix_opUpdateGeofenceCollectionMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateGeofenceCollectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGeofenceCollection(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opUpdateGeofenceCollectionMiddleware struct {
}

func (*endpointPrefix_opUpdateGeofenceCollectionMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opUpdateGeofenceCollectionMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "cp.geofencing." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opUpdateGeofenceCollectionMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opUpdateGeofenceCollectionMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opUpdateGeofenceCollection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateGeofenceCollection",
	}
}
