package tasks

import (
	"context"
	"errors"
	"log"
	"newsletter-go/internal"
	"newsletter-go/models"
	"time"

	"github.com/mailersend/mailersend-go"
)

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
		Email: "MS_tYtWtb@trial-3zxk54vnmmqljy6v.mlsender.net", // go to domains >> manage trial domain >> add SMTP user and get their address
	}

	// according to the MailerSend API, I can send emails to up to 10 recipients at once, but for some reason
	// I can only send to one at a time.
	for _, user := range newsletterObj.Users {
		var recipients []mailersend.Recipient
		recipientObj := mailersend.Recipient{
			Name:  user.Name,
			Email: user.Email,
		}
		recipients = append(recipients, recipientObj)

		msg := &mailersend.Message{
			From:       from,
			Recipients: recipients,
			Subject:    content.Subject,
			Text:       content.ContentText,
			HTML:       content.ContentHTML,
			Tags:       content.Tags,
		}

		_, err := mailersendInstance.Email.Send(ctx, msg)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}
