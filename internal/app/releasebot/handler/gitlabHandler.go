package handler

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/karriereat/release-bot/internal/pkg/config"
	"gopkg.in/go-playground/webhooks.v5/gitlab"
)

// GitlabHandler to process gitlab events
type GitlabHandler struct {
	Config *config.Config
}

func (h *GitlabHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	hook, _ := gitlab.New(gitlab.Options.Secret(h.Config.GitlabToken))

	payload, err := hook.Parse(r, gitlab.TagEvents, gitlab.PushEvents)

	fmt.Printf("%v", payload)

	w.WriteHeader(http.StatusOK)
	if err != nil {
		log.Error(err)
		if err == gitlab.ErrEventNotFound {
			// ok event wasn;t one of the ones asked to be parsed
		}
	}
	switch payload.(type) {

	case gitlab.TagEventPayload:

		log.Info("tag event payload")
		tag := payload.(gitlab.TagEventPayload)
		fmt.Fprintf(w, "%v", tag.Ref)
		// release := payload.(github.ReleasePayload)
		// // Do whatever you want from here...
		// fmt.Printf("%+v", release)
	}

	fmt.Fprintf(w, "Gitlab hook called")

}
