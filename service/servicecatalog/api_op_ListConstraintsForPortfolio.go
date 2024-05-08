// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the constraints for the specified portfolio and product.
func (c *Client) ListConstraintsForPortfolio(ctx context.Context, params *ListConstraintsForPortfolioInput, optFns ...func(*Options)) (*ListConstraintsForPortfolioOutput, error) {
	if params == nil {
		params = &ListConstraintsForPortfolioInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConstraintsForPortfolio", params, optFns, c.addOperationListConstraintsForPortfolioMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConstraintsForPortfolioOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConstraintsForPortfolioInput struct {

	// The portfolio identifier.
	//
	// This member is required.
	PortfolioId *string

	// The language code.
	//
	//   - jp - Japanese
	//
	//   - zh - Chinese
	AcceptLanguage *string

	// The maximum number of items to return with this call.
	PageSize int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	// The product identifier.
	ProductId *string

	noSmithyDocumentSerde
}

type ListConstraintsForPortfolioOutput struct {

	// Information about the constraints.
	ConstraintDetails []types.ConstraintDetail

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListConstraintsForPortfolioMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListConstraintsForPortfolio{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListConstraintsForPortfolio{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListConstraintsForPortfolio"); err != nil {
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
	if err = addOpListConstraintsForPortfolioValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConstraintsForPortfolio(options.Region), middleware.Before); err != nil {
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

// ListConstraintsForPortfolioAPIClient is a client that implements the
// ListConstraintsForPortfolio operation.
type ListConstraintsForPortfolioAPIClient interface {
	ListConstraintsForPortfolio(context.Context, *ListConstraintsForPortfolioInput, ...func(*Options)) (*ListConstraintsForPortfolioOutput, error)
}

var _ ListConstraintsForPortfolioAPIClient = (*Client)(nil)

// ListConstraintsForPortfolioPaginatorOptions is the paginator options for
// ListConstraintsForPortfolio
type ListConstraintsForPortfolioPaginatorOptions struct {
	// The maximum number of items to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListConstraintsForPortfolioPaginator is a paginator for
// ListConstraintsForPortfolio
type ListConstraintsForPortfolioPaginator struct {
	options   ListConstraintsForPortfolioPaginatorOptions
	client    ListConstraintsForPortfolioAPIClient
	params    *ListConstraintsForPortfolioInput
	nextToken *string
	firstPage bool
}

// NewListConstraintsForPortfolioPaginator returns a new
// ListConstraintsForPortfolioPaginator
func NewListConstraintsForPortfolioPaginator(client ListConstraintsForPortfolioAPIClient, params *ListConstraintsForPortfolioInput, optFns ...func(*ListConstraintsForPortfolioPaginatorOptions)) *ListConstraintsForPortfolioPaginator {
	if params == nil {
		params = &ListConstraintsForPortfolioInput{}
	}

	options := ListConstraintsForPortfolioPaginatorOptions{}
	if params.PageSize != 0 {
		options.Limit = params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListConstraintsForPortfolioPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.PageToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListConstraintsForPortfolioPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListConstraintsForPortfolio page.
func (p *ListConstraintsForPortfolioPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListConstraintsForPortfolioOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PageToken = p.nextToken

	params.PageSize = p.options.Limit

	result, err := p.client.ListConstraintsForPortfolio(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListConstraintsForPortfolio(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListConstraintsForPortfolio",
	}
}
