package intercom

import "net/url"

func (m IntercomManager) NotifyLeaveGuild(homeserver string) {
	target, err := url.Parse(homeserver)
	if err != nil {
		m.Logger.CheckException(err)
	}
	if target.Scheme == "sharmony" {
		target.Scheme = "https"
	} else {
		target.Scheme = "http"
	}
	target.Path = "/api/core/v1/"
}
