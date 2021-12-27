package main

import (
	"fmt"
	"gopractice/domain"
	"gopractice/repository"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var (
	dr *repository.DynamoDBRepo
)

func main() {
	cfgs := aws.NewConfig()
	cfgs.WithEndpoint("http://dynamodb:8000")
	cfgs.WithRegion("ap-northeast-1")
	cfgs.WithCredentials(credentials.NewStaticCredentials("dummy", "dummy", "dummy"))
	sess, err := session.NewSession(cfgs)
	logFatal(err)
	dynamodb := dynamo.New(session.Must(sess, err))
	dynamodb.CreateTable("test_list", domain.Data{}).Run()

	dr = repository.NewDynamoDBRepo(dynamodb, false)

	http.HandleFunc("/", requestHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		name := r.FormValue("name")
		data, err := dr.GetByName(name)
		if err == nil {
			message := fmt.Sprintf(`
				<h1>Name: %s</h1>
				<h3>Text: %s</h3>`,
				data.Name,
				data.Text,
			)
			fmt.Fprintf(w, message)
		} else {
			fmt.Fprintf(w, `<h1>No Item!</h1>`)
		}
	} else if r.Method == "POST" {
		data := domain.Data{
			Name: r.FormValue("name"),
			Text: r.FormValue("text"),
		}
		err := dr.Store(data)
		logFatatl(err)
		fmt.Fprintf(w)
	}
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
