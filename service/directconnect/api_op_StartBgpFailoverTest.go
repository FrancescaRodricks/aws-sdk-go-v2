// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts the virtual interface failover test that verifies your configuration
// meets your resiliency requirements by placing the BGP peering session in the
// DOWN state. You can then send traffic to verify that there are no outages.
//
// You can run the test on public, private, transit, and hosted virtual interfaces.
//
// You can use [ListVirtualInterfaceTestHistory] to view the virtual interface test history.
//
// If you need to stop the test before the test interval completes, use [StopBgpFailoverTest].
//
// [ListVirtualInterfaceTestHistory]: https://docs.aws.amazon.com/directconnect/latest/APIReference/API_ListVirtualInterfaceTestHistory.html
// [StopBgpFailoverTest]: https://docs.aws.amazon.com/directconnect/latest/APIReference/API_StopBgpFailoverTest.html
func (c *Client) StartBgpFailoverTest(ctx context.Context, params *StartBgpFailoverTestInput, optFns ...func(*Options)) (*StartBgpFailoverTestOutput, error) {
	if params == nil {
		params = &StartBgpFailoverTestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartBgpFailoverTest", params, optFns, c.addOperationStartBgpFailoverTestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartBgpFailoverTestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartBgpFailoverTestInput struct {

	// The ID of the virtual interface you want to test.
	//
	// This member is required.
	VirtualInterfaceId *string

	// The BGP peers to place in the DOWN state.
	BgpPeers []string

	// The time in minutes that the virtual interface failover test will last.
	//
	// Maximum value: 4,320 minutes (72 hours).
	//
	// Default: 180 minutes (3 hours).
	TestDurationInMinutes *int32

	noSmithyDocumentSerde
}

type StartBgpFailoverTestOutput struct {

	// Information about the virtual interface failover test.
	VirtualInterfaceTest *types.VirtualInterfaceTestHistory

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartBgpFailoverTestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartBgpFailoverTest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartBgpFailoverTest{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartBgpFailoverTest"); err != nil {
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
	if err = addOpStartBgpFailoverTestValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartBgpFailoverTest(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartBgpFailoverTest(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartBgpFailoverTest",
	}
}
