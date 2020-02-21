// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateTransitGatewayMulticastDomainInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The IDs of the subnets to associate with the transit gateway multicast domain.
	SubnetIds []string `locationNameList:"item" type:"list"`

	// The ID of the transit gateway attachment to associate with the transit gateway
	// multicast domain.
	TransitGatewayAttachmentId *string `type:"string"`

	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId *string `type:"string"`
}

// String returns the string representation
func (s AssociateTransitGatewayMulticastDomainInput) String() string {
	return awsutil.Prettify(s)
}

type AssociateTransitGatewayMulticastDomainOutput struct {
	_ struct{} `type:"structure"`

	// Information about the transit gateway multicast domain associations.
	Associations *TransitGatewayMulticastDomainAssociations `locationName:"associations" type:"structure"`
}

// String returns the string representation
func (s AssociateTransitGatewayMulticastDomainOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateTransitGatewayMulticastDomain = "AssociateTransitGatewayMulticastDomain"

// AssociateTransitGatewayMulticastDomainRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Associates the specified subnets and transit gateway attachments with the
// specified transit gateway multicast domain.
//
// The transit gateway attachment must be in the available state before you
// can add a resource. Use DescribeTransitGatewayAttachments (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayAttachments.html)
// to see the state of the attachment.
//
//    // Example sending a request using AssociateTransitGatewayMulticastDomainRequest.
//    req := client.AssociateTransitGatewayMulticastDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateTransitGatewayMulticastDomain
func (c *Client) AssociateTransitGatewayMulticastDomainRequest(input *AssociateTransitGatewayMulticastDomainInput) AssociateTransitGatewayMulticastDomainRequest {
	op := &aws.Operation{
		Name:       opAssociateTransitGatewayMulticastDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateTransitGatewayMulticastDomainInput{}
	}

	req := c.newRequest(op, input, &AssociateTransitGatewayMulticastDomainOutput{})
	return AssociateTransitGatewayMulticastDomainRequest{Request: req, Input: input, Copy: c.AssociateTransitGatewayMulticastDomainRequest}
}

// AssociateTransitGatewayMulticastDomainRequest is the request type for the
// AssociateTransitGatewayMulticastDomain API operation.
type AssociateTransitGatewayMulticastDomainRequest struct {
	*aws.Request
	Input *AssociateTransitGatewayMulticastDomainInput
	Copy  func(*AssociateTransitGatewayMulticastDomainInput) AssociateTransitGatewayMulticastDomainRequest
}

// Send marshals and sends the AssociateTransitGatewayMulticastDomain API request.
func (r AssociateTransitGatewayMulticastDomainRequest) Send(ctx context.Context) (*AssociateTransitGatewayMulticastDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateTransitGatewayMulticastDomainResponse{
		AssociateTransitGatewayMulticastDomainOutput: r.Request.Data.(*AssociateTransitGatewayMulticastDomainOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateTransitGatewayMulticastDomainResponse is the response type for the
// AssociateTransitGatewayMulticastDomain API operation.
type AssociateTransitGatewayMulticastDomainResponse struct {
	*AssociateTransitGatewayMulticastDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateTransitGatewayMulticastDomain request.
func (r *AssociateTransitGatewayMulticastDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}