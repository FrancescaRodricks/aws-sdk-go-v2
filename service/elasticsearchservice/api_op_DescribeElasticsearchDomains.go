// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns domain configuration information about the specified Elasticsearch
// domains, including the domain ID, domain endpoint, and domain ARN.
func (c *Client) DescribeElasticsearchDomains(ctx context.Context, params *DescribeElasticsearchDomainsInput, optFns ...func(*Options)) (*DescribeElasticsearchDomainsOutput, error) {
	if params == nil {
		params = &DescribeElasticsearchDomainsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeElasticsearchDomains", params, optFns, c.addOperationDescribeElasticsearchDomainsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeElasticsearchDomainsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeElasticsearchDomains operation. By default, the API returns the
// status of all Elasticsearch domains.
type DescribeElasticsearchDomainsInput struct {

	// The Elasticsearch domains for which you want information.
	//
	// This member is required.
	DomainNames []string

	noSmithyDocumentSerde
}

// The result of a DescribeElasticsearchDomains request. Contains the status of
// the specified domains or all domains owned by the account.
type DescribeElasticsearchDomainsOutput struct {

	// The status of the domains requested in the DescribeElasticsearchDomains request.
	//
	// This member is required.
	DomainStatusList []types.ElasticsearchDomainStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeElasticsearchDomainsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeElasticsearchDomains{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeElasticsearchDomains{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeElasticsearchDomains"); err != nil {
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
	if err = addOpDescribeElasticsearchDomainsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeElasticsearchDomains(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeElasticsearchDomains(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeElasticsearchDomains",
	}
}
