// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of IDs for all fleets that the vehicle is associated with.
//
// This API operation uses pagination. Specify the nextToken parameter in the
// request to return more results.
func (c *Client) ListFleetsForVehicle(ctx context.Context, params *ListFleetsForVehicleInput, optFns ...func(*Options)) (*ListFleetsForVehicleOutput, error) {
	if params == nil {
		params = &ListFleetsForVehicleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFleetsForVehicle", params, optFns, c.addOperationListFleetsForVehicleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFleetsForVehicleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFleetsForVehicleInput struct {

	//  The ID of the vehicle to retrieve information about.
	//
	// This member is required.
	VehicleName *string

	//  The maximum number of items to return, between 1 and 100, inclusive.
	MaxResults *int32

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

type ListFleetsForVehicleOutput struct {

	//  A list of fleet IDs that the vehicle is associated with.
	Fleets []string

	//  The token to retrieve the next set of results, or null if there are no more
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFleetsForVehicleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListFleetsForVehicle{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListFleetsForVehicle{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFleetsForVehicle"); err != nil {
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
	if err = addOpListFleetsForVehicleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFleetsForVehicle(options.Region), middleware.Before); err != nil {
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

// ListFleetsForVehicleAPIClient is a client that implements the
// ListFleetsForVehicle operation.
type ListFleetsForVehicleAPIClient interface {
	ListFleetsForVehicle(context.Context, *ListFleetsForVehicleInput, ...func(*Options)) (*ListFleetsForVehicleOutput, error)
}

var _ ListFleetsForVehicleAPIClient = (*Client)(nil)

// ListFleetsForVehiclePaginatorOptions is the paginator options for
// ListFleetsForVehicle
type ListFleetsForVehiclePaginatorOptions struct {
	//  The maximum number of items to return, between 1 and 100, inclusive.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFleetsForVehiclePaginator is a paginator for ListFleetsForVehicle
type ListFleetsForVehiclePaginator struct {
	options   ListFleetsForVehiclePaginatorOptions
	client    ListFleetsForVehicleAPIClient
	params    *ListFleetsForVehicleInput
	nextToken *string
	firstPage bool
}

// NewListFleetsForVehiclePaginator returns a new ListFleetsForVehiclePaginator
func NewListFleetsForVehiclePaginator(client ListFleetsForVehicleAPIClient, params *ListFleetsForVehicleInput, optFns ...func(*ListFleetsForVehiclePaginatorOptions)) *ListFleetsForVehiclePaginator {
	if params == nil {
		params = &ListFleetsForVehicleInput{}
	}

	options := ListFleetsForVehiclePaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFleetsForVehiclePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFleetsForVehiclePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFleetsForVehicle page.
func (p *ListFleetsForVehiclePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFleetsForVehicleOutput, error) {
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

	result, err := p.client.ListFleetsForVehicle(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFleetsForVehicle(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFleetsForVehicle",
	}
}
