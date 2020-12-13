package main

import (
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/disposable"
	"github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-email-validator/pkg/ev/free"
	"github.com/go-email-validator/go-email-validator/pkg/ev/role"
	"github.com/go-email-validator/go-email-validator/pkg/ev/smtp_checker"
)

func errPanic(err error) {
	if err != nil {
		panic(err)
	}

	ev.NewDepValidator(
		map[ev.ValidatorName]ev.ValidatorInterface{
			ev.FreeValidatorName:       ev.NewFreeValidator(free.NewWillWhiteSetFree()),
			ev.RoleValidatorName:       ev.NewRoleValidator(role.NewRBEASetRole()),
			ev.DisposableValidatorName: ev.NewDisposableValidator(disposable.MailCheckerDisposable{}),
			ev.SyntaxValidatorName:     &ev.SyntaxValidator{},
			ev.MXValidatorName:         &ev.MXValidator{},
			ev.SMTPValidatorName: ev.NewWarningsDecorator(
				ev.SMTPValidator{
					Checker: smtp_checker.Checker{
						GetConn:   smtp_checker.SimpleClientGetter,
						SendMail:  smtp_checker.NewSendMail(),
						FromEmail: ev_email.EmailFromString(smtp_checker.DefaultEmail),
					},
				},
				ev.NewIsWarning(hashset.New(smtp_checker.RandomRCPTStage), func(warningMap ev.WarningSet) ev.IsWarning {
					return func(err error) bool {
						return warningMap.Contains(err.(smtp_checker.SMTPError).Stage())
					}
				}),
			),
		},
	)
}
