package config

import (
	"html/template"
	"log"
)

// AppConfig przechowuje konfigurację aplikacji
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
