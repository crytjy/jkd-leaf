package protobuf

import (
	"errors"
	"fmt"
	"github.com/crytjy/jkd-leaf/chanrpc"
	"github.com/crytjy/jkd-leaf/log"
	"google.golang.org/protobuf/proto"
	"math"
	"reflect"
)

// -------------------------
// | msgId | protobuf message |
// -------------------------
type Processor struct {
	littleEndian bool
	msgInfo      map[uint16]*MsgInfo
}

type MsgInfo struct {
	msgType       reflect.Type
	msgRouter     *chanrpc.Server
	msgHandler    MsgHandler
	msgRawHandler MsgHandler
}

type MsgHandler func([]interface{})

type MsgRaw struct {
	msgID      uint16
	msgRawData []byte
}

type UserData struct {
	UserId  int64
	GuildId int32
	Req     interface{}
}

func NewProcessor() *Processor {
	p := new(Processor)
	p.littleEndian = false

	return p
}

// SetByteOrder It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) SetByteOrder(littleEndian bool) {
	p.littleEndian = littleEndian
}

// Register It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) Register(msgId uint16) uint16 {
	// 从 map 中获取值，并进行类型断言
	msg, ok := ReqMsgId[msgId]
	if !ok {
		fmt.Errorf("MSG:msgInstance Errorf", msg)
	}

	msgType := reflect.TypeOf(msg)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		log.Fatal("protobuf message pointer required")
	}

	if _, ok := p.msgInfo[msgId]; ok {
		log.Fatal("message %s is already registered", msgType)
	}

	if len(p.msgInfo) >= math.MaxUint16 {
		log.Fatal("too many protobuf messages (max = %v)", math.MaxUint16)
	}
	if len(p.msgInfo) == 0 {
		p.msgInfo = make(map[uint16]*MsgInfo)
	}

	i := new(MsgInfo)
	i.msgType = msgType

	p.msgInfo[msgId] = i
	return msgId
}

// SetRouter It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) SetRouter(msgId uint16, msgRouter *chanrpc.Server) {
	if _, ok := ReqMsgId[msgId]; !ok {
		log.Fatal("message %s is not registered", msgId)
	}

	p.msgInfo[msgId].msgRouter = msgRouter
}

// SetHandler It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) SetHandler(msgId interface{}, msgHandler MsgHandler) {
	umsgId := msgId.(uint16)
	_, ok := p.msgInfo[umsgId]
	if !ok {
		log.Fatal("message id %s not registered", umsgId)
	}

	p.msgInfo[umsgId].msgHandler = msgHandler
}

// goroutine safe
func (p *Processor) Route(msgId uint16, msg interface{}, userData interface{}) error {
	fmt.Println("Route p.msgInfo", p.msgInfo)
	// raw
	if msgRaw, ok := msg.(MsgRaw); ok {
		i := p.msgInfo[msgRaw.msgID]
		if i == nil {
			return fmt.Errorf("message id %v not registered", msgRaw.msgID)
		}
		if i.msgRawHandler != nil {
			i.msgRawHandler([]interface{}{msgRaw.msgID, msgRaw.msgRawData, userData})
		}
		return nil
	}

	fmt.Println("Route msgId", msgId)

	// protobuf
	i := p.msgInfo[msgId]
	if i == nil {
		return fmt.Errorf("message %s not registered", msgId)
	}
	if i.msgHandler != nil {
		i.msgHandler([]interface{}{msg, userData})
	}
	if i.msgRouter != nil {
		i.msgRouter.Go(msgId, msg, userData)
	}
	return nil
}

// Unmarshal goroutine safe
func (p *Processor) Unmarshal(data []byte) (interface{}, error, uint16) {
	if len(data) < 2 {
		return nil, errors.New("protobuf data too short"), 0
	}

	var m map[string]interface{}
	m = parsePacket(data, p.littleEndian)

	// id
	id, ok := m["cmd"].(uint16)
	if !ok {
		return nil, fmt.Errorf("Failed to assert cmd as uint16, err:%v", ok), 0
	}

	msgBody, msb := m["msgBody"].([]byte)
	if !msb {
		return nil, fmt.Errorf("Failed to assert msgBody as string, err:%v", ok), 0
	}

	// msg
	i := p.msgInfo[id]
	if i == nil {
		return nil, fmt.Errorf("message id %v not registered", id), 0
	}

	if i.msgRawHandler != nil {
		return MsgRaw{id, msgBody}, nil, 0
	} else {
		msg := reflect.New(i.msgType.Elem()).Interface()
		return msg, proto.Unmarshal(msgBody, msg.(proto.Message)), id
	}
}

// Marshal goroutine safe
func (p *Processor) Marshal(msg interface{}) ([][]byte, error) {
	pbMsg := msg.(proto.Message)
	msgType := reflect.TypeOf(msg)

	msgId, ok := RepMsgId[msgType]
	if !ok {
		err := fmt.Errorf("message msg %s not registered", pbMsg)
		return nil, err
	}

	// data
	data, err := proto.Marshal(pbMsg)

	var m []byte
	m = packetBuffer(msgId, data, p.littleEndian)

	return [][]byte{m}, err
}

// goroutine safe
func (p *Processor) Range(f func(id uint16, t reflect.Type)) {
	for id, i := range p.msgInfo {
		f(uint16(id), i.msgType)
	}
}
