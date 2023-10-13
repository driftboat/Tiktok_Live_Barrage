using System.Collections.Generic;

namespace Msgs
{
    [System.Serializable]
    public class WebcastMemberMessage
    {
        public string streamer;
        public Common common;
        public User user;
    }

}