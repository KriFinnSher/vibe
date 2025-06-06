// Code generated by ogen, DO NOT EDIT.

package api

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/Tintin
type Tintin struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Occupation OptString `json:"occupation"`
	Age        OptInt    `json:"age"`
}

// GetID returns the value of ID.
func (s *Tintin) GetID() string {
	return s.ID
}

// GetName returns the value of Name.
func (s *Tintin) GetName() string {
	return s.Name
}

// GetOccupation returns the value of Occupation.
func (s *Tintin) GetOccupation() OptString {
	return s.Occupation
}

// GetAge returns the value of Age.
func (s *Tintin) GetAge() OptInt {
	return s.Age
}

// SetID sets the value of ID.
func (s *Tintin) SetID(val string) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Tintin) SetName(val string) {
	s.Name = val
}

// SetOccupation sets the value of Occupation.
func (s *Tintin) SetOccupation(val OptString) {
	s.Occupation = val
}

// SetAge sets the value of Age.
func (s *Tintin) SetAge(val OptInt) {
	s.Age = val
}

func (*Tintin) tintinsIDGetRes() {}
func (*Tintin) tintinsIDPutRes() {}

// TintinsIDDeleteNoContent is response for TintinsIDDelete operation.
type TintinsIDDeleteNoContent struct{}

func (*TintinsIDDeleteNoContent) tintinsIDDeleteRes() {}

// TintinsIDDeleteNotFound is response for TintinsIDDelete operation.
type TintinsIDDeleteNotFound struct{}

func (*TintinsIDDeleteNotFound) tintinsIDDeleteRes() {}

// TintinsIDGetNotFound is response for TintinsIDGet operation.
type TintinsIDGetNotFound struct{}

func (*TintinsIDGetNotFound) tintinsIDGetRes() {}

// TintinsIDPutNotFound is response for TintinsIDPut operation.
type TintinsIDPutNotFound struct{}

func (*TintinsIDPutNotFound) tintinsIDPutRes() {}
