package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := map[uint16]string{
		104: "Ein super cooles Modul",
		117: "Netzwerkmodul",
		346: "Cloudmodul",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 104)
	modules[164] = "SQL-Modul"
	modules[117] = "Fortgeschrittenes Netzwerkmodul"
	// TODO: delete one
	// TODO: add one
	// TODO: replace one
	fmt.Println(modules)
}
