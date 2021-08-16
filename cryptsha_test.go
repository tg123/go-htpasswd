package htpasswd

import (
	"fmt"
	"testing"
)

type shacryptDatum struct {
	password string
	salt     string
	hashed   string
	prefix   string
}

var sha256TestData = []shacryptDatum{
	{"mickey", "123456", "2hClNSDw3lZ0X/9PFBSI2eCGMOS06v6IbChiRsjy6tA", PrefixCryptSha256},
	{"paul1", "654321", "yXh20wwTHRwjSLcw20kQtiViO9n7HXDgEvzWf.cDks4", PrefixCryptSha256},
	{"princessbuttercup", "gildor", "/96zrUL6Si5ApMDxKlIvMHefBZz04JXJeg.Lp1fjhg1", PrefixCryptSha256},
}

var sha512TestData = []shacryptDatum{
	{"vinnie6", "123456", "By3XGEfRf2RwFvWYR0kHRVJGq2/IKwLEGQxwyncoP88TGiBzHMBmvrTNxHgyqrmhZ/M7CGtkfIw0rBRfewW.y1", PrefixCryptSha512},
	{"mickey5", "iklkG8zV969+0x+f", "XKxre3pm8QNHezNxyEXj51AkNy5AXJQKifFhVWqhVaLLUAUAZkDy6Dp1Th/mTE/e/MkImK30.pByqu0CGsQZW1", PrefixCryptSha512},
	{"andrew1", "654321", "Qro3QWOs61UCarx1PAwAlL1.vJgZJXsIXml3.3vVhV.2xUwIRBmmyQzK9yFAqYY5iD1wkAdhUko6hl6T9N7s5.", PrefixCryptSha512},
	{"dreadpirateroberts", "98765432101234567890", "FPU3HtJ9RcPVUvxifkIJ/AlrBxWLqJQvyxK2f8x4qDX/A1RpcIvgjToU5erVkR6XUl7qwPsm7idpbMH5f/pBn0", PrefixCryptSha512},
}

func Test_CryptSha(t *testing.T) {

	for _, v := range sha256TestData {
		text := fmt.Sprintf(v.prefix+"%s$%s", v.salt, v.hashed)
		testParserGood(t, "crypt-sha512", AcceptCryptSha, RejectCryptSha, text, v.password)
	}

	for _, v := range sha512TestData {
		text := fmt.Sprintf(v.prefix+"%s$%s", v.salt, v.hashed)
		testParserGood(t, "crypt-sha512", AcceptCryptSha, RejectCryptSha, text, v.password)
	}

	testParserBad(t, "crypt-sha256", AcceptCryptSha, RejectCryptSha, "$5$nosalt")
	testParserBad(t, "crypt-sha512", AcceptCryptSha, RejectCryptSha, "$6$nosalt")
	testParserNot(t, "crypt-sha512", AcceptCryptSha, RejectCryptSha, "plain")
	testParserNot(t, "crypt-sha512", AcceptCryptSha, RejectCryptSha, "{SHA}plain")
}
