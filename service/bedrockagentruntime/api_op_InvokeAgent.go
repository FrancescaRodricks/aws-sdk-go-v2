// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagentruntime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime/types"
	"github.com/aws/smithy-go/middleware"
	smithysync "github.com/aws/smithy-go/sync"
	"sync"
)

// Sends a prompt for the agent to process and respond to. Use return control
// event type for function calling.
//
// The CLI doesn't support InvokeAgent .
//
//   - To continue the same conversation with an agent, use the same sessionId
//     value in the request.
//
//   - To activate trace enablement, turn enableTrace to true . Trace enablement
//     helps you follow the agent's reasoning process that led it to the information it
//     processed, the actions it took, and the final result it yielded. For more
//     information, see [Trace enablement].
//
//   - End a conversation by setting endSession to true .
//
//   - In the sessionState object, you can include attributes for the session or
//     prompt or parameters returned from the action group.
//
//   - Use return control event type for function calling.
//
// The response is returned in the bytes field of the chunk object.
//
//   - The attribution object contains citations for parts of the response.
//
//   - If you set enableTrace to true in the request, you can trace the agent's
//     steps and reasoning process that led it to the response.
//
//   - Errors are also surfaced in the response.
//
// [Trace enablement]: https://docs.aws.amazon.com/bedrock/latest/userguide/agents-test.html#trace-events
func (c *Client) InvokeAgent(ctx context.Context, params *InvokeAgentInput, optFns ...func(*Options)) (*InvokeAgentOutput, error) {
	if params == nil {
		params = &InvokeAgentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InvokeAgent", params, optFns, c.addOperationInvokeAgentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InvokeAgentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InvokeAgentInput struct {

	// The alias of the agent to use.
	//
	// This member is required.
	AgentAliasId *string

	// The unique identifier of the agent to use.
	//
	// This member is required.
	AgentId *string

	// The unique identifier of the session. Use the same value across requests to
	// continue the same conversation.
	//
	// This member is required.
	SessionId *string

	// Specifies whether to turn on the trace or not to track the agent's reasoning
	// process. For more information, see [Trace enablement].
	//
	// [Trace enablement]: https://docs.aws.amazon.com/bedrock/latest/userguide/agents-test.html#trace-events
	EnableTrace *bool

	// Specifies whether to end the session with the agent or not.
	EndSession *bool

	// The prompt text to send the agent.
	InputText *string

	// Contains parameters that specify various attributes of the session. For more
	// information, see [Control session context].
	//
	// [Control session context]: https://docs.aws.amazon.com/bedrock/latest/userguide/agents-session-state.html
	SessionState *types.SessionState

	noSmithyDocumentSerde
}

type InvokeAgentOutput struct {

	// The MIME type of the input data in the request. The default value is
	// application/json .
	//
	// This member is required.
	ContentType *string

	// The unique identifier of the session with the agent.
	//
	// This member is required.
	SessionId *string

	eventStream *InvokeAgentEventStream

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

// GetStream returns the type to interact with the event stream.
func (o *InvokeAgentOutput) GetStream() *InvokeAgentEventStream {
	return o.eventStream
}

func (c *Client) addOperationInvokeAgentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInvokeAgent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInvokeAgent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "InvokeAgent"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addEventStreamInvokeAgentMiddleware(stack, options); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpInvokeAgentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInvokeAgent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInvokeAgent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InvokeAgent",
	}
}

// InvokeAgentEventStream provides the event stream handling for the InvokeAgent operation.
//
// For testing and mocking the event stream this type should be initialized via
// the NewInvokeAgentEventStream constructor function. Using the functional options
// to pass in nested mock behavior.
type InvokeAgentEventStream struct {
	// ResponseStreamReader is the EventStream reader for the ResponseStream events.
	// This value is automatically set by the SDK when the API call is made Use this
	// member when unit testing your code with the SDK to mock out the EventStream
	// Reader.
	//
	// Must not be nil.
	Reader ResponseStreamReader

	done      chan struct{}
	closeOnce sync.Once
	err       *smithysync.OnceErr
}

// NewInvokeAgentEventStream initializes an InvokeAgentEventStream.
// This function should only be used for testing and mocking the InvokeAgentEventStream
// stream within your application.
//
// The Reader member must be set before reading events from the stream.
func NewInvokeAgentEventStream(optFns ...func(*InvokeAgentEventStream)) *InvokeAgentEventStream {
	es := &InvokeAgentEventStream{
		done: make(chan struct{}),
		err:  smithysync.NewOnceErr(),
	}
	for _, fn := range optFns {
		fn(es)
	}
	return es
}

// Events returns a channel to read events from.
func (es *InvokeAgentEventStream) Events() <-chan types.ResponseStream {
	return es.Reader.Events()
}

// Close closes the stream. This will also cause the stream to be closed.
// Close must be called when done using the stream API. Not calling Close
// may result in resource leaks.
//
// Will close the underlying EventStream writer and reader, and no more events can be
// sent or received.
func (es *InvokeAgentEventStream) Close() error {
	es.closeOnce.Do(es.safeClose)
	return es.Err()
}

func (es *InvokeAgentEventStream) safeClose() {
	close(es.done)

	es.Reader.Close()
}

// Err returns any error that occurred while reading or writing EventStream Events
// from the service API's response. Returns nil if there were no errors.
func (es *InvokeAgentEventStream) Err() error {
	if err := es.err.Err(); err != nil {
		return err
	}

	if err := es.Reader.Err(); err != nil {
		return err
	}

	return nil
}

func (es *InvokeAgentEventStream) waitStreamClose() {
	type errorSet interface {
		ErrorSet() <-chan struct{}
	}

	var outputErrCh <-chan struct{}
	if v, ok := es.Reader.(errorSet); ok {
		outputErrCh = v.ErrorSet()
	}
	var outputClosedCh <-chan struct{}
	if v, ok := es.Reader.(interface{ Closed() <-chan struct{} }); ok {
		outputClosedCh = v.Closed()
	}

	select {
	case <-es.done:
	case <-outputErrCh:
		es.err.SetError(es.Reader.Err())
		es.Close()

	case <-outputClosedCh:
		if err := es.Reader.Err(); err != nil {
			es.err.SetError(es.Reader.Err())
		}
		es.Close()

	}
}
