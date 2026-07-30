package verify

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	url2 "net/url"
	"os"
	hash2 "project/go-tasks/pkg/hash"
	"time"
)

func SendEmail(host string, token string, fileLink string, email string) (link string, errHead error) {
	hash, err := hash2.GenerateHash()
	if err != nil {
		return "", err
	}
	link = fmt.Sprintf("%s/verify/%s", host, hash)

	url, err := url2.Parse("https://send.api.mailtrap.io/api/send")
	if err != nil {
		return "", err
	}

	payload := SendingRequestBody{
		From: EmailAddress{
			Email: "hello@demomailtrap.co",
			Name:  "Mailtrap Test",
		},
		To: []EmailAddress{
			{
				Email: email,
			},
		},
		Subject: "Подтвердите регистрацию",
		Text:    fmt.Sprintf("Для потдтверждния электронной почты перейдите по ссылке: %s", link),
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("User-Agent", "MyApp/1.0 (https://myapp.com)")

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf(err.Error())
		}
	}()

	if resp.StatusCode != 200 {
		var respBody SendingErrorResponseBody
		err := json.NewDecoder(resp.Body).Decode(&respBody)
		if err != nil {
			return "", err
		}

		var errorsMes string
		for _, value := range respBody.Errors {
			errorsMes = errorsMes + value + ", "
		}

		return "", errors.New(errorsMes)
	}

	err = AddUser(email, hash, fileLink)
	if err != nil {
		return "", err
	}

	return link, nil
}

func AddUser(email string, hash string, fileLink string) error {
	file, err := os.ReadFile(fileLink)
	if err != nil {
		return err
	}

	var users []UserParams
	err = json.Unmarshal(file, &users)
	if err != nil {
		return err
	}

	users = append(users, UserParams{
		Email:  email,
		Data:   time.Now().Format("2006-01-02"),
		Hash:   hash,
		Status: "In process",
	})
	jsonByte, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(fileLink, jsonByte, 0666)
	if err != nil {
		return err
	}
	return nil
}

func VerifyLink(hash string, fileLink string) (bool, error) {
	file, err := os.ReadFile(fileLink)
	if err != nil {
		return false, err
	}

	var users []*UserParams
	err = json.Unmarshal(file, &users)
	if err != nil {
		return false, err
	}

	var checkSuccess = false
	for _, value := range users {
		if value.Hash == hash && value.Status != "Success" {
			value.Status = "Success"
			checkSuccess = true
		} else if value.Hash == hash && value.Status == "Success" {
			return true, errors.New("HASH IS VERIFIED")
		}
	}

	jsonByte, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		return false, err
	}
	err = os.WriteFile(fileLink, jsonByte, 0666)
	if err != nil {
		return false, err
	}

	return checkSuccess, nil
}
