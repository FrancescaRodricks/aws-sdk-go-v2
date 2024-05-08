// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Creates an uninitialized HSM instance.
//
// There is an upfront fee charged for each HSM instance that you create with the
// CreateHsm operation. If you accidentally provision an HSM and want to request a
// refund, delete the instance using the DeleteHsmoperation, go to the [AWS Support Center], create a new case,
// and select Account and Billing Support.
//
// It can take up to 20 minutes to create and provision an HSM. You can monitor
// the status of the HSM with the DescribeHsmoperation. The HSM is ready to be initialized
// when the status changes to RUNNING .
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS Support Center]: https://console.aws.amazon.com/support/home
func (c *Client) CreateHsm(ctx context.Context, params *CreateHsmInput, optFns ...func(*Options)) (*CreateHsmOutput, error) {
	if params == nil {
		params = &CreateHsmInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHsm", params, optFns, c.addOperationCreateHsmMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHsmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the CreateHsm operation.
type CreateHsmInput struct {

	// The ARN of an IAM role to enable the AWS CloudHSM service to allocate an ENI on
	// your behalf.
	//
	// This member is required.
	IamRoleArn *string

	// The SSH public key to install on the HSM.
	//
	// This member is required.
	SshKey *string

	// The identifier of the subnet in your VPC in which to place the HSM.
	//
	// This member is required.
	SubnetId *string

	// Specifies the type of subscription for the HSM.
	//
	//   - PRODUCTION - The HSM is being used in a production environment.
	//
	//   - TRIAL - The HSM is being used in a product trial.
	//
	// This member is required.
	SubscriptionType types.SubscriptionType

	// A user-defined token to ensure idempotence. Subsequent calls to this operation
	// with the same token will be ignored.
	ClientToken *string

	// The IP address to assign to the HSM's ENI.
	//
	// If an IP address is not specified, an IP address will be randomly chosen from
	// the CIDR range of the subnet.
	EniIp *string

	// The external ID from IamRoleArn , if present.
	ExternalId *string

	// The IP address for the syslog monitoring server. The AWS CloudHSM service only
	// supports one syslog monitoring server.
	SyslogIp *string

	noSmithyDocumentSerde
}

// Contains the output of the CreateHsm operation.
type CreateHsmOutput struct {

	// The ARN of the HSM.
	HsmArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateHsmMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateHsm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateHsm{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateHsm"); err != nil {
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
	if err = addOpCreateHsmValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHsm(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateHsm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateHsm",
	}
}
