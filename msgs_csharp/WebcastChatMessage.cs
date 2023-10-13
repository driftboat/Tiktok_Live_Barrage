using System;
using System.Collections.Generic;

namespace Msgs
{
    [Serializable]
    public class WebcastChatMessage
    {
        public string streamer;
        public string content;
        public string contentLanguage;
        public string fullScreenTextColor;
        public double inputType;
        public double quickChatScene;
        public double communityFlaggedStatus;
        public UserIdentity userIdentity;
        public Common common;
        public User user;
        public bool visibleToSender;
        public List<CommentQualityScores> commentQualityScoresList;
    }
    [Serializable]
    public class UserIdentity
    {
        public bool isModeratorOfAnchor;
        public bool isAnchor;
        public bool isGiftGiverOfAnchor;
        public bool isSubscriberOfAnchor;
        public bool isMutualFollowingWithAnchor;
        public bool isFollowerOfAnchor;
    }
    [Serializable]
    public class Common
    {
        public string msgProcessFilterV;
        public string anchorFoldTypeForWeb;
        public string roomId;
        public bool isShowMsg;
        public string describe;
        public string foldType;
        public string anchorFoldType;
        public string priorityScore;
        public string toIdc;
        public string clientSendTime;
        public string method;
        public double monitor;
        public string anchorPriorityScore;
        public string foldTypeForWeb;
        public double dispatchStrategy;
        public string createTime;
        public string msgProcessFilterK;
        public string fromIdc;
        public string roomMessageHeatLevel;
        public string msgId;
        public string logId;
    }
    [Serializable]
    public class User
    {
        public string verifiedReason;
        public double status;
        public string modifyTime;
        public string deprecated21;
        public bool allowShowInGossip;
        public string enableIchatImg;
        public string deprecated29;
        public bool deprecated17;
        public string nickname;
        public double deprecated8;
        public string ticketCount;
        public bool allowFindByContacts;
        public UserAttr userAttr;
        public bool allowUseLinkmic;
        public bool withCarManagementPermission;
        public AvatarThumb avatarThumb;
        public bool verified;
        public string createTime;
        public double userRole;
        public double commentRestrict;
        public double ichatRestrictType;
        public string scmLabel;
        public string shareQrcodeUri;
        public bool allowStrangeComment;
        public string deprecated18;
        public double exp;
        public bool pushVideoPost;
        public double deprecated3;
        public string deprecated5;
        public double linkMicStats;
        public bool isFollowing;
        public bool needProfileGuide;
        public bool pushStatus;
        public string displayId;
        public bool withFusionShopEntry;
        public bool deprecated19;
        public bool allowUnfollowerComment;
        public string deprecated7;
        public double secret;
        public string topVipNo;
        public string payScore;
        public string verifiedContent;
        public bool allowOthersDownloadWhenSharingVideo;
        public string bgImgUrl;
        public bool foldStrangerChat;
        public bool deprecated28;
        public bool isFollower;
        public string id;
        public bool withCommercePermission;
        public bool deprecated16;
        public string fanTicketCount;
        public string followStatus;
        public string payScores;
        public double deprecated2;
        public bool isBlock;
        public double deprecated12;
        public bool allowShowMyAction;
        public string idStr;
        public bool pushFollow;
        public string constellation;
        public double disableIchat;
        public bool pushDigg;
        public string bioDescription;
        public string deprecated4;
        public string secUid;
        public double deprecated15;
        public bool allowOthersDownloadVideo;
        public double blockStatus;
        public bool pushCommentStatus;
        public bool pushFriendAction;
        public bool pushVideoRecommend;
        public string deprecated1;
        public double deprecated6;
        public string specialId;
        public double deprecated13;
        public bool allowShareShowProfile;
        public bool pushIchat;
    }
    [Serializable]
    public class UserAttr
    {
        public bool isMuted;
        public bool isAdmin;
        public bool isSuperAdmin;
        public string muteDuration;
    }
    [Serializable]
    public class AvatarThumb
    {
        public string avgColor;
        public double imageType;
        public string openWebUrl;
        public bool isAnimated;
        public List<string> urlListList;
        public string uri;
        public string height;
        public string width;
    }
    [Serializable]
    public class CommentQualityScores
    {
        public string version;
        public string score;
    }
}