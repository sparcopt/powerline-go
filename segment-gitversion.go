package main

import (
	"encoding/json"
	"fmt"
	pwl "github.com/justjanne/powerline-go/powerline"
	"os/exec"
	"strings"
)

type GitVersion struct {
	MajorMinorPatch string
	PreReleaseLabel string
}

func segmentGitVersion(p *powerline) []pwl.Segment {
	if len(p.ignoreRepos) > 0 {
		out, err := runGitCommand("git", "rev-parse", "--show-toplevel")
		if err != nil {
			return []pwl.Segment{}
		}
		out = strings.TrimSpace(out)
		if p.ignoreRepos[out] {
			return []pwl.Segment{}
		}
	}

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
		Content:    fmt.Sprintf("%s-%s", gitVersionOutput.MajorMinorPatch, gitVersionOutput.PreReleaseLabel),
		Foreground: p.theme.GitVersionFg,
		Background: p.theme.GitVersionBg,
	}}
}
