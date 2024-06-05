// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmonitor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/networkmonitor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the details about a probe. This action requires both the monitorName
// and probeId parameters. Run ListMonitors to get a list of monitor names. Run
// GetMonitor to get a list of probes and probe IDs.
func (c *Client) GetProbe(ctx context.Context, params *GetProbeInput, optFns ...func(*Options)) (*GetProbeOutput, error) {
	if params == nil {
		params = &GetProbeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetProbe", params, optFns, c.addOperationGetProbeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetProbeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetProbeInput struct {

	// The name of the monitor associated with the probe. Run ListMonitors to get a
	// list of monitor names.
	//
	// This member is required.
	MonitorName *string

	// The ID of the probe to get information about. Run GetMonitor action to get a
	// list of probes and probe IDs for the monitor.
	//
	// This member is required.
	ProbeId *string

	noSmithyDocumentSerde
}

type GetProbeOutput struct {

	// The destination IP address for the monitor. This must be either an IPv4 or IPv6
	// address.
	//
	// This member is required.
	Destination *string

	// The protocol used for the network traffic between the source and destination .
	// This must be either TCP or ICMP .
	//
	// This member is required.
	Protocol types.Protocol

	// The ARN of the probe.
	//
	// This member is required.
	SourceArn *string

	// Indicates whether the IP address is IPV4 or IPV6 .
	AddressFamily types.AddressFamily

	// The time and date that the probe was created.
	CreatedAt *time.Time

	// The port associated with the destination . This is required only if the protocol
	// is TCP and must be a number between 1 and 65536 .
	DestinationPort *int32

	// The time and date that the probe was last modified.
	ModifiedAt *time.Time

	// The size of the packets sent between the source and destination. This must be a
	// number between 56 and 8500 .
	PacketSize *int32

	// The ARN of the probe.
	ProbeArn *string

	// The ID of the probe for which details are returned.
	ProbeId *string

	// The state of the probe.
	State types.ProbeState

	// The list of key-value pairs assigned to the probe.
	Tags map[string]string

	// The ID of the source VPC or subnet.
	VpcId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetProbeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetProbe{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetProbe{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetProbe"); err != nil {
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
	if err = addOpGetProbeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetProbe(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetProbe(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetProbe",
	}
}
