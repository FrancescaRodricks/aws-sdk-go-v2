// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the firewall domain lists that you have defined. For each firewall
// domain list, you can retrieve the domains that are defined for a list by calling
// ListFirewallDomains.
//
// A single call to this list operation might return only a partial list of the
// domain lists. For information, see MaxResults .
func (c *Client) ListFirewallDomainLists(ctx context.Context, params *ListFirewallDomainListsInput, optFns ...func(*Options)) (*ListFirewallDomainListsOutput, error) {
	if params == nil {
		params = &ListFirewallDomainListsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFirewallDomainLists", params, optFns, c.addOperationListFirewallDomainListsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFirewallDomainListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFirewallDomainListsInput struct {

	// The maximum number of objects that you want Resolver to return for this
	// request. If more objects are available, in the response, Resolver provides a
	// NextToken value that you can use in a subsequent call to get the next batch of
	// objects.
	//
	// If you don't specify a value for MaxResults , Resolver returns up to 100
	// objects.
	MaxResults *int32

	// For the first call to this list request, omit this value.
	//
	// When you request a list of objects, Resolver returns at most the number of
	// objects specified in MaxResults . If more objects are available for retrieval,
	// Resolver returns a NextToken value in the response. To retrieve the next batch
	// of objects, use the token that was returned for the prior request in your next
	// request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListFirewallDomainListsOutput struct {

	// A list of the domain lists that you have defined.
	//
	// This might be a partial list of the domain lists that you've defined. For
	// information, see MaxResults .
	FirewallDomainLists []types.FirewallDomainListMetadata

	// If objects are still available for retrieval, Resolver returns this token in
	// the response. To retrieve the next batch of objects, provide this token in your
	// next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFirewallDomainListsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFirewallDomainLists{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFirewallDomainLists{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFirewallDomainLists"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFirewallDomainLists(options.Region), middleware.Before); err != nil {
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

// ListFirewallDomainListsAPIClient is a client that implements the
// ListFirewallDomainLists operation.
type ListFirewallDomainListsAPIClient interface {
	ListFirewallDomainLists(context.Context, *ListFirewallDomainListsInput, ...func(*Options)) (*ListFirewallDomainListsOutput, error)
}

var _ ListFirewallDomainListsAPIClient = (*Client)(nil)

// ListFirewallDomainListsPaginatorOptions is the paginator options for
// ListFirewallDomainLists
type ListFirewallDomainListsPaginatorOptions struct {
	// The maximum number of objects that you want Resolver to return for this
	// request. If more objects are available, in the response, Resolver provides a
	// NextToken value that you can use in a subsequent call to get the next batch of
	// objects.
	//
	// If you don't specify a value for MaxResults , Resolver returns up to 100
	// objects.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFirewallDomainListsPaginator is a paginator for ListFirewallDomainLists
type ListFirewallDomainListsPaginator struct {
	options   ListFirewallDomainListsPaginatorOptions
	client    ListFirewallDomainListsAPIClient
	params    *ListFirewallDomainListsInput
	nextToken *string
	firstPage bool
}

// NewListFirewallDomainListsPaginator returns a new
// ListFirewallDomainListsPaginator
func NewListFirewallDomainListsPaginator(client ListFirewallDomainListsAPIClient, params *ListFirewallDomainListsInput, optFns ...func(*ListFirewallDomainListsPaginatorOptions)) *ListFirewallDomainListsPaginator {
	if params == nil {
		params = &ListFirewallDomainListsInput{}
	}

	options := ListFirewallDomainListsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFirewallDomainListsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFirewallDomainListsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFirewallDomainLists page.
func (p *ListFirewallDomainListsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFirewallDomainListsOutput, error) {
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

	result, err := p.client.ListFirewallDomainLists(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFirewallDomainLists(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFirewallDomainLists",
	}
}
