package cmd

import (
	"log"

	"github.com/justym/spocon/pkg"
	"github.com/justym/spocon/player"
	"github.com/spf13/cobra"
)

//NextCmd is command for next playback
var NextCmd = &cobra.Command{
	Use:   "next",
	Short: "next command makes lisetenning next playback",
	Run:   next,
}

func next(cmd *cobra.Command, args []string) {
	conf := pkg.NewConfig()
	result, err := pkg.NewClient(conf)
	if err != nil {
		log.Fatalln(err)
	}

	err = player.Next(result.Client)
	if err != nil {
		log.Fatalln(err)
	}
}
