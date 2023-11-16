using System;

namespace Msgs
{
    [Serializable]
    public class WebcastGiftMessage
    {
        public string streamer;
        public bool disableGiftTracking;
        public string giftId;
        public string repeatCount;
        public User user;
        public double repeatEnd;
        public Gift gift;
        public string fanTicketCount;
        public string groupCount;
        public double linkmicGiftExpressionStrategy;
        public string monitorExtra;
        public string orderId;
        public Common common;
        public MonitorInfo monitorInfo;

    }
    [Serializable]
    public class Gift
    {
        public bool combo;
        public string primaryEffectId;
        public string name;
        public bool isBoxGift;
        public string giftRankRecommendInfo;
        public double giftSubType;
        public bool forLinkmic;
        public double type;
        public double diamondCount;
        public bool isRandomGift;
        public string duration;
        public string id;
        public string goldEffect;
        public string describe;
        public bool isDisplayedOnPanel;
        public bool isBroadcastGift;
        public bool isEffectBefview;
        public bool canPutInGiftBox;
    }
    [Serializable]
    public class MonitorInfo
    { 
        public string anchorId;
        public string sendGiftProfitApiStartMs;
        public string fromVersion;
        public string fromPlatform;
        public string profitapiMessageDur;
        public string sendGiftProfitCoreStartMs;
        public string sendGiftReqStartMs;
        public string sendGiftSendMessageSuccessMs;
        public string sendProfitapiDur;
        public string toUserId;
        public string sendGiftStartClientLocalMs;
    }

}