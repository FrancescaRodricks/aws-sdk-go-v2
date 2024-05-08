// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the configuration of your Amazon Lightsail container service, such as
// its power, scale, and public domain names.
func (c *Client) UpdateContainerService(ctx context.Context, params *UpdateContainerServiceInput, optFns ...func(*Options)) (*UpdateContainerServiceOutput, error) {
	if params == nil {
		params = &UpdateContainerServiceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateContainerService", params, optFns, c.addOperationUpdateContainerServiceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateContainerServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateContainerServiceInput struct {

	// The name of the container service to update.
	//
	// This member is required.
	ServiceName *string

	// A Boolean value to indicate whether the container service is disabled.
	IsDisabled *bool

	// The power for the container service.
	//
	// The power specifies the amount of memory, vCPUs, and base monthly cost of each
	// node of the container service. The power and scale of a container service makes
	// up its configured capacity. To determine the monthly price of your container
	// service, multiply the base price of the power with the scale (the number of
	// nodes) of the service.
	//
	// Use the GetContainerServicePowers action to view the specifications of each
	// power option.
	Power types.ContainerServicePowerName

	// An object to describe the configuration for the container service to access
	// private container image repositories, such as Amazon Elastic Container Registry
	// (Amazon ECR) private repositories.
	//
	// For more information, see [Configuring access to an Amazon ECR private repository for an Amazon Lightsail container service] in the Amazon Lightsail Developer Guide.
	//
	// [Configuring access to an Amazon ECR private repository for an Amazon Lightsail container service]: https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-container-service-ecr-private-repo-access
	PrivateRegistryAccess *types.PrivateRegistryAccessRequest

	// The public domain names to use with the container service, such as example.com
	// and www.example.com .
	//
	// You can specify up to four public domain names for a container service. The
	// domain names that you specify are used when you create a deployment with a
	// container configured as the public endpoint of your container service.
	//
	// If you don't specify public domain names, then you can use the default domain
	// of the container service.
	//
	// You must create and validate an SSL/TLS certificate before you can use public
	// domain names with your container service. Use the CreateCertificate action to
	// create a certificate for the public domain names you want to use with your
	// container service.
	//
	// You can specify public domain names using a string to array map as shown in the
	// example later on this page.
	PublicDomainNames map[string][]string

	// The scale for the container service.
	//
	// The scale specifies the allocated compute nodes of the container service. The
	// power and scale of a container service makes up its configured capacity. To
	// determine the monthly price of your container service, multiply the base price
	// of the power with the scale (the number of nodes) of the service.
	Scale *int32

	noSmithyDocumentSerde
}

type UpdateContainerServiceOutput struct {

	// An object that describes a container service.
	ContainerService *types.ContainerService

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateContainerServiceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateContainerService{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateContainerService{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateContainerService"); err != nil {
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
	if err = addOpUpdateContainerServiceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateContainerService(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateContainerService(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateContainerService",
	}
}
