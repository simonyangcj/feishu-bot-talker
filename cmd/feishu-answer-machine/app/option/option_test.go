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
		if _, err := os.Stat("/tmp/feishu-answer-machine-config.yaml"); os.IsNotExist(err) {
			file, err := os.Create("/tmp/feishu-answer-machine-config.yaml")
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
				cmd.SetArgs([]string{"--config-file", "/tmp/feishu-answer-machine-config.yaml"})
				cmd.Execute()
				Expect(option.ConfigFile).To(Equal("/tmp/feishu-answer-machine-config.yaml"))
			})
			It("should failed option validation due to config file set to empty", func() {
				cmd := &cobra.Command{
					Use:  "test",
					Long: "test",
				}
				fs := cmd.PersistentFlags()
				option.BindFlags(fs)
				cmd.SetArgs([]string{"--config-file", ""})
				cmd.Execute()
				err := option.ValidateOptions()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("config-file is required"))
			})
		})
	})
})
