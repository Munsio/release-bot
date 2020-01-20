module github.com/karriereat/release-bot

go 1.13

replace gopkg.in/go-playground/webhooks.v5 v5.5.0 => github.com/Munsio/webhooks v5.5.0+incompatible

replace github.com/Sirupsen/logrus v1.3.0 => github.com/sirupsen/logrus v1.3.0

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/Sirupsen/logrus v1.3.0
	github.com/gorilla/websocket v1.4.0 // indirect
	github.com/karriereat/blackfriday-slack v0.0.0-20190109161150-1989d6c6eb35
	github.com/lusis/go-slackbot v0.0.0-20180109053408-401027ccfef5 // indirect
	github.com/lusis/slack-test v0.0.0-20190426140909-c40012f20018 // indirect
	github.com/nlopes/slack v0.4.0
	github.com/pkg/errors v0.8.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/spf13/afero v1.2.0 // indirect
	github.com/spf13/viper v1.3.1
	github.com/xanzy/go-gitlab v0.13.0
	golang.org/x/crypto v0.0.0-20190103213133-ff983b9c42bc // indirect
	golang.org/x/net v0.0.0-20190108225652-1e06a53dbb7e // indirect
	golang.org/x/oauth2 v0.0.0-20181203162652-d668ce993890 // indirect
	google.golang.org/appengine v1.4.0 // indirect
	gopkg.in/go-playground/webhooks.v5 v5.5.0
	gopkg.in/russross/blackfriday.v2 v2.0.0
)
