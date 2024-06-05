// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the supported DPU sizes for the supported application runtimes (for
// example, Athena notebook version 1 ).
func (c *Client) ListApplicationDPUSizes(ctx context.Context, params *ListApplicationDPUSizesInput, optFns ...func(*Options)) (*ListApplicationDPUSizesOutput, error) {
	if params == nil {
		params = &ListApplicationDPUSizesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListApplicationDPUSizes", params, optFns, c.addOperationListApplicationDPUSizesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListApplicationDPUSizesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListApplicationDPUSizesInput struct {

	// Specifies the maximum number of results to return.
	MaxResults *int32

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated.
	NextToken *string

	noSmithyDocumentSerde
}

type ListApplicationDPUSizesOutput struct {

	// A list of the supported DPU sizes that the application runtime supports.
	ApplicationDPUSizes []types.ApplicationDPUSizes

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListApplicationDPUSizesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListApplicationDPUSizes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListApplicationDPUSizes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListApplicationDPUSizes"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListApplicationDPUSizes(options.Region), middleware.Before); err != nil {
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

// ListApplicationDPUSizesAPIClient is a client that implements the
// ListApplicationDPUSizes operation.
type ListApplicationDPUSizesAPIClient interface {
	ListApplicationDPUSizes(context.Context, *ListApplicationDPUSizesInput, ...func(*Options)) (*ListApplicationDPUSizesOutput, error)
}

var _ ListApplicationDPUSizesAPIClient = (*Client)(nil)

// ListApplicationDPUSizesPaginatorOptions is the paginator options for
// ListApplicationDPUSizes
type ListApplicationDPUSizesPaginatorOptions struct {
	// Specifies the maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListApplicationDPUSizesPaginator is a paginator for ListApplicationDPUSizes
type ListApplicationDPUSizesPaginator struct {
	options   ListApplicationDPUSizesPaginatorOptions
	client    ListApplicationDPUSizesAPIClient
	params    *ListApplicationDPUSizesInput
	nextToken *string
	firstPage bool
}

// NewListApplicationDPUSizesPaginator returns a new
// ListApplicationDPUSizesPaginator
func NewListApplicationDPUSizesPaginator(client ListApplicationDPUSizesAPIClient, params *ListApplicationDPUSizesInput, optFns ...func(*ListApplicationDPUSizesPaginatorOptions)) *ListApplicationDPUSizesPaginator {
	if params == nil {
		params = &ListApplicationDPUSizesInput{}
	}

	options := ListApplicationDPUSizesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListApplicationDPUSizesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListApplicationDPUSizesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListApplicationDPUSizes page.
func (p *ListApplicationDPUSizesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListApplicationDPUSizesOutput, error) {
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

	result, err := p.client.ListApplicationDPUSizes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListApplicationDPUSizes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListApplicationDPUSizes",
	}
}
