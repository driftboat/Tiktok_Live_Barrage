#if !BESTHTTP_DISABLE_WEBSOCKET

using BestHTTP;
using BestHTTP.Examples;
using BestHTTP.WebSocket;
using System;
using System.Reflection;
using UnityEngine;
using Msgs;
using System.IO;

[System.Serializable]
public class ClientID
{
    public string clientId;
}

[System.Serializable]
public class CommonMsg
{
    public string msg; 
}
[System.Serializable]
public class ConnectedMsg
{ 
    public ClientID data;
}

[System.Serializable]
public class ChatMsg
{
    public Msgs.WebcastChatMessage data;
}
[System.Serializable]
public class GiftMsg
{
    public Msgs.WebcastGiftMessage data;
}
[System.Serializable]
public class LikeMsg
{
    public Msgs.WebcastLikeMessage data;
}
[System.Serializable]
public class MemberMsg
{
    public Msgs.WebcastMemberMessage data;
}

[System.Serializable]
public class SocialMsg
{
    public Msgs.WebcastSocialMessage data;
}
public class WebSocketManager : MonoBehaviour
{
    #region Private Fields

    /// <summary>
    /// The WebSocket address to connect
    /// </summary> 
    string address = "ws://127.0.0.1:9494/ws?systemId=tiktok";

    /// <summary>
    /// Default text to send
    /// </summary>
    string msgToSend = "Hello World!";


    /// <summary>
    /// Saved WebSocket instance
    /// </summary>
    WebSocket webSocket;

    /// <summary>
    /// GUI scroll position
    /// </summary>
    Vector2 scrollPos;


    #endregion

    #region Unity Events
    private void Start()
    {
        //string path = "Assets/test.txt";
        ////Read the text from directly from the test.txt file
        //StreamReader reader = new StreamReader(path);
        //var text = reader.ReadToEnd();
        //var giftMsg = JsonUtility.FromJson<Msgs.WebcastGiftMessage>(text);
        //Debug.Log(giftMsg.common.method);
        //reader.Close();
        webSocket = new WebSocket(new Uri(address));
#if !UNITY_WEBGL
        //   webSocket.StartPingThread = true;

#if !BESTHTTP_DISABLE_PROXY
        if (HTTPManager.Proxy != null)
            webSocket.InternalRequest.Proxy = new HTTPProxy(HTTPManager.Proxy.Address, HTTPManager.Proxy.Credentials, false);
#endif
#endif

        // Subscribe to the WS events
        webSocket.OnOpen += OnOpen;
        webSocket.OnMessage += OnMessageReceived;
        webSocket.OnClosed += OnClosed;
        webSocket.OnError += OnError;

        // Start connecting to the server
        webSocket.Open();
    }

    void OnDestroy()
    {
        if (webSocket != null)
        {
            webSocket.Close();
        }
    }

    
    #endregion

    #region WebSocket Event Handlers

    /// <summary>
    /// Called when the web socket is open, and we are ready to send and receive data
    /// </summary>
    void OnOpen(WebSocket ws)
    {
    }

    /// <summary>
    /// Called when we received a text message from the server
    /// </summary>
    void OnMessageReceived(WebSocket ws, string message)
    {
        Debug.Log(message);
        var commonMsg = JsonUtility.FromJson<CommonMsg>(message);
        if (commonMsg != null) {
            if (commonMsg.msg == "Connected") {
                var connectedMsg = JsonUtility.FromJson<ConnectedMsg>(message);
                Debug.Log(connectedMsg.data.clientId);
            }
            else if (commonMsg.msg == "Chat")
            {
                var chatMsg = JsonUtility.FromJson<ChatMsg>(message);
                Debug.Log(chatMsg.data.common.method);
            }
            else if (commonMsg.msg == "Gift")
            {
                var giftMsg = JsonUtility.FromJson<GiftMsg>(message);
                Debug.Log(giftMsg.data.common.method);
            }
            else if (commonMsg.msg == "Like")
            {
                var likeMsg = JsonUtility.FromJson<LikeMsg>(message);
                Debug.Log(likeMsg.data.common.method);
            }
            else if (commonMsg.msg == "Member") {
                var memberMsg = JsonUtility.FromJson<MemberMsg>(message);
                Debug.Log(memberMsg.data.common.method);
            }
            else if (commonMsg.msg == "Social"){
                var socialMsg = JsonUtility.FromJson<SocialMsg>(message);
                Debug.Log(socialMsg.data.common.method);
            }
            return;
        }
        
          
    }

    /// <summary>
    /// Called when the web socket closed
    /// </summary>
    void OnClosed(WebSocket ws, UInt16 code, string message)
    { 
        webSocket = null;
    }

    /// <summary>
    /// Called when an error occured on client side
    /// </summary>
    void OnError(WebSocket ws, Exception ex)
    {
        string errorMsg = string.Empty;
#if !UNITY_WEBGL || UNITY_EDITOR
        if (ws.InternalRequest.Response != null)
        {
            errorMsg = string.Format("Status Code from Server: {0} and Message: {1}", ws.InternalRequest.Response.StatusCode, ws.InternalRequest.Response.Message);
        }
#endif

        Debug.LogError(ex != null ? ex.Message +"\n" +ex.StackTrace : "Unknown Error " + errorMsg);
        webSocket = null;
    }

    #endregion
}

#endif