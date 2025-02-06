// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LoadBalancerSpec defines the desired state of LoadBalancer.
//
// Information about a load balancer.
type LoadBalancerSpec struct {

	// The load balancer attributes.
	Attributes []*LoadBalancerAttribute `json:"attributes,omitempty"`
	// [Application Load Balancers on Outposts] The ID of the customer-owned address
	// pool (CoIP pool).
	CustomerOwnedIPv4Pool *string `json:"customerOwnedIPv4Pool,omitempty"`
	// The IP address type. Internal load balancers must use ipv4.
	//
	// [Application Load Balancers] The possible values are ipv4 (IPv4 addresses),
	// dualstack (IPv4 and IPv6 addresses), and dualstack-without-public-ipv4 (public
	// IPv6 addresses and private IPv4 and IPv6 addresses).
	//
	// [Network Load Balancers and Gateway Load Balancers] The possible values are
	// ipv4 (IPv4 addresses) and dualstack (IPv4 and IPv6 addresses).
	IPAddressType *string `json:"ipAddressType,omitempty"`
	// The name of the load balancer.
	//
	// This name must be unique per region per account, can have a maximum of 32
	// characters, must contain only alphanumeric characters or hyphens, must not
	// begin or end with a hyphen, and must not begin with "internal-".
	Name *string `json:"name,omitempty"`
	// The nodes of an Internet-facing load balancer have public IP addresses. The
	// DNS name of an Internet-facing load balancer is publicly resolvable to the
	// public IP addresses of the nodes. Therefore, Internet-facing load balancers
	// can route requests from clients over the internet.
	//
	// The nodes of an internal load balancer have only private IP addresses. The
	// DNS name of an internal load balancer is publicly resolvable to the private
	// IP addresses of the nodes. Therefore, internal load balancers can route requests
	// only from clients with access to the VPC for the load balancer.
	//
	// The default is an Internet-facing load balancer.
	//
	// You can't specify a scheme for a Gateway Load Balancer.
	Scheme            *string                                    `json:"scheme,omitempty"`
	SecurityGroupRefs []*ackv1alpha1.AWSResourceReferenceWrapper `json:"securityGroupRefs,omitempty"`
	// [Application Load Balancers and Network Load Balancers] The IDs of the security
	// groups for the load balancer.
	SecurityGroups []*string `json:"securityGroups,omitempty"`
	// The IDs of the subnets. You can specify only one subnet per Availability
	// Zone. You must specify either subnets or subnet mappings, but not both.
	//
	// [Application Load Balancers] You must specify subnets from at least two Availability
	// Zones. You can't specify Elastic IP addresses for your subnets.
	//
	// [Application Load Balancers on Outposts] You must specify one Outpost subnet.
	//
	// [Application Load Balancers on Local Zones] You can specify subnets from
	// one or more Local Zones.
	//
	// [Network Load Balancers] You can specify subnets from one or more Availability
	// Zones. You can specify one Elastic IP address per subnet if you need static
	// IP addresses for your internet-facing load balancer. For internal load balancers,
	// you can specify one private IP address per subnet from the IPv4 range of
	// the subnet. For internet-facing load balancer, you can specify one IPv6 address
	// per subnet.
	//
	// [Gateway Load Balancers] You can specify subnets from one or more Availability
	// Zones. You can't specify Elastic IP addresses for your subnets.
	SubnetMappings []*SubnetMapping                           `json:"subnetMappings,omitempty"`
	SubnetRefs     []*ackv1alpha1.AWSResourceReferenceWrapper `json:"subnetRefs,omitempty"`
	// The IDs of the subnets. You can specify only one subnet per Availability
	// Zone. You must specify either subnets or subnet mappings, but not both. To
	// specify an Elastic IP address, specify subnet mappings instead of subnets.
	//
	// [Application Load Balancers] You must specify subnets from at least two Availability
	// Zones.
	//
	// [Application Load Balancers on Outposts] You must specify one Outpost subnet.
	//
	// [Application Load Balancers on Local Zones] You can specify subnets from
	// one or more Local Zones.
	//
	// [Network Load Balancers and Gateway Load Balancers] You can specify subnets
	// from one or more Availability Zones.
	Subnets []*string `json:"subnets,omitempty"`
	// The tags to assign to the load balancer.
	Tags []*Tag `json:"tags,omitempty"`
	// The type of load balancer. The default is application.
	Type *string `json:"type,omitempty"`
}

// LoadBalancerStatus defines the observed state of LoadBalancer
type LoadBalancerStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The subnets for the load balancer.
	// +kubebuilder:validation:Optional
	AvailabilityZones []*AvailabilityZone `json:"availabilityZones,omitempty"`
	// The ID of the Amazon Route 53 hosted zone associated with the load balancer.
	// +kubebuilder:validation:Optional
	CanonicalHostedZoneID *string `json:"canonicalHostedZoneID,omitempty"`
	// The date and time the load balancer was created.
	// +kubebuilder:validation:Optional
	CreatedTime *metav1.Time `json:"createdTime,omitempty"`
	// The public DNS name of the load balancer.
	// +kubebuilder:validation:Optional
	DNSName *string `json:"dnsName,omitempty"`
	// Indicates whether to evaluate inbound security group rules for traffic sent
	// to a Network Load Balancer through Amazon Web Services PrivateLink.
	// +kubebuilder:validation:Optional
	EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic *string `json:"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic,omitempty"`
	// The state of the load balancer.
	// +kubebuilder:validation:Optional
	State *LoadBalancerState `json:"state,omitempty"`
	// The ID of the VPC for the load balancer.
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcID,omitempty"`
}

// LoadBalancer is the Schema for the LoadBalancers API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type LoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerSpec   `json:"spec,omitempty"`
	Status            LoadBalancerStatus `json:"status,omitempty"`
}

// LoadBalancerList contains a list of LoadBalancer
// +kubebuilder:object:root=true
type LoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LoadBalancer{}, &LoadBalancerList{})
}
