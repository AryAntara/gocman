package main
import "fmt"

func Handle(i string,opt ...string){
	//log(i)
	switch i {
    case "init":
		log("initilize")
		Init()
		return
	case "run":
		//log(opt)
		if len(opt) == 0 {
			log("no command to run")
			return
		}
		command := opt[0]
		ReadConfiguration(command)
		break
	case "add":
		Name := Scan("Command Name : ")
		Command := Scan("Command use + as space and > before next command: ")
		prompt := Scan("Use env ?")
		i := false
		if prompt == "y" {
			i = true
		}
		env := []string{}
		for i {
				Name := Scan("ENV Name : ")
				Value := Scan("ENV Value : ")
				if Name == "" || Value == "" {
					break
				}
				structure := Name+"="+Value
				env = append(env,structure)
				prompt = Scan("add more?(y/n)")
				if prompt != "y" {
					break
				} else {
					break
				}
		}
		add(Name,Command,env)
	case "i","install":
		if len(opt) == 0 {
			install()
			WriteFile()
		}else {
			install(opt[0])
			WriteFile()
		}
	case "help","h":
		command := []string{"run <arg> -- run a command, command name as argument","add -- add a command","i/install -- install a module","h/help -- display this message"}
		log("usage command gocman <command> <arg> :")
		log("command :")
		for _,v := range command {
			fmt.Println("\t",v)
		}
	default:
		log("exit")
	}
}
