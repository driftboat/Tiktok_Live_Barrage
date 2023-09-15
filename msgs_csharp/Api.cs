public class Api
{
    private Uri BindToGroupURI = new Uri(  "http://127.0.0.1:9494/api/bind_to_group");
    public void BindToGroup(string groupName, string clientId, string sendUserId, Action<CommonResponse> complete)
    {

        HTTPRequest request = new HTTPRequest(BindToGroupURI, HTTPMethods.Post, (req, resp) =>
        {
            // All error handling omitted for clarity!

            Debug.Log("Response: " + resp.DataAsText);
            var cr = JsonUtility.FromJson<CommonResponse>(resp.DataAsText);
            complete(cr);
        });
        request.AddHeader("Content-Type", "application/json");
        request.AddHeader("SystemId", "tiktok");
        string data = $"{{\"clientId\":\"{clientId}\",\"sendUserId\":\"{sendUserId}\",\"groupName\":\"{groupName}\"}}";
        request.RawData = Encoding.UTF8.GetBytes(data);
        request.Send();
    }
}
