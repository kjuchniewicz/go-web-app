package models

// TemplateData przechowuje dane przesłane od obsługi do szablonów
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CDRFToken string
	Flash     string
	Warning   string
	Error     string
}
