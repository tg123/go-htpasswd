package htpasswd

import (
	"testing"
)

func Test_Bcrypt(t *testing.T) {
	testParserGood(t, "bcrypt", AcceptBcrypt, nil, "$2y$05$bWBMg3oUStnhfy5rFvoyreviPySU6hvEmBub5wIlM/D.c5FeYJQ6O", "bar")
	testParserBad(t, "bcrypt", nil, RejectBcrypt, "$2y$0")
	testParserNot(t, "bcrypt", nil, RejectBcrypt, "plaintext")
}
