// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmcontacts

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssmcontacts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the engagements to contact channels that occurred by engaging a contact.
func (c *Client) ListPagesByEngagement(ctx context.Context, params *ListPagesByEngagementInput, optFns ...func(*Options)) (*ListPagesByEngagementOutput, error) {
	if params == nil {
		params = &ListPagesByEngagementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPagesByEngagement", params, optFns, c.addOperationListPagesByEngagementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPagesByEngagementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPagesByEngagementInput struct {

	// The Amazon Resource Name (ARN) of the engagement.
	//
	// This member is required.
	EngagementId *string

	// The maximum number of engagements to contact channels to list per page of
	// results.
	MaxResults *int32

	// The pagination token to continue to the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPagesByEngagementOutput struct {

	// The list of engagements to contact channels.
	//
	// This member is required.
	Pages []types.Page

	// The pagination token to continue to the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPagesByEngagementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPagesByEngagement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPagesByEngagement{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPagesByEngagement"); err != nil {
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
	if err = addOpListPagesByEngagementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPagesByEngagement(options.Region), middleware.Before); err != nil {
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

// ListPagesByEngagementAPIClient is a client that implements the
// ListPagesByEngagement operation.
type ListPagesByEngagementAPIClient interface {
	ListPagesByEngagement(context.Context, *ListPagesByEngagementInput, ...func(*Options)) (*ListPagesByEngagementOutput, error)
}

var _ ListPagesByEngagementAPIClient = (*Client)(nil)

// ListPagesByEngagementPaginatorOptions is the paginator options for
// ListPagesByEngagement
type ListPagesByEngagementPaginatorOptions struct {
	// The maximum number of engagements to contact channels to list per page of
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPagesByEngagementPaginator is a paginator for ListPagesByEngagement
type ListPagesByEngagementPaginator struct {
	options   ListPagesByEngagementPaginatorOptions
	client    ListPagesByEngagementAPIClient
	params    *ListPagesByEngagementInput
	nextToken *string
	firstPage bool
}

// NewListPagesByEngagementPaginator returns a new ListPagesByEngagementPaginator
func NewListPagesByEngagementPaginator(client ListPagesByEngagementAPIClient, params *ListPagesByEngagementInput, optFns ...func(*ListPagesByEngagementPaginatorOptions)) *ListPagesByEngagementPaginator {
	if params == nil {
		params = &ListPagesByEngagementInput{}
	}

	options := ListPagesByEngagementPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPagesByEngagementPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPagesByEngagementPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPagesByEngagement page.
func (p *ListPagesByEngagementPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPagesByEngagementOutput, error) {
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

	result, err := p.client.ListPagesByEngagement(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPagesByEngagement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPagesByEngagement",
	}
}
