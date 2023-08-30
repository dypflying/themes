package warninglist

import (
	"html/template"

	adminTemplate "github.com/GoAdminGroup/go-admin/template"
)

type WarningList struct {
	*adminTemplate.BaseComponent

	Data []map[string]string
}

func New() WarningList {
	return WarningList{
		BaseComponent: &adminTemplate.BaseComponent{
			Name:     "warninglist",
			HTMLData: List["warninglist"],
		},
	}
}

func (p WarningList) SetData(value []map[string]string) WarningList {
	p.Data = value
	return p
}

func (p WarningList) GetContent() template.HTML { return p.GetContentWithData(p) }
