package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"text/template"

	"github.com/spf13/cobra"
)

var vueCmdFlags = &Flags{}

// vueCmd represents the vue command
var vueCmd = &cobra.Command{
	Use:   "vue [name]",
	Short: "Creates a new Buffalo / Vue.js application",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("you must enter a name for your new application")
		}
		vueCmdFlags.Name = args[0]

		flags := buildFlags([]string{"new", vueCmdFlags.Name})

		buffalo := exec.Command("buffalo", flags...)
		buffalo.Stdout = os.Stdout
		buffalo.Stderr = os.Stderr

		fmt.Printf("\nCreating new Buffalo app with: %s\n\n", buffalo.Args)
		err := buffalo.Run()
		if err != nil {
			return err
		}

		fmt.Println("Initial buffalo app has been created. Converting to Vue...")
		//app := meta.New(fmt.Sprintf("./%s", vueCmdFlags.Name))

		t, err := template.
			New("templates/package.json.tmpl").
			ParseFiles("templates/package.json.tmpl")
		if err != nil {
			return err
		}

		err = t.Execute(os.Stdout, nil)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(vueCmd)

	vueCmd.Flags().BoolVarP(&vueCmdFlags.Force, "force", "f", false, "delete and remake if the app already exists")
	vueCmd.Flags().BoolVarP(&vueCmdFlags.Verbose, "verbose", "v", false, "verbosely print out the go get commands")
	vueCmd.Flags().BoolVar(&vueCmdFlags.SkipPop, "skip-pop", false, "skips adding pop/soda to your app")
	vueCmd.Flags().BoolVar(&vueCmdFlags.WithDep, "with-dep", false, "adds github.com/golang/dep to your app")
	vueCmd.Flags().BoolVar(&vueCmdFlags.SkipYarn, "skip-yarn", false, "skip to use npm as the asset package manager")
	vueCmd.Flags().StringVar(&vueCmdFlags.DBType, "db-type", "", "specify the type of database you want to use [postgres, mysql, sqlite3]")
	vueCmd.Flags().StringVar(&vueCmdFlags.Docker, "docker", "", "specify the type of Docker file to generate [none, multi, standard]")
	vueCmd.Flags().StringVar(&vueCmdFlags.CIProvider, "ci-provider", "", "specify the type of ci file you would like buffalo to generate [none, travis, gitlab-ci]")
}

func buildFlags(flags []string) []string {
	if vueCmdFlags.Force {
		flags = append(flags, "--force")
	}

	if vueCmdFlags.Verbose {
		flags = append(flags, "--verbose")
	}

	if vueCmdFlags.SkipPop {
		flags = append(flags, "--skip-pop")
	}

	if vueCmdFlags.WithDep {
		flags = append(flags, "--with-dep")
	}

	if vueCmdFlags.SkipYarn {
		flags = append(flags, "--skip-yarn")
	}

	if vueCmdFlags.DBType != "" {
		flags = append(flags, "--db-type", vueCmdFlags.DBType)
	}

	if vueCmdFlags.Docker != "" {
		flags = append(flags, "--docker", vueCmdFlags.Docker)
	}

	if vueCmdFlags.CIProvider != "" {
		flags = append(flags, "--ci-provider", vueCmdFlags.CIProvider)
	}

	return flags
}

type Flags struct {
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
