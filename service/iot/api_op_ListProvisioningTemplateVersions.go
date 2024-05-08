// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A list of provisioning template versions.
//
// Requires permission to access the [ListProvisioningTemplateVersions] action.
//
// [ListProvisioningTemplateVersions]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) ListProvisioningTemplateVersions(ctx context.Context, params *ListProvisioningTemplateVersionsInput, optFns ...func(*Options)) (*ListProvisioningTemplateVersionsOutput, error) {
	if params == nil {
		params = &ListProvisioningTemplateVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProvisioningTemplateVersions", params, optFns, c.addOperationListProvisioningTemplateVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProvisioningTemplateVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProvisioningTemplateVersionsInput struct {

	// The name of the provisioning template.
	//
	// This member is required.
	TemplateName *string

	// The maximum number of results to return at one time.
	MaxResults *int32

	// A token to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListProvisioningTemplateVersionsOutput struct {

	// A token to retrieve the next set of results.
	NextToken *string

	// The list of provisioning template versions.
	Versions []types.ProvisioningTemplateVersionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProvisioningTemplateVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProvisioningTemplateVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProvisioningTemplateVersions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListProvisioningTemplateVersions"); err != nil {
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
	if err = addOpListProvisioningTemplateVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProvisioningTemplateVersions(options.Region), middleware.Before); err != nil {
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

// ListProvisioningTemplateVersionsAPIClient is a client that implements the
// ListProvisioningTemplateVersions operation.
type ListProvisioningTemplateVersionsAPIClient interface {
	ListProvisioningTemplateVersions(context.Context, *ListProvisioningTemplateVersionsInput, ...func(*Options)) (*ListProvisioningTemplateVersionsOutput, error)
}

var _ ListProvisioningTemplateVersionsAPIClient = (*Client)(nil)

// ListProvisioningTemplateVersionsPaginatorOptions is the paginator options for
// ListProvisioningTemplateVersions
type ListProvisioningTemplateVersionsPaginatorOptions struct {
	// The maximum number of results to return at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProvisioningTemplateVersionsPaginator is a paginator for
// ListProvisioningTemplateVersions
type ListProvisioningTemplateVersionsPaginator struct {
	options   ListProvisioningTemplateVersionsPaginatorOptions
	client    ListProvisioningTemplateVersionsAPIClient
	params    *ListProvisioningTemplateVersionsInput
	nextToken *string
	firstPage bool
}

// NewListProvisioningTemplateVersionsPaginator returns a new
// ListProvisioningTemplateVersionsPaginator
func NewListProvisioningTemplateVersionsPaginator(client ListProvisioningTemplateVersionsAPIClient, params *ListProvisioningTemplateVersionsInput, optFns ...func(*ListProvisioningTemplateVersionsPaginatorOptions)) *ListProvisioningTemplateVersionsPaginator {
	if params == nil {
		params = &ListProvisioningTemplateVersionsInput{}
	}

	options := ListProvisioningTemplateVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProvisioningTemplateVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProvisioningTemplateVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProvisioningTemplateVersions page.
func (p *ListProvisioningTemplateVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProvisioningTemplateVersionsOutput, error) {
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

	result, err := p.client.ListProvisioningTemplateVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProvisioningTemplateVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListProvisioningTemplateVersions",
	}
}
