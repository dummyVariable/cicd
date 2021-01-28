package cicd

//Job defines a structure of job instuctions
type Job struct {
	PreBuild  []string `yaml:"prebuild"`
	Build     []string `yaml:"build"`
	PostBuild []string `yaml:"postbuild"`
}
