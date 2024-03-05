package firestore

import (
	"context"
	"sync"
	"time"

	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
	log "github.com/sirupsen/logrus"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type client struct {
	cli *firestore.Client
}

var (
	onceClientLoaded sync.Once
	firestoreCli     *client
)

// NewClient returns a new instance to use firestore
func NewClient() *client {
	onceClientLoaded.Do(func() {
		ctx := context.Background()
		opt := option.WithCredentialsJSON([]byte(
			`{
          "type": "service_account",
          "project_id": "gofinance-416207",
          "private_key_id": "45f321d2f6d080583bac64b75733c7cf9431047e",
          "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCUKL+ac0ZCYFDf\nJBNnLdlVAqUEXqs0fM3GqISm7v/EhIDWdXd7nMBZJvaXB4shZMSBuCCyIkh7Ltnm\nlEURvgCscIbyKXC9CeJ96t91DETMS3XTMvHQdT1/g0ucsKSG8W2z9cKUkOnVUnhg\nmKenDqcg77IMAjll5cA4l7NK7ihczdQky393gbUk3sVk4+sdfFzalaAaQ+YaeSqs\n8FOMlIBY3HApGrcq2gdj9w8fkBWWK2edVe/0aXxsEraghGJAj1g3YDuIdPmaD6ZU\nnJ2fwEEgNPxKqspCV6X5xjpwkGWwcoTGaVDm5uu6Kj13b8IsVlo3UJRFUHP7Y+Ut\nOqIq5Y1DAgMBAAECggEAKEMIin7P9YK6B2GE2lGl0Um2MUFOsiVGQDmUTs7891jw\nsBbc4N0gkjmZ4kH9B5acJgo07GOPodYCBfx95ovlH/c2XMFxLxXsFD7K+cJbQc8x\n1jJubzElBK1X011XYIpkPvbGqVvWB4/xQYLKcIX1Weh5UHUcRvoP6fEjmnWsx0cI\nA5Jasp7kDdtfYV8uIwL+a57OuJqXswkGT8qes57F6lRpRnWupLCMFbDk47NwIh6l\nGZLcuYVY+I44pM6YlAnklXcR8ma8iksXTZh2snmiWlY/xSLwD6MFDZQWGkONL6dA\nDZHACLeSuZ8D1rJ6xZevaisbd21qyyHLgeK5uby8fQKBgQDG7PKECVg8aFxH5NfC\nGrw/QBqWJX61TZlz8PmupGpGK8keHNSWVh1CJwoTt+IAP0u6xzKDzQr5Dh1jdezQ\n1+zPPjm3+LqzwWSOs8ktE85i0A1pwJSv7GUk0VsDNFioOTsaeJwH6aHSyeo4vA9I\noKxtexPMb3Za1PYAn00EzFbwDQKBgQC+qwRLnjT9etz1gD2dcomz167uvVTB3UlC\neiTv1kE/TI7Lbxr7bkgN0V60HTHRyUhJ29adtC+9N85WuT8yNnCpVtGVllgpWyDH\n6Z756oBviZ8sM+5BHAZRHDaGI2BrlEodAijiWYmOXBjtkmJB9Xst459xt9cIiFxG\nZ7twgBfOjwKBgG8dJW2pRpcfeARUiEwM4P+3j+I0eW1ONhv0C5OIlTAy/vkKA5kh\n6t/rrI5NFTksIeHnqIjEfI+XWcUXtrsFEWrFiDoB/k+wA7pOJF5AhPTe388lWihW\nNDz3kA3fLHJy3Vew7P7iepiiXmloamUYEphouitkS0j7UZoRYVT/ysP1AoGAMQNn\nXET0HEEMMfPgYuwBqACG0IpNEnTfEz6w5hJuvWXhwzvxdLbZwOuUa7qhhus2MP5Z\nQjsClqKU4UhHjZHfDjjqMYjvatopKBrPKjF1CLcY+tTypzwcB5e4EG1hqOX601At\n0Fnx+W2FXvTlW01Od/Jul14J4gCjr9mUnovfEtMCgYEAvKqBkJI2zEWAWOp8obDN\ntQ45gmcHR9nZYgLYmzbkivua3BrXj6rma0ePGKBfOxmat5D8myMchCLOLRvIi7/h\nT1yCedfej353d1AJdqxiDp7YtvY401QDBapqivWCXJq9z6OPeSRmP4apL7mnjSVU\n7G8cDSS8EsorAJ8fLmk8DS8=\n-----END PRIVATE KEY-----\n",
          "client_email": "financegoapp@gofinance-416207.iam.gserviceaccount.com",
          "client_id": "104069885401176897112",
          "auth_uri": "https://accounts.google.com/o/oauth2/auth",
          "token_uri": "https://oauth2.googleapis.com/token",
          "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
          "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/financegoapp%40gofinance-416207.iam.gserviceaccount.com",
          "universe_domain": "googleapis.com"
        }`,
		))
		c, err := firestore.NewClient(ctx, "gofinance-416207", opt)
		if err != nil {
			panic("Field to connect firestore" + err.Error())
		}
		firestoreCli = &client{
			cli: c,
		}
		log.Info("Connected to firestore!!")
	})
	return firestoreCli
}

func (c *client) AddDocument(collection string, data interface{}) error {
	ctx := context.Background()
	_, _, err := c.cli.Collection(collection).Add(ctx, data)
	return err
}

func (c *client) AddExpense(expense *entity.Expense) error {
	ctx := context.Background()
	_, _, err := c.cli.Collection("expenses").Add(ctx, expense)
	return err
}

func (c *client) GetUserExpense(userID string, startDate, endDate time.Time) ([]*entity.Expense, error) {
	var expenses []*entity.Expense
	ctx := context.Background()
	expensesCollection := c.cli.Collection("expenses")
	query := expensesCollection.Where("userID", "==", userID).Where("date", ">=", startDate).Where("date", "<=", endDate)

	iter := query.Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		expense := &entity.Expense{}
		doc.DataTo(expense)
		expenses = append(expenses, expense)
	}
	return expenses, nil
}

func (c *client) AddDeposit(deposit *entity.Deposit) error {
	ctx := context.Background()
	_, _, err := c.cli.Collection("deposits").Add(ctx, deposit)
	return err
}

func (c *client) GetUserDeposit(userID string, startDate, endDate time.Time) ([]*entity.Deposit, error) {
	var deposits []*entity.Deposit
	ctx := context.Background()

	depositsCollection := c.cli.Collection("deposits")
	query := depositsCollection.Where("userID", "==", userID).Where("userID", "==", userID).Where("date", ">=", startDate).Where("date", "<=", endDate)

	iter := query.Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		deposit := &entity.Deposit{}
		doc.DataTo(deposit)
		deposits = append(deposits, deposit)
	}
	return deposits, nil
}
