package package_parameters

type VmdkPackageParameters struct {
	OSVersion string `yaml:"os_version"`
	OutputDir string `yaml:"output_dir"`
	Version   string `yaml:"version"`
	VMDKFile  string `yaml:"vmdk_file"`
}

// Copy into `d` the values in `s` which are empty in `d`.
func (d *VmdkPackageParameters) CopyFrom(s VmdkPackageParameters) {
	if d.OSVersion == "" {
		d.OSVersion = s.OSVersion
	}

	// ignore OutputDir from config file

	if d.Version == "" {
		d.Version = s.Version
	}

	if d.VMDKFile == "" {
		d.VMDKFile = s.VMDKFile
	}
}
