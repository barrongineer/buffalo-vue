package cmd

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

// availableCmd represents the available command
var availableCmd = &cobra.Command{
	Use:   "available",
	Short: "A list of available buffalo plugins",
	RunE: func(cmd *cobra.Command, args []string) error {
		plugs := []Plugin{
			{
				BuffaloCommand: "root",
				Description:    vueCmd.Short,
				Name:           "vue",
			},
		}

		return json.NewEncoder(os.Stdout).Encode(plugs)
	},
}

func init() {
	RootCmd.AddCommand(availableCmd)
}

type Plugin struct {
	BuffaloCommand string `json:"buffalo_command"`
	Description    string `json:"description"`
	Name           string `json:"name"`
}
