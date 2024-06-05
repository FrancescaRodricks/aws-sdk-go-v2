// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information, including performance data and capacity usage, which
// DataSync Discovery collects about a specific resource in your-premises storage
// system.
func (c *Client) DescribeStorageSystemResourceMetrics(ctx context.Context, params *DescribeStorageSystemResourceMetricsInput, optFns ...func(*Options)) (*DescribeStorageSystemResourceMetricsOutput, error) {
	if params == nil {
		params = &DescribeStorageSystemResourceMetricsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStorageSystemResourceMetrics", params, optFns, c.addOperationDescribeStorageSystemResourceMetricsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStorageSystemResourceMetricsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStorageSystemResourceMetricsInput struct {

	// Specifies the Amazon Resource Name (ARN) of the discovery job that collects
	// information about your on-premises storage system.
	//
	// This member is required.
	DiscoveryJobArn *string

	// Specifies the universally unique identifier (UUID) of the storage system
	// resource that you want information about.
	//
	// This member is required.
	ResourceId *string

	// Specifies the kind of storage system resource that you want information about.
	//
	// This member is required.
	ResourceType types.DiscoveryResourceType

	// Specifies a time within the total duration that the discovery job ran. To see
	// information gathered during a certain time frame, use this parameter with
	// StartTime .
	EndTime *time.Time

	// Specifies how many results that you want in the response.
	MaxResults *int32

	// Specifies an opaque string that indicates the position to begin the next list
	// of results in the response.
	NextToken *string

	// Specifies a time within the total duration that the discovery job ran. To see
	// information gathered during a certain time frame, use this parameter with
	// EndTime .
	StartTime *time.Time

	noSmithyDocumentSerde
}

type DescribeStorageSystemResourceMetricsOutput struct {

	// The details that your discovery job collected about your storage system
	// resource.
	Metrics []types.ResourceMetrics

	// The opaque string that indicates the position to begin the next list of results
	// in the response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStorageSystemResourceMetricsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeStorageSystemResourceMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeStorageSystemResourceMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeStorageSystemResourceMetrics"); err != nil {
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
	if err = addEndpointPrefix_opDescribeStorageSystemResourceMetricsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeStorageSystemResourceMetricsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStorageSystemResourceMetrics(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDescribeStorageSystemResourceMetricsMiddleware struct {
}

func (*endpointPrefix_opDescribeStorageSystemResourceMetricsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribeStorageSystemResourceMetricsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "discovery-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opDescribeStorageSystemResourceMetricsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opDescribeStorageSystemResourceMetricsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// DescribeStorageSystemResourceMetricsAPIClient is a client that implements the
// DescribeStorageSystemResourceMetrics operation.
type DescribeStorageSystemResourceMetricsAPIClient interface {
	DescribeStorageSystemResourceMetrics(context.Context, *DescribeStorageSystemResourceMetricsInput, ...func(*Options)) (*DescribeStorageSystemResourceMetricsOutput, error)
}

var _ DescribeStorageSystemResourceMetricsAPIClient = (*Client)(nil)

// DescribeStorageSystemResourceMetricsPaginatorOptions is the paginator options
// for DescribeStorageSystemResourceMetrics
type DescribeStorageSystemResourceMetricsPaginatorOptions struct {
	// Specifies how many results that you want in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeStorageSystemResourceMetricsPaginator is a paginator for
// DescribeStorageSystemResourceMetrics
type DescribeStorageSystemResourceMetricsPaginator struct {
	options   DescribeStorageSystemResourceMetricsPaginatorOptions
	client    DescribeStorageSystemResourceMetricsAPIClient
	params    *DescribeStorageSystemResourceMetricsInput
	nextToken *string
	firstPage bool
}

// NewDescribeStorageSystemResourceMetricsPaginator returns a new
// DescribeStorageSystemResourceMetricsPaginator
func NewDescribeStorageSystemResourceMetricsPaginator(client DescribeStorageSystemResourceMetricsAPIClient, params *DescribeStorageSystemResourceMetricsInput, optFns ...func(*DescribeStorageSystemResourceMetricsPaginatorOptions)) *DescribeStorageSystemResourceMetricsPaginator {
	if params == nil {
		params = &DescribeStorageSystemResourceMetricsInput{}
	}

	options := DescribeStorageSystemResourceMetricsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeStorageSystemResourceMetricsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeStorageSystemResourceMetricsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeStorageSystemResourceMetrics page.
func (p *DescribeStorageSystemResourceMetricsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeStorageSystemResourceMetricsOutput, error) {
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

	result, err := p.client.DescribeStorageSystemResourceMetrics(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeStorageSystemResourceMetrics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeStorageSystemResourceMetrics",
	}
}
