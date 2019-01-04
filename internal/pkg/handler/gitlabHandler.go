package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/karriereat/release-bot/internal/pkg/config"
	"github.com/karriereat/release-bot/pkg/notifier"
	gitlabApi "github.com/xanzy/go-gitlab"
	gitlabHook "gopkg.in/go-playground/webhooks.v5/gitlab"
)

// GitlabHandler to process gitlab events
type GitlabHandler struct {
	MessageChannel string
	notifyChan     chan notifier.Message
	hookParser     *gitlabHook.Webhook
	apiClient      *gitlabApi.Client
}

// NewGitlabHandler creates the handler to accept gitlab webhooks
func NewGitlabHandler(conf *config.Config, notifyChan chan notifier.Message) (*GitlabHandler, error) {
	hook, err := gitlabHook.New(gitlabHook.Options.Secret(conf.Gitlab.WebhookSecret))
	if err != nil {
		log.Error(err)
		return nil, errors.New("Cannot create gitlab webhook")
	}
	api := gitlabApi.NewClient(nil, conf.Gitlab.APIToken)
	err = api.SetBaseURL(conf.Gitlab.BaseURL)
	if err != nil {
		log.Error(err)
		return nil, errors.New("Cannot create gitlab api")
	}

	h := new(GitlabHandler)
	h.MessageChannel = conf.Slack.Channel
	h.notifyChan = notifyChan
	h.hookParser = hook
	h.apiClient = api

	return h, nil
}

// ServeHTTP gets called each time the desired route is called from external
func (h *GitlabHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("Recieved payload")
	payload, err := h.hookParser.Parse(r, gitlabHook.TagEvents, gitlabHook.PushEvents)
	log.Debug("Parsed payload")
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, ok := payload.(gitlabHook.TagEventPayload); ok {
		log.Debug("Payload is TagPush event")
		hooktag := payload.(gitlabHook.TagEventPayload)

		fullVersion := strings.Split(hooktag.Ref, "/")
		version := fullVersion[len(fullVersion)-1]
		log.Debug("Get Tag from Gitlab")
		tag, _, err := h.apiClient.Tags.GetTag(int(hooktag.ProjectID), version)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		url := fmt.Sprintf("[%s](%s)", hooktag.Project.Name, hooktag.Project.WebURL)
		text := fmt.Sprintf("**There is a new release %s for: %s**\n\n\n", version, url)

		log.Debug("Build message")

		message := notifier.NewMessage()
		message.PreText = text
		message.Text = tag.Release.Description
		message.Channel = h.MessageChannel
		message.ParseMarkdown = true

		h.notifyChan <- message
		log.Debug("Message sent")
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Gitlab hook called")

}
