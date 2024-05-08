// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Retrieves a list of summaries of created vehicles.
//
// This API operation uses pagination. Specify the nextToken parameter in the
// request to return more results.
func (c *Client) ListVehicles(ctx context.Context, params *ListVehiclesInput, optFns ...func(*Options)) (*ListVehiclesOutput, error) {
	if params == nil {
		params = &ListVehiclesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVehicles", params, optFns, c.addOperationListVehiclesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVehiclesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVehiclesInput struct {

	//  The maximum number of items to return, between 1 and 100, inclusive.
	MaxResults *int32

	//  The Amazon Resource Name (ARN) of a vehicle model (model manifest). You can
	// use this optional parameter to list only the vehicles created from a certain
	// vehicle model.
	ModelManifestArn *string

	// A pagination token for the next set of results.
	//
	// If the results of a search are large, only a portion of the results are
	// returned, and a nextToken pagination token is returned in the response. To
	// retrieve the next set of results, reissue the search request and include the
	// returned token. When all results have been returned, the response does not
	// contain a pagination token value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListVehiclesOutput struct {

	//  The token to retrieve the next set of results, or null if there are no more
	// results.
	NextToken *string

	//  A list of vehicles and information about them.
	VehicleSummaries []types.VehicleSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVehiclesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListVehicles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListVehicles{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListVehicles"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVehicles(options.Region), middleware.Before); err != nil {
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

// ListVehiclesAPIClient is a client that implements the ListVehicles operation.
type ListVehiclesAPIClient interface {
	ListVehicles(context.Context, *ListVehiclesInput, ...func(*Options)) (*ListVehiclesOutput, error)
}

var _ ListVehiclesAPIClient = (*Client)(nil)

// ListVehiclesPaginatorOptions is the paginator options for ListVehicles
type ListVehiclesPaginatorOptions struct {
	//  The maximum number of items to return, between 1 and 100, inclusive.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVehiclesPaginator is a paginator for ListVehicles
type ListVehiclesPaginator struct {
	options   ListVehiclesPaginatorOptions
	client    ListVehiclesAPIClient
	params    *ListVehiclesInput
	nextToken *string
	firstPage bool
}

// NewListVehiclesPaginator returns a new ListVehiclesPaginator
func NewListVehiclesPaginator(client ListVehiclesAPIClient, params *ListVehiclesInput, optFns ...func(*ListVehiclesPaginatorOptions)) *ListVehiclesPaginator {
	if params == nil {
		params = &ListVehiclesInput{}
	}

	options := ListVehiclesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListVehiclesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVehiclesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListVehicles page.
func (p *ListVehiclesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVehiclesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListVehicles(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListVehicles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListVehicles",
	}
}
