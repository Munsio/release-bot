package notifier

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/karriereat/blackfriday-slack"
	"github.com/nlopes/slack"
	bf "gopkg.in/russross/blackfriday.v2"
)

// SlackNotifier struct
type SlackNotifier struct {
	client     *slack.Client
	NotifyChan chan Message
}

// NewSlackNotifier builds an notifier depending on the configuration
func NewSlackNotifier(token string) *SlackNotifier {
	client := slack.New(token)
	NotifyChan := make(chan Message)
	return &SlackNotifier{client, NotifyChan}
}

// Run starts the channel listener
func (sn SlackNotifier) Run() {
	for {
		select {
		case msg := <-sn.NotifyChan:
			{
				sn.notify(msg)
			}
		}
	}
}

// notify sends the message to the slack channel defined in msg
func (sn SlackNotifier) notify(msg Message) {

	params := slack.PostMessageParameters{}
	if msg.ParseMarkdown {
		params.Markdown = true
		msg.Text = parseMarkdown(msg.Text)
	}
	_, _, err := sn.client.PostMessage(msg.Channel, msg.Text, params)

	if err != nil {
		log.Error(err)
	}
}

func parseMarkdown(text string) string {

	// replace \r with \n cause markdown parser bug - https://github.com/russross/blackfriday/pull/428
	parseText := strings.Replace(text, "\r\n", "\n", -1)

	renderer := &slackdown.Renderer{}
	extensions := bf.CommonExtensions
	md := bf.New(bf.WithRenderer(renderer), bf.WithExtensions(extensions))
	ast := md.Parse([]byte(parseText))
	output := renderer.Render(ast)

	return string(output)
}
