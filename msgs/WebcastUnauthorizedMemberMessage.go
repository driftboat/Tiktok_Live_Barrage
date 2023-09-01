package msgs
type WebcastUnauthorizedMemberMessage struct {
	Common struct {
		CreateTime string `json:"createTime"`
		AnchorFoldType string `json:"anchorFoldType"`
		LogId string `json:"logId"`
		FromIdc string `json:"fromIdc"`
		AnchorFoldTypeForWeb string `json:"anchorFoldTypeForWeb"`
		Method string `json:"method"`
		Monitor float64 `json:"monitor"`
		Describe string `json:"describe"`
		MsgProcessFilterK string `json:"msgProcessFilterK"`
		RoomMessageHeatLevel string `json:"roomMessageHeatLevel"`
		FoldTypeForWeb string `json:"foldTypeForWeb"`
		DispatchStrategy float64 `json:"dispatchStrategy"`
		MsgId string `json:"msgId"`
		RoomId string `json:"roomId"`
		IsShowMsg bool `json:"isShowMsg"`
		PriorityScore string `json:"priorityScore"`
		AnchorPriorityScore string `json:"anchorPriorityScore"`
		FoldType string `json:"foldType"`
		MsgProcessFilterV string `json:"msgProcessFilterV"`
		ToIdc string `json:"toIdc"`
		ClientSendTime string `json:"clientSendTime"`
	} `json:"common"`
	Action float64 `json:"action"`
	NickNamePrefix struct {
		Key string `json:"key"`
		DefaultPattern string `json:"defaultPattern"`
		DefaultFormat struct {
			UseHeighLightColor bool `json:"useHeighLightColor"`
			UseRemoteClor bool `json:"useRemoteClor"`
			Color string `json:"color"`
			Bold bool `json:"bold"`
			Italic bool `json:"italic"`
			Weight float64 `json:"weight"`
			ItalicAngle float64 `json:"italicAngle"`
			FontSize float64 `json:"fontSize"`
		} `json:"defaultFormat"`
	} `json:"nickNamePrefix"`
	NickName string `json:"nickName"`
	EnterText struct {
		Key string `json:"key"`
		DefaultPattern string `json:"defaultPattern"`
		DefaultFormat struct {
			FontSize float64 `json:"fontSize"`
			UseHeighLightColor bool `json:"useHeighLightColor"`
			UseRemoteClor bool `json:"useRemoteClor"`
			Color string `json:"color"`
			Bold bool `json:"bold"`
			Italic bool `json:"italic"`
			Weight float64 `json:"weight"`
			ItalicAngle float64 `json:"italicAngle"`
		} `json:"defaultFormat"`
	} `json:"enterText"`
}
