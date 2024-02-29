package distributor

type APKTester struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type APKRelease struct {
	FilePath     string       `json:"file_path"`
	ReleaseNotes string       `json:"release_notes"`
	APKTesters   []*APKTester `json:"apk_testers"`
}

type Distributor interface {
	Distribute(apkRelease *APKRelease) (string, error)
}
