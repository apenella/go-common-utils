package data

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

//
// LoadJSONFile method dumps data from the file f to object
func LoadJSONFile(f string, object interface{}) error {
	file, err := ioutil.ReadFile(f)

	if err != nil {
		return errors.New("(LoadJSONFile) Error on loading file '" + f + "'. " + err.Error())
	}
	err = json.Unmarshal(file, object)
	if err != nil {
		return errors.New("(LoadJSONFile) Error on " + f + " unmarshaling. " + err.Error())
	}

	return nil
}

// ObjectToJSONString converts any object to a json string
func ObjectToJSONString(object interface{}) (string, error) {
	var jsoned []byte
	var err error

	jsoned, err = json.Marshal(object)
	if err != nil {
		return err.Error(), err
	}

	return string(jsoned), nil
}

// ObjectToJSONStringPretty converts any object to a json string with a pretty format
func ObjectToJSONStringPretty(object interface{}) (string, error) {
	var jsoned []byte
	var err error

	jsoned, err = json.MarshalIndent(object, "", "  ")
	if err != nil {
		return err.Error(), err
	}

	return string(jsoned), nil

}

//
// ObjectToJSONByte converts any object to a json byte
func ObjectToJSONByte(object interface{}) ([]byte, error) {
	return json.Marshal(object)
}
