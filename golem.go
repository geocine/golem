package main

import (
    "fmt"
    "gopkg.in/AlecAivazis/survey.v1"
    "strings"
    "os/exec"
)

// the questions to ask
var qs = []*survey.Question{
    {
        Name: "type",
        Validate: survey.Required,
        Transform: GetResultString,
        Prompt: &survey.Select{
            Message: "Type of commit:",
            Options: []string{
                "feat     | new feature",
                "fix      | bug fix",
                "docs     | documentation",
                "style    | code formatting",
                "refactor | code refactor",
                "perf     | improve performance",
                "test     | add/update tests ",
                "chore    | other changes that doesn't modify src/test",
                "revert   | revert previous commit"},
            Default: "red",
        },
    },
    {
        Name:     "scope",
        Prompt:   &survey.Input{Message: "Scope of this change:"},
        Validate: survey.Required,
    },
    {
        Name:     "message",
        Prompt:   &survey.Input{Message: "Commit message:"},
        Validate: survey.Required,
    },
    {
        Name:     "jira",
        Prompt:   &survey.Input{Message: "JIRA issue(s) ID:"},
        Validate: survey.Required,
    },
}

func main() {
    // the answers will be written to this struct
    answers := struct {
        Type          string `survey:"type"`
        Scope         string                 
        Message       string
        Jira          string                    
    }{}

    // perform the questions
    err := survey.Ask(qs, &answers)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    out:= fmt.Sprintf("%s(%s): %s | %s", answers.Type, answers.Scope, answers.Jira, answers.Message)
    fmt.Printf(out)
    fmt.Println()
    output, err := exec.Command("git", "commit", "-m", out).CombinedOutput()
    if err != nil {
        fmt.Println(err.Error())
    }
    fmt.Println(string(output))
}

func GetResultString(ans interface{}) interface{} {
	transformer := survey.TransformString(SplitTrimString)
	return transformer(ans)
}

func SplitTrimString(ans string) string {
    s:= strings.Split(ans,"|")
    ans = strings.TrimSpace(s[0])
    return ans
}