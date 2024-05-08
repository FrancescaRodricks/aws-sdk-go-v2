// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutmetrics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lookoutmetrics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the detectors in the current AWS Region.
//
// Amazon Lookout for Metrics API actions are eventually consistent. If you do a
// read operation on a resource immediately after creating or modifying it, use
// retries to allow time for the write operation to complete.
func (c *Client) ListAnomalyDetectors(ctx context.Context, params *ListAnomalyDetectorsInput, optFns ...func(*Options)) (*ListAnomalyDetectorsOutput, error) {
	if params == nil {
		params = &ListAnomalyDetectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAnomalyDetectors", params, optFns, c.addOperationListAnomalyDetectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAnomalyDetectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAnomalyDetectorsInput struct {

	// The maximum number of results to return.
	MaxResults *int32

	// If the result of the previous request was truncated, the response includes a
	// NextToken . To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAnomalyDetectorsOutput struct {

	// A list of anomaly detectors in the account in the current region.
	AnomalyDetectorSummaryList []types.AnomalyDetectorSummary

	// If the response is truncated, the service returns this token. To retrieve the
	// next set of results, use the token in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAnomalyDetectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAnomalyDetectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAnomalyDetectors{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAnomalyDetectors"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAnomalyDetectors(options.Region), middleware.Before); err != nil {
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

// ListAnomalyDetectorsAPIClient is a client that implements the
// ListAnomalyDetectors operation.
type ListAnomalyDetectorsAPIClient interface {
	ListAnomalyDetectors(context.Context, *ListAnomalyDetectorsInput, ...func(*Options)) (*ListAnomalyDetectorsOutput, error)
}

var _ ListAnomalyDetectorsAPIClient = (*Client)(nil)

// ListAnomalyDetectorsPaginatorOptions is the paginator options for
// ListAnomalyDetectors
type ListAnomalyDetectorsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAnomalyDetectorsPaginator is a paginator for ListAnomalyDetectors
type ListAnomalyDetectorsPaginator struct {
	options   ListAnomalyDetectorsPaginatorOptions
	client    ListAnomalyDetectorsAPIClient
	params    *ListAnomalyDetectorsInput
	nextToken *string
	firstPage bool
}

// NewListAnomalyDetectorsPaginator returns a new ListAnomalyDetectorsPaginator
func NewListAnomalyDetectorsPaginator(client ListAnomalyDetectorsAPIClient, params *ListAnomalyDetectorsInput, optFns ...func(*ListAnomalyDetectorsPaginatorOptions)) *ListAnomalyDetectorsPaginator {
	if params == nil {
		params = &ListAnomalyDetectorsInput{}
	}

	options := ListAnomalyDetectorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAnomalyDetectorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAnomalyDetectorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAnomalyDetectors page.
func (p *ListAnomalyDetectorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAnomalyDetectorsOutput, error) {
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

	result, err := p.client.ListAnomalyDetectors(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAnomalyDetectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAnomalyDetectors",
	}
}
