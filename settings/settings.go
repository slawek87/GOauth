package settings

import (
	"github.com/slawek87/gophe/settings"
)

func Settings() *settings.Settings {
	return settings.SetSettings("./settings/default.cfg")
}