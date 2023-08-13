package models

import "github.com/KingKord/bookings/internal/forms"

// TemplateData holds data sent from handlers to template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // if we're not sure what data type is we use interfaces
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}

type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}
