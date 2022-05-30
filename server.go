package microprog

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// StartProgrammableServer ...
func StartProgrammableServer() (err error) {

	var ps *programmableServer
	ps, err = newProgrammableServer("../../programs/simple.yaml")
	if err != nil {
		return err
	}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        ps,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Server listening...")
	err = s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server shutdown successful")
	return nil
}

type (
	programmableServer struct {
		program *Program
	}
)

func (s programmableServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	printRequest(r)
	var (
		resCode int
		resBody string
	)

	for _, pRoute := range s.program.Routes {
		if pRoute.RequestMethod == r.Method && pRoute.RequestPath == r.URL.String() {
			// Request matches a programmed route
			resCode = pRoute.ResponseCode
			resBody = pRoute.ResponseBody
		}
	}

	if resCode <= 0 {
		rw.WriteHeader(404)
		return
	}

	if resCode != 200 {
		rw.WriteHeader(resCode)
	}
	if len(resBody) > 0 {
		rw.Write([]byte(resBody))
	}

	fmt.Printf("Response: %d %d\n", resCode, len(resBody))
}

func newProgrammableServer(programFilePath string) (*programmableServer, error) {
	program, err := serverProgramFromFile(programFilePath)
	if err != nil {
		return nil, err
	}

	return &programmableServer{
		program: program,
	}, nil
}

func printRequest(r *http.Request) {
	fmt.Printf("Request: %s %s\n", r.Method, r.URL)
}
