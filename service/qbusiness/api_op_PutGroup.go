// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create, or updates, a mapping of users—who have access to a document—to groups.
//
// You can also map sub groups to groups. For example, the group "Company
// Intellectual Property Teams" includes sub groups "Research" and "Engineering".
// These sub groups include their own list of users or people who work in these
// teams. Only users who work in research and engineering, and therefore belong in
// the intellectual property group, can see top-secret company documents in their
// Amazon Q Business chat results.
func (c *Client) PutGroup(ctx context.Context, params *PutGroupInput, optFns ...func(*Options)) (*PutGroupOutput, error) {
	if params == nil {
		params = &PutGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutGroup", params, optFns, c.addOperationPutGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutGroupInput struct {

	// The identifier of the application in which the user and group mapping belongs.
	//
	// This member is required.
	ApplicationId *string

	// A list of users or sub groups that belong to a group. This is for generating
	// Amazon Q Business chat results only from document a user has access to.
	//
	// This member is required.
	GroupMembers *types.GroupMembers

	// The list that contains your users or sub groups that belong the same group. For
	// example, the group "Company" includes the user "CEO" and the sub groups
	// "Research", "Engineering", and "Sales and Marketing".
	//
	// If you have more than 1000 users and/or sub groups for a single group, you need
	// to provide the path to the S3 file that lists your users and sub groups for a
	// group. Your sub groups can contain more than 1000 users, but the list of sub
	// groups that belong to a group (and/or users) must be no more than 1000.
	//
	// This member is required.
	GroupName *string

	// The identifier of the index in which you want to map users to their groups.
	//
	// This member is required.
	IndexId *string

	// The type of the group.
	//
	// This member is required.
	Type types.MembershipType

	// The identifier of the data source for which you want to map users to their
	// groups. This is useful if a group is tied to multiple data sources, but you only
	// want the group to access documents of a certain data source. For example, the
	// groups "Research", "Engineering", and "Sales and Marketing" are all tied to the
	// company's documents stored in the data sources Confluence and Salesforce.
	// However, "Sales and Marketing" team only needs access to customer-related
	// documents stored in Salesforce.
	DataSourceId *string

	noSmithyDocumentSerde
}

type PutGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutGroup"); err != nil {
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
	if err = addOpPutGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutGroup",
	}
}
