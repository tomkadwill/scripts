package main

import (
  "os"
  "os/exec"
  "strings"
  "path/filepath"
  "github.com/urfave/cli"
  "github.com/skratchdot/open-golang/open"
)

func main() {
  app := cli.NewApp()
  app.Name = "Open gitlab URL based on current directory and branch."
  app.Usage = "open-gitlab-branch"
  app.Action = func(c *cli.Context) error {
    open.Run("https://gitlab.ryaltoapp.com/ryalto/" + currentWorkingDir() + "/tree/" + branch())

    return nil
  }

  app.Run(os.Args)
}

func currentWorkingDir() string {
  currentWorkingDir, err := os.Getwd()
  if err != nil {
    panic(err)
  }

  return filepath.Base(currentWorkingDir)
}

func branch() string {
  out, err := exec.Command("git",  "rev-parse", "--abbrev-ref", "HEAD").Output()
  if err != nil {
    panic(err)
  }

  return strings.TrimSpace(string(out))
}
