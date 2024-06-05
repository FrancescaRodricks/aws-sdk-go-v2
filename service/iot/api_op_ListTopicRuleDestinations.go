// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the topic rule destinations in your Amazon Web Services account.
//
// Requires permission to access the [ListTopicRuleDestinations] action.
//
// [ListTopicRuleDestinations]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) ListTopicRuleDestinations(ctx context.Context, params *ListTopicRuleDestinationsInput, optFns ...func(*Options)) (*ListTopicRuleDestinationsOutput, error) {
	if params == nil {
		params = &ListTopicRuleDestinationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTopicRuleDestinations", params, optFns, c.addOperationListTopicRuleDestinationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTopicRuleDestinationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTopicRuleDestinationsInput struct {

	// The maximum number of results to return at one time.
	MaxResults *int32

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTopicRuleDestinationsOutput struct {

	// Information about a topic rule destination.
	DestinationSummaries []types.TopicRuleDestinationSummary

	// The token to use to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTopicRuleDestinationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTopicRuleDestinations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTopicRuleDestinations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTopicRuleDestinations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTopicRuleDestinations(options.Region), middleware.Before); err != nil {
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

// ListTopicRuleDestinationsAPIClient is a client that implements the
// ListTopicRuleDestinations operation.
type ListTopicRuleDestinationsAPIClient interface {
	ListTopicRuleDestinations(context.Context, *ListTopicRuleDestinationsInput, ...func(*Options)) (*ListTopicRuleDestinationsOutput, error)
}

var _ ListTopicRuleDestinationsAPIClient = (*Client)(nil)

// ListTopicRuleDestinationsPaginatorOptions is the paginator options for
// ListTopicRuleDestinations
type ListTopicRuleDestinationsPaginatorOptions struct {
	// The maximum number of results to return at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTopicRuleDestinationsPaginator is a paginator for ListTopicRuleDestinations
type ListTopicRuleDestinationsPaginator struct {
	options   ListTopicRuleDestinationsPaginatorOptions
	client    ListTopicRuleDestinationsAPIClient
	params    *ListTopicRuleDestinationsInput
	nextToken *string
	firstPage bool
}

// NewListTopicRuleDestinationsPaginator returns a new
// ListTopicRuleDestinationsPaginator
func NewListTopicRuleDestinationsPaginator(client ListTopicRuleDestinationsAPIClient, params *ListTopicRuleDestinationsInput, optFns ...func(*ListTopicRuleDestinationsPaginatorOptions)) *ListTopicRuleDestinationsPaginator {
	if params == nil {
		params = &ListTopicRuleDestinationsInput{}
	}

	options := ListTopicRuleDestinationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTopicRuleDestinationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTopicRuleDestinationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTopicRuleDestinations page.
func (p *ListTopicRuleDestinationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTopicRuleDestinationsOutput, error) {
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

	result, err := p.client.ListTopicRuleDestinations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTopicRuleDestinations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTopicRuleDestinations",
	}
}
