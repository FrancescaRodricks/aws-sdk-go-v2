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

// Retrieves information about workspaces in the current account.
func (c *Client) ListWorkspaces(ctx context.Context, params *ListWorkspacesInput, optFns ...func(*Options)) (*ListWorkspacesOutput, error) {
	if params == nil {
		params = &ListWorkspacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkspaces", params, optFns, c.addOperationListWorkspacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkspacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkspacesInput struct {

	// The maximum number of results to return at one time. The default is 25.
	//
	// Valid Range: Minimum value of 1. Maximum value of 250.
	MaxResults *int32

	// The string that specifies the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListWorkspacesOutput struct {

	// The string that specifies the next page of results.
	NextToken *string

	// A list of objects that contain information about the workspaces.
	WorkspaceSummaries []types.WorkspaceSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkspacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWorkspaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWorkspaces{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListWorkspaces"); err != nil {
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
	if err = addEndpointPrefix_opListWorkspacesMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkspaces(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListWorkspacesMiddleware struct {
}

func (*endpointPrefix_opListWorkspacesMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListWorkspacesMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
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
func addEndpointPrefix_opListWorkspacesMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListWorkspacesMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListWorkspacesAPIClient is a client that implements the ListWorkspaces
// operation.
type ListWorkspacesAPIClient interface {
	ListWorkspaces(context.Context, *ListWorkspacesInput, ...func(*Options)) (*ListWorkspacesOutput, error)
}

var _ ListWorkspacesAPIClient = (*Client)(nil)

// ListWorkspacesPaginatorOptions is the paginator options for ListWorkspaces
type ListWorkspacesPaginatorOptions struct {
	// The maximum number of results to return at one time. The default is 25.
	//
	// Valid Range: Minimum value of 1. Maximum value of 250.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkspacesPaginator is a paginator for ListWorkspaces
type ListWorkspacesPaginator struct {
	options   ListWorkspacesPaginatorOptions
	client    ListWorkspacesAPIClient
	params    *ListWorkspacesInput
	nextToken *string
	firstPage bool
}

// NewListWorkspacesPaginator returns a new ListWorkspacesPaginator
func NewListWorkspacesPaginator(client ListWorkspacesAPIClient, params *ListWorkspacesInput, optFns ...func(*ListWorkspacesPaginatorOptions)) *ListWorkspacesPaginator {
	if params == nil {
		params = &ListWorkspacesInput{}
	}

	options := ListWorkspacesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkspacesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkspacesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorkspaces page.
func (p *ListWorkspacesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkspacesOutput, error) {
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

	result, err := p.client.ListWorkspaces(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWorkspaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListWorkspaces",
	}
}
