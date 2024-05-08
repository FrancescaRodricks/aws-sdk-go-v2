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

// Paginated list of custom vocabulary items for a given bot locale's custom
// vocabulary.
func (c *Client) ListCustomVocabularyItems(ctx context.Context, params *ListCustomVocabularyItemsInput, optFns ...func(*Options)) (*ListCustomVocabularyItemsOutput, error) {
	if params == nil {
		params = &ListCustomVocabularyItemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCustomVocabularyItems", params, optFns, c.addOperationListCustomVocabularyItemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCustomVocabularyItemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCustomVocabularyItemsInput struct {

	// The identifier of the version of the bot associated with this custom vocabulary.
	//
	// This member is required.
	BotId *string

	// The bot version of the bot to the list custom vocabulary request.
	//
	// This member is required.
	BotVersion *string

	// The identifier of the language and locale where this custom vocabulary is used.
	// The string must match one of the supported locales. For more information, see
	// Supported languages
	// (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html).
	//
	// This member is required.
	LocaleId *string

	// The maximum number of items returned by the list operation.
	MaxResults *int32

	// The nextToken identifier to the list custom vocabulary request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCustomVocabularyItemsOutput struct {

	// The identifier of the bot associated with this custom vocabulary.
	BotId *string

	// The identifier of the version of the bot associated with this custom vocabulary.
	BotVersion *string

	// The custom vocabulary items from the list custom vocabulary response.
	CustomVocabularyItems []types.CustomVocabularyItem

	// The identifier of the language and locale where this custom vocabulary is used.
	// The string must match one of the supported locales. For more information, see [Supported Languages].
	//
	// [Supported Languages]: https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html
	LocaleId *string

	// The nextToken identifier to the list custom vocabulary response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCustomVocabularyItemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCustomVocabularyItems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCustomVocabularyItems{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCustomVocabularyItems"); err != nil {
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
	if err = addOpListCustomVocabularyItemsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCustomVocabularyItems(options.Region), middleware.Before); err != nil {
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

// ListCustomVocabularyItemsAPIClient is a client that implements the
// ListCustomVocabularyItems operation.
type ListCustomVocabularyItemsAPIClient interface {
	ListCustomVocabularyItems(context.Context, *ListCustomVocabularyItemsInput, ...func(*Options)) (*ListCustomVocabularyItemsOutput, error)
}

var _ ListCustomVocabularyItemsAPIClient = (*Client)(nil)

// ListCustomVocabularyItemsPaginatorOptions is the paginator options for
// ListCustomVocabularyItems
type ListCustomVocabularyItemsPaginatorOptions struct {
	// The maximum number of items returned by the list operation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCustomVocabularyItemsPaginator is a paginator for ListCustomVocabularyItems
type ListCustomVocabularyItemsPaginator struct {
	options   ListCustomVocabularyItemsPaginatorOptions
	client    ListCustomVocabularyItemsAPIClient
	params    *ListCustomVocabularyItemsInput
	nextToken *string
	firstPage bool
}

// NewListCustomVocabularyItemsPaginator returns a new
// ListCustomVocabularyItemsPaginator
func NewListCustomVocabularyItemsPaginator(client ListCustomVocabularyItemsAPIClient, params *ListCustomVocabularyItemsInput, optFns ...func(*ListCustomVocabularyItemsPaginatorOptions)) *ListCustomVocabularyItemsPaginator {
	if params == nil {
		params = &ListCustomVocabularyItemsInput{}
	}

	options := ListCustomVocabularyItemsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCustomVocabularyItemsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCustomVocabularyItemsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCustomVocabularyItems page.
func (p *ListCustomVocabularyItemsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCustomVocabularyItemsOutput, error) {
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

	result, err := p.client.ListCustomVocabularyItems(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCustomVocabularyItems(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCustomVocabularyItems",
	}
}
