package main

import (
	"github.com/gophermasters/bug-free-report/server"
)

func main(){
	bugsReportServer := server.NewServer()

	bugsReportServer.Run()
}