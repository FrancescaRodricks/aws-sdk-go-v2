// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the themes in the current Amazon Web Services account.
func (c *Client) ListThemes(ctx context.Context, params *ListThemesInput, optFns ...func(*Options)) (*ListThemesOutput, error) {
	if params == nil {
		params = &ListThemesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThemes", params, optFns, c.addOperationListThemesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThemesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThemesInput struct {

	// The ID of the Amazon Web Services account that contains the themes that you're
	// listing.
	//
	// This member is required.
	AwsAccountId *string

	// The maximum number of results to be returned per request.
	MaxResults int32

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// The type of themes that you want to list. Valid options include the
	// following:
	//
	// * ALL (default)- Display all existing themes.
	//
	// * CUSTOM - Display
	// only the themes created by people using Amazon QuickSight.
	//
	// * QUICKSIGHT -
	// Display only the starting themes defined by Amazon QuickSight.
	Type types.ThemeType

	noSmithyDocumentSerde
}

type ListThemesOutput struct {

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Information about the themes in the list.
	ThemeSummaryList []types.ThemeSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListThemesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListThemes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListThemes{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListThemesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListThemes(options.Region), middleware.Before); err != nil {
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
	return nil
}

// ListThemesAPIClient is a client that implements the ListThemes operation.
type ListThemesAPIClient interface {
	ListThemes(context.Context, *ListThemesInput, ...func(*Options)) (*ListThemesOutput, error)
}

var _ ListThemesAPIClient = (*Client)(nil)

// ListThemesPaginatorOptions is the paginator options for ListThemes
type ListThemesPaginatorOptions struct {
	// The maximum number of results to be returned per request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListThemesPaginator is a paginator for ListThemes
type ListThemesPaginator struct {
	options   ListThemesPaginatorOptions
	client    ListThemesAPIClient
	params    *ListThemesInput
	nextToken *string
	firstPage bool
}

// NewListThemesPaginator returns a new ListThemesPaginator
func NewListThemesPaginator(client ListThemesAPIClient, params *ListThemesInput, optFns ...func(*ListThemesPaginatorOptions)) *ListThemesPaginator {
	if params == nil {
		params = &ListThemesInput{}
	}

	options := ListThemesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListThemesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListThemesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListThemes page.
func (p *ListThemesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListThemesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListThemes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListThemes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "ListThemes",
	}
}