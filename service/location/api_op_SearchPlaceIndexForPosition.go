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

// Reverse geocodes a given coordinate and returns a legible address. Allows you
// to search for Places or points of interest near a given position.
func (c *Client) SearchPlaceIndexForPosition(ctx context.Context, params *SearchPlaceIndexForPositionInput, optFns ...func(*Options)) (*SearchPlaceIndexForPositionOutput, error) {
	if params == nil {
		params = &SearchPlaceIndexForPositionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchPlaceIndexForPosition", params, optFns, c.addOperationSearchPlaceIndexForPositionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchPlaceIndexForPositionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchPlaceIndexForPositionInput struct {

	// The name of the place index resource you want to use for the search.
	//
	// This member is required.
	IndexName *string

	// Specifies the longitude and latitude of the position to query.
	//
	// This parameter must contain a pair of numbers. The first number represents the
	// X coordinate, or longitude; the second number represents the Y coordinate, or
	// latitude.
	//
	// For example, [-123.1174, 49.2847] represents a position with longitude -123.1174
	// and latitude 49.2847 .
	//
	// This member is required.
	Position []float64

	// The optional [API key] to authorize the request.
	//
	// [API key]: https://docs.aws.amazon.com/location/latest/developerguide/using-apikeys.html
	Key *string

	// The preferred language used to return results. The value must be a valid [BCP 47]
	// language tag, for example, en for English.
	//
	// This setting affects the languages used in the results, but not the results
	// themselves. If no language is specified, or not supported for a particular
	// result, the partner automatically chooses a language for the result.
	//
	// For an example, we'll use the Greek language. You search for a location around
	// Athens, Greece, with the language parameter set to en . The city in the results
	// will most likely be returned as Athens .
	//
	// If you set the language parameter to el , for Greek, then the city in the
	// results will more likely be returned as Αθήνα .
	//
	// If the data provider does not have a value for Greek, the result will be in a
	// language that the provider does support.
	//
	// [BCP 47]: https://tools.ietf.org/search/bcp47
	Language *string

	// An optional parameter. The maximum number of results returned per request.
	//
	// Default value: 50
	MaxResults *int32

	noSmithyDocumentSerde
}

type SearchPlaceIndexForPositionOutput struct {

	// Returns a list of Places closest to the specified position. Each result
	// contains additional information about the Places returned.
	//
	// This member is required.
	Results []types.SearchForPositionResult

	// Contains a summary of the request. Echoes the input values for Position ,
	// Language , MaxResults , and the DataSource of the place index.
	//
	// This member is required.
	Summary *types.SearchPlaceIndexForPositionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchPlaceIndexForPositionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchPlaceIndexForPosition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchPlaceIndexForPosition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchPlaceIndexForPosition"); err != nil {
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
	if err = addEndpointPrefix_opSearchPlaceIndexForPositionMiddleware(stack); err != nil {
		return err
	}
	if err = addOpSearchPlaceIndexForPositionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchPlaceIndexForPosition(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opSearchPlaceIndexForPositionMiddleware struct {
}

func (*endpointPrefix_opSearchPlaceIndexForPositionMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opSearchPlaceIndexForPositionMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "places." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opSearchPlaceIndexForPositionMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opSearchPlaceIndexForPositionMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opSearchPlaceIndexForPosition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchPlaceIndexForPosition",
	}
}
