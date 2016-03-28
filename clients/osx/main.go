package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	var (
		cmdOut []byte
		err    error
	)

	// parse local

	cmdName := "security"
	cmdArgs := []string{"find-certificate", "-a", "/System/Library/Keychains/SystemRootCertificates.keychain"}

	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		log.Println("Error parsing root certificate")
		os.Exit(1)
	}

	lines := strings.Split(string(cmdOut), "\n")

	var certs []OSXCert
	var cert OSXCert

	counter := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if len(line) < 1 {
			continue
		}

		counter %= 13

		if MapLine(&cert, line, counter) {
			certs = append(certs, cert)
		}

		counter++
	}

	// parse remote

	// TODO get URL based on OS X version
	url := "https://s3.amazonaws.com/rootca-auditor/osx-elcapitan-trusted.json"
	response, err := http.Get(url)

	if err != nil {
		log.Println("Error parsing remote root certificate list")
		os.Exit(1)
	}

	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var remoteCerts []Cert
	err = decoder.Decode(&remoteCerts)

	if err != nil {
		log.Println("Error parsing remote root certificate list")
		os.Exit(1)
	}

	// compare certs

	for _, c := range certs {
		found := false

		cname := c.Attributes.Labl
		cserial := c.Attributes.Snbr

		for _, rc := range remoteCerts {
			rcname := rc.CertificateName
			rcserial := StripSpaces(rc.SerialNumber)

			if rcname == cname || rcserial == cserial {
				found = true
			}
		}

		if found {
			// TODO add flag to also show verified
			// fmt.Printf("VERIFIED: %s, %s\n", cname, cserial)
		} else {
			fmt.Printf("REMOTE NOT EXIST: %s, %s\n", cname, cserial)
		}
	}
}

func MapLine(cert *OSXCert, line string, counter int) bool {
	var re *regexp.Regexp

	switch counter {
	case 0:
		re = regexp.MustCompile(`^keychain: "(.+)"$`)
		match := re.FindStringSubmatch(line)
		cert.Keychain = match[1]
	case 1:
		re = regexp.MustCompile(`^version: (\d+)$`)
		match := re.FindStringSubmatch(line)
		cert.Version = match[1]
	case 2:
		re = regexp.MustCompile(`^class: 0x([0-9A-F]+)$`)
		match := re.FindStringSubmatch(line)
		cert.Class = match[1]
	case 3:
		re = regexp.MustCompile(`^attributes:$`)
	case 4:
		re = regexp.MustCompile(`^"alis"<blob>=(?:0x[0-9A-F]+\s+)?"(.+)"`)
		match := re.FindStringSubmatch(line)
		cert.Attributes.Alis = AsciiToText(match[1])
	case 5:
		re = regexp.MustCompile(`^"cenc"<uint32>=0x([0-9A-F]+)`)
		match := re.FindStringSubmatch(line)
		cert.Attributes.Cenc = match[1]
	case 6:
		re = regexp.MustCompile(`^"ctyp"<uint32>=0x([0-9A-F]+)`)
		match := re.FindStringSubmatch(line)
		cert.Attributes.Ctyp = match[1]
	case 7:
		re = regexp.MustCompile(`^"hpky"<blob>=0x([0-9A-F]+)`)
		match := re.FindStringSubmatch(line)
		cert.Attributes.Hpky = match[1]
	case 8:
		re = regexp.MustCompile(`^"issu"<blob>=0x([0-9A-F]+)`)
		match := re.FindStringSubmatch(line)
		cert.Attributes.Issu = match[1]
	case 9:
		re = regexp.MustCompile(`^"labl"<blob>=(?:0x[0-9A-F]+\s+)?"(.+)"`)
		match := re.FindStringSubmatch(line)
		cert.Attributes.Labl = AsciiToText(match[1])
	case 10:
		re = regexp.MustCompile(`^"skid"<blob>=(?:0x([0-9A-F]+)|<NULL>)`)
		match := re.FindStringSubmatch(line)
		cert.Attributes.Skid = match[1]
	case 11:
		re = regexp.MustCompile(`^"snbr"<blob>=(?:0x([0-9A-F]+)|"(.+)")`)
		match := re.FindStringSubmatch(line)
		if len(match) == 3 {
			cert.Attributes.Snbr = match[2]
		}
		cert.Attributes.Snbr = match[1]
	case 12:
		re = regexp.MustCompile(`^"subj"<blob>=0x([0-9A-F]+)`)
		match := re.FindStringSubmatch(line)
		cert.Attributes.Subj = match[1]
		if re.MatchString(line) {
			return true
		}
	}

	return false
}
