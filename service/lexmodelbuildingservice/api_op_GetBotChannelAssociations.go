// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Returns a list of all of the channels associated with the specified bot.
//
// The GetBotChannelAssociations operation requires permissions for the
// lex:GetBotChannelAssociations action.
func (c *Client) GetBotChannelAssociations(ctx context.Context, params *GetBotChannelAssociationsInput, optFns ...func(*Options)) (*GetBotChannelAssociationsOutput, error) {
	if params == nil {
		params = &GetBotChannelAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBotChannelAssociations", params, optFns, c.addOperationGetBotChannelAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBotChannelAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBotChannelAssociationsInput struct {

	// An alias pointing to the specific version of the Amazon Lex bot to which this
	// association is being made.
	//
	// This member is required.
	BotAlias *string

	// The name of the Amazon Lex bot in the association.
	//
	// This member is required.
	BotName *string

	// The maximum number of associations to return in the response. The default is
	// 50.
	MaxResults *int32

	// Substring to match in channel association names. An association will be
	// returned if any part of its name matches the substring. For example, "xyz"
	// matches both "xyzabc" and "abcxyz." To return all bot channel associations, use
	// a hyphen ("-") as the nameContains parameter.
	NameContains *string

	// A pagination token for fetching the next page of associations. If the response
	// to this call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of associations, specify the pagination token
	// in the next request.
	NextToken *string

	noSmithyDocumentSerde
}

type GetBotChannelAssociationsOutput struct {

	// An array of objects, one for each association, that provides information about
	// the Amazon Lex bot and its association with the channel.
	BotChannelAssociations []types.BotChannelAssociation

	// A pagination token that fetches the next page of associations. If the response
	// to this call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of associations, specify the pagination token
	// in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBotChannelAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBotChannelAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBotChannelAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetBotChannelAssociations"); err != nil {
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
	if err = addOpGetBotChannelAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBotChannelAssociations(options.Region), middleware.Before); err != nil {
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

// GetBotChannelAssociationsAPIClient is a client that implements the
// GetBotChannelAssociations operation.
type GetBotChannelAssociationsAPIClient interface {
	GetBotChannelAssociations(context.Context, *GetBotChannelAssociationsInput, ...func(*Options)) (*GetBotChannelAssociationsOutput, error)
}

var _ GetBotChannelAssociationsAPIClient = (*Client)(nil)

// GetBotChannelAssociationsPaginatorOptions is the paginator options for
// GetBotChannelAssociations
type GetBotChannelAssociationsPaginatorOptions struct {
	// The maximum number of associations to return in the response. The default is
	// 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetBotChannelAssociationsPaginator is a paginator for GetBotChannelAssociations
type GetBotChannelAssociationsPaginator struct {
	options   GetBotChannelAssociationsPaginatorOptions
	client    GetBotChannelAssociationsAPIClient
	params    *GetBotChannelAssociationsInput
	nextToken *string
	firstPage bool
}

// NewGetBotChannelAssociationsPaginator returns a new
// GetBotChannelAssociationsPaginator
func NewGetBotChannelAssociationsPaginator(client GetBotChannelAssociationsAPIClient, params *GetBotChannelAssociationsInput, optFns ...func(*GetBotChannelAssociationsPaginatorOptions)) *GetBotChannelAssociationsPaginator {
	if params == nil {
		params = &GetBotChannelAssociationsInput{}
	}

	options := GetBotChannelAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetBotChannelAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetBotChannelAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetBotChannelAssociations page.
func (p *GetBotChannelAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetBotChannelAssociationsOutput, error) {
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

	result, err := p.client.GetBotChannelAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetBotChannelAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetBotChannelAssociations",
	}
}
