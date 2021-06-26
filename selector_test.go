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
				Facilities: map[Facility][]Priority{
					User: {Crit}, Mail: {Crit}, Daemon: {Crit},
					Auth: {Crit}, Syslog: {Crit}, Lpr: {Crit},
					News: {Crit}, Uucp: {Crit}, Cron: {Crit},
					AuthPriv: {Crit}, Ftp: {Crit}, Local0: {Crit},
					Local1: {Crit}, Local2: {Crit}, Local3: {Crit},
					Local4: {Crit}, Local5: {Crit}, Local6: {Crit},
					Local7: {Crit},
				},
			},
		},
		{
			line: "kern.*",
			expectedSelector: Selector{
				Facilities: map[Facility][]Priority{
					Kern: {Emerg, Alert, Crit, Err, Warning, Notice, Info, Debug},
				},
			},
		},
		{
			line: "kern.crit",
			expectedSelector: Selector{
				Facilities: map[Facility][]Priority{
					Kern: {Emerg, Alert, Crit},
				},
			},
		},
		{
			line: "kern.info;kern.!err",
			expectedSelector: Selector{
				Facilities: map[Facility][]Priority{
					Kern: {Warning, Notice, Info},
				},
			},
		},
		{
			line: "mail.=info",
			expectedSelector: Selector{
				Facilities: map[Facility][]Priority{
					Mail: {Info},
				},
			},
		},
		{
			line: "mail.*;mail.!=info",
			expectedSelector: Selector{
				Facilities: map[Facility][]Priority{
					Mail: {Emerg, Alert, Crit, Err, Warning, Notice, Debug},
				},
			},
		},
		{
			line: "mail,news.=info",
			expectedSelector: Selector{
				Facilities: map[Facility][]Priority{
					Mail: {Info},
					News: {Info},
				},
			},
		},
		{
			line: `*.=info;*.=notice;mail.none`,
			expectedSelector: Selector{
				Facilities: map[Facility][]Priority{
					User: {Info, Notice}, Daemon: {Info, Notice},
					Auth: {Info, Notice}, Syslog: {Info, Notice}, Lpr: {Info, Notice},
					News: {Info, Notice}, Uucp: {Info, Notice}, Cron: {Info, Notice},
					AuthPriv: {Info, Notice}, Ftp: {Info, Notice}, Local0: {Info, Notice},
					Local1: {Info, Notice}, Local2: {Info, Notice}, Local3: {Info, Notice},
					Local4: {Info, Notice}, Local5: {Info, Notice}, Local6: {Info, Notice},
					Local7: {Info, Notice}, Kern: {Info, Notice},
				},
			},
		},
		{
			line: "*.=info;mail,news.none",
			expectedSelector: Selector{
				Facilities: map[Facility][]Priority{
					User: {Info}, Daemon: {Info}, Auth: {Info},
					Syslog: {Info}, Lpr: {Info}, Uucp: {Info},
					Cron: {Info}, AuthPriv: {Info}, Ftp: {Info},
					Local0: {Info}, Local1: {Info}, Local2: {Info},
					Local3: {Info}, Local4: {Info}, Local5: {Info},
					Local6: {Info}, Local7: {Info}, Kern: {Info},
				},
			},
		},
		{
			line: `*.*`,
			expectedSelector: Selector{
				Facilities: map[Facility][]Priority{
					User: prioritiesCopy(), Mail: prioritiesCopy(), Daemon: prioritiesCopy(),
					Auth: prioritiesCopy(), Syslog: prioritiesCopy(), Lpr: prioritiesCopy(),
					News: prioritiesCopy(), Uucp: prioritiesCopy(), Cron: prioritiesCopy(),
					AuthPriv: prioritiesCopy(), Ftp: prioritiesCopy(), Local0: prioritiesCopy(),
					Local1: prioritiesCopy(), Local2: prioritiesCopy(), Local3: prioritiesCopy(),
					Local4: prioritiesCopy(), Local5: prioritiesCopy(), Local6: prioritiesCopy(),
					Local7: prioritiesCopy(), Kern: prioritiesCopy(),
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.line, func(t *testing.T) {
			var selector Selector
			err := selector.Parse(tC.line)
			require.NoError(t, err, tC.line)
			require.Equal(t, tC.expectedSelector, selector)
		})
	}
}
