// Code generated by smithy-go-codegen DO NOT EDIT.

package oam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/oam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of source account links that are linked to this monitoring
// account sink.
//
// To use this operation, provide the sink ARN. To retrieve a list of sink ARNs,
// use [ListSinks].
//
// To find a list of links for one source account, use [ListLinks].
//
// [ListLinks]: https://docs.aws.amazon.com/OAM/latest/APIReference/API_ListLinks.html
// [ListSinks]: https://docs.aws.amazon.com/OAM/latest/APIReference/API_ListSinks.html
func (c *Client) ListAttachedLinks(ctx context.Context, params *ListAttachedLinksInput, optFns ...func(*Options)) (*ListAttachedLinksOutput, error) {
	if params == nil {
		params = &ListAttachedLinksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAttachedLinks", params, optFns, c.addOperationListAttachedLinksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAttachedLinksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAttachedLinksInput struct {

	// The ARN of the sink that you want to retrieve links for.
	//
	// This member is required.
	SinkIdentifier *string

	// Limits the number of returned links to the specified number.
	MaxResults *int32

	// The token for the next set of items to return. You received this token from a
	// previous call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAttachedLinksOutput struct {

	// An array of structures that contain the information about the attached links.
	//
	// This member is required.
	Items []types.ListAttachedLinksItem

	// The token to use when requesting the next set of links.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAttachedLinksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAttachedLinks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAttachedLinks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAttachedLinks"); err != nil {
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
	if err = addOpListAttachedLinksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAttachedLinks(options.Region), middleware.Before); err != nil {
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

// ListAttachedLinksAPIClient is a client that implements the ListAttachedLinks
// operation.
type ListAttachedLinksAPIClient interface {
	ListAttachedLinks(context.Context, *ListAttachedLinksInput, ...func(*Options)) (*ListAttachedLinksOutput, error)
}

var _ ListAttachedLinksAPIClient = (*Client)(nil)

// ListAttachedLinksPaginatorOptions is the paginator options for ListAttachedLinks
type ListAttachedLinksPaginatorOptions struct {
	// Limits the number of returned links to the specified number.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAttachedLinksPaginator is a paginator for ListAttachedLinks
type ListAttachedLinksPaginator struct {
	options   ListAttachedLinksPaginatorOptions
	client    ListAttachedLinksAPIClient
	params    *ListAttachedLinksInput
	nextToken *string
	firstPage bool
}

// NewListAttachedLinksPaginator returns a new ListAttachedLinksPaginator
func NewListAttachedLinksPaginator(client ListAttachedLinksAPIClient, params *ListAttachedLinksInput, optFns ...func(*ListAttachedLinksPaginatorOptions)) *ListAttachedLinksPaginator {
	if params == nil {
		params = &ListAttachedLinksInput{}
	}

	options := ListAttachedLinksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAttachedLinksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAttachedLinksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAttachedLinks page.
func (p *ListAttachedLinksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAttachedLinksOutput, error) {
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

	result, err := p.client.ListAttachedLinks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAttachedLinks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAttachedLinks",
	}
}
