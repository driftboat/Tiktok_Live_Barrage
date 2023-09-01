package msgs
type SystemMessage struct {
	Source string `json:"source"`
	Scene string `json:"scene"`
	IsWelcome bool `json:"isWelcome"`
	Common struct {
		FromIdc string `json:"fromIdc"`
		FoldTypeForWeb string `json:"foldTypeForWeb"`
		ClientSendTime string `json:"clientSendTime"`
		Method string `json:"method"`
		MsgId string `json:"msgId"`
		RoomId string `json:"roomId"`
		Monitor float64 `json:"monitor"`
		AnchorFoldType string `json:"anchorFoldType"`
		DispatchStrategy float64 `json:"dispatchStrategy"`
		Describe string `json:"describe"`
		FoldType string `json:"foldType"`
		LogId string `json:"logId"`
		MsgProcessFilterK string `json:"msgProcessFilterK"`
		AnchorFoldTypeForWeb string `json:"anchorFoldTypeForWeb"`
		AnchorPriorityScore string `json:"anchorPriorityScore"`
		RoomMessageHeatLevel string `json:"roomMessageHeatLevel"`
		CreateTime string `json:"createTime"`
		IsShowMsg bool `json:"isShowMsg"`
		PriorityScore string `json:"priorityScore"`
		MsgProcessFilterV string `json:"msgProcessFilterV"`
		ToIdc string `json:"toIdc"`
	} `json:"common"`
	Content string `json:"content"`
	SupprotLandscape bool `json:"supprotLandscape"`
}
