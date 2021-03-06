package scm

import (
	"capsulecd/pkg/config"
	"capsulecd/pkg/pipeline"
	"net/http"
)

// Create mock using:
// mockgen -source=pkg/scm/interface.go -destination=pkg/scm/mock/mock_scm.go
type Interface interface {

	// init method will generate an authenticated client that can be used to comunicate with Scm
	// MUST set pipelineData.GitParentPath
	Init(pipelineData *pipeline.Data, config config.Interface, client *http.Client) error

	// Determine if this is a pull request or a push.
	// if it's a pull request the scm must retrieve the pull request payload and return it
	// if its a push, the scm must retrieve the push payload and return it
	// CAN NOT override
	// MUST set pipelineData.IsPullRequest
	// RETURNS scm.Payload
	RetrievePayload() (*Payload, error)

	// start processing the payload, which should result in a local git repository that we
	// can begin to test. Since this is a push, no packaging is required
	// CAN NOT override
	// MUST set pipelineData.GitLocalPath
	// MUST set pipelineData.GitLocalBranch
	// MUST set pipelienData.GitRemote
	// MUST set pipelineData.GitHeadInfo
	// SHOULD set pipelineData.NearestTagDetails
	// REQUIRES pipelineData.GitParentPath
	CheckoutPushPayload(payload *Payload) error

	// all capsule CD processing will be kicked off via a payload. In Github's case, the payload is the pull request data.
	// should check if the pull request opener even has permissions to create a release.
	// all sources should process the payload by downloading a git repository that contains the master branch merged with the test branch
	// CAN NOT override
	// MUST set pipelineData.GitLocalPath
	// MUST set pipelineData.GitLocalBranch
	// MUST set pipelienData.GitRemote
	// MUST set pipelineData.GitBaseInfo
	// MUST set pipelineData.GitHeadInfo
	// SHOULD set pipelineData.NearestTagDetails
	// REQUIRES pipelineData.GitParentPath
	CheckoutPullRequestPayload(payload *Payload) error

	// The repository should now contain code that has been the merged, tested and version bumped.
	// This method will push these changes to the source code repository
	// this step should also do any scm specific releases (github release, asset uploading, etc)
	// CAN override
	// REQUIRES config.scm_repo_full_name
	// REQUIRES pipelineData.ScmReleaseCommit
	// REQUIRES pipelineData.GitLocalPath
	// REQUIRES pipelineData.GitLocalBranch
	// REQUIRES pipelineData.GitBaseInfo
	// REQUIRES pipelineData.GitHeadInfo
	// REQUIRES pipelineData.ReleaseArtifacts
	// REQUIRES pipelineData.ReleaseVersion
	// REQUIRES pipelineData.ReleaseCommit
	// REQUIRES pipelineData.GitParentPath
	// USES set pipelineData.NearestTagDetails
	Publish() error //create release.

	//Upload assets to SCM, and attach to SCM release if possible.
	//Failing to upload Assets to SCM will not fail the publish (we'll retry 5 times)
	//Should not be called directly, will be called via Publish()
	//ReleaseData will be different for each SCM, but is probably a release ID that we can attach files to.
	//REQUIRES config.scm_repo_full_name
	//REQUIRES pipelineData.ReleaseAssets
	//REQUIRES pipelineData.GitLocalPath
	PublishAssets(releaseData interface{}) error

	// optionally delete the PR branch after the code has been merged into master.
	// only do so if:
	// - "scm_enable_branch_cleanup" is true
	// - HEAD PR branch is in the same repository as the BASE
	// - branch is not the default branch or "master" for this repository
	// - branch is not protected (SCM specific feature)
	// CAN override
	// USES scm_enable_branch_cleanup
	// REQUIRES config.scm_repo_full_name
	// REQUIRES pipelineData.GitBaseInfo.Repo.FullName
	// REQUIRES pipelineData.GitHeadInfo.Repo.FullName
	// REQUIRES pipelineData.GitHeadInfo.Ref
	Cleanup() error

	// Notify should update the scm with the build status at each stage.
	// If the scm does not support notifications this should be a no-op
	// In general, if the Notify method returns an error, we'll ignore it, and continue the pipeline.
	// REQUIRES config.scm_repo_full_name
	Notify(ref string, state string, message string) error
}
