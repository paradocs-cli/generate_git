package gengit

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type GitOptions struct {
	CommitOptions struct {
		Name    string
		Message string
		Email   string
	}
	RemoteOptions struct {
		Pat string
		Provider string
		Repo string
	}
	Branches string
}

var (
	repo *git.Repository
)

// InitRepo takes a string argument for repo path and then checks to see if it has already been initialized
// if it has it returns the Repository, if not it initializes it and returns the repository
func InitRepo(w string) (*git.Repository, error) {
	check, err := CheckForGit()
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

// AddRefs adds the changes made to the staging area to be committed
func AddRefs(r git.Repository) (*git.Worktree, error) {
	tree, err := r.Worktree()
	if err != nil {
		return tree, fmt.Errorf("%v", err.Error())
	}
	_, err = tree.Add(".")
	return tree, nil
}

// CommitObjs commits updates for the current repository
func CommitObjs(w git.Worktree, o GitOptions) (string, error) {
	commit, err := w.Commit(o.CommitOptions.Message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  o.CommitOptions.Name,
			Email: o.CommitOptions.Email,
			When:  time.Now(),
		},
	})
	if err != nil {
		return commit.String(), fmt.Errorf("%v", err.Error())
	}
	return commit.String(), nil
}

// CheckForGit checks for a .git directory to avoid reinitialization of an already existsing git repository
func CheckForGit() (bool, error) {
	var dirs []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("error checking filepath.Walk: %v\n", err)
		}

		dirs = append(dirs, path)
		return nil
	})
	if err != nil {
		log.Fatalf("Unable to parse directory: %v", err)
	}

	for _, v := range dirs {
		dir, err := os.ReadDir(v)
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

// CreateBranch takes arguments and creates x amount of refs from main based on the arguments passed and returns a slice of string with said references short names
func CreateBranch(o GitOptions, g git.Repository)([]string, error){
	var refs []string
	spl := strings.Split(o.Branches, ",")

	r, err := g.Head()
	if err != nil {
		return refs, fmt.Errorf("%v",err.Error())
	}

	for _,v := range spl {
		plumb := fmt.Sprintf("refs/heads/%s", v)
		ref := plumbing.NewHashReference(plumbing.ReferenceName(plumb), r.Hash())

		store := g.Storer.SetReference(ref)
		if store != nil {
			return refs, fmt.Errorf("%v",err.Error())
		}
	}
	getRefs, err := g.References()
	if err != nil {
		return refs, fmt.Errorf("%v",err.Error())
	}

	err = getRefs.ForEach(func(ref *plumbing.Reference) error {
		if ref.Type() == plumbing.HashReference {
			refs = append(refs, ref.Name().Short())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return refs, nil
}