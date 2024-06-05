// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Describes the settings that a bot has for a specific locale.
func (c *Client) DescribeBotLocale(ctx context.Context, params *DescribeBotLocaleInput, optFns ...func(*Options)) (*DescribeBotLocaleOutput, error) {
	if params == nil {
		params = &DescribeBotLocaleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBotLocale", params, optFns, c.addOperationDescribeBotLocaleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBotLocaleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBotLocaleInput struct {

	// The identifier of the bot associated with the locale.
	//
	// This member is required.
	BotId *string

	// The version of the bot associated with the locale.
	//
	// This member is required.
	BotVersion *string

	// The unique identifier of the locale to describe. The string must match one of
	// the supported locales. For more information, see [Supported languages].
	//
	// [Supported languages]: https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html
	//
	// This member is required.
	LocaleId *string

	noSmithyDocumentSerde
}

type DescribeBotLocaleOutput struct {

	// The identifier of the bot associated with the locale.
	BotId *string

	// History of changes, such as when a locale is used in an alias, that have taken
	// place for the locale.
	BotLocaleHistoryEvents []types.BotLocaleHistoryEvent

	// The status of the bot. If the status is Failed , the reasons for the failure are
	// listed in the failureReasons field.
	BotLocaleStatus types.BotLocaleStatus

	// The version of the bot associated with the locale.
	BotVersion *string

	// The date and time that the locale was created.
	CreationDateTime *time.Time

	// The description of the locale.
	Description *string

	// if botLocaleStatus is Failed , Amazon Lex explains why it failed to build the
	// bot.
	FailureReasons []string

	// Contains settings for Amazon Bedrock's generative AI features for your bot
	// locale.
	GenerativeAISettings *types.GenerativeAISettings

	// The number of intents defined for the locale.
	IntentsCount *int32

	// The date and time that the locale was last submitted for building.
	LastBuildSubmittedDateTime *time.Time

	// The date and time that the locale was last updated.
	LastUpdatedDateTime *time.Time

	// The unique identifier of the described locale.
	LocaleId *string

	// The name of the locale.
	LocaleName *string

	// The confidence threshold where Amazon Lex inserts the AMAZON.FallbackIntent and
	// AMAZON.KendraSearchIntent intents in the list of possible intents for an
	// utterance.
	NluIntentConfidenceThreshold *float64

	// Recommended actions to take to resolve an error in the failureReasons field.
	RecommendedActions []string

	// The number of slot types defined for the locale.
	SlotTypesCount *int32

	// The Amazon Polly voice Amazon Lex uses for voice interaction with the user.
	VoiceSettings *types.VoiceSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBotLocaleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBotLocale{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBotLocale{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeBotLocale"); err != nil {
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
	if err = addOpDescribeBotLocaleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBotLocale(options.Region), middleware.Before); err != nil {
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

// DescribeBotLocaleAPIClient is a client that implements the DescribeBotLocale
// operation.
type DescribeBotLocaleAPIClient interface {
	DescribeBotLocale(context.Context, *DescribeBotLocaleInput, ...func(*Options)) (*DescribeBotLocaleOutput, error)
}

var _ DescribeBotLocaleAPIClient = (*Client)(nil)

// BotLocaleBuiltWaiterOptions are waiter options for BotLocaleBuiltWaiter
type BotLocaleBuiltWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	//
	// Passing options here is functionally equivalent to passing values to this
	// config's ClientOptions field that extend the inner client's APIOptions directly.
	APIOptions []func(*middleware.Stack) error

	// Functional options to be passed to all operations invoked by this client.
	//
	// Function values that modify the inner APIOptions are applied after the waiter
	// config's own APIOptions modifiers.
	ClientOptions []func(*Options)

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// BotLocaleBuiltWaiter will use default minimum delay of 10 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, BotLocaleBuiltWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state.
	//
	// By default service-modeled logic will populate this option. This option can
	// thus be used to define a custom waiter state with fall-back to service-modeled
	// waiter state mutators.The function returns an error in case of a failure state.
	// In case of retry state, this function returns a bool value of true and nil
	// error, while in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeBotLocaleInput, *DescribeBotLocaleOutput, error) (bool, error)
}

// BotLocaleBuiltWaiter defines the waiters for BotLocaleBuilt
type BotLocaleBuiltWaiter struct {
	client DescribeBotLocaleAPIClient

	options BotLocaleBuiltWaiterOptions
}

// NewBotLocaleBuiltWaiter constructs a BotLocaleBuiltWaiter.
func NewBotLocaleBuiltWaiter(client DescribeBotLocaleAPIClient, optFns ...func(*BotLocaleBuiltWaiterOptions)) *BotLocaleBuiltWaiter {
	options := BotLocaleBuiltWaiterOptions{}
	options.MinDelay = 10 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = botLocaleBuiltStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &BotLocaleBuiltWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for BotLocaleBuilt waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *BotLocaleBuiltWaiter) Wait(ctx context.Context, params *DescribeBotLocaleInput, maxWaitDur time.Duration, optFns ...func(*BotLocaleBuiltWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for BotLocaleBuilt waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *BotLocaleBuiltWaiter) WaitForOutput(ctx context.Context, params *DescribeBotLocaleInput, maxWaitDur time.Duration, optFns ...func(*BotLocaleBuiltWaiterOptions)) (*DescribeBotLocaleOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeBotLocale(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
			for _, opt := range options.ClientOptions {
				opt(o)
			}
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for BotLocaleBuilt waiter")
}

func botLocaleBuiltStateRetryable(ctx context.Context, input *DescribeBotLocaleInput, output *DescribeBotLocaleOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("botLocaleStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Built"
		value, ok := pathValue.(types.BotLocaleStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.BotLocaleStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("botLocaleStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Deleting"
		value, ok := pathValue.(types.BotLocaleStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.BotLocaleStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("botLocaleStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Failed"
		value, ok := pathValue.(types.BotLocaleStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.BotLocaleStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("botLocaleStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "NotBuilt"
		value, ok := pathValue.(types.BotLocaleStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.BotLocaleStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// BotLocaleCreatedWaiterOptions are waiter options for BotLocaleCreatedWaiter
type BotLocaleCreatedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	//
	// Passing options here is functionally equivalent to passing values to this
	// config's ClientOptions field that extend the inner client's APIOptions directly.
	APIOptions []func(*middleware.Stack) error

	// Functional options to be passed to all operations invoked by this client.
	//
	// Function values that modify the inner APIOptions are applied after the waiter
	// config's own APIOptions modifiers.
	ClientOptions []func(*Options)

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// BotLocaleCreatedWaiter will use default minimum delay of 10 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, BotLocaleCreatedWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state.
	//
	// By default service-modeled logic will populate this option. This option can
	// thus be used to define a custom waiter state with fall-back to service-modeled
	// waiter state mutators.The function returns an error in case of a failure state.
	// In case of retry state, this function returns a bool value of true and nil
	// error, while in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeBotLocaleInput, *DescribeBotLocaleOutput, error) (bool, error)
}

// BotLocaleCreatedWaiter defines the waiters for BotLocaleCreated
type BotLocaleCreatedWaiter struct {
	client DescribeBotLocaleAPIClient

	options BotLocaleCreatedWaiterOptions
}

// NewBotLocaleCreatedWaiter constructs a BotLocaleCreatedWaiter.
func NewBotLocaleCreatedWaiter(client DescribeBotLocaleAPIClient, optFns ...func(*BotLocaleCreatedWaiterOptions)) *BotLocaleCreatedWaiter {
	options := BotLocaleCreatedWaiterOptions{}
	options.MinDelay = 10 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = botLocaleCreatedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &BotLocaleCreatedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for BotLocaleCreated waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *BotLocaleCreatedWaiter) Wait(ctx context.Context, params *DescribeBotLocaleInput, maxWaitDur time.Duration, optFns ...func(*BotLocaleCreatedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for BotLocaleCreated waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *BotLocaleCreatedWaiter) WaitForOutput(ctx context.Context, params *DescribeBotLocaleInput, maxWaitDur time.Duration, optFns ...func(*BotLocaleCreatedWaiterOptions)) (*DescribeBotLocaleOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeBotLocale(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
			for _, opt := range options.ClientOptions {
				opt(o)
			}
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for BotLocaleCreated waiter")
}

func botLocaleCreatedStateRetryable(ctx context.Context, input *DescribeBotLocaleInput, output *DescribeBotLocaleOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("botLocaleStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Built"
		value, ok := pathValue.(types.BotLocaleStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.BotLocaleStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("botLocaleStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "ReadyExpressTesting"
		value, ok := pathValue.(types.BotLocaleStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.BotLocaleStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("botLocaleStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "NotBuilt"
		value, ok := pathValue.(types.BotLocaleStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.BotLocaleStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("botLocaleStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Deleting"
		value, ok := pathValue.(types.BotLocaleStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.BotLocaleStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("botLocaleStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Failed"
		value, ok := pathValue.(types.BotLocaleStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.BotLocaleStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// BotLocaleExpressTestingAvailableWaiterOptions are waiter options for
// BotLocaleExpressTestingAvailableWaiter
type BotLocaleExpressTestingAvailableWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	//
	// Passing options here is functionally equivalent to passing values to this
	// config's ClientOptions field that extend the inner client's APIOptions directly.
	APIOptions []func(*middleware.Stack) error

	// Functional options to be passed to all operations invoked by this client.
	//
	// Function values that modify the inner APIOptions are applied after the waiter
	// config's own APIOptions modifiers.
	ClientOptions []func(*Options)

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// BotLocaleExpressTestingAvailableWaiter will use default minimum delay of 10
	// seconds. Note that MinDelay must resolve to a value lesser than or equal to the
	// MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, BotLocaleExpressTestingAvailableWaiter will use default max delay
	// of 120 seconds. Note that MaxDelay must resolve to value greater than or equal
	// to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state.
	//
	// By default service-modeled logic will populate this option. This option can
	// thus be used to define a custom waiter state with fall-back to service-modeled
	// waiter state mutators.The function returns an error in case of a failure state.
	// In case of retry state, this function returns a bool value of true and nil
	// error, while in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeBotLocaleInput, *DescribeBotLocaleOutput, error) (bool, error)
}

// BotLocaleExpressTestingAvailableWaiter defines the waiters for
// BotLocaleExpressTestingAvailable
type BotLocaleExpressTestingAvailableWaiter struct {
	client DescribeBotLocaleAPIClient

	options BotLocaleExpressTestingAvailableWaiterOptions
}

// NewBotLocaleExpressTestingAvailableWaiter constructs a
// BotLocaleExpressTestingAvailableWaiter.
func NewBotLocaleExpressTestingAvailableWaiter(client DescribeBotLocaleAPIClient, optFns ...func(*BotLocaleExpressTestingAvailableWaiterOptions)) *BotLocaleExpressTestingAvailableWaiter {
	options := BotLocaleExpressTestingAvailableWaiterOptions{}
	options.MinDelay = 10 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = botLocaleExpressTestingAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &BotLocaleExpressTestingAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for BotLocaleExpressTestingAvailable waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *BotLocaleExpressTestingAvailableWaiter) Wait(ctx context.Context, params *DescribeBotLocaleInput, maxWaitDur time.Duration, optFns ...func(*BotLocaleExpressTestingAvailableWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for BotLocaleExpressTestingAvailable
// waiter and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *BotLocaleExpressTestingAvailableWaiter) WaitForOutput(ctx context.Context, params *DescribeBotLocaleInput, maxWaitDur time.Duration, optFns ...func(*BotLocaleExpressTestingAvailableWaiterOptions)) (*DescribeBotLocaleOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeBotLocale(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
			for _, opt := range options.ClientOptions {
				opt(o)
			}
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for BotLocaleExpressTestingAvailable waiter")
}

func botLocaleExpressTestingAvailableStateRetryable(ctx context.Context, input *DescribeBotLocaleInput, output *DescribeBotLocaleOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("botLocaleStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Built"
		value, ok := pathValue.(types.BotLocaleStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.BotLocaleStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("botLocaleStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "ReadyExpressTesting"
		value, ok := pathValue.(types.BotLocaleStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.BotLocaleStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("botLocaleStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Deleting"
		value, ok := pathValue.(types.BotLocaleStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.BotLocaleStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("botLocaleStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Failed"
		value, ok := pathValue.(types.BotLocaleStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.BotLocaleStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("botLocaleStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "NotBuilt"
		value, ok := pathValue.(types.BotLocaleStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.BotLocaleStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeBotLocale(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeBotLocale",
	}
}
