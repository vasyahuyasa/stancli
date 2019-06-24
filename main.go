package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nats-io/stan.go"
)

const defaultClusterID = "nats"
const defaultClientID = "stan_cli"
const defaultSubject = "default"

func main() {
	var subject, clusterID, clientID, natsUrl string
	var help bool

	flag.StringVar(&subject, "subject", defaultSubject, "Subject")
	flag.StringVar(&clusterID, "cluster", defaultClusterID, "Cluster ID")
	flag.StringVar(&clientID, "client_id", defaultClientID, "Client ID")
	flag.StringVar(&natsUrl, "url", "nats://client:123456@localhost:4222", "NATS server url")
	flag.BoolVar(&help, "help", false, "This help text")
	flag.Usage = usage

	flag.Parse()
	if help {
		flag.Usage()
	}

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsUrl))
	if err != nil {
		fmt.Println("Can not connect to nats streaming:", err)
		os.Exit(1)
	}

	r := os.Stdin
	if file := flag.Arg(0); file != "" && file != "-" {
		r, err = os.Open(file)
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println("Can not read from STDIN:", err)
		os.Exit(1)
	}

	err = sc.Publish(subject, data)
	if err != nil {
		fmt.Println("Can not send to nats streaming:", err)
		os.Exit(1)
	}

	err = sc.Close()
	if err != nil {
		fmt.Println("Can not close nats streaming connection:", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func usage() {
	fmt.Fprintf(os.Stdout, "Usage: %s [OPTIONS] [FILE]\n\nSend data to Nats streaming server.\nWith no FILE, or when FILE is -, read standard input.\n\nOptions:\n", os.Args[0])
	flag.PrintDefaults()
}
