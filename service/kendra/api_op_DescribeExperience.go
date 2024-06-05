// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about your Amazon Kendra experience such as a search
// application. For more information on creating a search application experience,
// see [Building a search experience with no code].
//
// [Building a search experience with no code]: https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html
func (c *Client) DescribeExperience(ctx context.Context, params *DescribeExperienceInput, optFns ...func(*Options)) (*DescribeExperienceOutput, error) {
	if params == nil {
		params = &DescribeExperienceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeExperience", params, optFns, c.addOperationDescribeExperienceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeExperienceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeExperienceInput struct {

	// The identifier of your Amazon Kendra experience you want to get information on.
	//
	// This member is required.
	Id *string

	// The identifier of the index for your Amazon Kendra experience.
	//
	// This member is required.
	IndexId *string

	noSmithyDocumentSerde
}

type DescribeExperienceOutput struct {

	// Shows the configuration information for your Amazon Kendra experience. This
	// includes ContentSourceConfiguration , which specifies the data source IDs and/or
	// FAQ IDs, and UserIdentityConfiguration , which specifies the user or group
	// information to grant access to your Amazon Kendra experience.
	Configuration *types.ExperienceConfiguration

	// The Unix timestamp when your Amazon Kendra experience was created.
	CreatedAt *time.Time

	// Shows the description for your Amazon Kendra experience.
	Description *string

	// Shows the endpoint URLs for your Amazon Kendra experiences. The URLs are unique
	// and fully hosted by Amazon Web Services.
	Endpoints []types.ExperienceEndpoint

	// The reason your Amazon Kendra experience could not properly process.
	ErrorMessage *string

	// Shows the identifier of your Amazon Kendra experience.
	Id *string

	// Shows the identifier of the index for your Amazon Kendra experience.
	IndexId *string

	// Shows the name of your Amazon Kendra experience.
	Name *string

	// Shows the Amazon Resource Name (ARN) of a role with permission to access Query
	// API, QuerySuggestions API, SubmitFeedback API, and IAM Identity Center that
	// stores your user and group information.
	RoleArn *string

	// The current processing status of your Amazon Kendra experience. When the status
	// is ACTIVE , your Amazon Kendra experience is ready to use. When the status is
	// FAILED , the ErrorMessage field contains the reason that this failed.
	Status types.ExperienceStatus

	// The Unix timestamp when your Amazon Kendra experience was last updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeExperienceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeExperience{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeExperience{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeExperience"); err != nil {
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
	if err = addOpDescribeExperienceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeExperience(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeExperience(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeExperience",
	}
}
