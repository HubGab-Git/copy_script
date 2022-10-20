package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const (
	defaultDir          = "."
	sourceDirUsage      = "Source directory from where files should be copied"
	destinationDirUsage = "Destination directory where files should be copied"
)

var sourceDir string
var destinationDir string

func init() {

	flag.StringVar(&sourceDir, "source-dir", defaultDir, sourceDirUsage)
	flag.StringVar(&sourceDir, "s", defaultDir, sourceDirUsage)
	flag.StringVar(&destinationDir, "destination-dir", defaultDir, destinationDirUsage)
	flag.StringVar(&destinationDir, "d", defaultDir, destinationDirUsage)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Copy files form one folder to another\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		// fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "   -s, --source-dir:\n")
		fmt.Fprintf(os.Stderr, "      "+sourceDirUsage+"\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "   -d, --destination-dir:\n")
		fmt.Fprintf(os.Stderr, "      "+destinationDirUsage+"\n")
		fmt.Fprintf(os.Stderr, "\n")
	}
}

func main() {
	flag.Parse()
	if sourceDir == "." && destinationDir == "." {
		flag.Usage()
	} else {
		var cmd *exec.Cmd
		if runtime.GOOS != "windows" {
			cmd = exec.Command("cp", "-R", sourceDir, destinationDir)
		} else {
			cmd = exec.Command("cmd", "/C", "xcopy", "/s", "/e", sourceDir, destinationDir)
		}
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		} else {
			fmt.Println("Copied files from '" + sourceDir + "' to '" + destinationDir + "'")
		}
	}
}
