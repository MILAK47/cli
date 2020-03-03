package cmd

import (
	"fmt"

	helper "github.com/home-assistant/cli/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var audioReloadCmd = &cobra.Command{
	Use:     "reload",
	Aliases: []string{"refresh", "re"},
	Short:   "Reload the Home Assistant Audio updating information",
	Long: `
Reloading the Home Assistant Audio, triggers the to regather
all data and devices it currently has.`,
	Example: `
  ha audio reload`,
	Run: func(cmd *cobra.Command, args []string) {
		log.WithField("args", args).Debug("audio reload")

		section := "audio"
		command := "reload"
		base := viper.GetString("endpoint")

		resp, err := helper.GenericJSONPost(base, section, command, nil)
		if err != nil {
			fmt.Println(err)
			ExitWithError = true
		} else {
			ExitWithError = !helper.ShowJSONResponse(resp)
		}

		return
	},
}

func init() {
	audioCmd.AddCommand(audioReloadCmd)
}
