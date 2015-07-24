package main

import (
        "net/http"
        "runtime"
        "github.com/juleinschmit/httprouter"
        )
        
func main() {
  
  cpuCount := runtime.NumCPU()
  runtime.GOMAXPROCS(cpuCount)
  
  router := httprouter.New()
  
  http.ListenAndServe(":"+ config.Serve.Port, router)

}
