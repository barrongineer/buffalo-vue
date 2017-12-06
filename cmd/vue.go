package cmd

import (
	"errors"
	"os"
	"os/exec"

	"fmt"

	"github.com/spf13/cobra"
)

var app = &BuffaloApp{}

// vueCmd represents the vue command
var vueCmd = &cobra.Command{
	Use:   "vue [name]",
	Short: "Creates a new Buffalo / Vue.js application",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("you must enter a name for your new application")
		}
		app.Name = args[0]

		flags := buildFlags([]string{"new", app.Name})

		buffalo := exec.Command("buffalo", flags...)
		buffalo.Stdout = os.Stdout
		buffalo.Stderr = os.Stderr

		fmt.Printf("\nCreating new Buffalo app with: %s\n\n", buffalo.Args)
		err := buffalo.Run()
		if err != nil {
			return err
		}

		fmt.Println("Initial buffalo app has been created. Converting to Vue...")
		// convert to vue

		return nil
	},
}

func init() {
	RootCmd.AddCommand(vueCmd)

	vueCmd.Flags().BoolVarP(&app.Force, "force", "f", false, "delete and remake if the app already exists")
	vueCmd.Flags().BoolVarP(&app.Verbose, "verbose", "v", false, "verbosely print out the go get commands")
	vueCmd.Flags().BoolVar(&app.SkipPop, "skip-pop", false, "skips adding pop/soda to your app")
	vueCmd.Flags().BoolVar(&app.WithDep, "with-dep", false, "adds github.com/golang/dep to your app")
	vueCmd.Flags().BoolVar(&app.SkipYarn, "skip-yarn", false, "skip to use npm as the asset package manager")
	vueCmd.Flags().StringVar(&app.DBType, "db-type", "", "specify the type of database you want to use [postgres, mysql, sqlite3]")
	vueCmd.Flags().StringVar(&app.Docker, "docker", "", "specify the type of Docker file to generate [none, multi, standard]")
	vueCmd.Flags().StringVar(&app.CIProvider, "ci-provider", "", "specify the type of ci file you would like buffalo to generate [none, travis, gitlab-ci]")
}

func buildFlags(flags []string) []string {
	if app.Force {
		flags = append(flags, "--force")
	}

	if app.Verbose {
		flags = append(flags, "--verbose")
	}

	if app.SkipPop {
		flags = append(flags, "--skip-pop")
	}

	if app.WithDep {
		flags = append(flags, "--with-dep")
	}

	if app.SkipYarn {
		flags = append(flags, "--skip-yarn")
	}

	if app.DBType != "" {
		flags = append(flags, "--db-type", app.DBType)
	}

	if app.Docker != "" {
		flags = append(flags, "--docker", app.Docker)
	}

	if app.CIProvider != "" {
		flags = append(flags, "--ci-provider", app.CIProvider)
	}

	return flags
}

type BuffaloApp struct {
	Name       string
	Force      bool
	Verbose    bool
	SkipPop    bool
	WithDep    bool
	SkipYarn   bool
	DBType     string
	Docker     string
	CIProvider string
}
