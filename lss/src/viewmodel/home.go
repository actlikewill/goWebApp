package viewmodel

// Home represents the data for the home page.
type Home struct {
	Title string
	Active string
}

// NewHome creates a new instance of the Home struct
func NewHome() Home {
	result := Home{
		Active: "home",
		Title: "Lemonade Stand Supply",
	}
	return result
}