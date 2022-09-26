package password_test

import (
	"github.com/nelsonlai-go/password"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hash Password", func() {
	It("should hash the password", func() {
		pwd := "password"

		h := password.Hash(pwd)
		Expect(h).ToNot(BeEmpty())

		ok := password.Verify(pwd, h)
		Expect(ok).To(BeTrue())
	})

	It("should not verify the wrong password", func() {
		pwd := "password"

		h := password.Hash(pwd)
		Expect(h).ToNot(BeEmpty())

		ok := password.Verify("password2", h)
		Expect(ok).To(BeFalse())
	})
})
