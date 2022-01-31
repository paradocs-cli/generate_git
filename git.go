package gengit

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	repo *git.Repository
)

// InitRepo takes a string argument for repo path and then checks to see if it has already been initialized
// if it has it returns the Repository, if not it initializes it and returns the repository
func InitRepo(w string)(*git.Repository,error){
	check , err := CheckForGit()
	if err != nil {
		return nil, fmt.Errorf("%v", err.Error())
	}
	if check {
		repo, err = git.PlainOpen(w)
		if err != nil {
			return repo, fmt.Errorf("%v", err.Error())
		}
		return repo, nil
	} else {
		repo, err = git.PlainInit(w, false)
		if err != nil {
			return repo, fmt.Errorf("%v", err.Error())

		}
	}
	return repo, nil
}

func AddRefs(){

}

func CommitObjs(){

}

// CheckForGit checks for a .git directory to avoid reinitialization of an already existsing git repository
func CheckForGit() (bool, error){
	var dirs []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("error checking filepath.Walk: %v\n",err)
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
				return false, nil
			}
		}

	}
	return true, nil
}