package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	CSRFToken string
	Error     string
	Flash     string
	Warning   string
	Data      map[string]interface{}
	FloatMap  map[string]float32
	IntMap    map[string]int
	StringMap map[string]string
}
