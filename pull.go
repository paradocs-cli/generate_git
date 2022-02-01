package gengit

import (
	"fmt"
	"github.com/go-git/go-git/v5"
)

func CheckRemote(o GitOptions, g git.Repository) (bool, error) {
 rem, err := g.Remote("origin")
 if err != nil {
 	return false, fmt.Errorf("%v", err.Error())
 }
 if len(rem.Config().URLs) != 0 {
	 rep := fmt.Sprintf("%v.git",o.RemoteOptions.RepoUrl)
	 for _,v := range rem.Config().URLs {
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

func SetRemote(o GitOptions, g git.Repository){
	check, err := CheckRemote(o, g)
	if err != nil {
		return
	}
	
	if check {

	} else if !check {

	}
}