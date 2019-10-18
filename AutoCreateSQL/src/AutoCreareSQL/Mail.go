package main

import (
	"fmt"
	"mailSender"
	"net/mail"
	"time"
)

func SetSendMail() (*mailSender.MailSender, *mail.Address)  {

	if config.Int(CFG_MAIL_ENABLE,0) == 1 {
		mailGenera := &mailSender.MailSender{
			Server: config.Str(CFG_MAIL_SERVER, ""),
			Port:   config.Str(CFG_MAIL_PORT,""),
			User:   config.Str(CFG_MAIL_USER,""),
			Pass:   config.Str(CFG_MAIL_PASS,""),

			From:   mail.Address{
				Name:     config.Str(CFG_MAIL_FROM_NAME,""),
				Address:  config.Str(CFG_MAIL_FROM_ADDRESS,""),
			},
		}

		mailTo := &mail.Address{
			Name:     config.Str(CFG_MAIL_TO_NAME,""),
			Address:  config.Str(CFG_MAIL_TO_ADDRESS,""),
		}

		return mailGenera, mailTo
	}
	return nil, nil
}

func SendMailError(err string, timeBegin, timeEnd time.Time)  {
	mailGenera, mailTo := SetSendMail()
	if mailGenera == nil || mailTo == nil {
		return
	}

	PG := config.Str(PG_NAME,"")

	subject := "[ERROR]" + PG
	content :=  fmt.Sprintln("Process Auto Create SQL")
	content += "\r\n Notice from log"
	content += "\r\n Time start : " +timeBegin.Format(FORMAT_TIME_NOW)
	content += "\r\n Time end   : " +timeEnd.Format(FORMAT_TIME_NOW)
	content +=  fmt.Sprintf("\r\n%v",err)
	content += "\r\n Please go to the log file to see the error details"
	content += "\r\n---------------------------------------------\r\n"
	content += gLogBuf.String()

	e := mailGenera.SendText(*mailTo, subject, content)

	if e != nil {
		gLog.Println(e)
	}
}

func SendMailSuccess(timeBegin, timeEnd time.Time)  {
	mailGenera, mailTo := SetSendMail()
	if mailGenera == nil || mailTo == nil {
		return
	}

	PG := config.Str(PG_NAME,"")

	subject := "[SUCCESS]" + PG
	content :=  fmt.Sprintln("Process Auto Create SQL")
	content += "\r\n Notice from log"
	content += "\r\n Time start : " +timeBegin.Format(FORMAT_TIME_NOW)
	content += "\r\n Time end   : " +timeEnd.Format(FORMAT_TIME_NOW)
	content += "\r\n---------------------------------------------\r\n"
	content += gLogBuf.String()

	e := mailGenera.SendText(*mailTo, subject, content)

	if e != nil {
		gLog.Println(e)
	}
}