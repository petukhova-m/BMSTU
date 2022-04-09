package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/tls"
	"fmt"
	"github.com/skorobogatov/input"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
)

func main() {

	info := ParseFile("info.txt")

	login := info["login"]
	pass := decrypt(ScanCryptFile("config.txt"))
	host := info["host"]

	auth := smtp.PlainAuth("", login, pass, host)

	fmt.Println("Enter receiver address: ")
	receiver := input.Gets()

	fmt.Println("Enter subject: ")
	subject := input.Gets()

	fmt.Println("Enter your message: ")
	body := input.Gets()

	message := []byte ("To: " + receiver + "\r\n"+
		"From: " + login + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	config := &tls.Config {
		InsecureSkipVerify: true,
		ServerName: host,
	}

	conn, err := tls.Dial("tcp", host + ":465", config)
	if err != nil {
		log.Fatal(err)
	}

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Auth(auth); err != nil {
		log.Fatal(err)
	}

	if err = client.Mail(login); err != nil {
		log.Fatal(err)
	}

	if err = client.Rcpt(receiver); err != nil {
		log.Fatal(err)
	}

	wc, err := client.Data()
	if err != nil {
		log.Fatal(err)
	}

	_, err = wc.Write(message)
	if err != nil {
		log.Fatal(err)
	}

	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = client.Quit()
	if err != nil {
		log.Fatal(err)
	}
}

func decrypt(pass string) string  {
	text := []byte(pass)
	key := []byte("okaylet'strytomake32bytesphrase!")

	c, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Fatal(err)
	}

	if len(text) < gcm.NonceSize() {
		log.Fatal("length error!")
	}

	nonce, text := text[:gcm.NonceSize()], text[gcm.NonceSize():]
	decryptpass, err := gcm.Open(nil, nonce, text, nil)
	if err != nil {
		log.Fatal(err)
	}

	return string(decryptpass)
}

func ScanCryptFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
