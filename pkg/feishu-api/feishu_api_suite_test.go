package feishu_api_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFeishuApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FeishuApi Suite")
}
