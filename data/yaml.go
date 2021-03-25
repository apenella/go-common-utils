package data

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func LoadYAMLFile(file string, object interface{}) error {

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return errors.New("(LoadYAMLFile) Error loading file " + file + ". " + err.Error())
	}
	err = yaml.Unmarshal(yamlFile, object)
	if err != nil {
		return errors.New("(LoadYAMLFile) Error on " + file + " unmarshaling. " + err.Error())
	}

	return nil
}

// ObjectToYamlString converts any object to a yaml string
func ObjectToYamlString(object interface{}) (string, error) {
	var yamled []byte
	var err error

	yamled, err = yaml.Marshal(object)
	if err != nil {
		return err.Error(), err
	}

	return string(yamled), nil
}
