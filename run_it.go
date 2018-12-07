package main

import (
	"fmt"
	"go_sandbox/serverconnector"
	"os"
)

func main() {
	//serverConnector, serverConnectorErr := serverconnector.New("http://localhost:3000/")
	serverConnector, serverConnectorErr := serverconnector.New()
	if serverConnectorErr == nil {
		serverConnector.ApiKey = os.Getenv("APPLITOOLS_API_KEY")
		session := serverconnector.NewSession(
			serverconnector.SessionStartInfo{
				AgentId:          "GO_SDK_CLIENT",
				AppIdOrName:      "GO_SDK_DEVELOPMENT",
				ScenarioIdOrName: "Fake",
			},
			*serverConnector)
		err := session.StartSession()

		if err == nil {
			fmt.Println(session.RunningSession)
		} else {
			fmt.Println("Error")
			fmt.Println(err.Error())
		}
	}
}
