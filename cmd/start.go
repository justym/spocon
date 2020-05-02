package cmd

import (
	"log"

	"github.com/justym/spocon/pkg"
	"github.com/justym/spocon/player"
	"github.com/spf13/cobra"
)

//StartCmd is command to start playback
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start/resume playback",
	Run:   start,
}

func start(_ *cobra.Command, _ []string) {
	conf := pkg.NewConfig()
	result, err := pkg.NewClient(conf)
	if err != nil {
		log.Fatalln(err)
	}

	if err := player.Start(result.Client); err != nil {
		log.Fatalln(err)
	}
}
