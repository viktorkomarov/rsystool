package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// from https://www.rsyslog.com/doc/v8-stable/configuration/sysklogd_format.html
func TestSelectorParseHappyCases(t *testing.T) {
	testCases := []struct {
		line             string
		expectedSelector Selector
	}{
		{
			line: "*.=crit;kern.none",
			expectedSelector: Selector{
				Facilities: []Facility{User, Mail, Daemon, Auth, Syslog, Lpr, News, Uucp, Cron, AuthPriv, Ftp,
					Local0, Local1, Local2, Local3, Local4, Local5, Local6, Local7},
				Priorities: []Priority{Crit},
			},
		},
		{
			line: "kern.*",
			expectedSelector: Selector{
				Facilities: []Facility{Kern},
				Priorities: []Priority{Emerg, Alert, Crit, Err, Warning, Notice, Info, Debug},
			},
		},
		{
			line: "kern.crit",
			expectedSelector: Selector{
				Facilities: []Facility{Kern},
				Priorities: []Priority{Emerg, Alert, Crit},
			},
		},
		{
			line: "kern.info;kern.!err",
			expectedSelector: Selector{
				Facilities: []Facility{Kern},
				Priorities: []Priority{Info, Notice, Warning},
			},
		},
		{
			line: "mail.=info",
			expectedSelector: Selector{
				Facilities: []Facility{Mail},
				Priorities: []Priority{Info},
			},
		},
		{
			line: "mail.*;mail.!=info",
			expectedSelector: Selector{
				Facilities: []Facility{Mail},
				Priorities: []Priority{Emerg, Alert, Crit, Err, Warning, Notice, Debug},
			},
		},
		{
			line: "mail,news.=info",
			expectedSelector: Selector{
				Facilities: []Facility{Mail, News},
				Priorities: []Priority{Info},
			},
		},
		{
			line: `*.=info;*.=notice;mail.none`,
			expectedSelector: Selector{
				Facilities: []Facility{User, Daemon, Auth, Syslog, Lpr, News, Uucp, Cron, AuthPriv, Ftp,
					Local0, Local1, Local2, Local3, Local4, Local5, Local6, Local7},
				Priorities: []Priority{Info, Notice},
			},
		},
		{
			line: "*.=info;mail,news.none",
			expectedSelector: Selector{
				Facilities: []Facility{User, Daemon, Auth, Syslog, Lpr, Uucp, Cron, AuthPriv, Ftp,
					Local0, Local1, Local2, Local3, Local4, Local5, Local6, Local7},
				Priorities: []Priority{Emerg, Alert, Crit, Err, Warning, Notice, Info, Debug},
			},
		},
		{
			line: `*.*`,
			expectedSelector: Selector{
				Facilities: []Facility{User, Mail, Daemon, Auth, Syslog, Lpr, News, Uucp, Cron, AuthPriv, Ftp,
					Local0, Local1, Local2, Local3, Local4, Local5, Local6, Local7},
				Priorities: []Priority{Emerg, Alert, Crit, Err, Warning, Notice, Info, Debug},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.line, func(t *testing.T) {
			actual, err := ParseSelector(tC.line)
			require.NoError(t, err, tC.line)
			require.Equal(t, tC.expectedSelector, actual)
		})
	}
}
