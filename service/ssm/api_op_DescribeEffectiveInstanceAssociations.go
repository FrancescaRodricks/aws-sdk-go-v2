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

// All associations for the managed nodes.
func (c *Client) DescribeEffectiveInstanceAssociations(ctx context.Context, params *DescribeEffectiveInstanceAssociationsInput, optFns ...func(*Options)) (*DescribeEffectiveInstanceAssociationsOutput, error) {
	if params == nil {
		params = &DescribeEffectiveInstanceAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEffectiveInstanceAssociations", params, optFns, c.addOperationDescribeEffectiveInstanceAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEffectiveInstanceAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEffectiveInstanceAssociationsInput struct {

	// The managed node ID for which you want to view all associations.
	//
	// This member is required.
	InstanceId *string

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeEffectiveInstanceAssociationsOutput struct {

	// The associations for the requested managed node.
	Associations []types.InstanceAssociation

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEffectiveInstanceAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEffectiveInstanceAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEffectiveInstanceAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEffectiveInstanceAssociations"); err != nil {
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
	if err = addOpDescribeEffectiveInstanceAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEffectiveInstanceAssociations(options.Region), middleware.Before); err != nil {
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

// DescribeEffectiveInstanceAssociationsAPIClient is a client that implements the
// DescribeEffectiveInstanceAssociations operation.
type DescribeEffectiveInstanceAssociationsAPIClient interface {
	DescribeEffectiveInstanceAssociations(context.Context, *DescribeEffectiveInstanceAssociationsInput, ...func(*Options)) (*DescribeEffectiveInstanceAssociationsOutput, error)
}

var _ DescribeEffectiveInstanceAssociationsAPIClient = (*Client)(nil)

// DescribeEffectiveInstanceAssociationsPaginatorOptions is the paginator options
// for DescribeEffectiveInstanceAssociations
type DescribeEffectiveInstanceAssociationsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEffectiveInstanceAssociationsPaginator is a paginator for
// DescribeEffectiveInstanceAssociations
type DescribeEffectiveInstanceAssociationsPaginator struct {
	options   DescribeEffectiveInstanceAssociationsPaginatorOptions
	client    DescribeEffectiveInstanceAssociationsAPIClient
	params    *DescribeEffectiveInstanceAssociationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeEffectiveInstanceAssociationsPaginator returns a new
// DescribeEffectiveInstanceAssociationsPaginator
func NewDescribeEffectiveInstanceAssociationsPaginator(client DescribeEffectiveInstanceAssociationsAPIClient, params *DescribeEffectiveInstanceAssociationsInput, optFns ...func(*DescribeEffectiveInstanceAssociationsPaginatorOptions)) *DescribeEffectiveInstanceAssociationsPaginator {
	if params == nil {
		params = &DescribeEffectiveInstanceAssociationsInput{}
	}

	options := DescribeEffectiveInstanceAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEffectiveInstanceAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEffectiveInstanceAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEffectiveInstanceAssociations page.
func (p *DescribeEffectiveInstanceAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEffectiveInstanceAssociationsOutput, error) {
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

	result, err := p.client.DescribeEffectiveInstanceAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeEffectiveInstanceAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEffectiveInstanceAssociations",
	}
}
