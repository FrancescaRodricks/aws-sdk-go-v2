// Code generated by smithy-go-codegen DO NOT EDIT.

package grafana

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/grafana/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this operation to define the identity provider (IdP) that this workspace
// authenticates users from, using SAML. You can also map SAML assertion attributes
// to workspace user information and define which groups in the assertion attribute
// are to have the Admin and Editor roles in the workspace.
//
// Changes to the authentication method for a workspace may take a few minutes to
// take effect.
func (c *Client) UpdateWorkspaceAuthentication(ctx context.Context, params *UpdateWorkspaceAuthenticationInput, optFns ...func(*Options)) (*UpdateWorkspaceAuthenticationOutput, error) {
	if params == nil {
		params = &UpdateWorkspaceAuthenticationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWorkspaceAuthentication", params, optFns, c.addOperationUpdateWorkspaceAuthenticationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWorkspaceAuthenticationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWorkspaceAuthenticationInput struct {

	// Specifies whether this workspace uses SAML 2.0, IAM Identity Center, or both to
	// authenticate users for using the Grafana console within a workspace. For more
	// information, see [User authentication in Amazon Managed Grafana].
	//
	// [User authentication in Amazon Managed Grafana]: https://docs.aws.amazon.com/grafana/latest/userguide/authentication-in-AMG.html
	//
	// This member is required.
	AuthenticationProviders []types.AuthenticationProviderTypes

	// The ID of the workspace to update the authentication for.
	//
	// This member is required.
	WorkspaceId *string

	// If the workspace uses SAML, use this structure to map SAML assertion attributes
	// to workspace user information and define which groups in the assertion attribute
	// are to have the Admin and Editor roles in the workspace.
	SamlConfiguration *types.SamlConfiguration

	noSmithyDocumentSerde
}

type UpdateWorkspaceAuthenticationOutput struct {

	// A structure that describes the user authentication for this workspace after the
	// update is made.
	//
	// This member is required.
	Authentication *types.AuthenticationDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateWorkspaceAuthenticationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateWorkspaceAuthentication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateWorkspaceAuthentication{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateWorkspaceAuthentication"); err != nil {
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
	if err = addOpUpdateWorkspaceAuthenticationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWorkspaceAuthentication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateWorkspaceAuthentication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateWorkspaceAuthentication",
	}
}
