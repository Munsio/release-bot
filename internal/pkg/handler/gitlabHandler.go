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

	payload, err := h.hookParser.Parse(r, gitlabHook.TagEvents, gitlabHook.PushEvents)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch payload.(type) {

	case gitlabHook.TagEventPayload:
		hooktag := payload.(gitlabHook.TagEventPayload)

		version := strings.Split(hooktag.Ref, "/")
		tag, _, err := h.apiClient.Tags.GetTag(int(hooktag.ProjectID), version[len(version)-1])
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		message := notifier.NewMessage()
		message.Text = tag.Release.Description
		message.Channel = h.MessageChannel
		message.ParseMarkdown = true

		h.notifyChan <- message
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Gitlab hook called")

}
