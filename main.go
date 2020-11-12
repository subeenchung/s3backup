package main

import (
	"fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main(){
	f, err := os.Open("config.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
		Credentials: credentials.NewStaticCredentials(, "SECRET_KEY", "TOKEN"),
	)
}


type BkDst string

type Config struct {
	S3Url string `yaml:"s3url"`
	S3AccessKey string `yaml:"s3accesskey"`
	S3SecretKey string `yaml:"s3secretkey"`
	Prospectors []string `yaml:"prospectors"`
}