// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the target networks associated with the specified Client VPN endpoint.
func (c *Client) DescribeClientVpnTargetNetworks(ctx context.Context, params *DescribeClientVpnTargetNetworksInput, optFns ...func(*Options)) (*DescribeClientVpnTargetNetworksOutput, error) {
	if params == nil {
		params = &DescribeClientVpnTargetNetworksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeClientVpnTargetNetworks", params, optFns, c.addOperationDescribeClientVpnTargetNetworksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClientVpnTargetNetworksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClientVpnTargetNetworksInput struct {

	// The ID of the Client VPN endpoint.
	//
	// This member is required.
	ClientVpnEndpointId *string

	// The IDs of the target network associations.
	AssociationIds []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters. Filter names and values are case-sensitive.
	//
	//   - association-id - The ID of the association.
	//
	//   - target-network-id - The ID of the subnet specified as the target network.
	//
	//   - vpc-id - The ID of the VPC in which the target network is located.
	Filters []types.Filter

	// The maximum number of results to return for the request in a single page. The
	// remaining results can be seen by sending another request with the nextToken
	// value.
	MaxResults *int32

	// The token to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeClientVpnTargetNetworksOutput struct {

	// Information about the associated target networks.
	ClientVpnTargetNetworks []types.TargetNetwork

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeClientVpnTargetNetworksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeClientVpnTargetNetworks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeClientVpnTargetNetworks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeClientVpnTargetNetworks"); err != nil {
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
	if err = addOpDescribeClientVpnTargetNetworksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClientVpnTargetNetworks(options.Region), middleware.Before); err != nil {
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

// DescribeClientVpnTargetNetworksAPIClient is a client that implements the
// DescribeClientVpnTargetNetworks operation.
type DescribeClientVpnTargetNetworksAPIClient interface {
	DescribeClientVpnTargetNetworks(context.Context, *DescribeClientVpnTargetNetworksInput, ...func(*Options)) (*DescribeClientVpnTargetNetworksOutput, error)
}

var _ DescribeClientVpnTargetNetworksAPIClient = (*Client)(nil)

// DescribeClientVpnTargetNetworksPaginatorOptions is the paginator options for
// DescribeClientVpnTargetNetworks
type DescribeClientVpnTargetNetworksPaginatorOptions struct {
	// The maximum number of results to return for the request in a single page. The
	// remaining results can be seen by sending another request with the nextToken
	// value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeClientVpnTargetNetworksPaginator is a paginator for
// DescribeClientVpnTargetNetworks
type DescribeClientVpnTargetNetworksPaginator struct {
	options   DescribeClientVpnTargetNetworksPaginatorOptions
	client    DescribeClientVpnTargetNetworksAPIClient
	params    *DescribeClientVpnTargetNetworksInput
	nextToken *string
	firstPage bool
}

// NewDescribeClientVpnTargetNetworksPaginator returns a new
// DescribeClientVpnTargetNetworksPaginator
func NewDescribeClientVpnTargetNetworksPaginator(client DescribeClientVpnTargetNetworksAPIClient, params *DescribeClientVpnTargetNetworksInput, optFns ...func(*DescribeClientVpnTargetNetworksPaginatorOptions)) *DescribeClientVpnTargetNetworksPaginator {
	if params == nil {
		params = &DescribeClientVpnTargetNetworksInput{}
	}

	options := DescribeClientVpnTargetNetworksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeClientVpnTargetNetworksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeClientVpnTargetNetworksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeClientVpnTargetNetworks page.
func (p *DescribeClientVpnTargetNetworksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeClientVpnTargetNetworksOutput, error) {
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

	result, err := p.client.DescribeClientVpnTargetNetworks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeClientVpnTargetNetworks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeClientVpnTargetNetworks",
	}
}
