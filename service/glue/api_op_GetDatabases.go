// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves all databases defined in a given Data Catalog.
func (c *Client) GetDatabases(ctx context.Context, params *GetDatabasesInput, optFns ...func(*Options)) (*GetDatabasesOutput, error) {
	if params == nil {
		params = &GetDatabasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDatabases", params, optFns, c.addOperationGetDatabasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDatabasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDatabasesInput struct {

	// The ID of the Data Catalog from which to retrieve Databases . If none is
	// provided, the Amazon Web Services account ID is used by default.
	CatalogId *string

	// The maximum number of databases to return in one response.
	MaxResults *int32

	// A continuation token, if this is a continuation call.
	NextToken *string

	// Allows you to specify that you want to list the databases shared with your
	// account. The allowable values are FEDERATED , FOREIGN or ALL .
	//
	//   - If set to FEDERATED , will list the federated databases (referencing an
	//   external entity) shared with your account.
	//
	//   - If set to FOREIGN , will list the databases shared with your account.
	//
	//   - If set to ALL , will list the databases shared with your account, as well as
	//   the databases in yor local account.
	ResourceShareType types.ResourceShareType

	noSmithyDocumentSerde
}

type GetDatabasesOutput struct {

	// A list of Database objects from the specified catalog.
	//
	// This member is required.
	DatabaseList []types.Database

	// A continuation token for paginating the returned list of tokens, returned if
	// the current segment of the list is not the last.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDatabasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDatabases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDatabases{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDatabases"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDatabases(options.Region), middleware.Before); err != nil {
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

// GetDatabasesAPIClient is a client that implements the GetDatabases operation.
type GetDatabasesAPIClient interface {
	GetDatabases(context.Context, *GetDatabasesInput, ...func(*Options)) (*GetDatabasesOutput, error)
}

var _ GetDatabasesAPIClient = (*Client)(nil)

// GetDatabasesPaginatorOptions is the paginator options for GetDatabases
type GetDatabasesPaginatorOptions struct {
	// The maximum number of databases to return in one response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetDatabasesPaginator is a paginator for GetDatabases
type GetDatabasesPaginator struct {
	options   GetDatabasesPaginatorOptions
	client    GetDatabasesAPIClient
	params    *GetDatabasesInput
	nextToken *string
	firstPage bool
}

// NewGetDatabasesPaginator returns a new GetDatabasesPaginator
func NewGetDatabasesPaginator(client GetDatabasesAPIClient, params *GetDatabasesInput, optFns ...func(*GetDatabasesPaginatorOptions)) *GetDatabasesPaginator {
	if params == nil {
		params = &GetDatabasesInput{}
	}

	options := GetDatabasesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetDatabasesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetDatabasesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetDatabases page.
func (p *GetDatabasesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetDatabasesOutput, error) {
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

	result, err := p.client.GetDatabases(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetDatabases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDatabases",
	}
}
