package ui

func Message(level int, msg string) {
	RootPage.SendMessage("EID0", "message", map[string]interface{}{"level": level, "msg": msg})
}
func Notify(level int, title, text string) {
	RootPage.SendMessage("EID0", "notify", map[string]interface{}{"level": level, "title": title, "text": text})

}
