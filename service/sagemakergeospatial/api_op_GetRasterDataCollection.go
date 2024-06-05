// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakergeospatial

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemakergeospatial/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this operation to get details of a specific raster data collection.
func (c *Client) GetRasterDataCollection(ctx context.Context, params *GetRasterDataCollectionInput, optFns ...func(*Options)) (*GetRasterDataCollectionOutput, error) {
	if params == nil {
		params = &GetRasterDataCollectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRasterDataCollection", params, optFns, c.addOperationGetRasterDataCollectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRasterDataCollectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRasterDataCollectionInput struct {

	// The Amazon Resource Name (ARN) of the raster data collection.
	//
	// This member is required.
	Arn *string

	noSmithyDocumentSerde
}

type GetRasterDataCollectionOutput struct {

	// The Amazon Resource Name (ARN) of the raster data collection.
	//
	// This member is required.
	Arn *string

	// A description of the raster data collection.
	//
	// This member is required.
	Description *string

	// The URL of the description page.
	//
	// This member is required.
	DescriptionPageUrl *string

	// The list of image source bands in the raster data collection.
	//
	// This member is required.
	ImageSourceBands []string

	// The name of the raster data collection.
	//
	// This member is required.
	Name *string

	// The filters supported by the raster data collection.
	//
	// This member is required.
	SupportedFilters []types.Filter

	// The raster data collection type.
	//
	// This member is required.
	Type types.DataCollectionType

	// Each tag consists of a key and a value.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRasterDataCollectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRasterDataCollection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRasterDataCollection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetRasterDataCollection"); err != nil {
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
	if err = addOpGetRasterDataCollectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRasterDataCollection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRasterDataCollection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetRasterDataCollection",
	}
}
