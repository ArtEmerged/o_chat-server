// Code generated by http://github.com/gojuno/minimock (v3.3.14). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/ArtEmerged/o_chat-server/internal/service.MessageService -o message_service_minimock.go -n MessageServiceMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/ArtEmerged/o_chat-server/internal/model"
	"github.com/gojuno/minimock/v3"
)

// MessageServiceMock implements service.MessageService
type MessageServiceMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcSendMessage          func(ctx context.Context, in *model.SendMessageRequest) (err error)
	inspectFuncSendMessage   func(ctx context.Context, in *model.SendMessageRequest)
	afterSendMessageCounter  uint64
	beforeSendMessageCounter uint64
	SendMessageMock          mMessageServiceMockSendMessage
}

// NewMessageServiceMock returns a mock for service.MessageService
func NewMessageServiceMock(t minimock.Tester) *MessageServiceMock {
	m := &MessageServiceMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SendMessageMock = mMessageServiceMockSendMessage{mock: m}
	m.SendMessageMock.callArgs = []*MessageServiceMockSendMessageParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mMessageServiceMockSendMessage struct {
	optional           bool
	mock               *MessageServiceMock
	defaultExpectation *MessageServiceMockSendMessageExpectation
	expectations       []*MessageServiceMockSendMessageExpectation

	callArgs []*MessageServiceMockSendMessageParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// MessageServiceMockSendMessageExpectation specifies expectation struct of the MessageService.SendMessage
type MessageServiceMockSendMessageExpectation struct {
	mock      *MessageServiceMock
	params    *MessageServiceMockSendMessageParams
	paramPtrs *MessageServiceMockSendMessageParamPtrs
	results   *MessageServiceMockSendMessageResults
	Counter   uint64
}

// MessageServiceMockSendMessageParams contains parameters of the MessageService.SendMessage
type MessageServiceMockSendMessageParams struct {
	ctx context.Context
	in  *model.SendMessageRequest
}

// MessageServiceMockSendMessageParamPtrs contains pointers to parameters of the MessageService.SendMessage
type MessageServiceMockSendMessageParamPtrs struct {
	ctx *context.Context
	in  **model.SendMessageRequest
}

// MessageServiceMockSendMessageResults contains results of the MessageService.SendMessage
type MessageServiceMockSendMessageResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmSendMessage *mMessageServiceMockSendMessage) Optional() *mMessageServiceMockSendMessage {
	mmSendMessage.optional = true
	return mmSendMessage
}

// Expect sets up expected params for MessageService.SendMessage
func (mmSendMessage *mMessageServiceMockSendMessage) Expect(ctx context.Context, in *model.SendMessageRequest) *mMessageServiceMockSendMessage {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("MessageServiceMock.SendMessage mock is already set by Set")
	}

	if mmSendMessage.defaultExpectation == nil {
		mmSendMessage.defaultExpectation = &MessageServiceMockSendMessageExpectation{}
	}

	if mmSendMessage.defaultExpectation.paramPtrs != nil {
		mmSendMessage.mock.t.Fatalf("MessageServiceMock.SendMessage mock is already set by ExpectParams functions")
	}

	mmSendMessage.defaultExpectation.params = &MessageServiceMockSendMessageParams{ctx, in}
	for _, e := range mmSendMessage.expectations {
		if minimock.Equal(e.params, mmSendMessage.defaultExpectation.params) {
			mmSendMessage.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSendMessage.defaultExpectation.params)
		}
	}

	return mmSendMessage
}

// ExpectCtxParam1 sets up expected param ctx for MessageService.SendMessage
func (mmSendMessage *mMessageServiceMockSendMessage) ExpectCtxParam1(ctx context.Context) *mMessageServiceMockSendMessage {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("MessageServiceMock.SendMessage mock is already set by Set")
	}

	if mmSendMessage.defaultExpectation == nil {
		mmSendMessage.defaultExpectation = &MessageServiceMockSendMessageExpectation{}
	}

	if mmSendMessage.defaultExpectation.params != nil {
		mmSendMessage.mock.t.Fatalf("MessageServiceMock.SendMessage mock is already set by Expect")
	}

	if mmSendMessage.defaultExpectation.paramPtrs == nil {
		mmSendMessage.defaultExpectation.paramPtrs = &MessageServiceMockSendMessageParamPtrs{}
	}
	mmSendMessage.defaultExpectation.paramPtrs.ctx = &ctx

	return mmSendMessage
}

