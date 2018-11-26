package notifier

import (
	"regexp"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/karriereat/release-bot/internal/pkg/config"
	bfslack "github.com/karriereat/release-bot/pkg/slackdown"
	"github.com/nlopes/slack"
	bf "gopkg.in/russross/blackfriday.v2"
)

// Notifier base interface
type Notifier struct {
	Config *config.Config
}

// SlackNotifier struct
type SlackNotifier struct {
	client   *slack.Client
	notifier *Notifier
}

// Message is the struct for all outgoing messages
type Message struct {
	Channel string
	Text    string
}

// NewSlackNotifier builds an notifier depending on the configuration
func NewSlackNotifier(conf *config.Config) *SlackNotifier {
	client := slack.New(conf.SlackToken)
	notifier := &Notifier{Config: conf}
	return &SlackNotifier{client, notifier}
}

// Run starts the channel listener
func (sn SlackNotifier) Run(notifyChan chan Message) {

	for {
		select {
		case msg := <-notifyChan:
			{
				sn.notifyTag(msg)
			}
		}
	}

}

// NotifyTag sends an tag information to an slack channel
func (sn SlackNotifier) notifyTag(msg Message) {

	//msg.convertMarkdown()
	renderer := &bfslack.Renderer{}
	extensions := bf.CommonExtensions
	md := bf.New(bf.WithRenderer(renderer), bf.WithExtensions(extensions))
	ast := md.Parse([]byte(msg.Text))
	output := renderer.Render(ast)

	log.Info(string(output))

	params := slack.PostMessageParameters{Markdown: true}
	_, _, err := sn.client.PostMessage(msg.Channel, string(output), params)

	if err != nil {
		log.Error(err)
	}
}

func (msg *Message) convertMarkdown() {

	log.Debug(msg.Text)
	re := regexp.MustCompile("#+.*")
	headlines := re.FindAllStringSubmatchIndex(msg.Text, -1)
	text := msg.Text

	log.Debug(headlines)

	// reverse headline slice
	log.Debug(headlines)
	for i := len(headlines)/2 - 1; i >= 0; i-- {
		opp := len(headlines) - 1 - i
		headlines[i], headlines[opp] = headlines[opp], headlines[i]
	}
	log.Debug(headlines)

	for _, l := range headlines {
		line := text[l[0] : l[1]-1]
		replacement := "*" + strings.Trim(line, "#") + "*"
		text = text[:l[0]] + replacement + text[l[1]-1:]
		log.Debug(text)
	}

	msg.Text = text
}
