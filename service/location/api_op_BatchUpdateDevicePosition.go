// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Uploads position update data for one or more devices to a tracker resource (up
// to 10 devices per batch). Amazon Location uses the data when it reports the last
// known device position and position history. Amazon Location retains location
// data for 30 days.
//
// Position updates are handled based on the PositionFiltering property of the
// tracker. When PositionFiltering is set to TimeBased , updates are evaluated
// against linked geofence collections, and location data is stored at a maximum of
// one position per 30 second interval. If your update frequency is more often than
// every 30 seconds, only one update per 30 seconds is stored for each unique
// device ID.
//
// When PositionFiltering is set to DistanceBased filtering, location data is
// stored and evaluated against linked geofence collections only if the device has
// moved more than 30 m (98.4 ft).
//
// When PositionFiltering is set to AccuracyBased filtering, location data is
// stored and evaluated against linked geofence collections only if the device has
// moved more than the measured accuracy. For example, if two consecutive updates
// from a device have a horizontal accuracy of 5 m and 10 m, the second update is
// neither stored or evaluated if the device has moved less than 15 m. If
// PositionFiltering is set to AccuracyBased filtering, Amazon Location uses the
// default value { "Horizontal": 0} when accuracy is not provided on a
// DevicePositionUpdate .
func (c *Client) BatchUpdateDevicePosition(ctx context.Context, params *BatchUpdateDevicePositionInput, optFns ...func(*Options)) (*BatchUpdateDevicePositionOutput, error) {
	if params == nil {
		params = &BatchUpdateDevicePositionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchUpdateDevicePosition", params, optFns, c.addOperationBatchUpdateDevicePositionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchUpdateDevicePositionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUpdateDevicePositionInput struct {

	// The name of the tracker resource to update.
	//
	// This member is required.
	TrackerName *string

	// Contains the position update details for each device, up to 10 devices.
	//
	// This member is required.
	Updates []types.DevicePositionUpdate

	noSmithyDocumentSerde
}

type BatchUpdateDevicePositionOutput struct {

	// Contains error details for each device that failed to update its position.
	//
	// This member is required.
	Errors []types.BatchUpdateDevicePositionError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchUpdateDevicePositionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchUpdateDevicePosition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchUpdateDevicePosition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchUpdateDevicePosition"); err != nil {
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
	if err = addEndpointPrefix_opBatchUpdateDevicePositionMiddleware(stack); err != nil {
		return err
	}
	if err = addOpBatchUpdateDevicePositionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUpdateDevicePosition(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opBatchUpdateDevicePositionMiddleware struct {
}

func (*endpointPrefix_opBatchUpdateDevicePositionMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opBatchUpdateDevicePositionMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "tracking." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opBatchUpdateDevicePositionMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opBatchUpdateDevicePositionMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opBatchUpdateDevicePosition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchUpdateDevicePosition",
	}
}
