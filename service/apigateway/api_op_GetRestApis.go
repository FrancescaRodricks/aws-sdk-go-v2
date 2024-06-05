// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the RestApis resources for your collection.
func (c *Client) GetRestApis(ctx context.Context, params *GetRestApisInput, optFns ...func(*Options)) (*GetRestApisOutput, error) {
	if params == nil {
		params = &GetRestApisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRestApis", params, optFns, c.addOperationGetRestApisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRestApisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The GET request to list existing RestApis defined for your collection.
type GetRestApisInput struct {

	// The maximum number of returned results per page. The default value is 25 and
	// the maximum value is 500.
	Limit *int32

	// The current pagination position in the paged result set.
	Position *string

	noSmithyDocumentSerde
}

// Contains references to your APIs and links that guide you in how to interact
// with your collection. A collection offers a paginated view of your APIs.
type GetRestApisOutput struct {

	// The current page of elements from this collection.
	Items []types.RestApi

	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRestApisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRestApis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRestApis{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetRestApis"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRestApis(options.Region), middleware.Before); err != nil {
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
	if err = addAcceptHeader(stack); err != nil {
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

// GetRestApisAPIClient is a client that implements the GetRestApis operation.
type GetRestApisAPIClient interface {
	GetRestApis(context.Context, *GetRestApisInput, ...func(*Options)) (*GetRestApisOutput, error)
}

var _ GetRestApisAPIClient = (*Client)(nil)

// GetRestApisPaginatorOptions is the paginator options for GetRestApis
type GetRestApisPaginatorOptions struct {
	// The maximum number of returned results per page. The default value is 25 and
	// the maximum value is 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetRestApisPaginator is a paginator for GetRestApis
type GetRestApisPaginator struct {
	options   GetRestApisPaginatorOptions
	client    GetRestApisAPIClient
	params    *GetRestApisInput
	nextToken *string
	firstPage bool
}

// NewGetRestApisPaginator returns a new GetRestApisPaginator
func NewGetRestApisPaginator(client GetRestApisAPIClient, params *GetRestApisInput, optFns ...func(*GetRestApisPaginatorOptions)) *GetRestApisPaginator {
	if params == nil {
		params = &GetRestApisInput{}
	}

	options := GetRestApisPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetRestApisPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Position,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetRestApisPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetRestApis page.
func (p *GetRestApisPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetRestApisOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Position = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.GetRestApis(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Position

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetRestApis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetRestApis",
	}
}
