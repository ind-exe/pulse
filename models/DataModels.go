package models

import (
	"fmt"
	"strings"
	"time"
)

type Notifier interface{
	ToString() string
}

type DnsModel struct {
	Type string
	Question string
	Timestamp time.Time
	IP string
}

func (dm *DnsModel) ToString() string {
	return fmt.Sprintf(
		"📡 NEW DNS REQUEST RECEIVED\n"+
		"---------------------------\n"+
		"🕒 Time     : %s\n"+
		"📍 Address  : %s\n"+
		"📌 Type     : %s\n"+
		"❓ Question : %s\n"+
		"---------------------------",
		dm.Timestamp.Format("2006-01-02 15:04:05"),
		dm.IP,
		dm.Type,
		dm.Question,
	)
	
}

type UrlModel struct {
	Method string
	HostName string
	Path string
	Port string
	Timestamp time.Time
	IP string
	Headers map[string][]string
	Body string
}

func headerToString(header map[string][]string) string {
	var sb strings.Builder
	for key, values := range header {
		sb.WriteString(fmt.Sprintf("\n\t%s: %s", key, strings.Join(values, ", ")))
	}
	return sb.String()
}

func (um *UrlModel) ToString() string {
	return fmt.Sprintf(
		"🌍 **NEW HTTP REQUEST RECEIVED**\n"+
		"---------------------------------\n"+
		"📅 Time     : %s\n"+
		"🌐 Address  : %s\n"+
		"➡️ Method   : %s\n"+
		"📌 Path     : %s\n"+
		"🔢 Port     : %s\n"+
		"📋 Headers  :\n%s\n"+
		"📝 Body     :\n```%s```\n"+
		"---------------------------------",
		um.Timestamp.Format("2006-01-02 15:04:05"),
		um.IP,
		um.Method,
		um.Path,
		um.Port,
		headerToString(um.Headers),
		um.Body,
	)
	
}