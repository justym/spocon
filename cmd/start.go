package cmd

import (
	"log"

	"github.com/justym/spocon/pkg"
	"github.com/spf13/cobra"
)

//StartCmd is command to start playback
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start command start/resume playback",
	Run:   start,
}

//TODO
func start(cmd *cobra.Command, args []string) {
	conf := pkg.NewConfig()
	_, err := pkg.NewClient(conf)
	if err != nil {
		log.Fatalln(err)
	}
}
