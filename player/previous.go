package player

import (
	"errors"
	"net/http"
	"strconv"
)

const (
	prevEndpoint    = "https://api.spotify.com/v1/me/player/previous"
	prevSuccessCode = 204
)

func Previous(client *http.Client) error {
	req, err := http.NewRequest(http.MethodPost, prevEndpoint, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != prevSuccessCode {
		return errors.New("Invalid status code: " + strconv.Itoa(resp.StatusCode))
	}

	return nil
}

// エンドポイントは変わるわけではないので、pathのみを変更できるようにした方がいいかも！
// 今はhttp.Clientを受けているけど、独自のplayerのクライアントを定義して、（http.Clientとエンドポイントをもつみたいな）そいつが各メソッドを持っていると
// 最初にclientとエンドポイントを渡してあげてclientを作ってそれをcmdをNewするときに渡して、あげてればかなりいい感じになるかなと思いました！
// 現状でもplayerとcmdが別れていていいと思うので、そこを直せるとまとまりが出るし、関数さえ実装すれば他の人たちも簡単に同じインターフェースのコマンドを作れるようになるかなと思いました！
//type CmdClient struct {
//	Client PlayerClient
//}
//
//func NewCmdClient(client PlayerClient) *CmdClient {
//	return &CmdClient{
//		Client: client,
//	}
//}
//
//func (s *CmdClient) Next(cmd *cobra.Command, args []string) {
//	err := s.Client.Next()
//	if err != nil {
//		log.Fatalln(err)
//	}
//}
//
//
//type PlayerClient interface {
//	Previous() error
//	Pause() error
//	Next() error
//}
//
//type playerClient struct {
//	Client   *http.Client
//	Endpoint string
//}
//
//func NewPlayerClient(client *http.Client, endpoint string) PlayerClient {
//	return &playerClient{Client: client, Endpoint: endpoint}
//}
//
//func (p *playerClient) Previous() error {
//	req, err := http.NewRequest(http.MethodPost, p.Endpoint+"/previous", nil)
//	if err != nil {
//		return err
//	}
//
//	resp, err := p.Client.Do(req)
//	if err != nil {
//		return err
//	}
//
//	if resp.StatusCode != prevSuccessCode {
//		return errors.New("Invalid status code: " + strconv.Itoa(resp.StatusCode))
//	}
//
//	return nil
//}
//
//func (p *playerClient) Pause() error {
//	req, err := http.NewRequest(http.MethodPut, pauseEndpoint, nil)
//	if err != nil {
//		return err
//	}
//
//	resp, err := p.Client.Do(req)
//	if err != nil {
//		return err
//	}
//
//	if resp.StatusCode != pauseSuccesCode {
//		return errors.New("Invalid status code: " + strconv.Itoa(resp.StatusCode))
//	}
//
//	return err
//}
//
//func (p *playerClient) Next() error {
//	req, err := http.NewRequest(http.MethodPost, nextEndpoint, nil)
//	if err != nil {
//		return err
//	}
//
//	resp, err := p.Client.Do(req)
//	if err != nil {
//		return err
//	}
//
//	if resp.StatusCode != nextSuccessCode {
//		return errors.New("Invalid status code: " + strconv.Itoa(resp.StatusCode))
//	}
//
//	return nil
//}
