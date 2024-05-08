// Code generated by smithy-go-codegen DO NOT EDIT.

package vpclattice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/vpclattice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a listener rule. Each listener has a default rule for checking
// connection requests, but you can define additional rules. Each rule consists of
// a priority, one or more actions, and one or more conditions. For more
// information, see [Listener rules]in the Amazon VPC Lattice User Guide.
//
// [Listener rules]: https://docs.aws.amazon.com/vpc-lattice/latest/ug/listeners.html#listener-rules
func (c *Client) CreateRule(ctx context.Context, params *CreateRuleInput, optFns ...func(*Options)) (*CreateRuleOutput, error) {
	if params == nil {
		params = &CreateRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRule", params, optFns, c.addOperationCreateRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRuleInput struct {

	// The action for the default rule.
	//
	// This member is required.
	Action types.RuleAction

	// The ID or Amazon Resource Name (ARN) of the listener.
	//
	// This member is required.
	ListenerIdentifier *string

	// The rule match.
	//
	// This member is required.
	Match types.RuleMatch

	// The name of the rule. The name must be unique within the listener. The valid
	// characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or
	// last character, or immediately after another hyphen.
	//
	// This member is required.
	Name *string

	// The priority assigned to the rule. Each rule for a specific listener must have
	// a unique priority. The lower the priority number the higher the priority.
	//
	// This member is required.
	Priority *int32

	// The ID or Amazon Resource Name (ARN) of the service.
	//
	// This member is required.
	ServiceIdentifier *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If you retry a request that completed successfully using the
	// same client token and parameters, the retry succeeds without performing any
	// actions. If the parameters aren't identical, the retry fails.
	ClientToken *string

	// The tags for the rule.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateRuleOutput struct {

	// The rule action. Each rule must include exactly one of the following types of
	// actions: forward or fixed-response , and it must be the last action to be
	// performed.
	Action types.RuleAction

	// The Amazon Resource Name (ARN) of the rule.
	Arn *string

	// The ID of the rule.
	Id *string

	// The rule match. The RuleMatch must be an HttpMatch . This means that the rule
	// should be an exact match on HTTP constraints which are made up of the HTTP
	// method, path, and header.
	Match types.RuleMatch

	// The name of the rule.
	Name *string

	// The priority assigned to the rule. The lower the priority number the higher the
	// priority.
	Priority *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateRule"); err != nil {
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
	if err = addIdempotencyToken_opCreateRuleMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRule(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateRule struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateRule) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateRuleInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateRuleMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateRule{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRule",
	}
}
