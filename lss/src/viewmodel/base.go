package	viewmodel

//Base provides a simple base struct for the data context.
type Base struct {
	Title string
}
//NewBase is the constructor for the data context.
func NewBase() Base {
	return Base{
		Title: "Lemonade Stand Supply",
	}
}