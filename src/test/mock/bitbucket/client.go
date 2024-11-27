// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/zawa-t/pr/src/platform/bitbucket"
	"sync"
)

// Ensure, that ClientMock does implement bitbucket.Client.
// If this is not the case, regenerate this file with moq.
var _ bitbucket.Client = &ClientMock{}

// ClientMock is a mock implementation of bitbucket.Client.
//
//	func TestSomethingThatUsesClient(t *testing.T) {
//
//		// make and configure a mocked bitbucket.Client
//		mockedClient := &ClientMock{
//			BulkUpsertAnnotationsFunc: func(ctx context.Context, datas []bitbucket.AnnotationData, reportID string) error {
//				panic("mock out the BulkUpsertAnnotations method")
//			},
//			DeleteReportFunc: func(ctx context.Context, reportID string) error {
//				panic("mock out the DeleteReport method")
//			},
//			GetCommentsFunc: func(ctx context.Context) ([]bitbucket.Comment, error) {
//				panic("mock out the GetComments method")
//			},
//			GetReportFunc: func(ctx context.Context, reportID string) (*bitbucket.AnnotationResponse, error) {
//				panic("mock out the GetReport method")
//			},
//			PostCommentFunc: func(ctx context.Context, data bitbucket.CommentData) error {
//				panic("mock out the PostComment method")
//			},
//			UpsertReportFunc: func(ctx context.Context, reportID string, data bitbucket.ReportData) error {
//				panic("mock out the UpsertReport method")
//			},
//		}
//
//		// use mockedClient in code that requires bitbucket.Client
//		// and then make assertions.
//
//	}
type ClientMock struct {
	// BulkUpsertAnnotationsFunc mocks the BulkUpsertAnnotations method.
	BulkUpsertAnnotationsFunc func(ctx context.Context, datas []bitbucket.AnnotationData, reportID string) error

	// DeleteReportFunc mocks the DeleteReport method.
	DeleteReportFunc func(ctx context.Context, reportID string) error

	// GetCommentsFunc mocks the GetComments method.
	GetCommentsFunc func(ctx context.Context) ([]bitbucket.Comment, error)

	// GetReportFunc mocks the GetReport method.
	GetReportFunc func(ctx context.Context, reportID string) (*bitbucket.AnnotationResponse, error)

	// PostCommentFunc mocks the PostComment method.
	PostCommentFunc func(ctx context.Context, data bitbucket.CommentData) error

	// UpsertReportFunc mocks the UpsertReport method.
	UpsertReportFunc func(ctx context.Context, reportID string, data bitbucket.ReportData) error

	// calls tracks calls to the methods.
	calls struct {
		// BulkUpsertAnnotations holds details about calls to the BulkUpsertAnnotations method.
		BulkUpsertAnnotations []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Datas is the datas argument value.
			Datas []bitbucket.AnnotationData
			// ReportID is the reportID argument value.
			ReportID string
		}
		// DeleteReport holds details about calls to the DeleteReport method.
		DeleteReport []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ReportID is the reportID argument value.
			ReportID string
		}
		// GetComments holds details about calls to the GetComments method.
		GetComments []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetReport holds details about calls to the GetReport method.
		GetReport []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ReportID is the reportID argument value.
			ReportID string
		}
		// PostComment holds details about calls to the PostComment method.
		PostComment []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Data is the data argument value.
			Data bitbucket.CommentData
		}
		// UpsertReport holds details about calls to the UpsertReport method.
		UpsertReport []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ReportID is the reportID argument value.
			ReportID string
			// Data is the data argument value.
			Data bitbucket.ReportData
		}
	}
	lockBulkUpsertAnnotations sync.RWMutex
	lockDeleteReport          sync.RWMutex
	lockGetComments           sync.RWMutex
	lockGetReport             sync.RWMutex
	lockPostComment           sync.RWMutex
	lockUpsertReport          sync.RWMutex
}

