package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "spocon",
}

func cmdInit() error {
	rootCmd.AddCommand(AuthCmd, NextCmd, PauseCmd, PrevCmd, StartCmd, StatusCmd)
	return nil
}

//Execute execute spocon commands
func Execute() {
	if err := cmdInit(); err != nil {
		log.Fatalln(err)
	}
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

// 結構シングルトンでいろんな物を用意しているけれど、本当にシングルトンにする必要性があるのかどうかを考えた方が良さそう！
// シングルトンは確かにインスタンスをプロジェクト内で一つだけ用意することはできるけれど、グローバル変数でになることには変わりがない、のでスコープが大きくなってしまうって言うのと、
// 非公開であったとして同じパッケージないならどこからでもさわれてしまうし、逆になぜここで使えるのかとかなんでここで使えないのかとかが出てきて、可読性を大きく下げる要因になりえる
// さらに大量にシングルトンを使うと、シングルトンである変数とそうでない変数をわかるようにしないとミスの元になるのと、グローバル変数だから管理が大変になる。
// 個人的にはCLIではコマンド一回でプログラムが終了するからシングルトンじゃなくて、クライアントを毎回精製して、それを使わせるようにするでいいかなと思う！
