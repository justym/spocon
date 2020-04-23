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

// コマンド全体的に言えるけど、シングルトンなのだからConfigを毎回読んだり、NewClientを読んだとしてもインスタンスができるわけではないけれど、
// でも使う側はそれは隠蔽されていてわからないし、あとはここのClient作成は毎回する必要性はないような気がする？
// 例えばオプショナルパターンチックにするなら、以下のようにするとかかな？
//func NewNextCmd(client *a2ns.AuthorizedClient) *cobra.Command {
//	return &cobra.Command{
//		Use:   "next",
//		Short: "next command makes lisetenning next playback",
//		Run:   newNextCmdRun(client),
//	}
//}
//
//func newNextCmdRun(client *a2ns.AuthorizedClient) func(cmd *cobra.Command, args []string) {
//	return func(cmd *cobra.Command, args []string) {
//		err := player.Next(client.Client)
//		if err != nil {
//			log.Fatalln(err)
//		}
//	}
//}

// もしくは別の型を用意して以下のようにコマンドを持たせてもいいかもね！
//type CmdClient a2ns.AuthorizedClient
//
//func (s *CmdClient) Next(cmd *cobra.Command, args []string) {
//	err := player.Next(s.Client)
//	if err != nil {
//		log.Fatalln(err)
//	}
//}
//
//func NewCmd(client *SpotifyClient) []*cobra.Command {
//	return []*cobra.Command{{
//		Use:   "next",
//		Short: "next command makes lisetenning next playback",
//		Run:   client.Next,
//	}}
//}
// rootCmd.AddCommand(...NewCmd(client))
