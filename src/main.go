package main

import (
		"fmt"
        "net/http"
        "runtime"
        "github.com/julienschmidt/httprouter"
        )
     
func main() {
  InitConfig()        

  cpuCount := runtime.NumCPU()
  runtime.GOMAXPROCS(cpuCount)
  

  router := httprouter.New()
  router.GET("/dice", GetDice)
  http.ListenAndServe(":"+ "1234", router)
}
