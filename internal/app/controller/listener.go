// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/clivern/hamster/internal/app/event"
	"github.com/clivern/hamster/internal/app/listener"
	"github.com/clivern/hamster/internal/app/pkg/logger"
	"github.com/clivern/hamster/plugin"
	"github.com/gin-gonic/gin"
	"os"
)

// Listen controller
func Listen(c *gin.Context) {
	var actions listener.Action
	var commands listener.Commands

	rawBody, _ := c.GetRawData()
	body := string(rawBody)

	parser := &listener.Parser{
		UserAgent:      c.GetHeader("User-Agent"),
		GithubDelivery: c.GetHeader("X-GitHub-Delivery"),
		GitHubEvent:    c.GetHeader("X-GitHub-Event"),
		HubSignature:   c.GetHeader("X-Hub-Signature"),
		Body:           body,
	}

	ok := parser.VerifySignature(os.Getenv("GithubWebhookSecret"))
	evt := parser.GetGitHubEvent()

	logger.Infof("Incoming event %s with payload %s!", evt, body)

	if ok {
		switch evt {
		case "status":
			var status event.Status
			status.LoadFromJSON(rawBody)
			actions.RegisterStatusAction(plugin.StatusListener)
			actions.ExecuteStatusActions(status)
		case "watch":
			var watch event.Watch
			watch.LoadFromJSON(rawBody)
			actions.RegisterWatchAction(plugin.WatchListener)
			actions.ExecuteWatchActions(watch)
		case "issues":
			var issues event.Issues
			issues.LoadFromJSON(rawBody)
			actions.RegisterIssuesAction(plugin.IssuesListener)
			actions.ExecuteIssuesActions(issues)

			// Commands Listeners
			commands.RegisterIssuesAction("test", plugin.IssuesTestCommandListener)
			commands.ExecuteIssuesActions(issues)
		case "push":
			var push event.Push
			push.LoadFromJSON(rawBody)
			actions.RegisterPushAction(plugin.PushListener)
			actions.ExecutePushActions(push)
		case "issue_comment":
			var issueComment event.IssueComment
			issueComment.LoadFromJSON(rawBody)
			actions.RegisterIssueCommentAction(plugin.IssueCommentListener)
			actions.ExecuteIssueCommentActions(issueComment)

			// Commands Listeners
			commands.RegisterIssueCommentAction("test", plugin.IssueCommentTestCommandListener)
			commands.ExecuteIssueCommentActions(issueComment)
		case "create":
			var create event.Create
			create.LoadFromJSON(rawBody)
			actions.RegisterCreateAction(plugin.CreateListener)
			actions.ExecuteCreateActions(create)
		case "label":
			var label event.Label
			label.LoadFromJSON(rawBody)
			actions.RegisterLabelAction(plugin.LabelListener)
			actions.ExecuteLabelActions(label)
		case "delete":
			var delete event.Delete
			delete.LoadFromJSON(rawBody)
			actions.RegisterDeleteAction(plugin.DeleteListener)
			actions.ExecuteDeleteActions(delete)
		case "milestone":
			var milestone event.Milestone
			milestone.LoadFromJSON(rawBody)
			actions.RegisterMilestoneAction(plugin.MilestoneListener)
			actions.ExecuteMilestoneActions(milestone)
		case "pull_request":
			var pullRequest event.PullRequest
			pullRequest.LoadFromJSON(rawBody)
			actions.RegisterPullRequestAction(plugin.PullRequestListener)
			actions.ExecutePullRequestActions(pullRequest)
		case "pull_request_review":
			var pullRequestReview event.PullRequestReview
			pullRequestReview.LoadFromJSON(rawBody)
			actions.RegisterPullRequestReviewAction(plugin.PullRequestReviewListener)
			actions.ExecutePullRequestReviewActions(pullRequestReview)
		case "pull_request_review_comment":
			var pullRequestReviewComment event.PullRequestReviewComment
			pullRequestReviewComment.LoadFromJSON(rawBody)
			actions.RegisterPullRequestReviewCommentAction(plugin.PullRequestReviewCommentListener)
			actions.ExecutePullRequestReviewCommentActions(pullRequestReviewComment)
		default:
			logger.Infof("Unknown or unsupported event %s!", evt)
		}

		var raw event.Raw
		raw.SetEvent(evt)
		raw.SetBody(body)
		actions.RegisterRawAction(plugin.RawListener)
		actions.ExecuteRawActions(raw)

		c.JSON(200, gin.H{
			"status": "Nice!",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "Oops!",
		})
	}
}
