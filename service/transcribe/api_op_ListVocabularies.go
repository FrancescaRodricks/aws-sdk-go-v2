// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a list of custom vocabularies that match the specified criteria. If no
// criteria are specified, all custom vocabularies are returned.
//
// To get detailed information about a specific custom vocabulary, use the
// operation.
func (c *Client) ListVocabularies(ctx context.Context, params *ListVocabulariesInput, optFns ...func(*Options)) (*ListVocabulariesOutput, error) {
	if params == nil {
		params = &ListVocabulariesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVocabularies", params, optFns, c.addOperationListVocabulariesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVocabulariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVocabulariesInput struct {

	// The maximum number of custom vocabularies to return in each page of results. If
	// there are fewer results than the value that you specify, only the actual results
	// are returned. If you do not specify a value, a default of 5 is used.
	MaxResults *int32

	// Returns only the custom vocabularies that contain the specified string. The
	// search is not case sensitive.
	NameContains *string

	// If your ListVocabularies request returns more results than can be displayed,
	// NextToken is displayed in the response with an associated string. To get the
	// next page of results, copy this string and repeat your request, including
	// NextToken with the value of the copied string. Repeat as needed to view all your
	// results.
	NextToken *string

	// Returns only custom vocabularies with the specified state. Vocabularies are
	// ordered by creation date, with the newest vocabulary first. If you do not
	// include StateEquals , all custom medical vocabularies are returned.
	StateEquals types.VocabularyState

	noSmithyDocumentSerde
}

type ListVocabulariesOutput struct {

	// If NextToken is present in your response, it indicates that not all results are
	// displayed. To view the next set of results, copy the string associated with the
	// NextToken parameter in your results output, then run your request again
	// including NextToken with the value of the copied string. Repeat as needed to
	// view all your results.
	NextToken *string

	// Lists all custom vocabularies that have the status specified in your request.
	// Vocabularies are ordered by creation date, with the newest vocabulary first.
	Status types.VocabularyState

	// Provides information about the custom vocabularies that match the criteria
	// specified in your request.
	Vocabularies []types.VocabularyInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVocabulariesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListVocabularies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListVocabularies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListVocabularies"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVocabularies(options.Region), middleware.Before); err != nil {
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

// ListVocabulariesAPIClient is a client that implements the ListVocabularies
// operation.
type ListVocabulariesAPIClient interface {
	ListVocabularies(context.Context, *ListVocabulariesInput, ...func(*Options)) (*ListVocabulariesOutput, error)
}

var _ ListVocabulariesAPIClient = (*Client)(nil)

// ListVocabulariesPaginatorOptions is the paginator options for ListVocabularies
type ListVocabulariesPaginatorOptions struct {
	// The maximum number of custom vocabularies to return in each page of results. If
	// there are fewer results than the value that you specify, only the actual results
	// are returned. If you do not specify a value, a default of 5 is used.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVocabulariesPaginator is a paginator for ListVocabularies
type ListVocabulariesPaginator struct {
	options   ListVocabulariesPaginatorOptions
	client    ListVocabulariesAPIClient
	params    *ListVocabulariesInput
	nextToken *string
	firstPage bool
}

// NewListVocabulariesPaginator returns a new ListVocabulariesPaginator
func NewListVocabulariesPaginator(client ListVocabulariesAPIClient, params *ListVocabulariesInput, optFns ...func(*ListVocabulariesPaginatorOptions)) *ListVocabulariesPaginator {
	if params == nil {
		params = &ListVocabulariesInput{}
	}

	options := ListVocabulariesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListVocabulariesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVocabulariesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListVocabularies page.
func (p *ListVocabulariesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVocabulariesOutput, error) {
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

	result, err := p.client.ListVocabularies(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListVocabularies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListVocabularies",
	}
}