// BulkUpsertAnnotations calls BulkUpsertAnnotationsFunc.
func (mock *ClientMock) BulkUpsertAnnotations(ctx context.Context, datas []bitbucket.AnnotationData, reportID string) error {
	if mock.BulkUpsertAnnotationsFunc == nil {
		panic("ClientMock.BulkUpsertAnnotationsFunc: method is nil but Client.BulkUpsertAnnotations was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Datas    []bitbucket.AnnotationData
		ReportID string
	}{
		Ctx:      ctx,
		Datas:    datas,
		ReportID: reportID,
	}
	mock.lockBulkUpsertAnnotations.Lock()
	mock.calls.BulkUpsertAnnotations = append(mock.calls.BulkUpsertAnnotations, callInfo)
	mock.lockBulkUpsertAnnotations.Unlock()
	return mock.BulkUpsertAnnotationsFunc(ctx, datas, reportID)
}

// BulkUpsertAnnotationsCalls gets all the calls that were made to BulkUpsertAnnotations.
// Check the length with:
//
//	len(mockedClient.BulkUpsertAnnotationsCalls())
func (mock *ClientMock) BulkUpsertAnnotationsCalls() []struct {
	Ctx      context.Context
	Datas    []bitbucket.AnnotationData
	ReportID string
} {
	var calls []struct {
		Ctx      context.Context
		Datas    []bitbucket.AnnotationData
		ReportID string
	}
	mock.lockBulkUpsertAnnotations.RLock()
	calls = mock.calls.BulkUpsertAnnotations
	mock.lockBulkUpsertAnnotations.RUnlock()
	return calls
}

// DeleteReport calls DeleteReportFunc.
func (mock *ClientMock) DeleteReport(ctx context.Context, reportID string) error {
	if mock.DeleteReportFunc == nil {
		panic("ClientMock.DeleteReportFunc: method is nil but Client.DeleteReport was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		ReportID string
	}{
		Ctx:      ctx,
		ReportID: reportID,
	}
	mock.lockDeleteReport.Lock()
	mock.calls.DeleteReport = append(mock.calls.DeleteReport, callInfo)
	mock.lockDeleteReport.Unlock()
	return mock.DeleteReportFunc(ctx, reportID)
}

// DeleteReportCalls gets all the calls that were made to DeleteReport.
// Check the length with:
//
//	len(mockedClient.DeleteReportCalls())
func (mock *ClientMock) DeleteReportCalls() []struct {
	Ctx      context.Context
	ReportID string
} {
	var calls []struct {
		Ctx      context.Context
		ReportID string
	}
	mock.lockDeleteReport.RLock()
	calls = mock.calls.DeleteReport
	mock.lockDeleteReport.RUnlock()
	return calls
}

// GetComments calls GetCommentsFunc.
func (mock *ClientMock) GetComments(ctx context.Context) ([]bitbucket.Comment, error) {
	if mock.GetCommentsFunc == nil {
		panic("ClientMock.GetCommentsFunc: method is nil but Client.GetComments was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetComments.Lock()
	mock.calls.GetComments = append(mock.calls.GetComments, callInfo)
	mock.lockGetComments.Unlock()
	return mock.GetCommentsFunc(ctx)
}

// GetCommentsCalls gets all the calls that were made to GetComments.
// Check the length with:
//
//	len(mockedClient.GetCommentsCalls())
func (mock *ClientMock) GetCommentsCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetComments.RLock()
	calls = mock.calls.GetComments
	mock.lockGetComments.RUnlock()
	return calls
}

// GetReport calls GetReportFunc.
func (mock *ClientMock) GetReport(ctx context.Context, reportID string) (*bitbucket.AnnotationResponse, error) {
	if mock.GetReportFunc == nil {
		panic("ClientMock.GetReportFunc: method is nil but Client.GetReport was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		ReportID string
	}{
		Ctx:      ctx,
		ReportID: reportID,
	}
	mock.lockGetReport.Lock()
	mock.calls.GetReport = append(mock.calls.GetReport, callInfo)
	mock.lockGetReport.Unlock()
	return mock.GetReportFunc(ctx, reportID)
}

// GetReportCalls gets all the calls that were made to GetReport.
// Check the length with:
//
//	len(mockedClient.GetReportCalls())
func (mock *ClientMock) GetReportCalls() []struct {
	Ctx      context.Context
	ReportID string
} {
	var calls []struct {
		Ctx      context.Context
		ReportID string
	}
	mock.lockGetReport.RLock()
	calls = mock.calls.GetReport
	mock.lockGetReport.RUnlock()
	return calls
}

// PostComment calls PostCommentFunc.
func (mock *ClientMock) PostComment(ctx context.Context, data bitbucket.CommentData) error {
	if mock.PostCommentFunc == nil {
		panic("ClientMock.PostCommentFunc: method is nil but Client.PostComment was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Data bitbucket.CommentData
	}{
		Ctx:  ctx,
		Data: data,
	}
	mock.lockPostComment.Lock()
	mock.calls.PostComment = append(mock.calls.PostComment, callInfo)
	mock.lockPostComment.Unlock()
	return mock.PostCommentFunc(ctx, data)
}

// PostCommentCalls gets all the calls that were made to PostComment.
// Check the length with:
//
//	len(mockedClient.PostCommentCalls())
func (mock *ClientMock) PostCommentCalls() []struct {
	Ctx  context.Context
	Data bitbucket.CommentData
} {
	var calls []struct {
		Ctx  context.Context
		Data bitbucket.CommentData
	}
	mock.lockPostComment.RLock()
	calls = mock.calls.PostComment
	mock.lockPostComment.RUnlock()
	return calls
}

// UpsertReport calls UpsertReportFunc.
func (mock *ClientMock) UpsertReport(ctx context.Context, reportID string, data bitbucket.ReportData) error {
	if mock.UpsertReportFunc == nil {
		panic("ClientMock.UpsertReportFunc: method is nil but Client.UpsertReport was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		ReportID string
		Data     bitbucket.ReportData
	}{
		Ctx:      ctx,
		ReportID: reportID,
		Data:     data,
	}
	mock.lockUpsertReport.Lock()
	mock.calls.UpsertReport = append(mock.calls.UpsertReport, callInfo)
	mock.lockUpsertReport.Unlock()
	return mock.UpsertReportFunc(ctx, reportID, data)
}

// UpsertReportCalls gets all the calls that were made to UpsertReport.
// Check the length with:
//
//	len(mockedClient.UpsertReportCalls())
func (mock *ClientMock) UpsertReportCalls() []struct {
	Ctx      context.Context
	ReportID string
	Data     bitbucket.ReportData
} {
	var calls []struct {
		Ctx      context.Context
		ReportID string
		Data     bitbucket.ReportData
	}
	mock.lockUpsertReport.RLock()
	calls = mock.calls.UpsertReport
	mock.lockUpsertReport.RUnlock()
	return calls
}