// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoverycontrolconfig

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Displays details about a control panel.
func (c *Client) DescribeControlPanel(ctx context.Context, params *DescribeControlPanelInput, optFns ...func(*Options)) (*DescribeControlPanelOutput, error) {
	if params == nil {
		params = &DescribeControlPanelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeControlPanel", params, optFns, c.addOperationDescribeControlPanelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeControlPanelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeControlPanelInput struct {

	// The Amazon Resource Name (ARN) of the control panel.
	//
	// This member is required.
	ControlPanelArn *string

	noSmithyDocumentSerde
}

type DescribeControlPanelOutput struct {

	// Information about the control panel.
	ControlPanel *types.ControlPanel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeControlPanelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeControlPanel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeControlPanel{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeControlPanel"); err != nil {
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
	if err = addOpDescribeControlPanelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeControlPanel(options.Region), middleware.Before); err != nil {
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

// DescribeControlPanelAPIClient is a client that implements the
// DescribeControlPanel operation.
type DescribeControlPanelAPIClient interface {
	DescribeControlPanel(context.Context, *DescribeControlPanelInput, ...func(*Options)) (*DescribeControlPanelOutput, error)
}

var _ DescribeControlPanelAPIClient = (*Client)(nil)

// ControlPanelCreatedWaiterOptions are waiter options for
// ControlPanelCreatedWaiter
type ControlPanelCreatedWaiterOptions struct {

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
	// ControlPanelCreatedWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, ControlPanelCreatedWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
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
	Retryable func(context.Context, *DescribeControlPanelInput, *DescribeControlPanelOutput, error) (bool, error)
}

// ControlPanelCreatedWaiter defines the waiters for ControlPanelCreated
type ControlPanelCreatedWaiter struct {
	client DescribeControlPanelAPIClient

	options ControlPanelCreatedWaiterOptions
}

// NewControlPanelCreatedWaiter constructs a ControlPanelCreatedWaiter.
func NewControlPanelCreatedWaiter(client DescribeControlPanelAPIClient, optFns ...func(*ControlPanelCreatedWaiterOptions)) *ControlPanelCreatedWaiter {
	options := ControlPanelCreatedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = controlPanelCreatedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ControlPanelCreatedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ControlPanelCreated waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *ControlPanelCreatedWaiter) Wait(ctx context.Context, params *DescribeControlPanelInput, maxWaitDur time.Duration, optFns ...func(*ControlPanelCreatedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ControlPanelCreated waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *ControlPanelCreatedWaiter) WaitForOutput(ctx context.Context, params *DescribeControlPanelInput, maxWaitDur time.Duration, optFns ...func(*ControlPanelCreatedWaiterOptions)) (*DescribeControlPanelOutput, error) {
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

		out, err := w.client.DescribeControlPanel(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for ControlPanelCreated waiter")
}

func controlPanelCreatedStateRetryable(ctx context.Context, input *DescribeControlPanelInput, output *DescribeControlPanelOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("ControlPanel.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DEPLOYED"
		value, ok := pathValue.(types.Status)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.Status value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("ControlPanel.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "PENDING"
		value, ok := pathValue.(types.Status)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.Status value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err != nil {
		var errorType *types.InternalServerException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	return true, nil
}

// ControlPanelDeletedWaiterOptions are waiter options for
// ControlPanelDeletedWaiter
type ControlPanelDeletedWaiterOptions struct {

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
	// ControlPanelDeletedWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, ControlPanelDeletedWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
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
	Retryable func(context.Context, *DescribeControlPanelInput, *DescribeControlPanelOutput, error) (bool, error)
}

// ControlPanelDeletedWaiter defines the waiters for ControlPanelDeleted
type ControlPanelDeletedWaiter struct {
	client DescribeControlPanelAPIClient

	options ControlPanelDeletedWaiterOptions
}

// NewControlPanelDeletedWaiter constructs a ControlPanelDeletedWaiter.
func NewControlPanelDeletedWaiter(client DescribeControlPanelAPIClient, optFns ...func(*ControlPanelDeletedWaiterOptions)) *ControlPanelDeletedWaiter {
	options := ControlPanelDeletedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = controlPanelDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ControlPanelDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ControlPanelDeleted waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *ControlPanelDeletedWaiter) Wait(ctx context.Context, params *DescribeControlPanelInput, maxWaitDur time.Duration, optFns ...func(*ControlPanelDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ControlPanelDeleted waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *ControlPanelDeletedWaiter) WaitForOutput(ctx context.Context, params *DescribeControlPanelInput, maxWaitDur time.Duration, optFns ...func(*ControlPanelDeletedWaiterOptions)) (*DescribeControlPanelOutput, error) {
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

		out, err := w.client.DescribeControlPanel(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for ControlPanelDeleted waiter")
}

func controlPanelDeletedStateRetryable(ctx context.Context, input *DescribeControlPanelInput, output *DescribeControlPanelOutput, err error) (bool, error) {

	if err != nil {
		var errorType *types.ResourceNotFoundException
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("ControlPanel.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "PENDING_DELETION"
		value, ok := pathValue.(types.Status)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.Status value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err != nil {
		var errorType *types.InternalServerException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeControlPanel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeControlPanel",
	}
}
