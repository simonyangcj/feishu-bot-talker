package config

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("config test", func() {
	var config *Config
	BeforeEach(func() {
		config = &Config{}
	})

	Describe("validation test", func() {
		Context("when validation config", func() {
			It("should failed config validation due to app_id set to empty", func() {
				err := config.validationConfig()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("app_id is required"))
			})
			It("should failed config validation due to app_secret set to empty", func() {
				config.AppID = "test-id"
				err := config.validationConfig()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("app_secret is required"))
			})
			It("should failed config validation due to encrypt_key set to empty", func() {
				config.AppID = "test-id"
				config.AppSecret = "test-secret"
				err := config.validationConfig()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("encrypt_key is required"))
			})
			It("should failed config validation due to verification_token set to empty", func() {
				config.AppID = "test-id"
				config.AppSecret = "test-secret"
				config.EncryptKey = "test-encrypt-key"
				err := config.validationConfig()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("verification_token is required"))
			})
			It("should pass config validation", func() {
				config.AppID = "test-id"
				config.AppSecret = "test-secret"
				config.EncryptKey = "test-encrypt-key"
				config.VerificationToken = "test-verification-token"
				err := config.validationConfig()
				Expect(err).To(BeNil())
			})
		})
	})
})
