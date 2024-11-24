// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/zawa-t/pr/commentator/src/report"
	"sync"
)

// Ensure, that ReporterMock does implement report.Reporter.
// If this is not the case, regenerate this file with moq.
var _ report.Reporter = &ReporterMock{}

// ReporterMock is a mock implementation of report.Reporter.
//
//	func TestSomethingThatUsesReporter(t *testing.T) {
//
//		// make and configure a mocked report.Reporter
//		mockedReporter := &ReporterMock{
//			ReportFunc: func(ctx context.Context, data report.Data) error {
//				panic("mock out the Report method")
//			},
//		}
//
//		// use mockedReporter in code that requires report.Reporter
//		// and then make assertions.
//
//	}
type ReporterMock struct {
	// ReportFunc mocks the Report method.
	ReportFunc func(ctx context.Context, data report.Data) error

	// calls tracks calls to the methods.
	calls struct {
		// Report holds details about calls to the Report method.
		Report []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Data is the data argument value.
			Data report.Data
		}
	}
	lockReport sync.RWMutex
}

// Report calls ReportFunc.
func (mock *ReporterMock) Report(ctx context.Context, data report.Data) error {
	if mock.ReportFunc == nil {
		panic("ReporterMock.ReportFunc: method is nil but Reporter.Report was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Data report.Data
	}{
		Ctx:  ctx,
		Data: data,
	}
	mock.lockReport.Lock()
	mock.calls.Report = append(mock.calls.Report, callInfo)
	mock.lockReport.Unlock()
	return mock.ReportFunc(ctx, data)
}

// ReportCalls gets all the calls that were made to Report.
// Check the length with:
//
//	len(mockedReporter.ReportCalls())
func (mock *ReporterMock) ReportCalls() []struct {
	Ctx  context.Context
	Data report.Data
} {
	var calls []struct {
		Ctx  context.Context
		Data report.Data
	}
	mock.lockReport.RLock()
	calls = mock.calls.Report
	mock.lockReport.RUnlock()
	return calls
}
