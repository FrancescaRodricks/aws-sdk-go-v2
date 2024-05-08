// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunegraph

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/neptunegraph/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"strconv"
	"time"
)

// Retrieves information about a specified private endpoint.
func (c *Client) GetPrivateGraphEndpoint(ctx context.Context, params *GetPrivateGraphEndpointInput, optFns ...func(*Options)) (*GetPrivateGraphEndpointOutput, error) {
	if params == nil {
		params = &GetPrivateGraphEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPrivateGraphEndpoint", params, optFns, c.addOperationGetPrivateGraphEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPrivateGraphEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPrivateGraphEndpointInput struct {

	// The unique identifier of the Neptune Analytics graph.
	//
	// This member is required.
	GraphIdentifier *string

	// The ID of the VPC where the private endpoint is located.
	//
	// This member is required.
	VpcId *string

	noSmithyDocumentSerde
}

func (in *GetPrivateGraphEndpointInput) bindEndpointParams(p *EndpointParameters) {

	p.ApiType = ptr.String("ControlPlane")
}

type GetPrivateGraphEndpointOutput struct {

	// The current status of the private endpoint.
	//
	// This member is required.
	Status types.PrivateGraphEndpointStatus

	// The subnet IDs involved.
	//
	// This member is required.
	SubnetIds []string

	// The ID of the VPC where the private endpoint is located.
	//
	// This member is required.
	VpcId *string

	// The ID of the private endpoint.
	VpcEndpointId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPrivateGraphEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetPrivateGraphEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPrivateGraphEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetPrivateGraphEndpoint"); err != nil {
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
	if err = addOpGetPrivateGraphEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPrivateGraphEndpoint(options.Region), middleware.Before); err != nil {
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

// GetPrivateGraphEndpointAPIClient is a client that implements the
// GetPrivateGraphEndpoint operation.
type GetPrivateGraphEndpointAPIClient interface {
	GetPrivateGraphEndpoint(context.Context, *GetPrivateGraphEndpointInput, ...func(*Options)) (*GetPrivateGraphEndpointOutput, error)
}

var _ GetPrivateGraphEndpointAPIClient = (*Client)(nil)

// PrivateGraphEndpointAvailableWaiterOptions are waiter options for
// PrivateGraphEndpointAvailableWaiter
type PrivateGraphEndpointAvailableWaiterOptions struct {

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
	// PrivateGraphEndpointAvailableWaiter will use default minimum delay of 10
	// seconds. Note that MinDelay must resolve to a value lesser than or equal to the
	// MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, PrivateGraphEndpointAvailableWaiter will use default max delay of
	// 1800 seconds. Note that MaxDelay must resolve to value greater than or equal to
	// the MinDelay.
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
	Retryable func(context.Context, *GetPrivateGraphEndpointInput, *GetPrivateGraphEndpointOutput, error) (bool, error)
}

// PrivateGraphEndpointAvailableWaiter defines the waiters for
// PrivateGraphEndpointAvailable
type PrivateGraphEndpointAvailableWaiter struct {
	client GetPrivateGraphEndpointAPIClient

	options PrivateGraphEndpointAvailableWaiterOptions
}

// NewPrivateGraphEndpointAvailableWaiter constructs a
// PrivateGraphEndpointAvailableWaiter.
func NewPrivateGraphEndpointAvailableWaiter(client GetPrivateGraphEndpointAPIClient, optFns ...func(*PrivateGraphEndpointAvailableWaiterOptions)) *PrivateGraphEndpointAvailableWaiter {
	options := PrivateGraphEndpointAvailableWaiterOptions{}
	options.MinDelay = 10 * time.Second
	options.MaxDelay = 1800 * time.Second
	options.Retryable = privateGraphEndpointAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &PrivateGraphEndpointAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for PrivateGraphEndpointAvailable waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *PrivateGraphEndpointAvailableWaiter) Wait(ctx context.Context, params *GetPrivateGraphEndpointInput, maxWaitDur time.Duration, optFns ...func(*PrivateGraphEndpointAvailableWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for PrivateGraphEndpointAvailable
// waiter and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *PrivateGraphEndpointAvailableWaiter) WaitForOutput(ctx context.Context, params *GetPrivateGraphEndpointInput, maxWaitDur time.Duration, optFns ...func(*PrivateGraphEndpointAvailableWaiterOptions)) (*GetPrivateGraphEndpointOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 1800 * time.Second
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

		out, err := w.client.GetPrivateGraphEndpoint(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for PrivateGraphEndpointAvailable waiter")
}

func privateGraphEndpointAvailableStateRetryable(ctx context.Context, input *GetPrivateGraphEndpointInput, output *GetPrivateGraphEndpointOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DELETING"
		value, ok := pathValue.(types.PrivateGraphEndpointStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.PrivateGraphEndpointStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "FAILED"
		value, ok := pathValue.(types.PrivateGraphEndpointStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.PrivateGraphEndpointStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "AVAILABLE"
		value, ok := pathValue.(types.PrivateGraphEndpointStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.PrivateGraphEndpointStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	return true, nil
}

// PrivateGraphEndpointDeletedWaiterOptions are waiter options for
// PrivateGraphEndpointDeletedWaiter
type PrivateGraphEndpointDeletedWaiterOptions struct {

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
	// PrivateGraphEndpointDeletedWaiter will use default minimum delay of 10 seconds.
	// Note that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, PrivateGraphEndpointDeletedWaiter will use default max delay of
	// 1800 seconds. Note that MaxDelay must resolve to value greater than or equal to
	// the MinDelay.
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
	Retryable func(context.Context, *GetPrivateGraphEndpointInput, *GetPrivateGraphEndpointOutput, error) (bool, error)
}

// PrivateGraphEndpointDeletedWaiter defines the waiters for
// PrivateGraphEndpointDeleted
type PrivateGraphEndpointDeletedWaiter struct {
	client GetPrivateGraphEndpointAPIClient

	options PrivateGraphEndpointDeletedWaiterOptions
}

// NewPrivateGraphEndpointDeletedWaiter constructs a
// PrivateGraphEndpointDeletedWaiter.
func NewPrivateGraphEndpointDeletedWaiter(client GetPrivateGraphEndpointAPIClient, optFns ...func(*PrivateGraphEndpointDeletedWaiterOptions)) *PrivateGraphEndpointDeletedWaiter {
	options := PrivateGraphEndpointDeletedWaiterOptions{}
	options.MinDelay = 10 * time.Second
	options.MaxDelay = 1800 * time.Second
	options.Retryable = privateGraphEndpointDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &PrivateGraphEndpointDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for PrivateGraphEndpointDeleted waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *PrivateGraphEndpointDeletedWaiter) Wait(ctx context.Context, params *GetPrivateGraphEndpointInput, maxWaitDur time.Duration, optFns ...func(*PrivateGraphEndpointDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for PrivateGraphEndpointDeleted waiter
// and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *PrivateGraphEndpointDeletedWaiter) WaitForOutput(ctx context.Context, params *GetPrivateGraphEndpointInput, maxWaitDur time.Duration, optFns ...func(*PrivateGraphEndpointDeletedWaiterOptions)) (*GetPrivateGraphEndpointOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 1800 * time.Second
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

		out, err := w.client.GetPrivateGraphEndpoint(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for PrivateGraphEndpointDeleted waiter")
}

func privateGraphEndpointDeletedStateRetryable(ctx context.Context, input *GetPrivateGraphEndpointInput, output *GetPrivateGraphEndpointOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("status != 'DELETING'", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "true"
		bv, err := strconv.ParseBool(expectedValue)
		if err != nil {
			return false, fmt.Errorf("error parsing boolean from string %w", err)
		}
		value, ok := pathValue.(bool)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected bool value got %T", pathValue)
		}

		if value == bv {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err != nil {
		var errorType *types.ResourceNotFoundException
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetPrivateGraphEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetPrivateGraphEndpoint",
	}
}
