package linenotcat

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

var (
	ver string
	rev string
)

type opts struct {
	Message    string `short:"m" long:"message" default:"" description:"Send a message directly"`
	ImageFile  string `short:"i" long:"image" default:"" description:"Upload image file"`
	Tee        bool   `short:"t" long:"tee" default:"false" description:"Print STDIN to screen before posting"`
	Stream     bool   `short:"s" long:"stream" default:"false" description:"Post messages to LINE Notify continuously"`
	ConfigFile string `long:"config_file" default:"" description:"Load the specified configuration file"`
}

func parseArgs(args []string) (opt *opts, remainArgs []string) {
	o := &opts{}
	p := flags.NewParser(o, flags.Default)
	p.Usage = fmt.Sprintf("\n\nVersion: %s\nRevision: %s", ver, rev)
	remainArgs, err := p.ParseArgs(args)
	if err != nil {
		os.Exit(1)
	}
	return o, remainArgs
}

func Run(args []string) {
	o, remainArgs := parseArgs(args)

	var token string
	var err error
	if o.ConfigFile == "" {
		token, err = readDefaultToken()
	} else {
		token, err = readToken(o.ConfigFile)
	}

	if err != nil {
		panic(err)
	}

	ln := &lineNotifier{
		token: token,
	}

	if o.ImageFile != "" {
		warnIfStreamMode(o)
		warnIfArgumentRemained(remainArgs)

		msg := o.Message
		if msg == "" {
			msg = "Image file"
		}
		err := ln.notifyImage(o.ImageFile, msg, o.Tee)
		if err != nil {
			panic(err)
		}
		return
	}

	if o.Message != "" {
		// Send text message directly
		warnIfStreamMode(o)
		ln.notifyMessage(o.Message, o.Tee)
		return
	}

	if o.Stream {
		// Stream mode
		warnIfArgumentRemained(remainArgs)

		s := newStream(ln)
		go s.processStreamQueue(o.Tee)
		go s.watchStdin()
		go s.trap()
		select {}

		return
	}

	if len(remainArgs) > 0 {
		// Send file contents
		warnIfStreamMode(o)
		ln.notifyFile(remainArgs[0], o.Tee)
		return
	}

	// Send messages from STDIN
	lines := make(chan string)
	go readFromStdin(lines)

	tmpFilePath, err := writeTemp(lines)
	if err != nil {
		panic(err)
	}

	defer os.Remove(tmpFilePath)
	ln.notifyFile(tmpFilePath, o.Tee)
}

func warnIfStreamMode(o *opts) {
	if o.Stream {
		fmt.Println("Given stream option, but it is ignored when image sending mode")
	}
}

func warnIfArgumentRemained(remainArgs []string) {
	if len(remainArgs) > 0 {
		fmt.Println("Given file, but it is ignored when stream mode")
	}
}