// ExpectInParam2 sets up expected param in for MessageService.SendMessage
func (mmSendMessage *mMessageServiceMockSendMessage) ExpectInParam2(in *model.SendMessageRequest) *mMessageServiceMockSendMessage {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("MessageServiceMock.SendMessage mock is already set by Set")
	}

	if mmSendMessage.defaultExpectation == nil {
		mmSendMessage.defaultExpectation = &MessageServiceMockSendMessageExpectation{}
	}

	if mmSendMessage.defaultExpectation.params != nil {
		mmSendMessage.mock.t.Fatalf("MessageServiceMock.SendMessage mock is already set by Expect")
	}

	if mmSendMessage.defaultExpectation.paramPtrs == nil {
		mmSendMessage.defaultExpectation.paramPtrs = &MessageServiceMockSendMessageParamPtrs{}
	}
	mmSendMessage.defaultExpectation.paramPtrs.in = &in

	return mmSendMessage
}

// Inspect accepts an inspector function that has same arguments as the MessageService.SendMessage
func (mmSendMessage *mMessageServiceMockSendMessage) Inspect(f func(ctx context.Context, in *model.SendMessageRequest)) *mMessageServiceMockSendMessage {
	if mmSendMessage.mock.inspectFuncSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("Inspect function is already set for MessageServiceMock.SendMessage")
	}

	mmSendMessage.mock.inspectFuncSendMessage = f

	return mmSendMessage
}

// Return sets up results that will be returned by MessageService.SendMessage
func (mmSendMessage *mMessageServiceMockSendMessage) Return(err error) *MessageServiceMock {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("MessageServiceMock.SendMessage mock is already set by Set")
	}

	if mmSendMessage.defaultExpectation == nil {
		mmSendMessage.defaultExpectation = &MessageServiceMockSendMessageExpectation{mock: mmSendMessage.mock}
	}
	mmSendMessage.defaultExpectation.results = &MessageServiceMockSendMessageResults{err}
	return mmSendMessage.mock
}

// Set uses given function f to mock the MessageService.SendMessage method
func (mmSendMessage *mMessageServiceMockSendMessage) Set(f func(ctx context.Context, in *model.SendMessageRequest) (err error)) *MessageServiceMock {
	if mmSendMessage.defaultExpectation != nil {
		mmSendMessage.mock.t.Fatalf("Default expectation is already set for the MessageService.SendMessage method")
	}

	if len(mmSendMessage.expectations) > 0 {
		mmSendMessage.mock.t.Fatalf("Some expectations are already set for the MessageService.SendMessage method")
	}

	mmSendMessage.mock.funcSendMessage = f
	return mmSendMessage.mock
}

// When sets expectation for the MessageService.SendMessage which will trigger the result defined by the following
// Then helper
func (mmSendMessage *mMessageServiceMockSendMessage) When(ctx context.Context, in *model.SendMessageRequest) *MessageServiceMockSendMessageExpectation {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("MessageServiceMock.SendMessage mock is already set by Set")
	}

	expectation := &MessageServiceMockSendMessageExpectation{
		mock:   mmSendMessage.mock,
		params: &MessageServiceMockSendMessageParams{ctx, in},
	}
	mmSendMessage.expectations = append(mmSendMessage.expectations, expectation)
	return expectation
}

// Then sets up MessageService.SendMessage return parameters for the expectation previously defined by the When method
func (e *MessageServiceMockSendMessageExpectation) Then(err error) *MessageServiceMock {
	e.results = &MessageServiceMockSendMessageResults{err}
	return e.mock
}

