// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/detective/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified member accounts from the behavior graph. The removed
// accounts no longer contribute data to the behavior graph. This operation can
// only be called by the administrator account for the behavior graph.
//
// For invited accounts, the removed accounts are deleted from the list of
// accounts in the behavior graph. To restore the account, the administrator
// account must send another invitation.
//
// For organization accounts in the organization behavior graph, the Detective
// administrator account can always enable the organization account again.
// Organization accounts that are not enabled as member accounts are not included
// in the ListMembers results for the organization behavior graph.
//
// An administrator account cannot use DeleteMembers to remove their own account
// from the behavior graph. To disable a behavior graph, the administrator account
// uses the DeleteGraph API method.
func (c *Client) DeleteMembers(ctx context.Context, params *DeleteMembersInput, optFns ...func(*Options)) (*DeleteMembersOutput, error) {
	if params == nil {
		params = &DeleteMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteMembers", params, optFns, c.addOperationDeleteMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteMembersInput struct {

	// The list of Amazon Web Services account identifiers for the member accounts to
	// remove from the behavior graph. You can remove up to 50 member accounts at a
	// time.
	//
	// This member is required.
	AccountIds []string

	// The ARN of the behavior graph to remove members from.
	//
	// This member is required.
	GraphArn *string

	noSmithyDocumentSerde
}

type DeleteMembersOutput struct {

	// The list of Amazon Web Services account identifiers for the member accounts
	// that Detective successfully removed from the behavior graph.
	AccountIds []string

	// The list of member accounts that Detective was not able to remove from the
	// behavior graph. For each member account, provides the reason that the deletion
	// could not be processed.
	UnprocessedAccounts []types.UnprocessedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteMembers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteMembers"); err != nil {
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
	if err = addOpDeleteMembersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMembers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteMembers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteMembers",
	}
}
