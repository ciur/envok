package profiles

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func loadConfig(filename string) (*Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func convertToStringMap(data interface{}) (Vars, error) {
	result := make(Vars)

	rawMap, ok := data.(Config)
	if !ok {
		return nil, fmt.Errorf("expected Config, got %T", data)
	}

	for key, value := range rawMap {
		switch v := value.(type) {
		case string:
			result[key] = v
		case int, float64, bool:
			result[key] = fmt.Sprintf("%v", v)
		default:
			return nil, fmt.Errorf("unsupported value type: %T for key %s", v, key)
		}
	}

	return result, nil
}

func Load(fileName string) ([]Profile, error) {
	var profiles []Profile

	config, err := loadConfig(fileName)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	for key, data := range *config {
		envVars, err := convertToStringMap(data)
		if err != nil {
			return nil, fmt.Errorf("Conversion error: %s\n", err)
		}

		profile := Profile{Name: key, Vars: envVars}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}
