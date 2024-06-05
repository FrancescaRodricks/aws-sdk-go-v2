// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the metadata for the firewalls that you have defined. If you provide
// VPC identifiers in your request, this returns only the firewalls for those VPCs.
//
// Depending on your setting for max results and the number of firewalls, a single
// call might not return the full list.
func (c *Client) ListFirewalls(ctx context.Context, params *ListFirewallsInput, optFns ...func(*Options)) (*ListFirewallsOutput, error) {
	if params == nil {
		params = &ListFirewallsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFirewalls", params, optFns, c.addOperationListFirewallsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFirewallsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFirewallsInput struct {

	// The maximum number of objects that you want Network Firewall to return for this
	// request. If more objects are available, in the response, Network Firewall
	// provides a NextToken value that you can use in a subsequent call to get the
	// next batch of objects.
	MaxResults *int32

	// When you request a list of objects with a MaxResults setting, if the number of
	// objects that are still available for retrieval exceeds the maximum you
	// requested, Network Firewall returns a NextToken value in the response. To
	// retrieve the next batch of objects, use the token returned from the prior
	// request in your next request.
	NextToken *string

	// The unique identifiers of the VPCs that you want Network Firewall to retrieve
	// the firewalls for. Leave this blank to retrieve all firewalls that you have
	// defined.
	VpcIds []string

	noSmithyDocumentSerde
}

type ListFirewallsOutput struct {

	// The firewall metadata objects for the VPCs that you specified. Depending on
	// your setting for max results and the number of firewalls you have, a single call
	// might not be the full list.
	Firewalls []types.FirewallMetadata

	// When you request a list of objects with a MaxResults setting, if the number of
	// objects that are still available for retrieval exceeds the maximum you
	// requested, Network Firewall returns a NextToken value in the response. To
	// retrieve the next batch of objects, use the token returned from the prior
	// request in your next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFirewallsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListFirewalls{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListFirewalls{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFirewalls"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFirewalls(options.Region), middleware.Before); err != nil {
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

// ListFirewallsAPIClient is a client that implements the ListFirewalls operation.
type ListFirewallsAPIClient interface {
	ListFirewalls(context.Context, *ListFirewallsInput, ...func(*Options)) (*ListFirewallsOutput, error)
}

var _ ListFirewallsAPIClient = (*Client)(nil)

// ListFirewallsPaginatorOptions is the paginator options for ListFirewalls
type ListFirewallsPaginatorOptions struct {
	// The maximum number of objects that you want Network Firewall to return for this
	// request. If more objects are available, in the response, Network Firewall
	// provides a NextToken value that you can use in a subsequent call to get the
	// next batch of objects.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFirewallsPaginator is a paginator for ListFirewalls
type ListFirewallsPaginator struct {
	options   ListFirewallsPaginatorOptions
	client    ListFirewallsAPIClient
	params    *ListFirewallsInput
	nextToken *string
	firstPage bool
}

// NewListFirewallsPaginator returns a new ListFirewallsPaginator
func NewListFirewallsPaginator(client ListFirewallsAPIClient, params *ListFirewallsInput, optFns ...func(*ListFirewallsPaginatorOptions)) *ListFirewallsPaginator {
	if params == nil {
		params = &ListFirewallsInput{}
	}

	options := ListFirewallsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFirewallsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFirewallsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFirewalls page.
func (p *ListFirewallsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFirewallsOutput, error) {
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

	result, err := p.client.ListFirewalls(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFirewalls(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFirewalls",
	}
}
