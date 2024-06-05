// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the imports for a bot, bot locale, or custom vocabulary. Imports are kept
// in the list for 7 days.
func (c *Client) ListImports(ctx context.Context, params *ListImportsInput, optFns ...func(*Options)) (*ListImportsOutput, error) {
	if params == nil {
		params = &ListImportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListImports", params, optFns, c.addOperationListImportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListImportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListImportsInput struct {

	// The unique identifier that Amazon Lex assigned to the bot.
	BotId *string

	// The version of the bot to list imports for.
	BotVersion *string

	// Provides the specification of a filter used to limit the bots in the response
	// to only those that match the filter specification. You can only specify one
	// filter and one string to filter on.
	Filters []types.ImportFilter

	// Specifies the locale that should be present in the list. If you don't specify a
	// resource type in the filters parameter, the list contains both bot locales and
	// custom vocabularies.
	LocaleId *string

	// The maximum number of imports to return in each page of results. If there are
	// fewer results than the max page size, only the actual number of results are
	// returned.
	MaxResults *int32

	// If the response from the ListImports operation contains more results than
	// specified in the maxResults parameter, a token is returned in the response.
	//
	// Use the returned token in the nextToken parameter of a ListImports request to
	// return the next page of results. For a complete set of results, call the
	// ListImports operation until the nextToken returned in the response is null.
	NextToken *string

	// Determines the field that the list of imports is sorted by. You can sort by the
	// LastUpdatedDateTime field in ascending or descending order.
	SortBy *types.ImportSortBy

	noSmithyDocumentSerde
}

type ListImportsOutput struct {

	// The unique identifier assigned by Amazon Lex to the bot.
	BotId *string

	// The version of the bot that was imported. It will always be DRAFT .
	BotVersion *string

	// Summary information for the imports that meet the filter criteria specified in
	// the request. The length of the list is specified in the maxResults parameter.
	// If there are more imports available, the nextToken field contains a token to
	// get the next page of results.
	ImportSummaries []types.ImportSummary

	// The locale specified in the request.
	LocaleId *string

	// A token that indicates whether there are more results to return in a response
	// to the ListImports operation. If the nextToken field is present, you send the
	// contents as the nextToken parameter of a ListImports operation request to get
	// the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListImportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListImports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListImports{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListImports"); err != nil {
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
	if err = addOpListImportsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListImports(options.Region), middleware.Before); err != nil {
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

// ListImportsAPIClient is a client that implements the ListImports operation.
type ListImportsAPIClient interface {
	ListImports(context.Context, *ListImportsInput, ...func(*Options)) (*ListImportsOutput, error)
}

var _ ListImportsAPIClient = (*Client)(nil)

// ListImportsPaginatorOptions is the paginator options for ListImports
type ListImportsPaginatorOptions struct {
	// The maximum number of imports to return in each page of results. If there are
	// fewer results than the max page size, only the actual number of results are
	// returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListImportsPaginator is a paginator for ListImports
type ListImportsPaginator struct {
	options   ListImportsPaginatorOptions
	client    ListImportsAPIClient
	params    *ListImportsInput
	nextToken *string
	firstPage bool
}

// NewListImportsPaginator returns a new ListImportsPaginator
func NewListImportsPaginator(client ListImportsAPIClient, params *ListImportsInput, optFns ...func(*ListImportsPaginatorOptions)) *ListImportsPaginator {
	if params == nil {
		params = &ListImportsInput{}
	}

	options := ListImportsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListImportsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListImportsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListImports page.
func (p *ListImportsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListImportsOutput, error) {
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

	result, err := p.client.ListImports(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListImports(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListImports",
	}
}
