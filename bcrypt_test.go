package htpasswd

import (
	"testing"
)

func Test_Bcrypt(t *testing.T) {
	testParserGood(t, "bcrypt", AcceptBcrypt, nil, "$2y$05$bWBMg3oUStnhfy5rFvoyreviPySU6hvEmBub5wIlM/D.c5FeYJQ6O", "bar")
	testParserGood(t, "bcrypt", AcceptBcrypt, nil, "$2b$08$hQbZuw.cHsECArUAP9mOjehaJxTG9NMJfioQIHcbC0YyXpVybhoQa", "bar")
	testParserBad(t, "bcrypt", nil, RejectBcrypt, "$2y$0")
	testParserNot(t, "bcrypt", nil, RejectBcrypt, "plaintext")
}
