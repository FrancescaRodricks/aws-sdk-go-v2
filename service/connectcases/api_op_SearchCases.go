// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcases

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connectcases/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches for cases within their associated Cases domain. Search results are
// returned as a paginated list of abridged case documents.
//
// For customer_id you must provide the full customer profile ARN in this format:
// arn:aws:profile:your AWS Region:your AWS account ID:domains/profiles domain
// name/profiles/profile ID .
func (c *Client) SearchCases(ctx context.Context, params *SearchCasesInput, optFns ...func(*Options)) (*SearchCasesOutput, error) {
	if params == nil {
		params = &SearchCasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchCases", params, optFns, c.addOperationSearchCasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchCasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchCasesInput struct {

	// The unique identifier of the Cases domain.
	//
	// This member is required.
	DomainId *string

	// The list of field identifiers to be returned as part of the response.
	Fields []types.FieldIdentifier

	// A list of filter objects.
	Filter types.CaseFilter

	// The maximum number of cases to return. The current maximum supported value is
	// 25. This is also the default value when no other value is provided.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// A word or phrase used to perform a quick search.
	SearchTerm *string

	// A list of sorts where each sort specifies a field and their sort order to be
	// applied to the results.
	Sorts []types.Sort

	noSmithyDocumentSerde
}

type SearchCasesOutput struct {

	// A list of case documents where each case contains the properties CaseId and
	// Fields where each field is a complex union structure.
	//
	// This member is required.
	Cases []*types.SearchCasesResponseItem

	// The token for the next set of results. This is null if there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchCasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchCases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchCases{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchCases"); err != nil {
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
	if err = addOpSearchCasesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchCases(options.Region), middleware.Before); err != nil {
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

// SearchCasesAPIClient is a client that implements the SearchCases operation.
type SearchCasesAPIClient interface {
	SearchCases(context.Context, *SearchCasesInput, ...func(*Options)) (*SearchCasesOutput, error)
}

var _ SearchCasesAPIClient = (*Client)(nil)

// SearchCasesPaginatorOptions is the paginator options for SearchCases
type SearchCasesPaginatorOptions struct {
	// The maximum number of cases to return. The current maximum supported value is
	// 25. This is also the default value when no other value is provided.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchCasesPaginator is a paginator for SearchCases
type SearchCasesPaginator struct {
	options   SearchCasesPaginatorOptions
	client    SearchCasesAPIClient
	params    *SearchCasesInput
	nextToken *string
	firstPage bool
}

// NewSearchCasesPaginator returns a new SearchCasesPaginator
func NewSearchCasesPaginator(client SearchCasesAPIClient, params *SearchCasesInput, optFns ...func(*SearchCasesPaginatorOptions)) *SearchCasesPaginator {
	if params == nil {
		params = &SearchCasesInput{}
	}

	options := SearchCasesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchCasesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchCasesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchCases page.
func (p *SearchCasesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchCasesOutput, error) {
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

	result, err := p.client.SearchCases(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchCases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchCases",
	}
}
