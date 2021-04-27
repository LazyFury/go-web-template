package layout

import (
	"bytes"
	"html/template"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-template/response"
)

type (
	Layout struct {
		BaseTemplate string
		Tmpl         *template.Template
	}

	LayoutParams struct {
		Header       map[string]interface{}
		Data         map[string]interface{}
		Footer       map[string]interface{}
		TemplateName string
		Layout       *Layout
	}
)

func New(name string, tmpl *template.Template) *Layout {
	return &Layout{
		BaseTemplate: name,
		Tmpl:         tmpl,
	}
}

func (p *LayoutParams) bind(c *gin.Context) {
	// header
	p.Header["path"] = c.Request.URL.String()
	// footer
}

var m = &sync.RWMutex{}

func (p *LayoutParams) RenderPage() template.HTML {
	m.Lock()
	defer m.Unlock()
	var w = bytes.NewBuffer([]byte(""))

	if err := p.Layout.Tmpl.ExecuteTemplate(w, p.TemplateName, p.Data); err != nil {
		response.Error(err)
	}

	html := template.HTML(w.String())
	return html
}

func (l *Layout) Render(c *gin.Context, params *LayoutParams) {
	params.Layout = l
	params.bind(c)
	c.HTML(http.StatusOK, l.BaseTemplate, params)
}
