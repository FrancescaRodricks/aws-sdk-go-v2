// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of past predictions. The list can be filtered by detector ID,
// detector version ID, event ID, event type, or by specifying a time period. If
// filter is not specified, the most recent prediction is returned.
//
// For example, the following filter lists all past predictions for xyz event type
// - { "eventType":{ "value": "xyz" }” }
//
// This is a paginated API. If you provide a null maxResults , this action will
// retrieve a maximum of 10 records per page. If you provide a maxResults , the
// value must be between 50 and 100. To get the next page results, provide the
// nextToken from the response as part of your request. A null nextToken fetches
// the records from the beginning.
func (c *Client) ListEventPredictions(ctx context.Context, params *ListEventPredictionsInput, optFns ...func(*Options)) (*ListEventPredictionsOutput, error) {
	if params == nil {
		params = &ListEventPredictionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEventPredictions", params, optFns, c.addOperationListEventPredictionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEventPredictionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEventPredictionsInput struct {

	//  The detector ID.
	DetectorId *types.FilterCondition

	//  The detector version ID.
	DetectorVersionId *types.FilterCondition

	//  The event ID.
	EventId *types.FilterCondition

	//  The event type associated with the detector.
	EventType *types.FilterCondition

	//  The maximum number of predictions to return for the request.
	MaxResults *int32

	//  Identifies the next page of results to return. Use the token to make the call
	// again to retrieve the next page. Keep all other arguments unchanged. Each
	// pagination token expires after 24 hours.
	NextToken *string

	//  The time period for when the predictions were generated.
	PredictionTimeRange *types.PredictionTimeRange

	noSmithyDocumentSerde
}

type ListEventPredictionsOutput struct {

	//  The summary of the past predictions.
	EventPredictionSummaries []types.EventPredictionSummary

	//  Identifies the next page of results to return. Use the token to make the call
	// again to retrieve the next page. Keep all other arguments unchanged. Each
	// pagination token expires after 24 hours.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEventPredictionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEventPredictions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEventPredictions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEventPredictions"); err != nil {
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
	if err = addOpListEventPredictionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEventPredictions(options.Region), middleware.Before); err != nil {
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

// ListEventPredictionsAPIClient is a client that implements the
// ListEventPredictions operation.
type ListEventPredictionsAPIClient interface {
	ListEventPredictions(context.Context, *ListEventPredictionsInput, ...func(*Options)) (*ListEventPredictionsOutput, error)
}

var _ ListEventPredictionsAPIClient = (*Client)(nil)

// ListEventPredictionsPaginatorOptions is the paginator options for
// ListEventPredictions
type ListEventPredictionsPaginatorOptions struct {
	//  The maximum number of predictions to return for the request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEventPredictionsPaginator is a paginator for ListEventPredictions
type ListEventPredictionsPaginator struct {
	options   ListEventPredictionsPaginatorOptions
	client    ListEventPredictionsAPIClient
	params    *ListEventPredictionsInput
	nextToken *string
	firstPage bool
}

// NewListEventPredictionsPaginator returns a new ListEventPredictionsPaginator
func NewListEventPredictionsPaginator(client ListEventPredictionsAPIClient, params *ListEventPredictionsInput, optFns ...func(*ListEventPredictionsPaginatorOptions)) *ListEventPredictionsPaginator {
	if params == nil {
		params = &ListEventPredictionsInput{}
	}

	options := ListEventPredictionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEventPredictionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEventPredictionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEventPredictions page.
func (p *ListEventPredictionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEventPredictionsOutput, error) {
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

	result, err := p.client.ListEventPredictions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEventPredictions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEventPredictions",
	}
}
