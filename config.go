package main
import "os"

var C = os.Args
type GocmanConfig struct {
  Command string `json:"cmd"`
  Env []string  `json:"env"`
}
type Cmd map[string]GocmanConfig 
var Pkg []string
var Env []string
