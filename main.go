package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {

	loadFile := flag.String("l", "", "File path to load from")
	flag.Parse()

	servicepath := os.Getenv("SERVICEPATH")

	if *loadFile == "" {

		sess, err := session.NewSession()
		ssmSession := ssm.New(sess)
		input := &ssm.GetParametersByPathInput{
			Path:           aws.String(servicepath),
			WithDecryption: aws.Bool(true),
			Recursive:      aws.Bool(true),
		}

		resp, err := ssmSession.GetParametersByPath(input)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, param := range resp.Parameters {
			paramName := strings.Split(*param.Name, "/")
			fmt.Printf("%s=%v\n", strings.ToUpper(paramName[3]), *param.Value)
		}

		input2 := &ssm.GetParametersByPathInput{
			NextToken:      aws.String(*resp.NextToken),
			Path:           aws.String(servicepath),
			WithDecryption: aws.Bool(true),
			Recursive:      aws.Bool(true),
		}

		if resp.NextToken != nil {

			resp2, err := ssmSession.GetParametersByPath(input2)
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, param := range resp2.Parameters {
				paramName := strings.Split(*param.Name, "/")
				fmt.Printf("%s=%v\n", strings.ToUpper(paramName[3]), *param.Value)
			}
		}
	} else {
		ssmSession := ssm.New(session.New(), aws.NewConfig())
		file, err := os.Open(*loadFile)

		if err != nil {
			panic(err)
		}

		reader := bufio.NewReader(file)

		for {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			}
			env := strings.Split(string(line), ",")
			input := &ssm.PutParameterInput{
				Name:        aws.String(servicepath + "/" + env[0]),
				Type:        aws.String("String"),
				Value:       aws.String(env[1]),
				Description: aws.String(env[2]),
				Overwrite:   aws.Bool(true),
			}
			result, err := ssmSession.PutParameter(input)
			if err != nil {
				fmt.Println(result)
				fmt.Println(err)
			}

		}
	}

	return
}
