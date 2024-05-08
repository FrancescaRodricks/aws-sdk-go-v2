// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified member Amazon Web Services account as a delegated
// administrator for the specified Amazon Web Services service.
//
// Deregistering a delegated administrator can have unintended impacts on the
// functionality of the enabled Amazon Web Services service. See the documentation
// for the enabled service before you deregister a delegated administrator so that
// you understand any potential impacts.
//
// You can run this action only for Amazon Web Services services that support this
// feature. For a current list of services that support it, see the column Supports
// Delegated Administrator in the table at [Amazon Web Services Services that you can use with Organizations]in the Organizations User Guide.
//
// This operation can be called only from the organization's management account.
//
// [Amazon Web Services Services that you can use with Organizations]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services_list.html
func (c *Client) DeregisterDelegatedAdministrator(ctx context.Context, params *DeregisterDelegatedAdministratorInput, optFns ...func(*Options)) (*DeregisterDelegatedAdministratorOutput, error) {
	if params == nil {
		params = &DeregisterDelegatedAdministratorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterDelegatedAdministrator", params, optFns, c.addOperationDeregisterDelegatedAdministratorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterDelegatedAdministratorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterDelegatedAdministratorInput struct {

	// The account ID number of the member account in the organization that you want
	// to deregister as a delegated administrator.
	//
	// This member is required.
	AccountId *string

	// The service principal name of an Amazon Web Services service for which the
	// account is a delegated administrator.
	//
	// Delegated administrator privileges are revoked for only the specified Amazon
	// Web Services service from the member account. If the specified service is the
	// only service for which the member account is a delegated administrator, the
	// operation also revokes Organizations read action permissions.
	//
	// This member is required.
	ServicePrincipal *string

	noSmithyDocumentSerde
}

type DeregisterDelegatedAdministratorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterDelegatedAdministratorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterDelegatedAdministrator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterDelegatedAdministrator{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeregisterDelegatedAdministrator"); err != nil {
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
	if err = addOpDeregisterDelegatedAdministratorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterDelegatedAdministrator(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterDelegatedAdministrator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeregisterDelegatedAdministrator",
	}
}
