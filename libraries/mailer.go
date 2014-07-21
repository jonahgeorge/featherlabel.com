package mailer

// import (
// 	"bytes"
// 	"fmt"
// 	"log"
// 	"text/template"
//
// 	"github.com/sendgrid/sendgrid-go"
// )
//
// const (
// 	from_email = "hello@featherlabel.com"
// )
//
// var (
// 	templates *template.Template
// 	sg        *sendgrid.SGClient
// )
//
// func init() {
// 	sg := sendgrid.NewSendGridClient("sendgrid_user", "sendgrid_key")
//
// 	// parse templates
// 	// templates, err = template.ParseFiles("template.html")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// }
//
// func ExampleFunction(recipient string) {
//
// }
//
// func SendEmailVerification(name string, email string) {
// 	message := sendgrid.NewMail()
//
// 	message.AddTo(recipient_email)
// 	message.AddToName(recipient_name)
//
// 	message.SetSubject("SendGrid Testing")
// 	message.SetText("Test, Test, Test")
//
// 	message.SetFrom(from_email)
//
// 	if r := sg.Send(message); r == nil {
// 		fmt.Println("Email sent!")
// 	} else {
// 		fmt.Println(r)
// 	}
// }
//
// // Sends an email confirmation letter to the recipient argument
func SendConfirmationEmail(recipient string) {
	// 	log.Println("Email would be sent here to " + recipient)
	//
	// 	// Set the sender and recipient
	// 	client.Mail(address)
	// 	err = client.Rcpt(recipient)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	// Retrieve writer from postfix
	// 	wc, err := client.Data()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer wc.Close()
	//
	// 	buffer := new(bytes.Buffer)
	// 	templates.Execute(buffer, map[string]interface{}{
	// 		"Subject": "Welcome to Feather Label!",
	// 	})
	//
	// 	if _, err = buffer.WriteTo(wc); err != nil {
	// 		log.Fatal(err)
	// 	}
}
