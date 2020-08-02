package cameo

import (
	"fmt"
	"github.com/gofiber/template/html"
	"html/template"
	"io"
)

type Engine html.Engine

func CreateTemplateEngine() *Engine {
	return &Engine{
		Templates: template.Must(template.New("form").Parse(FormTemplate)),
	}
}

func (e *Engine) Load() error {
	return nil
}

func (e *Engine) Render(out io.Writer, template string, binding interface{}, layout ...string) error {
	tmpl := e.Templates.Lookup(template)
	if tmpl == nil {
		return fmt.Errorf("render: template %s does not exist", template)
	}
	return tmpl.Execute(out, binding)
}
