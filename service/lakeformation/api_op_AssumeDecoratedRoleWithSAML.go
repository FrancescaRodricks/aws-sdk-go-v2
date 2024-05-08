// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Allows a caller to assume an IAM role decorated as the SAML user specified in
// the SAML assertion included in the request. This decoration allows Lake
// Formation to enforce access policies against the SAML users and groups. This API
// operation requires SAML federation setup in the caller’s account as it can only
// be called with valid SAML assertions. Lake Formation does not scope down the
// permission of the assumed role. All permissions attached to the role via the
// SAML federation setup will be included in the role session.
//
// This decorated role is expected to access data in Amazon S3 by getting
// temporary access from Lake Formation which is authorized via the virtual API
// GetDataAccess . Therefore, all SAML roles that can be assumed via
// AssumeDecoratedRoleWithSAML must at a minimum include
// lakeformation:GetDataAccess in their role policies. A typical IAM policy
// attached to such a role would look as follows:
func (c *Client) AssumeDecoratedRoleWithSAML(ctx context.Context, params *AssumeDecoratedRoleWithSAMLInput, optFns ...func(*Options)) (*AssumeDecoratedRoleWithSAMLOutput, error) {
	if params == nil {
		params = &AssumeDecoratedRoleWithSAMLInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssumeDecoratedRoleWithSAML", params, optFns, c.addOperationAssumeDecoratedRoleWithSAMLMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssumeDecoratedRoleWithSAMLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssumeDecoratedRoleWithSAMLInput struct {

	// The Amazon Resource Name (ARN) of the SAML provider in IAM that describes the
	// IdP.
	//
	// This member is required.
	PrincipalArn *string

	// The role that represents an IAM principal whose scope down policy allows it to
	// call credential vending APIs such as GetTemporaryTableCredentials . The caller
	// must also have iam:PassRole permission on this role.
	//
	// This member is required.
	RoleArn *string

	// A SAML assertion consisting of an assertion statement for the user who needs
	// temporary credentials. This must match the SAML assertion that was issued to
	// IAM. This must be Base64 encoded.
	//
	// This member is required.
	SAMLAssertion *string

	// The time period, between 900 and 43,200 seconds, for the timeout of the
	// temporary credentials.
	DurationSeconds *int32

	noSmithyDocumentSerde
}

type AssumeDecoratedRoleWithSAMLOutput struct {

	// The access key ID for the temporary credentials. (The access key consists of an
	// access key ID and a secret key).
	AccessKeyId *string

	// The date and time when the temporary credentials expire.
	Expiration *time.Time

	// The secret key for the temporary credentials. (The access key consists of an
	// access key ID and a secret key).
	SecretAccessKey *string

	// The session token for the temporary credentials.
	SessionToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssumeDecoratedRoleWithSAMLMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssumeDecoratedRoleWithSAML{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssumeDecoratedRoleWithSAML{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssumeDecoratedRoleWithSAML"); err != nil {
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
	if err = addOpAssumeDecoratedRoleWithSAMLValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssumeDecoratedRoleWithSAML(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssumeDecoratedRoleWithSAML(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssumeDecoratedRoleWithSAML",
	}
}
