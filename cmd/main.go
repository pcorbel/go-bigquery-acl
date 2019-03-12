package main

import (
	"context"
	"flag"
	"os"

	"fmt"
	"os/user"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/mgutz/ansi"
	"github.com/pkg/errors"
)

var (
	confPath string
	statusOK = true
)

func init() {

	// Parse flags
	flag.StringVar(&confPath, "conf", "config/config.yaml", "The path of the configuration file")
	flag.Parse()

	// Print information on current update
	fmt.Println(ansi.Color("BigQuery update information", "yellow+b"))

	// Print creation date
	fmt.Println(fmt.Sprintf("  Created at:           %s", time.Now().UTC()))

	// Print current username
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println(errors.Wrap(err, "cannot get current user"), "red")
	}
	fmt.Println(fmt.Sprintf("  Author:               %s", currentUser.Username))

	// Print GCP credentials file
	credEnv, ok := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS")
	if ok {
		fmt.Println(fmt.Sprintf("  Credentials:          %s", credEnv))
	} else {
		fmt.Println("  Credentials:          default")
	}

	// Print configuration file
	if confPath == "config.yaml" {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(ansi.Color(fmt.Sprint(errors.Wrap(err, "cannot get current working directory")), "red"))
		}
		fmt.Print(fmt.Sprintf("  Configuration file:   %s/%s", dir, confPath))
	} else {
		fmt.Print(fmt.Sprintf("  Configuration file:   %s", confPath))
	}
}

func main() {

	if err := run(); err != nil {
		fmt.Println(ansi.Color(fmt.Sprintf("%v", err), "red"))
		statusOK = false
	}

	fmt.Println(ansi.Color("\n\nBigQuery update result", "yellow+b"))

	if statusOK {
		fmt.Println("  Status:               " + ansi.Color("success", "green"))
	} else {
		fmt.Println("  Status:               " + ansi.Color("failure", "red"))
	}
}

func run() error {

	var conf Config
	err := conf.LoadFromFile(confPath)
	if err != nil {
		return errors.Wrap(err, "cannot load configuration")
	}

	client, err := bigquery.NewClient(context.Background(), conf.Project)
	if err != nil {
		return errors.Wrap(err, "cannot create client")
	}

	err = updateAccessControl(client, conf)
	if err != nil {
		return errors.Wrap(err, "cannot update accesses")
	}

	return nil
}
