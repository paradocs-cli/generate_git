package gengit

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type GitScaffold struct {
	RootPath string

}
func InitRepo(){

}

func AddRefs(){

}

func CommitObjs(){

}

// CheckForGit checks for a .git directory to avoid reinitialization of an already existsing git repository
func CheckForGit() bool{
	var dirs []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		dirs = append(dirs, path)
		return nil
	})
	if err != nil {
		log.Fatalf("Unable to parse directory: %v", err)
	}

	for _, v := range dirs {
		dir, err  := os.ReadDir(v)
		if err != nil {
			log.Fatal(err)
		}
		for _, x := range dir {
			if strings.Compare(x.Name(), ".git") == 0 {
				return false
			}
		}

	}
	return true
}