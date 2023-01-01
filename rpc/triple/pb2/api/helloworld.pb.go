/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: rpc/triple/pb2/api/helloworld.proto

package api

import (
	models "github.com/apache/dubbo-go-samples/rpc/triple/pb2/models"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_rpc_triple_pb2_api_helloworld_proto protoreflect.FileDescriptor

var file_rpc_triple_pb2_api_helloworld_proto_rawDesc = []byte{
	0x0a, 0x23, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x32,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x67, 0x6f, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x62, 0x32,
	0x2e, 0x61, 0x70, 0x69, 0x1a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2d, 0x67, 0x6f,
	0x2d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x72, 0x69,
	0x70, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x32, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc1,
	0x02, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x94, 0x01, 0x0a, 0x08, 0x53,
	0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x46, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x64, 0x75, 0x62, 0x62,
	0x6f, 0x5f, 0x67, 0x6f, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x74, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x62, 0x32, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3e, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x61,
	0x63, 0x68, 0x65, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x5f, 0x67, 0x6f, 0x5f, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x2e,
	0x70, 0x62, 0x32, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22,
	0x00, 0x12, 0x9e, 0x01, 0x0a, 0x0e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x46, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x5f, 0x67,
	0x6f, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x72,
	0x69, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x62, 0x32, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x5f, 0x67, 0x6f, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x62, 0x32,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x61, 0x70, 0x69,
}

var file_rpc_triple_pb2_api_helloworld_proto_goTypes = []interface{}{
	(*models.HelloRequest)(nil), // 0: github.com.apache.dubbo_go_samples.rpc.triple.pb2.models.HelloRequest
	(*models.User)(nil),         // 1: github.com.apache.dubbo_go_samples.rpc.triple.pb2.models.User
}
var file_rpc_triple_pb2_api_helloworld_proto_depIdxs = []int32{
	0, // 0: org.apache.dubbogo.samples.rpc.triple.pb2.api.Greeter.SayHello:input_type -> github.com.apache.dubbo_go_samples.rpc.triple.pb2.models.HelloRequest
	0, // 1: org.apache.dubbogo.samples.rpc.triple.pb2.api.Greeter.SayHelloStream:input_type -> github.com.apache.dubbo_go_samples.rpc.triple.pb2.models.HelloRequest
	1, // 2: org.apache.dubbogo.samples.rpc.triple.pb2.api.Greeter.SayHello:output_type -> github.com.apache.dubbo_go_samples.rpc.triple.pb2.models.User
	1, // 3: org.apache.dubbogo.samples.rpc.triple.pb2.api.Greeter.SayHelloStream:output_type -> github.com.apache.dubbo_go_samples.rpc.triple.pb2.models.User
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_triple_pb2_api_helloworld_proto_init() }
func file_rpc_triple_pb2_api_helloworld_proto_init() {
	if File_rpc_triple_pb2_api_helloworld_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_triple_pb2_api_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_triple_pb2_api_helloworld_proto_goTypes,
		DependencyIndexes: file_rpc_triple_pb2_api_helloworld_proto_depIdxs,
	}.Build()
	File_rpc_triple_pb2_api_helloworld_proto = out.File
	file_rpc_triple_pb2_api_helloworld_proto_rawDesc = nil
	file_rpc_triple_pb2_api_helloworld_proto_goTypes = nil
	file_rpc_triple_pb2_api_helloworld_proto_depIdxs = nil
}