// Times sets number of times MessageService.SendMessage should be invoked
func (mmSendMessage *mMessageServiceMockSendMessage) Times(n uint64) *mMessageServiceMockSendMessage {
	if n == 0 {
		mmSendMessage.mock.t.Fatalf("Times of MessageServiceMock.SendMessage mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmSendMessage.expectedInvocations, n)
	return mmSendMessage
}

func (mmSendMessage *mMessageServiceMockSendMessage) invocationsDone() bool {
	if len(mmSendMessage.expectations) == 0 && mmSendMessage.defaultExpectation == nil && mmSendMessage.mock.funcSendMessage == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmSendMessage.mock.afterSendMessageCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmSendMessage.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// SendMessage implements service.MessageService
func (mmSendMessage *MessageServiceMock) SendMessage(ctx context.Context, in *model.SendMessageRequest) (err error) {
	mm_atomic.AddUint64(&mmSendMessage.beforeSendMessageCounter, 1)
	defer mm_atomic.AddUint64(&mmSendMessage.afterSendMessageCounter, 1)

	if mmSendMessage.inspectFuncSendMessage != nil {
		mmSendMessage.inspectFuncSendMessage(ctx, in)
	}

	mm_params := MessageServiceMockSendMessageParams{ctx, in}

	// Record call args
	mmSendMessage.SendMessageMock.mutex.Lock()
	mmSendMessage.SendMessageMock.callArgs = append(mmSendMessage.SendMessageMock.callArgs, &mm_params)
	mmSendMessage.SendMessageMock.mutex.Unlock()

	for _, e := range mmSendMessage.SendMessageMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSendMessage.SendMessageMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSendMessage.SendMessageMock.defaultExpectation.Counter, 1)
		mm_want := mmSendMessage.SendMessageMock.defaultExpectation.params
		mm_want_ptrs := mmSendMessage.SendMessageMock.defaultExpectation.paramPtrs

		mm_got := MessageServiceMockSendMessageParams{ctx, in}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmSendMessage.t.Errorf("MessageServiceMock.SendMessage got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.in != nil && !minimock.Equal(*mm_want_ptrs.in, mm_got.in) {
				mmSendMessage.t.Errorf("MessageServiceMock.SendMessage got unexpected parameter in, want: %#v, got: %#v%s\n", *mm_want_ptrs.in, mm_got.in, minimock.Diff(*mm_want_ptrs.in, mm_got.in))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSendMessage.t.Errorf("MessageServiceMock.SendMessage got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSendMessage.SendMessageMock.defaultExpectation.results
		if mm_results == nil {
			mmSendMessage.t.Fatal("No results are set for the MessageServiceMock.SendMessage")
		}
		return (*mm_results).err
	}
	if mmSendMessage.funcSendMessage != nil {
		return mmSendMessage.funcSendMessage(ctx, in)
	}
	mmSendMessage.t.Fatalf("Unexpected call to MessageServiceMock.SendMessage. %v %v", ctx, in)
	return
}

// SendMessageAfterCounter returns a count of finished MessageServiceMock.SendMessage invocations
func (mmSendMessage *MessageServiceMock) SendMessageAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendMessage.afterSendMessageCounter)
}

// SendMessageBeforeCounter returns a count of MessageServiceMock.SendMessage invocations
func (mmSendMessage *MessageServiceMock) SendMessageBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendMessage.beforeSendMessageCounter)
}

// Calls returns a list of arguments used in each call to MessageServiceMock.SendMessage.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSendMessage *mMessageServiceMockSendMessage) Calls() []*MessageServiceMockSendMessageParams {
	mmSendMessage.mutex.RLock()

	argCopy := make([]*MessageServiceMockSendMessageParams, len(mmSendMessage.callArgs))
	copy(argCopy, mmSendMessage.callArgs)

	mmSendMessage.mutex.RUnlock()

	return argCopy
}

// MinimockSendMessageDone returns true if the count of the SendMessage invocations corresponds
// the number of defined expectations
func (m *MessageServiceMock) MinimockSendMessageDone() bool {
	if m.SendMessageMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.SendMessageMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.SendMessageMock.invocationsDone()
}

// MinimockSendMessageInspect logs each unmet expectation
func (m *MessageServiceMock) MinimockSendMessageInspect() {
	for _, e := range m.SendMessageMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MessageServiceMock.SendMessage with params: %#v", *e.params)
		}
	}

	afterSendMessageCounter := mm_atomic.LoadUint64(&m.afterSendMessageCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.SendMessageMock.defaultExpectation != nil && afterSendMessageCounter < 1 {
		if m.SendMessageMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MessageServiceMock.SendMessage")
		} else {
			m.t.Errorf("Expected call to MessageServiceMock.SendMessage with params: %#v", *m.SendMessageMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendMessage != nil && afterSendMessageCounter < 1 {
		m.t.Error("Expected call to MessageServiceMock.SendMessage")
	}

	if !m.SendMessageMock.invocationsDone() && afterSendMessageCounter > 0 {
		m.t.Errorf("Expected %d calls to MessageServiceMock.SendMessage but found %d calls",
			mm_atomic.LoadUint64(&m.SendMessageMock.expectedInvocations), afterSendMessageCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MessageServiceMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockSendMessageInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MessageServiceMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *MessageServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSendMessageDone()
}
