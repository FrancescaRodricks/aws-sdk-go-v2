// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists operations that match the criteria that you specify.
func (c *Client) ListOperations(ctx context.Context, params *ListOperationsInput, optFns ...func(*Options)) (*ListOperationsOutput, error) {
	if params == nil {
		params = &ListOperationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOperations", params, optFns, c.addOperationListOperationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOperationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOperationsInput struct {

	// A complex type that contains specifications for the operations that you want to
	// list, for example, operations that you started between a specified start date
	// and end date.
	//
	// If you specify more than one filter, an operation must match all filters to be
	// returned by ListOperations .
	Filters []types.OperationFilter

	// The maximum number of items that you want Cloud Map to return in the response
	// to a ListOperations request. If you don't specify a value for MaxResults , Cloud
	// Map returns up to 100 operations.
	MaxResults *int32

	// For the first ListOperations request, omit this value.
	//
	// If the response contains NextToken , submit another ListOperations request to
	// get the next group of results. Specify the value of NextToken from the previous
	// response in the next request.
	//
	// Cloud Map gets MaxResults operations and then filters them based on the
	// specified criteria. It's possible that no operations in the first MaxResults
	// operations matched the specified criteria but that subsequent groups of
	// MaxResults operations do contain operations that match the criteria.
	NextToken *string

	noSmithyDocumentSerde
}

type ListOperationsOutput struct {

	// If the response contains NextToken , submit another ListOperations request to
	// get the next group of results. Specify the value of NextToken from the previous
	// response in the next request.
	//
	// Cloud Map gets MaxResults operations and then filters them based on the
	// specified criteria. It's possible that no operations in the first MaxResults
	// operations matched the specified criteria but that subsequent groups of
	// MaxResults operations do contain operations that match the criteria.
	NextToken *string

	// Summary information about the operations that match the specified criteria.
	Operations []types.OperationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOperationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListOperations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListOperations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListOperations"); err != nil {
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
	if err = addOpListOperationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOperations(options.Region), middleware.Before); err != nil {
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

// ListOperationsAPIClient is a client that implements the ListOperations
// operation.
type ListOperationsAPIClient interface {
	ListOperations(context.Context, *ListOperationsInput, ...func(*Options)) (*ListOperationsOutput, error)
}

var _ ListOperationsAPIClient = (*Client)(nil)

// ListOperationsPaginatorOptions is the paginator options for ListOperations
type ListOperationsPaginatorOptions struct {
	// The maximum number of items that you want Cloud Map to return in the response
	// to a ListOperations request. If you don't specify a value for MaxResults , Cloud
	// Map returns up to 100 operations.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOperationsPaginator is a paginator for ListOperations
type ListOperationsPaginator struct {
	options   ListOperationsPaginatorOptions
	client    ListOperationsAPIClient
	params    *ListOperationsInput
	nextToken *string
	firstPage bool
}

// NewListOperationsPaginator returns a new ListOperationsPaginator
func NewListOperationsPaginator(client ListOperationsAPIClient, params *ListOperationsInput, optFns ...func(*ListOperationsPaginatorOptions)) *ListOperationsPaginator {
	if params == nil {
		params = &ListOperationsInput{}
	}

	options := ListOperationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOperationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOperationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOperations page.
func (p *ListOperationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOperationsOutput, error) {
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

	result, err := p.client.ListOperations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListOperations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListOperations",
	}
}
