package serverconnector

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	ErrorParsingResponse = "Error while parsing server response"
)

type Session struct {
	StartInfo       SessionStartInfo
	RunningSession  sessionInfo
	ServerConnector ServerConnectorSettings
}

type SessionStartInfo struct {
	AgentId                 string
	AppIdOrName             string
	VerId                   string
	ScenarioIdOrName        string
	EnvName                 string
	DefaultMatchSettings    string
	BranchName              string
	ParentBranchName        string
	CompareWithParentBranch string
	Properties              string
}

type sessionInfo struct {
	Id         string
	SessionId  string
	BatchId    string
	BaselineId string
	Url        string
}

func (session *Session) StartSession() error {
	data := map[string]SessionStartInfo{"StartInfo": session.StartInfo}
	body, err := json.Marshal(data)
	if err == nil {
		runningSession, err := parseSessionResponse(session.ServerConnector.post(body))
		if err == nil {
			session.RunningSession = runningSession
			return nil
		}
		return err
	}
	return errors.New("Error while parsing SessionStartInfo data")
}

func parseSessionResponse(response *http.Response, err error) (sessionInfo, error) {
	if err == nil {
		defer response.Body.Close()
		result := sessionInfo{}
		if response.StatusCode != http.StatusCreated {
			return result, errors.New(response.Status)
		} else {
			decodeError := json.NewDecoder(response.Body).Decode(&result)
			if decodeError != nil {
				return result, errors.New(ErrorParsingResponse + " " + decodeError.Error())
			}
			return result, nil
		}
	}
	return sessionInfo{}, err
}
