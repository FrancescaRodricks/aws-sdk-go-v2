// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides information regarding the user.
func (c *Client) DescribeUser(ctx context.Context, params *DescribeUserInput, optFns ...func(*Options)) (*DescribeUserOutput, error) {
	if params == nil {
		params = &DescribeUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeUser", params, optFns, c.addOperationDescribeUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUserInput struct {

	// The identifier for the organization under which the user exists.
	//
	// This member is required.
	OrganizationId *string

	// The identifier for the user to be described.
	//
	// The identifier can be the UserId, Username, or email. The following identity
	// formats are available:
	//
	//   - User ID: 12345678-1234-1234-1234-123456789012 or
	//   S-1-1-12-1234567890-123456789-123456789-1234
	//
	//   - Email address: user@domain.tld
	//
	//   - User name: user
	//
	// This member is required.
	UserId *string

	noSmithyDocumentSerde
}

type DescribeUserOutput struct {

	// City where the user is located.
	City *string

	// Company of the user.
	Company *string

	// Country where the user is located.
	Country *string

	// Department of the user.
	Department *string

	// The date and time at which the user was disabled for WorkMail usage, in UNIX
	// epoch time format.
	DisabledDate *time.Time

	// The display name of the user.
	DisplayName *string

	// The email of the user.
	Email *string

	// The date and time at which the user was enabled for WorkMailusage, in UNIX
	// epoch time format.
	EnabledDate *time.Time

	// First name of the user.
	FirstName *string

	// If enabled, the user is hidden from the global address list.
	HiddenFromGlobalAddressList bool

	// Initials of the user.
	Initials *string

	// Job title of the user.
	JobTitle *string

	// Last name of the user.
	LastName *string

	// The date when the mailbox was removed for the user.
	MailboxDeprovisionedDate *time.Time

	// The date when the mailbox was created for the user.
	MailboxProvisionedDate *time.Time

	// The name for the user.
	Name *string

	// Office where the user is located.
	Office *string

	// The state of a user: enabled (registered to WorkMail) or disabled (deregistered
	// or never registered to WorkMail).
	State types.EntityState

	// Street where the user is located.
	Street *string

	// User's contact number.
	Telephone *string

	// The identifier for the described user.
	UserId *string

	// In certain cases, other entities are modeled as users. If interoperability is
	// enabled, resources are imported into WorkMail as users. Because different
	// WorkMail organizations rely on different directory types, administrators can
	// distinguish between an unregistered user (account is disabled and has a user
	// role) and the directory administrators. The values are USER, RESOURCE,
	// SYSTEM_USER, and REMOTE_USER.
	UserRole types.UserRole

	// Zip code of the user.
	ZipCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeUser{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeUser"); err != nil {
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
	if err = addOpDescribeUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeUser",
	}
}
