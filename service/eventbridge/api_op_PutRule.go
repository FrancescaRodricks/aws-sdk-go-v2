// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates the specified rule. Rules are enabled by default, or based
// on value of the state. You can disable a rule using [DisableRule].
//
// A single rule watches for events from a single event bus. Events generated by
// Amazon Web Services services go to your account's default event bus. Events
// generated by SaaS partner services or applications go to the matching partner
// event bus. If you have custom applications or services, you can specify whether
// their events go to your default event bus or a custom event bus that you have
// created. For more information, see [CreateEventBus].
//
// If you are updating an existing rule, the rule is replaced with what you
// specify in this PutRule command. If you omit arguments in PutRule , the old
// values for those arguments are not kept. Instead, they are replaced with null
// values.
//
// When you create or update a rule, incoming events might not immediately start
// matching to new or updated rules. Allow a short period of time for changes to
// take effect.
//
// A rule must contain at least an EventPattern or ScheduleExpression. Rules with
// EventPatterns are triggered when a matching event is observed. Rules with
// ScheduleExpressions self-trigger based on the given schedule. A rule can have
// both an EventPattern and a ScheduleExpression, in which case the rule triggers
// on matching events as well as on a schedule.
//
// When you initially create a rule, you can optionally assign one or more tags to
// the rule. Tags can help you organize and categorize your resources. You can also
// use them to scope user permissions, by granting a user permission to access or
// change only rules with certain tag values. To use the PutRule operation and
// assign tags, you must have both the events:PutRule and events:TagResource
// permissions.
//
// If you are updating an existing rule, any tags you specify in the PutRule
// operation are ignored. To update the tags of an existing rule, use [TagResource]and [UntagResource].
//
// Most services in Amazon Web Services treat : or / as the same character in
// Amazon Resource Names (ARNs). However, EventBridge uses an exact match in event
// patterns and rules. Be sure to use the correct ARN characters when creating
// event patterns so that they match the ARN syntax in the event you want to match.
//
// In EventBridge, it is possible to create rules that lead to infinite loops,
// where a rule is fired repeatedly. For example, a rule might detect that ACLs
// have changed on an S3 bucket, and trigger software to change them to the desired
// state. If the rule is not written carefully, the subsequent change to the ACLs
// fires the rule again, creating an infinite loop.
//
// To prevent this, write the rules so that the triggered actions do not re-fire
// the same rule. For example, your rule could fire only if ACLs are found to be in
// a bad state, instead of after any change.
//
// An infinite loop can quickly cause higher than expected charges. We recommend
// that you use budgeting, which alerts you when charges exceed your specified
// limit. For more information, see [Managing Your Costs with Budgets].
//
// [CreateEventBus]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_CreateEventBus.html
// [TagResource]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_UntagResource.html
// [DisableRule]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_DisableRule.html
// [Managing Your Costs with Budgets]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/budgets-managing-costs.html
func (c *Client) PutRule(ctx context.Context, params *PutRuleInput, optFns ...func(*Options)) (*PutRuleOutput, error) {
	if params == nil {
		params = &PutRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRule", params, optFns, c.addOperationPutRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRuleInput struct {

	// The name of the rule that you are creating or updating.
	//
	// This member is required.
	Name *string

	// A description of the rule.
	Description *string

	// The name or ARN of the event bus to associate with this rule. If you omit this,
	// the default event bus is used.
	EventBusName *string

	// The event pattern. For more information, see [Amazon EventBridge event patterns] in the Amazon EventBridge User
	// Guide.
	//
	// [Amazon EventBridge event patterns]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-event-patterns.html
	EventPattern *string

	// The Amazon Resource Name (ARN) of the IAM role associated with the rule.
	//
	// If you're setting an event bus in another account as the target and that
	// account granted permission to your account through an organization instead of
	// directly by the account ID, you must specify a RoleArn with proper permissions
	// in the Target structure, instead of here in this parameter.
	RoleArn *string

	// The scheduling expression. For example, "cron(0 20 * * ? *)" or "rate(5
	// minutes)".
	ScheduleExpression *string

	// Indicates whether the rule is enabled or disabled.
	State types.RuleState

	// The list of key-value pairs to associate with the rule.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type PutRuleOutput struct {

	// The Amazon Resource Name (ARN) of the rule.
	RuleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutRule"); err != nil {
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
	if err = addOpPutRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutRule",
	}
}
