package timeline

type Option struct {
	Label   string `json:"label,omitempty"`
	Content string `json:"content,omitempty"`
	Detail  string `json:"detail,omitempty"`
	Color   string `json:"color,omitempty"`
}

func Primary(label, content string) *Option {
	return &Option{Label: label, Content: content, Color: "primary"}
}
func Success(label, content string) *Option {
	return &Option{Label: label, Content: content, Color: "success"}
}
func Warning(label, content string) *Option {
	return &Option{Label: label, Content: content, Color: "warning"}
}
func Error(label, content string) *Option {
	return &Option{Label: label, Content: content, Color: "error"}
}
func (o *Option) WithDetail(detail string) *Option {
	o.Detail = detail
	return o
}
