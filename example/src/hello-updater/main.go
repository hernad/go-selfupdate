package main

import (
	"os"
	"log"
	"github.com/sanbornm/go-selfupdate/selfupdate"
	"fmt"
	"os/exec"
)

// "os"
// "syscall"
// "time"

var version string
var variant string

var updater = &selfupdate.Updater{
    CurrentVersion: version,                   // Manually update the const, or set it using `go build -ldflags="-X main.VERSION=<newver>" -o hello-updater src/hello-updater/main.go`
    ApiURL:         "http://localhost:8080/",  // The server hosting `$CmdName/$GOOS-$ARCH.json` which contains the checksum for the binary
    BinURL:         "http://localhost:8080/",  // The server hosting the zip file containing the binary application which is a fallback for the patch method
    DiffURL:        "http://localhost:8080/",  // The server hosting the binary patch diff for incremental updates
    Dir:            "update/",                 // The directory created by the app when run which stores the cktime file
    CmdName:        "hello-updater",           // The app name which is appended to the ApiURL to look for an update
    ForceCheck: true,                          // For this example, always check for an update unless the version is "dev"
}

func main() {
	log.Printf("Hello world I am currently version Husremović %v %v", updater.CurrentVersion, variant )
	updater.BackgroundRun()
	log.Printf("Next run, I should be %v", updater.Info.Version)
	launch_f18( "3.1.99" )
}


func launch_f18( cVersion string ) {

    cCmd := "./F18_" + cVersion
    binary, lookErr := exec.LookPath( cCmd )
    if lookErr != nil {
        panic(lookErr)
    }
    fmt.Println(binary)

    //os.Remove("/tmp/stdin")
    //os.Remove("/tmp/stdout")
    //os.Remove("/tmp/stderr")

    //fstdin, err1 := os.Create("/tmp/stdin")
    //fstdout, err2 := os.Create("/tmp/stdout")
    //fstderr, err3 := os.Create("/tmp/stderr")
    //if err3 != nil {
    //    fmt.Println(err3)
    //    panic("WOW")
    //}

    //args := []string{}
		//args := []string{ "2>F18.err.log" }

		/*

    procAttr := syscall.ProcAttr{
        Dir:   "/tmp",
        Files: []uintptr{ uintptr(syscall.Stdin), uintptr(syscall.Stdout), fstderr.Fd()},
        Env:   []string{"VAR1=ABC123"},
        Sys: &syscall.SysProcAttr{
            Foreground: true,
        },
    }
		*/

/*
		env := os.Environ()

    //pid, err := syscall.ForkExec(binary, argv, &procAttr)
    //fmt.Println("Spawned proc", pid, err)

		execErr := syscall.Exec(binary, args, env)
		if execErr != nil {
        panic(execErr)
    }
*/

		//cmd := exec.Command( "sh", "-c", "./F18.sh" )

		cmd := exec.Command( "./F18_" + cVersion )

		cmd.Stdin = os.Stdin
 	  cmd.Stdout = os.Stdout

	        //f18In, _ := cmd.StdinPipe()
	        //f18Out, _ := cmd.StdoutPipe()
	        cmd.Start()
	        //fmt.Println( f18In )
	        //fmt.Println( f18Out )
	        cmd.Wait()

    //time.Sleep(time.Second * 100)


}
