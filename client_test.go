package ding_talk

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewClient(t *testing.T) {
	Convey("测试发送通知到叮叮机器人", t, func() {
		client := NewClient("https://oapi.dingtalk.com/robot/send?access_token=f1ccc82a6df21007962649774eea13fbe8be3b91df822872b0f4eb9529581cdf")
		// text request
		text := TextMessage{
			MsgType: Text,
			Text: TextData{
				Content: "我就是我, 是不一样的烟火",
			},
			At: &At{
				IsAtAll: true,
			},
		}
		// link request
		link := LinkMessage{
			MsgType: Link,
			Link: LinkData{
				Title:      "时代的火车向前开",
				Text:       "这个即将发布的新版本，创始人陈航（花名“无招”）称它为“红树林”。\n而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，这一次，为什么是“红树林",
				PicUrl:     "http://img.alicdn.com/imgextra/i4/3087155153/O1CN01Ixo0P21nw7TNo8pbH_!!3087155153.jpg",
				MessageUrl: "https://www.dingtalk.com",
			},
		}
		// markdown request
		markdown := MarkdownMessage{
			MsgType: Markdown,
			Markdown: MarkdownData{
				Title: "杭州天气",
				Text: "#### 杭州天气 @156xxxx8827\n" +
					"> 9度，西北风1级，空气良89，相对温度73%\n\n" +
					"> ![screenshot](https://gw.alipayobjects.com/zos/skylark-tools/public/files/84111bbeba74743d2771ed4f062d1f25.png)\n" +
					"> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n",
			},
			At: &At{
				IsAtAll: true,
			},
		}
		// 整体跳转ActionCard request
		singleActionCard := SingleActionCardMessage{
			MsgType: ActionCard,
			ActionCard: SingleActionCardData{
				Title:          "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
				Text:           "![screenshot](serverapi2/@lADOpwk3K80C0M0FoA) ### 乔布斯 20 年前想打造的苹果咖啡厅 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
				HideAvatar:     "0",
				BtnOrientation: "0",
				SingleTitle:    "阅读全文",
				SingleURL:      "https://www.dingtalk.com",
			},
		}
		// 独立跳转ActionCard request
		actionCard := ActionCardMessage{
			MsgType: ActionCard,
			ActionCard: ActionCardData{
				Title:          "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
				Text:           "![screenshot](serverapi2/@lADOpwk3K80C0M0FoA) ### 乔布斯 20 年前想打造的苹果咖啡厅 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
				HideAvatar:     "0",
				BtnOrientation: "0",
				Btns: []BtnsData{
					{
						Title:     "内容不错",
						ActionURL: "https://www.dingtalk.com",
					}, {
						Title:     "不感兴趣",
						ActionURL: "https://www.dingtalk.com",
					},
				},
			},
		}
		feedCard := FeedCardMessage{
			MsgType: FeedCard,
			FeedCard: FeedCardMsg{
				Links: []FeedCardData{
					{
						Title:      "时代的火车向前开",
						MessageURL: "https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI",
						PicURL:     "https://www.dingtalk.com",
					}, {
						Title:      "时代的火车向前开2",
						MessageURL: "https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI",
						PicURL:     "https://www.dingtalk.com",
					},
				},
			},
		}

		Convey("测试发送 text 通知到叮叮机器人", func() {
			response, err := client.Execute(text)
			if err != nil {
				panic(err)
			}
			So(0, ShouldEqual, response.ErrCode)
		})
		Convey("测试发送 link 通知到叮叮机器人", func() {
			response, err := client.Execute(link)
			if err != nil {
				panic(err)
			}
			So(0, ShouldEqual, response.ErrCode)
		})
		Convey("测试发送 markdown 通知到叮叮机器人", func() {
			response, err := client.Execute(markdown)
			if err != nil {
				panic(err)
			}
			So(0, ShouldEqual, response.ErrCode)
		})
		Convey("测试发送整体跳转的 ActionCard 通知到叮叮机器人", func() {
			response, err := client.Execute(singleActionCard)
			if err != nil {
				panic(err)
			}
			So(0, ShouldEqual, response.ErrCode)
		})
		Convey("测试发送独立跳转的 ActionCard 通知到叮叮机器人", func() {
			response, err := client.Execute(actionCard)
			if err != nil {
				panic(err)
			}
			So(0, ShouldEqual, response.ErrCode)
		})
		Convey("测试发送 FeedCard 通知到叮叮机器人", func() {
			response, err := client.Execute(feedCard)
			if err != nil {
				panic(err)
			}
			So(0, ShouldEqual, response.ErrCode)
		})
	})

}
