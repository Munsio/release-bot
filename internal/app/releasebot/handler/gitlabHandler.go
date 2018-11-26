package handler

import (
	"fmt"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/karriereat/release-bot/internal/pkg/config"
	"github.com/karriereat/release-bot/internal/pkg/notifier"
	gitlabApi "github.com/xanzy/go-gitlab"
	gitlabHook "gopkg.in/go-playground/webhooks.v5/gitlab"
)

// GitlabHandler to process gitlab events
type GitlabHandler struct {
	Config     *config.Config
	NotifyChan chan notifier.Message
}

func (h *GitlabHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	hook, _ := gitlabHook.New(gitlabHook.Options.Secret(h.Config.GitlabToken))
	payload, err := hook.Parse(r, gitlabHook.TagEvents, gitlabHook.PushEvents)

	api := gitlabApi.NewClient(nil, h.Config.GitlabApiToken)
	api.SetBaseURL("https://gitlab/api/v4")

	if err != nil {
		log.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)

	switch payload.(type) {

	case gitlabHook.TagEventPayload:
		hooktag := payload.(gitlabHook.TagEventPayload)

		version := strings.Split(hooktag.Ref, "/")
		log.Debug("Version: " + version[len(version)-1])
		log.Debug("ProjectId: ", hooktag.ProjectID)
		tag, _, err := api.Tags.GetTag(int(hooktag.ProjectID), version[len(version)-1])
		if err != nil {
			log.Error(err)
			return
		}

		// replace \r with \n cause markdown parser bug - https://github.com/russross/blackfriday/pull/428
		releaseText := strings.Replace(tag.Release.Description, "\r\n", "\n", -1)
		message := notifier.Message{Text: releaseText, Channel: "@martin.treml"}

		h.NotifyChan <- message
	}

	fmt.Fprintf(w, "Gitlab hook called")

}
