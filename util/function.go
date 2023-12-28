package util

import (
	"bytes"
	"fmt"
	"github.com/sony/sonyflake"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

// CheckErr 错误处理 | panic
func CheckErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

// FmtErr 错误处理 | 打印
func FmtErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

// PrintData 打印数据
func PrintData(data []map[string]interface{}) {
	for _, row := range data {
		for column, value := range row {
			fmt.Printf("%s: %v\t", column, value)
		}
		fmt.Println()
	}
}

// RandId 雪花ID
func RandId() uint64 {
	// 创建一个新的雪花ID生成器
	sf := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Now(),
	})

	// 生成雪花随机数
	id, err := sf.NextID()
	if err != nil {
		stamp := GetMilliTimeStamp()
		return uint64(stamp)
	}

	return id
}

// RandIdByPre 雪花ID
func RandIdByPre(prefix uint16) uint64 {
	stamp := GetMilliTimeStamp()
	// 创建一个新的雪花ID生成器
	sf := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Now(),
		MachineID: func() (uint16, error) {
			prefixId := prefix + uint16(stamp)
			return prefixId, nil
		},
	})

	// 生成雪花随机数
	id, err := sf.NextID()
	if err != nil {
		return uint64(stamp)
	}

	return id
}

// GetGoroutineId 获取协程ID
func GetGoroutineId() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)

	return n
}

// SwapKeysAndValues 互换Key Value
func SwapKeysAndValues(originalMap map[string]int) map[int]string {
	swappedMap := make(map[int]string)

	for key, value := range originalMap {
		swappedMap[value] = key
	}

	return swappedMap
}

// FormatGoFile 格式化 Go 文件
func FormatGoFile(filename string) error {
	// 使用 exec 包运行 gofmt 命令
	cmd := exec.Command("gofmt", "-w", filename)

	// 设置输出和错误输出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 运行 gofmt 命令
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
