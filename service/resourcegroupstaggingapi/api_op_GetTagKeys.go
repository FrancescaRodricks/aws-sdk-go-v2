// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroupstaggingapi

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns all tag keys currently in use in the specified Amazon Web Services
// Region for the calling account.
//
// This operation supports pagination, where the response can be sent in multiple
// pages. You should check the PaginationToken response parameter to determine if
// there are additional results available to return. Repeat the query, passing the
// PaginationToken response parameter value as an input to the next request until
// you recieve a null value. A null value for PaginationToken indicates that there
// are no more results waiting to be returned.
func (c *Client) GetTagKeys(ctx context.Context, params *GetTagKeysInput, optFns ...func(*Options)) (*GetTagKeysOutput, error) {
	if params == nil {
		params = &GetTagKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTagKeys", params, optFns, c.addOperationGetTagKeysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTagKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTagKeysInput struct {

	// Specifies a PaginationToken response value from a previous request to indicate
	// that you want the next page of results. Leave this parameter empty in your
	// initial request.
	PaginationToken *string

	noSmithyDocumentSerde
}

type GetTagKeysOutput struct {

	// A string that indicates that there is more data available than this response
	// contains. To receive the next part of the response, specify this response value
	// as the PaginationToken value in the request for the next page.
	PaginationToken *string

	// A list of all tag keys in the Amazon Web Services account.
	TagKeys []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTagKeysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetTagKeys{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTagKeys{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetTagKeys"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTagKeys(options.Region), middleware.Before); err != nil {
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

// GetTagKeysAPIClient is a client that implements the GetTagKeys operation.
type GetTagKeysAPIClient interface {
	GetTagKeys(context.Context, *GetTagKeysInput, ...func(*Options)) (*GetTagKeysOutput, error)
}

var _ GetTagKeysAPIClient = (*Client)(nil)

// GetTagKeysPaginatorOptions is the paginator options for GetTagKeys
type GetTagKeysPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetTagKeysPaginator is a paginator for GetTagKeys
type GetTagKeysPaginator struct {
	options   GetTagKeysPaginatorOptions
	client    GetTagKeysAPIClient
	params    *GetTagKeysInput
	nextToken *string
	firstPage bool
}

// NewGetTagKeysPaginator returns a new GetTagKeysPaginator
func NewGetTagKeysPaginator(client GetTagKeysAPIClient, params *GetTagKeysInput, optFns ...func(*GetTagKeysPaginatorOptions)) *GetTagKeysPaginator {
	if params == nil {
		params = &GetTagKeysInput{}
	}

	options := GetTagKeysPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetTagKeysPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.PaginationToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetTagKeysPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetTagKeys page.
func (p *GetTagKeysPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetTagKeysOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PaginationToken = p.nextToken

	result, err := p.client.GetTagKeys(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.PaginationToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetTagKeys(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetTagKeys",
	}
}
