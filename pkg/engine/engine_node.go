package engine

import (
	"capsulecd/pkg/config"
	"capsulecd/pkg/errors"
	"capsulecd/pkg/pipeline"
	"capsulecd/pkg/scm"
	"capsulecd/pkg/utils"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

type nodeMetadata struct {
	Version string
}
type engineNode struct {
	*EngineBase

	PipelineData *pipeline.Data
	Scm          scm.Scm //Interface
}

func (n *engineNode) ValidateTools() error {
	if _, kerr := exec.LookPath("npm"); kerr != nil {
		return errors.EngineValidateToolError("npm binary is missing")
	}

	return nil
}

func (n *engineNode) Init(pipelineData *pipeline.Data, sourceScm scm.Scm) error {
	n.Scm = sourceScm
	n.PipelineData = pipelineData
	return nil
}

func (n *engineNode) BuildStep() error {
	//validate that the npm package.json file exists
	if !utils.FileExists(path.Join(n.PipelineData.GitLocalPath, "package.json")) {
		return errors.EngineBuildPackageInvalid("package.json file is required to process Node package")
	}

	// no need to bump up the version here. It will automatically be bumped up via the npm version patch command.
	// however we need to read the version from the package.json file and check if a npm module already exists.

	//TODO: check if this module name and version already exist.

	// check for/create any required missing folders/files
	if derr := os.MkdirAll(path.Join(n.PipelineData.GitLocalPath, "test"), 0644); derr != nil {
		return derr
	}

	//TODO: add gitignore content.
	//if !utils.FileExists(path.Join(g.PipelineData.GitLocalPath, ".gitignore")) {
	//	ioutil.WriteFile(path.Join(g.PipelineData.GitLocalPath, ".gitignore"),
	//		[]byte(""),
	//		0644,
	//	)
	//}
	return nil
}

func (n *engineNode) TestStep() error {

	// the module has already been downloaded. lets make sure all its dependencies are available.
	if derr := utils.BashCmdExec("npm install", n.PipelineData.GitLocalPath, ""); derr != nil {
		return errors.EngineTestRunnerError("npm install failed. Check module dependencies")
	}

	// create a shrinkwrap file.
	if derr := utils.BashCmdExec("npm shrinkwrap", n.PipelineData.GitLocalPath, ""); derr != nil {
		return errors.EngineTestRunnerError("npm shrinkwrap failed. Check log for exact error")
	}

	//run test command
	var testCmd string
	if config.IsSet("engine_cmd_test") {
		testCmd = config.GetString("engine_cmd_test")
	} else {
		testCmd = "npm test"
	}
	//running tox will install all dependencies in a virtual env, and then run unit tests.
	if derr := utils.BashCmdExec(testCmd, n.PipelineData.GitLocalPath, ""); derr != nil {
		return errors.EngineTestRunnerError(fmt.Sprintf("Test command (%s) failed. Check log for more details.", testCmd))
	}
	return nil
}

func (n *engineNode) PackageStep() error {
	// commit changes to the cookbook. (test run occurs before this, and it should clean up any instrumentation files, created,
	// as they will be included in the commmit and any release artifacts)
	if cerr := utils.GitCommit(n.PipelineData.GitLocalPath, "Committing automated changes before packaging."); cerr != nil {
		return cerr
	}

	// run npm publish
	versionCmd := fmt.Sprintf("npm version %s -m '(v%%s) Automated packaging of release by CapsuleCD'",
		config.GetString("engine_version_bump_type"),
	)
	if verr := utils.BashCmdExec(versionCmd, n.PipelineData.GitLocalPath, ""); verr != nil {
		return errors.EngineTestRunnerError("npm version bump failed")
	}

	tagCommit, terr := utils.GitLatestTaggedCommit(n.PipelineData.GitLocalPath)
	if terr != nil {
		return terr
	}

	n.PipelineData.ReleaseCommit = tagCommit.CommitSha
	n.PipelineData.ReleaseVersion = tagCommit.TagShortName
	return nil
}

func (n *engineNode) DistStep() error {
	if !config.IsSet("npm_auth_token") {
		return errors.EngineDistCredentialsMissing("cannot deploy page to npm, credentials missing")
	}

	npmrcFile, _ := ioutil.TempFile("", ".npmrc")
	defer os.Remove(npmrcFile.Name())

	// write the .npmrc config jfile.
	npmrcContent := fmt.Sprintf(
		"//registry.npmjs.org/:_authToken=%s",
		config.GetString("npm_auth_token"),
	)

	if _, werr := npmrcFile.Write([]byte(npmrcContent)); werr != nil {
		return werr
	}

	npmPublishCmd := "npm publish ."
	derr := utils.BashCmdExec(npmPublishCmd, n.PipelineData.GitLocalPath, "")
	if derr != nil {
		return errors.EngineDistPackageError("npm publish failed. Check log for exact error")
	}
	return nil
}