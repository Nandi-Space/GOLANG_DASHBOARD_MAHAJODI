package store

import (
	"Mahajodi_GOLANG_Dashboard/models"
	"context"
	"database/sql"
	"fmt"
	"path/filepath"
	"time"

	firebase "firebase.google.com/go"
	"github.com/aws/aws-sdk-go/aws"
	awsCredentials "github.com/aws/aws-sdk-go/aws/credentials"
	awsSession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

var DBState State

type State struct {
	db          *sql.DB
	Config      *models.Config
	fApp        *firebase.App
	awsSession  *awsSession.Session
	emailClient *ses.SES
}

func InitState(config *models.Config) {
	awsConf := aws.Config{
		Credentials: awsCredentials.NewStaticCredentials(config.AWS.ID, config.AWS.Secret, ""),
		Region:      &config.AWS.Region,
	}
	sess := awsSession.Must(awsSession.NewSession(&awsConf))

	db := dbConn(config)
	DBState = State{
		db:         db,
		Config:     config,
		fApp:       initializeAppWithServiceAccount(config),
		awsSession: sess,

		emailClient: ses.New(sess, &awsConf),
	}
}

func initializeAppWithServiceAccount(config *models.Config) *firebase.App {
	absPath, err := filepath.Abs(config.Firebase.CredentialPath)
	if err != nil || config.Firebase.CredentialPath == "" {
		logrus.Fatal("firebase credential file not found in " + filepath.Join(config.Firebase.CredentialPath))
	}
	opt := option.WithCredentialsFile(absPath)
	fConfig := &firebase.Config{ProjectID: "mahajodi-matrimony"}
	app, err := firebase.NewApp(context.Background(), fConfig, opt)
	if err != nil {
		logrus.Fatal(err)
	}
	return app
}

func dbConn(config *models.Config) *sql.DB {

	var str string

	str = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name)

	db, err := sql.Open("mysql", str)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	//Check if the connection is successful by establishing a connection.
	//Retry upto 10 times if connection is not successful
	for retryCount := 0; retryCount < 10; retryCount++ {
		err := db.Ping()
		if err == nil {
			fmt.Println("database connection successful")
			return db
		}

		fmt.Println(err)
		fmt.Println("could not connect to database: retrying...")
		time.Sleep(time.Second)
	}

	logrus.Fatal("could not connect to database")
	return nil
}

func initEmailClient() {

}
