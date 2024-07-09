package option

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

var _ = Describe("option test", func() {
	var option *Option
	BeforeEach(func() {
		option = &Option{}
		// create a tmp file in /tmp/fake-content.txt if not exists
		if _, err := os.Stat("/tmp/fake-content.txt"); os.IsNotExist(err) {
			file, err := os.Create("/tmp/fake-content.txt")
			if err != nil {
				panic(err)
			}
			defer file.Close()
		}
	})

	Describe("binding test", func() {
		Context("when binding flags", func() {
			It("should bind all flags", func() {
				cmd := &cobra.Command{
					Use:  "test",
					Long: "test",
				}
				fs := cmd.PersistentFlags()
				option.BindFlags(fs)
				cmd.SetArgs([]string{"--app-id", "test-id", "--app-secret", "test-secret", "--receive-id-type", "chat_id", "--message-type", "text", "--content-file", "/tmp/fake-content.txt"})
				cmd.Execute()
				Expect(option.AppID).To(Equal("test-id"))
				Expect(option.AppSecret).To(Equal("test-secret"))
				Expect(option.ReceiveIdType).To(Equal("chat_id"))
			})
			It("should failed option validation due to no content provider", func() {
				cmd := &cobra.Command{
					Use:  "test",
					Long: "test",
				}
				fs := cmd.PersistentFlags()
				option.BindFlags(fs)
				cmd.SetArgs([]string{"--app-id", "test-id", "--app-secret", "test-secret", "--receive-id-type", "chat_id", "--receive-id-value", "test-value", "--message-type", "text"})
				cmd.Execute()
				err := option.ValidateOptions()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("content-file is required"))
			})
			It("should pass option validation", func() {
				cmd := &cobra.Command{
					Use:  "test",
					Long: "test",
				}
				fs := cmd.PersistentFlags()
				option.BindFlags(fs)
				cmd.SetArgs([]string{"--app-id", "test-id", "--app-secret", "test-secret", "--receive-id-type", "chat_id", "--message-type", "text", "--content-file", "/tmp/fake-content.txt", "--receive-id-value", "test-value"})
				cmd.Execute()
				err := option.ValidateOptions()
				Expect(err).To(BeNil())
			})
			It("should failed to validate option with ReceiveIdType", func() {
				cmd := &cobra.Command{
					Use:  "test",
					Long: "test",
				}
				fs := cmd.PersistentFlags()
				option.BindFlags(fs)
				cmd.SetArgs([]string{"--app-id", "test-id", "--app-secret", "test-secret", "--receive-id-type", "test_id", "--message-type", "text", "--content-file", "test.txt", "--receive-id-value", "test-value"})
				cmd.Execute()
				err := option.ValidateOptions()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("receive-id-type test_id is not supported"))
			})
			It("should failed to due to content file not found", func() {
				cmd := &cobra.Command{
					Use:  "test",
					Long: "test",
				}
				fs := cmd.PersistentFlags()
				option.BindFlags(fs)
				cmd.SetArgs([]string{"--app-id", "test-id", "--app-secret", "test-secret", "--receive-id-type", "chat_id", "--message-type", "text", "--content-file", "test.txt", "--receive-id-value", "test-value"})
				cmd.Execute()
				err := option.ValidateOptions()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("content-file test.txt does not exist"))
			})
			It("should failed to validate option with MessageType", func() {
				cmd := &cobra.Command{
					Use:  "test",
					Long: "test",
				}
				fs := cmd.PersistentFlags()
				option.BindFlags(fs)
				cmd.SetArgs([]string{"--app-id", "test-id", "--app-secret", "test-secret", "--receive-id-type", "chat_id", "--message-type", "image", "--content-file", "/tmp/fake-content.txt", "--receive-id-value", "test-value"})
				cmd.Execute()
				err := option.ValidateOptions()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("message-type image is not supported"))
			})
			It("should not bind any flags when not provided", func() {
				cmd := &cobra.Command{
					Use:  "test",
					Long: "test",
				}
				fs := cmd.PersistentFlags()
				option.BindFlags(fs)
				cmd.SetArgs([]string{})
				cmd.Execute()
				Expect(option.AppID).To(Equal(""))
				Expect(option.AppSecret).To(Equal(""))
				Expect(option.ReceiveIdType).To(Equal("open_id"))
				Expect(option.MessageType).To(Equal("text"))
				Expect(option.ContentFile).To(Equal(""))
			})
		})
	})
})
