using System.Collections.Generic;

namespace Msgs
{
    public class WebcastChatMessage
    {
        public string Content { get; set; }
        public string ContentLanguage { get; set; }
        public string FullScreenTextColor { get; set; }
        public double InputType { get; set; }
        public double QuickChatScene { get; set; }
        public double CommunityFlaggedStatus { get; set; }
        public UserIdentity UserIdentity { get; set; }
        public Common Common { get; set; }
        public User User { get; set; }
        public bool VisibleToSender { get; set; }
        public List<CommentQualityScores> CommentQualityScoresList { get; set; }
    }

    public class UserIdentity
    {
        public bool IsModeratorOfAnchor { get; set; }
        public bool IsAnchor { get; set; }
        public bool IsGiftGiverOfAnchor { get; set; }
        public bool IsSubscriberOfAnchor { get; set; }
        public bool IsMutualFollowingWithAnchor { get; set; }
        public bool IsFollowerOfAnchor { get; set; }
    }

    public class Common
    {
        public string MsgProcessFilterV { get; set; }
        public string AnchorFoldTypeForWeb { get; set; }
        public string RoomId { get; set; }
        public bool IsShowMsg { get; set; }
        public string Describe { get; set; }
        public string FoldType { get; set; }
        public string AnchorFoldType { get; set; }
        public string PriorityScore { get; set; }
        public string ToIdc { get; set; }
        public string ClientSendTime { get; set; }
        public string Method { get; set; }
        public double Monitor { get; set; }
        public string AnchorPriorityScore { get; set; }
        public string FoldTypeForWeb { get; set; }
        public double DispatchStrategy { get; set; }
        public string CreateTime { get; set; }
        public string MsgProcessFilterK { get; set; }
        public string FromIdc { get; set; }
        public string RoomMessageHeatLevel { get; set; }
        public string MsgId { get; set; }
        public string LogId { get; set; }
    }

    public class User
    {
        public string VerifiedReason { get; set; }
        public double Status { get; set; }
        public string ModifyTime { get; set; }
        public string Deprecated21 { get; set; }
        public bool AllowShowInGossip { get; set; }
        public string EnableIchatImg { get; set; }
        public string Deprecated29 { get; set; }
        public bool Deprecated17 { get; set; }
        public string Nickname { get; set; }
        public double Deprecated8 { get; set; }
        public string TicketCount { get; set; }
        public bool AllowFindByContacts { get; set; }
        public UserAttr UserAttr { get; set; }
        public bool AllowUseLinkmic { get; set; }
        public bool WithCarManagementPermission { get; set; }
        public AvatarThumb AvatarThumb { get; set; }
        public bool Verified { get; set; }
        public string CreateTime { get; set; }
        public double UserRole { get; set; }
        public double CommentRestrict { get; set; }
        public double IchatRestrictType { get; set; }
        public string ScmLabel { get; set; }
        public string ShareQrcodeUri { get; set; }
        public bool AllowStrangeComment { get; set; }
        public string Deprecated18 { get; set; }
        public double Exp { get; set; }
        public bool PushVideoPost { get; set; }
        public double Deprecated3 { get; set; }
        public string Deprecated5 { get; set; }
        public double LinkMicStats { get; set; }
        public bool IsFollowing { get; set; }
        public bool NeedProfileGuide { get; set; }
        public bool PushStatus { get; set; }
        public string DisplayId { get; set; }
        public bool WithFusionShopEntry { get; set; }
        public bool Deprecated19 { get; set; }
        public bool AllowUnfollowerComment { get; set; }
        public string Deprecated7 { get; set; }
        public double Secret { get; set; }
        public string TopVipNo { get; set; }
        public string PayScore { get; set; }
        public string VerifiedContent { get; set; }
        public bool AllowOthersDownloadWhenSharingVideo { get; set; }
        public string BgImgUrl { get; set; }
        public bool FoldStrangerChat { get; set; }
        public bool Deprecated28 { get; set; }
        public bool IsFollower { get; set; }
        public string Id { get; set; }
        public bool WithCommercePermission { get; set; }
        public bool Deprecated16 { get; set; }
        public string FanTicketCount { get; set; }
        public string FollowStatus { get; set; }
        public string PayScores { get; set; }
        public double Deprecated2 { get; set; }
        public bool IsBlock { get; set; }
        public double Deprecated12 { get; set; }
        public bool AllowShowMyAction { get; set; }
        public string IdStr { get; set; }
        public bool PushFollow { get; set; }
        public string Constellation { get; set; }
        public double DisableIchat { get; set; }
        public bool PushDigg { get; set; }
        public string BioDescription { get; set; }
        public string Deprecated4 { get; set; }
        public string SecUid { get; set; }
        public double Deprecated15 { get; set; }
        public bool AllowOthersDownloadVideo { get; set; }
        public double BlockStatus { get; set; }
        public bool PushCommentStatus { get; set; }
        public bool PushFriendAction { get; set; }
        public bool PushVideoRecommend { get; set; }
        public string Deprecated1 { get; set; }
        public double Deprecated6 { get; set; }
        public string SpecialId { get; set; }
        public double Deprecated13 { get; set; }
        public bool AllowShareShowProfile { get; set; }
        public bool PushIchat { get; set; }
    }

    public class UserAttr
    {
        public bool IsMuted { get; set; }
        public bool IsAdmin { get; set; }
        public bool IsSuperAdmin { get; set; }
        public string MuteDuration { get; set; }
    }

    public class AvatarThumb
    {
        public string AvgColor { get; set; }
        public double ImageType { get; set; }
        public string OpenWebUrl { get; set; }
        public bool IsAnimated { get; set; }
        public List<string> UrlListList { get; set; }
        public string Uri { get; set; }
        public string Height { get; set; }
        public string Width { get; set; }
    }

    public class CommentQualityScores
    {
        public string Version { get; set; }
        public string Score { get; set; }
    }
}