package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/bramvdbogaerde/go-scp"
	"golang.org/x/crypto/ssh"
)

const (
	pathToRPi = "raspberrypi.local:22"

	localAgentFile  = "./build/agent"
	remoteAgentFile = "/home/pi/agent"

	localFindAndKillScript  = "./scripts/find_and_kill_agent.sh"
	remoteFindAndKillScript = "/home/pi/find_and_kill_agent.sh"
)

var config = &ssh.ClientConfig{
	User: "pi",
	Auth: []ssh.AuthMethod{
		ssh.Password("raspberry"),
	},
	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
}

func main() {
	if err := uploadFile(localFindAndKillScript, remoteFindAndKillScript, "0755"); err != nil {
		log.Fatal("Failed to upload find and kill script.", err)
	}

	if err := executeRemoteCommand(remoteFindAndKillScript); err != nil {
		log.Fatal("Error while executing find and kill command.", err)
	}

	if err := uploadFile(localAgentFile, remoteAgentFile, "0755"); err != nil {
		log.Fatal("Failed to upload agent.", err)
	}

	if err := fireAndForget(remoteAgentFile); err != nil {
		log.Fatal("Error while executing find and kill command.", err)
	}
}

func estabilishSessionAndExecute(callback func(*ssh.Session) error) error {
	client, err := ssh.Dial("tcp", pathToRPi, config)
	if err != nil {
		return err
	}

	session, err := client.NewSession()
	if err != nil {
		return err
	}

	defer session.Close()

	return callback(session)
}

func executeRemoteCommand(command string) error {
	return estabilishSessionAndExecute(func(session *ssh.Session) error {
		var b bytes.Buffer
		session.Stdout = &b
		if err := session.Run(command); err != nil {
			return err
		}
		fmt.Println(b.String())

		return nil
	})
}

func fireAndForget(command string) error {
	return estabilishSessionAndExecute(func(session *ssh.Session) error {
		if err := session.Start(command); err != nil {
			return err
		}

		return nil
	})
}

func uploadFile(pathToFile string, remoteName string, permission string) error {
	client := scp.NewClient(pathToRPi, config)

	err := client.Connect()
	if err != nil {
		return err
	}

	defer client.Close()

	f, _ := os.Open(pathToFile)
	defer f.Close()

	client.CopyFile(f, remoteName, permission)

	return nil
}
