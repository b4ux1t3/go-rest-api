package config

// Configuration contains the values for initializing the HTTP server.
// 	- Port is the TCP port on which teh server will listen
//	- IP is the IP address that the server will respond to
//	- Res_Tmpl is the HTML template file location for the browser-facing response
//	- Ind_Tmpl is the HTML template file location for the browser-facing root-level path
type Configuration struct {
	Port    string
	IP      string
	ResTmpl string
	IndTmpl string
}

// LoadConfig takes a file path as an optional parameter. If that file is passed as nil,
// a Configuration will be made which references the local directory for the template
// files and a default port of 8080, with nil for IP.
func LoadConfig(filePath string) Configuration {
	if filePath == "" {
		config := Configuration{Port: "8080", IP: "", ResTmpl: "./result.html", IndTmpl: "./index.html"}
		return config
	}

	config := Configuration{Port: "8080", IP: "", ResTmpl: "./result.html", IndTmpl: "./index.html"}
	return config

}
