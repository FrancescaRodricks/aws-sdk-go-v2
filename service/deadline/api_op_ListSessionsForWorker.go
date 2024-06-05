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

// Lists sessions for a worker.
func (c *Client) ListSessionsForWorker(ctx context.Context, params *ListSessionsForWorkerInput, optFns ...func(*Options)) (*ListSessionsForWorkerOutput, error) {
	if params == nil {
		params = &ListSessionsForWorkerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSessionsForWorker", params, optFns, c.addOperationListSessionsForWorkerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSessionsForWorkerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSessionsForWorkerInput struct {

	// The farm ID for the session.
	//
	// This member is required.
	FarmId *string

	// The fleet ID for the session.
	//
	// This member is required.
	FleetId *string

	// The worker ID for the session.
	//
	// This member is required.
	WorkerId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	MaxResults *int32

	// The token for the next set of results, or null to start from the beginning.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSessionsForWorkerOutput struct {

	// The sessions in the response.
	//
	// This member is required.
	Sessions []types.WorkerSessionSummary

	// The token for the next set of results, or null to start from the beginning.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSessionsForWorkerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSessionsForWorker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSessionsForWorker{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSessionsForWorker"); err != nil {
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
	if err = addEndpointPrefix_opListSessionsForWorkerMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListSessionsForWorkerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSessionsForWorker(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListSessionsForWorkerMiddleware struct {
}

func (*endpointPrefix_opListSessionsForWorkerMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListSessionsForWorkerMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
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
func addEndpointPrefix_opListSessionsForWorkerMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListSessionsForWorkerMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListSessionsForWorkerAPIClient is a client that implements the
// ListSessionsForWorker operation.
type ListSessionsForWorkerAPIClient interface {
	ListSessionsForWorker(context.Context, *ListSessionsForWorkerInput, ...func(*Options)) (*ListSessionsForWorkerOutput, error)
}

var _ ListSessionsForWorkerAPIClient = (*Client)(nil)

// ListSessionsForWorkerPaginatorOptions is the paginator options for
// ListSessionsForWorker
type ListSessionsForWorkerPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSessionsForWorkerPaginator is a paginator for ListSessionsForWorker
type ListSessionsForWorkerPaginator struct {
	options   ListSessionsForWorkerPaginatorOptions
	client    ListSessionsForWorkerAPIClient
	params    *ListSessionsForWorkerInput
	nextToken *string
	firstPage bool
}

// NewListSessionsForWorkerPaginator returns a new ListSessionsForWorkerPaginator
func NewListSessionsForWorkerPaginator(client ListSessionsForWorkerAPIClient, params *ListSessionsForWorkerInput, optFns ...func(*ListSessionsForWorkerPaginatorOptions)) *ListSessionsForWorkerPaginator {
	if params == nil {
		params = &ListSessionsForWorkerInput{}
	}

	options := ListSessionsForWorkerPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSessionsForWorkerPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSessionsForWorkerPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSessionsForWorker page.
func (p *ListSessionsForWorkerPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSessionsForWorkerOutput, error) {
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

	result, err := p.client.ListSessionsForWorker(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSessionsForWorker(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSessionsForWorker",
	}
}
