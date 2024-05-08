// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes read and write permissions for a dashboard.
func (c *Client) DescribeDashboardPermissions(ctx context.Context, params *DescribeDashboardPermissionsInput, optFns ...func(*Options)) (*DescribeDashboardPermissionsOutput, error) {
	if params == nil {
		params = &DescribeDashboardPermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDashboardPermissions", params, optFns, c.addOperationDescribeDashboardPermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDashboardPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDashboardPermissionsInput struct {

	// The ID of the Amazon Web Services account that contains the dashboard that
	// you're describing permissions for.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the dashboard, also added to the IAM policy.
	//
	// This member is required.
	DashboardId *string

	noSmithyDocumentSerde
}

type DescribeDashboardPermissionsOutput struct {

	// The Amazon Resource Name (ARN) of the dashboard.
	DashboardArn *string

	// The ID for the dashboard.
	DashboardId *string

	// A structure that contains the configuration of a shareable link that grants
	// access to the dashboard. Your users can use the link to view and interact with
	// the dashboard, if the dashboard has been shared with them. For more information
	// about sharing dashboards, see [Sharing Dashboards].
	//
	// [Sharing Dashboards]: https://docs.aws.amazon.com/quicksight/latest/user/sharing-a-dashboard.html
	LinkSharingConfiguration *types.LinkSharingConfiguration

	// A structure that contains the permissions for the dashboard.
	Permissions []types.ResourcePermission

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDashboardPermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDashboardPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDashboardPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDashboardPermissions"); err != nil {
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
	if err = addOpDescribeDashboardPermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDashboardPermissions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDashboardPermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDashboardPermissions",
	}
}
