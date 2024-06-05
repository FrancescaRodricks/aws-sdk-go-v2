// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides temporary, HTTP basic credentials that are associated with a given
// runtime IAM role and used by a cluster with fine-grained access control
// activated. You can use these credentials to connect to cluster endpoints that
// support username and password authentication.
func (c *Client) GetClusterSessionCredentials(ctx context.Context, params *GetClusterSessionCredentialsInput, optFns ...func(*Options)) (*GetClusterSessionCredentialsOutput, error) {
	if params == nil {
		params = &GetClusterSessionCredentialsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetClusterSessionCredentials", params, optFns, c.addOperationGetClusterSessionCredentialsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetClusterSessionCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetClusterSessionCredentialsInput struct {

	// The unique identifier of the cluster.
	//
	// This member is required.
	ClusterId *string

	// The Amazon Resource Name (ARN) of the runtime role for interactive workload
	// submission on the cluster. The runtime role can be a cross-account IAM role. The
	// runtime role ARN is a combination of account ID, role name, and role type using
	// the following format: arn:partition:service:region:account:resource .
	ExecutionRoleArn *string

	noSmithyDocumentSerde
}

type GetClusterSessionCredentialsOutput struct {

	// The credentials that you can use to connect to cluster endpoints that support
	// username and password authentication.
	Credentials types.Credentials

	// The time when the credentials that are returned by the
	// GetClusterSessionCredentials API expire.
	ExpiresAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetClusterSessionCredentialsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetClusterSessionCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetClusterSessionCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetClusterSessionCredentials"); err != nil {
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
	if err = addOpGetClusterSessionCredentialsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetClusterSessionCredentials(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetClusterSessionCredentials(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetClusterSessionCredentials",
	}
}
