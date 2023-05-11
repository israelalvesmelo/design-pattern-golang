package main

import (
	"fmt"
	"strings"
)

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	validEmail(from)
	b.email.from = from

	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	validEmail(to)
	b.email.to = to
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}

func validEmail(email string) {
	if !strings.Contains(email, "@") {
		panic("email should contain @")
	}

}

func sendEmailImpl(email *email) {
	fmt.Print(fmt.Sprintf("Sending email about %s to %s", email.subject, email.to))
}

type build func(*EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendEmailImpl(&builder.email)
}

func main() {
	SendEmail(func(b *EmailBuilder) {
		b.From("foo@bar.com").
			To("bar@baz.com").
			Subject("Meeting").
			Body("Hello, do you want to meet?")
	})
}
