package main

import  (
        "code.google.com/p/gcfg"
        )
        

type Config struct {
  Serve struct {
    Port string
  }
}

var (
  config Config
)

func InitConfig() {
  gcfg.ReadFileInto(&config, "path/to/config.cfg")
}
