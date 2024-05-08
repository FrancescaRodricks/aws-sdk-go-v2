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

// Updates a monitor probe. This action requires both the monitorName and probeId
// parameters. Run ListMonitors to get a list of monitor names. Run GetMonitor to
// get a list of probes and probe IDs.
//
// You can update the following para create a monitor with probes using this
// command. For each probe, you define the following:
//
//   - state —The state of the probe.
//
//   - destination — The target destination IP address for the probe.
//
//   - destinationPort —Required only if the protocol is TCP .
//
//   - protocol —The communication protocol between the source and destination.
//     This will be either TCP or ICMP .
//
//   - packetSize —The size of the packets. This must be a number between 56 and
//     8500 .
//
//   - (Optional) tags —Key-value pairs created and assigned to the probe.
func (c *Client) UpdateProbe(ctx context.Context, params *UpdateProbeInput, optFns ...func(*Options)) (*UpdateProbeOutput, error) {
	if params == nil {
		params = &UpdateProbeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateProbe", params, optFns, c.addOperationUpdateProbeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateProbeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateProbeInput struct {

	// The name of the monitor that the probe was updated for.
	//
	// This member is required.
	MonitorName *string

	// The ID of the probe to update.
	//
	// This member is required.
	ProbeId *string

	// The updated IP address for the probe destination. This must be either an IPv4
	// or IPv6 address.
	Destination *string

	// The updated port for the probe destination. This is required only if the
	// protocol is TCP and must be a number between 1 and 65536 .
	DestinationPort *int32

	// he updated packets size for network traffic between the source and destination.
	// This must be a number between 56 and 8500 .
	PacketSize *int32

	// The updated network protocol for the destination. This can be either TCP or ICMP
	// . If the protocol is TCP , then port is also required.
	Protocol types.Protocol

	// The state of the probe update.
	State types.ProbeState

	noSmithyDocumentSerde
}

type UpdateProbeOutput struct {

	// The updated destination IP address for the probe.
	//
	// This member is required.
	Destination *string

	// The updated protocol for the probe.
	//
	// This member is required.
	Protocol types.Protocol

	// The updated ARN of the source subnet.
	//
	// This member is required.
	SourceArn *string

	// The updated IP address family. This must be either IPV4 or IPV6 .
	AddressFamily types.AddressFamily

	// The time and date that the probe was created.
	CreatedAt *time.Time

	// The updated destination port. This must be a number between 1 and 65536 .
	DestinationPort *int32

	// The time and date that the probe was last updated.
	ModifiedAt *time.Time

	// The updated packet size for the probe.
	PacketSize *int32

	// The updated ARN of the probe.
	ProbeArn *string

	// The updated ID of the probe.
	ProbeId *string

	// The state of the updated probe.
	State types.ProbeState

	// Update tags for a probe.
	Tags map[string]string

	// The updated ID of the source VPC subnet ID.
	VpcId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateProbeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateProbe{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateProbe{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateProbe"); err != nil {
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
	if err = addOpUpdateProbeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateProbe(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateProbe(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateProbe",
	}
}
