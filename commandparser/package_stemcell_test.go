package commandparser_test

import (
	"flag"

	. "github.com/cloudfoundry-incubator/stembuild/commandparser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("package_stemcell", func() {
	// Focus of this test is not to test the Flags.Parse functionality as much
	// as to test that the command line flags values are stored in the expected
	// struct variables. This adds a bit of protection when renaming flag parameters.
	Describe("SetFlags", func() {

		var f *flag.FlagSet
		var PkgCmd *PackageCmd

		BeforeEach(func() {
			f = flag.NewFlagSet("test", flag.ExitOnError)
			PkgCmd = &PackageCmd{}
			PkgCmd.SetFlags(f)
		})

		var longformArgs = []string{"-vmdk", "some_vmdk_file",
			"-outputDir", "some_output_dir",
		}
		var shortformArgs = []string{"-vmdk", "some_vmdk_file",
			"-o", "some_output_dir",
		}

		Context("a vmdk file is specified as a flag parameter", func() {
			It("then the vmdk file name is stored", func() {
				err := f.Parse(longformArgs)
				Expect(err).ToNot(HaveOccurred())
				Expect(PkgCmd.GetVMDK()).To(Equal("some_vmdk_file"))
			})
		})

		Context("an output directory is specified as a flag parameter", func() {
			It("when using the long form the directory is stored", func() {
				err := f.Parse(longformArgs)
				Expect(err).ToNot(HaveOccurred())
				Expect(PkgCmd.GetOutputDir()).To(Equal("some_output_dir"))
			})

			It("when using the short form the directory is stored", func() {
				err := f.Parse(shortformArgs)
				Expect(err).ToNot(HaveOccurred())
				Expect(PkgCmd.GetOutputDir()).To(Equal("some_output_dir"))
			})

		})

	})

})
