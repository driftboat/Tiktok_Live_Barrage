#if !BESTHTTP_DISABLE_WEBSOCKET

using BestHTTP;
using BestHTTP.Examples;
using BestHTTP.WebSocket;
using System;
using System.Reflection;
using UnityEngine;
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
        SignalBus.GlobalSignal.Dispatch(SignalDefine.onWebSocketOpened);
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
            return;
        }
        var wsMsg = JsonUtility.FromJson<Msgs.WebcastBaseMessage>(message);
        if (wsMsg.common.Method == "WebcastChatMessage")
        {
             var chatMsg = (JsonUtility.FromJson<Msgs.WebcastChatMessage>(message));
        }
        else if (wsMsg.common.Method == "WebcastGiftMessage")
        {
            var giftMsg = (JsonUtility.FromJson<Msgs.WebcastGiftMessage>(message));
        }
        else if (wsMsg.common.Method == "WebcastLikeMessage") {
            var likeMsg = (JsonUtility.FromJson<Msgs.WebcastLikeMessage>(message));
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