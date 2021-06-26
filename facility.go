package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrIllegalFacility error = errors.New("illegal facility")
)

type Facility int

const (
	Kern        Facility = iota << 3
	User        Facility = iota << 3
	Mail        Facility = iota << 3
	Daemon      Facility = iota << 3
	Auth        Facility = iota << 3
	Syslog      Facility = iota << 3
	Lpr         Facility = iota << 3
	News        Facility = iota << 3
	Uucp        Facility = iota << 3
	Cron        Facility = iota << 3
	AuthPriv    Facility = iota << 3
	Ftp         Facility = iota << 3
	Local0      Facility = iota << 3
	Local1      Facility = iota << 3
	Local2      Facility = iota << 3
	Local3      Facility = iota << 3
	Local4      Facility = iota << 3
	Local5      Facility = iota << 3
	Local6      Facility = iota << 3
	Local7      Facility = iota << 3
	AllFacility Facility = -2
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

func facilitiesCopy() []Facility {
	return []Facility{User, Mail, Daemon, Auth, Syslog, Lpr, News, Uucp, Cron, AuthPriv, Ftp,
		Local0, Local1, Local2, Local3, Local4, Local5, Local6, Local7}
}

func FacilityParse(line string) ([]Facility, error) {
	parts := strings.Split(line, ",")
	result := make([]Facility, 0)
	for _, part := range parts {
		if isContainOnlyDigit(part) {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("%w: %s", ErrIllegalFacility, err)
			}

			facility := Facility(num)
			if !facilityOneOf[facility] {
				return nil, fmt.Errorf("%w: unknown num %d", ErrIllegalFacility, num)
			}

			result = append(result, facility)
			continue
		}

		if facility, ok := facilityByName[part]; ok {
			result = append(result, facility)
		} else {
			return nil, fmt.Errorf("%w: unknown %s", ErrIllegalFacility, part)
		}
	}

	for _, f := range result {
		if f == AllFacility {
			return facilitiesCopy(), nil
		}
	}

	return result, nil
}
