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

// Gets all the triggers associated with a job.
func (c *Client) GetTriggers(ctx context.Context, params *GetTriggersInput, optFns ...func(*Options)) (*GetTriggersOutput, error) {
	if params == nil {
		params = &GetTriggersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTriggers", params, optFns, c.addOperationGetTriggersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTriggersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTriggersInput struct {

	// The name of the job to retrieve triggers for. The trigger that can start this
	// job is returned, and if there is no such trigger, all triggers are returned.
	DependentJobName *string

	// The maximum size of the response.
	MaxResults *int32

	// A continuation token, if this is a continuation call.
	NextToken *string

	noSmithyDocumentSerde
}

type GetTriggersOutput struct {

	// A continuation token, if not all the requested triggers have yet been returned.
	NextToken *string

	// A list of triggers for the specified job.
	Triggers []types.Trigger

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTriggersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetTriggers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTriggers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetTriggers"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTriggers(options.Region), middleware.Before); err != nil {
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

// GetTriggersAPIClient is a client that implements the GetTriggers operation.
type GetTriggersAPIClient interface {
	GetTriggers(context.Context, *GetTriggersInput, ...func(*Options)) (*GetTriggersOutput, error)
}

var _ GetTriggersAPIClient = (*Client)(nil)

// GetTriggersPaginatorOptions is the paginator options for GetTriggers
type GetTriggersPaginatorOptions struct {
	// The maximum size of the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetTriggersPaginator is a paginator for GetTriggers
type GetTriggersPaginator struct {
	options   GetTriggersPaginatorOptions
	client    GetTriggersAPIClient
	params    *GetTriggersInput
	nextToken *string
	firstPage bool
}

// NewGetTriggersPaginator returns a new GetTriggersPaginator
func NewGetTriggersPaginator(client GetTriggersAPIClient, params *GetTriggersInput, optFns ...func(*GetTriggersPaginatorOptions)) *GetTriggersPaginator {
	if params == nil {
		params = &GetTriggersInput{}
	}

	options := GetTriggersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetTriggersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetTriggersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetTriggers page.
func (p *GetTriggersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetTriggersOutput, error) {
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

	result, err := p.client.GetTriggers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetTriggers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetTriggers",
	}
}
