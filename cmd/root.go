package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "spocon",
}

func cmdInit() error {
	rootCmd.AddCommand(AuthCmd,NextCmd,PauseCmd,PrevCmd,StatusCmd,)
	return nil
}

//Execute execute spocon commands
func Execute(){
	if err := cmdInit(); err != nil {
		log.Fatalln(err)
	}
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

	




