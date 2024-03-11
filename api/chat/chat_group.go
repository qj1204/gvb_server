package chat

import (
	"encoding/json"
	"fmt"
	"github.com/DanPlayer/randomname"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/ctype"
	"gvb_server/models/common/response"
	"net/http"
	"strings"
	"time"
)

type ChatUser struct {
	Conn     *websocket.Conn
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}

var ConnGroup = make(map[string]ChatUser)

type GroupRequest struct {
	Content string        `json:"content"`  // 内容
	MsgType ctype.MsgType `json:"msg_type"` // 消息类型
}

type GroupResponse struct {
	NickName    string        `json:"nick_name"`    // 昵称
	Avatar      string        `json:"avatar"`       // 头像
	Content     string        `json:"content"`      // 内容
	MsgType     ctype.MsgType `json:"msg_type"`     // 消息类型
	OnlineCount int           `json:"online_count"` // 在线人数
	Date        time.Time     `json:"date"`         // 发送时间
}

func (this *ChatApi) ChatGroupView(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// 鉴权，true表示放行，false表示拦截
			return true
		},
	}
	// 将http请求升级为websocket请求
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		response.FailWithMessage("升级为websocket失败", c)
		return
	}

	curAddr := conn.RemoteAddr().String()
	// 需要生成昵称，根据昵称首字关联头像地址
	nickName := randomname.GenerateName()
	nickNameFirst := string([]rune(nickName)[0])
	chatUser := ChatUser{
		Conn:     conn,
		NickName: nickName,
		Avatar:   fmt.Sprintf("static/chat_avatar/%s.png", nickNameFirst),
	}
	ConnGroup[curAddr] = chatUser

	for {
		// 消息类型，消息，错误
		_, p, err := conn.ReadMessage() // p就是客户端发给服务端的消息
		if err != nil {
			// 用户断开连接
			SendGroupMsg(curAddr, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				Content:     fmt.Sprintf("%s 离开聊天室", chatUser.NickName),
				MsgType:     ctype.OutRoomMsg,
				OnlineCount: len(ConnGroup) - 1,
				Date:        time.Now(),
			})
			break
		}
		// 参数绑定
		var gRequest GroupRequest
		err = json.Unmarshal(p, &gRequest)
		if err != nil {
			SendMsg(curAddr, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				Content:     "参数绑定失败",
				MsgType:     ctype.SystemMsg,
				OnlineCount: len(ConnGroup),
				Date:        time.Now(),
			})
			continue
		}
		// 判断消息类型
		switch gRequest.MsgType {
		case ctype.TextMsg:
			if strings.TrimSpace(gRequest.Content) == "" {
				SendMsg(curAddr, GroupResponse{
					NickName:    chatUser.NickName,
					Avatar:      chatUser.Avatar,
					Content:     "消息不能为空",
					MsgType:     ctype.SystemMsg,
					OnlineCount: len(ConnGroup),
					Date:        time.Now(),
				})
				continue
			}
			SendGroupMsg(curAddr, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				Content:     gRequest.Content,
				MsgType:     ctype.TextMsg,
				OnlineCount: len(ConnGroup),
				Date:        time.Now(),
			})
		case ctype.InRoomMsg:
			SendGroupMsg(curAddr, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				Content:     fmt.Sprintf("%s 加入聊天室", chatUser.NickName),
				MsgType:     ctype.InRoomMsg,
				OnlineCount: len(ConnGroup),
				Date:        time.Now(),
			})
		default:
			SendMsg(curAddr, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				Content:     "消息类型错误",
				MsgType:     ctype.SystemMsg,
				OnlineCount: len(ConnGroup),
				Date:        time.Now(),
			})
		}

	}
	defer conn.Close()
	delete(ConnGroup, curAddr)
}

// SendGroupMsg 发送消息给所有客户端
func SendGroupMsg(curAddr string, gResponse GroupResponse) {
	byteData, _ := json.Marshal(gResponse)
	ip, addr := GetIPAndAddr(curAddr)
	global.DB.Create(&models.ChatModel{
		NickName: gResponse.NickName,
		Avatar:   gResponse.Avatar,
		Content:  gResponse.Content,
		MsgType:  gResponse.MsgType,
		IP:       ip,
		Addr:     addr,
		IsGroup:  true,
	})
	for _, chatUser := range ConnGroup {
		// 服务端发送消息给客户端
		chatUser.Conn.WriteMessage(websocket.TextMessage, byteData)
	}
}

// SendMsg 发送消息给当前客户端
func SendMsg(curAddr string, gResponse GroupResponse) {
	byteData, _ := json.Marshal(gResponse)
	chatUser := ConnGroup[curAddr]
	ip, addr := GetIPAndAddr(curAddr)
	global.DB.Create(&models.ChatModel{
		NickName: gResponse.NickName,
		Avatar:   gResponse.Avatar,
		Content:  gResponse.Content,
		MsgType:  gResponse.MsgType,
		IP:       ip,
		Addr:     addr,
		IsGroup:  false,
	})
	chatUser.Conn.WriteMessage(websocket.TextMessage, byteData)
}

// GetIPAndAddr 获取用户的ip和地址
func GetIPAndAddr(curAddr string) (ip, addr string) {
	userAddr := "内网"
	return strings.Split(curAddr, ":")[0], userAddr
}
