// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cluster

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type HostInfo to the protobuf v3 wire format
func (val *HostInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type HostInfo from the protobuf v3 wire format
func (val *HostInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *HostInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two HostInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *HostInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *HostInfo
	switch t := that.(type) {
	case *HostInfo:
		that1 = t
	case HostInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RingInfo to the protobuf v3 wire format
func (val *RingInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RingInfo from the protobuf v3 wire format
func (val *RingInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RingInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RingInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RingInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RingInfo
	switch t := that.(type) {
	case *RingInfo:
		that1 = t
	case RingInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type MembershipInfo to the protobuf v3 wire format
func (val *MembershipInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type MembershipInfo from the protobuf v3 wire format
func (val *MembershipInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *MembershipInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two MembershipInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *MembershipInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *MembershipInfo
	switch t := that.(type) {
	case *MembershipInfo:
		that1 = t
	case MembershipInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ClusterMember to the protobuf v3 wire format
func (val *ClusterMember) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ClusterMember from the protobuf v3 wire format
func (val *ClusterMember) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ClusterMember) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ClusterMember values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ClusterMember) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ClusterMember
	switch t := that.(type) {
	case *ClusterMember:
		that1 = t
	case ClusterMember:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}