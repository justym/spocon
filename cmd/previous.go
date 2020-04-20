package cmd

import (
	"log"

	"github.com/justym/spocon/pkg"
	"github.com/justym/spocon/player"
	"github.com/spf13/cobra"
)

//PrevCmd is command for pagin previous playback
var PrevCmd = &cobra.Command{
	Use:   "prev",
	Short: "prev command go to previous palyback",
	Run:   previous,
}

func previous(cmd *cobra.Command, args []string) {
	conf := pkg.NewConfig()
	result, err := pkg.NewClient(conf)
	if err != nil {
		log.Fatalln(err)
	}

	err = player.Previous(result.Client)
	if err != nil {
		log.Fatalln(err)
	}
}
