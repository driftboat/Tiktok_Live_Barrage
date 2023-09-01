package msgs
type WebcastMemberMessage struct {
	RankScore string `json:"rankScore"`
	TopUserNo string `json:"topUserNo"`
	DisplayStyle float64 `json:"displayStyle"`
	User struct {
		Deprecated2 float64 `json:"deprecated2"`
		Deprecated8 float64 `json:"deprecated8"`
		WithCommercePermission bool `json:"withCommercePermission"`
		AllowShowInGossip bool `json:"allowShowInGossip"`
		PushIchat bool `json:"pushIchat"`
		PushVideoPost bool `json:"pushVideoPost"`
		Status float64 `json:"status"`
		AllowStrangeComment bool `json:"allowStrangeComment"`
		NeedProfileGuide bool `json:"needProfileGuide"`
		PushVideoRecommend bool `json:"pushVideoRecommend"`
		PayScores string `json:"payScores"`
		Nickname string `json:"nickname"`
		BioDescription string `json:"bioDescription"`
		Deprecated5 string `json:"deprecated5"`
		SpecialId string `json:"specialId"`
		AllowUseLinkmic bool `json:"allowUseLinkmic"`
		BgImgUrl string `json:"bgImgUrl"`
		FoldStrangerChat bool `json:"foldStrangerChat"`
		VerifiedReason string `json:"verifiedReason"`
		PayGrade struct {
			Name string `json:"name"`
			Deprecated24 string `json:"deprecated24"`
			ScreenChatType string `json:"screenChatType"`
			Deprecated26 string `json:"deprecated26"`
			Score string `json:"score"`
			Deprecated20 string `json:"deprecated20"`
			Level string `json:"level"`
			Deprecated23 string `json:"deprecated23"`
			NextPrivileges string `json:"nextPrivileges"`
			NextName string `json:"nextName"`
			GradeDescribe string `json:"gradeDescribe"`
			UpgradeNeedConsume string `json:"upgradeNeedConsume"`
			GradeBanner string `json:"gradeBanner"`
			Deprecated22 string `json:"deprecated22"`
			Deprecated25 string `json:"deprecated25"`
		} `json:"payGrade"`
		Constellation string `json:"constellation"`
		PushFriendAction bool `json:"pushFriendAction"`
		ShareQrcodeUri string `json:"shareQrcodeUri"`
		Deprecated12 float64 `json:"deprecated12"`
		AllowShowMyAction bool `json:"allowShowMyAction"`
		CommentRestrict float64 `json:"commentRestrict"`
		FollowInfo struct {
			FollowingCount string `json:"followingCount"`
			FollowerCount string `json:"followerCount"`
			FollowStatus string `json:"followStatus"`
			PushStatus string `json:"pushStatus"`
		} `json:"followInfo"`
		TopVipNo string `json:"topVipNo"`
		PayScore string `json:"payScore"`
		TicketCount string `json:"ticketCount"`
		VerifiedContent string `json:"verifiedContent"`
		AllowShareShowProfile bool `json:"allowShareShowProfile"`
		Deprecated28 bool `json:"deprecated28"`
		PushCommentStatus bool `json:"pushCommentStatus"`
		LinkMicStats float64 `json:"linkMicStats"`
		WithFusionShopEntry bool `json:"withFusionShopEntry"`
		SecUid string `json:"secUid"`
		UserRole float64 `json:"userRole"`
		Deprecated16 bool `json:"deprecated16"`
		IsFollower bool `json:"isFollower"`
		Verified bool `json:"verified"`
		AllowOthersDownloadWhenSharingVideo bool `json:"allowOthersDownloadWhenSharingVideo"`
		AllowUnfollowerComment bool `json:"allowUnfollowerComment"`
		BlockStatus float64 `json:"blockStatus"`
		WithCarManagementPermission bool `json:"withCarManagementPermission"`
		ScmLabel string `json:"scmLabel"`
		Deprecated4 string `json:"deprecated4"`
		ModifyTime string `json:"modifyTime"`
		Deprecated19 bool `json:"deprecated19"`
		FollowStatus string `json:"followStatus"`
		IsFollowing bool `json:"isFollowing"`
		PushFollow bool `json:"pushFollow"`
		Deprecated3 float64 `json:"deprecated3"`
		Deprecated6 float64 `json:"deprecated6"`
		Deprecated13 float64 `json:"deprecated13"`
		AllowFindByContacts bool `json:"allowFindByContacts"`
		EnableIchatImg string `json:"enableIchatImg"`
		PushStatus bool `json:"pushStatus"`
		Deprecated17 bool `json:"deprecated17"`
		Id string `json:"id"`
		Deprecated1 string `json:"deprecated1"`
		AvatarThumb struct {
			AvgColor string `json:"avgColor"`
			ImageType float64 `json:"imageType"`
			OpenWebUrl string `json:"openWebUrl"`
			IsAnimated bool `json:"isAnimated"`
			UrlListList []string `json:"urlListList"`
			Uri string `json:"uri"`
			Height string `json:"height"`
			Width string `json:"width"`
		} `json:"avatarThumb"`
		Secret float64 `json:"secret"`
		Exp float64 `json:"exp"`
		IchatRestrictType float64 `json:"ichatRestrictType"`
		IsBlock bool `json:"isBlock"`
		UserAttr struct {
			IsMuted bool `json:"isMuted"`
			IsAdmin bool `json:"isAdmin"`
			IsSuperAdmin bool `json:"isSuperAdmin"`
			MuteDuration string `json:"muteDuration"`
		} `json:"userAttr"`
		Deprecated18 string `json:"deprecated18"`
		DisableIchat float64 `json:"disableIchat"`
		Deprecated29 string `json:"deprecated29"`
		PushDigg bool `json:"pushDigg"`
		Deprecated21 string `json:"deprecated21"`
		Deprecated15 float64 `json:"deprecated15"`
		AllowOthersDownloadVideo bool `json:"allowOthersDownloadVideo"`
		FanTicketCount string `json:"fanTicketCount"`
		Deprecated7 string `json:"deprecated7"`
		CreateTime string `json:"createTime"`
		DisplayId string `json:"displayId"`
		IdStr string `json:"idStr"`
	} `json:"user"`
	MemberCount string `json:"memberCount"`
	IsSetToAdmin bool `json:"isSetToAdmin"`
	Action string `json:"action"`
	ActionDescription string `json:"actionDescription"`
	UserId string `json:"userId"`
	PopStr string `json:"popStr"`
	ActionDuration string `json:"actionDuration"`
	Common struct {
		AnchorFoldType string `json:"anchorFoldType"`
		PriorityScore string `json:"priorityScore"`
		MsgProcessFilterV string `json:"msgProcessFilterV"`
		FromIdc string `json:"fromIdc"`
		AnchorFoldTypeForWeb string `json:"anchorFoldTypeForWeb"`
		DispatchStrategy float64 `json:"dispatchStrategy"`
		Method string `json:"method"`
		Monitor float64 `json:"monitor"`
		LogId string `json:"logId"`
		ToIdc string `json:"toIdc"`
		IsShowMsg bool `json:"isShowMsg"`
		Describe string `json:"describe"`
		DisplayText struct {
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
			PiecesList []*struct {
				UserValue struct {
					User struct {
						Deprecated8 float64 `json:"deprecated8"`
						FollowStatus string `json:"followStatus"`
						TopVipNo string `json:"topVipNo"`
						PayScore string `json:"payScore"`
						IsFollowing bool `json:"isFollowing"`
						PushCommentStatus bool `json:"pushCommentStatus"`
						VerifiedReason string `json:"verifiedReason"`
						FollowInfo struct {
							FollowingCount string `json:"followingCount"`
							FollowerCount string `json:"followerCount"`
							FollowStatus string `json:"followStatus"`
							PushStatus string `json:"pushStatus"`
						} `json:"followInfo"`
						Deprecated13 float64 `json:"deprecated13"`
						AllowShowMyAction bool `json:"allowShowMyAction"`
						AllowStrangeComment bool `json:"allowStrangeComment"`
						FanTicketCount string `json:"fanTicketCount"`
						Deprecated28 bool `json:"deprecated28"`
						Nickname string `json:"nickname"`
						Secret float64 `json:"secret"`
						Deprecated21 string `json:"deprecated21"`
						PayScores string `json:"payScores"`
						Deprecated17 bool `json:"deprecated17"`
						Id string `json:"id"`
						Deprecated5 string `json:"deprecated5"`
						AvatarThumb struct {
							AvgColor string `json:"avgColor"`
							ImageType float64 `json:"imageType"`
							OpenWebUrl string `json:"openWebUrl"`
							IsAnimated bool `json:"isAnimated"`
							UrlListList []string `json:"urlListList"`
							Uri string `json:"uri"`
							Height string `json:"height"`
							Width string `json:"width"`
						} `json:"avatarThumb"`
						ModifyTime string `json:"modifyTime"`
						ShareQrcodeUri string `json:"shareQrcodeUri"`
						TicketCount string `json:"ticketCount"`
						DisplayId string `json:"displayId"`
						UserRole float64 `json:"userRole"`
						Exp float64 `json:"exp"`
						ScmLabel string `json:"scmLabel"`
						WithFusionShopEntry bool `json:"withFusionShopEntry"`
						PushDigg bool `json:"pushDigg"`
						PushVideoPost bool `json:"pushVideoPost"`
						Deprecated2 float64 `json:"deprecated2"`
						WithCommercePermission bool `json:"withCommercePermission"`
						IsFollower bool `json:"isFollower"`
						PushFriendAction bool `json:"pushFriendAction"`
						PushStatus bool `json:"pushStatus"`
						Verified bool `json:"verified"`
						Deprecated18 string `json:"deprecated18"`
						BlockStatus float64 `json:"blockStatus"`
						PushFollow bool `json:"pushFollow"`
						BioDescription string `json:"bioDescription"`
						Deprecated3 float64 `json:"deprecated3"`
						Deprecated4 string `json:"deprecated4"`
						SpecialId string `json:"specialId"`
						LinkMicStats float64 `json:"linkMicStats"`
						AllowShareShowProfile bool `json:"allowShareShowProfile"`
						EnableIchatImg string `json:"enableIchatImg"`
						Deprecated29 string `json:"deprecated29"`
						PushIchat bool `json:"pushIchat"`
						IsBlock bool `json:"isBlock"`
						IchatRestrictType float64 `json:"ichatRestrictType"`
						Deprecated7 string `json:"deprecated7"`
						Status float64 `json:"status"`
						PayGrade struct {
							Deprecated26 string `json:"deprecated26"`
							GradeDescribe string `json:"gradeDescribe"`
							Name string `json:"name"`
							Deprecated22 string `json:"deprecated22"`
							UpgradeNeedConsume string `json:"upgradeNeedConsume"`
							NextName string `json:"nextName"`
							Deprecated23 string `json:"deprecated23"`
							Deprecated24 string `json:"deprecated24"`
							Deprecated25 string `json:"deprecated25"`
							Deprecated20 string `json:"deprecated20"`
							Level string `json:"level"`
							NextPrivileges string `json:"nextPrivileges"`
							Score string `json:"score"`
							GradeBanner string `json:"gradeBanner"`
							ScreenChatType string `json:"screenChatType"`
						} `json:"payGrade"`
						UserAttr struct {
							IsMuted bool `json:"isMuted"`
							IsAdmin bool `json:"isAdmin"`
							IsSuperAdmin bool `json:"isSuperAdmin"`
							MuteDuration string `json:"muteDuration"`
						} `json:"userAttr"`
						SecUid string `json:"secUid"`
						Deprecated19 bool `json:"deprecated19"`
						BgImgUrl string `json:"bgImgUrl"`
						Deprecated16 bool `json:"deprecated16"`
						VerifiedContent string `json:"verifiedContent"`
						CommentRestrict float64 `json:"commentRestrict"`
						Constellation string `json:"constellation"`
						FoldStrangerChat bool `json:"foldStrangerChat"`
						AllowShowInGossip bool `json:"allowShowInGossip"`
						AllowUnfollowerComment bool `json:"allowUnfollowerComment"`
						WithCarManagementPermission bool `json:"withCarManagementPermission"`
						Deprecated1 string `json:"deprecated1"`
						CreateTime string `json:"createTime"`
						Deprecated15 float64 `json:"deprecated15"`
						AllowOthersDownloadWhenSharingVideo bool `json:"allowOthersDownloadWhenSharingVideo"`
						DisableIchat float64 `json:"disableIchat"`
						Deprecated12 float64 `json:"deprecated12"`
						AllowOthersDownloadVideo bool `json:"allowOthersDownloadVideo"`
						AllowUseLinkmic bool `json:"allowUseLinkmic"`
						IdStr string `json:"idStr"`
						PushVideoRecommend bool `json:"pushVideoRecommend"`
						Deprecated6 float64 `json:"deprecated6"`
						AllowFindByContacts bool `json:"allowFindByContacts"`
						NeedProfileGuide bool `json:"needProfileGuide"`
					} `json:"user"`
					WithColon bool `json:"withColon"`
				} `json:"userValue"`
				Type float64 `json:"type"`
				Format struct {
					UseHeighLightColor bool `json:"useHeighLightColor"`
					UseRemoteClor bool `json:"useRemoteClor"`
					Color string `json:"color"`
					Bold bool `json:"bold"`
					Italic bool `json:"italic"`
					Weight float64 `json:"weight"`
					ItalicAngle float64 `json:"italicAngle"`
					FontSize float64 `json:"fontSize"`
				} `json:"format"`
				StringValue string `json:"stringValue"`
			} `json:"piecesList"`
		} `json:"displayText"`
		FoldType string `json:"foldType"`
		MsgProcessFilterK string `json:"msgProcessFilterK"`
		ClientSendTime string `json:"clientSendTime"`
		MsgId string `json:"msgId"`
		RoomId string `json:"roomId"`
		CreateTime string `json:"createTime"`
		AnchorPriorityScore string `json:"anchorPriorityScore"`
		RoomMessageHeatLevel string `json:"roomMessageHeatLevel"`
		FoldTypeForWeb string `json:"foldTypeForWeb"`
	} `json:"common"`
	IsTopUser bool `json:"isTopUser"`
	EnterType string `json:"enterType"`
	KickSource float64 `json:"kickSource"`
	AllowPreviewTime string `json:"allowPreviewTime"`
	ClientEnterSource string `json:"clientEnterSource"`
	ClientEnterType string `json:"clientEnterType"`
	AnchorDisplayText struct {
		Key string `json:"key"`
		DefaultPattern string `json:"defaultPattern"`
		DefaultFormat struct {
			Italic bool `json:"italic"`
			Weight float64 `json:"weight"`
			ItalicAngle float64 `json:"italicAngle"`
			FontSize float64 `json:"fontSize"`
			UseHeighLightColor bool `json:"useHeighLightColor"`
			UseRemoteClor bool `json:"useRemoteClor"`
			Color string `json:"color"`
			Bold bool `json:"bold"`
		} `json:"defaultFormat"`
		PiecesList []*struct {
			StringValue string `json:"stringValue"`
			UserValue struct {
				User struct {
					TopVipNo string `json:"topVipNo"`
					UserAttr struct {
						MuteDuration string `json:"muteDuration"`
						IsMuted bool `json:"isMuted"`
						IsAdmin bool `json:"isAdmin"`
						IsSuperAdmin bool `json:"isSuperAdmin"`
					} `json:"userAttr"`
					PayScore string `json:"payScore"`
					AllowUnfollowerComment bool `json:"allowUnfollowerComment"`
					PushFriendAction bool `json:"pushFriendAction"`
					Deprecated4 string `json:"deprecated4"`
					AvatarThumb struct {
						IsAnimated bool `json:"isAnimated"`
						UrlListList []string `json:"urlListList"`
						Uri string `json:"uri"`
						Height string `json:"height"`
						Width string `json:"width"`
						AvgColor string `json:"avgColor"`
						ImageType float64 `json:"imageType"`
						OpenWebUrl string `json:"openWebUrl"`
					} `json:"avatarThumb"`
					ScmLabel string `json:"scmLabel"`
					AllowShowMyAction bool `json:"allowShowMyAction"`
					ShareQrcodeUri string `json:"shareQrcodeUri"`
					SecUid string `json:"secUid"`
					FanTicketCount string `json:"fanTicketCount"`
					Deprecated29 string `json:"deprecated29"`
					PushDigg bool `json:"pushDigg"`
					PushVideoRecommend bool `json:"pushVideoRecommend"`
					Deprecated17 bool `json:"deprecated17"`
					WithCarManagementPermission bool `json:"withCarManagementPermission"`
					WithCommercePermission bool `json:"withCommercePermission"`
					AllowStrangeComment bool `json:"allowStrangeComment"`
					ModifyTime string `json:"modifyTime"`
					DisplayId string `json:"displayId"`
					Deprecated12 float64 `json:"deprecated12"`
					AllowOthersDownloadWhenSharingVideo bool `json:"allowOthersDownloadWhenSharingVideo"`
					Deprecated18 string `json:"deprecated18"`
					Nickname string `json:"nickname"`
					CreateTime string `json:"createTime"`
					FoldStrangerChat bool `json:"foldStrangerChat"`
					SpecialId string `json:"specialId"`
					Deprecated15 float64 `json:"deprecated15"`
					AllowUseLinkmic bool `json:"allowUseLinkmic"`
					DisableIchat float64 `json:"disableIchat"`
					EnableIchatImg string `json:"enableIchatImg"`
					Verified bool `json:"verified"`
					PayGrade struct {
						Level string `json:"level"`
						Deprecated26 string `json:"deprecated26"`
						ScreenChatType string `json:"screenChatType"`
						Score string `json:"score"`
						Deprecated20 string `json:"deprecated20"`
						NextName string `json:"nextName"`
						Deprecated22 string `json:"deprecated22"`
						Deprecated24 string `json:"deprecated24"`
						UpgradeNeedConsume string `json:"upgradeNeedConsume"`
						GradeBanner string `json:"gradeBanner"`
						Name string `json:"name"`
						Deprecated23 string `json:"deprecated23"`
						NextPrivileges string `json:"nextPrivileges"`
						Deprecated25 string `json:"deprecated25"`
						GradeDescribe string `json:"gradeDescribe"`
					} `json:"payGrade"`
					FollowStatus string `json:"followStatus"`
					IchatRestrictType float64 `json:"ichatRestrictType"`
					Deprecated2 float64 `json:"deprecated2"`
					Deprecated13 float64 `json:"deprecated13"`
					FollowInfo struct {
						PushStatus string `json:"pushStatus"`
						FollowingCount string `json:"followingCount"`
						FollowerCount string `json:"followerCount"`
						FollowStatus string `json:"followStatus"`
					} `json:"followInfo"`
					Deprecated16 bool `json:"deprecated16"`
					Deprecated28 bool `json:"deprecated28"`
					NeedProfileGuide bool `json:"needProfileGuide"`
					PushCommentStatus bool `json:"pushCommentStatus"`
					Deprecated3 float64 `json:"deprecated3"`
					Deprecated5 string `json:"deprecated5"`
					AllowShareShowProfile bool `json:"allowShareShowProfile"`
					Exp float64 `json:"exp"`
					Deprecated7 string `json:"deprecated7"`
					PushIchat bool `json:"pushIchat"`
					BioDescription string `json:"bioDescription"`
					AllowShowInGossip bool `json:"allowShowInGossip"`
					VerifiedContent string `json:"verifiedContent"`
					BgImgUrl string `json:"bgImgUrl"`
					PushStatus bool `json:"pushStatus"`
					Secret float64 `json:"secret"`
					WithFusionShopEntry bool `json:"withFusionShopEntry"`
					PushFollow bool `json:"pushFollow"`
					AllowOthersDownloadVideo bool `json:"allowOthersDownloadVideo"`
					IdStr string `json:"idStr"`
					CommentRestrict float64 `json:"commentRestrict"`
					IsFollower bool `json:"isFollower"`
					IsFollowing bool `json:"isFollowing"`
					IsBlock bool `json:"isBlock"`
					Deprecated1 string `json:"deprecated1"`
					Deprecated6 float64 `json:"deprecated6"`
					UserRole float64 `json:"userRole"`
					PayScores string `json:"payScores"`
					Id string `json:"id"`
					Deprecated8 float64 `json:"deprecated8"`
					Deprecated19 bool `json:"deprecated19"`
					Constellation string `json:"constellation"`
					Deprecated21 string `json:"deprecated21"`
					AllowFindByContacts bool `json:"allowFindByContacts"`
					VerifiedReason string `json:"verifiedReason"`
					TicketCount string `json:"ticketCount"`
					BlockStatus float64 `json:"blockStatus"`
					PushVideoPost bool `json:"pushVideoPost"`
					Status float64 `json:"status"`
					LinkMicStats float64 `json:"linkMicStats"`
				} `json:"user"`
				WithColon bool `json:"withColon"`
			} `json:"userValue"`
			Type float64 `json:"type"`
			Format struct {
				Color string `json:"color"`
				Bold bool `json:"bold"`
				Italic bool `json:"italic"`
				Weight float64 `json:"weight"`
				ItalicAngle float64 `json:"italicAngle"`
				FontSize float64 `json:"fontSize"`
				UseHeighLightColor bool `json:"useHeighLightColor"`
				UseRemoteClor bool `json:"useRemoteClor"`
			} `json:"format"`
		} `json:"piecesList"`
	} `json:"anchorDisplayText"`
	ClientLiveReason string `json:"clientLiveReason"`
	UserShareType string `json:"userShareType"`
}
