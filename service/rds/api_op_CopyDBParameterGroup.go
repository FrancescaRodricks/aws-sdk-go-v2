// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Copies the specified DB parameter group.
//
// You can't copy a default DB parameter group. Instead, create a new custom DB
// parameter group, which copies the default parameters and values for the
// specified DB parameter group family.
func (c *Client) CopyDBParameterGroup(ctx context.Context, params *CopyDBParameterGroupInput, optFns ...func(*Options)) (*CopyDBParameterGroupOutput, error) {
	if params == nil {
		params = &CopyDBParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyDBParameterGroup", params, optFns, c.addOperationCopyDBParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyDBParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyDBParameterGroupInput struct {

	// The identifier or ARN for the source DB parameter group. For information about
	// creating an ARN, see [Constructing an ARN for Amazon RDS]in the Amazon RDS User Guide.
	//
	// Constraints:
	//
	//   - Must specify a valid DB parameter group.
	//
	// [Constructing an ARN for Amazon RDS]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.ARN.html#USER_Tagging.ARN.Constructing
	//
	// This member is required.
	SourceDBParameterGroupIdentifier *string

	// A description for the copied DB parameter group.
	//
	// This member is required.
	TargetDBParameterGroupDescription *string

	// The identifier for the copied DB parameter group.
	//
	// Constraints:
	//
	//   - Can't be null, empty, or blank
	//
	//   - Must contain from 1 to 255 letters, numbers, or hyphens
	//
	//   - First character must be a letter
	//
	//   - Can't end with a hyphen or contain two consecutive hyphens
	//
	// Example: my-db-parameter-group
	//
	// This member is required.
	TargetDBParameterGroupIdentifier *string

	// A list of tags. For more information, see [Tagging Amazon RDS Resources] in the Amazon RDS User Guide.
	//
	// [Tagging Amazon RDS Resources]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CopyDBParameterGroupOutput struct {

	// Contains the details of an Amazon RDS DB parameter group.
	//
	// This data type is used as a response element in the DescribeDBParameterGroups
	// action.
	DBParameterGroup *types.DBParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCopyDBParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCopyDBParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCopyDBParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CopyDBParameterGroup"); err != nil {
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
	if err = addOpCopyDBParameterGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopyDBParameterGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCopyDBParameterGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CopyDBParameterGroup",
	}
}
