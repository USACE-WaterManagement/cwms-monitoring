package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"text/template"
)

//go:embed prometheus.yml.template
var ProfileTemplate string

func main() {
	var toRun = os.Args[1]
	var err error
	// build template
	tmpl, err := template.New("config").Parse(ProfileTemplate)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("/tmp/prometheus.yml", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	env := os.Environ()

	toMap := make(map[string]string)

	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if (len(pair) == 2) {
			toMap[pair[0]] = pair[1]
		}
	}

	err = tmpl.Execute(file, toMap)
	if err != nil {
		panic(err)
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}


	binaryPath, err := exec.LookPath(toRun)
	if err != nil {
		panic(err)
	}	

	
	args := []string{toRun,
			 "--config.file=/tmp/prometheus.yml",
			 "--storage.tsdb.path=/prometheus",
             "--web.console.libraries=/etc/prometheus/console_libraries",
             "--web.console.templates=/etc/prometheus/consoles"}	
	err = syscall.Exec(binaryPath, args, env)

	if err != nil {
		panic(err)
	}
}