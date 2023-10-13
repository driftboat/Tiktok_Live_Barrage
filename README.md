# tkbarrage
Tiktok弹幕抓取 发送websocket服务器  
Receive live stream events such as comments and gifts and send to your websocket server tool in Go language.   
请联系 qq: 174389107  telegram：db_444444  
Go语言实现，js注入  
  
![avatar](images/tk1.png)


# 抓取流程说明
## 抓取服务启动
- [下载安装包](https://github.com/driftboat/TiktokBarrage/releases/latest)
- 解压运行 Barrage.exe
- 关闭所有Chrome浏览器(可以不关闭，但必须在窗口中设置一个单独的Chrome用户目录，专门给Barrage使用）  
- 点Start运行
## 客户端接收弹幕消息模拟
- 开启模拟ws客户端 用这个在线ws测试网站，连接 ws://127.0.0.1:9494/ws?systemId=tiktok ，也可以用其他ws测试网站 http://www.jsons.cn/websocket/  
![image](https://github.com/driftboat/TiktokBarrage/assets/247809/22a97a4f-3222-4d1c-ad4d-8981751f32ef)
 - Unity csharp客户端连接代码在 msgs_csharp (基于besthttp)，已测试通过  
   ![P`KP815`F%$QQ@M A0_SU45](https://github.com/driftboat/Tiktok_Live_Barrage/assets/247809/6be22911-ca24-4a8f-a7a5-cb72444ab328)

## 抓取方案应用简述
- 启动抓取后，客户端用websocket连接 ws://127.0.0.1:9494/ws?systemId=tiktok
- 接收并处理弹幕数据
## Barrage功能说明
- Barrage启动ws服务，打开chrome  
- Barrage注入js，将消息发送到设置的ws服务地址
- 任意连接到Barrage的ws客户端，都会收到转发的消息
- 如果要发送到自己的ws服务器，可以关闭“开启本机ws服务”  ，填入自己的ws服务器地址  
 
# Server
 服务器端(单台服务器24小时压测支持1000主播，1000 qps) Window测试版： https://github.com/driftboat/TiktokBarrage/releases/download/1.0.0/BarrageServer.zip  
 - 运行需要redis， 修改conf/app.ini下的redis配置 如果redis6.0(ACL)， 开启username设置用户名  
 - 周排行榜支持，至多2个排行榜

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


# Live message data structures
All data structures are in the msgs directory
![avatar](images/tk2.png)

# ToDo

