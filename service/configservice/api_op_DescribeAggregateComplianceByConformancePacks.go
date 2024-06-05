// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the conformance packs and their associated compliance status
// with the count of compliant and noncompliant Config rules within each
// conformance pack. Also returns the total rule count which includes compliant
// rules, noncompliant rules, and rules that cannot be evaluated due to
// insufficient data.
//
// The results can return an empty result page, but if you have a nextToken , the
// results are displayed on the next page.
func (c *Client) DescribeAggregateComplianceByConformancePacks(ctx context.Context, params *DescribeAggregateComplianceByConformancePacksInput, optFns ...func(*Options)) (*DescribeAggregateComplianceByConformancePacksOutput, error) {
	if params == nil {
		params = &DescribeAggregateComplianceByConformancePacksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAggregateComplianceByConformancePacks", params, optFns, c.addOperationDescribeAggregateComplianceByConformancePacksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAggregateComplianceByConformancePacksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAggregateComplianceByConformancePacksInput struct {

	// The name of the configuration aggregator.
	//
	// This member is required.
	ConfigurationAggregatorName *string

	// Filters the result by AggregateConformancePackComplianceFilters object.
	Filters *types.AggregateConformancePackComplianceFilters

	// The maximum number of conformance packs compliance details returned on each
	// page. The default is maximum. If you specify 0, Config uses the default.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeAggregateComplianceByConformancePacksOutput struct {

	// Returns the AggregateComplianceByConformancePack object.
	AggregateComplianceByConformancePacks []types.AggregateComplianceByConformancePack

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAggregateComplianceByConformancePacksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAggregateComplianceByConformancePacks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAggregateComplianceByConformancePacks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAggregateComplianceByConformancePacks"); err != nil {
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
	if err = addOpDescribeAggregateComplianceByConformancePacksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAggregateComplianceByConformancePacks(options.Region), middleware.Before); err != nil {
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

// DescribeAggregateComplianceByConformancePacksAPIClient is a client that
// implements the DescribeAggregateComplianceByConformancePacks operation.
type DescribeAggregateComplianceByConformancePacksAPIClient interface {
	DescribeAggregateComplianceByConformancePacks(context.Context, *DescribeAggregateComplianceByConformancePacksInput, ...func(*Options)) (*DescribeAggregateComplianceByConformancePacksOutput, error)
}

var _ DescribeAggregateComplianceByConformancePacksAPIClient = (*Client)(nil)

// DescribeAggregateComplianceByConformancePacksPaginatorOptions is the paginator
// options for DescribeAggregateComplianceByConformancePacks
type DescribeAggregateComplianceByConformancePacksPaginatorOptions struct {
	// The maximum number of conformance packs compliance details returned on each
	// page. The default is maximum. If you specify 0, Config uses the default.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAggregateComplianceByConformancePacksPaginator is a paginator for
// DescribeAggregateComplianceByConformancePacks
type DescribeAggregateComplianceByConformancePacksPaginator struct {
	options   DescribeAggregateComplianceByConformancePacksPaginatorOptions
	client    DescribeAggregateComplianceByConformancePacksAPIClient
	params    *DescribeAggregateComplianceByConformancePacksInput
	nextToken *string
	firstPage bool
}

// NewDescribeAggregateComplianceByConformancePacksPaginator returns a new
// DescribeAggregateComplianceByConformancePacksPaginator
func NewDescribeAggregateComplianceByConformancePacksPaginator(client DescribeAggregateComplianceByConformancePacksAPIClient, params *DescribeAggregateComplianceByConformancePacksInput, optFns ...func(*DescribeAggregateComplianceByConformancePacksPaginatorOptions)) *DescribeAggregateComplianceByConformancePacksPaginator {
	if params == nil {
		params = &DescribeAggregateComplianceByConformancePacksInput{}
	}

	options := DescribeAggregateComplianceByConformancePacksPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeAggregateComplianceByConformancePacksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAggregateComplianceByConformancePacksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeAggregateComplianceByConformancePacks page.
func (p *DescribeAggregateComplianceByConformancePacksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAggregateComplianceByConformancePacksOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.DescribeAggregateComplianceByConformancePacks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeAggregateComplianceByConformancePacks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAggregateComplianceByConformancePacks",
	}
}
