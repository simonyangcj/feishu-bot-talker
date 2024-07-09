package feishu_api

import (
	"encoding/json"
	"feishu-bot-talker/cmd/feishu-bot-talker/app/option"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("option test", func() {
	var optionTest *option.Option
	BeforeEach(func() {
		optionTest = &option.Option{
			AppID:         "test-id",
			AppSecret:     "test-secret",
			ReceiveIdType: "chat_id",
			MessageType:   "text",
			ContentFile:   "/tmp/fake-content1.txt",
		}
		// create a tmp file in /tmp/fake-content.txt if not exists
		if _, err := os.Stat("/tmp/fake-content1.txt"); os.IsNotExist(err) {
			file, err := os.Create("/tmp/fake-content1.txt")
			if err != nil {
				panic(err)
			}
			defer file.Close()
		}
	})

	AfterEach(func() {
		err := os.Remove("/tmp/fake-content1.txt")
		Expect(err).To(BeNil())
	})

	Describe("render test", func() {
		Context("MessageType text test", func() {
			It("should render text", func() {
				text := &ContentRenderText{
					Text: "hi",
				}
				result, err := json.Marshal(text)
				Expect(err).To(BeNil())
				// write to file /tmp/fake-content1.txt
				err = os.WriteFile("/tmp/fake-content1.txt", result, 0644)
				Expect(err).To(BeNil())
				caller, err := CreateFeiShuDataPost(optionTest)
				Expect(err).To(BeNil())
				Expect(caller.Context).To(Equal(`{"text":"hi"}`))
			})
			It("should failed to render due to MessageType not proper set ", func() {
				optionTest.MessageType = "image"
				text := &ContentRenderText{
					Text: "hi",
				}
				result, err := json.Marshal(text)
				Expect(err).To(BeNil())
				// write to file /tmp/fake-content1.txt
				err = os.WriteFile("/tmp/fake-content1.txt", result, 0644)
				Expect(err).To(BeNil())
				_, err = CreateFeiShuDataPost(optionTest)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("message-type image is not supported"))
			})
			It("should faile to render due to content unable to parse", func() {
				err := os.WriteFile("/tmp/fake-content1.txt", []byte("hi"), 0644)
				Expect(err).To(BeNil())
				_, err = CreateFeiShuDataPost(optionTest)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("token test", func() {
		Context("Call api get token", func() {
			It("should be get token as expected", func() {
				optionTest.AppID = "cli_a602c1f9f93a500e"
				optionTest.AppSecret = "kZgjf3frn3fEQ9Zsk6Uh0eEhoSPDZZsE"
				_, err := GetTenantAccessToken(optionTest)
				Expect(err).To(BeNil())
			})
			It("should be send message as expected", func() {
				text := &ContentRenderText{
					Text: " hi",
				}
				result, err := json.Marshal(text)
				Expect(err).To(BeNil())
				err = os.WriteFile("/tmp/fake-content1.txt", result, 0644)
				Expect(err).To(BeNil())
				optionTest.AppID = "cli_a602c1f9f93a500e"
				optionTest.AppSecret = "kZgjf3frn3fEQ9Zsk6Uh0eEhoSPDZZsE"
				optionTest.ReceiveIdValue = "oc_ed741e22603db268f855388c3cfc7f59"
				caller, err := CreateFeiShuDataPost(optionTest)
				Expect(err).To(BeNil())
				err = SendMessage(optionTest, caller)
				Expect(err).To(BeNil())
			})
		})
	})
})
