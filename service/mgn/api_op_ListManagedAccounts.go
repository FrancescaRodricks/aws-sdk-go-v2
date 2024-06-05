// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List Managed Accounts.
func (c *Client) ListManagedAccounts(ctx context.Context, params *ListManagedAccountsInput, optFns ...func(*Options)) (*ListManagedAccountsOutput, error) {
	if params == nil {
		params = &ListManagedAccountsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListManagedAccounts", params, optFns, c.addOperationListManagedAccountsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListManagedAccountsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// List managed accounts request.
type ListManagedAccountsInput struct {

	// List managed accounts request max results.
	MaxResults *int32

	// List managed accounts request next token.
	NextToken *string

	noSmithyDocumentSerde
}

// List managed accounts response.
type ListManagedAccountsOutput struct {

	// List managed accounts response items.
	//
	// This member is required.
	Items []types.ManagedAccount

	// List managed accounts response next token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListManagedAccountsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListManagedAccounts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListManagedAccounts{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListManagedAccounts"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListManagedAccounts(options.Region), middleware.Before); err != nil {
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

// ListManagedAccountsAPIClient is a client that implements the
// ListManagedAccounts operation.
type ListManagedAccountsAPIClient interface {
	ListManagedAccounts(context.Context, *ListManagedAccountsInput, ...func(*Options)) (*ListManagedAccountsOutput, error)
}

var _ ListManagedAccountsAPIClient = (*Client)(nil)

// ListManagedAccountsPaginatorOptions is the paginator options for
// ListManagedAccounts
type ListManagedAccountsPaginatorOptions struct {
	// List managed accounts request max results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListManagedAccountsPaginator is a paginator for ListManagedAccounts
type ListManagedAccountsPaginator struct {
	options   ListManagedAccountsPaginatorOptions
	client    ListManagedAccountsAPIClient
	params    *ListManagedAccountsInput
	nextToken *string
	firstPage bool
}

// NewListManagedAccountsPaginator returns a new ListManagedAccountsPaginator
func NewListManagedAccountsPaginator(client ListManagedAccountsAPIClient, params *ListManagedAccountsInput, optFns ...func(*ListManagedAccountsPaginatorOptions)) *ListManagedAccountsPaginator {
	if params == nil {
		params = &ListManagedAccountsInput{}
	}

	options := ListManagedAccountsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListManagedAccountsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListManagedAccountsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListManagedAccounts page.
func (p *ListManagedAccountsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListManagedAccountsOutput, error) {
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

	result, err := p.client.ListManagedAccounts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListManagedAccounts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListManagedAccounts",
	}
}
