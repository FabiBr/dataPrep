package main

import (
	"io"
	"log"
	"strings"
	"unicode"
)

var (
	// Trace is a logging handler
	Trace *log.Logger
	// Info is a logging handler
	Info *log.Logger
	// Warning is a logging handler
	Warning *log.Logger
	// Debug is a logging handler
	Debug *log.Logger
	// Error is a logging handler
	Error *log.Logger
)

func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	debugHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE : ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO : ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Debug = log.New(debugHandle,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARN : ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERR  : ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func NewPaths() *Paths {
	return &Paths{
		TweetList:       "",
		PartyHandleList: "",
		PartyResultList: "",
		OldResultList:   "",
		NewResultList:   "",
	}

}
