// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes an existing Amazon Redshift IAM Identity Center application.
func (c *Client) ModifyRedshiftIdcApplication(ctx context.Context, params *ModifyRedshiftIdcApplicationInput, optFns ...func(*Options)) (*ModifyRedshiftIdcApplicationOutput, error) {
	if params == nil {
		params = &ModifyRedshiftIdcApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyRedshiftIdcApplication", params, optFns, c.addOperationModifyRedshiftIdcApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyRedshiftIdcApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyRedshiftIdcApplicationInput struct {

	// The ARN for the Redshift application that integrates with IAM Identity Center.
	//
	// This member is required.
	RedshiftIdcApplicationArn *string

	// The authorized token issuer list for the Amazon Redshift IAM Identity Center
	// application to change.
	AuthorizedTokenIssuerList []types.AuthorizedTokenIssuer

	// The IAM role ARN associated with the Amazon Redshift IAM Identity Center
	// application to change. It has the required permissions to be assumed and invoke
	// the IDC Identity Center API.
	IamRoleArn *string

	// The display name for the Amazon Redshift IAM Identity Center application to
	// change. It appears on the console.
	IdcDisplayName *string

	// The namespace for the Amazon Redshift IAM Identity Center application to
	// change. It determines which managed application verifies the connection token.
	IdentityNamespace *string

	// A collection of service integrations associated with the application.
	ServiceIntegrations []types.ServiceIntegrationsUnion

	noSmithyDocumentSerde
}

type ModifyRedshiftIdcApplicationOutput struct {

	// Contains properties for the Redshift IDC application.
	RedshiftIdcApplication *types.RedshiftIdcApplication

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyRedshiftIdcApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyRedshiftIdcApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyRedshiftIdcApplication{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyRedshiftIdcApplication"); err != nil {
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
	if err = addOpModifyRedshiftIdcApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyRedshiftIdcApplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyRedshiftIdcApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyRedshiftIdcApplication",
	}
}
