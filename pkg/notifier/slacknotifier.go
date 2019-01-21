package notifier

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/karriereat/blackfriday-slack"
	"github.com/karriereat/release-bot/internal/pkg/config"
	"github.com/nlopes/slack"
	bf "gopkg.in/russross/blackfriday.v2"
)

// SlackNotifier struct
type SlackNotifier struct {
	client     *slack.Client
	NotifyChan chan Message
	iconEmoji  string
}

// NewSlackNotifier builds an notifier depending on the configuration
func NewSlackNotifier(conf *config.Config) *SlackNotifier {
	client := slack.New(conf.Slack.Token)
	NotifyChan := make(chan Message)
	return &SlackNotifier{client, NotifyChan, conf.Slack.IconEmoji}
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
	log.Debug("Recieved message")
	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{}

	if msg.ParseMarkdown {
		log.Debug("Parse markdown")
		msg.Text = parseMarkdown(msg.Text)
		msg.PreText = parseMarkdown(msg.PreText)
		attachment.MarkdownIn = []string{"text", "pretext"}
		log.Debug("Parsing markdown finished")
	}

	attachment.Pretext = msg.PreText
	attachment.Text = msg.Text
	params.Attachments = append(params.Attachments, attachment)
	params.IconEmoji = sn.iconEmoji

	log.Debug("Send message to slack")
	_, _, err := sn.client.PostMessage(msg.Channel, "", params)
	if err != nil {
		log.Error(err)
	} else {
		log.Debug("Message send sucessfully")
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
