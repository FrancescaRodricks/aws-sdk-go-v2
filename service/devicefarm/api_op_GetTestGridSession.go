// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A session is an instance of a browser created through a RemoteWebDriver with
// the URL from CreateTestGridUrlResult$url. You can use the following to look up sessions:
//
//   - The session ARN (GetTestGridSessionRequest$sessionArn ).
//
//   - The project ARN and a session ID (GetTestGridSessionRequest$projectArn and GetTestGridSessionRequest$sessionId).
func (c *Client) GetTestGridSession(ctx context.Context, params *GetTestGridSessionInput, optFns ...func(*Options)) (*GetTestGridSessionOutput, error) {
	if params == nil {
		params = &GetTestGridSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTestGridSession", params, optFns, c.addOperationGetTestGridSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTestGridSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTestGridSessionInput struct {

	// The ARN for the project that this session belongs to. See CreateTestGridProject and ListTestGridProjects.
	ProjectArn *string

	// An ARN that uniquely identifies a TestGridSession.
	SessionArn *string

	// An ID associated with this session.
	SessionId *string

	noSmithyDocumentSerde
}

type GetTestGridSessionOutput struct {

	// The TestGridSession that was requested.
	TestGridSession *types.TestGridSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTestGridSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetTestGridSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTestGridSession{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetTestGridSession"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTestGridSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTestGridSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetTestGridSession",
	}
}
