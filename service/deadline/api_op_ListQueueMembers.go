// Code generated by smithy-go-codegen DO NOT EDIT.

package deadline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/deadline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the members in a queue.
func (c *Client) ListQueueMembers(ctx context.Context, params *ListQueueMembersInput, optFns ...func(*Options)) (*ListQueueMembersOutput, error) {
	if params == nil {
		params = &ListQueueMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListQueueMembers", params, optFns, c.addOperationListQueueMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListQueueMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListQueueMembersInput struct {

	// The farm ID for the queue.
	//
	// This member is required.
	FarmId *string

	// The queue ID to include on the list.
	//
	// This member is required.
	QueueId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	MaxResults *int32

	// The token for the next set of results, or null to start from the beginning.
	NextToken *string

	noSmithyDocumentSerde
}

type ListQueueMembersOutput struct {

	// The members on the list.
	//
	// This member is required.
	Members []types.QueueMember

	// If Deadline Cloud returns nextToken , then there are more results available. The
	// value of nextToken is a unique pagination token for each page. To retrieve the
	// next page, call the operation again using the returned token. Keep all other
	// arguments unchanged. If no results remain, then nextToken is set to null . Each
	// pagination token expires after 24 hours. If you provide a token that isn't
	// valid, then you receive an HTTP 400 ValidationException error.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListQueueMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListQueueMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListQueueMembers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListQueueMembers"); err != nil {
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
	if err = addEndpointPrefix_opListQueueMembersMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListQueueMembersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListQueueMembers(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListQueueMembersMiddleware struct {
}

func (*endpointPrefix_opListQueueMembersMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListQueueMembersMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "management." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListQueueMembersMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListQueueMembersMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListQueueMembersAPIClient is a client that implements the ListQueueMembers
// operation.
type ListQueueMembersAPIClient interface {
	ListQueueMembers(context.Context, *ListQueueMembersInput, ...func(*Options)) (*ListQueueMembersOutput, error)
}

var _ ListQueueMembersAPIClient = (*Client)(nil)

// ListQueueMembersPaginatorOptions is the paginator options for ListQueueMembers
type ListQueueMembersPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListQueueMembersPaginator is a paginator for ListQueueMembers
type ListQueueMembersPaginator struct {
	options   ListQueueMembersPaginatorOptions
	client    ListQueueMembersAPIClient
	params    *ListQueueMembersInput
	nextToken *string
	firstPage bool
}

// NewListQueueMembersPaginator returns a new ListQueueMembersPaginator
func NewListQueueMembersPaginator(client ListQueueMembersAPIClient, params *ListQueueMembersInput, optFns ...func(*ListQueueMembersPaginatorOptions)) *ListQueueMembersPaginator {
	if params == nil {
		params = &ListQueueMembersInput{}
	}

	options := ListQueueMembersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListQueueMembersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListQueueMembersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListQueueMembers page.
func (p *ListQueueMembersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListQueueMembersOutput, error) {
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

	result, err := p.client.ListQueueMembers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListQueueMembers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListQueueMembers",
	}
}
