package profiles

type Config map[string]interface{}
type Vars map[string]string

type Profile struct {
	Name string
	Vars Vars
}
