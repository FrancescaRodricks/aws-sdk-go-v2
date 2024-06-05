// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the replication configuration for a specific file system. If a file
// system is not specified, all of the replication configurations for the Amazon
// Web Services account in an Amazon Web Services Region are retrieved.
func (c *Client) DescribeReplicationConfigurations(ctx context.Context, params *DescribeReplicationConfigurationsInput, optFns ...func(*Options)) (*DescribeReplicationConfigurationsOutput, error) {
	if params == nil {
		params = &DescribeReplicationConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReplicationConfigurations", params, optFns, c.addOperationDescribeReplicationConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReplicationConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeReplicationConfigurationsInput struct {

	// You can retrieve the replication configuration for a specific file system by
	// providing its file system ID.
	FileSystemId *string

	// (Optional) To limit the number of objects returned in a response, you can
	// specify the MaxItems parameter. The default value is 100.
	MaxResults *int32

	// NextToken is present if the response is paginated. You can use NextToken in a
	// subsequent request to fetch the next page of output.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeReplicationConfigurationsOutput struct {

	// You can use the NextToken from the previous response in a subsequent request to
	// fetch the additional descriptions.
	NextToken *string

	// The collection of replication configurations that is returned.
	Replications []types.ReplicationConfigurationDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReplicationConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeReplicationConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeReplicationConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeReplicationConfigurations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReplicationConfigurations(options.Region), middleware.Before); err != nil {
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

// DescribeReplicationConfigurationsAPIClient is a client that implements the
// DescribeReplicationConfigurations operation.
type DescribeReplicationConfigurationsAPIClient interface {
	DescribeReplicationConfigurations(context.Context, *DescribeReplicationConfigurationsInput, ...func(*Options)) (*DescribeReplicationConfigurationsOutput, error)
}

var _ DescribeReplicationConfigurationsAPIClient = (*Client)(nil)

// DescribeReplicationConfigurationsPaginatorOptions is the paginator options for
// DescribeReplicationConfigurations
type DescribeReplicationConfigurationsPaginatorOptions struct {
	// (Optional) To limit the number of objects returned in a response, you can
	// specify the MaxItems parameter. The default value is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReplicationConfigurationsPaginator is a paginator for
// DescribeReplicationConfigurations
type DescribeReplicationConfigurationsPaginator struct {
	options   DescribeReplicationConfigurationsPaginatorOptions
	client    DescribeReplicationConfigurationsAPIClient
	params    *DescribeReplicationConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeReplicationConfigurationsPaginator returns a new
// DescribeReplicationConfigurationsPaginator
func NewDescribeReplicationConfigurationsPaginator(client DescribeReplicationConfigurationsAPIClient, params *DescribeReplicationConfigurationsInput, optFns ...func(*DescribeReplicationConfigurationsPaginatorOptions)) *DescribeReplicationConfigurationsPaginator {
	if params == nil {
		params = &DescribeReplicationConfigurationsInput{}
	}

	options := DescribeReplicationConfigurationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReplicationConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReplicationConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeReplicationConfigurations page.
func (p *DescribeReplicationConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReplicationConfigurationsOutput, error) {
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

	result, err := p.client.DescribeReplicationConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeReplicationConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeReplicationConfigurations",
	}
}
