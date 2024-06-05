// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches quick connects in an Amazon Connect instance, with optional filtering.
func (c *Client) SearchQuickConnects(ctx context.Context, params *SearchQuickConnectsInput, optFns ...func(*Options)) (*SearchQuickConnectsOutput, error) {
	if params == nil {
		params = &SearchQuickConnectsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchQuickConnects", params, optFns, c.addOperationSearchQuickConnectsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchQuickConnectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchQuickConnectsInput struct {

	// The identifier of the Amazon Connect instance. You can [find the instance ID] in the Amazon Resource
	// Name (ARN) of the instance.
	//
	// [find the instance ID]: https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html
	//
	// This member is required.
	InstanceId *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The search criteria to be used to return quick connects.
	SearchCriteria *types.QuickConnectSearchCriteria

	// Filters to be applied to search results.
	SearchFilter *types.QuickConnectSearchFilter

	noSmithyDocumentSerde
}

type SearchQuickConnectsOutput struct {

	// The total number of quick connects which matched your search query.
	ApproximateTotalCount *int64

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Information about the quick connects.
	QuickConnects []types.QuickConnect

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchQuickConnectsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchQuickConnects{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchQuickConnects{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchQuickConnects"); err != nil {
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
	if err = addOpSearchQuickConnectsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchQuickConnects(options.Region), middleware.Before); err != nil {
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

// SearchQuickConnectsAPIClient is a client that implements the
// SearchQuickConnects operation.
type SearchQuickConnectsAPIClient interface {
	SearchQuickConnects(context.Context, *SearchQuickConnectsInput, ...func(*Options)) (*SearchQuickConnectsOutput, error)
}

var _ SearchQuickConnectsAPIClient = (*Client)(nil)

// SearchQuickConnectsPaginatorOptions is the paginator options for
// SearchQuickConnects
type SearchQuickConnectsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchQuickConnectsPaginator is a paginator for SearchQuickConnects
type SearchQuickConnectsPaginator struct {
	options   SearchQuickConnectsPaginatorOptions
	client    SearchQuickConnectsAPIClient
	params    *SearchQuickConnectsInput
	nextToken *string
	firstPage bool
}

// NewSearchQuickConnectsPaginator returns a new SearchQuickConnectsPaginator
func NewSearchQuickConnectsPaginator(client SearchQuickConnectsAPIClient, params *SearchQuickConnectsInput, optFns ...func(*SearchQuickConnectsPaginatorOptions)) *SearchQuickConnectsPaginator {
	if params == nil {
		params = &SearchQuickConnectsInput{}
	}

	options := SearchQuickConnectsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchQuickConnectsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchQuickConnectsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchQuickConnects page.
func (p *SearchQuickConnectsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchQuickConnectsOutput, error) {
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

	result, err := p.client.SearchQuickConnects(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchQuickConnects(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchQuickConnects",
	}
}
