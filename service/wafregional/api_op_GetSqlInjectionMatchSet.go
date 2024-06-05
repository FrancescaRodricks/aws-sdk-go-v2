// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns the SqlInjectionMatchSet that is specified by SqlInjectionMatchSetId .
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
func (c *Client) GetSqlInjectionMatchSet(ctx context.Context, params *GetSqlInjectionMatchSetInput, optFns ...func(*Options)) (*GetSqlInjectionMatchSetOutput, error) {
	if params == nil {
		params = &GetSqlInjectionMatchSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSqlInjectionMatchSet", params, optFns, c.addOperationGetSqlInjectionMatchSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSqlInjectionMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to get a SqlInjectionMatchSet.
type GetSqlInjectionMatchSetInput struct {

	// The SqlInjectionMatchSetId of the SqlInjectionMatchSet that you want to get. SqlInjectionMatchSetId
	// is returned by CreateSqlInjectionMatchSetand by ListSqlInjectionMatchSets.
	//
	// This member is required.
	SqlInjectionMatchSetId *string

	noSmithyDocumentSerde
}

// The response to a GetSqlInjectionMatchSet request.
type GetSqlInjectionMatchSetOutput struct {

	// Information about the SqlInjectionMatchSet that you specified in the GetSqlInjectionMatchSet
	// request. For more information, see the following topics:
	//
	// SqlInjectionMatchSet
	//   - : Contains Name , SqlInjectionMatchSetId , and an array of
	//   SqlInjectionMatchTuple objects
	//
	// SqlInjectionMatchTuple
	//   - : Each SqlInjectionMatchTuple object contains FieldToMatch and
	//   TextTransformation
	//
	// FieldToMatch
	//   - : Contains Data and Type
	SqlInjectionMatchSet *types.SqlInjectionMatchSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSqlInjectionMatchSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSqlInjectionMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSqlInjectionMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSqlInjectionMatchSet"); err != nil {
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
	if err = addOpGetSqlInjectionMatchSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSqlInjectionMatchSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSqlInjectionMatchSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSqlInjectionMatchSet",
	}
}
