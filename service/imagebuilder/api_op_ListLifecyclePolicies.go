// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get a list of lifecycle policies in your Amazon Web Services account.
func (c *Client) ListLifecyclePolicies(ctx context.Context, params *ListLifecyclePoliciesInput, optFns ...func(*Options)) (*ListLifecyclePoliciesOutput, error) {
	if params == nil {
		params = &ListLifecyclePoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLifecyclePolicies", params, optFns, c.addOperationListLifecyclePoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLifecyclePoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLifecyclePoliciesInput struct {

	// Streamline results based on one of the following values: Name , Status .
	Filters []types.Filter

	// The maximum items to return in a request.
	MaxResults *int32

	// A token to specify where to start paginating. This is the nextToken from a
	// previously truncated response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLifecyclePoliciesOutput struct {

	// A list of lifecycle policies in your Amazon Web Services account that meet the
	// criteria specified in the request.
	LifecyclePolicySummaryList []types.LifecyclePolicySummary

	// The next token used for paginated responses. When this field isn't empty, there
	// are additional elements that the service hasn't included in this request. Use
	// this token with the next request to retrieve additional objects.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLifecyclePoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLifecyclePolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLifecyclePolicies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListLifecyclePolicies"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLifecyclePolicies(options.Region), middleware.Before); err != nil {
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

// ListLifecyclePoliciesAPIClient is a client that implements the
// ListLifecyclePolicies operation.
type ListLifecyclePoliciesAPIClient interface {
	ListLifecyclePolicies(context.Context, *ListLifecyclePoliciesInput, ...func(*Options)) (*ListLifecyclePoliciesOutput, error)
}

var _ ListLifecyclePoliciesAPIClient = (*Client)(nil)

// ListLifecyclePoliciesPaginatorOptions is the paginator options for
// ListLifecyclePolicies
type ListLifecyclePoliciesPaginatorOptions struct {
	// The maximum items to return in a request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLifecyclePoliciesPaginator is a paginator for ListLifecyclePolicies
type ListLifecyclePoliciesPaginator struct {
	options   ListLifecyclePoliciesPaginatorOptions
	client    ListLifecyclePoliciesAPIClient
	params    *ListLifecyclePoliciesInput
	nextToken *string
	firstPage bool
}

// NewListLifecyclePoliciesPaginator returns a new ListLifecyclePoliciesPaginator
func NewListLifecyclePoliciesPaginator(client ListLifecyclePoliciesAPIClient, params *ListLifecyclePoliciesInput, optFns ...func(*ListLifecyclePoliciesPaginatorOptions)) *ListLifecyclePoliciesPaginator {
	if params == nil {
		params = &ListLifecyclePoliciesInput{}
	}

	options := ListLifecyclePoliciesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLifecyclePoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLifecyclePoliciesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLifecyclePolicies page.
func (p *ListLifecyclePoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLifecyclePoliciesOutput, error) {
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

	result, err := p.client.ListLifecyclePolicies(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLifecyclePolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListLifecyclePolicies",
	}
}
