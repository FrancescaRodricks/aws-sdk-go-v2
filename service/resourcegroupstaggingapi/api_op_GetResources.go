// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroupstaggingapi

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns all the tagged or previously tagged resources that are located in the
// specified Amazon Web Services Region for the account.
//
// Depending on what information you want returned, you can also specify the
// following:
//
//   - Filters that specify what tags and resource types you want returned. The
//     response includes all tags that are associated with the requested resources.
//
//   - Information about compliance with the account's effective tag policy. For
//     more information on tag policies, see [Tag Policies]in the Organizations User Guide.
//
// This operation supports pagination, where the response can be sent in multiple
// pages. You should check the PaginationToken response parameter to determine if
// there are additional results available to return. Repeat the query, passing the
// PaginationToken response parameter value as an input to the next request until
// you recieve a null value. A null value for PaginationToken indicates that there
// are no more results waiting to be returned.
//
// [Tag Policies]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html
func (c *Client) GetResources(ctx context.Context, params *GetResourcesInput, optFns ...func(*Options)) (*GetResourcesOutput, error) {
	if params == nil {
		params = &GetResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResources", params, optFns, c.addOperationGetResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourcesInput struct {

	// Specifies whether to exclude resources that are compliant with the tag policy.
	// Set this to true if you are interested in retrieving information on
	// noncompliant resources only.
	//
	// You can use this parameter only if the IncludeComplianceDetails parameter is
	// also set to true .
	ExcludeCompliantResources *bool

	// Specifies whether to include details regarding the compliance with the
	// effective tag policy. Set this to true to determine whether resources are
	// compliant with the tag policy and to get details.
	IncludeComplianceDetails *bool

	// Specifies a PaginationToken response value from a previous request to indicate
	// that you want the next page of results. Leave this parameter empty in your
	// initial request.
	PaginationToken *string

	// Specifies a list of ARNs of resources for which you want to retrieve tag data.
	// You can't specify both this parameter and any of the pagination parameters (
	// ResourcesPerPage , TagsPerPage , PaginationToken ) in the same request. If you
	// specify both, you get an Invalid Parameter exception.
	//
	// If a resource specified by this parameter doesn't exist, it doesn't generate an
	// error; it simply isn't included in the response.
	//
	// An ARN (Amazon Resource Name) uniquely identifies a resource. For more
	// information, see [Amazon Resource Names (ARNs) and Amazon Web Services Service Namespaces]in the Amazon Web Services General Reference.
	//
	// [Amazon Resource Names (ARNs) and Amazon Web Services Service Namespaces]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	ResourceARNList []string

	// Specifies the resource types that you want included in the response. The format
	// of each resource type is service[:resourceType] . For example, specifying a
	// resource type of ec2 returns all Amazon EC2 resources (which includes EC2
	// instances). Specifying a resource type of ec2:instance returns only EC2
	// instances.
	//
	// The string for each service name and resource type is the same as that embedded
	// in a resource's Amazon Resource Name (ARN). For the list of services whose
	// resources you can use in this parameter, see [Services that support the Resource Groups Tagging API].
	//
	// You can specify multiple resource types by using an array. The array can
	// include up to 100 items. Note that the length constraint requirement applies to
	// each resource type filter. For example, the following string would limit the
	// response to only Amazon EC2 instances, Amazon S3 buckets, or any Audit Manager
	// resource:
	//
	//     ec2:instance,s3:bucket,auditmanager
	//
	// [Services that support the Resource Groups Tagging API]: https://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/supported-services.html
	ResourceTypeFilters []string

	// Specifies the maximum number of results to be returned in each page. A query
	// can return fewer than this maximum, even if there are more results still to
	// return. You should always check the PaginationToken response value to see if
	// there are more results. You can specify a minimum of 1 and a maximum value of
	// 100.
	ResourcesPerPage *int32

	// Specifies a list of TagFilters (keys and values) to restrict the output to only
	// those resources that have tags with the specified keys and, if included, the
	// specified values. Each TagFilter must contain a key with values optional. A
	// request can include up to 50 keys, and each key can include up to 20 values.
	//
	// Note the following when deciding how to use TagFilters:
	//
	//   - If you don't specify a TagFilter , the response includes all resources that
	//   are currently tagged or ever had a tag. Resources that currently don't have tags
	//   are shown with an empty tag set, like this: "Tags": [] .
	//
	//   - If you specify more than one filter in a single request, the response
	//   returns only those resources that satisfy all filters.
	//
	//   - If you specify a filter that contains more than one value for a key, the
	//   response returns resources that match any of the specified values for that key.
	//
	//   - If you don't specify a value for a key, the response returns all resources
	//   that are tagged with that key, with any or no value.
	//
	// For example, for the following filters: filter1= {keyA,{value1}} ,
	//   filter2={keyB,{value2,value3,value4}} , filter3= {keyC} :
	//
	//   - GetResources({filter1}) returns resources tagged with key1=value1
	//
	//   - GetResources({filter2}) returns resources tagged with key2=value2 or
	//   key2=value3 or key2=value4
	//
	//   - GetResources({filter3}) returns resources tagged with any tag with the key
	//   key3 , and with any or no value
	//
	//   - GetResources({filter1,filter2,filter3}) returns resources tagged with
	//   (key1=value1) and (key2=value2 or key2=value3 or key2=value4) and (key3, any or
	//   no value)
	TagFilters []types.TagFilter

	// Amazon Web Services recommends using ResourcesPerPage instead of this parameter.
	//
	// A limit that restricts the number of tags (key and value pairs) returned by
	// GetResources in paginated output. A resource with no tags is counted as having
	// one tag (one key and value pair).
	//
	// GetResources does not split a resource and its associated tags across pages. If
	// the specified TagsPerPage would cause such a break, a PaginationToken is
	// returned in place of the affected resource and its tags. Use that token in
	// another request to get the remaining data. For example, if you specify a
	// TagsPerPage of 100 and the account has 22 resources with 10 tags each (meaning
	// that each resource has 10 key and value pairs), the output will consist of three
	// pages. The first page displays the first 10 resources, each with its 10 tags.
	// The second page displays the next 10 resources, each with its 10 tags. The third
	// page displays the remaining 2 resources, each with its 10 tags.
	//
	// You can set TagsPerPage to a minimum of 100 items up to a maximum of 500 items.
	TagsPerPage *int32

	noSmithyDocumentSerde
}

type GetResourcesOutput struct {

	// A string that indicates that there is more data available than this response
	// contains. To receive the next part of the response, specify this response value
	// as the PaginationToken value in the request for the next page.
	PaginationToken *string

	// A list of resource ARNs and the tags (keys and values) associated with each.
	ResourceTagMappingList []types.ResourceTagMapping

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetResources"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResources(options.Region), middleware.Before); err != nil {
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

// GetResourcesAPIClient is a client that implements the GetResources operation.
type GetResourcesAPIClient interface {
	GetResources(context.Context, *GetResourcesInput, ...func(*Options)) (*GetResourcesOutput, error)
}

var _ GetResourcesAPIClient = (*Client)(nil)

// GetResourcesPaginatorOptions is the paginator options for GetResources
type GetResourcesPaginatorOptions struct {
	// Specifies the maximum number of results to be returned in each page. A query
	// can return fewer than this maximum, even if there are more results still to
	// return. You should always check the PaginationToken response value to see if
	// there are more results. You can specify a minimum of 1 and a maximum value of
	// 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetResourcesPaginator is a paginator for GetResources
type GetResourcesPaginator struct {
	options   GetResourcesPaginatorOptions
	client    GetResourcesAPIClient
	params    *GetResourcesInput
	nextToken *string
	firstPage bool
}

// NewGetResourcesPaginator returns a new GetResourcesPaginator
func NewGetResourcesPaginator(client GetResourcesAPIClient, params *GetResourcesInput, optFns ...func(*GetResourcesPaginatorOptions)) *GetResourcesPaginator {
	if params == nil {
		params = &GetResourcesInput{}
	}

	options := GetResourcesPaginatorOptions{}
	if params.ResourcesPerPage != nil {
		options.Limit = *params.ResourcesPerPage
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.PaginationToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetResources page.
func (p *GetResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetResourcesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PaginationToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.ResourcesPerPage = limit

	result, err := p.client.GetResources(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.PaginationToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetResources",
	}
}
