// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the Amazon OpenSearch Service instances that you have reserved in a
// given Region. For more information, see [Reserved Instances in Amazon OpenSearch Service].
//
// [Reserved Instances in Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ri.html
func (c *Client) DescribeReservedInstances(ctx context.Context, params *DescribeReservedInstancesInput, optFns ...func(*Options)) (*DescribeReservedInstancesOutput, error) {
	if params == nil {
		params = &DescribeReservedInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReservedInstances", params, optFns, c.addOperationDescribeReservedInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReservedInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the request parameters to the DescribeReservedInstances operation.
type DescribeReservedInstancesInput struct {

	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	MaxResults int32

	// If your initial DescribeReservedInstances operation returns a nextToken , you
	// can include the returned nextToken in subsequent DescribeReservedInstances
	// operations, which returns results in the next page.
	NextToken *string

	// The reserved instance identifier filter value. Use this parameter to show only
	// the reservation that matches the specified reserved OpenSearch instance ID.
	ReservedInstanceId *string

	noSmithyDocumentSerde
}

// Container for results from DescribeReservedInstances
type DescribeReservedInstancesOutput struct {

	// When nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Send the request again
	// using the returned token to retrieve the next page.
	NextToken *string

	// List of Reserved Instances in the current Region.
	ReservedInstances []types.ReservedInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReservedInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeReservedInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeReservedInstances{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeReservedInstances"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservedInstances(options.Region), middleware.Before); err != nil {
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

// DescribeReservedInstancesAPIClient is a client that implements the
// DescribeReservedInstances operation.
type DescribeReservedInstancesAPIClient interface {
	DescribeReservedInstances(context.Context, *DescribeReservedInstancesInput, ...func(*Options)) (*DescribeReservedInstancesOutput, error)
}

var _ DescribeReservedInstancesAPIClient = (*Client)(nil)

// DescribeReservedInstancesPaginatorOptions is the paginator options for
// DescribeReservedInstances
type DescribeReservedInstancesPaginatorOptions struct {
	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReservedInstancesPaginator is a paginator for DescribeReservedInstances
type DescribeReservedInstancesPaginator struct {
	options   DescribeReservedInstancesPaginatorOptions
	client    DescribeReservedInstancesAPIClient
	params    *DescribeReservedInstancesInput
	nextToken *string
	firstPage bool
}

// NewDescribeReservedInstancesPaginator returns a new
// DescribeReservedInstancesPaginator
func NewDescribeReservedInstancesPaginator(client DescribeReservedInstancesAPIClient, params *DescribeReservedInstancesInput, optFns ...func(*DescribeReservedInstancesPaginatorOptions)) *DescribeReservedInstancesPaginator {
	if params == nil {
		params = &DescribeReservedInstancesInput{}
	}

	options := DescribeReservedInstancesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReservedInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReservedInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeReservedInstances page.
func (p *DescribeReservedInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReservedInstancesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeReservedInstances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeReservedInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeReservedInstances",
	}
}
