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

// Allows or restricts mirroring network services.
//
// By default, Amazon DNS network services are not eligible for Traffic Mirror.
// Use AddNetworkServices to add network services to a Traffic Mirror filter. When
// a network service is added to the Traffic Mirror filter, all traffic related to
// that network service will be mirrored. When you no longer want to mirror network
// services, use RemoveNetworkServices to remove the network services from the
// Traffic Mirror filter.
func (c *Client) ModifyTrafficMirrorFilterNetworkServices(ctx context.Context, params *ModifyTrafficMirrorFilterNetworkServicesInput, optFns ...func(*Options)) (*ModifyTrafficMirrorFilterNetworkServicesOutput, error) {
	if params == nil {
		params = &ModifyTrafficMirrorFilterNetworkServicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyTrafficMirrorFilterNetworkServices", params, optFns, c.addOperationModifyTrafficMirrorFilterNetworkServicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyTrafficMirrorFilterNetworkServicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyTrafficMirrorFilterNetworkServicesInput struct {

	// The ID of the Traffic Mirror filter.
	//
	// This member is required.
	TrafficMirrorFilterId *string

	// The network service, for example Amazon DNS, that you want to mirror.
	AddNetworkServices []types.TrafficMirrorNetworkService

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The network service, for example Amazon DNS, that you no longer want to mirror.
	RemoveNetworkServices []types.TrafficMirrorNetworkService

	noSmithyDocumentSerde
}

type ModifyTrafficMirrorFilterNetworkServicesOutput struct {

	// The Traffic Mirror filter that the network service is associated with.
	TrafficMirrorFilter *types.TrafficMirrorFilter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyTrafficMirrorFilterNetworkServicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyTrafficMirrorFilterNetworkServices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyTrafficMirrorFilterNetworkServices{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyTrafficMirrorFilterNetworkServices"); err != nil {
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
	if err = addOpModifyTrafficMirrorFilterNetworkServicesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyTrafficMirrorFilterNetworkServices(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyTrafficMirrorFilterNetworkServices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyTrafficMirrorFilterNetworkServices",
	}
}
