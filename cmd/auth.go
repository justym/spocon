package cmd

import (
	"log"

	"github.com/justym/spocon/pkg"
	"github.com/spf13/cobra"
)

//AuthCmd is command for authentication
var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "auth command do authenticating to spotify",
	Run:   authenticate,
}

// 使ってないなら_にしておくといいかも！
func authenticate(cmd *cobra.Command, args []string) {
	conf := pkg.NewConfig()
	_, err := pkg.NewClient(conf)
	if err != nil {
		log.Fatalln(err)
	}
}
