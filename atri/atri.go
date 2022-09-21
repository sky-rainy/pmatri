/*
Package atri 本文件基于 https://github.com/Kyomotoi/ATRI
为 Golang 移植版，语料、素材均来自上述项目
本项目遵守 AGPL v3 协议进行开源
*/
package atri

import (
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"regexp"
	"strconv"
	"time"

	"github.com/FloatTech/AnimeAPI/aireply"
	"github.com/FloatTech/zbputils/process"
	"github.com/pkumza/numcn"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

const (
	// 服务名
	servicename = "pm"
	// ATRI 表情的 codechina 镜像
	res = "https://gitcode.net/acfunghost/pm/-/raw/pmimg/"
)

var (
	re    = regexp.MustCompile(`(\-|\+)?\d+(\.\d+)?`)
	cnapi = "http://233366.proxy.nscc-gz.cn:8888?speaker=%s&text=%s"
)

func init() {
	// 被喊名字
	zero.OnRegex(`^派蒙(.*)$`, zero.OnlyToMe, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			var nickname = zero.BotConfig.NickName[0]
			msg := ctx.State["regex_matched"].([]string)[1]
			process.SleepAbout1sTo2s()
			// 获取回复模式
			r := aireply.NewAIReply("小爱")
			// 获取回复的文本
			reply := r.TalkPlain(msg, zero.BotConfig.NickName[0])

			// 获取语言
			record := message.Record(fmt.Sprintf(cnapi, url.QueryEscape(nickname), url.QueryEscape(
				// 将数字转文字
				re.ReplaceAllStringFunc(reply, func(s string) string {
					f, err := strconv.ParseFloat(s, 64)
					if err != nil {
						log.Println("获取语音err : ", err)
						return s
					}
					return numcn.EncodeFromFloat64(f)
				}),
			))).Add("cache", 0)
			if record.Data == nil {
				ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text(reply))
				return
			}
			// 发送语音
			ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text(reply))
			// switch rand.Intn(2) {
			// case 0:

			// case 1:
			// 	ctx.SendChain(message.Text("我在呢，怎么啦？"), randImage("pmzixinshuai.png", "pmzai.png"))
			// }
		})
	// 戳一戳
	zero.On("notice/notify/poke", zero.OnlyToMe, isAtriSleeping).SetBlock(false).
		Handle(func(ctx *zero.Ctx) {
			var nickname = zero.BotConfig.NickName[0]
			process.SleepAbout1sTo2s()
			switch rand.Intn(3) {
			case 0:
				ctx.SendChain(message.Text("请不要戳", nickname, " ！不然我就生气了！"), randImage("pmshengqi.png", "pmfeijie.png"))
			case 1:
				process.SleepAbout1sTo2s()
				ctx.SendChain(message.Text("喂(#`O′) 戳", nickname, "干嘛！"), randImage("pmmiaoshi.png", "pmnadao.png"))
			default:
				// 频繁触发，不回复
			}
		})
	zero.OnKeywordGroup([]string{"草你妈", "操你妈", "脑瘫", "脑残", "废柴", "fw", "five", "废物", "战斗", "爬", "爪巴", "sb", "SB", "傻B", "垃圾", "傻逼", "傻屌", "2B", "2b", "傻叉"}, isAtriSleeping, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			ctx.SetGroupBan(
				ctx.Event.GroupID,
				ctx.Event.UserID, // 禁言的人的qq
				1*600,
			)
			ctx.SendChain(message.Text("大坏蛋！！闭嘴~"), randImage("pmshengqi3.png", "pmshengqi4.png", "pmshengqi2.png"))
		})
	zero.OnKeywordGroup([]string{"谢谢", "3q", "非常感谢"}, zero.OnlyToMe, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			switch rand.Intn(3) {
			case 0:
				ctx.SendChain(message.Text("不客气"), randImage("pmwunai.png", "pmweixiao.png"))
			case 1:
				process.SleepAbout1sTo2s()
				ctx.SendChain(message.Text("客气了，下次请我好吃的就行"), randImage("pmwunai.png", "pmweixiao.png"))
			default:
				// 频繁触发，不回复
			}
		})
	zero.OnKeywordGroup([]string{"啊这", "阿这"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			if rand.Intn(2) == 0 {
				ctx.SendChain(randImage("pmfeijie.png", "pmdanxin.png"))
			}
		})
	zero.OnKeywordGroup([]string{"你好可爱", "可爱", "好可爱啊"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			ctx.SendChain(randText("谢谢", "哎呀不要老夸我，我会害羞的~", "哈哈~谢谢"), randImage("pmsaobaoqing.png", "pmmiaoshi3.png"))
		})
	zero.OnKeywordGroup([]string{"大佬", "厉害", "666", "感谢大佬"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			if rand.Intn(2) == 0 {
				ctx.SendChain(randImage("pmdalao2.png", "pmdalao.png"))
			}
		})
	zero.OnKeywordGroup([]string{"是不是", "你说呢"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			switch rand.Intn(5) {
			case 0:
				ctx.SendChain(randText("对对", "嗯嗯", "可能是吧"), randImage("pmmiaoshi3.png", "pmmiaoshi.png"))
			case 1, 2:
				ctx.SendChain(randImage("pmwenhao.png", "pmwenhao2.png", "pmwenhao3.png", "pmwenhao4.png"))
			}
		})
	zero.OnFullMatchGroup([]string{"？", "?", "¿"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			switch rand.Intn(5) {
			case 0:
				ctx.SendChain(randText("?", "？", "嗯？", "(。´・ω・)ん?", "ん？"))
			case 1, 2:
				ctx.SendChain(randImage("pmwenhao.png", "pmwenhao2.png", "pmwenhao3.png", "pmwenhao4.png"))
			}
		})
	zero.OnKeyword("离谱", isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			switch rand.Intn(5) {
			case 0:
				ctx.SendChain(randText("?", "？", "嗯？", "(。´・ω・)?", "？？？"))
			case 1, 2:
				ctx.SendChain(randImage("pmfeijie2.png", "pmasir.png"))
			}
		})
}

func randText(text ...string) message.MessageSegment {
	return message.Text(text[rand.Intn(len(text))])
}

func randImage(file ...string) message.MessageSegment {
	return message.Image(res + file[rand.Intn(len(file))])
}

func randRecord(file ...string) message.MessageSegment {
	return message.Record(res + file[rand.Intn(len(file))])
}

// isAtriSleeping 凌晨0点到6点，ATRI 在睡觉，不回应任何请求
func isAtriSleeping(ctx *zero.Ctx) bool {
	if now := time.Now().Hour(); now >= 1 && now < 6 {
		return false
	}
	return true
}
