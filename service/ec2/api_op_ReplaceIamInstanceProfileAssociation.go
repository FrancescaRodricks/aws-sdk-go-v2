// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Replaces an IAM instance profile for the specified running instance. You can
// use this action to change the IAM instance profile that's associated with an
// instance without having to disassociate the existing IAM instance profile first.
//
// Use DescribeIamInstanceProfileAssociations to get the association ID.
func (c *Client) ReplaceIamInstanceProfileAssociation(ctx context.Context, params *ReplaceIamInstanceProfileAssociationInput, optFns ...func(*Options)) (*ReplaceIamInstanceProfileAssociationOutput, error) {
	if params == nil {
		params = &ReplaceIamInstanceProfileAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReplaceIamInstanceProfileAssociation", params, optFns, c.addOperationReplaceIamInstanceProfileAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReplaceIamInstanceProfileAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReplaceIamInstanceProfileAssociationInput struct {

	// The ID of the existing IAM instance profile association.
	//
	// This member is required.
	AssociationId *string

	// The IAM instance profile.
	//
	// This member is required.
	IamInstanceProfile *types.IamInstanceProfileSpecification

	noSmithyDocumentSerde
}

type ReplaceIamInstanceProfileAssociationOutput struct {

	// Information about the IAM instance profile association.
	IamInstanceProfileAssociation *types.IamInstanceProfileAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationReplaceIamInstanceProfileAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpReplaceIamInstanceProfileAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpReplaceIamInstanceProfileAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ReplaceIamInstanceProfileAssociation"); err != nil {
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
	if err = addOpReplaceIamInstanceProfileAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReplaceIamInstanceProfileAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opReplaceIamInstanceProfileAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ReplaceIamInstanceProfileAssociation",
	}
}
