package main

import (
	"strconv"
)

type Facility int

const (
	Kern         Facility = iota << 3
	User         Facility = iota << 3
	Mail         Facility = iota << 3
	Daemon       Facility = iota << 3
	Auth         Facility = iota << 3
	Syslog       Facility = iota << 3
	Lpr          Facility = iota << 3
	News         Facility = iota << 3
	Uucp         Facility = iota << 3
	Cron         Facility = iota << 3
	AuthPriv     Facility = iota << 3
	Ftp          Facility = iota << 3
	Local0       Facility = iota << 3
	Local1       Facility = iota << 3
	Local2       Facility = iota << 3
	Local3       Facility = iota << 3
	Local4       Facility = iota << 3
	Local5       Facility = iota << 3
	Local6       Facility = iota << 3
	Local7       Facility = iota << 3
	NoneFacility Facility = -1
	AllFacility  Facility = -2
)

var facilityOneOf = map[Facility]bool{
	Kern: true, User: true, Mail: true,
	Daemon: true, Auth: true, Syslog: true,
	Lpr: true, News: true, Uucp: true,
	Cron: true, AuthPriv: true, Ftp: true,
	Local0: true, Local1: true, Local2: true,
	Local3: true, Local4: true, Local5: true,
	Local6: true, Local7: true,
}

var facilityByName = map[string]Facility{
	"auth": Auth, "authpriv": AuthPriv,
	"cron": Cron, "daemon": Daemon,
	"ftp": Ftp, "kern": Kern, "lpr": Lpr,
	"mail": Mail, "news": News,
	"security": Auth, "syslog": Syslog,
	"user": User, "uucp": Uucp, "local0": Local0,
	"local1": Local1, "local2": Local2,
	"local3": Local3, "local4": Local4,
	"local5": Local5, "local6": Local6,
	"local7": Local7, "*": AllFacility,
}

func FacilityParse(facility string) Facility {
	if isContainOnlyDigit(facility) {
		num, err := strconv.Atoi(facility)
		if err != nil {
			return NoneFacility
		}

		facility := Facility(num)
		if facilityOneOf[facility] {
			return facility
		}

		return NoneFacility
	}

	if fac, ok := facilityByName[facility]; ok {
		return fac
	}

	return NoneFacility
}
