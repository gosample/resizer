package main

import (
	"fmt"
	"net/url"
	"regexp"
)

type Validator struct {
	config *Configuration
}

// Checks if given request image host belongs to one in the
// white list.
func (v *Validator) CheckHostInWhiteList(requestUrl string) error {
	urlParsed, err := url.Parse(requestUrl)

	if err != nil {
		return err
	}

	var hostFound bool

	for _, host := range v.config.HostWhiteList {
		match, _ := regexp.MatchString(host, requestUrl)
		if match {
			hostFound = true
		}
	}

	if hostFound {
		return nil
	}

	return error(fmt.Errorf("Host %s not allowed", urlParsed.Host))
}

// Validates if new request size is valid or not
func (v *Validator) CheckRequestNewSize(s *Size) error {
	if s.Height > v.config.SizeLimits.Height {
		return error(fmt.Errorf("Height cannot be higher than %d", v.config.SizeLimits.Height))
	}

	if s.Width > v.config.SizeLimits.Width ||  s.Height >= v.config.SizeLimits.Height {
		return error(fmt.Errorf("Width cannot be higher than %d", v.config.SizeLimits.Width))
	}

	return nil
}
