package date

import "regexp"

var monthStrLong = "(january|february|march|april|may|june|july|august|september|october|november|december)"
var monthStrShort = "(jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec)"
var op = "(?i)(^|\\s|\\()"
var cl = "(\\s|,|\\.|\\)|$)"

var patterns = map[string]*regexp.Regexp{
	"02-01-2006": regexp.MustCompile(op + "(0?[1-9]|1[0-9]|2[0-9]|3[0-1])-(0?[1-9]|1[0-2])-(\\d{4})" + cl),
	"02.01.2006": regexp.MustCompile(op + "(0?[1-9]|1[0-9]|2[0-9]|3[0-1])[.](0?[1-9]|1[0-2])[.](\\d{4})" + cl),
	"02/01/2006": regexp.MustCompile(op + "(0?[1-9]|1[0-9]|2[0-9]|3[0-1])/(0?[1-9]|1[0-2])/(\\d{4})" + cl),
	"02-01-06":   regexp.MustCompile(op + "(0?[1-9]|1[0-9]|2[0-9]|3[0-1])-(0?[1-9]|1[0-2])-(\\d{2})" + cl),
	"02.01.06":   regexp.MustCompile(op + "(0?[1-9]|1[0-9]|2[0-9]|3[0-1])[.](0?[1-9]|1[0-2])[.](\\d{2})" + cl),
	"02/01/06":   regexp.MustCompile(op + "(0?[1-9]|1[0-9]|2[0-9]|3[0-1])/(0?[1-9]|1[0-2])/(\\d{2})" + cl),

	"January 02, 2006": regexp.MustCompile(op + monthStrLong + "\\s+(0[1-9]|1[0-9]|2[0-9]|3[0-1])(\\s+)?,\\s+(\\d{4})" + cl),
	"January 02 2006":  regexp.MustCompile(op + monthStrLong + "\\s+(0[1-9]|1[0-9]|2[0-9]|3[0-1])\\s+(\\d{4})" + cl),
	"January 2, 2006":  regexp.MustCompile(op + monthStrLong + "\\s+[1-9](\\s+)?,\\s+(\\d{4})" + cl),
	"January 2 2006":   regexp.MustCompile(op + monthStrLong + "\\s+[1-9]\\s+(\\d{4})" + cl),
	"January 2006":     regexp.MustCompile(op + monthStrLong + "\\s+(\\d{4})" + cl),
}
