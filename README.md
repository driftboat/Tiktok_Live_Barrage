# tkbarrage
2024,2025年全年可用，无需key  
Tiktok弹幕抓取 发送websocket服务器 
Go语言实现，js注入模式和代理两种模式可选，都可以连接使用
  
![avatar](images/tk1.jpg)


# 抓取流程说明
## 抓取服务启动（注入模式）
- [下载安装包](https://github.com/driftboat/TiktokBarrage/releases/latest)
- 解压运行 Barrage.exe
- 关闭所有Chrome浏览器(可以不关闭，但必须在窗口中设置一个单独的Chrome用户目录，专门给Barrage使用）  
- 输入正确的key，点开始(注入模式)运行
## 抓取服务启动（代理模式）
- [下载安装包](https://github.com/driftboat/TiktokBarrage/releases/latest)
- 解压运行 Barrage.exe
- 在chrome浏览器方式启动命令修改，在结尾增加 --ignore-certificate-errors  
   ![image](https://github.com/user-attachments/assets/7d6dcd17-6192-4c5a-a273-ee085129da2c)
- chrome安装Proxy SwitchyOmega插件(从Chrome应用商店安装)，并设置代理(注意这个代理是发给barrage开的7802端口的，必须按下面截图设置为7802)
  ![image](https://github.com/user-attachments/assets/1d30e201-ffdf-4a5c-a56f-240455e04a22)
- 点开Proxy SwitchOmega插件，选择proxy情景  
  ![image](https://github.com/user-attachments/assets/81983325-fa64-41dd-9a40-ffb893796db0)
- 点开始(代理模式)运行，如果你需要科学上网，需要输入代理地址，目前只支持socekt5或者http，例如 socket5://admin:123456@127.0.0.1:8888 或 http://admin:123456@127.0.0.1:8888
## 客户端接收弹幕消息模拟
- 开启模拟ws客户端 用这个在线ws测试网站，连接 ws://127.0.0.1:9494/ws?systemId=tiktok ，也可以用其他ws测试网站 http://www.websocket-test.com/ http://www.jsons.cn/websocket/  
![image](https://github.com/driftboat/TiktokBarrage/assets/247809/22a97a4f-3222-4d1c-ad4d-8981751f32ef)
 - Unity csharp客户端连接代码在 msgs_csharp (基于besthttp)，已测试通过  
   ![P`KP815`F%$QQ@M A0_SU45](https://github.com/driftboat/Tiktok_Live_Barrage/assets/247809/6be22911-ca24-4a8f-a7a5-cb72444ab328)

## 抓取方案应用简述
- 启动抓取后，客户端用websocket连接 ws://127.0.0.1:9494/ws?systemId=tiktok
- 接收并处理弹幕数据
## Barrage功能说明(注入模式)
- Barrage启动ws服务，打开chrome  
- Barrage注入js，将消息发送到设置的ws服务地址
- 任意连接到Barrage的ws客户端，都会收到转发的消息
- 如果要发送到自己的ws服务器，可以关闭“开启本机ws服务”  ，填入自己的ws服务器地址  
 
## Barrage功能说明(代理模式)
- Barrage启动ws服务，开启7802端口  
- 浏览器用Proxy SwitchyOmega等插件，将对应的网络请求转发到barrage的7802端口 
- barrage将对应的网络请求通过配置的科学上网代理转发出去(不需要科学上网忽略) 
- 任意连接到Barrage的ws客户端，都会收到转发的消息
- 如果要发送到自己的ws服务器，可以关闭“开启本机ws服务”  ，填入自己的ws服务器地址  
 

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
消息json格式为：msg 消息标识（Chat，Gift，Like，Member....） data 为具体数据如下  
msg_csharp里面部分字段不全，根据需要可自行填补  
All data structures are in the msgs directory
![avatar](images/tk2.png)

# 错误解决

## 注入模式，无法抓取
- 检查chrome版本，重置chrome设置到默认
- 用获取fun值按钮，获取新的值，修改fun.txt中的值
	- 点开后在chrome的开发者控制台，寻找以下值  
   	![image](https://github.com/user-attachments/assets/6bb52534-228f-488a-9c1e-5820a696f472)
- 联系 qq: 174389107 或者提交issue  

   		 
