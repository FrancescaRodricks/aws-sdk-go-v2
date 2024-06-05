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

// You can call DescribeValidDBInstanceModifications to learn what modifications
// you can make to your DB instance. You can use this information when you call
// ModifyDBInstance .
//
// This command doesn't apply to RDS Custom.
func (c *Client) DescribeValidDBInstanceModifications(ctx context.Context, params *DescribeValidDBInstanceModificationsInput, optFns ...func(*Options)) (*DescribeValidDBInstanceModificationsOutput, error) {
	if params == nil {
		params = &DescribeValidDBInstanceModificationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeValidDBInstanceModifications", params, optFns, c.addOperationDescribeValidDBInstanceModificationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeValidDBInstanceModificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeValidDBInstanceModificationsInput struct {

	// The customer identifier or the ARN of your DB instance.
	//
	// This member is required.
	DBInstanceIdentifier *string

	noSmithyDocumentSerde
}

type DescribeValidDBInstanceModificationsOutput struct {

	// Information about valid modifications that you can make to your DB instance.
	// Contains the result of a successful call to the
	// DescribeValidDBInstanceModifications action. You can use this information when
	// you call ModifyDBInstance .
	ValidDBInstanceModificationsMessage *types.ValidDBInstanceModificationsMessage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeValidDBInstanceModificationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeValidDBInstanceModifications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeValidDBInstanceModifications{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeValidDBInstanceModifications"); err != nil {
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
	if err = addOpDescribeValidDBInstanceModificationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeValidDBInstanceModifications(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeValidDBInstanceModifications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeValidDBInstanceModifications",
	}
}
