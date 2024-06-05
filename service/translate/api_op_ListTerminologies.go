// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a list of custom terminologies associated with your account.
func (c *Client) ListTerminologies(ctx context.Context, params *ListTerminologiesInput, optFns ...func(*Options)) (*ListTerminologiesOutput, error) {
	if params == nil {
		params = &ListTerminologiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTerminologies", params, optFns, c.addOperationListTerminologiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTerminologiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTerminologiesInput struct {

	// The maximum number of custom terminologies returned per list request.
	MaxResults *int32

	// If the result of the request to ListTerminologies was truncated, include the
	// NextToken to fetch the next group of custom terminologies.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTerminologiesOutput struct {

	//  If the response to the ListTerminologies was truncated, the NextToken fetches
	// the next group of custom terminologies.
	NextToken *string

	// The properties list of the custom terminologies returned on the list request.
	TerminologyPropertiesList []types.TerminologyProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTerminologiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTerminologies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTerminologies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTerminologies"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTerminologies(options.Region), middleware.Before); err != nil {
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

// ListTerminologiesAPIClient is a client that implements the ListTerminologies
// operation.
type ListTerminologiesAPIClient interface {
	ListTerminologies(context.Context, *ListTerminologiesInput, ...func(*Options)) (*ListTerminologiesOutput, error)
}

var _ ListTerminologiesAPIClient = (*Client)(nil)

// ListTerminologiesPaginatorOptions is the paginator options for ListTerminologies
type ListTerminologiesPaginatorOptions struct {
	// The maximum number of custom terminologies returned per list request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTerminologiesPaginator is a paginator for ListTerminologies
type ListTerminologiesPaginator struct {
	options   ListTerminologiesPaginatorOptions
	client    ListTerminologiesAPIClient
	params    *ListTerminologiesInput
	nextToken *string
	firstPage bool
}

// NewListTerminologiesPaginator returns a new ListTerminologiesPaginator
func NewListTerminologiesPaginator(client ListTerminologiesAPIClient, params *ListTerminologiesInput, optFns ...func(*ListTerminologiesPaginatorOptions)) *ListTerminologiesPaginator {
	if params == nil {
		params = &ListTerminologiesInput{}
	}

	options := ListTerminologiesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTerminologiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTerminologiesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTerminologies page.
func (p *ListTerminologiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTerminologiesOutput, error) {
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

	result, err := p.client.ListTerminologies(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTerminologies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTerminologies",
	}
}
