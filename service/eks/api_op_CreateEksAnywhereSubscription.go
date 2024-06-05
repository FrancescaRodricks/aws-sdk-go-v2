// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an EKS Anywhere subscription. When a subscription is created, it is a
// contract agreement for the length of the term specified in the request. Licenses
// that are used to validate support are provisioned in Amazon Web Services License
// Manager and the caller account is granted access to EKS Anywhere Curated
// Packages.
func (c *Client) CreateEksAnywhereSubscription(ctx context.Context, params *CreateEksAnywhereSubscriptionInput, optFns ...func(*Options)) (*CreateEksAnywhereSubscriptionOutput, error) {
	if params == nil {
		params = &CreateEksAnywhereSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEksAnywhereSubscription", params, optFns, c.addOperationCreateEksAnywhereSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEksAnywhereSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEksAnywhereSubscriptionInput struct {

	// The unique name for your subscription. It must be unique in your Amazon Web
	// Services account in the Amazon Web Services Region you're creating the
	// subscription in. The name can contain only alphanumeric characters
	// (case-sensitive), hyphens, and underscores. It must start with an alphabetic
	// character and can't be longer than 100 characters.
	//
	// This member is required.
	Name *string

	// An object representing the term duration and term unit type of your
	// subscription. This determines the term length of your subscription. Valid values
	// are MONTHS for term unit and 12 or 36 for term duration, indicating a 12 month
	// or 36 month subscription. This value cannot be changed after creating the
	// subscription.
	//
	// This member is required.
	Term *types.EksAnywhereSubscriptionTerm

	// A boolean indicating whether the subscription auto renews at the end of the
	// term.
	AutoRenew bool

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientRequestToken *string

	// The number of licenses to purchase with the subscription. Valid values are
	// between 1 and 100. This value can't be changed after creating the subscription.
	LicenseQuantity int32

	// The license type for all licenses in the subscription. Valid value is CLUSTER.
	// With the CLUSTER license type, each license covers support for a single EKS
	// Anywhere cluster.
	LicenseType types.EksAnywhereSubscriptionLicenseType

	// The metadata for a subscription to assist with categorization and organization.
	// Each tag consists of a key and an optional value. Subscription tags don't
	// propagate to any other resources associated with the subscription.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateEksAnywhereSubscriptionOutput struct {

	// The full description of the subscription.
	Subscription *types.EksAnywhereSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEksAnywhereSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateEksAnywhereSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateEksAnywhereSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateEksAnywhereSubscription"); err != nil {
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
	if err = addIdempotencyToken_opCreateEksAnywhereSubscriptionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateEksAnywhereSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEksAnywhereSubscription(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateEksAnywhereSubscription struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateEksAnywhereSubscription) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateEksAnywhereSubscription) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateEksAnywhereSubscriptionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateEksAnywhereSubscriptionInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateEksAnywhereSubscriptionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateEksAnywhereSubscription{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateEksAnywhereSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateEksAnywhereSubscription",
	}
}
