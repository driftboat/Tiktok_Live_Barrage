using System.Collections.Generic;

namespace Msgs
{
    public class WebcastLikeMessage
    {
        public Common Common { get; set; }
        public User User { get; set; }
        public string Icon { get; set; }
        public string EffectCnt { get; set; }
        public string Count { get; set; }
        public string Total { get; set; }
        public string Color { get; set; }
    }

    public class User
    {
        public string TopVipNo { get; set; }
        public bool Deprecated19 { get; set; }
        public bool AllowOthersDownloadVideo { get; set; }
        public double IchatRestrictType { get; set; }
        public string ModifyTime { get; set; }
        public string TicketCount { get; set; }
        public bool Deprecated28 { get; set; }
        public string IdStr { get; set; }
        public bool PushIchat { get; set; }
        public double Deprecated3 { get; set; }
        public string PayScore { get; set; }
        public string Deprecated21 { get; set; }
        public string Deprecated29 { get; set; }
        public bool WithCarManagementPermission { get; set; }
        public double Deprecated8 { get; set; }
        public AvatarThumb AvatarThumb { get; set; }
        public string SecUid { get; set; }
        public double Deprecated13 { get; set; }
        public bool AllowUnfollowerComment { get; set; }
        public bool Deprecated16 { get; set; }
        public string FanTicketCount { get; set; }
        public double Deprecated6 { get; set; }
        public string Deprecated1 { get; set; }
        public string Nickname { get; set; }
    }

    public class AvatarThumb
    {
        public string Width { get; set; }
        public string AvgColor { get; set; }
        public double ImageType { get; set; }
        public string OpenWebUrl { get; set; }
        public bool IsAnimated { get; set; }
        public List<string> UrlListList { get; set; }
        public string Uri { get; set; }
        public string Height { get; set; }
    }

    public class Common
    {
        public string PriorityScore { get; set; }
        public string MsgProcessFilterK { get; set; }
        public string FromIdc { get; set; }
        public string ToIdc { get; set; }
        public string RoomMessageHeatLevel { get; set; }
        public double DispatchStrategy { get; set; }
        public string MsgId { get; set; }
        public string RoomId { get; set; }
        public string CreateTime { get; set; }
        public double Monitor { get; set; }
        public string AnchorFoldType { get; set; }
        public string MsgProcessFilterV { get; set; }
        public string Method { get; set; }
        public bool IsShowMsg { get; set; }
        public string LogId { get; set; }
        public string ClientSendTime { get; set; }
        public string Describe { get; set; }
        public string FoldType { get; set; }
        public string AnchorPriorityScore { get; set; }
        public string FoldTypeForWeb { get; set; }
        public string AnchorFoldTypeForWeb { get; set; }
    }
}