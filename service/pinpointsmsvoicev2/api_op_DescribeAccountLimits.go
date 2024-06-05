// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the current Amazon Pinpoint SMS Voice V2 resource quotas for your
// account. The description for a quota includes the quota name, current usage
// toward that quota, and the quota's maximum value.
//
// When you establish an Amazon Web Services account, the account has initial
// quotas on the maximum number of configuration sets, opt-out lists, phone
// numbers, and pools that you can create in a given Region. For more information
// see [Amazon Pinpoint quotas]in the Amazon Pinpoint Developer Guide.
//
// [Amazon Pinpoint quotas]: https://docs.aws.amazon.com/pinpoint/latest/developerguide/quotas.html
func (c *Client) DescribeAccountLimits(ctx context.Context, params *DescribeAccountLimitsInput, optFns ...func(*Options)) (*DescribeAccountLimitsOutput, error) {
	if params == nil {
		params = &DescribeAccountLimitsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccountLimits", params, optFns, c.addOperationDescribeAccountLimitsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAccountLimitsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccountLimitsInput struct {

	// The maximum number of results to return per each request.
	MaxResults *int32

	// The token to be used for the next set of paginated results. You don't need to
	// supply a value for this field in the initial request.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeAccountLimitsOutput struct {

	// An array of AccountLimit objects that show the current spend limits.
	AccountLimits []types.AccountLimit

	// The token to be used for the next set of paginated results. If this field is
	// empty then there are no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAccountLimitsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeAccountLimits{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeAccountLimits{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAccountLimits"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccountLimits(options.Region), middleware.Before); err != nil {
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

// DescribeAccountLimitsAPIClient is a client that implements the
// DescribeAccountLimits operation.
type DescribeAccountLimitsAPIClient interface {
	DescribeAccountLimits(context.Context, *DescribeAccountLimitsInput, ...func(*Options)) (*DescribeAccountLimitsOutput, error)
}

var _ DescribeAccountLimitsAPIClient = (*Client)(nil)

// DescribeAccountLimitsPaginatorOptions is the paginator options for
// DescribeAccountLimits
type DescribeAccountLimitsPaginatorOptions struct {
	// The maximum number of results to return per each request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAccountLimitsPaginator is a paginator for DescribeAccountLimits
type DescribeAccountLimitsPaginator struct {
	options   DescribeAccountLimitsPaginatorOptions
	client    DescribeAccountLimitsAPIClient
	params    *DescribeAccountLimitsInput
	nextToken *string
	firstPage bool
}

// NewDescribeAccountLimitsPaginator returns a new DescribeAccountLimitsPaginator
func NewDescribeAccountLimitsPaginator(client DescribeAccountLimitsAPIClient, params *DescribeAccountLimitsInput, optFns ...func(*DescribeAccountLimitsPaginatorOptions)) *DescribeAccountLimitsPaginator {
	if params == nil {
		params = &DescribeAccountLimitsInput{}
	}

	options := DescribeAccountLimitsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeAccountLimitsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAccountLimitsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeAccountLimits page.
func (p *DescribeAccountLimitsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAccountLimitsOutput, error) {
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

	result, err := p.client.DescribeAccountLimits(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeAccountLimits(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAccountLimits",
	}
}
