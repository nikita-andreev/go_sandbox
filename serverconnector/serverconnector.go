package serverconnector

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
)

type ServerConnectorSettings struct {
	ServerUrl url.URL
	ApiKey    string
}

func New(serverUrl ...string) (*ServerConnectorSettings, error) {
	var baseUrl string
	baseUrl = "https://eyessdk.applitools.com"
	if serverUrl != nil && serverUrl[0] != "" {
		baseUrl = serverUrl[0]
	}
	parsedUrl, parseError := url.Parse(baseUrl)
	if parseError == nil {
		return &ServerConnectorSettings{ServerUrl: *parsedUrl}, nil
	}
	return nil, errors.New("Setting server url is failed!" + " " + parseError.Error())
}

func NewSession(sessionStartInfo SessionStartInfo, serverConnector ServerConnectorSettings) Session {
	return Session{
		StartInfo:       sessionStartInfo,
		ServerConnector: serverConnector}
}

func (settings *ServerConnectorSettings) endpointUrl() (*url.URL, error) {
	return settings.ServerUrl.Parse("api/sessions/running")
}

func (settings *ServerConnectorSettings) post(body []byte) (*http.Response, error) {
	buf := bytes.NewBuffer(body)
	postUrl, parseError := settings.endpointUrl()
	if parseError == nil {
		settings.setApiKey(postUrl)
		return http.Post(postUrl.String(), "application/json", buf)
	}
	return nil, errors.New("JOPA!")
}

func (settings *ServerConnectorSettings) setApiKey(u *url.URL) {
	if settings.ApiKey != "" {
		query := u.Query()
		query.Set("apiKey", settings.ApiKey)
		u.RawQuery = query.Encode()
	}
}
