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
// Returns an array of ByteMatchSetSummary objects.
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
func (c *Client) ListByteMatchSets(ctx context.Context, params *ListByteMatchSetsInput, optFns ...func(*Options)) (*ListByteMatchSetsOutput, error) {
	if params == nil {
		params = &ListByteMatchSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListByteMatchSets", params, optFns, c.addOperationListByteMatchSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListByteMatchSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListByteMatchSetsInput struct {

	// Specifies the number of ByteMatchSet objects that you want AWS WAF to return
	// for this request. If you have more ByteMatchSets objects than the number you
	// specify for Limit , the response includes a NextMarker value that you can use
	// to get another batch of ByteMatchSet objects.
	Limit int32

	// If you specify a value for Limit and you have more ByteMatchSets than the value
	// of Limit , AWS WAF returns a NextMarker value in the response that allows you
	// to list another group of ByteMatchSets . For the second and subsequent
	// ListByteMatchSets requests, specify the value of NextMarker from the previous
	// response to get information about another batch of ByteMatchSets .
	NextMarker *string

	noSmithyDocumentSerde
}

type ListByteMatchSetsOutput struct {

	// An array of ByteMatchSetSummary objects.
	ByteMatchSets []types.ByteMatchSetSummary

	// If you have more ByteMatchSet objects than the number that you specified for
	// Limit in the request, the response includes a NextMarker value. To list more
	// ByteMatchSet objects, submit another ListByteMatchSets request, and specify the
	// NextMarker value from the response in the NextMarker value in the next request.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListByteMatchSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListByteMatchSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListByteMatchSets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListByteMatchSets"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListByteMatchSets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListByteMatchSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListByteMatchSets",
	}
}
