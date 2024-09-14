package models

// Data??
// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // interface ermöglicht es mir alles zu erhalten was es gibt
	CSRFToken string                 // placeholder
	Flash     string                 // placeholder
	Warning   string                 // placeholder
	Error     string                 // placeholder
}
