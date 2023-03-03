package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	const SSH_ADDRESS = "0.0.0.0:22"
	const SSH_USERNAME = "user2"
	const SSH_PASSWORD = "password"

	sshConfig := &ssh.ClientConfig{
		User:            SSH_USERNAME,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            []ssh.AuthMethod{ssh.Password(SSH_PASSWORD)},
	}

	client, err := ssh.Dial("tcp", SSH_ADDRESS, sshConfig)
	if client != nil {
		defer client.Close()
	}

	if err != nil {
		log.Fatal("Failed to dial. " + err.Error())
	}

	session, err := client.NewSession()

	if session != nil {
		defer session.Close()
	}

	if err != nil {
		log.Fatal("Failed to create session. " + err.Error())
	}

	session.Stdout = os.Stdout
	session.Stdin = os.Stdin
	session.Stderr = os.Stderr

	// err = session.Run("pwd")
	// if err != nil {
	// 	log.Fatal("Command execution error")
	// }

	// err = session.Start("/bin/bash")
	// if err != nil {
	// 	log.Fatal("Error start bash")
	// }

	// commands := []string{
	// 	"cd folder",
	// 	"cd src/myproject",
	// 	"ls",
	// 	"exit",
	// }

	// for _, cmd := range commands {
	// 	if _, err := fmt.Println(stdin, cmd); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// err = session.Wait()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// outputErr := os.Stderr.Strin

	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		log.Fatal(err)
	}

	fDestination, err := sftpClient.Create("test-file.txt")

	if err != nil {
		log.Fatal("Failed to create destination file." + err.Error())
	}

	fSource, err := os.Open("file.txt")

	if err != nil {
		log.Fatal("Failed open file. " + err.Error())
	}

	_, err = io.Copy(fDestination, fSource)
	if err != nil {
		log.Fatal("Failed copy file. " + err.Error())
	}

	log.Println("File copied")

}

func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)

	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}

	return ssh.PublicKeys(key)
}
