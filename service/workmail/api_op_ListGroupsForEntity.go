// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns all the groups to which an entity belongs.
func (c *Client) ListGroupsForEntity(ctx context.Context, params *ListGroupsForEntityInput, optFns ...func(*Options)) (*ListGroupsForEntityOutput, error) {
	if params == nil {
		params = &ListGroupsForEntityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGroupsForEntity", params, optFns, c.addOperationListGroupsForEntityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGroupsForEntityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGroupsForEntityInput struct {

	// The identifier for the entity.
	//
	// The entity ID can accept UserId or GroupID, Username or Groupname, or email.
	//
	//   - Entity ID: 12345678-1234-1234-1234-123456789012 or
	//   S-1-1-12-1234567890-123456789-123456789-1234
	//
	//   - Email address: entity@domain.tld
	//
	//   - Entity name: entity
	//
	// This member is required.
	EntityId *string

	// The identifier for the organization under which the entity exists.
	//
	// This member is required.
	OrganizationId *string

	// Limit the search results based on the filter criteria.
	Filters *types.ListGroupsForEntityFilters

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token to use to retrieve the next page of results. The first call does not
	// contain any tokens.
	NextToken *string

	noSmithyDocumentSerde
}

type ListGroupsForEntityOutput struct {

	// The overview of groups in an organization.
	Groups []types.GroupIdentifier

	// The token to use to retrieve the next page of results. This value is `null`
	// when there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGroupsForEntityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListGroupsForEntity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListGroupsForEntity{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListGroupsForEntity"); err != nil {
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
	if err = addOpListGroupsForEntityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGroupsForEntity(options.Region), middleware.Before); err != nil {
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

// ListGroupsForEntityAPIClient is a client that implements the
// ListGroupsForEntity operation.
type ListGroupsForEntityAPIClient interface {
	ListGroupsForEntity(context.Context, *ListGroupsForEntityInput, ...func(*Options)) (*ListGroupsForEntityOutput, error)
}

var _ ListGroupsForEntityAPIClient = (*Client)(nil)

// ListGroupsForEntityPaginatorOptions is the paginator options for
// ListGroupsForEntity
type ListGroupsForEntityPaginatorOptions struct {
	// The maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGroupsForEntityPaginator is a paginator for ListGroupsForEntity
type ListGroupsForEntityPaginator struct {
	options   ListGroupsForEntityPaginatorOptions
	client    ListGroupsForEntityAPIClient
	params    *ListGroupsForEntityInput
	nextToken *string
	firstPage bool
}

// NewListGroupsForEntityPaginator returns a new ListGroupsForEntityPaginator
func NewListGroupsForEntityPaginator(client ListGroupsForEntityAPIClient, params *ListGroupsForEntityInput, optFns ...func(*ListGroupsForEntityPaginatorOptions)) *ListGroupsForEntityPaginator {
	if params == nil {
		params = &ListGroupsForEntityInput{}
	}

	options := ListGroupsForEntityPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGroupsForEntityPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGroupsForEntityPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGroupsForEntity page.
func (p *ListGroupsForEntityPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGroupsForEntityOutput, error) {
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

	result, err := p.client.ListGroupsForEntity(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListGroupsForEntity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListGroupsForEntity",
	}
}
