package indexbox

import (
	"html/template"
	"strings"

	adminTemplate "github.com/GoAdminGroup/go-admin/template"
)

type IndexBox struct {
	*adminTemplate.BaseComponent

	Title      template.HTML
	Value      template.HTML
	Url        string
	Color      template.HTML
	IsSvg      bool
	IsHexColor bool
	Icon       template.HTML
	IconColor  template.HTML
}

func New() IndexBox {
	return IndexBox{
		BaseComponent: &adminTemplate.BaseComponent{
			Name:     "indexbox",
			HTMLData: List["indexbox"],
		},
	}
}

func (s IndexBox) SetIconColor(value template.HTML) IndexBox {
	s.IconColor = value
	return s
}

func (s IndexBox) SetTitle(value template.HTML) IndexBox {
	s.Title = value
	return s
}

func (s IndexBox) SetValue(value template.HTML) IndexBox {
	s.Value = value
	return s
}

func (s IndexBox) SetColor(value template.HTML) IndexBox {
	s.Color = value
	if strings.Contains(string(value), "#") {
		s.IsHexColor = true
	}
	return s
}

func (s IndexBox) SetIcon(value template.HTML) IndexBox {
	s.Icon = value
	if strings.Contains(string(value), "svg") {
		s.IsSvg = true
	}
	return s
}

func (s IndexBox) SetUrl(value string) IndexBox {
	s.Url = value
	return s
}

func (s IndexBox) GetContent() template.HTML { return s.GetContentWithData(s) }
