package infra

import (
	"context"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/slack-go/slack/slacktest"
)

// SlackLocalNotifier はローカル用のSlack通知の実装です
// slacktestのTestServerを使って実際のSlack通知をすることなくモックしています
// TestServerはStart()してもサーバは永続的に待機してくれるわけではないのでNotifyメソッド内で呼び出してます
// Testserverは呼び出すたびに空いたport番号をランダムに選びます
// 本来はSlackNotifier内で繋ぎ込むURLを変更できるのが良いですが、上記のようなTestServerの都合上ローカル用の実装を作っています
type SlackLocalNotifier struct{}

// NewSlackLocalNotifier ...
func NewSlackLocalNotifier() *SlackLocalNotifier {
	return &SlackLocalNotifier{}
}

// Notify ...
func (s *SlackLocalNotifier) Notify(ctx context.Context, message string) error {
	ts := slacktest.NewTestServer()
	defer ts.Stop()
	ts.Start()
	endpoint := ts.GetAPIURL() + "chat.postMessage"
	// NOTE: 本来はSlackが期待するjsonフォーマットでbodyを定義する必要がありますが簡単のために省略しています
	res, err := http.Post(endpoint, "application/json; charset=UTF-8", strings.NewReader(message))
	if err != nil {
		return errors.Wrap(err, "")
	}
	// 疎通確認
	b, _ := io.ReadAll(res.Body)
	log.Printf("%v", string(b))
	return nil
}
