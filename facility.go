package main

import (
	"strconv"
)

type FacilityKeyword int

const (
	Kern         FacilityKeyword = iota << 3
	User         FacilityKeyword = iota << 3
	Mail         FacilityKeyword = iota << 3
	Daemon       FacilityKeyword = iota << 3
	Auth         FacilityKeyword = iota << 3
	Syslog       FacilityKeyword = iota << 3
	Lpr          FacilityKeyword = iota << 3
	News         FacilityKeyword = iota << 3
	Uucp         FacilityKeyword = iota << 3
	Cron         FacilityKeyword = iota << 3
	AuthPriv     FacilityKeyword = iota << 3
	Ftp          FacilityKeyword = iota << 3
	Local0       FacilityKeyword = iota << 3
	Local1       FacilityKeyword = iota << 3
	Local2       FacilityKeyword = iota << 3
	Local3       FacilityKeyword = iota << 3
	Local4       FacilityKeyword = iota << 3
	Local5       FacilityKeyword = iota << 3
	Local6       FacilityKeyword = iota << 3
	Local7       FacilityKeyword = iota << 3
	NoneFacility FacilityKeyword = -1
	AllFacility  FacilityKeyword = -2
)

var facilityOneOf = map[FacilityKeyword]bool{
	Kern: true, User: true, Mail: true,
	Daemon: true, Auth: true, Syslog: true,
	Lpr: true, News: true, Uucp: true,
	Cron: true, AuthPriv: true, Ftp: true,
	Local0: true, Local1: true, Local2: true,
	Local3: true, Local4: true, Local5: true,
	Local6: true, Local7: true,
}

var facilityByName = map[string]FacilityKeyword{
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

func FacilityParse(facility string) FacilityKeyword {
	if isContainOnlyDigit(facility) {
		num, err := strconv.Atoi(facility)
		if err != nil {
			return NoneFacility
		}

		facility := FacilityKeyword(num)
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
