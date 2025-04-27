package profiles

type Config map[string]interface{}
type Vars map[string]string

type Profile struct {
	Name string
	Vars Vars
}

type ByName []Profile

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
