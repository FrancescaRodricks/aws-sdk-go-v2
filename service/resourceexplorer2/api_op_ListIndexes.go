// Code generated by smithy-go-codegen DO NOT EDIT.

package resourceexplorer2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/resourceexplorer2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of all of the indexes in Amazon Web Services Regions that are
// currently collecting resource information for Amazon Web Services Resource
// Explorer.
func (c *Client) ListIndexes(ctx context.Context, params *ListIndexesInput, optFns ...func(*Options)) (*ListIndexesOutput, error) {
	if params == nil {
		params = &ListIndexesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIndexes", params, optFns, c.addOperationListIndexesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIndexesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIndexesInput struct {

	// The maximum number of results that you want included on each page of the
	// response. If you do not include this parameter, it defaults to a value
	// appropriate to the operation. If additional items exist beyond those included in
	// the current response, the NextToken response element is present and has a value
	// (is not null). Include that value as the NextToken request parameter in the
	// next call to the operation to get the next part of the results.
	//
	// An API operation can return fewer results than the maximum even when there are
	// more results available. You should check NextToken after every operation to
	// ensure that you receive all of the results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from. The pagination
	// tokens expire after 24 hours.
	NextToken *string

	// If specified, limits the response to only information about the index in the
	// specified list of Amazon Web Services Regions.
	Regions []string

	// If specified, limits the output to only indexes of the specified Type, either
	// LOCAL or AGGREGATOR .
	//
	// Use this option to discover the aggregator index for your account.
	Type types.IndexType

	noSmithyDocumentSerde
}

type ListIndexesOutput struct {

	// A structure that contains the details and status of each index.
	Indexes []types.Index

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null . The
	// pagination tokens expire after 24 hours.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListIndexesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListIndexes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListIndexes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListIndexes"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListIndexes(options.Region), middleware.Before); err != nil {
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

// ListIndexesAPIClient is a client that implements the ListIndexes operation.
type ListIndexesAPIClient interface {
	ListIndexes(context.Context, *ListIndexesInput, ...func(*Options)) (*ListIndexesOutput, error)
}

var _ ListIndexesAPIClient = (*Client)(nil)

// ListIndexesPaginatorOptions is the paginator options for ListIndexes
type ListIndexesPaginatorOptions struct {
	// The maximum number of results that you want included on each page of the
	// response. If you do not include this parameter, it defaults to a value
	// appropriate to the operation. If additional items exist beyond those included in
	// the current response, the NextToken response element is present and has a value
	// (is not null). Include that value as the NextToken request parameter in the
	// next call to the operation to get the next part of the results.
	//
	// An API operation can return fewer results than the maximum even when there are
	// more results available. You should check NextToken after every operation to
	// ensure that you receive all of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListIndexesPaginator is a paginator for ListIndexes
type ListIndexesPaginator struct {
	options   ListIndexesPaginatorOptions
	client    ListIndexesAPIClient
	params    *ListIndexesInput
	nextToken *string
	firstPage bool
}

// NewListIndexesPaginator returns a new ListIndexesPaginator
func NewListIndexesPaginator(client ListIndexesAPIClient, params *ListIndexesInput, optFns ...func(*ListIndexesPaginatorOptions)) *ListIndexesPaginator {
	if params == nil {
		params = &ListIndexesInput{}
	}

	options := ListIndexesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListIndexesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListIndexesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListIndexes page.
func (p *ListIndexesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListIndexesOutput, error) {
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

	result, err := p.client.ListIndexes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListIndexes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListIndexes",
	}
}
