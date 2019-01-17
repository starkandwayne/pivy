package metadata_test

import (
	"io/ioutil"
	"path/filepath"
	"runtime"

	. "github.com/starkandwayne/pivy/metadata"
	. "github.com/starkandwayne/pivy/string_diff_matcher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

var _ = Describe("GolangRenderer", func() {
	//	format.TruncatedDiff = true
	var (
		out string
	)
	// standardizeSpaces: remove redudant spaces for more context in diff
	standardizeSpaces := func(s string) string {
		return s
		// return strings.Join(strings.Fields(s), " ")
	}
	parse := func(m string) {
		parser, err := NewParser([]byte(m))
		Expect(err).ToNot(HaveOccurred())
		out, err = parser.RenderToGolang()
		Expect(err).ToNot(HaveOccurred())
		out = standardizeSpaces(out)
	}
	readAsset := func(file string) string {
		_, filename, _, _ := runtime.Caller(0)
		assetdir := filepath.Join(filepath.Dir(filename), "assets")
		data, err := ioutil.ReadFile(filepath.Join(assetdir, file))
		Expect(err).ToNot(HaveOccurred())

		return string(data)
	}

	Describe("given metadata with property_blueprints", func() {
		format.TruncatedDiff = false

		It("renders a go file with property struct", func() {
			parse(readAsset("property_blueprints.yml"))
			Expect(out).To(EqualWithDiff(standardizeSpaces(readAsset("property_blueprints.go"))))
		})
	})
})