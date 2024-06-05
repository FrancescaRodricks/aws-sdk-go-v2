// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists metric attributions.
func (c *Client) ListMetricAttributions(ctx context.Context, params *ListMetricAttributionsInput, optFns ...func(*Options)) (*ListMetricAttributionsOutput, error) {
	if params == nil {
		params = &ListMetricAttributionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMetricAttributions", params, optFns, c.addOperationListMetricAttributionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMetricAttributionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMetricAttributionsInput struct {

	// The metric attributions' dataset group Amazon Resource Name (ARN).
	DatasetGroupArn *string

	// The maximum number of metric attributions to return in one page of results.
	MaxResults *int32

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMetricAttributionsOutput struct {

	// The list of metric attributions.
	MetricAttributions []types.MetricAttributionSummary

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMetricAttributionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListMetricAttributions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMetricAttributions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMetricAttributions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMetricAttributions(options.Region), middleware.Before); err != nil {
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

// ListMetricAttributionsAPIClient is a client that implements the
// ListMetricAttributions operation.
type ListMetricAttributionsAPIClient interface {
	ListMetricAttributions(context.Context, *ListMetricAttributionsInput, ...func(*Options)) (*ListMetricAttributionsOutput, error)
}

var _ ListMetricAttributionsAPIClient = (*Client)(nil)

// ListMetricAttributionsPaginatorOptions is the paginator options for
// ListMetricAttributions
type ListMetricAttributionsPaginatorOptions struct {
	// The maximum number of metric attributions to return in one page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMetricAttributionsPaginator is a paginator for ListMetricAttributions
type ListMetricAttributionsPaginator struct {
	options   ListMetricAttributionsPaginatorOptions
	client    ListMetricAttributionsAPIClient
	params    *ListMetricAttributionsInput
	nextToken *string
	firstPage bool
}

// NewListMetricAttributionsPaginator returns a new ListMetricAttributionsPaginator
func NewListMetricAttributionsPaginator(client ListMetricAttributionsAPIClient, params *ListMetricAttributionsInput, optFns ...func(*ListMetricAttributionsPaginatorOptions)) *ListMetricAttributionsPaginator {
	if params == nil {
		params = &ListMetricAttributionsInput{}
	}

	options := ListMetricAttributionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMetricAttributionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMetricAttributionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMetricAttributions page.
func (p *ListMetricAttributionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMetricAttributionsOutput, error) {
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

	result, err := p.client.ListMetricAttributions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMetricAttributions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMetricAttributions",
	}
}
