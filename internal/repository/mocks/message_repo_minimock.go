// Code generated by http://github.com/gojuno/minimock (v3.3.14). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/ArtEmerged/o_chat-server/internal/repository.MessageRepo -o message_repo_minimock.go -n MessageRepoMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/ArtEmerged/o_chat-server/internal/model"
	"github.com/gojuno/minimock/v3"
)

// MessageRepoMock implements repository.MessageRepo
type MessageRepoMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcSendMessage          func(ctx context.Context, in *model.SendMessageRequest) (err error)
	inspectFuncSendMessage   func(ctx context.Context, in *model.SendMessageRequest)
	afterSendMessageCounter  uint64
	beforeSendMessageCounter uint64
	SendMessageMock          mMessageRepoMockSendMessage
}

// NewMessageRepoMock returns a mock for repository.MessageRepo
func NewMessageRepoMock(t minimock.Tester) *MessageRepoMock {
	m := &MessageRepoMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SendMessageMock = mMessageRepoMockSendMessage{mock: m}
	m.SendMessageMock.callArgs = []*MessageRepoMockSendMessageParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mMessageRepoMockSendMessage struct {
	optional           bool
	mock               *MessageRepoMock
	defaultExpectation *MessageRepoMockSendMessageExpectation
	expectations       []*MessageRepoMockSendMessageExpectation

	callArgs []*MessageRepoMockSendMessageParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// MessageRepoMockSendMessageExpectation specifies expectation struct of the MessageRepo.SendMessage
type MessageRepoMockSendMessageExpectation struct {
	mock      *MessageRepoMock
	params    *MessageRepoMockSendMessageParams
	paramPtrs *MessageRepoMockSendMessageParamPtrs
	results   *MessageRepoMockSendMessageResults
	Counter   uint64
}

// MessageRepoMockSendMessageParams contains parameters of the MessageRepo.SendMessage
type MessageRepoMockSendMessageParams struct {
	ctx context.Context
	in  *model.SendMessageRequest
}

// MessageRepoMockSendMessageParamPtrs contains pointers to parameters of the MessageRepo.SendMessage
type MessageRepoMockSendMessageParamPtrs struct {
	ctx *context.Context
	in  **model.SendMessageRequest
}

// MessageRepoMockSendMessageResults contains results of the MessageRepo.SendMessage
type MessageRepoMockSendMessageResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmSendMessage *mMessageRepoMockSendMessage) Optional() *mMessageRepoMockSendMessage {
	mmSendMessage.optional = true
	return mmSendMessage
}

// Expect sets up expected params for MessageRepo.SendMessage
func (mmSendMessage *mMessageRepoMockSendMessage) Expect(ctx context.Context, in *model.SendMessageRequest) *mMessageRepoMockSendMessage {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("MessageRepoMock.SendMessage mock is already set by Set")
	}

	if mmSendMessage.defaultExpectation == nil {
		mmSendMessage.defaultExpectation = &MessageRepoMockSendMessageExpectation{}
	}

	if mmSendMessage.defaultExpectation.paramPtrs != nil {
		mmSendMessage.mock.t.Fatalf("MessageRepoMock.SendMessage mock is already set by ExpectParams functions")
	}

	mmSendMessage.defaultExpectation.params = &MessageRepoMockSendMessageParams{ctx, in}
	for _, e := range mmSendMessage.expectations {
		if minimock.Equal(e.params, mmSendMessage.defaultExpectation.params) {
			mmSendMessage.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSendMessage.defaultExpectation.params)
		}
	}

	return mmSendMessage
}

// ExpectCtxParam1 sets up expected param ctx for MessageRepo.SendMessage
func (mmSendMessage *mMessageRepoMockSendMessage) ExpectCtxParam1(ctx context.Context) *mMessageRepoMockSendMessage {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("MessageRepoMock.SendMessage mock is already set by Set")
	}

	if mmSendMessage.defaultExpectation == nil {
		mmSendMessage.defaultExpectation = &MessageRepoMockSendMessageExpectation{}
	}

	if mmSendMessage.defaultExpectation.params != nil {
		mmSendMessage.mock.t.Fatalf("MessageRepoMock.SendMessage mock is already set by Expect")
	}

	if mmSendMessage.defaultExpectation.paramPtrs == nil {
		mmSendMessage.defaultExpectation.paramPtrs = &MessageRepoMockSendMessageParamPtrs{}
	}
	mmSendMessage.defaultExpectation.paramPtrs.ctx = &ctx

	return mmSendMessage
}

