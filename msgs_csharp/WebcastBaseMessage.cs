using System.Collections.Generic;

namespace Msgs
{
    public class WebcastBaseMessage
    {
        public Common Common { get; set; } 
    }
     
    public class Common
    {
        public string Method { get; set; }
    }
}