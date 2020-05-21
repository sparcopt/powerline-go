package main

import (
	"encoding/json"
	"fmt"
	pwl "github.com/justjanne/powerline-go/powerline"
	"os/exec"
)

func segmentGitVersion(p *powerline) []pwl.Segment {

	out, err := exec.Command("dotnet-gitversion").Output()

	if err != nil {
        return []pwl.Segment{}
    }

	var result map[string]interface{}

	json.Unmarshal([]byte(out), &result)
	semVersion := fmt.Sprintf("%v.%v.%v", result["Major"], result["Minor"], result["Patch"])

	return []pwl.Segment{{
		Name:       "gitversion",
		Content:    semVersion,
		Foreground: p.theme.GitVersionFg,
		Background: p.theme.GitVersionBg,
	}}
}
