// config for the app
package config

import "html/template"

// INFO
// UseCache wird ATM nicht verwendet

// AppConfig holds the application config
type AppConfig struct {
	UseCache              bool
	TemplateCache         []string
	TemplateCacheAdvanced map[string]*template.Template
}
