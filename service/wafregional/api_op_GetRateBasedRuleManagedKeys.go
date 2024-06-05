// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
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
// Returns an array of IP addresses currently being blocked by the RateBasedRule that is
// specified by the RuleId . The maximum number of managed keys that will be
// blocked is 10,000. If more than 10,000 addresses exceed the rate limit, the
// 10,000 addresses with the highest rates will be blocked.
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
func (c *Client) GetRateBasedRuleManagedKeys(ctx context.Context, params *GetRateBasedRuleManagedKeysInput, optFns ...func(*Options)) (*GetRateBasedRuleManagedKeysOutput, error) {
	if params == nil {
		params = &GetRateBasedRuleManagedKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRateBasedRuleManagedKeys", params, optFns, c.addOperationGetRateBasedRuleManagedKeysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRateBasedRuleManagedKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRateBasedRuleManagedKeysInput struct {

	// The RuleId of the RateBasedRule for which you want to get a list of ManagedKeys . RuleId is
	// returned by CreateRateBasedRuleand by ListRateBasedRules.
	//
	// This member is required.
	RuleId *string

	// A null value and not currently used. Do not include this in your request.
	NextMarker *string

	noSmithyDocumentSerde
}

type GetRateBasedRuleManagedKeysOutput struct {

	// An array of IP addresses that currently are blocked by the specified RateBasedRule.
	ManagedKeys []string

	// A null value and not currently used.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRateBasedRuleManagedKeysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRateBasedRuleManagedKeys{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRateBasedRuleManagedKeys{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetRateBasedRuleManagedKeys"); err != nil {
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
	if err = addOpGetRateBasedRuleManagedKeysValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRateBasedRuleManagedKeys(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRateBasedRuleManagedKeys(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetRateBasedRuleManagedKeys",
	}
}
