// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Provides cluster-level details including status, hardware and software
// configuration, VPC settings, and so on.
func (c *Client) DescribeCluster(ctx context.Context, params *DescribeClusterInput, optFns ...func(*Options)) (*DescribeClusterOutput, error) {
	if params == nil {
		params = &DescribeClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCluster", params, optFns, c.addOperationDescribeClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This input determines which cluster to describe.
type DescribeClusterInput struct {

	// The identifier of the cluster to describe.
	//
	// This member is required.
	ClusterId *string

	noSmithyDocumentSerde
}

// This output contains the description of the cluster.
type DescribeClusterOutput struct {

	// This output contains the details for the requested cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeCluster"); err != nil {
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
	if err = addOpDescribeClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCluster(options.Region), middleware.Before); err != nil {
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

// DescribeClusterAPIClient is a client that implements the DescribeCluster
// operation.
type DescribeClusterAPIClient interface {
	DescribeCluster(context.Context, *DescribeClusterInput, ...func(*Options)) (*DescribeClusterOutput, error)
}

var _ DescribeClusterAPIClient = (*Client)(nil)

// ClusterRunningWaiterOptions are waiter options for ClusterRunningWaiter
type ClusterRunningWaiterOptions struct {

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
	// ClusterRunningWaiter will use default minimum delay of 30 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, ClusterRunningWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *DescribeClusterInput, *DescribeClusterOutput, error) (bool, error)
}

// ClusterRunningWaiter defines the waiters for ClusterRunning
type ClusterRunningWaiter struct {
	client DescribeClusterAPIClient

	options ClusterRunningWaiterOptions
}

// NewClusterRunningWaiter constructs a ClusterRunningWaiter.
func NewClusterRunningWaiter(client DescribeClusterAPIClient, optFns ...func(*ClusterRunningWaiterOptions)) *ClusterRunningWaiter {
	options := ClusterRunningWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = clusterRunningStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ClusterRunningWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ClusterRunning waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *ClusterRunningWaiter) Wait(ctx context.Context, params *DescribeClusterInput, maxWaitDur time.Duration, optFns ...func(*ClusterRunningWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ClusterRunning waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *ClusterRunningWaiter) WaitForOutput(ctx context.Context, params *DescribeClusterInput, maxWaitDur time.Duration, optFns ...func(*ClusterRunningWaiterOptions)) (*DescribeClusterOutput, error) {
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

		out, err := w.client.DescribeCluster(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for ClusterRunning waiter")
}

func clusterRunningStateRetryable(ctx context.Context, input *DescribeClusterInput, output *DescribeClusterOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Cluster.Status.State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "RUNNING"
		value, ok := pathValue.(types.ClusterState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ClusterState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Cluster.Status.State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "WAITING"
		value, ok := pathValue.(types.ClusterState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ClusterState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Cluster.Status.State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "TERMINATING"
		value, ok := pathValue.(types.ClusterState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ClusterState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Cluster.Status.State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "TERMINATED"
		value, ok := pathValue.(types.ClusterState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ClusterState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Cluster.Status.State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "TERMINATED_WITH_ERRORS"
		value, ok := pathValue.(types.ClusterState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ClusterState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// ClusterTerminatedWaiterOptions are waiter options for ClusterTerminatedWaiter
type ClusterTerminatedWaiterOptions struct {

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
	// ClusterTerminatedWaiter will use default minimum delay of 30 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, ClusterTerminatedWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *DescribeClusterInput, *DescribeClusterOutput, error) (bool, error)
}

// ClusterTerminatedWaiter defines the waiters for ClusterTerminated
type ClusterTerminatedWaiter struct {
	client DescribeClusterAPIClient

	options ClusterTerminatedWaiterOptions
}

// NewClusterTerminatedWaiter constructs a ClusterTerminatedWaiter.
func NewClusterTerminatedWaiter(client DescribeClusterAPIClient, optFns ...func(*ClusterTerminatedWaiterOptions)) *ClusterTerminatedWaiter {
	options := ClusterTerminatedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = clusterTerminatedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ClusterTerminatedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ClusterTerminated waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *ClusterTerminatedWaiter) Wait(ctx context.Context, params *DescribeClusterInput, maxWaitDur time.Duration, optFns ...func(*ClusterTerminatedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ClusterTerminated waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *ClusterTerminatedWaiter) WaitForOutput(ctx context.Context, params *DescribeClusterInput, maxWaitDur time.Duration, optFns ...func(*ClusterTerminatedWaiterOptions)) (*DescribeClusterOutput, error) {
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

		out, err := w.client.DescribeCluster(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for ClusterTerminated waiter")
}

func clusterTerminatedStateRetryable(ctx context.Context, input *DescribeClusterInput, output *DescribeClusterOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Cluster.Status.State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "TERMINATED"
		value, ok := pathValue.(types.ClusterState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ClusterState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Cluster.Status.State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "TERMINATED_WITH_ERRORS"
		value, ok := pathValue.(types.ClusterState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ClusterState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeCluster",
	}
}
