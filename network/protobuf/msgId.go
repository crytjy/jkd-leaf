package base

import (
	"google.golang.org/protobuf/proto"
	"reflect"
)

// ReqMsgId 请求PB协议ID
var ReqMsgId = map[uint16]proto.Message{}

// RepMsgId 响应PB协议ID
var RepMsgId = map[reflect.Type]uint16{}

func Handle(reqId, repId uint16, reqProto, repProto proto.Message) {
	ReqMsgId[reqId] = reqProto
	RepMsgId[reflect.TypeOf(repProto)] = repId
}
