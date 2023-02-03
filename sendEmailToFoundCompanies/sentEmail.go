package sendEmailToFoundCompanies

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"strings"
)

func SendEmail() {
	// to be able to read .env file we should load first
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Something went wrong with env file Env file ")
		return
	}

	emptyEmailFile := excelize.NewFile()

	f, err := excelize.OpenFile("employerH1B.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	nonEmailCount := 1
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		companyName := row[1]
		companyEmails := row[4]
		companyWebsite := row[3]
		if companyEmails != "[]" {
			// Remove square brackets from string
			array := companyEmails[1 : len(companyEmails)-1]

			// Split the string into a slice of strings using the space character as a delimiter

			emailArray := strings.Split(array, " ")
			fmt.Println(len(emailArray))
			if len(emailArray) != 0 {
				for _, email := range emailArray {
					EmailTemplate(companyName, email)

				}
			}
		} else {
			cellA := "A" + strconv.Itoa(nonEmailCount)
			cellB := "B" + strconv.Itoa(nonEmailCount)
			if err := emptyEmailFile.SetCellValue("Sheet1", cellA, companyName); err != nil {
				fmt.Println("cellA")
			}
			if err := emptyEmailFile.SetCellValue("Sheet1", cellB, companyWebsite); err != nil {
				fmt.Println("cellB")
			}
			if err := emptyEmailFile.SaveAs("nonEmailFile.xlsx"); err != nil {
				log.Fatal(err)
			}
			nonEmailCount += 1
		}

	}
}

//by default, it is .env, so we don't have to write

func EmailTemplate(companyName, email string) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Something went wrong with env file Env file ")
		return
	}
	from := os.Getenv("MAIL")
	password := os.Getenv("PASSWD")
	host := "smtp.gmail.com"
	port := "587"
	name := "Sultan Guvenbas"
	body := "To Whom It May Concern,\n\n" +
		"I am writing to express my interest in the software engineering role at " + companyName + "." +
		" As a recently graduated computer engineer with a 3.45 GPA, I am eager to bring my technical skills and passion for software engineering to a dynamic work environment." +
		"\n\nMy goal is to pursue a career in the US, and I believe that your company would provide me the opportunities I am seeking. I am motivated, ambitious, and eager to take on new challenges, " +
		"and I believe I would be a valuable asset to your team." +
		"\n\n" +
		"I have extensive experience in programming languages such as Go, Java, JavaScript, TypeScript, SQL, and Python, as well as front-end frameworks such as VueJS, React, Angular, NextJS, and PHP. " +
		"Additionally, I am proficient in creating scalable web applications using backend frameworks such as NodeJS, Gin, NestJS, ExpressJS, and TypeORM." +
		"\n\nMy technical expertise, combined with my commitment to continuous learning and growth, make me a strong candidate for this role. " +
		"I would be honored to bring my skills and experience to your team and be a part of a company making a difference in the industry." +
		"\n\n" + "I would be grateful for the opportunity to discuss my qualifications further and how I can contribute to your company's success. Please find attached my resume and portfolio for your review." +
		"\n\nThank you for considering my application. I look forward to hearing from you soon." +
		"\n\n" +
		"Note: I was searching for a company that offers H1B visa sponsorship, and I came across the website h1bgrader," +
		" which lists all the companies that have sponsored visas in the past. However, the website did not provide any " +
		"email addresses to contact these companies. So, I teamed up with my friend Ali and wrote a code to find the websites" +
		" of these companies using their names. We then searched for email addresses on these websites by writing a function that" +
		" searched for them in the text. To send the emails, we created a template email and wrote another function to send the emails. " +
		"Although writing the code wasn't particularly difficult, I believe it demonstrates my dedication and determination to make my dream of" +
		" starting my career in the US a reality. The code is available on GitHub! \n\n" +
		"\n\nBest regards," +
		"\n\n" + name + "\r\n"
	// Connect to the SMTP server
	auth := smtp.PlainAuth("", from, password, host)
	to := email
	// Read the attachment file
	file, err := ioutil.ReadFile("resume.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Encode the attachment to base64
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(file)))
	base64.StdEncoding.Encode(encoded, file)
	// Create the message
	var message bytes.Buffer
	message.WriteString("From: " + from + "\r\n")
	message.WriteString("To: " + to + "\r\n")
	message.WriteString("Subject: H1B Visa Sponsorship\r\n")
	message.WriteString("MIME-Version: 1.0\r\n")
	message.WriteString("Content-Type: multipart/mixed; boundary=frontier\r\n")
	message.WriteString("\r\n")
	message.WriteString("--frontier\r\n")
	message.WriteString("Content-Type: text/plain\r\n")
	message.WriteString("\r\n")
	message.WriteString(body)
	message.WriteString("\r\n")
	message.WriteString("--frontier\r\n")
	message.WriteString("Content-Type: application/pdf\r\n")
	message.WriteString("Content-Transfer-Encoding: base64\r\n")
	message.WriteString("Content-Disposition: attachment; filename=resume.pdf\r\n")
	message.WriteString("\r\n")
	message.Write(encoded)
	message.WriteString("\r\n")
	message.WriteString("--frontier--\r\n")
	// Send the email
	err = smtp.SendMail(host+":"+port, auth, from, []string{to}, message.Bytes())
	if err != nil {
		fmt.Println("Email Couldn't send" + companyName)
		return
	}

	fmt.Println("Email sent successfully!")
}

var count = 0
var sendEmailCompanies = excelize.NewFile()

func saveCompaniesThatReceivedEmails(companyName, email string) {
	count = count + 1
	cellA := "A" + strconv.Itoa(count)
	cellB := "B" + strconv.Itoa(count)

	err := sendEmailCompanies.SetCellValue("Sheet1", cellA, companyName)
	if err != nil {
		fmt.Println("cella", err.Error())
		return
	}
	err = sendEmailCompanies.SetCellValue("Sheet1", cellB, email)
	if err != nil {
		fmt.Println("cella", err.Error())
		return
	}
	if err := sendEmailCompanies.SaveAs("successfulEmails.xlsx"); err != nil {
		log.Fatal(err)
	}

}
