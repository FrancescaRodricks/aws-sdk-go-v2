// Code generated by smithy-go-codegen DO NOT EDIT.

package amp

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The ListScrapers operation lists all of the scrapers in your account. This
// includes scrapers being created or deleted. You can optionally filter the
// returned list.
func (c *Client) ListScrapers(ctx context.Context, params *ListScrapersInput, optFns ...func(*Options)) (*ListScrapersOutput, error) {
	if params == nil {
		params = &ListScrapersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListScrapers", params, optFns, c.addOperationListScrapersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListScrapersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ListScrapers operation.
type ListScrapersInput struct {

	// (Optional) A list of key-value pairs to filter the list of scrapers returned.
	// Keys include status , sourceArn , destinationArn , and alias .
	//
	// Filters on the same key are OR 'd together, and filters on different keys are
	// AND 'd together. For example, status=ACTIVE&status=CREATING&alias=Test , will
	// return all scrapers that have the alias Test, and are either in status ACTIVE or
	// CREATING.
	//
	// To find all active scrapers that are sending metrics to a specific Amazon
	// Managed Service for Prometheus workspace, you would use the ARN of the workspace
	// in a query:
	//
	//     status=ACTIVE&destinationArn=arn:aws:aps:us-east-1:123456789012:workspace/ws-example1-1234-abcd-56ef-123456789012
	//
	// If this is included, it filters the results to only the scrapers that match the
	// filter.
	Filters map[string][]string

	// Optional) The maximum number of scrapers to return in one ListScrapers
	// operation. The range is 1-1000.
	//
	// If you omit this parameter, the default of 100 is used.
	MaxResults *int32

	// (Optional) The token for the next set of items to return. (You received this
	// token from a previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

// Represents the output of a ListScrapers operation.
type ListScrapersOutput struct {

	// A list of ScraperSummary structures giving information about scrapers in the
	// account that match the filters provided.
	//
	// This member is required.
	Scrapers []types.ScraperSummary

	// A token indicating that there are more results to retrieve. You can use this
	// token as part of your next ListScrapers operation to retrieve those results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListScrapersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListScrapers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListScrapers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListScrapers"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListScrapers(options.Region), middleware.Before); err != nil {
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

// ListScrapersAPIClient is a client that implements the ListScrapers operation.
type ListScrapersAPIClient interface {
	ListScrapers(context.Context, *ListScrapersInput, ...func(*Options)) (*ListScrapersOutput, error)
}

var _ ListScrapersAPIClient = (*Client)(nil)

// ListScrapersPaginatorOptions is the paginator options for ListScrapers
type ListScrapersPaginatorOptions struct {
	// Optional) The maximum number of scrapers to return in one ListScrapers
	// operation. The range is 1-1000.
	//
	// If you omit this parameter, the default of 100 is used.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListScrapersPaginator is a paginator for ListScrapers
type ListScrapersPaginator struct {
	options   ListScrapersPaginatorOptions
	client    ListScrapersAPIClient
	params    *ListScrapersInput
	nextToken *string
	firstPage bool
}

// NewListScrapersPaginator returns a new ListScrapersPaginator
func NewListScrapersPaginator(client ListScrapersAPIClient, params *ListScrapersInput, optFns ...func(*ListScrapersPaginatorOptions)) *ListScrapersPaginator {
	if params == nil {
		params = &ListScrapersInput{}
	}

	options := ListScrapersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListScrapersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListScrapersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListScrapers page.
func (p *ListScrapersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListScrapersOutput, error) {
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

	result, err := p.client.ListScrapers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListScrapers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListScrapers",
	}
}
