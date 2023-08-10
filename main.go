package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	DEFAULT_DIR_NAME  string = "new_project"
	GO_MOD_INIT       string = "root"
	DEFAULT_GO_MAIN   string = "main.go"
	CONTENT_MAIN_FILE string = "package main \n\nfunc main(\n\n)"
)

func ErrorMsgParser(msg []string) string {
	if len(msg) < 1 {
		return ""
	} else {
		return msg[0]
	}
}

func ErrorHandler(err error, msg ...string) {
	errMsg := ErrorMsgParser(msg)

	if err != nil {
		log.Println(errMsg, err)
		panic(err)
	}
}

func main() {
	var err error

	currentDir, err := os.Getwd()
	ErrorHandler(err, "Greška prilikom dobijanja radnog direktorijuma:")

	dirName := ArgParser()

	err = os.Mkdir(dirName, 0777)
	ErrorHandler(err, "Problem pri stvaranju novog foldera:")
	fmt.Printf("Stvoren novi projekat (%s) na mjestu %s\n", dirName, currentDir)

	newDir := filepath.Join(currentDir, dirName)

	err = os.Chdir(newDir)
	ErrorHandler(err, "Greška prilikom promene radnog direktorijuma:")

	err = CreateGoMod()
	ErrorHandler(err, "Greška prilikom izvršavanja komande:")

	file, err := CreateMain()
	ErrorHandler(err, "Greška pri kreiranju main.go fajla")

	defer file.Close()

	_, err = file.WriteString(CONTENT_MAIN_FILE)
	ErrorHandler(err, "Greška prilikom pisanja u fajl:")
}

func CreateMain() (*os.File, error) {
	return os.Create(DEFAULT_GO_MAIN)
}

func CreateGoMod() error {
	cmd := exec.Command("go", "mod", "init", GO_MOD_INIT)

	return cmd.Run()
}

func ArgParser() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}

	return DEFAULT_DIR_NAME
}
