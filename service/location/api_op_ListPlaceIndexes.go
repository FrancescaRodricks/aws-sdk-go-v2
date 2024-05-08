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

// Lists place index resources in your Amazon Web Services account.
func (c *Client) ListPlaceIndexes(ctx context.Context, params *ListPlaceIndexesInput, optFns ...func(*Options)) (*ListPlaceIndexesOutput, error) {
	if params == nil {
		params = &ListPlaceIndexesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPlaceIndexes", params, optFns, c.addOperationListPlaceIndexesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPlaceIndexesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPlaceIndexesInput struct {

	// An optional limit for the maximum number of results returned in a single call.
	//
	// Default value: 100
	MaxResults *int32

	// The pagination token specifying which page of results to return in the
	// response. If no token is provided, the default page is the first page.
	//
	// Default value: null
	NextToken *string

	noSmithyDocumentSerde
}

type ListPlaceIndexesOutput struct {

	// Lists the place index resources that exist in your Amazon Web Services account
	//
	// This member is required.
	Entries []types.ListPlaceIndexesResponseEntry

	// A pagination token indicating that there are additional pages available. You
	// can use the token in a new request to fetch the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPlaceIndexesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPlaceIndexes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPlaceIndexes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPlaceIndexes"); err != nil {
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
	if err = addEndpointPrefix_opListPlaceIndexesMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPlaceIndexes(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListPlaceIndexesMiddleware struct {
}

func (*endpointPrefix_opListPlaceIndexesMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListPlaceIndexesMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "cp.places." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListPlaceIndexesMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListPlaceIndexesMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListPlaceIndexesAPIClient is a client that implements the ListPlaceIndexes
// operation.
type ListPlaceIndexesAPIClient interface {
	ListPlaceIndexes(context.Context, *ListPlaceIndexesInput, ...func(*Options)) (*ListPlaceIndexesOutput, error)
}

var _ ListPlaceIndexesAPIClient = (*Client)(nil)

// ListPlaceIndexesPaginatorOptions is the paginator options for ListPlaceIndexes
type ListPlaceIndexesPaginatorOptions struct {
	// An optional limit for the maximum number of results returned in a single call.
	//
	// Default value: 100
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPlaceIndexesPaginator is a paginator for ListPlaceIndexes
type ListPlaceIndexesPaginator struct {
	options   ListPlaceIndexesPaginatorOptions
	client    ListPlaceIndexesAPIClient
	params    *ListPlaceIndexesInput
	nextToken *string
	firstPage bool
}

// NewListPlaceIndexesPaginator returns a new ListPlaceIndexesPaginator
func NewListPlaceIndexesPaginator(client ListPlaceIndexesAPIClient, params *ListPlaceIndexesInput, optFns ...func(*ListPlaceIndexesPaginatorOptions)) *ListPlaceIndexesPaginator {
	if params == nil {
		params = &ListPlaceIndexesInput{}
	}

	options := ListPlaceIndexesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPlaceIndexesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPlaceIndexesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPlaceIndexes page.
func (p *ListPlaceIndexesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPlaceIndexesOutput, error) {
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

	result, err := p.client.ListPlaceIndexes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPlaceIndexes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPlaceIndexes",
	}
}
