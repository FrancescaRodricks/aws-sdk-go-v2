// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists configured tables.
func (c *Client) ListConfiguredTables(ctx context.Context, params *ListConfiguredTablesInput, optFns ...func(*Options)) (*ListConfiguredTablesOutput, error) {
	if params == nil {
		params = &ListConfiguredTablesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConfiguredTables", params, optFns, c.addOperationListConfiguredTablesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConfiguredTablesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConfiguredTablesInput struct {

	// The maximum size of the results that is returned per call.
	MaxResults *int32

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListConfiguredTablesOutput struct {

	// The configured tables listed by the request.
	//
	// This member is required.
	ConfiguredTableSummaries []types.ConfiguredTableSummary

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListConfiguredTablesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListConfiguredTables{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListConfiguredTables{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListConfiguredTables"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConfiguredTables(options.Region), middleware.Before); err != nil {
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

// ListConfiguredTablesAPIClient is a client that implements the
// ListConfiguredTables operation.
type ListConfiguredTablesAPIClient interface {
	ListConfiguredTables(context.Context, *ListConfiguredTablesInput, ...func(*Options)) (*ListConfiguredTablesOutput, error)
}

var _ ListConfiguredTablesAPIClient = (*Client)(nil)

// ListConfiguredTablesPaginatorOptions is the paginator options for
// ListConfiguredTables
type ListConfiguredTablesPaginatorOptions struct {
	// The maximum size of the results that is returned per call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListConfiguredTablesPaginator is a paginator for ListConfiguredTables
type ListConfiguredTablesPaginator struct {
	options   ListConfiguredTablesPaginatorOptions
	client    ListConfiguredTablesAPIClient
	params    *ListConfiguredTablesInput
	nextToken *string
	firstPage bool
}

// NewListConfiguredTablesPaginator returns a new ListConfiguredTablesPaginator
func NewListConfiguredTablesPaginator(client ListConfiguredTablesAPIClient, params *ListConfiguredTablesInput, optFns ...func(*ListConfiguredTablesPaginatorOptions)) *ListConfiguredTablesPaginator {
	if params == nil {
		params = &ListConfiguredTablesInput{}
	}

	options := ListConfiguredTablesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListConfiguredTablesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListConfiguredTablesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListConfiguredTables page.
func (p *ListConfiguredTablesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListConfiguredTablesOutput, error) {
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

	result, err := p.client.ListConfiguredTables(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListConfiguredTables(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListConfiguredTables",
	}
}
