package simple

import (
	"html/template"

	adminTemplate "github.com/GoAdminGroup/go-admin/template"
)

type Simple struct {
	*adminTemplate.BaseComponent

	Border string
	Number template.HTML
	Title  template.HTML
	Color  template.HTML
	Url    string
}

func New() Simple {
	return Simple{
		BaseComponent: &adminTemplate.BaseComponent{
			Name:     "simple",
			HTMLData: List["simple"],
		},
	}
}

func (c Simple) SetUrl(value string) Simple {
	c.Url = value
	return c
}

func (c Simple) SetTitle(value template.HTML) Simple {
	c.Title = value
	return c
}

func (c Simple) SetNumber(value template.HTML) Simple {
	c.Number = value
	return c
}

func (c Simple) SetColor(value template.HTML) Simple {
	c.Color = value
	return c
}

func (c Simple) SetBorder(value string) Simple {
	c.Border = value
	return c
}

func (c Simple) GetContent() template.HTML { return c.GetContentWithData(c) }
