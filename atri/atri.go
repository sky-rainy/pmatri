/*
Package atri 本文件基于 https://github.com/Kyomotoi/ATRI
为 Golang 移植版，语料、素材均来自上述项目
本项目遵守 AGPL v3 协议进行开源
*/
package atri

import (
	"math/rand"
	"time"

	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/FloatTech/zbputils/process"
	"github.com/wdvxdr1123/ZeroBot/extension/rate"
)

var (
	// 戳一戳
	poke = rate.NewManager[int64](time.Minute*5, 8)
)

const (
	// 服务名
	servicename = "pm"
	// ATRI 表情的 codechina 镜像
	res = "https://gitcode.net/acfunghost/pm/-/raw/pmimg/"
)

func init() { // 插件主体
	engine := control.Register(servicename, &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Help:             "想了解派蒙更多吗？\n",
		OnEnable: func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			ctx.SendChain(message.Text("嗯呜呜……？"))
		},
		OnDisable: func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			ctx.SendChain(message.Text("Zzz……Zzz……"))
		},
	})
	// 被喊名字
	engine.OnFullMatch("", zero.OnlyToMe, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			var nickname = zero.BotConfig.NickName[0]
			process.SleepAbout1sTo2s()

			switch rand.Intn(2) {
			case 0:
				ctx.SendChain(message.Text(
					[]string{
						nickname + "在此，有何贵干？~",
						"什么？什么？",
						"在呢！在呢！",
						nickname + "不在呢~",
					}[rand.Intn(4)],
				))
			case 1:
				ctx.SendChain(message.Text("我在呢，怎么啦？"), randImage("pmzixinshuai.png", "pmzai.png"))
			}
		})
	// 戳一戳
	engine.On("notice/notify/poke", zero.OnlyToMe, isAtriSleeping).SetBlock(false).
		Handle(func(ctx *zero.Ctx) {
			var nickname = zero.BotConfig.NickName[0]
			switch {
			case poke.Load(ctx.Event.GroupID).AcquireN(3):
				process.SleepAbout1sTo2s()
				ctx.SendChain(message.Text("请不要戳", nickname, " ！不然我就生气了！"), randImage("pmshengqi.png", "pmfeijie.png"))
			case poke.Load(ctx.Event.GroupID).Acquire():
				process.SleepAbout1sTo2s()
				ctx.SendChain(message.Text("喂(#`O′) 戳", nickname, "干嘛！"), randImage("pmmiaoshi.png", "pmnadao.png"))
			default:
				// 频繁触发，不回复
			}
		})
	engine.OnKeywordGroup([]string{"草你妈", "操你妈", "脑瘫", "脑残", "废柴", "fw", "five", "废物", "战斗", "爬", "爪巴", "sb", "SB", "傻B", "垃圾", "傻逼", "傻屌", "2B", "2b", "傻叉"}, isAtriSleeping, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			ctx.SetGroupBan(
				ctx.Event.GroupID,
				ctx.Event.UserID, // 禁言的人的qq
				1*600,
			)
			ctx.SendChain(message.Text("大坏蛋！！闭嘴~"), randImage("pmshengqi3.png", "pmshengqi4.png", "pmshengqi2.png"))
		})
	engine.OnKeywordGroup([]string{"啊这"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			if rand.Intn(2) == 0 {
				ctx.SendChain(randImage("pmfeijie.png", "pmdanxin.png"))
			}
		})
	engine.OnFullMatchGroup([]string{"？", "?", "¿"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			switch rand.Intn(5) {
			case 0:
				ctx.SendChain(randText("?", "？", "嗯？", "(。´・ω・)ん?", "ん？"))
			case 1, 2:
				ctx.SendChain(randImage("pmwenhao.png", "pmwenhao2.png", "pmwenhao3.png", "pmwenhao4.png"))
			}
		})
	engine.OnKeyword("离谱", isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			switch rand.Intn(5) {
			case 0:
				ctx.SendChain(randText("?", "？", "嗯？", "(。´・ω・)ん?", "ん？"))
			case 1, 2:
				ctx.SendChain(randImage("pmfeijie2.png"))
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