// ExpectInParam2 sets up expected param in for MessageRepo.SendMessage
func (mmSendMessage *mMessageRepoMockSendMessage) ExpectInParam2(in *model.SendMessageRequest) *mMessageRepoMockSendMessage {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("MessageRepoMock.SendMessage mock is already set by Set")
	}

	if mmSendMessage.defaultExpectation == nil {
		mmSendMessage.defaultExpectation = &MessageRepoMockSendMessageExpectation{}
	}

	if mmSendMessage.defaultExpectation.params != nil {
		mmSendMessage.mock.t.Fatalf("MessageRepoMock.SendMessage mock is already set by Expect")
	}

	if mmSendMessage.defaultExpectation.paramPtrs == nil {
		mmSendMessage.defaultExpectation.paramPtrs = &MessageRepoMockSendMessageParamPtrs{}
	}
	mmSendMessage.defaultExpectation.paramPtrs.in = &in

	return mmSendMessage
}

// Inspect accepts an inspector function that has same arguments as the MessageRepo.SendMessage
func (mmSendMessage *mMessageRepoMockSendMessage) Inspect(f func(ctx context.Context, in *model.SendMessageRequest)) *mMessageRepoMockSendMessage {
	if mmSendMessage.mock.inspectFuncSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("Inspect function is already set for MessageRepoMock.SendMessage")
	}

	mmSendMessage.mock.inspectFuncSendMessage = f

	return mmSendMessage
}

// Return sets up results that will be returned by MessageRepo.SendMessage
func (mmSendMessage *mMessageRepoMockSendMessage) Return(err error) *MessageRepoMock {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("MessageRepoMock.SendMessage mock is already set by Set")
	}

	if mmSendMessage.defaultExpectation == nil {
		mmSendMessage.defaultExpectation = &MessageRepoMockSendMessageExpectation{mock: mmSendMessage.mock}
	}
	mmSendMessage.defaultExpectation.results = &MessageRepoMockSendMessageResults{err}
	return mmSendMessage.mock
}

// Set uses given function f to mock the MessageRepo.SendMessage method
func (mmSendMessage *mMessageRepoMockSendMessage) Set(f func(ctx context.Context, in *model.SendMessageRequest) (err error)) *MessageRepoMock {
	if mmSendMessage.defaultExpectation != nil {
		mmSendMessage.mock.t.Fatalf("Default expectation is already set for the MessageRepo.SendMessage method")
	}

	if len(mmSendMessage.expectations) > 0 {
		mmSendMessage.mock.t.Fatalf("Some expectations are already set for the MessageRepo.SendMessage method")
	}

	mmSendMessage.mock.funcSendMessage = f
	return mmSendMessage.mock
}

// When sets expectation for the MessageRepo.SendMessage which will trigger the result defined by the following
// Then helper
func (mmSendMessage *mMessageRepoMockSendMessage) When(ctx context.Context, in *model.SendMessageRequest) *MessageRepoMockSendMessageExpectation {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("MessageRepoMock.SendMessage mock is already set by Set")
	}

	expectation := &MessageRepoMockSendMessageExpectation{
		mock:   mmSendMessage.mock,
		params: &MessageRepoMockSendMessageParams{ctx, in},
	}
	mmSendMessage.expectations = append(mmSendMessage.expectations, expectation)
	return expectation
}

// Then sets up MessageRepo.SendMessage return parameters for the expectation previously defined by the When method
func (e *MessageRepoMockSendMessageExpectation) Then(err error) *MessageRepoMock {
	e.results = &MessageRepoMockSendMessageResults{err}
	return e.mock
}

