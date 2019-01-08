package htpasswd

import (
	"testing"
)

func Test_Bcrypt(t *testing.T) {
	testParserGood(t, "bcrypt", AcceptBcrypt, nil, "$2y$05$bWBMg3oUStnhfy5rFvoyreviPySU6hvEmBub5wIlM/D.c5FeYJQ6O", "bar")
	testParserGood(t, "bcrypt", AcceptBcrypt, nil, "$2b$08$hQbZuw.cHsECArUAP9mOjehaJxTG9NMJfioQIHcbC0YyXpVybhoQa", "bar")
	testParserGood(t, "bcrypt", AcceptBcrypt, nil, "$2x$05$/OK.fbVrR/bpIqNJ5ianF.CE5elHaaO4EbggVDjb8P19RukzXSM3e", "\xff\xff\xa3")
	testParserBad(t, "bcrypt", nil, RejectBcrypt, "$2y$0")
	testParserNot(t, "bcrypt", nil, RejectBcrypt, "plaintext")
}
