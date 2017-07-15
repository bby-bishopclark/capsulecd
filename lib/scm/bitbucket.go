package scm

import (
"github.com/google/go-github/github"
"context"
"golang.org/x/oauth2"
)

type scmBitbucket struct {
	client *github.Client
	options *ScmOptions
}

// configure method will generate an authenticated client that can be used to comunicate with Github
// MUST set @git_parent_path
// MUST set @client field
func (b *scmBitbucket) Configure() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "... your access token ..."},
	)
	tc := oauth2.NewClient(ctx, ts)

	b.client = github.NewClient(tc)
	return
}

func (b *scmBitbucket) RetrievePayload() *ScmPayload {
	return &ScmPayload{}
}

func (b *scmBitbucket) ProcessPushPayload() {
	return
}

func (b *scmBitbucket) ProcessPullRequestPayload() {
	return
}

func (b *scmBitbucket) Publish() {
	return
}

func (b *scmBitbucket) Notify() {
	return
}

func (b *scmBitbucket) Options() *ScmOptions {
	return b.options
}