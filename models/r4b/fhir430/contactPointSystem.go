
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ContactPointSystem is documented here http://hl7.org/fhir/ValueSet/contact-point-system
type ContactPointSystem int

const (
	ContactPointSystemPhone ContactPointSystem = iota
	ContactPointSystemFax
	ContactPointSystemEmail
	ContactPointSystemPager
	ContactPointSystemUrl
	ContactPointSystemSms
	ContactPointSystemOther
)

func (code ContactPointSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ContactPointSystem) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ContactPointSystem code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "phone":
		*code = ContactPointSystemPhone
	case "fax":
		*code = ContactPointSystemFax
	case "email":
		*code = ContactPointSystemEmail
	case "pager":
		*code = ContactPointSystemPager
	case "url":
		*code = ContactPointSystemUrl
	case "sms":
		*code = ContactPointSystemSms
	case "other":
		*code = ContactPointSystemOther
	default:
		return fmt.Errorf("unknown ContactPointSystem code `%s`", s)
	}
	return nil
}
func (code ContactPointSystem) String() string {
	return code.Code()
}
func (code ContactPointSystem) Code() string {
	switch code {
	case ContactPointSystemPhone:
		return "phone"
	case ContactPointSystemFax:
		return "fax"
	case ContactPointSystemEmail:
		return "email"
	case ContactPointSystemPager:
		return "pager"
	case ContactPointSystemUrl:
		return "url"
	case ContactPointSystemSms:
		return "sms"
	case ContactPointSystemOther:
		return "other"
	}
	return "<unknown>"
}
func (code ContactPointSystem) Display() string {
	switch code {
	case ContactPointSystemPhone:
		return "Phone"
	case ContactPointSystemFax:
		return "Fax"
	case ContactPointSystemEmail:
		return "Email"
	case ContactPointSystemPager:
		return "Pager"
	case ContactPointSystemUrl:
		return "URL"
	case ContactPointSystemSms:
		return "SMS"
	case ContactPointSystemOther:
		return "Other"
	}
	return "<unknown>"
}
func (code ContactPointSystem) Definition() string {
	switch code {
	case ContactPointSystemPhone:
		return "The value is a telephone number used for voice calls. Use of full international numbers starting with + is recommended to enable automatic dialing support but not required."
	case ContactPointSystemFax:
		return "The value is a fax machine. Use of full international numbers starting with + is recommended to enable automatic dialing support but not required."
	case ContactPointSystemEmail:
		return "The value is an email address."
	case ContactPointSystemPager:
		return "The value is a pager number. These may be local pager numbers that are only usable on a particular pager system."
	case ContactPointSystemUrl:
		return "A contact that is not a phone, fax, pager or email address and is expressed as a URL.  This is intended for various institutional or personal contacts including web sites, blogs, Skype, Twitter, Facebook, etc. Do not use for email addresses."
	case ContactPointSystemSms:
		return "A contact that can be used for sending an sms message (e.g. mobile phones, some landlines)."
	case ContactPointSystemOther:
		return "A contact that is not a phone, fax, page or email address and is not expressible as a URL.  E.g. Internal mail address.  This SHOULD NOT be used for contacts that are expressible as a URL (e.g. Skype, Twitter, Facebook, etc.)  Extensions may be used to distinguish \"other\" contact types."
	}
	return "<unknown>"
}
