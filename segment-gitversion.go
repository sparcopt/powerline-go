package main

import (
	"encoding/json"
	pwl "github.com/justjanne/powerline-go/powerline"
	"os/exec"
)

type GitVersion struct {
	MajorMinorPatch string
}

func segmentGitVersion(p *powerline) []pwl.Segment {

	out, err := exec.Command("dotnet-gitversion").Output()
	
	if err != nil {
        return []pwl.Segment{}
    }

	var gitVersionOutput GitVersion
	err = json.Unmarshal([]byte(out), &gitVersionOutput)

	if err != nil || gitVersionOutput.MajorMinorPatch == "" {
		return []pwl.Segment{}
	}

	return []pwl.Segment{{
		Name:       "gitversion",
		Content:    gitVersionOutput.MajorMinorPatch,
		Foreground: p.theme.GitVersionFg,
		Background: p.theme.GitVersionBg,
	}}
}
