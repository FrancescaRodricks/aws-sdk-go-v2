// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the expressions configured for the search domain. Can be limited to
// specific expressions by name. By default, shows all expressions and includes any
// pending changes to the configuration. Set the Deployed option to true to show
// the active configuration and exclude pending changes. For more information, see [Configuring Expressions]
// in the Amazon CloudSearch Developer Guide.
//
// [Configuring Expressions]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-expressions.html
func (c *Client) DescribeExpressions(ctx context.Context, params *DescribeExpressionsInput, optFns ...func(*Options)) (*DescribeExpressionsOutput, error) {
	if params == nil {
		params = &DescribeExpressionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeExpressions", params, optFns, c.addOperationDescribeExpressionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeExpressionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeDomains operation. Specifies the name of the
// domain you want to describe. To restrict the response to particular expressions,
// specify the names of the expressions you want to describe. To show the active
// configuration and exclude any pending changes, set the Deployed option to true .
type DescribeExpressionsInput struct {

	// The name of the domain you want to describe.
	//
	// This member is required.
	DomainName *string

	// Whether to display the deployed configuration ( true ) or include any pending
	// changes ( false ). Defaults to false .
	Deployed *bool

	// Limits the DescribeExpressions response to the specified expressions. If not specified, all
	// expressions are shown.
	ExpressionNames []string

	noSmithyDocumentSerde
}

// The result of a DescribeExpressions request. Contains the expressions
// configured for the domain specified in the request.
type DescribeExpressionsOutput struct {

	// The expressions configured for the domain.
	//
	// This member is required.
	Expressions []types.ExpressionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeExpressionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeExpressions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeExpressions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeExpressions"); err != nil {
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
	if err = addOpDescribeExpressionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeExpressions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeExpressions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeExpressions",
	}
}
