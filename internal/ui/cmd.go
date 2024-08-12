package ui

func Message(level int, msg string) {
	RootElement.SendRootMessage("message", map[string]interface{}{"level": level, "msg": msg})
}
func Notify(level int, title, text string) {
	RootElement.SendRootMessage("notify", map[string]interface{}{"level": level, "title": title, "text": text})

}
