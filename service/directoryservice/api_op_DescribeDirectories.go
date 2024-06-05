// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Obtains information about the directories that belong to this account.
//
// You can retrieve information about specific directories by passing the
// directory identifiers in the DirectoryIds parameter. Otherwise, all directories
// that belong to the current account are returned.
//
// This operation supports pagination with the use of the NextToken request and
// response parameters. If more results are available, the
// DescribeDirectoriesResult.NextToken member contains a token that you pass in the
// next call to DescribeDirectoriesto retrieve the next set of items.
//
// You can also specify a maximum number of return results with the Limit
// parameter.
func (c *Client) DescribeDirectories(ctx context.Context, params *DescribeDirectoriesInput, optFns ...func(*Options)) (*DescribeDirectoriesOutput, error) {
	if params == nil {
		params = &DescribeDirectoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDirectories", params, optFns, c.addOperationDescribeDirectoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDirectoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the DescribeDirectories operation.
type DescribeDirectoriesInput struct {

	// A list of identifiers of the directories for which to obtain the information.
	// If this member is null, all directories that belong to the current account are
	// returned.
	//
	// An empty list results in an InvalidParameterException being thrown.
	DirectoryIds []string

	// The maximum number of items to return. If this value is zero, the maximum
	// number of items is specified by the limitations of the operation.
	Limit *int32

	// The DescribeDirectoriesResult.NextToken value from a previous call to DescribeDirectories. Pass
	// null if this is the first call.
	NextToken *string

	noSmithyDocumentSerde
}

// Contains the results of the DescribeDirectories operation.
type DescribeDirectoriesOutput struct {

	// The list of DirectoryDescription objects that were retrieved.
	//
	// It is possible that this list contains less than the number of items specified
	// in the Limit member of the request. This occurs if there are less than the
	// requested number of items left to retrieve, or if the limitations of the
	// operation have been exceeded.
	DirectoryDescriptions []types.DirectoryDescription

	// If not null, more results are available. Pass this value for the NextToken
	// parameter in a subsequent call to DescribeDirectoriesto retrieve the next set of items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDirectoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDirectories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDirectories{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDirectories"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDirectories(options.Region), middleware.Before); err != nil {
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

// DescribeDirectoriesAPIClient is a client that implements the
// DescribeDirectories operation.
type DescribeDirectoriesAPIClient interface {
	DescribeDirectories(context.Context, *DescribeDirectoriesInput, ...func(*Options)) (*DescribeDirectoriesOutput, error)
}

var _ DescribeDirectoriesAPIClient = (*Client)(nil)

// DescribeDirectoriesPaginatorOptions is the paginator options for
// DescribeDirectories
type DescribeDirectoriesPaginatorOptions struct {
	// The maximum number of items to return. If this value is zero, the maximum
	// number of items is specified by the limitations of the operation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDirectoriesPaginator is a paginator for DescribeDirectories
type DescribeDirectoriesPaginator struct {
	options   DescribeDirectoriesPaginatorOptions
	client    DescribeDirectoriesAPIClient
	params    *DescribeDirectoriesInput
	nextToken *string
	firstPage bool
}

// NewDescribeDirectoriesPaginator returns a new DescribeDirectoriesPaginator
func NewDescribeDirectoriesPaginator(client DescribeDirectoriesAPIClient, params *DescribeDirectoriesInput, optFns ...func(*DescribeDirectoriesPaginatorOptions)) *DescribeDirectoriesPaginator {
	if params == nil {
		params = &DescribeDirectoriesInput{}
	}

	options := DescribeDirectoriesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDirectoriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDirectoriesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDirectories page.
func (p *DescribeDirectoriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDirectoriesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeDirectories(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeDirectories(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDirectories",
	}
}
