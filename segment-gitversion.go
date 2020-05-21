package main

import (
	"encoding/json"
	"fmt"
	pwl "github.com/justjanne/powerline-go/powerline"
	"os/exec"
)

func segmentGitVersion(p *powerline) []pwl.Segment {

	out, _ := exec.Command("dotnet-gitversion").Output()
	var result map[string]interface{}

	json.Unmarshal([]byte(out), &result)
	semVersion := fmt.Sprintf("%v.%v.%v", result["Major"], result["Minor"], result["Patch"])

	return []pwl.Segment{{
		Name:       "gitVersion",
		Content:    semVersion,
		Foreground: p.theme.GitVersionFg,
		Background: p.theme.GitVersionBg,
	}}
}
