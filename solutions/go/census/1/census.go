// Package census simulates a system used to collect census data.
package census
import "unicode/utf8"

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
    return &Resident{name, age, address}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
    return utf8.RuneCountInString(r.Name) > 1 && utf8.RuneCountInString(r.Address["street"]) > 1
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
    r.Name = ""
    r.Age = 0
    r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
    count := 0
	for _, e := range residents {
        if e.HasRequiredInfo() {
            count++
        }
    }
    return count
}
