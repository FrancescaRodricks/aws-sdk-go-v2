// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the versions for an add-on.
//
// Information such as the Kubernetes versions that you can use the add-on with,
// the owner , publisher , and the type of the add-on are returned.
func (c *Client) DescribeAddonVersions(ctx context.Context, params *DescribeAddonVersionsInput, optFns ...func(*Options)) (*DescribeAddonVersionsOutput, error) {
	if params == nil {
		params = &DescribeAddonVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAddonVersions", params, optFns, c.addOperationDescribeAddonVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAddonVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAddonVersionsInput struct {

	// The name of the add-on. The name must match one of the names returned by [ListAddons]
	// ListAddons .
	//
	// [ListAddons]: https://docs.aws.amazon.com/eks/latest/APIReference/API_ListAddons.html
	AddonName *string

	// The Kubernetes versions that you can use the add-on with.
	KubernetesVersion *string

	// The maximum number of results, returned in paginated output. You receive
	// maxResults in a single page, along with a nextToken response element. You can
	// see the remaining results of the initial request by sending another request with
	// the returned nextToken value. This value can be between 1 and 100. If you don't
	// use this parameter, 100 results and a nextToken value, if applicable, are
	// returned.
	MaxResults *int32

	// The nextToken value returned from a previous paginated request, where maxResults
	// was used and the results exceeded the value of that parameter. Pagination
	// continues from the end of the previous results that returned the nextToken
	// value. This value is null when there are no more results to return.
	//
	// This token should be treated as an opaque identifier that is used only to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string

	// The owner of the add-on. For valid owners , don't specify a value for this
	// property.
	Owners []string

	// The publisher of the add-on. For valid publishers , don't specify a value for
	// this property.
	Publishers []string

	// The type of the add-on. For valid types , don't specify a value for this
	// property.
	Types []string

	noSmithyDocumentSerde
}

type DescribeAddonVersionsOutput struct {

	// The list of available versions with Kubernetes version compatibility and other
	// properties.
	Addons []types.AddonInfo

	// The nextToken value to include in a future DescribeAddonVersions request. When
	// the results of a DescribeAddonVersions request exceed maxResults , you can use
	// this value to retrieve the next page of results. This value is null when there
	// are no more results to return.
	//
	// This token should be treated as an opaque identifier that is used only to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAddonVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAddonVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAddonVersions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAddonVersions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAddonVersions(options.Region), middleware.Before); err != nil {
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

// DescribeAddonVersionsAPIClient is a client that implements the
// DescribeAddonVersions operation.
type DescribeAddonVersionsAPIClient interface {
	DescribeAddonVersions(context.Context, *DescribeAddonVersionsInput, ...func(*Options)) (*DescribeAddonVersionsOutput, error)
}

var _ DescribeAddonVersionsAPIClient = (*Client)(nil)

// DescribeAddonVersionsPaginatorOptions is the paginator options for
// DescribeAddonVersions
type DescribeAddonVersionsPaginatorOptions struct {
	// The maximum number of results, returned in paginated output. You receive
	// maxResults in a single page, along with a nextToken response element. You can
	// see the remaining results of the initial request by sending another request with
	// the returned nextToken value. This value can be between 1 and 100. If you don't
	// use this parameter, 100 results and a nextToken value, if applicable, are
	// returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAddonVersionsPaginator is a paginator for DescribeAddonVersions
type DescribeAddonVersionsPaginator struct {
	options   DescribeAddonVersionsPaginatorOptions
	client    DescribeAddonVersionsAPIClient
	params    *DescribeAddonVersionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeAddonVersionsPaginator returns a new DescribeAddonVersionsPaginator
func NewDescribeAddonVersionsPaginator(client DescribeAddonVersionsAPIClient, params *DescribeAddonVersionsInput, optFns ...func(*DescribeAddonVersionsPaginatorOptions)) *DescribeAddonVersionsPaginator {
	if params == nil {
		params = &DescribeAddonVersionsInput{}
	}

	options := DescribeAddonVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeAddonVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAddonVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeAddonVersions page.
func (p *DescribeAddonVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAddonVersionsOutput, error) {
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

	result, err := p.client.DescribeAddonVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeAddonVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAddonVersions",
	}
}
