// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a RuleGroup per the specifications provided.
//
// A rule group defines a collection of rules to inspect and control web requests
// that you can use in a WebACL. When you create a rule group, you define an immutable
// capacity limit. If you update a rule group, you must stay within the capacity.
// This allows others to reuse the rule group with confidence in its capacity
// requirements.
func (c *Client) CreateRuleGroup(ctx context.Context, params *CreateRuleGroupInput, optFns ...func(*Options)) (*CreateRuleGroupOutput, error) {
	if params == nil {
		params = &CreateRuleGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRuleGroup", params, optFns, c.addOperationCreateRuleGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRuleGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRuleGroupInput struct {

	// The web ACL capacity units (WCUs) required for this rule group.
	//
	// When you create your own rule group, you define this, and you cannot change it
	// after creation. When you add or modify the rules in a rule group, WAF enforces
	// this limit. You can check the capacity for a set of rules using CheckCapacity.
	//
	// WAF uses WCUs to calculate and control the operating resources that are used to
	// run your rules, rule groups, and web ACLs. WAF calculates capacity differently
	// for each rule type, to reflect the relative cost of each rule. Simple rules that
	// cost little to run use fewer WCUs than more complex rules that use more
	// processing power. Rule group capacity is fixed at creation, which helps users
	// plan their web ACL WCU usage when they use a rule group. For more information,
	// see [WAF web ACL capacity units (WCU)]in the WAF Developer Guide.
	//
	// [WAF web ACL capacity units (WCU)]: https://docs.aws.amazon.com/waf/latest/developerguide/aws-waf-capacity-units.html
	//
	// This member is required.
	Capacity *int64

	// The name of the rule group. You cannot change the name of a rule group after
	// you create it.
	//
	// This member is required.
	Name *string

	// Specifies whether this is for an Amazon CloudFront distribution or for a
	// regional application. A regional application can be an Application Load Balancer
	// (ALB), an Amazon API Gateway REST API, an AppSync GraphQL API, an Amazon Cognito
	// user pool, an App Runner service, or an Amazon Web Services Verified Access
	// instance.
	//
	// To work with CloudFront, you must also specify the Region US East (N. Virginia)
	// as follows:
	//
	//   - CLI - Specify the Region when you use the CloudFront scope:
	//   --scope=CLOUDFRONT --region=us-east-1 .
	//
	//   - API and SDKs - For all calls, use the Region endpoint us-east-1.
	//
	// This member is required.
	Scope types.Scope

	// Defines and enables Amazon CloudWatch metrics and web request sample
	// collection.
	//
	// This member is required.
	VisibilityConfig *types.VisibilityConfig

	// A map of custom response keys and content bodies. When you create a rule with a
	// block action, you can send a custom response to the web request. You define
	// these for the rule group, and then use them in the rules that you define in the
	// rule group.
	//
	// For information about customizing web requests and responses, see [Customizing web requests and responses in WAF] in the WAF
	// Developer Guide.
	//
	// For information about the limits on count and size for custom request and
	// response settings, see [WAF quotas]in the WAF Developer Guide.
	//
	// [WAF quotas]: https://docs.aws.amazon.com/waf/latest/developerguide/limits.html
	// [Customizing web requests and responses in WAF]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html
	CustomResponseBodies map[string]types.CustomResponseBody

	// A description of the rule group that helps with identification.
	Description *string

	// The Rule statements used to identify the web requests that you want to manage. Each
	// rule includes one top-level statement that WAF uses to identify matching web
	// requests, and parameters that govern how WAF handles them.
	Rules []types.Rule

	// An array of key:value pairs to associate with the resource.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateRuleGroupOutput struct {

	// High-level information about a RuleGroup, returned by operations like create and list.
	// This provides information like the ID, that you can use to retrieve and manage a
	// RuleGroup , and the ARN, that you provide to the RuleGroupReferenceStatement to use the rule group in a Rule.
	Summary *types.RuleGroupSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRuleGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRuleGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRuleGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateRuleGroup"); err != nil {
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
	if err = addOpCreateRuleGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRuleGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRuleGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRuleGroup",
	}
}
