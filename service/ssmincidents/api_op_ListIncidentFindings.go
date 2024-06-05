// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmincidents

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssmincidents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of the IDs of findings, plus their last modified times, that
// have been identified for a specified incident. A finding represents a recent
// application environment change made by an CloudFormation stack creation or
// update or an CodeDeploy deployment that can be investigated as a potential cause
// of the incident.
func (c *Client) ListIncidentFindings(ctx context.Context, params *ListIncidentFindingsInput, optFns ...func(*Options)) (*ListIncidentFindingsOutput, error) {
	if params == nil {
		params = &ListIncidentFindingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIncidentFindings", params, optFns, c.addOperationListIncidentFindingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIncidentFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIncidentFindingsInput struct {

	// The Amazon Resource Name (ARN) of the incident for which you want to view
	// associated findings.
	//
	// This member is required.
	IncidentRecordArn *string

	// The maximum number of findings to retrieve per call.
	MaxResults *int32

	// The pagination token for the next set of items to return. (You received this
	// token from a previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type ListIncidentFindingsOutput struct {

	// A list of findings that represent deployments that might be the potential cause
	// of the incident.
	//
	// This member is required.
	Findings []types.FindingSummary

	// The pagination token to use when requesting the next set of items. If there are
	// no additional items to return, the string is null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListIncidentFindingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListIncidentFindings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListIncidentFindings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListIncidentFindings"); err != nil {
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
	if err = addOpListIncidentFindingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListIncidentFindings(options.Region), middleware.Before); err != nil {
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

// ListIncidentFindingsAPIClient is a client that implements the
// ListIncidentFindings operation.
type ListIncidentFindingsAPIClient interface {
	ListIncidentFindings(context.Context, *ListIncidentFindingsInput, ...func(*Options)) (*ListIncidentFindingsOutput, error)
}

var _ ListIncidentFindingsAPIClient = (*Client)(nil)

// ListIncidentFindingsPaginatorOptions is the paginator options for
// ListIncidentFindings
type ListIncidentFindingsPaginatorOptions struct {
	// The maximum number of findings to retrieve per call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListIncidentFindingsPaginator is a paginator for ListIncidentFindings
type ListIncidentFindingsPaginator struct {
	options   ListIncidentFindingsPaginatorOptions
	client    ListIncidentFindingsAPIClient
	params    *ListIncidentFindingsInput
	nextToken *string
	firstPage bool
}

// NewListIncidentFindingsPaginator returns a new ListIncidentFindingsPaginator
func NewListIncidentFindingsPaginator(client ListIncidentFindingsAPIClient, params *ListIncidentFindingsInput, optFns ...func(*ListIncidentFindingsPaginatorOptions)) *ListIncidentFindingsPaginator {
	if params == nil {
		params = &ListIncidentFindingsInput{}
	}

	options := ListIncidentFindingsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListIncidentFindingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListIncidentFindingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListIncidentFindings page.
func (p *ListIncidentFindingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListIncidentFindingsOutput, error) {
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

	result, err := p.client.ListIncidentFindings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListIncidentFindings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListIncidentFindings",
	}
}
