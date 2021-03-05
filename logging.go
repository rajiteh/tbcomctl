package tbcomctl

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rusq/dlog"

	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	None         = "<none>"
	notAvailable = "N/A"
	chatPrivate  = "private"
)

// lg is the package logger.
var lg Logger = dlog.FromContext(context.Background()) // getting default logger
// dlg is the debug logger.
var dlg Logger = nologger{}

// Logger is the interface for logging.
type Logger interface {
	Print(v ...interface{})
	Println(v ...interface{})
	Printf(format string, a ...interface{})
}

// Userinfo returns the user info.
func Userinfo(u *tb.User) string {
	if u == nil {
		return None
	}

	return fmt.Sprintf("<[ID %d] %s (%s)>", u.ID, u.Username, Nvlstring(u.LanguageCode, notAvailable))
}

// ChatInfo returns the chat info.
func ChatInfo(ch *tb.Chat) string {
	if ch == nil {
		return None
	}
	if ch.Type == chatPrivate {
		return chatPrivate + ":" + Userinfo(&tb.User{
			ID:        int(ch.ID),
			FirstName: ch.FirstName,
			LastName:  ch.LastName,
			Username:  ch.Username,
		})
	}
	var title = ""
	if ch.Title != "" {
		title = fmt.Sprintf(" (%q)", ch.Title)
	}
	return fmt.Sprintf("<[%d] %s%s>", ch.ID, ch.Type, title)
}

// Sdump dumps the structure.
func Sdump(m interface{}) string {
	var buf strings.Builder
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")
	if err := enc.Encode(m); err != nil {
		lg.Println("failed to marshal")
	}
	return buf.String()
}

// SetLogger sets the current logger.
func SetLogger(l Logger) {
	if l == nil {
		return
	}
	lg = l
}

// SetDebugLogger sets the debug logger which is used to output debug messages,
// if you must.  By default, debug logging is disabled.
func SetDebugLogger(l Logger) {
	if l == nil {
		return
	}
	dlg = l
}

// NoDebugLogger switches off debug messages.
func NoDebugLogger() {
	dlg = nologger{}
}

// GetLogger returns current logger.
func GetLogger() Logger {
	return lg
}

// NoLogging switches off default logging, if you're brave.
func NoLogging() {
	lg = nologger{}
}

type nologger struct{}

func (nologger) Print(v ...interface{})                 {}
func (nologger) Println(v ...interface{})               {}
func (nologger) Printf(format string, a ...interface{}) {}
