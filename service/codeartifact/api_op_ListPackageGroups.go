// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of package groups in the requested domain.
func (c *Client) ListPackageGroups(ctx context.Context, params *ListPackageGroupsInput, optFns ...func(*Options)) (*ListPackageGroupsOutput, error) {
	if params == nil {
		params = &ListPackageGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPackageGroups", params, optFns, c.addOperationListPackageGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPackageGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPackageGroupsInput struct {

	//  The domain for which you want to list package groups.
	//
	// This member is required.
	Domain *string

	//  The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	//  The maximum number of results to return per page.
	MaxResults *int32

	//  The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	//  A prefix for which to search package groups. When included, ListPackageGroups
	// will return only package groups with patterns that match the prefix.
	Prefix *string

	noSmithyDocumentSerde
}

type ListPackageGroupsOutput struct {

	//  The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	//  The list of package groups in the requested domain.
	PackageGroups []types.PackageGroupSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPackageGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPackageGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPackageGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPackageGroups"); err != nil {
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
	if err = addOpListPackageGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPackageGroups(options.Region), middleware.Before); err != nil {
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

// ListPackageGroupsAPIClient is a client that implements the ListPackageGroups
// operation.
type ListPackageGroupsAPIClient interface {
	ListPackageGroups(context.Context, *ListPackageGroupsInput, ...func(*Options)) (*ListPackageGroupsOutput, error)
}

var _ ListPackageGroupsAPIClient = (*Client)(nil)

// ListPackageGroupsPaginatorOptions is the paginator options for ListPackageGroups
type ListPackageGroupsPaginatorOptions struct {
	//  The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPackageGroupsPaginator is a paginator for ListPackageGroups
type ListPackageGroupsPaginator struct {
	options   ListPackageGroupsPaginatorOptions
	client    ListPackageGroupsAPIClient
	params    *ListPackageGroupsInput
	nextToken *string
	firstPage bool
}

// NewListPackageGroupsPaginator returns a new ListPackageGroupsPaginator
func NewListPackageGroupsPaginator(client ListPackageGroupsAPIClient, params *ListPackageGroupsInput, optFns ...func(*ListPackageGroupsPaginatorOptions)) *ListPackageGroupsPaginator {
	if params == nil {
		params = &ListPackageGroupsInput{}
	}

	options := ListPackageGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPackageGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPackageGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPackageGroups page.
func (p *ListPackageGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPackageGroupsOutput, error) {
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

	result, err := p.client.ListPackageGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPackageGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPackageGroups",
	}
}