// Times sets number of times MessageRepo.SendMessage should be invoked
func (mmSendMessage *mMessageRepoMockSendMessage) Times(n uint64) *mMessageRepoMockSendMessage {
	if n == 0 {
		mmSendMessage.mock.t.Fatalf("Times of MessageRepoMock.SendMessage mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmSendMessage.expectedInvocations, n)
	return mmSendMessage
}

func (mmSendMessage *mMessageRepoMockSendMessage) invocationsDone() bool {
	if len(mmSendMessage.expectations) == 0 && mmSendMessage.defaultExpectation == nil && mmSendMessage.mock.funcSendMessage == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmSendMessage.mock.afterSendMessageCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmSendMessage.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// SendMessage implements repository.MessageRepo
func (mmSendMessage *MessageRepoMock) SendMessage(ctx context.Context, in *model.SendMessageRequest) (err error) {
	mm_atomic.AddUint64(&mmSendMessage.beforeSendMessageCounter, 1)
	defer mm_atomic.AddUint64(&mmSendMessage.afterSendMessageCounter, 1)

	if mmSendMessage.inspectFuncSendMessage != nil {
		mmSendMessage.inspectFuncSendMessage(ctx, in)
	}

	mm_params := MessageRepoMockSendMessageParams{ctx, in}

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

		mm_got := MessageRepoMockSendMessageParams{ctx, in}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmSendMessage.t.Errorf("MessageRepoMock.SendMessage got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.in != nil && !minimock.Equal(*mm_want_ptrs.in, mm_got.in) {
				mmSendMessage.t.Errorf("MessageRepoMock.SendMessage got unexpected parameter in, want: %#v, got: %#v%s\n", *mm_want_ptrs.in, mm_got.in, minimock.Diff(*mm_want_ptrs.in, mm_got.in))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSendMessage.t.Errorf("MessageRepoMock.SendMessage got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSendMessage.SendMessageMock.defaultExpectation.results
		if mm_results == nil {
			mmSendMessage.t.Fatal("No results are set for the MessageRepoMock.SendMessage")
		}
		return (*mm_results).err
	}
	if mmSendMessage.funcSendMessage != nil {
		return mmSendMessage.funcSendMessage(ctx, in)
	}
	mmSendMessage.t.Fatalf("Unexpected call to MessageRepoMock.SendMessage. %v %v", ctx, in)
	return
}

// SendMessageAfterCounter returns a count of finished MessageRepoMock.SendMessage invocations
func (mmSendMessage *MessageRepoMock) SendMessageAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendMessage.afterSendMessageCounter)
}

// SendMessageBeforeCounter returns a count of MessageRepoMock.SendMessage invocations
func (mmSendMessage *MessageRepoMock) SendMessageBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendMessage.beforeSendMessageCounter)
}

// Calls returns a list of arguments used in each call to MessageRepoMock.SendMessage.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSendMessage *mMessageRepoMockSendMessage) Calls() []*MessageRepoMockSendMessageParams {
	mmSendMessage.mutex.RLock()

	argCopy := make([]*MessageRepoMockSendMessageParams, len(mmSendMessage.callArgs))
	copy(argCopy, mmSendMessage.callArgs)

	mmSendMessage.mutex.RUnlock()

	return argCopy
}

// MinimockSendMessageDone returns true if the count of the SendMessage invocations corresponds
// the number of defined expectations
func (m *MessageRepoMock) MinimockSendMessageDone() bool {
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
func (m *MessageRepoMock) MinimockSendMessageInspect() {
	for _, e := range m.SendMessageMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MessageRepoMock.SendMessage with params: %#v", *e.params)
		}
	}

	afterSendMessageCounter := mm_atomic.LoadUint64(&m.afterSendMessageCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.SendMessageMock.defaultExpectation != nil && afterSendMessageCounter < 1 {
		if m.SendMessageMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MessageRepoMock.SendMessage")
		} else {
			m.t.Errorf("Expected call to MessageRepoMock.SendMessage with params: %#v", *m.SendMessageMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendMessage != nil && afterSendMessageCounter < 1 {
		m.t.Error("Expected call to MessageRepoMock.SendMessage")
	}

	if !m.SendMessageMock.invocationsDone() && afterSendMessageCounter > 0 {
		m.t.Errorf("Expected %d calls to MessageRepoMock.SendMessage but found %d calls",
			mm_atomic.LoadUint64(&m.SendMessageMock.expectedInvocations), afterSendMessageCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MessageRepoMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockSendMessageInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MessageRepoMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *MessageRepoMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSendMessageDone()
}