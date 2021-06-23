package main

type RsyslogFormat int

const (
	Basic RsyslogFormat = iota
	Advanced
	ObsoleteLegacy
)

type Selector struct {
	Facility string // insensetive
	Priority string // insensetive
}

type FacilityKeyword string

const (
	AuthFacility     FacilityKeyword = "auth"
	AuthPrivFacility FacilityKeyword = "authpriv"
	CronFacility     FacilityKeyword = "cron"
	DaemonFacility   FacilityKeyword = "daemon"
	FtpFacility      FacilityKeyword = "ftp"
	KernFacility     FacilityKeyword = "kern"
	LprFacility      FacilityKeyword = "lpr"
	MailFacility     FacilityKeyword = "mail"
	MarkFacility     FacilityKeyword = "mark"
	NewsFacility     FacilityKeyword = "news"
	SecurityFacility FacilityKeyword = "security" // Deprecated
	SyslogFacility   FacilityKeyword = "syslog"
	UucpFacility     FacilityKeyword = "uucp"
)

/*
#define	LOG_KERN	(0<<3)	/* kernel messages
#define	LOG_USER	(1<<3)	/* random user-level messages
#define	LOG_MAIL	(2<<3)	/* mail system
#define	LOG_DAEMON	(3<<3)	/* system daemons
#define	LOG_AUTH	(4<<3)	/* security/authorization messages
#define	LOG_SYSLOG	(5<<3)	/* messages generated internally by syslogd
#define	LOG_LPR		(6<<3)	/* line printer subsystem
#define	LOG_NEWS	(7<<3)	/* network news subsystem
#define	LOG_UUCP	(8<<3)	/* UUCP subsystem
#define	LOG_CRON	(9<<3)	/* clock daemon
#define	LOG_AUTHPRIV	(10<<3)	/* security/authorization messages (private)
#define	LOG_FTP		(11<<3)	/* ftp daemon
	/* other codes through 15 reserved for system use
#define	LOG_LOCAL0	(16<<3)	/* reserved for local use
#define	LOG_LOCAL1	(17<<3)	/* reserved for local use
#define	LOG_LOCAL2	(18<<3)	/* reserved for local use
#define	LOG_LOCAL3	(19<<3)	/* reserved for local use
#define	LOG_LOCAL4	(20<<3)	/* reserved for local use
#define	LOG_LOCAL5	(21<<3)	/* reserved for local use
#define	LOG_LOCAL6	(22<<3)	/* reserved for local use
#define	LOG_LOCAL7	(23<<3)	/* reserved for local use
*/

type Priority int

const (
	LogEmerg Priority = iota
	LogAlert
	LogCrit
	LogErr
	LogWarning
	LogNotice
	LogInfo
	LogDebug
	LogNone
)

var logName = map[Priority]string{
	LogEmerg:   "emerg",
	LogAlert:   "alert",
	LogCrit:    "crit",
	LogDebug:   "debug",
	LogInfo:    "info",
	LogNotice:  "notice",
	LogWarning: "warn",
	LogErr:     "err",
	LogNone:    "none",
}

func PriorityName(p Priority) string {
	if n, ok := logName[p]; ok {
		return n
	}

	return "unknown"
}

func main() {
	// write test TDD
	// to int for easire ordering
}
