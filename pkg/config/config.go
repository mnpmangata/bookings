package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

//avoid import cycle bug. do not import anything unless necessary

// Appconfig holds the application config
type Appconfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
