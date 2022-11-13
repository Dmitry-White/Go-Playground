package core

// If you want to export a struct and
//  init this struct value outside of it's package
// than all fields of the struct must start with capital letter,
// otherwise you will get an error "Unexported field 'fieldName' usage"
type Flags struct {
	Name     string
	Greeting string
	Prompt   bool
	Preview  bool
}
