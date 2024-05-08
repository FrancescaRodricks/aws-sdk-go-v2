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

// Retrieves the geofence details from a geofence collection.
func (c *Client) GetGeofence(ctx context.Context, params *GetGeofenceInput, optFns ...func(*Options)) (*GetGeofenceOutput, error) {
	if params == nil {
		params = &GetGeofenceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGeofence", params, optFns, c.addOperationGetGeofenceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGeofenceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGeofenceInput struct {

	// The geofence collection storing the target geofence.
	//
	// This member is required.
	CollectionName *string

	// The geofence you're retrieving details for.
	//
	// This member is required.
	GeofenceId *string

	noSmithyDocumentSerde
}

type GetGeofenceOutput struct {

	// The timestamp for when the geofence collection was created in [ISO 8601] format:
	// YYYY-MM-DDThh:mm:ss.sssZ
	//
	// [ISO 8601]: https://www.iso.org/iso-8601-date-and-time-format.html
	//
	// This member is required.
	CreateTime *time.Time

	// The geofence identifier.
	//
	// This member is required.
	GeofenceId *string

	// Contains the geofence geometry details describing a polygon or a circle.
	//
	// This member is required.
	Geometry *types.GeofenceGeometry

	// Identifies the state of the geofence. A geofence will hold one of the following
	// states:
	//
	//   - ACTIVE — The geofence has been indexed by the system.
	//
	//   - PENDING — The geofence is being processed by the system.
	//
	//   - FAILED — The geofence failed to be indexed by the system.
	//
	//   - DELETED — The geofence has been deleted from the system index.
	//
	//   - DELETING — The geofence is being deleted from the system index.
	//
	// This member is required.
	Status *string

	// The timestamp for when the geofence collection was last updated in [ISO 8601] format:
	// YYYY-MM-DDThh:mm:ss.sssZ
	//
	// [ISO 8601]: https://www.iso.org/iso-8601-date-and-time-format.html
	//
	// This member is required.
	UpdateTime *time.Time

	// User defined properties of the geofence. A property is a key-value pair stored
	// with the geofence and added to any geofence event triggered with that geofence.
	//
	// Format: "key" : "value"
	GeofenceProperties map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetGeofenceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetGeofence{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetGeofence{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetGeofence"); err != nil {
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
	if err = addEndpointPrefix_opGetGeofenceMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetGeofenceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetGeofence(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetGeofenceMiddleware struct {
}

func (*endpointPrefix_opGetGeofenceMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetGeofenceMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "geofencing." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetGeofenceMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetGeofenceMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetGeofence(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetGeofence",
	}
}
