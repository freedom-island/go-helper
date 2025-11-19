package helper

import (
	"github.com/gorilla/websocket"
	"log"
	"strings"
)

// CheckWebSocketError 检查 WebSocket 错误
func CheckWebSocketError(err error) bool {
	if strings.Contains(err.Error(), "forcibly closed by the remote host") || strings.Contains(err.Error(), "established connection was aborted by the software in your host machine") {
		log.Println("连接被远程主机关闭，尝试重连")
		return true
	} else if strings.Contains(err.Error(), "A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond") {
		log.Println("连接因网络故障关闭，尝试重连(远端已关闭)")
		return true
	} else if strings.Contains(err.Error(), " i/o timeout") {
		log.Println("连接因网络故障关闭，尝试重连(超时)")
		return true
	} else if strings.Contains(err.Error(), "use of closed network connection") {
		log.Println("未知原因关闭，尝试重连")
		return true
	} else if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
		log.Println("读取数据异常中断，尝试重连")
		return true
	}

	return false
}
