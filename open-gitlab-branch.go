package main

import (
  "os"
  "path/filepath"
  "github.com/urfave/cli"
  "github.com/skratchdot/open-golang/open"
)

func main() {
  app := cli.NewApp()
  app.Name = "Open gitlab URL based on current directory and branch."
  app.Usage = "open-gitlab-branch"
  app.Action = func(c *cli.Context) error {

    currentWorkingDir, err := os.Getwd()
    if err != nil {
      panic(err)
    }
    dir := filepath.Base(currentWorkingDir)

    // TODO: Get current branch
    branch := "master"

    open.Run("https://gitlab.ryaltoapp.com/ryalto/" + dir + "/tree/" + branch)

    return nil
  }

  app.Run(os.Args)
}
