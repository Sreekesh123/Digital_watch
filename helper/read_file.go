package helper

import (
"io/ioutil"
"digital_watch/service"
)


func ReadHTMLFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(content), nil
}


func ShowCurrentMode() string {
	
	if service.CurrentMode == 0 {
		return "Letters"
	}
	return "Numbers"
	
}