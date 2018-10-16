package helpers

import (
	"github.com/moisespsena/go-error-wrap"
	"gopkg.in/yaml.v2"
)

func ParseMap(data map[string]interface{}, out interface{}) (err error) {
	b, err := yaml.Marshal(data)
	if err != nil {
		return errwrap.Wrap(err, "YAML Marshall data")
	}
	if err = yaml.Unmarshal(b, out); err != nil {
		return errwrap.Wrap(err, "YAML Umarshall config")
	}
	return nil
}
