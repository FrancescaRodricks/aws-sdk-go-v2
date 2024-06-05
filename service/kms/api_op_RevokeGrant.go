// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified grant. You revoke a grant to terminate the permissions
// that the grant allows. For more information, see [Retiring and revoking grants]in the Key Management Service
// Developer Guide .
//
// When you create, retire, or revoke a grant, there might be a brief delay,
// usually less than five minutes, until the grant is available throughout KMS.
// This state is known as eventual consistency. For details, see [Eventual consistency]in the Key
// Management Service Developer Guide .
//
// For detailed information about grants, including grant terminology, see [Grants in KMS] in the
// Key Management Service Developer Guide . For examples of working with grants in
// several programming languages, see [Programming grants].
//
// Cross-account use: Yes. To perform this operation on a KMS key in a different
// Amazon Web Services account, specify the key ARN in the value of the KeyId
// parameter.
//
// Required permissions: [kms:RevokeGrant] (key policy).
//
// Related operations:
//
// # CreateGrant
//
// # ListGrants
//
// # ListRetirableGrants
//
// # RetireGrant
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#terms-eventual-consistency
// [Programming grants]: https://docs.aws.amazon.com/kms/latest/developerguide/programming-grants.html
// [kms:RevokeGrant]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Retiring and revoking grants]: https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#grant-delete
// [Grants in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/programming-eventual-consistency.html
func (c *Client) RevokeGrant(ctx context.Context, params *RevokeGrantInput, optFns ...func(*Options)) (*RevokeGrantOutput, error) {
	if params == nil {
		params = &RevokeGrantInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RevokeGrant", params, optFns, c.addOperationRevokeGrantMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RevokeGrantOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RevokeGrantInput struct {

	// Identifies the grant to revoke. To get the grant ID, use CreateGrant, ListGrants, or ListRetirableGrants.
	//
	// This member is required.
	GrantId *string

	// A unique identifier for the KMS key associated with the grant. To get the key
	// ID and key ARN for a KMS key, use ListKeysor DescribeKey.
	//
	// Specify the key ID or key ARN of the KMS key. To specify a KMS key in a
	// different Amazon Web Services account, you must use the key ARN.
	//
	// For example:
	//
	//   - Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//   - Key ARN:
	//   arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey.
	//
	// This member is required.
	KeyId *string

	// Checks if your request will succeed. DryRun is an optional parameter.
	//
	// To learn more about how to use this parameter, see [Testing your KMS API calls] in the Key Management
	// Service Developer Guide.
	//
	// [Testing your KMS API calls]: https://docs.aws.amazon.com/kms/latest/developerguide/programming-dryrun.html
	DryRun *bool

	noSmithyDocumentSerde
}

type RevokeGrantOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRevokeGrantMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRevokeGrant{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRevokeGrant{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RevokeGrant"); err != nil {
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
	if err = addOpRevokeGrantValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRevokeGrant(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRevokeGrant(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RevokeGrant",
	}
}
