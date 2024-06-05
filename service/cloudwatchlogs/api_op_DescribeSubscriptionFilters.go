// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the subscription filters for the specified log group. You can list all
// the subscription filters or filter the results by prefix. The results are
// ASCII-sorted by filter name.
func (c *Client) DescribeSubscriptionFilters(ctx context.Context, params *DescribeSubscriptionFiltersInput, optFns ...func(*Options)) (*DescribeSubscriptionFiltersOutput, error) {
	if params == nil {
		params = &DescribeSubscriptionFiltersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSubscriptionFilters", params, optFns, c.addOperationDescribeSubscriptionFiltersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSubscriptionFiltersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSubscriptionFiltersInput struct {

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string

	// The prefix to match. If you don't specify a value, no prefix filter is applied.
	FilterNamePrefix *string

	// The maximum number of items returned. If you don't specify a value, the default
	// is up to 50 items.
	Limit *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeSubscriptionFiltersOutput struct {

	// The token for the next set of items to return. The token expires after 24 hours.
	NextToken *string

	// The subscription filters.
	SubscriptionFilters []types.SubscriptionFilter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSubscriptionFiltersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSubscriptionFilters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSubscriptionFilters{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeSubscriptionFilters"); err != nil {
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
	if err = addOpDescribeSubscriptionFiltersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSubscriptionFilters(options.Region), middleware.Before); err != nil {
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

// DescribeSubscriptionFiltersAPIClient is a client that implements the
// DescribeSubscriptionFilters operation.
type DescribeSubscriptionFiltersAPIClient interface {
	DescribeSubscriptionFilters(context.Context, *DescribeSubscriptionFiltersInput, ...func(*Options)) (*DescribeSubscriptionFiltersOutput, error)
}

var _ DescribeSubscriptionFiltersAPIClient = (*Client)(nil)

// DescribeSubscriptionFiltersPaginatorOptions is the paginator options for
// DescribeSubscriptionFilters
type DescribeSubscriptionFiltersPaginatorOptions struct {
	// The maximum number of items returned. If you don't specify a value, the default
	// is up to 50 items.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSubscriptionFiltersPaginator is a paginator for
// DescribeSubscriptionFilters
type DescribeSubscriptionFiltersPaginator struct {
	options   DescribeSubscriptionFiltersPaginatorOptions
	client    DescribeSubscriptionFiltersAPIClient
	params    *DescribeSubscriptionFiltersInput
	nextToken *string
	firstPage bool
}

// NewDescribeSubscriptionFiltersPaginator returns a new
// DescribeSubscriptionFiltersPaginator
func NewDescribeSubscriptionFiltersPaginator(client DescribeSubscriptionFiltersAPIClient, params *DescribeSubscriptionFiltersInput, optFns ...func(*DescribeSubscriptionFiltersPaginatorOptions)) *DescribeSubscriptionFiltersPaginator {
	if params == nil {
		params = &DescribeSubscriptionFiltersInput{}
	}

	options := DescribeSubscriptionFiltersPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeSubscriptionFiltersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSubscriptionFiltersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeSubscriptionFilters page.
func (p *DescribeSubscriptionFiltersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSubscriptionFiltersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeSubscriptionFilters(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeSubscriptionFilters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeSubscriptionFilters",
	}
}
