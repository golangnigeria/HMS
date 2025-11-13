package render

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"strings"
)

 
// ✅ Add custom functions here
var functions = template.FuncMap{
	// Creates a map (dictionary) from key-value pairs 
	"dict": func(values ...interface{}) (map[string]interface{}, error) {
		if len(values)%2 != 0 {
			return nil, fmt.Errorf("invalid dict call: odd number of args")
		}
		dict := make(map[string]interface{}, len(values)/2)
		for i := 0; i < len(values); i += 2 {
			key, ok := values[i].(string)
			if !ok {
				return nil, fmt.Errorf("dict keys must be strings")
			}
			dict[key] = values[i+1]
		}
		return dict, nil
	},

	// Creates a list (slice) from given arguments
	"list": func(args ...interface{}) []interface{} {
		return args
	},

	// Converts a value to its JSON string representation
	"toJson": func(v interface{}) string {
		b, err := json.Marshal(v)
		if err != nil {
			log.Println("Error in toJson:", err)
			return "{}"
		}
		return string(b)
	},

	// Splits a string into lines based on newline characters
	"splitLines": func(s string) []string {
		return strings.Split(s, "\n")
	},
}
