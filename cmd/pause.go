package cmd

import (
	"log"

	"github.com/justym/spocon/pkg"
	"github.com/justym/spocon/player"
	"github.com/spf13/cobra"
)

//PauseCmd is command for pause playback
var PauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "pause command pauses current playback",
	Run:   pause,
}

func pause(cmd *cobra.Command, args []string) {
	conf := pkg.NewConfig()
	result, err := pkg.NewClient(conf)
	if err != nil {
		log.Fatalln(err)
	}

	err = player.Pause(result.Client)
	if err != nil {
		log.Fatalln(err)
	}
}
