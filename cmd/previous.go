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
	Short: "play the previous song",
	Run:   previous,
}

func previous(_ *cobra.Command, _ []string) {
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
