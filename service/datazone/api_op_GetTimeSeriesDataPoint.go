// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the existing data point for the asset.
func (c *Client) GetTimeSeriesDataPoint(ctx context.Context, params *GetTimeSeriesDataPointInput, optFns ...func(*Options)) (*GetTimeSeriesDataPointOutput, error) {
	if params == nil {
		params = &GetTimeSeriesDataPointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTimeSeriesDataPoint", params, optFns, c.addOperationGetTimeSeriesDataPointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTimeSeriesDataPointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTimeSeriesDataPointInput struct {

	// The ID of the Amazon DataZone domain that houses the asset for which you want
	// to get the data point.
	//
	// This member is required.
	DomainIdentifier *string

	// The ID of the asset for which you want to get the data point.
	//
	// This member is required.
	EntityIdentifier *string

	// The type of the asset for which you want to get the data point.
	//
	// This member is required.
	EntityType types.TimeSeriesEntityType

	// The name of the time series form that houses the data point that you want to
	// get.
	//
	// This member is required.
	FormName *string

	// The ID of the data point that you want to get.
	//
	// This member is required.
	Identifier *string

	noSmithyDocumentSerde
}

type GetTimeSeriesDataPointOutput struct {

	// The ID of the Amazon DataZone domain that houses the asset data point that you
	// want to get.
	DomainId *string

	// The ID of the asset for which you want to get the data point.
	EntityId *string

	// The type of the asset for which you want to get the data point.
	EntityType types.TimeSeriesEntityType

	// The time series form that houses the data point that you want to get.
	Form *types.TimeSeriesDataPointFormOutput

	// The name of the time series form that houses the data point that you want to
	// get.
	FormName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTimeSeriesDataPointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTimeSeriesDataPoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTimeSeriesDataPoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetTimeSeriesDataPoint"); err != nil {
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
	if err = addOpGetTimeSeriesDataPointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTimeSeriesDataPoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTimeSeriesDataPoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetTimeSeriesDataPoint",
	}
}
