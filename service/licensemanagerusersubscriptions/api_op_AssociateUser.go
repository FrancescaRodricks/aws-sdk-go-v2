// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanagerusersubscriptions

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/licensemanagerusersubscriptions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates the user to an EC2 instance to utilize user-based subscriptions.
//
// Your estimated bill for charges on the number of users and related costs will
// take 48 hours to appear for billing periods that haven't closed (marked as
// Pending billing status) in Amazon Web Services Billing. For more information,
// see [Viewing your monthly charges]in the Amazon Web Services Billing User Guide.
//
// [Viewing your monthly charges]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/invoice.html
func (c *Client) AssociateUser(ctx context.Context, params *AssociateUserInput, optFns ...func(*Options)) (*AssociateUserOutput, error) {
	if params == nil {
		params = &AssociateUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateUser", params, optFns, c.addOperationAssociateUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateUserInput struct {

	// The identity provider of the user.
	//
	// This member is required.
	IdentityProvider types.IdentityProvider

	// The ID of the EC2 instance, which provides user-based subscriptions.
	//
	// This member is required.
	InstanceId *string

	// The user name from the identity provider for the user.
	//
	// This member is required.
	Username *string

	// The domain name of the user.
	Domain *string

	noSmithyDocumentSerde
}

type AssociateUserOutput struct {

	// Metadata that describes the associate user operation.
	//
	// This member is required.
	InstanceUserSummary *types.InstanceUserSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateUser{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateUser"); err != nil {
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
	if err = addOpAssociateUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateUser",
	}
}
