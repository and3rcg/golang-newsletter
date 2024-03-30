package tasks

import (
	"context"
	"errors"
	"log"
	"newsletter-go/internal"
	"newsletter-go/models"
	"os"
	"time"

	"github.com/hibiken/asynq"
	"github.com/mailersend/mailersend-go"
)

const (
	SendEmailTaskType = "email:send"
)

type SendEmailTaskPayload struct {
	Newsletter models.Newsletter
	Content    models.EmailContent
}

func NewSendEmailTask() (*asynq.Task, error) {
	// wip

	return nil, nil
}

func SendNewsletterEmailsTask(r *internal.Repository, newsletterObj *models.Newsletter, content models.EmailContent) error {
	mailersendInstance := r.MS

	if len(newsletterObj.Users) <= 0 {
		return errors.New("newsletter is empty")
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	from := mailersend.From{
		Name:  newsletterObj.Name,
		Email: os.Getenv("MAILERSEND_DOMAIN"), // go to domains >> manage trial domain >> add SMTP user and get their address
	}

	emailContent := content.ContentHTML + `
	<hr>
	<p>
		<a href="/api/newsletter/unsubscribe?email={{ Email }}&newsletter_id={{ ID }}">Click here to unsubscribe from the newsletter</a>
	</p>
	`

	// according to the MailerSend API, I can send emails to up to 10 recipients at once, but for some reason
	// I can only send to one at a time.
	for _, user := range newsletterObj.Users {
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
					"Email": user.Email,
					"ID":    newsletterObj.ID,
				},
			},
		}

		msg := &mailersend.Message{
			From:            from,
			Recipients:      recipients,
			Subject:         content.Subject,
			Text:            content.ContentText,
			HTML:            emailContent,
			Tags:            content.Tags,
			Personalization: personalization,
		}

		_, err := mailersendInstance.Email.Send(ctx, msg)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}
