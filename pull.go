package gengit

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"os"
)

type RemoteStats struct {
	Relationships struct {
		RemoteUrls []string
		LocalTree string
	}
	IsClean bool
}

// CheckRemote checks for the existence of a remote 'origin' and returns a boolean
func CheckRemote(o GitOptions, g git.Repository) (bool, error) {
	rem, err := g.Remote("origin")
	if err != nil {
		return false, fmt.Errorf("%v", err.Error())
	}
	if len(rem.Config().URLs) != 0 {
		rep := fmt.Sprintf("%v.git", o.RemoteOptions.RepoUrl)
		for _, v := range rem.Config().URLs {
			if rep == v {
				return true, nil
			} else {
				return false, nil
			}
		}
		return true, nil
	}
	return true, nil
}

// SetRemote uses the CheckRemote function and returns statistics of the remote repository. If CheckRemote
// returns false a new remote is created for the repository.
func SetRemote(o GitOptions, g git.Repository) (RemoteStats, error) {
	var remoteS RemoteStats
	check, err := CheckRemote(o, g)
	if err != nil {
		return remoteS, fmt.Errorf("%v", err.Error())
	}

	if check {
		tree, err := g.Worktree()
		if err != nil {
			return remoteS, fmt.Errorf("%v", err.Error())
		}

		status, err := tree.Status()
		if err != nil {
			return remoteS, fmt.Errorf("%v", err.Error())
		}

		conf, err := g.Config()
		if err != nil {
			return remoteS, fmt.Errorf("%v", err.Error())
		}

		for _,v := range conf.Remotes {
			remoteS.Relationships.RemoteUrls = append(remoteS.Relationships.RemoteUrls, v.URLs...)
		}

		remoteS.Relationships.LocalTree = os.Getenv("PWD")

		remoteS.IsClean = status.IsClean()

		return remoteS, nil
	} else if !check {
		tree, err := g.Worktree()
		if err != nil {
			return remoteS, fmt.Errorf("%v", err.Error())
		}

		status, err := tree.Status()
		if err != nil {
			return remoteS, fmt.Errorf("%v", err.Error())
		}

		rem := &config.RemoteConfig{
			Name: "origin",
			URLs: []string{
				o.RemoteOptions.RepoUrl,
			},
		}

		remote, err := g.CreateRemote(rem)
		if err != nil {
			return remoteS, err
		}

		remoteS.Relationships.RemoteUrls = remote.Config().URLs

		remoteS.Relationships.LocalTree = os.Getenv("PWD")

		remoteS.IsClean = status.IsClean()

		return remoteS, nil
	}
	return remoteS, nil
}
