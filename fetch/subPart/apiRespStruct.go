package dailyboard

type ApiResp struct {
	Result    int        `json:"result"`
	HostName  string     `json:"host-name"`
	RankList  []RankList `json:"rankList"`
	RequestID string     `json:"requestId"`
}
type RecoReason struct {
	Desc       string      `json:"desc"`
	Href       string      `json:"href"`
	Tag        string      `json:"tag"`
	LayoutType int         `json:"layoutType"`
	DescType   interface{} `json:"descType"`
	Type       int         `json:"type"`
}
type VideoList struct {
	Priority             int    `json:"priority"`
	VisibleType          int    `json:"visibleType"`
	DurationMillis       int    `json:"durationMillis"`
	UploadTime           int64  `json:"uploadTime"`
	DanmakuCount         int    `json:"danmakuCount"`
	Title                string `json:"title"`
	SourceStatus         int    `json:"sourceStatus"`
	SizeType             int    `json:"sizeType"`
	DanmakuGuidePosition int    `json:"danmakuGuidePosition"`
	DanmakuCountShow     string `json:"danmakuCountShow"`
	FileName             string `json:"fileName"`
	ID                   string `json:"id"`
}
type TagList struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
type HeadCdnUrls struct {
	URL            string `json:"url"`
	FreeTrafficCdn bool   `json:"freeTrafficCdn"`
}
type SocialMedal struct {
}
type CdnUrls struct {
	URL            string `json:"url"`
	FreeTrafficCdn bool   `json:"freeTrafficCdn"`
}
type ThumbnailImage struct {
	CdnUrls []CdnUrls `json:"cdnUrls"`
}
type UserHeadImgInfo struct {
	Width                int            `json:"width"`
	Height               int            `json:"height"`
	Size                 int            `json:"size"`
	Type                 int            `json:"type"`
	Animated             bool           `json:"animated"`
	ThumbnailImage       ThumbnailImage `json:"thumbnailImage"`
	ThumbnailImageCdnURL string         `json:"thumbnailImageCdnUrl"`
}
type User struct {
	Action               int             `json:"action"`
	Href                 string          `json:"href"`
	SexTrend             int             `json:"sexTrend"`
	NameColor            int             `json:"nameColor"`
	VerifiedTypes        []int           `json:"verifiedTypes"`
	AvatarFramePcImg     string          `json:"avatarFramePcImg"`
	AvatarFrameMobileImg string          `json:"avatarFrameMobileImg"`
	IsFollowing          bool            `json:"isFollowing"`
	ContributeCount      string          `json:"contributeCount"`
	FollowingStatus      int             `json:"followingStatus"`
	AvatarFrame          int             `json:"avatarFrame"`
	HeadURL              string          `json:"headUrl"`
	FanCountValue        int             `json:"fanCountValue"`
	VerifiedType         int             `json:"verifiedType"`
	VerifiedText         string          `json:"verifiedText"`
	Gender               int             `json:"gender"`
	FollowingCount       string          `json:"followingCount"`
	HeadCdnUrls          []HeadCdnUrls   `json:"headCdnUrls"`
	IsJoinUpCollege      bool            `json:"isJoinUpCollege"`
	FollowingCountValue  int             `json:"followingCountValue"`
	ContributeCountValue int             `json:"contributeCountValue"`
	FanCount             string          `json:"fanCount"`
	SocialMedal          SocialMedal     `json:"socialMedal"`
	AvatarImage          string          `json:"avatarImage"`
	UserHeadImgInfo      UserHeadImgInfo `json:"userHeadImgInfo"`
	IsFollowed           bool            `json:"isFollowed"`
	Name                 string          `json:"name"`
	Signature            string          `json:"signature"`
	ID                   string          `json:"id"`
}
type CoverCdnUrls struct {
	URL            string `json:"url"`
	FreeTrafficCdn bool   `json:"freeTrafficCdn"`
}
type CoverImgInfo struct {
	Width                int            `json:"width"`
	Height               int            `json:"height"`
	Size                 int            `json:"size"`
	Type                 int            `json:"type"`
	Animated             bool           `json:"animated"`
	ThumbnailImage       ThumbnailImage `json:"thumbnailImage"`
	ThumbnailImageCdnURL string         `json:"thumbnailImageCdnUrl"`
}
type Channel struct {
	ParentID   int    `json:"parentId"`
	ParentName string `json:"parentName"`
	Name       string `json:"name"`
	ID         int    `json:"id"`
}
type RankList struct {
	GroupID                     string         `json:"groupId"`
	UserID                      int            `json:"userId"`
	DougaID                     string         `json:"dougaId"`
	IsFollowing                 bool           `json:"isFollowing"`
	UserImg                     string         `json:"userImg"`
	ChannelID                   int            `json:"channelId"`
	ChannelName                 string         `json:"channelName"`
	ContentID                   int            `json:"contentId"`
	FansCount                   int            `json:"fansCount"`
	ContentTitle                string         `json:"contentTitle"`
	ContentDesc                 string         `json:"contentDesc,omitempty"`
	VideoCover                  string         `json:"videoCover"`
	UserSignature               string         `json:"userSignature"`
	ContributionCount           int            `json:"contributionCount"`
	DanmuCount                  int            `json:"danmuCount"`
	ContributeTime              int64          `json:"contributeTime"`
	UserName                    string         `json:"userName"`
	Duration                    int            `json:"duration"`
	ContentType                 int            `json:"contentType"`
	CreateTime                  string         `json:"createTime"`
	RecoReason                  RecoReason     `json:"recoReason"`
	CommentCountRealValue       int            `json:"commentCountRealValue"`
	DurationMillis              int            `json:"durationMillis"`
	ShareCount                  int            `json:"shareCount"`
	GiftPeachCount              int            `json:"giftPeachCount"`
	LikeCount                   int            `json:"likeCount"`
	ViewCount                   int            `json:"viewCount"`
	CommentCount                int            `json:"commentCount"`
	StowCount                   int            `json:"stowCount"`
	BananaCount                 int            `json:"bananaCount"`
	CoverURL                    string         `json:"coverUrl"`
	CreateTimeMillis            int64          `json:"createTimeMillis"`
	DanmakuCount                int            `json:"danmakuCount"`
	VideoList                   []VideoList    `json:"videoList"`
	TagList                     []TagList      `json:"tagList,omitempty"`
	Title                       string         `json:"title"`
	User                        User           `json:"user,omitempty"`
	HasHotComment               bool           `json:"hasHotComment"`
	CommentCountShow            string         `json:"commentCountShow"`
	LikeCountShow               string         `json:"likeCountShow"`
	IsLike                      bool           `json:"isLike"`
	CoverCdnUrls                []CoverCdnUrls `json:"coverCdnUrls"`
	CoverImgInfo                CoverImgInfo   `json:"coverImgInfo"`
	CommentCountTenThousandShow string         `json:"commentCountTenThousandShow"`
	ShareURL                    string         `json:"shareUrl"`
	PicShareURL                 string         `json:"picShareUrl"`
	IsDislike                   bool           `json:"isDislike"`
	ViewCountShow               string         `json:"viewCountShow"`
	ShareCountShow              string         `json:"shareCountShow"`
	StowCountShow               string         `json:"stowCountShow"`
	DanmakuCountShow            string         `json:"danmakuCountShow"`
	BananaCountShow             string         `json:"bananaCountShow"`
	GiftPeachCountShow          string         `json:"giftPeachCountShow"`
	DisableEdit                 bool           `json:"disableEdit"`
	BelongToSpecifyArubamu      bool           `json:"belongToSpecifyArubamu"`
	Description                 string         `json:"description"`
	Status                      int            `json:"status"`
	Channel                     Channel        `json:"channel"`
	IsFavorite                  bool           `json:"isFavorite"`
	SuperUbb                    bool           `json:"superUbb"`
	IsRewardSupportted          bool           `json:"isRewardSupportted"`
	IsThrowBanana               bool           `json:"isThrowBanana"`
	OriginalDeclare             int            `json:"originalDeclare,omitempty"`
	StaffContribute             bool           `json:"staffContribute,omitempty"`
}

