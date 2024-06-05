// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a custom set of DHCP options. After you create a DHCP option set, you
// associate it with a VPC. After you associate a DHCP option set with a VPC, all
// existing and newly launched instances in the VPC use this set of DHCP options.
//
// The following are the individual DHCP options you can specify. For more
// information, see [DHCP options sets]in the Amazon VPC User Guide.
//
//   - domain-name - If you're using AmazonProvidedDNS in us-east-1 , specify
//     ec2.internal . If you're using AmazonProvidedDNS in any other Region, specify
//     region.compute.internal . Otherwise, specify a custom domain name. This value
//     is used to complete unqualified DNS hostnames.
//
// Some Linux operating systems accept multiple domain names separated by spaces.
//
//	However, Windows and other Linux operating systems treat the value as a single
//	domain, which results in unexpected behavior. If your DHCP option set is
//	associated with a VPC that has instances running operating systems that treat
//	the value as a single domain, specify only one domain name.
//
//	- domain-name-servers - The IP addresses of up to four DNS servers, or
//	AmazonProvidedDNS. To specify multiple domain name servers in a single
//	parameter, separate the IP addresses using commas. To have your instances
//	receive custom DNS hostnames as specified in domain-name , you must specify a
//	custom DNS server.
//
//	- ntp-servers - The IP addresses of up to eight Network Time Protocol (NTP)
//	servers (four IPv4 addresses and four IPv6 addresses).
//
//	- netbios-name-servers - The IP addresses of up to four NetBIOS name servers.
//
//	- netbios-node-type - The NetBIOS node type (1, 2, 4, or 8). We recommend that
//	you specify 2. Broadcast and multicast are not supported. For more information
//	about NetBIOS node types, see [RFC 2132].
//
//	- ipv6-address-preferred-lease-time - A value (in seconds, minutes, hours, or
//	years) for how frequently a running instance with an IPv6 assigned to it goes
//	through DHCPv6 lease renewal. Acceptable values are between 140 and 2147483647
//	seconds (approximately 68 years). If no value is entered, the default lease time
//	is 140 seconds. If you use long-term addressing for EC2 instances, you can
//	increase the lease time and avoid frequent lease renewal requests. Lease renewal
//	typically occurs when half of the lease time has elapsed.
//
// [DHCP options sets]: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_DHCP_Options.html
//
// [RFC 2132]: http://www.ietf.org/rfc/rfc2132.txt
func (c *Client) CreateDhcpOptions(ctx context.Context, params *CreateDhcpOptionsInput, optFns ...func(*Options)) (*CreateDhcpOptionsOutput, error) {
	if params == nil {
		params = &CreateDhcpOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDhcpOptions", params, optFns, c.addOperationCreateDhcpOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDhcpOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDhcpOptionsInput struct {

	// A DHCP configuration option.
	//
	// This member is required.
	DhcpConfigurations []types.NewDhcpConfiguration

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The tags to assign to the DHCP option.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreateDhcpOptionsOutput struct {

	// A set of DHCP options.
	DhcpOptions *types.DhcpOptions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDhcpOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateDhcpOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateDhcpOptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDhcpOptions"); err != nil {
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
	if err = addOpCreateDhcpOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDhcpOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDhcpOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDhcpOptions",
	}
}
