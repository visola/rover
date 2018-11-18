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
)

var config = &ssh.ClientConfig{
	User: "pi",
	Auth: []ssh.AuthMethod{
		ssh.Password("raspberry"),
	},
	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
}

func main() {
	files := [][]string{
		[]string{"build/agent", "/home/pi/agent", "0755"},
	}

	for _, f := range files {
		local := f[0]
		remote := f[1]
		permissions := f[2]

		fmt.Printf("Uploading %s to %s (%s)", local, remote, permissions)
		if err := uploadFile(local, remote, permissions); err != nil {
			log.Fatal("Failed to upload agent.", err)
		}
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
