package base

import (
	"encoding/binary"
)

const PACKET_HEADER_SIZE = 4 //包头

// packetBuffer 根据指定的命令、缓冲字符串和缓冲区大小创建一个数据包。
func packetBuffer(cmd uint16, bufferStr []byte, littleEndian bool) []byte {
	bufferSize := len(bufferStr)
	lenBytes := make([]byte, 2)
	cmdBytes := make([]byte, 2)

	if littleEndian {
		binary.LittleEndian.PutUint16(lenBytes, uint16(bufferSize+PACKET_HEADER_SIZE))
		binary.LittleEndian.PutUint16(cmdBytes, cmd)
	} else {
		binary.BigEndian.PutUint16(lenBytes, uint16(bufferSize+PACKET_HEADER_SIZE))
		binary.BigEndian.PutUint16(cmdBytes, cmd)
	}

	buffer := append(lenBytes, cmdBytes...)
	buffer = append(buffer, bufferStr...)
	return buffer
}

// 定义解析数据包的函数
func parsePacket(buffer []byte, littleEndian bool) map[string]interface{} {
	header := buffer[:PACKET_HEADER_SIZE]
	var lenBytes uint16
	var cmdBytes uint16
	if littleEndian {
		lenBytes = binary.LittleEndian.Uint16(header[:2])
		cmdBytes = binary.LittleEndian.Uint16(header[2:])
	} else {
		lenBytes = binary.BigEndian.Uint16(header[:2])
		cmdBytes = binary.BigEndian.Uint16(header[2:])
	}

	return map[string]interface{}{
		"cmd":     cmdBytes,
		"msgBody": buffer[PACKET_HEADER_SIZE:lenBytes],
	}
}
