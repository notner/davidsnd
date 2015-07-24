package main

import (
        "net/http"
        "runtime"
        "github.com/juleinschmit/httprouter"
        )
     
func init() {
  initConfig()        

  cpuCount := runtime.NumCPU()
  runtime.GOMAXPROCS(cpuCount)
  
}
   
func main() {
  router := httprouter.New()
  http.ListenAndServe(":"+ config.Serve.Port, router)

}
