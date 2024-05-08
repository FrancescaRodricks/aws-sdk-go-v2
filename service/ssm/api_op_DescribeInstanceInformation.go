// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides information about one or more of your managed nodes, including the
// operating system platform, SSM Agent version, association status, and IP
// address. This operation does not return information for nodes that are either
// Stopped or Terminated.
//
// If you specify one or more node IDs, the operation returns information for
// those managed nodes. If you don't specify node IDs, it returns information for
// all your managed nodes. If you specify a node ID that isn't valid or a node that
// you don't own, you receive an error.
//
// The IamRole field returned for this API operation is the Identity and Access
// Management (IAM) role assigned to on-premises managed nodes. This operation does
// not return the IAM role for EC2 instances.
func (c *Client) DescribeInstanceInformation(ctx context.Context, params *DescribeInstanceInformationInput, optFns ...func(*Options)) (*DescribeInstanceInformationOutput, error) {
	if params == nil {
		params = &DescribeInstanceInformationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstanceInformation", params, optFns, c.addOperationDescribeInstanceInformationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstanceInformationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstanceInformationInput struct {

	// One or more filters. Use a filter to return a more specific list of managed
	// nodes. You can filter based on tags applied to your managed nodes. Tag filters
	// can't be combined with other filter types. Use this Filters data type instead
	// of InstanceInformationFilterList , which is deprecated.
	Filters []types.InstanceInformationStringFilter

	// This is a legacy method. We recommend that you don't use this method. Instead,
	// use the Filters data type. Filters enables you to return node information by
	// filtering based on tags applied to managed nodes.
	//
	// Attempting to use InstanceInformationFilterList and Filters leads to an
	// exception error.
	InstanceInformationFilterList []types.InstanceInformationFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	// The default value is 10 items.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeInstanceInformationOutput struct {

	// The managed node information list.
	InstanceInformationList []types.InstanceInformation

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInstanceInformationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeInstanceInformation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeInstanceInformation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeInstanceInformation"); err != nil {
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
	if err = addOpDescribeInstanceInformationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceInformation(options.Region), middleware.Before); err != nil {
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

// DescribeInstanceInformationAPIClient is a client that implements the
// DescribeInstanceInformation operation.
type DescribeInstanceInformationAPIClient interface {
	DescribeInstanceInformation(context.Context, *DescribeInstanceInformationInput, ...func(*Options)) (*DescribeInstanceInformationOutput, error)
}

var _ DescribeInstanceInformationAPIClient = (*Client)(nil)

// DescribeInstanceInformationPaginatorOptions is the paginator options for
// DescribeInstanceInformation
type DescribeInstanceInformationPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	// The default value is 10 items.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeInstanceInformationPaginator is a paginator for
// DescribeInstanceInformation
type DescribeInstanceInformationPaginator struct {
	options   DescribeInstanceInformationPaginatorOptions
	client    DescribeInstanceInformationAPIClient
	params    *DescribeInstanceInformationInput
	nextToken *string
	firstPage bool
}

// NewDescribeInstanceInformationPaginator returns a new
// DescribeInstanceInformationPaginator
func NewDescribeInstanceInformationPaginator(client DescribeInstanceInformationAPIClient, params *DescribeInstanceInformationInput, optFns ...func(*DescribeInstanceInformationPaginatorOptions)) *DescribeInstanceInformationPaginator {
	if params == nil {
		params = &DescribeInstanceInformationInput{}
	}

	options := DescribeInstanceInformationPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeInstanceInformationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeInstanceInformationPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeInstanceInformation page.
func (p *DescribeInstanceInformationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeInstanceInformationOutput, error) {
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

	result, err := p.client.DescribeInstanceInformation(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeInstanceInformation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeInstanceInformation",
	}
}
