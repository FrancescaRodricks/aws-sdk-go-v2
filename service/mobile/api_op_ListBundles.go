// Code generated by smithy-go-codegen DO NOT EDIT.

package mobile

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mobile/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List all available bundles.
func (c *Client) ListBundles(ctx context.Context, params *ListBundlesInput, optFns ...func(*Options)) (*ListBundlesOutput, error) {
	if params == nil {
		params = &ListBundlesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBundles", params, optFns, c.addOperationListBundlesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBundlesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request structure to request all available bundles.
type ListBundlesInput struct {

	// Maximum number of records to list in a single response.
	MaxResults int32

	// Pagination token. Set to null to start listing bundles from start. If non-null
	// pagination token is returned in a result, then pass its value in here in another
	// request to list more bundles.
	NextToken *string

	noSmithyDocumentSerde
}

// Result structure contains a list of all available bundles with details.
type ListBundlesOutput struct {

	// A list of bundles.
	BundleList []types.BundleDetails

	// Pagination token. If non-null pagination token is returned in a result, then
	// pass its value in another request to fetch more entries.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBundlesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBundles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBundles{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBundles(options.Region), middleware.Before); err != nil {
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
	return nil
}

// ListBundlesAPIClient is a client that implements the ListBundles operation.
type ListBundlesAPIClient interface {
	ListBundles(context.Context, *ListBundlesInput, ...func(*Options)) (*ListBundlesOutput, error)
}

var _ ListBundlesAPIClient = (*Client)(nil)

// ListBundlesPaginatorOptions is the paginator options for ListBundles
type ListBundlesPaginatorOptions struct {
	// Maximum number of records to list in a single response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBundlesPaginator is a paginator for ListBundles
type ListBundlesPaginator struct {
	options   ListBundlesPaginatorOptions
	client    ListBundlesAPIClient
	params    *ListBundlesInput
	nextToken *string
	firstPage bool
}

// NewListBundlesPaginator returns a new ListBundlesPaginator
func NewListBundlesPaginator(client ListBundlesAPIClient, params *ListBundlesInput, optFns ...func(*ListBundlesPaginatorOptions)) *ListBundlesPaginator {
	if params == nil {
		params = &ListBundlesInput{}
	}

	options := ListBundlesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBundlesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBundlesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBundles page.
func (p *ListBundlesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBundlesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListBundles(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBundles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "awsmobilehubservice",
		OperationName: "ListBundles",
	}
}