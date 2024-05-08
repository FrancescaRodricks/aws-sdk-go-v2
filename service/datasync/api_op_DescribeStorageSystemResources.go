// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information that DataSync Discovery collects about resources in your
// on-premises storage system.
func (c *Client) DescribeStorageSystemResources(ctx context.Context, params *DescribeStorageSystemResourcesInput, optFns ...func(*Options)) (*DescribeStorageSystemResourcesOutput, error) {
	if params == nil {
		params = &DescribeStorageSystemResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStorageSystemResources", params, optFns, c.addOperationDescribeStorageSystemResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStorageSystemResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStorageSystemResourcesInput struct {

	// Specifies the Amazon Resource Name (ARN) of the discovery job that's collecting
	// data from your on-premises storage system.
	//
	// This member is required.
	DiscoveryJobArn *string

	// Specifies what kind of storage system resources that you want information about.
	//
	// This member is required.
	ResourceType types.DiscoveryResourceType

	// Filters the storage system resources that you want returned. For example, this
	// might be volumes associated with a specific storage virtual machine (SVM).
	Filter map[string][]string

	// Specifies the maximum number of storage system resources that you want to list
	// in a response.
	MaxResults *int32

	// Specifies an opaque string that indicates the position to begin the next list
	// of results in the response.
	NextToken *string

	// Specifies the universally unique identifiers (UUIDs) of the storage system
	// resources that you want information about. You can't use this parameter in
	// combination with the Filter parameter.
	ResourceIds []string

	noSmithyDocumentSerde
}

type DescribeStorageSystemResourcesOutput struct {

	// The opaque string that indicates the position to begin the next list of results
	// in the response.
	NextToken *string

	// The information collected about your storage system's resources. A response can
	// also include Amazon Web Services storage service recommendations.
	//
	// For more information, see [storage resource information] collected by and [recommendations] provided by DataSync Discovery.
	//
	// [storage resource information]: https://docs.aws.amazon.com/datasync/latest/userguide/discovery-understand-findings.html
	// [recommendations]: https://docs.aws.amazon.com/datasync/latest/userguide/discovery-understand-recommendations.html
	ResourceDetails *types.ResourceDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStorageSystemResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeStorageSystemResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeStorageSystemResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeStorageSystemResources"); err != nil {
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
	if err = addEndpointPrefix_opDescribeStorageSystemResourcesMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeStorageSystemResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStorageSystemResources(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDescribeStorageSystemResourcesMiddleware struct {
}

func (*endpointPrefix_opDescribeStorageSystemResourcesMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribeStorageSystemResourcesMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
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
func addEndpointPrefix_opDescribeStorageSystemResourcesMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opDescribeStorageSystemResourcesMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// DescribeStorageSystemResourcesAPIClient is a client that implements the
// DescribeStorageSystemResources operation.
type DescribeStorageSystemResourcesAPIClient interface {
	DescribeStorageSystemResources(context.Context, *DescribeStorageSystemResourcesInput, ...func(*Options)) (*DescribeStorageSystemResourcesOutput, error)
}

var _ DescribeStorageSystemResourcesAPIClient = (*Client)(nil)

// DescribeStorageSystemResourcesPaginatorOptions is the paginator options for
// DescribeStorageSystemResources
type DescribeStorageSystemResourcesPaginatorOptions struct {
	// Specifies the maximum number of storage system resources that you want to list
	// in a response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeStorageSystemResourcesPaginator is a paginator for
// DescribeStorageSystemResources
type DescribeStorageSystemResourcesPaginator struct {
	options   DescribeStorageSystemResourcesPaginatorOptions
	client    DescribeStorageSystemResourcesAPIClient
	params    *DescribeStorageSystemResourcesInput
	nextToken *string
	firstPage bool
}

// NewDescribeStorageSystemResourcesPaginator returns a new
// DescribeStorageSystemResourcesPaginator
func NewDescribeStorageSystemResourcesPaginator(client DescribeStorageSystemResourcesAPIClient, params *DescribeStorageSystemResourcesInput, optFns ...func(*DescribeStorageSystemResourcesPaginatorOptions)) *DescribeStorageSystemResourcesPaginator {
	if params == nil {
		params = &DescribeStorageSystemResourcesInput{}
	}

	options := DescribeStorageSystemResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeStorageSystemResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeStorageSystemResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeStorageSystemResources page.
func (p *DescribeStorageSystemResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeStorageSystemResourcesOutput, error) {
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

	result, err := p.client.DescribeStorageSystemResources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeStorageSystemResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeStorageSystemResources",
	}
}