type ArticleApi struct {
	Result    int               `json:"result"`
	HostName  string            `json:"host-name"`
	RankList  []ArticleRankList `json:"rankList"`
	RequestID string            `json:"requestId"`
}
type ArticleTagList struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}
type ArticleRankList struct {
	GroupID                     string           `json:"groupId"`
	SourcePlatform              string           `json:"sourcePlatform"`
	IsFollowing                 bool             `json:"isFollowing"`
	UserImg                     string           `json:"userImg"`
	ChannelID                   int              `json:"channelId"`
	ViewCount                   int              `json:"viewCount"`
	CommentCount                int              `json:"commentCount"`
	BananaCount                 int              `json:"bananaCount"`
	ResourceID                  int              `json:"resourceId"`
	ResourceType                string           `json:"resourceType"`
	ChannelName                 string           `json:"channelName"`
	ContentID                   int              `json:"contentId"`
	FansCount                   int              `json:"fansCount"`
	TagList                     []ArticleTagList `json:"tagList,omitempty"`
	AuthorID                    int              `json:"authorId"`
	UserID                      int              `json:"userId"`
	ContentTitle                string           `json:"contentTitle"`
	VideoCover                  string           `json:"videoCover"`
	UserSignature               string           `json:"userSignature"`
	ContributionCount           int              `json:"contributionCount"`
	DanmuCount                  int              `json:"danmuCount"`
	ContributeTime              int64            `json:"contributeTime"`
	UserName                    string           `json:"userName"`
	CoverImgInfo                CoverImgInfo     `json:"coverImgInfo"`
	CommentCountTenThousandShow string           `json:"commentCountTenThousandShow"`
	ShareURL                    string           `json:"shareUrl"`
	ViewCountShow               string           `json:"viewCountShow"`
	DanmakuCountShow            string           `json:"danmakuCountShow"`
	BananaCountShow             string           `json:"bananaCountShow"`
	ContentType                 int              `json:"contentType"`
	Status                      int              `json:"status"`
	UserHeadImgInfo             UserHeadImgInfo  `json:"userHeadImgInfo"`
	Channel                     Channel          `json:"channel"`
	ContentDesc                 string           `json:"contentDesc,omitempty"`
}
