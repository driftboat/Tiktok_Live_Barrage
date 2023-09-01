# tkbarrage
Tiktok弹幕抓取 发送websocket服务器   
Receive live stream events such as comments and gifts and send to your websocket server tool in Go language.  free for one month 
免费试用一个月 下载客户端  https://github.com/driftboat/TiktokBarrage/releases/download/1.0.0/Barrage.zip  
续期,请联系 qq: 174389107  telegram：db_444444  
Go语言实现，js注入  
弹幕游戏服务器开发，房间消息转发,积分榜7天排名，单台服务器24小时压测支持1000主播同时播，1000 qps  
![avatar](images/tk1.png)

# How to run
- Download https://github.com/driftboat/TiktokBarrage/releases/download/1.0.0/Barrage.zip  
- Unzip and run Barrage.exe  
- Clik on the start button
- Waiting for Chrome to open

# Create your websocket server
Write a go server to receive and parse data like this
```go

func (c *Client) Read() {
	go func() {
		for {
			messageType, msg, err := c.Socket.ReadMessage()
			if err != nil {
				if messageType == -1 && websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived) {
					Manager.DisConnect <- c
					return
				} else if messageType != websocket.PingMessage {
					return
				}
			} else {
				msgString := string(msg)
				var jsonData map[string]interface{}
				err := json.Unmarshal([]byte(msgString), &jsonData)
				if err != nil {
					logrus.Error("Failed to unmarshal JSON:", err)

				}

				if common, ok := jsonData["common"].(map[string]interface{}); ok {
					if method, ok := common["method"].(string); ok {
						logrus.Info("Method:", method)
 

						if method == "WebcastGiftMessage" {
							var giftMessage msgs.WebcastGiftMessage
							err := json.Unmarshal([]byte(msgString), &giftMessage)
							if err != nil {
								logrus.Error("Failed to unmarshal JSON into WebcastGiftMessage:", err)
							} else {
								c.LogChan <- giftMessage.User.Nickname + "使用了" + giftMessage.Gift.Describe + "数量" + giftMessage.GroupCount
							}
							// Process the giftMessage struct as needed

						}

						if method == "WebcastLikeMessage" {
							var likeMessage msgs.WebcastLikeMessage
							err := json.Unmarshal([]byte(msgString), &likeMessage)
							if err != nil {
								logrus.Error("Failed to unmarshal JSON into WebcastGiftMessage:", err)
							} else {
								c.LogChan <- likeMessage.User.Nickname + "点赞" + likeMessage.Count + "次，总数" + likeMessage.Total
							}
							// Process the giftMessage struct as needed

						}

						if method == "WebcastChatMessage" {
							var chatMessage msgs.WebcastChatMessage
							err := json.Unmarshal([]byte(msgString), &chatMessage)
							if err != nil {
								logrus.Error("Failed to unmarshal JSON into WebcastGiftMessage:", err)
							} else {
								c.LogChan <- chatMessage.User.Nickname + ":" + chatMessage.Content
							}
							// Process the giftMessage struct as needed

						}

					}
				}

			}
		}
	}()
}

```


# 数据结构示例
All data structures are in the msgs directory
![avatar](images/tk2.png)

