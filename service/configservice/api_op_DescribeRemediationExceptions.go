// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the details of one or more remediation exceptions. A detailed view of a
// remediation exception for a set of resources that includes an explanation of an
// exception and the time when the exception will be deleted. When you specify the
// limit and the next token, you receive a paginated response.
//
// Config generates a remediation exception when a problem occurs executing a
// remediation action to a specific resource. Remediation exceptions blocks
// auto-remediation until the exception is cleared.
//
// When you specify the limit and the next token, you receive a paginated
// response.
//
// Limit and next token are not applicable if you request resources in batch. It
// is only applicable, when you request all resources.
func (c *Client) DescribeRemediationExceptions(ctx context.Context, params *DescribeRemediationExceptionsInput, optFns ...func(*Options)) (*DescribeRemediationExceptionsOutput, error) {
	if params == nil {
		params = &DescribeRemediationExceptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRemediationExceptions", params, optFns, c.addOperationDescribeRemediationExceptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRemediationExceptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRemediationExceptionsInput struct {

	// The name of the Config rule.
	//
	// This member is required.
	ConfigRuleName *string

	// The maximum number of RemediationExceptionResourceKey returned on each page.
	// The default is 25. If you specify 0, Config uses the default.
	Limit int32

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string

	// An exception list of resource exception keys to be processed with the current
	// request. Config adds exception for each resource key. For example, Config adds 3
	// exceptions for 3 resource keys.
	ResourceKeys []types.RemediationExceptionResourceKey

	noSmithyDocumentSerde
}

type DescribeRemediationExceptionsOutput struct {

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string

	// Returns a list of remediation exception objects.
	RemediationExceptions []types.RemediationException

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRemediationExceptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeRemediationExceptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeRemediationExceptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeRemediationExceptions"); err != nil {
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
	if err = addOpDescribeRemediationExceptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRemediationExceptions(options.Region), middleware.Before); err != nil {
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

// DescribeRemediationExceptionsAPIClient is a client that implements the
// DescribeRemediationExceptions operation.
type DescribeRemediationExceptionsAPIClient interface {
	DescribeRemediationExceptions(context.Context, *DescribeRemediationExceptionsInput, ...func(*Options)) (*DescribeRemediationExceptionsOutput, error)
}

var _ DescribeRemediationExceptionsAPIClient = (*Client)(nil)

// DescribeRemediationExceptionsPaginatorOptions is the paginator options for
// DescribeRemediationExceptions
type DescribeRemediationExceptionsPaginatorOptions struct {
	// The maximum number of RemediationExceptionResourceKey returned on each page.
	// The default is 25. If you specify 0, Config uses the default.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeRemediationExceptionsPaginator is a paginator for
// DescribeRemediationExceptions
type DescribeRemediationExceptionsPaginator struct {
	options   DescribeRemediationExceptionsPaginatorOptions
	client    DescribeRemediationExceptionsAPIClient
	params    *DescribeRemediationExceptionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeRemediationExceptionsPaginator returns a new
// DescribeRemediationExceptionsPaginator
func NewDescribeRemediationExceptionsPaginator(client DescribeRemediationExceptionsAPIClient, params *DescribeRemediationExceptionsInput, optFns ...func(*DescribeRemediationExceptionsPaginatorOptions)) *DescribeRemediationExceptionsPaginator {
	if params == nil {
		params = &DescribeRemediationExceptionsInput{}
	}

	options := DescribeRemediationExceptionsPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeRemediationExceptionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeRemediationExceptionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeRemediationExceptions page.
func (p *DescribeRemediationExceptionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeRemediationExceptionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.DescribeRemediationExceptions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeRemediationExceptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeRemediationExceptions",
	}
}
