// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/detective/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Detective investigations lets you investigate IAM users and IAM roles using
// indicators of compromise. An indicator of compromise (IOC) is an artifact
// observed in or on a network, system, or environment that can (with a high level
// of confidence) identify malicious activity or a security incident.
// ListInvestigations lists all active Detective investigations.
func (c *Client) ListInvestigations(ctx context.Context, params *ListInvestigationsInput, optFns ...func(*Options)) (*ListInvestigationsOutput, error) {
	if params == nil {
		params = &ListInvestigationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInvestigations", params, optFns, c.addOperationListInvestigationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInvestigationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInvestigationsInput struct {

	// The Amazon Resource Name (ARN) of the behavior graph.
	//
	// This member is required.
	GraphArn *string

	// Filters the investigation results based on a criteria.
	FilterCriteria *types.FilterCriteria

	// Lists the maximum number of investigations in a page.
	MaxResults *int32

	// Lists if there are more results available. The value of nextToken is a unique
	// pagination token for each page. Repeat the call using the returned token to
	// retrieve the next page. Keep all other arguments unchanged.
	//
	// Each pagination token expires after 24 hours. Using an expired pagination token
	// will return a Validation Exception error.
	NextToken *string

	// Sorts the investigation results based on a criteria.
	SortCriteria *types.SortCriteria

	noSmithyDocumentSerde
}

type ListInvestigationsOutput struct {

	// Lists the summary of uncommon behavior or malicious activity which indicates a
	// compromise.
	InvestigationDetails []types.InvestigationDetail

	// Lists if there are more results available. The value of nextToken is a unique
	// pagination token for each page. Repeat the call using the returned token to
	// retrieve the next page. Keep all other arguments unchanged.
	//
	// Each pagination token expires after 24 hours.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInvestigationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListInvestigations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListInvestigations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListInvestigations"); err != nil {
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
	if err = addOpListInvestigationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInvestigations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListInvestigations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListInvestigations",
	}
}
