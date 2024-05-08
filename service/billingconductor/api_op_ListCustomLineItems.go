// Code generated by smithy-go-codegen DO NOT EDIT.

package billingconductor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/billingconductor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	A paginated call to get a list of all custom line items (FFLIs) for the given
//
// billing period. If you don't provide a billing period, the current billing
// period is used.
func (c *Client) ListCustomLineItems(ctx context.Context, params *ListCustomLineItemsInput, optFns ...func(*Options)) (*ListCustomLineItemsOutput, error) {
	if params == nil {
		params = &ListCustomLineItemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCustomLineItems", params, optFns, c.addOperationListCustomLineItemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCustomLineItemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCustomLineItemsInput struct {

	//  The preferred billing period to get custom line items (FFLIs).
	BillingPeriod *string

	// A ListCustomLineItemsFilter that specifies the custom line item names and/or
	// billing group Amazon Resource Names (ARNs) to retrieve FFLI information.
	Filters *types.ListCustomLineItemsFilter

	//  The maximum number of billing groups to retrieve.
	MaxResults *int32

	//  The pagination token that's used on subsequent calls to get custom line items
	// (FFLIs).
	NextToken *string

	noSmithyDocumentSerde
}

type ListCustomLineItemsOutput struct {

	//  A list of FreeFormLineItemListElements received.
	CustomLineItems []types.CustomLineItemListElement

	//  The pagination token that's used on subsequent calls to get custom line items
	// (FFLIs).
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCustomLineItemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCustomLineItems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCustomLineItems{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCustomLineItems"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCustomLineItems(options.Region), middleware.Before); err != nil {
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

// ListCustomLineItemsAPIClient is a client that implements the
// ListCustomLineItems operation.
type ListCustomLineItemsAPIClient interface {
	ListCustomLineItems(context.Context, *ListCustomLineItemsInput, ...func(*Options)) (*ListCustomLineItemsOutput, error)
}

var _ ListCustomLineItemsAPIClient = (*Client)(nil)

// ListCustomLineItemsPaginatorOptions is the paginator options for
// ListCustomLineItems
type ListCustomLineItemsPaginatorOptions struct {
	//  The maximum number of billing groups to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCustomLineItemsPaginator is a paginator for ListCustomLineItems
type ListCustomLineItemsPaginator struct {
	options   ListCustomLineItemsPaginatorOptions
	client    ListCustomLineItemsAPIClient
	params    *ListCustomLineItemsInput
	nextToken *string
	firstPage bool
}

// NewListCustomLineItemsPaginator returns a new ListCustomLineItemsPaginator
func NewListCustomLineItemsPaginator(client ListCustomLineItemsAPIClient, params *ListCustomLineItemsInput, optFns ...func(*ListCustomLineItemsPaginatorOptions)) *ListCustomLineItemsPaginator {
	if params == nil {
		params = &ListCustomLineItemsInput{}
	}

	options := ListCustomLineItemsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCustomLineItemsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCustomLineItemsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCustomLineItems page.
func (p *ListCustomLineItemsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCustomLineItemsOutput, error) {
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

	result, err := p.client.ListCustomLineItems(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCustomLineItems(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCustomLineItems",
	}
}
