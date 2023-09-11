# tkbarrage
Tiktok弹幕抓取 发送websocket服务器   （之前发布的Barrage.exe存在Bug无法连接自己设置的ws url，请重新下载  ）  
Receive live stream events such as comments and gifts and send to your websocket server tool in Go language.  free for one month   
免费试用一个月 下载客户端  https://github.com/driftboat/TiktokBarrage/releases/download/1.0.0/Barrage.zip  
请联系 qq: 174389107  telegram：db_444444  
Go语言实现，js注入  
  
![avatar](images/tk1.png)

# Server
 服务器端(单台服务器24小时压测支持1000主播，1000 qps) Window测试版： https://github.com/driftboat/TiktokBarrage/releases/download/1.0.0/BarrageServer.zip  
 - 运行需要redis， 修改conf/app.ini下的redis配置 如果redis6.0(ACL)， 开启username设置用户名  

 - 启动服务器，从服务器log获取房间id（如：（{"data":7275920301647465259,"level":"info","msg":"tiktok_chat","time":"2023-09-07 14:29:33"}）（实际使用中由主播在真实客户端填入）
 - 关闭所有chrome浏览器，启动Barrage客户端，连上服务器才能获得上面的房间id  
   ![image](https://github.com/driftboat/TiktokBarrage/assets/247809/17cf5828-0f6d-4496-ad23-e22fb78aaf0b)
 - 开启模拟ws客户端 用这个在线ws测试网站，也可以用其他ws测试网站 http://www.jsons.cn/websocket/  
![image](https://github.com/driftboat/TiktokBarrage/assets/247809/22a97a4f-3222-4d1c-ad4d-8981751f32ef)
 - 绑定模拟客户端到对应房间，修改clientId为前面ws模拟客户端返回的clientid，groupName为服务器log获取的房间id
  ```
   curl --location 'http://127.0.0.1:9494/api/bind_to_group' \
--header 'systemId: tiktok' \
--header 'Content-Type: application/json' \
--data '{
"sendUserId":"123",
"groupName":"7275920301647465259",	
"clientId":"fGpD6UCRFYoKD7BfRCNbq+n/+ltgdWGH9jlHMsaD72SnOJD7ZF5z3WfOk6xSrTww"
}'
```


# How to run
- Download https://github.com/driftboat/TiktokBarrage/releases/download/1.0.0/Barrage.zip  
- Unzip and run Barrage.exe
- Close all chrome browser
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


# Live message data structures
All data structures are in the msgs directory
![avatar](images/tk2.png)

# ToDo

