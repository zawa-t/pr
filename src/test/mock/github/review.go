// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/zawa-t/pr-commentator/src/platform/github"
	"sync"
)

// Ensure, that ReviewMock does implement github.Review.
// If this is not the case, regenerate this file with moq.
var _ github.Review = &ReviewMock{}

// ReviewMock is a mock implementation of github.Review.
//
//	func TestSomethingThatUsesReview(t *testing.T) {
//
//		// make and configure a mocked github.Review
//		mockedReview := &ReviewMock{
//			CreateCommentFunc: func(ctx context.Context, data github.CommentData) error {
//				panic("mock out the CreateComment method")
//			},
//		}
//
//		// use mockedReview in code that requires github.Review
//		// and then make assertions.
//
//	}
type ReviewMock struct {
	// CreateCommentFunc mocks the CreateComment method.
	CreateCommentFunc func(ctx context.Context, data github.CommentData) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateComment holds details about calls to the CreateComment method.
		CreateComment []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Data is the data argument value.
			Data github.CommentData
		}
	}
	lockCreateComment sync.RWMutex
}

// CreateComment calls CreateCommentFunc.
func (mock *ReviewMock) CreateComment(ctx context.Context, data github.CommentData) error {
	if mock.CreateCommentFunc == nil {
		panic("ReviewMock.CreateCommentFunc: method is nil but Review.CreateComment was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Data github.CommentData
	}{
		Ctx:  ctx,
		Data: data,
	}
	mock.lockCreateComment.Lock()
	mock.calls.CreateComment = append(mock.calls.CreateComment, callInfo)
	mock.lockCreateComment.Unlock()
	return mock.CreateCommentFunc(ctx, data)
}

// CreateCommentCalls gets all the calls that were made to CreateComment.
// Check the length with:
//
//	len(mockedReview.CreateCommentCalls())
func (mock *ReviewMock) CreateCommentCalls() []struct {
	Ctx  context.Context
	Data github.CommentData
} {
	var calls []struct {
		Ctx  context.Context
		Data github.CommentData
	}
	mock.lockCreateComment.RLock()
	calls = mock.calls.CreateComment
	mock.lockCreateComment.RUnlock()
	return calls
}