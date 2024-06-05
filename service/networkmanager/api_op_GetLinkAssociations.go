// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the link associations for a device or a link. Either the device ID or the
// link ID must be specified.
func (c *Client) GetLinkAssociations(ctx context.Context, params *GetLinkAssociationsInput, optFns ...func(*Options)) (*GetLinkAssociationsOutput, error) {
	if params == nil {
		params = &GetLinkAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLinkAssociations", params, optFns, c.addOperationGetLinkAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLinkAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLinkAssociationsInput struct {

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// The ID of the device.
	DeviceId *string

	// The ID of the link.
	LinkId *string

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetLinkAssociationsOutput struct {

	// The link associations.
	LinkAssociations []types.LinkAssociation

	// The token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLinkAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetLinkAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLinkAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetLinkAssociations"); err != nil {
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
	if err = addOpGetLinkAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLinkAssociations(options.Region), middleware.Before); err != nil {
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

// GetLinkAssociationsAPIClient is a client that implements the
// GetLinkAssociations operation.
type GetLinkAssociationsAPIClient interface {
	GetLinkAssociations(context.Context, *GetLinkAssociationsInput, ...func(*Options)) (*GetLinkAssociationsOutput, error)
}

var _ GetLinkAssociationsAPIClient = (*Client)(nil)

// GetLinkAssociationsPaginatorOptions is the paginator options for
// GetLinkAssociations
type GetLinkAssociationsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetLinkAssociationsPaginator is a paginator for GetLinkAssociations
type GetLinkAssociationsPaginator struct {
	options   GetLinkAssociationsPaginatorOptions
	client    GetLinkAssociationsAPIClient
	params    *GetLinkAssociationsInput
	nextToken *string
	firstPage bool
}

// NewGetLinkAssociationsPaginator returns a new GetLinkAssociationsPaginator
func NewGetLinkAssociationsPaginator(client GetLinkAssociationsAPIClient, params *GetLinkAssociationsInput, optFns ...func(*GetLinkAssociationsPaginatorOptions)) *GetLinkAssociationsPaginator {
	if params == nil {
		params = &GetLinkAssociationsInput{}
	}

	options := GetLinkAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetLinkAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetLinkAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetLinkAssociations page.
func (p *GetLinkAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetLinkAssociationsOutput, error) {
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

	result, err := p.client.GetLinkAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetLinkAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetLinkAssociations",
	}
}
