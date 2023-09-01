package msgs
type WebcastRoomMessage struct {
	Common struct {
		MsgId string `json:"msgId"`
		CreateTime string `json:"createTime"`
		IsShowMsg bool `json:"isShowMsg"`
		DisplayText struct {
			Key string `json:"key"`
			DefaultPattern string `json:"defaultPattern"`
			DefaultFormat struct {
				Bold bool `json:"bold"`
				Italic bool `json:"italic"`
				Weight float64 `json:"weight"`
				ItalicAngle float64 `json:"italicAngle"`
				FontSize float64 `json:"fontSize"`
				UseHeighLightColor bool `json:"useHeighLightColor"`
				UseRemoteClor bool `json:"useRemoteClor"`
				Color string `json:"color"`
			} `json:"defaultFormat"`
		} `json:"displayText"`
		FoldType string `json:"foldType"`
		FromIdc string `json:"fromIdc"`
		ToIdc string `json:"toIdc"`
		AnchorFoldTypeForWeb string `json:"anchorFoldTypeForWeb"`
		DispatchStrategy float64 `json:"dispatchStrategy"`
		AnchorFoldType string `json:"anchorFoldType"`
		PriorityScore string `json:"priorityScore"`
		LogId string `json:"logId"`
		FoldTypeForWeb string `json:"foldTypeForWeb"`
		ClientSendTime string `json:"clientSendTime"`
		Method string `json:"method"`
		RoomId string `json:"roomId"`
		MsgProcessFilterK string `json:"msgProcessFilterK"`
		AnchorPriorityScore string `json:"anchorPriorityScore"`
		RoomMessageHeatLevel string `json:"roomMessageHeatLevel"`
		Monitor float64 `json:"monitor"`
		Describe string `json:"describe"`
		MsgProcessFilterV string `json:"msgProcessFilterV"`
	} `json:"common"`
	Content string `json:"content"`
	SupprotLandscape bool `json:"supprotLandscape"`
	Source string `json:"source"`
	Scene string `json:"scene"`
	IsWelcome bool `json:"isWelcome"`
}
