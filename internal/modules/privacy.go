/*
 * ● ArcMusic
 * ○ A high-performance engine for streaming music in Telegram voicechats.
 *
 * Copyright (C) 2026 Team Arc
 */

package modules

import (
	"fmt"

	tg "github.com/amarnathcjd/gogram/telegram"

	"main/internal/config"
)

func init() {
	helpTexts["/privacy"] = `<i>Show the bot's privacy policy.</i>`
}

func privacyHandler(m *tg.NewMessage) error {
	privacyText := fmt.Sprintf(`<b>🛡 ကိုယ်ရေးအချက်အလက်မူဝါဒ (Privacy Policy)</b>

သင်၏ ကိုယ်ရေးအချက်အလက်ကို ကျွန်ုပ်တို့ အလေးထားပါသည်။ ဤ Bot ကို ကိုယ်ရေးအချက်အလက် လုံခြုံရေးကို အဓိကထား၍ ရေးဆွဲထားပါသည်။

<b>📊 ကျွန်ုပ်တို့ စုဆောင်းသော အချက်အလက်များ</b>
<blockquote>Bot လုပ်ဆောင်ချက်အတွက် လိုအပ်သော အချက်အလက်များသာ သိမ်းဆည်းပါသည်:
• <b>User & Chat IDs:</b> Group များအား ခွဲခြားသတ်မှတ်ရန်နှင့် ဆက်တင်များ စီမံခန့်ခွဲရန်။
• <b>Preferences:</b> ဘာသာစကားနှင့် Bot ဆက်တင်များ။
• <b>Access Control:</b> သင်၏ Group အတွင်း ခွင့်ပြုထားသော အသုံးပြုသူစာရင်း။
• <b>RTMP Config:</b> RTMP streaming အသုံးပြုမှသာ သိမ်းဆည်းပါသည်။</blockquote>

<b>📩 မက်ဆေ့ခ်ျ လုံခြုံရေး</b>
<blockquote>• Bot သည် <code>/play</code> ကဲ့သို့သော Command များ (သို့) Bot ၏ Button များကို အသုံးပြုသည့်အခါမှသာ မက်ဆေ့ခ်ျများကို ဖတ်ရှုပါသည်။
• သင်၏ ပုဂ္ဂလိက ပြောဆိုမှုများ သို့မဟုတ် Group အတွင်းရှိ အခြား မက်ဆေ့ခ်ျများကို ဖတ်ရှုခြင်း၊ သိမ်းဆည်းခြင်း သို့မဟုတ် စောင့်ကြည့်ခြင်း <b>လုံးဝ မပြုလုပ်ပါ</b>။</blockquote>

<b>🌐 ပြင်ပ ဝန်ဆောင်မှုများ</b>
<blockquote>• သင်တောင်းဆိုသော သီချင်းများကို ရှာဖွေရန်နှင့် Stream လုပ်ရန်အတွက် <b>YouTube</b> နှင့် <b>Spotify</b> ကဲ့သို့သော ပြင်ပ ဝန်ဆောင်မှုများကိုသာ အသုံးပြုပါသည်။
• ရှာဖွေမှု အချက်အလက် (Search query) မှလွဲ၍ အခြား ပုဂ္ဂိုလ်ရေး အချက်အလက်များကို ၎င်းတို့ထံ ပေးပို့ခြင်း မရှိပါ။</blockquote>

<b>🤝 အချက်အလက် မျှဝေခြင်း</b>
<blockquote>• ကျွန်ုပ်တို့သည် သင်၏ အချက်အလက်များကို တတိယအဖွဲ့အစည်းသို့ ရောင်းချခြင်း၊ မျှဝေခြင်း သို့မဟုတ် လဲလှယ်ခြင်း <b>လုံးဝ မပြုလုပ်ပါ</b>။
• အချက်အလက်အားလုံးကို Bot ၏ တေးဂီတဝန်ဆောင်မှုများ တိုးတက်ကောင်းမွန်စေရန်အတွက်သာ အသုံးပြုပါသည်။</blockquote>

<b>✨ ကျွန်ုပ်တို့၏ ကတိကဝတ်</b>
ဤ Bot သည် အသုံးပြုသူများ၏ လုံခြုံရေးကို လေးစားလျက် အရည်အသွေးမြင့် တေးဂီတဝန်ဆောင်မှုများ ပေးအပ်ရန် ရည်ရွယ်ထားသော <a href="https://t.me/myanmarbot_music">Open-source ပရောဂျက်</a> ဖြစ်ပါသည်။ ကျွန်တော်ကတော့ @HANTHAR999 ဖြစ်ပါသည်။

<i>မေးမြန်းလိုသည်များ ရှိပါက ကျွန်ုပ်တို့၏ <a href="%s">Support Chat</a> တွင် ဝင်ရောက်မေးမြန်းနိုင်ပါသည်။</i>`, config.SupportChat)

	_, err := m.Reply(privacyText)
	return err
}
