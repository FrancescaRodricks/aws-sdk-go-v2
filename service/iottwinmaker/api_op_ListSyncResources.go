// Code generated by smithy-go-codegen DO NOT EDIT.

package iottwinmaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iottwinmaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the sync resources.
func (c *Client) ListSyncResources(ctx context.Context, params *ListSyncResourcesInput, optFns ...func(*Options)) (*ListSyncResourcesOutput, error) {
	if params == nil {
		params = &ListSyncResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSyncResources", params, optFns, c.addOperationListSyncResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSyncResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSyncResourcesInput struct {

	// The sync source.
	//
	// Currently the only supported syncSource is SITEWISE .
	//
	// This member is required.
	SyncSource *string

	// The ID of the workspace that contains the sync job.
	//
	// This member is required.
	WorkspaceId *string

	// A list of objects that filter the request.
	//
	// The following filter combinations are supported:
	//
	//   - Filter with state
	//
	//   - Filter with ResourceType and ResourceId
	//
	//   - Filter with ResourceType and ExternalId
	Filters []types.SyncResourceFilter

	// The maximum number of results to return at one time. The default is 50.
	//
	// Valid Range: Minimum value of 0. Maximum value of 200.
	MaxResults *int32

	// The string that specifies the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSyncResourcesOutput struct {

	// The string that specifies the next page of results.
	NextToken *string

	// The sync resources.
	SyncResources []types.SyncResourceSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSyncResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSyncResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSyncResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSyncResources"); err != nil {
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
	if err = addEndpointPrefix_opListSyncResourcesMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListSyncResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSyncResources(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListSyncResourcesMiddleware struct {
}

func (*endpointPrefix_opListSyncResourcesMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListSyncResourcesMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListSyncResourcesMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListSyncResourcesMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListSyncResourcesAPIClient is a client that implements the ListSyncResources
// operation.
type ListSyncResourcesAPIClient interface {
	ListSyncResources(context.Context, *ListSyncResourcesInput, ...func(*Options)) (*ListSyncResourcesOutput, error)
}

var _ ListSyncResourcesAPIClient = (*Client)(nil)

// ListSyncResourcesPaginatorOptions is the paginator options for ListSyncResources
type ListSyncResourcesPaginatorOptions struct {
	// The maximum number of results to return at one time. The default is 50.
	//
	// Valid Range: Minimum value of 0. Maximum value of 200.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSyncResourcesPaginator is a paginator for ListSyncResources
type ListSyncResourcesPaginator struct {
	options   ListSyncResourcesPaginatorOptions
	client    ListSyncResourcesAPIClient
	params    *ListSyncResourcesInput
	nextToken *string
	firstPage bool
}

// NewListSyncResourcesPaginator returns a new ListSyncResourcesPaginator
func NewListSyncResourcesPaginator(client ListSyncResourcesAPIClient, params *ListSyncResourcesInput, optFns ...func(*ListSyncResourcesPaginatorOptions)) *ListSyncResourcesPaginator {
	if params == nil {
		params = &ListSyncResourcesInput{}
	}

	options := ListSyncResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSyncResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSyncResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSyncResources page.
func (p *ListSyncResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSyncResourcesOutput, error) {
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

	result, err := p.client.ListSyncResources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSyncResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSyncResources",
	}
}
