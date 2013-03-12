package main
 
import (
        "log"
        "net/smtp"
        "flag"
        "fmt"
        "strings"
)
 
var (
    subject = flag.String( "s","","subject of the mail" )
    body= flag.String( "b","","body of themail" )
    reciMail = flag.String( "m","","recipient mail address" )
)
 
func main() {
        // Set up authentication information.
        flag.Parse()
        sub := fmt.Sprintf("subject: %s\r\n\r\n",*subject)
        content :=  *body
        mailList := strings.Split( *reciMail,",")
 
        auth := smtp.PlainAuth(
                "",
                "joycn@163.com",
                "password",
                "smtp.163.com",
                //"smtp.gmail.com",
        )
        // Connect to the server, authenticate, set the sender and recipient,
        // and send the email all in one step.
        err := smtp.SendMail(
                "smtp.163.com:25",
                auth,
                "joycn@163.com",
                mailList,
                []byte(sub+content),
        )
        if err != nil {
                log.Fatal(err)
        }
}


