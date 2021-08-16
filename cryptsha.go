package htpasswd

import (
	"fmt"
	"strings"

	"github.com/GehirnInc/crypt"
	_ "github.com/GehirnInc/crypt/sha256_crypt"
	_ "github.com/GehirnInc/crypt/sha512_crypt"
)

type cryptPassword struct {
	salt   string
	hashed string
	prefix string
}

// Prefixes
const PrefixCryptSha256 = "$5$"
const PrefixCryptSha512 = "$6$"

// Accepts valid passwords
func AcceptCryptSha(src string) (EncodedPasswd, error) {
	if !strings.HasPrefix(src, PrefixCryptSha256) && !strings.HasPrefix(src, PrefixCryptSha512) {
		return nil, nil
	}

	prefix := PrefixCryptSha512
	if strings.HasPrefix(src, PrefixCryptSha256) {
		prefix = PrefixCryptSha256
	}

	rest := strings.TrimPrefix(src, prefix)
	mparts := strings.SplitN(rest, "$", 2)
	if len(mparts) != 2 {
		return nil, fmt.Errorf("malformed crypt-SHA password: %s", src)
	}

	salt, hashed := mparts[0], mparts[1]
	if len(salt) > 16 {
		salt = salt[0:16]
	}
	return &cryptPassword{salt, hashed, prefix}, nil
}

// RejectCryptSha known indexes
func RejectCryptSha(src string) (EncodedPasswd, error) {
	if !strings.HasPrefix(src, PrefixCryptSha512) && !strings.HasPrefix(src, PrefixCryptSha256) {
		return nil, nil
	}
	return nil, fmt.Errorf("crypt-sha password rejected: %s", src)
}

func shaCrypt(password string, salt string, prefix string) string {

	var ret string
	var sb strings.Builder
	sb.WriteString(prefix)
	sb.WriteString(salt)
	totalSalt := sb.String()

	if prefix == PrefixCryptSha512 {
		crypt := crypt.SHA512.New()
		ret, _ = crypt.Generate([]byte(password), []byte(totalSalt))

	} else if prefix == PrefixCryptSha256 {
		crypt := crypt.SHA256.New()
		ret, _ = crypt.Generate([]byte(password), []byte(totalSalt))
	}

	return ret[len(totalSalt)+1:]
}

func (m *cryptPassword) MatchesPassword(pw string) bool {
	hashed := shaCrypt(pw, m.salt, m.prefix)
	return constantTimeEquals(hashed, m.hashed)
}
