package microprog

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Program struct {
		Routes []Route `yaml:"routes"`
	}

	Route struct {
		RequestMethod string `yaml:"method"`
		RequestPath   string `yaml:"path"`
		ResponseCode  int    `yaml:"response_code"`
		ResponseBody  string `yaml:"response_body"`
	}
)

func serverProgramFromFile(programFilePath string) (p *Program, err error) {
	data, err := os.ReadFile(programFilePath)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, &p)
	if err != nil {
		return
	}

	fmt.Printf("Server program read (%d routes)\n", len(p.Routes))
	for i, r := range p.Routes {
		fmt.Printf("%d) %s %s %d %d\n", i, r.RequestMethod, r.RequestPath, r.ResponseCode, len(r.ResponseBody))
	}

	// TODO validate

	return p, nil
}
