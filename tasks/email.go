package tasks

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"newsletter-go/internal"
	"newsletter-go/models"
	"os"
	"time"

	"github.com/hibiken/asynq"
	"github.com/mailersend/mailersend-go"
)

// this will contain the identifiers for each task handler in this file
const (
	TypeTaskSendNewsletterEmails = "email:send"
)

// the struct that contains a task handler's arguments
type SendEmailTaskPayload struct {
	Newsletter models.Newsletter
	Content    models.EmailContent
}

// this one will create the payload from the arguments and return the task object to be enqueued
func NewTaskSendNewsletterEmails(n models.Newsletter, content models.EmailContent) (*asynq.Task, error) {
	payload := SendEmailTaskPayload{
		Newsletter: n,
		Content:    content,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// the server makes the association between TypeTaskSendNewsletterEmails and HandlerTaskSendNewsletterEmails
	task := asynq.NewTask(TypeTaskSendNewsletterEmails, payloadBytes)

	return task, nil
}

// this is the task itself
func HandlerTaskSendNewsletterEmails(ctx context.Context, t *asynq.Task) error {
	var args SendEmailTaskPayload
	err := json.Unmarshal(t.Payload(), &args)
	if err != nil {
		return err
	}

	// getting a new MailerSend instance here since I couldn't send the repo in the payload
	mailersendInstance, err := internal.StartMailerSendInstance()
	if err != nil {
		return err
	}

	if len(args.Newsletter.Users) <= 0 {
		return errors.New("newsletter is empty")
	}

	backgroundCtx := context.Background()
	backgroundCtx, cancel := context.WithTimeout(backgroundCtx, 5*time.Second)
	defer cancel()

	from := mailersend.From{
		Name:  args.Newsletter.Name,
		Email: os.Getenv("MAILERSEND_DOMAIN"), // go to domains >> manage trial domain >> add SMTP user and get their address
	}

	emailContent := args.Content.ContentHTML + `
	<hr>
	<p>
		<a href="{{ Domain }}/api/newsletter/unsubscribe?email={{ Email }}&newsletter_id={{ ID }}">Click here to unsubscribe from the newsletter</a>
	</p>
	`

	// according to the MailerSend API, I can send emails to up to 10 recipients at once, but for some reason
	// I can only send to one at a time.
	for _, user := range args.Newsletter.Users {
		log.Println("Sending e-mail to", user.Name)
		var recipients []mailersend.Recipient
		recipientObj := mailersend.Recipient{
			Name:  user.Name,
			Email: user.Email,
		}
		recipients = append(recipients, recipientObj)

		// set up personalization for a link to unsubscribe from the newsletter
		personalization := []mailersend.Personalization{
			{
				Email: user.Email,
				Data: map[string]interface{}{
					"Email":  user.Email,
					"ID":     args.Newsletter.ID,
					"Domain": os.Getenv("THIS_DOMAIN"),
				},
			},
		}

		msg := &mailersend.Message{
			From:            from,
			Recipients:      recipients,
			Subject:         args.Content.Subject,
			Text:            args.Content.ContentText,
			HTML:            emailContent,
			Tags:            args.Content.Tags,
			Personalization: personalization,
		}

		_, err := mailersendInstance.Email.Send(backgroundCtx, msg)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}
