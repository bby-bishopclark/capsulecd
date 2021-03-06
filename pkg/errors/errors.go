package errors

import (
	"fmt"
)

// Raised when there is an issue with the filesystem for scm checkout
type ScmFilesystemError string

func (str ScmFilesystemError) Error() string {
	return fmt.Sprintf("ScmFilesystemError: %q", string(str))
}

// Raised when the scm is not recognized
type ScmUnspecifiedError string

func (str ScmUnspecifiedError) Error() string {
	return fmt.Sprintf("ScmUnspecifiedError: %q", string(str))
}

// Raised when capsule cannot create an authenticated client for the source.
type ScmAuthenticationFailed string

func (str ScmAuthenticationFailed) Error() string {
	return fmt.Sprintf("ScmAuthenticationFailed: %q", string(str))
}

// Raised when there is an error parsing the repo payload format.
type ScmPayloadFormatError string

func (str ScmPayloadFormatError) Error() string {
	return fmt.Sprintf("ScmPayloadFormatError: %q", string(str))
}

// Raised when a source payload is unsupported/action is invalid
type ScmPayloadUnsupported string

func (str ScmPayloadUnsupported) Error() string {
	return fmt.Sprintf("ScmPayloadUnsupported: %q", string(str))
}

// Raised when the user who started the packaging is unauthorized (non-collaborator)
type ScmUnauthorizedUser string

func (str ScmUnauthorizedUser) Error() string {
	return fmt.Sprintf("ScmUnauthorizedUser: %q", string(str))
}

// Raised when the scm cleanup failed
type ScmCleanupFailed string

func (str ScmCleanupFailed) Error() string {
	return fmt.Sprintf("ScmCleanupFailed: %q", string(str))
}

// Raised when the scm pr already merged
type ScmMergeNothingToMergeError string

func (str ScmMergeNothingToMergeError) Error() string {
	return fmt.Sprintf("ScmMergeNothingToMergeError: %q", string(str))
}

// Raised during a PR merge when there is a merge conflict
type ScmMergeConflictError string
func (str ScmMergeConflictError) Error() string {
	return fmt.Sprintf("ScmMergeConflictError: %q", string(str))
}

// Raised during a PR merge where the analysis returns a result that we do not understand.
type ScmMergeAnalysisUnknownError string
func (str ScmMergeAnalysisUnknownError) Error() string {
	return fmt.Sprintf("ScmMergeAnalysisUnknownError: %q", string(str))
}

// Raised when the config file specifies a hook/override for a step when the type is :repo
type EngineTransformUnavailableStep string

func (str EngineTransformUnavailableStep) Error() string {
	return fmt.Sprintf("EngineTransformUnavailableStep: %q", string(str))
}

// Raised when the environment is missing a required tool/binary
type EngineValidateToolError string

func (str EngineValidateToolError) Error() string {
	return fmt.Sprintf("EngineValidateToolError: %q", string(str))
}

// Raised when the engine is not recognized
type EngineUnspecifiedError string

func (str EngineUnspecifiedError) Error() string {
	return fmt.Sprintf("EngineUnspecifiedError: %q", string(str))
}

// Raised when the package is missing certain required files (ie metadata.rb, package.json, setup.py, etc)
type EngineBuildPackageInvalid string

func (str EngineBuildPackageInvalid) Error() string {
	return fmt.Sprintf("EngineBuildPackageInvalid: %q", string(str))
}

// Raised when the source could not be compiled or build for any reason
type EngineBuildPackageFailed string

func (str EngineBuildPackageFailed) Error() string {
	return fmt.Sprintf("EngineBuildPackageFailed: %q", string(str))
}

// Raised when package dependencies fail to install correctly.
type EngineTestDependenciesError string

func (str EngineTestDependenciesError) Error() string {
	return fmt.Sprintf("EngineTestDependenciesError: %q", string(str))
}

// Raised when the package test runner fails
type EngineTestRunnerError string

func (str EngineTestRunnerError) Error() string {
	return fmt.Sprintf("EngineTestRunnerError: %q", string(str))
}

// Raised when package manager asseble step fails.
type MgrAssembleError string

func (str MgrAssembleError) Error() string {
	return fmt.Sprintf("MgrAssembleError: %q", string(str))
}

// Raised when credentials required to upload/deploy new package are missing.
type MgrDistCredentialsMissing string

func (str MgrDistCredentialsMissing) Error() string {
	return fmt.Sprintf("MgrDistCredentialsMissing: %q", string(str))
}

// Raised when an error occurs while uploading package.
type MgrDistPackageError string

func (str MgrDistPackageError) Error() string {
	return fmt.Sprintf("MgrDistPackageError: %q", string(str))
}
