package ding_talk

// @指定的人
type At struct {
	AtMobiles []string `json:"atMobiles"` // 被@人的手机号(在content里添加@人的手机号)
	IsAtAll   bool     `json:"isAtAll"`   // @所有人时：true，否则为：false
}

// text 消息
type TextMessage struct {
	MsgType string   `json:"msgtype"` // 消息类型，此时固定为：text
	Text    TextData `json:"text"`
	At      *At      `json:"at"`
}
type TextData struct {
	Content string `json:"content"` // 消息内容
}

// link 消息
type LinkMessage struct {
	MsgType string   `json:"msgtype"` // 消息类型，此时固定为：link
	Link    LinkData `json:"link"`
}
type LinkData struct {
	Title      string `json:"title"`      // 消息标题
	Text       string `json:"text"`       // 消息内容。如果太长只会部分展示
	PicUrl     string `json:"picUrl"`     // 点击消息跳转的URL
	MessageUrl string `json:"messageUrl"` // 图片URL
}

// markdown 消息
type MarkdownMessage struct {
	MsgType  string       `json:"msgtype"` // 消息类型，此时固定为：markdown
	Markdown MarkdownData `json:"markdown"`
	At       *At          `json:"at"`
}
type MarkdownData struct {
	Title string `json:"title"` // 消息标题
	Text  string `json:"text"`  // 消息内容。如果太长只会部分展示
}

// 整体跳转的 actionCard 消息
type SingleActionCardMessage struct {
	MsgType    string               `json:"msgtype"` // 消息类型，此时固定为：actionCard
	ActionCard SingleActionCardData `json:"actionCard"`
}
type SingleActionCardData struct {
	Title          string `json:"title"`          // 消息标题
	Text           string `json:"text"`           // 消息内容。如果太长只会部分展示
	SingleTitle    string `json:"singleTitle"`    // 单个按钮的方案。(设置此项和singleURL后btns无效)
	SingleURL      string `json:"singleUrl"`      // 点击singleTitle按钮触发的URL
	BtnOrientation string `json:"btnOrientation"` // 0-按钮竖直排列，1-按钮横向排列
	HideAvatar     string `json:"hideAvatar"`     // 0-正常发消息者头像，1-隐藏发消息者头像
}

// 独立跳转的 actionCard 消息
type ActionCardMessage struct {
	MsgType    string         `json:"msgtype"` // 消息类型，此时固定为：actionCard
	ActionCard ActionCardData `json:"actionCard"`
}
type ActionCardData struct {
	Title          string     `json:"title"`          // 消息标题
	Text           string     `json:"text"`           // 消息内容。如果太长只会部分展示
	Btns           []BtnsData `json:"btns"`           // 按钮的信息
	BtnOrientation string     `json:"btnOrientation"` // 0-按钮竖直排列，1-按钮横向排列
	HideAvatar     string     `json:"hideAvatar"`     // 0-正常发消息者头像，1-隐藏发消息者头像
}
type BtnsData struct {
	Title     string `json:"title"`     // 按钮方案
	ActionURL string `json:"actionUrl"` // 点击按钮触发的URL
}

// FeedCard 消息
type FeedCardMessage struct {
	MsgType  string      `json:"msgtype"` // 消息类型，此时固定为：feedCard
	FeedCard FeedCardMsg `json:"feedCard"`
}
type FeedCardMsg struct {
	Links []FeedCardData `json:"links"`
}
type FeedCardData struct {
	Title      string `json:"title"`      // 单条信息文本
	MessageURL string `json:"messageUrl"` // 点击单条信息到跳转链接
	PicURL     string `json:"picUrl"`     // 单条信息后面图片的URL
}
