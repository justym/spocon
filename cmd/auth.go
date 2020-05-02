package cmd

import (
	"log"

	"github.com/justym/spocon/pkg"
	"github.com/spf13/cobra"
)

//AuthCmd is command for authentication
var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "authenticate to spotify",
	Run:   authenticate,
}

func authenticate(_ *cobra.Command, _ []string) {
	conf := pkg.NewConfig()
	_, err := pkg.NewClient(conf)
	if err != nil {
		log.Fatalln(err)
	}
}
