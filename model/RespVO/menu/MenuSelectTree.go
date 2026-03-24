package menu

type MenuTreeSelect struct {
	ID       int64             `json:"id"`
	Label    string            `json:"label"`
	Disabled bool              `json:"disabled"`
	Children []*MenuTreeSelect `json:"children,omitempty"`
}
