// Code generated by smithy-go-codegen DO NOT EDIT.

package dataexchange

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation lists a revision's assets sorted alphabetically in descending
// order.
func (c *Client) ListRevisionAssets(ctx context.Context, params *ListRevisionAssetsInput, optFns ...func(*Options)) (*ListRevisionAssetsOutput, error) {
	if params == nil {
		params = &ListRevisionAssetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRevisionAssets", params, optFns, c.addOperationListRevisionAssetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRevisionAssetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRevisionAssetsInput struct {

	// The unique identifier for a data set.
	//
	// This member is required.
	DataSetId *string

	// The unique identifier for a revision.
	//
	// This member is required.
	RevisionId *string

	// The maximum number of results returned by a single call.
	MaxResults int32

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRevisionAssetsOutput struct {

	// The asset objects listed by the request.
	Assets []types.AssetEntry

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRevisionAssetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRevisionAssets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRevisionAssets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRevisionAssets"); err != nil {
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
	if err = addOpListRevisionAssetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRevisionAssets(options.Region), middleware.Before); err != nil {
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

// ListRevisionAssetsAPIClient is a client that implements the ListRevisionAssets
// operation.
type ListRevisionAssetsAPIClient interface {
	ListRevisionAssets(context.Context, *ListRevisionAssetsInput, ...func(*Options)) (*ListRevisionAssetsOutput, error)
}

var _ ListRevisionAssetsAPIClient = (*Client)(nil)

// ListRevisionAssetsPaginatorOptions is the paginator options for
// ListRevisionAssets
type ListRevisionAssetsPaginatorOptions struct {
	// The maximum number of results returned by a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRevisionAssetsPaginator is a paginator for ListRevisionAssets
type ListRevisionAssetsPaginator struct {
	options   ListRevisionAssetsPaginatorOptions
	client    ListRevisionAssetsAPIClient
	params    *ListRevisionAssetsInput
	nextToken *string
	firstPage bool
}

// NewListRevisionAssetsPaginator returns a new ListRevisionAssetsPaginator
func NewListRevisionAssetsPaginator(client ListRevisionAssetsAPIClient, params *ListRevisionAssetsInput, optFns ...func(*ListRevisionAssetsPaginatorOptions)) *ListRevisionAssetsPaginator {
	if params == nil {
		params = &ListRevisionAssetsInput{}
	}

	options := ListRevisionAssetsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRevisionAssetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRevisionAssetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRevisionAssets page.
func (p *ListRevisionAssetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRevisionAssetsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListRevisionAssets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRevisionAssets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRevisionAssets",
	}
}
