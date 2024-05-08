// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a vector data tile from the map resource. Map tiles are used by
// clients to render a map. they're addressed using a grid arrangement with an X
// coordinate, Y coordinate, and Z (zoom) level.
//
// The origin (0, 0) is the top left of the map. Increasing the zoom level by 1
// doubles both the X and Y dimensions, so a tile containing data for the entire
// world at (0/0/0) will be split into 4 tiles at zoom 1 (1/0/0, 1/0/1, 1/1/0,
// 1/1/1).
func (c *Client) GetMapTile(ctx context.Context, params *GetMapTileInput, optFns ...func(*Options)) (*GetMapTileOutput, error) {
	if params == nil {
		params = &GetMapTileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMapTile", params, optFns, c.addOperationGetMapTileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMapTileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMapTileInput struct {

	// The map resource to retrieve the map tiles from.
	//
	// This member is required.
	MapName *string

	// The X axis value for the map tile.
	//
	// This member is required.
	X *string

	// The Y axis value for the map tile.
	//
	// This member is required.
	Y *string

	// The zoom value for the map tile.
	//
	// This member is required.
	Z *string

	// The optional [API key] to authorize the request.
	//
	// [API key]: https://docs.aws.amazon.com/location/latest/developerguide/using-apikeys.html
	Key *string

	noSmithyDocumentSerde
}

type GetMapTileOutput struct {

	// Contains Mapbox Vector Tile (MVT) data.
	Blob []byte

	// The HTTP Cache-Control directive for the value.
	CacheControl *string

	// The map tile's content type. For example, application/vnd.mapbox-vector-tile .
	ContentType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMapTileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMapTile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMapTile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMapTile"); err != nil {
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
	if err = addEndpointPrefix_opGetMapTileMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetMapTileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMapTile(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetMapTileMiddleware struct {
}

func (*endpointPrefix_opGetMapTileMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetMapTileMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "maps." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetMapTileMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetMapTileMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetMapTile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMapTile",
	}
}
