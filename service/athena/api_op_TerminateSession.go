// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Terminates an active session. A TerminateSession call on a session that is
// already inactive (for example, in a FAILED , TERMINATED or TERMINATING state)
// succeeds but has no effect. Calculations running in the session when
// TerminateSession is called are forcefully stopped, but may display as FAILED
// instead of STOPPED .
func (c *Client) TerminateSession(ctx context.Context, params *TerminateSessionInput, optFns ...func(*Options)) (*TerminateSessionOutput, error) {
	if params == nil {
		params = &TerminateSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TerminateSession", params, optFns, c.addOperationTerminateSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TerminateSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TerminateSessionInput struct {

	// The session ID.
	//
	// This member is required.
	SessionId *string

	noSmithyDocumentSerde
}

type TerminateSessionOutput struct {

	// The state of the session. A description of each state follows.
	//
	// CREATING - The session is being started, including acquiring resources.
	//
	// CREATED - The session has been started.
	//
	// IDLE - The session is able to accept a calculation.
	//
	// BUSY - The session is processing another task and is unable to accept a
	// calculation.
	//
	// TERMINATING - The session is in the process of shutting down.
	//
	// TERMINATED - The session and its resources are no longer running.
	//
	// DEGRADED - The session has no healthy coordinators.
	//
	// FAILED - Due to a failure, the session and its resources are no longer running.
	State types.SessionState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTerminateSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTerminateSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTerminateSession{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "TerminateSession"); err != nil {
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
	if err = addOpTerminateSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTerminateSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTerminateSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "TerminateSession",
	}
}
