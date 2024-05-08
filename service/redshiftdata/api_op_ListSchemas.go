// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftdata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the schemas in a database. A token is returned to page through the schema
// list. Depending on the authorization method, use one of the following
// combinations of request parameters:
//
//   - Secrets Manager - when connecting to a cluster, provide the secret-arn of a
//     secret stored in Secrets Manager which has username and password . The
//     specified secret contains credentials to connect to the database you specify.
//     When you are connecting to a cluster, you also supply the database name, If you
//     provide a cluster identifier ( dbClusterIdentifier ), it must match the
//     cluster identifier stored in the secret. When you are connecting to a serverless
//     workgroup, you also supply the database name.
//
//   - Temporary credentials - when connecting to your data warehouse, choose one
//     of the following options:
//
//   - When connecting to a serverless workgroup, specify the workgroup name and
//     database name. The database user name is derived from the IAM identity. For
//     example, arn:iam::123456789012:user:foo has the database user name IAM:foo .
//     Also, permission to call the redshift-serverless:GetCredentials operation is
//     required.
//
//   - When connecting to a cluster as an IAM identity, specify the cluster
//     identifier and the database name. The database user name is derived from the IAM
//     identity. For example, arn:iam::123456789012:user:foo has the database user
//     name IAM:foo . Also, permission to call the
//     redshift:GetClusterCredentialsWithIAM operation is required.
//
//   - When connecting to a cluster as a database user, specify the cluster
//     identifier, the database name, and the database user name. Also, permission to
//     call the redshift:GetClusterCredentials operation is required.
//
// For more information about the Amazon Redshift Data API and CLI usage examples,
// see [Using the Amazon Redshift Data API]in the Amazon Redshift Management Guide.
//
// [Using the Amazon Redshift Data API]: https://docs.aws.amazon.com/redshift/latest/mgmt/data-api.html
func (c *Client) ListSchemas(ctx context.Context, params *ListSchemasInput, optFns ...func(*Options)) (*ListSchemasOutput, error) {
	if params == nil {
		params = &ListSchemasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSchemas", params, optFns, c.addOperationListSchemasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSchemasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSchemasInput struct {

	// The name of the database that contains the schemas to list. If ConnectedDatabase
	// is not specified, this is also the database to connect to with your
	// authentication credentials.
	//
	// This member is required.
	Database *string

	// The cluster identifier. This parameter is required when connecting to a cluster
	// and authenticating using either Secrets Manager or temporary credentials.
	ClusterIdentifier *string

	// A database name. The connected database is specified when you connect with your
	// authentication credentials.
	ConnectedDatabase *string

	// The database user name. This parameter is required when connecting to a cluster
	// as a database user and authenticating using temporary credentials.
	DbUser *string

	// The maximum number of schemas to return in the response. If more schemas exist
	// than fit in one response, then NextToken is returned to page through the
	// results.
	MaxResults int32

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned NextToken value in the next
	// NextToken parameter and retrying the command. If the NextToken field is empty,
	// all response records have been retrieved for the request.
	NextToken *string

	// A pattern to filter results by schema name. Within a schema pattern, "%" means
	// match any substring of 0 or more characters and "_" means match any one
	// character. Only schema name entries matching the search pattern are returned.
	SchemaPattern *string

	// The name or ARN of the secret that enables access to the database. This
	// parameter is required when authenticating using Secrets Manager.
	SecretArn *string

	// The serverless workgroup name or Amazon Resource Name (ARN). This parameter is
	// required when connecting to a serverless workgroup and authenticating using
	// either Secrets Manager or temporary credentials.
	WorkgroupName *string

	noSmithyDocumentSerde
}

type ListSchemasOutput struct {

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned NextToken value in the next
	// NextToken parameter and retrying the command. If the NextToken field is empty,
	// all response records have been retrieved for the request.
	NextToken *string

	// The schemas that match the request pattern.
	Schemas []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSchemasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSchemas{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSchemas{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSchemas"); err != nil {
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
	if err = addOpListSchemasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSchemas(options.Region), middleware.Before); err != nil {
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

// ListSchemasAPIClient is a client that implements the ListSchemas operation.
type ListSchemasAPIClient interface {
	ListSchemas(context.Context, *ListSchemasInput, ...func(*Options)) (*ListSchemasOutput, error)
}

var _ ListSchemasAPIClient = (*Client)(nil)

// ListSchemasPaginatorOptions is the paginator options for ListSchemas
type ListSchemasPaginatorOptions struct {
	// The maximum number of schemas to return in the response. If more schemas exist
	// than fit in one response, then NextToken is returned to page through the
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSchemasPaginator is a paginator for ListSchemas
type ListSchemasPaginator struct {
	options   ListSchemasPaginatorOptions
	client    ListSchemasAPIClient
	params    *ListSchemasInput
	nextToken *string
	firstPage bool
}

// NewListSchemasPaginator returns a new ListSchemasPaginator
func NewListSchemasPaginator(client ListSchemasAPIClient, params *ListSchemasInput, optFns ...func(*ListSchemasPaginatorOptions)) *ListSchemasPaginator {
	if params == nil {
		params = &ListSchemasInput{}
	}

	options := ListSchemasPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSchemasPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSchemasPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSchemas page.
func (p *ListSchemasPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSchemasOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListSchemas(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSchemas(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSchemas",
	}
}
