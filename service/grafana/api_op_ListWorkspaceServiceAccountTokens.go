// Code generated by smithy-go-codegen DO NOT EDIT.

package grafana

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/grafana/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of tokens for a workspace service account.
//
// This does not return the key for each token. You cannot access keys after they
// are created. To create a new key, delete the token and recreate it.
//
// Service accounts are only available for workspaces that are compatible with
// Grafana version 9 and above.
func (c *Client) ListWorkspaceServiceAccountTokens(ctx context.Context, params *ListWorkspaceServiceAccountTokensInput, optFns ...func(*Options)) (*ListWorkspaceServiceAccountTokensOutput, error) {
	if params == nil {
		params = &ListWorkspaceServiceAccountTokensInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkspaceServiceAccountTokens", params, optFns, c.addOperationListWorkspaceServiceAccountTokensMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkspaceServiceAccountTokensOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkspaceServiceAccountTokensInput struct {

	// The ID of the service account for which to return tokens.
	//
	// This member is required.
	ServiceAccountId *string

	// The ID of the workspace for which to return tokens.
	//
	// This member is required.
	WorkspaceId *string

	// The maximum number of tokens to include in the results.
	MaxResults *int32

	// The token for the next set of service accounts to return. (You receive this
	// token from a previous ListWorkspaceServiceAccountTokens operation.)
	NextToken *string

	noSmithyDocumentSerde
}

type ListWorkspaceServiceAccountTokensOutput struct {

	// The ID of the service account where the tokens reside.
	//
	// This member is required.
	ServiceAccountId *string

	// An array of structures containing information about the tokens.
	//
	// This member is required.
	ServiceAccountTokens []types.ServiceAccountTokenSummary

	// The ID of the workspace where the tokens reside.
	//
	// This member is required.
	WorkspaceId *string

	// The token to use when requesting the next set of service accounts.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkspaceServiceAccountTokensMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWorkspaceServiceAccountTokens{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWorkspaceServiceAccountTokens{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListWorkspaceServiceAccountTokens"); err != nil {
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
	if err = addOpListWorkspaceServiceAccountTokensValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkspaceServiceAccountTokens(options.Region), middleware.Before); err != nil {
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

// ListWorkspaceServiceAccountTokensAPIClient is a client that implements the
// ListWorkspaceServiceAccountTokens operation.
type ListWorkspaceServiceAccountTokensAPIClient interface {
	ListWorkspaceServiceAccountTokens(context.Context, *ListWorkspaceServiceAccountTokensInput, ...func(*Options)) (*ListWorkspaceServiceAccountTokensOutput, error)
}

var _ ListWorkspaceServiceAccountTokensAPIClient = (*Client)(nil)

// ListWorkspaceServiceAccountTokensPaginatorOptions is the paginator options for
// ListWorkspaceServiceAccountTokens
type ListWorkspaceServiceAccountTokensPaginatorOptions struct {
	// The maximum number of tokens to include in the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkspaceServiceAccountTokensPaginator is a paginator for
// ListWorkspaceServiceAccountTokens
type ListWorkspaceServiceAccountTokensPaginator struct {
	options   ListWorkspaceServiceAccountTokensPaginatorOptions
	client    ListWorkspaceServiceAccountTokensAPIClient
	params    *ListWorkspaceServiceAccountTokensInput
	nextToken *string
	firstPage bool
}

// NewListWorkspaceServiceAccountTokensPaginator returns a new
// ListWorkspaceServiceAccountTokensPaginator
func NewListWorkspaceServiceAccountTokensPaginator(client ListWorkspaceServiceAccountTokensAPIClient, params *ListWorkspaceServiceAccountTokensInput, optFns ...func(*ListWorkspaceServiceAccountTokensPaginatorOptions)) *ListWorkspaceServiceAccountTokensPaginator {
	if params == nil {
		params = &ListWorkspaceServiceAccountTokensInput{}
	}

	options := ListWorkspaceServiceAccountTokensPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkspaceServiceAccountTokensPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkspaceServiceAccountTokensPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorkspaceServiceAccountTokens page.
func (p *ListWorkspaceServiceAccountTokensPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkspaceServiceAccountTokensOutput, error) {
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

	result, err := p.client.ListWorkspaceServiceAccountTokens(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWorkspaceServiceAccountTokens(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListWorkspaceServiceAccountTokens",
	}
}
