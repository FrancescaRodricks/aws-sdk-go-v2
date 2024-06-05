// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	The CreateAdditionalAssignmentsForHIT operation increases the maximum number
//
// of assignments of an existing HIT.
//
// To extend the maximum number of assignments, specify the number of additional
// assignments.
//
//   - HITs created with fewer than 10 assignments cannot be extended to have 10
//     or more assignments. Attempting to add assignments in a way that brings the
//     total number of assignments for a HIT from fewer than 10 assignments to 10 or
//     more assignments will result in an
//     AWS.MechanicalTurk.InvalidMaximumAssignmentsIncrease exception.
//
//   - HITs that were created before July 22, 2015 cannot be extended. Attempting
//     to extend HITs that were created before July 22, 2015 will result in an
//     AWS.MechanicalTurk.HITTooOldForExtension exception.
func (c *Client) CreateAdditionalAssignmentsForHIT(ctx context.Context, params *CreateAdditionalAssignmentsForHITInput, optFns ...func(*Options)) (*CreateAdditionalAssignmentsForHITOutput, error) {
	if params == nil {
		params = &CreateAdditionalAssignmentsForHITInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAdditionalAssignmentsForHIT", params, optFns, c.addOperationCreateAdditionalAssignmentsForHITMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAdditionalAssignmentsForHITOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAdditionalAssignmentsForHITInput struct {

	// The ID of the HIT to extend.
	//
	// This member is required.
	HITId *string

	// The number of additional assignments to request for this HIT.
	//
	// This member is required.
	NumberOfAdditionalAssignments *int32

	//  A unique identifier for this request, which allows you to retry the call on
	// error without extending the HIT multiple times. This is useful in cases such as
	// network timeouts where it is unclear whether or not the call succeeded on the
	// server. If the extend HIT already exists in the system from a previous call
	// using the same UniqueRequestToken , subsequent calls will return an error with a
	// message containing the request ID.
	UniqueRequestToken *string

	noSmithyDocumentSerde
}

type CreateAdditionalAssignmentsForHITOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAdditionalAssignmentsForHITMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAdditionalAssignmentsForHIT{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAdditionalAssignmentsForHIT{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAdditionalAssignmentsForHIT"); err != nil {
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
	if err = addOpCreateAdditionalAssignmentsForHITValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAdditionalAssignmentsForHIT(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAdditionalAssignmentsForHIT(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAdditionalAssignmentsForHIT",
	}
}
