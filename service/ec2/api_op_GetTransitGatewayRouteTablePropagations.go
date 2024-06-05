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

// Gets information about the route table propagations for the specified transit
// gateway route table.
func (c *Client) GetTransitGatewayRouteTablePropagations(ctx context.Context, params *GetTransitGatewayRouteTablePropagationsInput, optFns ...func(*Options)) (*GetTransitGatewayRouteTablePropagationsOutput, error) {
	if params == nil {
		params = &GetTransitGatewayRouteTablePropagationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTransitGatewayRouteTablePropagations", params, optFns, c.addOperationGetTransitGatewayRouteTablePropagationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTransitGatewayRouteTablePropagationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTransitGatewayRouteTablePropagationsInput struct {

	// The ID of the transit gateway route table.
	//
	// This member is required.
	TransitGatewayRouteTableId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters. The possible values are:
	//
	//   - resource-id - The ID of the resource.
	//
	//   - resource-type - The resource type. Valid values are vpc | vpn |
	//   direct-connect-gateway | peering | connect .
	//
	//   - transit-gateway-attachment-id - The ID of the attachment.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetTransitGatewayRouteTablePropagationsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the route table propagations.
	TransitGatewayRouteTablePropagations []types.TransitGatewayRouteTablePropagation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTransitGatewayRouteTablePropagationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetTransitGatewayRouteTablePropagations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetTransitGatewayRouteTablePropagations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetTransitGatewayRouteTablePropagations"); err != nil {
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
	if err = addOpGetTransitGatewayRouteTablePropagationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTransitGatewayRouteTablePropagations(options.Region), middleware.Before); err != nil {
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

// GetTransitGatewayRouteTablePropagationsAPIClient is a client that implements
// the GetTransitGatewayRouteTablePropagations operation.
type GetTransitGatewayRouteTablePropagationsAPIClient interface {
	GetTransitGatewayRouteTablePropagations(context.Context, *GetTransitGatewayRouteTablePropagationsInput, ...func(*Options)) (*GetTransitGatewayRouteTablePropagationsOutput, error)
}

var _ GetTransitGatewayRouteTablePropagationsAPIClient = (*Client)(nil)

// GetTransitGatewayRouteTablePropagationsPaginatorOptions is the paginator
// options for GetTransitGatewayRouteTablePropagations
type GetTransitGatewayRouteTablePropagationsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetTransitGatewayRouteTablePropagationsPaginator is a paginator for
// GetTransitGatewayRouteTablePropagations
type GetTransitGatewayRouteTablePropagationsPaginator struct {
	options   GetTransitGatewayRouteTablePropagationsPaginatorOptions
	client    GetTransitGatewayRouteTablePropagationsAPIClient
	params    *GetTransitGatewayRouteTablePropagationsInput
	nextToken *string
	firstPage bool
}

// NewGetTransitGatewayRouteTablePropagationsPaginator returns a new
// GetTransitGatewayRouteTablePropagationsPaginator
func NewGetTransitGatewayRouteTablePropagationsPaginator(client GetTransitGatewayRouteTablePropagationsAPIClient, params *GetTransitGatewayRouteTablePropagationsInput, optFns ...func(*GetTransitGatewayRouteTablePropagationsPaginatorOptions)) *GetTransitGatewayRouteTablePropagationsPaginator {
	if params == nil {
		params = &GetTransitGatewayRouteTablePropagationsInput{}
	}

	options := GetTransitGatewayRouteTablePropagationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetTransitGatewayRouteTablePropagationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetTransitGatewayRouteTablePropagationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetTransitGatewayRouteTablePropagations page.
func (p *GetTransitGatewayRouteTablePropagationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetTransitGatewayRouteTablePropagationsOutput, error) {
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

	result, err := p.client.GetTransitGatewayRouteTablePropagations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetTransitGatewayRouteTablePropagations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetTransitGatewayRouteTablePropagations",
	}
}
