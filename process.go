package main
import (
	"os"
	"fmt"
	"encoding/json"
	"strings"
	"os/exec"
)
func ReadConfiguration(c string) {
	ConfigName := "gocman.config"
	ReadFile,err := os.ReadFile(ConfigName)
	if err != nil {
		log(err)
		return
	}
	Config := Cmd{}
	json.Unmarshal(ReadFile,&Config)
	if Config[c].Command == "" {
		return
	}
	if len(Config[c].Env) != 0 {
		Env = Config[c].Env
		//log(Env)
	}
	RunEnv()
	Command := strings.Split(Config[c].Command,">")
	for _,v := range Command {
		Pas := strings.Split(v," ")
		allArgs := Pas[1:]
		Run := exec.Command(Pas[0],allArgs...)
		log("run command "+v)
		err = Run.Run()
		out := ""
		if string(out) != "" {
			fmt.Println(strings.ReplaceAll(string(out),"\n",""))
		}
		//err = Run.Run()
		if err != nil {
			log(err)
		}
	}
}
func Init(){
	if _,err := os.Stat("./gocman.config"); err == nil {
		return
	}
	config := Cmd{}
	config["run"] = GocmanConfig{
		Command : "go run .",
	}
	config["build"] = GocmanConfig{
		Command : "go build",
	}
	//log(config)
	ex,_ := json.Marshal(config)
	os.WriteFile("gocman.config",ex,0700)
	return
}
func add(c string,v string,env []string){
	ConfigName := "gocman.config"
	ReadFile,err := os.ReadFile(ConfigName)
	if err != nil {
		os.WriteFile("gocman.config",[]byte(""),0700)
	}
	Config := Cmd{}
	json.Unmarshal(ReadFile,&Config)
	Config[c] = GocmanConfig{
		Command : strings.ReplaceAll(v,"+"," "),
		Env : env,
	}
	C,_ := json.Marshal(Config)
	os.WriteFile("gocman.config",C,0700)
	log("command added")
}
func log(c interface{}) {
	fmt.Println("::",c)
}
func Scan(c string) string {
	fmt.Print(c)
	name := ""
	fmt.Scan(&name)
	return name
}
func RunEnv(){
	configEnv := Env
	if len(configEnv) == 0 {
		log("no env to use")
		return 
	}
	splitEnvConf := configEnv
	for _,v := range splitEnvConf {
		splitEnv := strings.Split(v,"=")
		name := splitEnv[0]
		value := splitEnv[1]
		os.Setenv(name,value)
		log("env set `"+name+"` with value `"+value+"`")
	}
}
func install(c ...string){
	if len(c) == 0 {
		Config ,err:= os.ReadFile("gocman.package")
		if err != nil {
			log("file gocman.package not found")
			return 
		}
		mode := []string{}
		json.Unmarshal(Config,&mode)
		for _,v := range mode {
			install(v)
		}
		return 
	}
	cmd := exec.Command("go","get",c[0])
	err := cmd.Run()
	if err != nil {
		log("connection error")
		return
	}
	Pkg = append(Pkg,c[0])
}
func WriteFile(){
	if len(Pkg) == 0 {
		log("package empty")
		return
	}
	str,err := json.Marshal(Pkg)
	if err != nil {
		log("something err")
		return
	}
	os.WriteFile("gocman.package",str,0700)
}

