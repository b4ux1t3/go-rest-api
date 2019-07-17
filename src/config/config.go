package config

import (
	"encoding/json"
	"io/ioutil"
)

// Configuration contains the values for initializing the HTTP server.
// 	- Port is the TCP port on which teh server will listen
//	- IP is the IP address that the server will respond to
//	- ResTmpl is the HTML template file location for the browser-facing response
//	- IndTmpl is the HTML template file location for the browser-facing root-level path
type Configuration struct {
	IP      string
	Port    string
	IndTmpl string
	ResTmpl string
}

// LoadConfig takes a file path as an optional parameter. If that file is passed as nil,
// a Configuration will be made which references the local directory for the template
// files and a default port of 8080, with nil for IP.
func LoadConfig(filePath string) Configuration {
	var config Configuration
	if filePath == "" {
		config = Configuration{Port: "8080", IP: "", ResTmpl: "./result.html", IndTmpl: "./index.html"}
		return config
	}

	// First we need to give the decoder a place to work.
	// We could just use a string to string map, but that
	// would preclude us from using integers later on.
	var data map[string]interface{}

	// Then we need to grab the actual bytes from the file.
	fileBytes, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	// Then we unmarshal the data into our map.
	err = json.Unmarshal(fileBytes, &data)

	if err != nil {
		panic(err)
	}

	// Now we can break the data out into our config.
	config.IP = data["IP"].(string)
	config.Port = data["Port"].(string)
	config.IndTmpl = data["IndTmpl"].(string)
	config.ResTmpl = data["ResTmpl"].(string)

	return config

}
