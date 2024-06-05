// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of Explainability resources created using the CreateExplainability operation. This
// operation returns a summary for each Explainability. You can filter the list
// using an array of Filterobjects.
//
// To retrieve the complete set of properties for a particular Explainability
// resource, use the ARN with the DescribeExplainabilityoperation.
func (c *Client) ListExplainabilities(ctx context.Context, params *ListExplainabilitiesInput, optFns ...func(*Options)) (*ListExplainabilitiesOutput, error) {
	if params == nil {
		params = &ListExplainabilitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListExplainabilities", params, optFns, c.addOperationListExplainabilitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListExplainabilitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExplainabilitiesInput struct {

	// An array of filters. For each filter, provide a condition and a match
	// statement. The condition is either IS or IS_NOT , which specifies whether to
	// include or exclude the resources that match the statement from the list. The
	// match statement consists of a key and a value.
	//
	// Filter properties
	//
	//   - Condition - The condition to apply. Valid values are IS and IS_NOT .
	//
	//   - Key - The name of the parameter to filter on. Valid values are ResourceArn
	//   and Status .
	//
	//   - Value - The value to match.
	Filters []types.Filter

	// The number of items returned in the response.
	MaxResults *int32

	// If the result of the previous request was truncated, the response includes a
	// NextToken. To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string

	noSmithyDocumentSerde
}

type ListExplainabilitiesOutput struct {

	// An array of objects that summarize the properties of each Explainability
	// resource.
	Explainabilities []types.ExplainabilitySummary

	// Returns this token if the response is truncated. To retrieve the next set of
	// results, use the token in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListExplainabilitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListExplainabilities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListExplainabilities{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListExplainabilities"); err != nil {
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
	if err = addOpListExplainabilitiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListExplainabilities(options.Region), middleware.Before); err != nil {
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

// ListExplainabilitiesAPIClient is a client that implements the
// ListExplainabilities operation.
type ListExplainabilitiesAPIClient interface {
	ListExplainabilities(context.Context, *ListExplainabilitiesInput, ...func(*Options)) (*ListExplainabilitiesOutput, error)
}

var _ ListExplainabilitiesAPIClient = (*Client)(nil)

// ListExplainabilitiesPaginatorOptions is the paginator options for
// ListExplainabilities
type ListExplainabilitiesPaginatorOptions struct {
	// The number of items returned in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListExplainabilitiesPaginator is a paginator for ListExplainabilities
type ListExplainabilitiesPaginator struct {
	options   ListExplainabilitiesPaginatorOptions
	client    ListExplainabilitiesAPIClient
	params    *ListExplainabilitiesInput
	nextToken *string
	firstPage bool
}

// NewListExplainabilitiesPaginator returns a new ListExplainabilitiesPaginator
func NewListExplainabilitiesPaginator(client ListExplainabilitiesAPIClient, params *ListExplainabilitiesInput, optFns ...func(*ListExplainabilitiesPaginatorOptions)) *ListExplainabilitiesPaginator {
	if params == nil {
		params = &ListExplainabilitiesInput{}
	}

	options := ListExplainabilitiesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListExplainabilitiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListExplainabilitiesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListExplainabilities page.
func (p *ListExplainabilitiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListExplainabilitiesOutput, error) {
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

	result, err := p.client.ListExplainabilities(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListExplainabilities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListExplainabilities",
	}
}
