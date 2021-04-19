package layout

import (
	"bytes"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-template/response"
	tools "github.com/lazyfury/go-web-template/tools/template"
)

type (
	LayoutParams struct {
		Header             map[string]interface{}
		Data               map[string]interface{}
		Footer             map[string]interface{}
		TemplateName       string
		HeaderTemplateName string
		FooterTemplateName string
		Template           *template.Template
	}
)

var Bootstrap *template.Template

func InitBootstrap(dir string, parent string, funcs template.FuncMap) {
	Bootstrap = template.Must(tools.ParseGlob(template.New("main").Funcs(funcs), dir, parent))
}

func New(name string, args ...map[string]interface{}) *LayoutParams {
	data := map[string]interface{}{}

	for _, arg := range args {
		for k := range arg {
			data[k] = arg[k]
		}
	}

	return &LayoutParams{
		Data:         data,
		TemplateName: name,
		Template:     Bootstrap,
		Header:       map[string]interface{}{},
		Footer:       map[string]interface{}{},
	}
}
func Render(c *gin.Context, name string, args ...map[string]interface{}) {
	layout := New(name, args...)
	layout.Render(c)
}

func (p *LayoutParams) bind(c *gin.Context) {
	// header
	p.Header["path"] = c.Request.URL.String()

	// footer
}
func (p *LayoutParams) Render(c *gin.Context) {

	w := bytes.NewBuffer([]byte(""))
	if p.HeaderTemplateName == "" {
		p.HeaderTemplateName = "header"
	}
	if p.FooterTemplateName == "" {
		p.FooterTemplateName = "footer"
	}

	p.bind(c)

	tmpl := p.Template.New("layout")
	err := tmpl.ExecuteTemplate(w, p.HeaderTemplateName, p.Header)
	if err != nil {
		response.Error(err)
	}
	err = tmpl.ExecuteTemplate(w, p.TemplateName, p.Data)
	if err != nil {
		response.Error(err)
	}
	err = tmpl.ExecuteTemplate(w, p.FooterTemplateName, p.Footer)
	if err != nil {
		response.Error(err)
	}
	html := template.HTML(w.String())
	_, err = c.Writer.Write([]byte(html))
	if err != nil {
		response.Error(err)
	}
}
