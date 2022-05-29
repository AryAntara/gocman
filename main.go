package main

func main(){
	if len(C) == 1 {
		log("error no command")
		return
	}
	if len(C) == 2 {
		Handle(C[1])
	} else if len(C) == 3 {
		Handle(C[1],C[2])
	}
}

