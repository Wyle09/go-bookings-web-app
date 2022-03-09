package models

// Holds data sent from handlers to templates
type TemplateData struct {
	StringMap     map[string]string
	IntMap        map[string]int
	FLoatMap      map[string]float32
	Data          map[string]interface{}
	CSRFToken     string
	SucessMessage string
	Warning       string
	Error         string
}
