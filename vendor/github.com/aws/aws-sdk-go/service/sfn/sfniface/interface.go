// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package sfniface provides an interface to enable mocking the AWS Step Functions service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package sfniface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sfn"
)

// SFNAPI provides an interface to enable mocking the
// sfn.SFN service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Step Functions.
//	func myFunc(svc sfniface.SFNAPI) bool {
//	    // Make svc.CreateActivity request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := sfn.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockSFNClient struct {
//	    sfniface.SFNAPI
//	}
//	func (m *mockSFNClient) CreateActivity(input *sfn.CreateActivityInput) (*sfn.CreateActivityOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockSFNClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type SFNAPI interface {
	CreateActivity(*sfn.CreateActivityInput) (*sfn.CreateActivityOutput, error)
	CreateActivityWithContext(aws.Context, *sfn.CreateActivityInput, ...request.Option) (*sfn.CreateActivityOutput, error)
	CreateActivityRequest(*sfn.CreateActivityInput) (*request.Request, *sfn.CreateActivityOutput)

	CreateStateMachine(*sfn.CreateStateMachineInput) (*sfn.CreateStateMachineOutput, error)
	CreateStateMachineWithContext(aws.Context, *sfn.CreateStateMachineInput, ...request.Option) (*sfn.CreateStateMachineOutput, error)
	CreateStateMachineRequest(*sfn.CreateStateMachineInput) (*request.Request, *sfn.CreateStateMachineOutput)

	CreateStateMachineAlias(*sfn.CreateStateMachineAliasInput) (*sfn.CreateStateMachineAliasOutput, error)
	CreateStateMachineAliasWithContext(aws.Context, *sfn.CreateStateMachineAliasInput, ...request.Option) (*sfn.CreateStateMachineAliasOutput, error)
	CreateStateMachineAliasRequest(*sfn.CreateStateMachineAliasInput) (*request.Request, *sfn.CreateStateMachineAliasOutput)

	DeleteActivity(*sfn.DeleteActivityInput) (*sfn.DeleteActivityOutput, error)
	DeleteActivityWithContext(aws.Context, *sfn.DeleteActivityInput, ...request.Option) (*sfn.DeleteActivityOutput, error)
	DeleteActivityRequest(*sfn.DeleteActivityInput) (*request.Request, *sfn.DeleteActivityOutput)

	DeleteStateMachine(*sfn.DeleteStateMachineInput) (*sfn.DeleteStateMachineOutput, error)
	DeleteStateMachineWithContext(aws.Context, *sfn.DeleteStateMachineInput, ...request.Option) (*sfn.DeleteStateMachineOutput, error)
	DeleteStateMachineRequest(*sfn.DeleteStateMachineInput) (*request.Request, *sfn.DeleteStateMachineOutput)

	DeleteStateMachineAlias(*sfn.DeleteStateMachineAliasInput) (*sfn.DeleteStateMachineAliasOutput, error)
	DeleteStateMachineAliasWithContext(aws.Context, *sfn.DeleteStateMachineAliasInput, ...request.Option) (*sfn.DeleteStateMachineAliasOutput, error)
	DeleteStateMachineAliasRequest(*sfn.DeleteStateMachineAliasInput) (*request.Request, *sfn.DeleteStateMachineAliasOutput)

	DeleteStateMachineVersion(*sfn.DeleteStateMachineVersionInput) (*sfn.DeleteStateMachineVersionOutput, error)
	DeleteStateMachineVersionWithContext(aws.Context, *sfn.DeleteStateMachineVersionInput, ...request.Option) (*sfn.DeleteStateMachineVersionOutput, error)
	DeleteStateMachineVersionRequest(*sfn.DeleteStateMachineVersionInput) (*request.Request, *sfn.DeleteStateMachineVersionOutput)

	DescribeActivity(*sfn.DescribeActivityInput) (*sfn.DescribeActivityOutput, error)
	DescribeActivityWithContext(aws.Context, *sfn.DescribeActivityInput, ...request.Option) (*sfn.DescribeActivityOutput, error)
	DescribeActivityRequest(*sfn.DescribeActivityInput) (*request.Request, *sfn.DescribeActivityOutput)

	DescribeExecution(*sfn.DescribeExecutionInput) (*sfn.DescribeExecutionOutput, error)
	DescribeExecutionWithContext(aws.Context, *sfn.DescribeExecutionInput, ...request.Option) (*sfn.DescribeExecutionOutput, error)
	DescribeExecutionRequest(*sfn.DescribeExecutionInput) (*request.Request, *sfn.DescribeExecutionOutput)

	DescribeMapRun(*sfn.DescribeMapRunInput) (*sfn.DescribeMapRunOutput, error)
	DescribeMapRunWithContext(aws.Context, *sfn.DescribeMapRunInput, ...request.Option) (*sfn.DescribeMapRunOutput, error)
	DescribeMapRunRequest(*sfn.DescribeMapRunInput) (*request.Request, *sfn.DescribeMapRunOutput)

	DescribeStateMachine(*sfn.DescribeStateMachineInput) (*sfn.DescribeStateMachineOutput, error)
	DescribeStateMachineWithContext(aws.Context, *sfn.DescribeStateMachineInput, ...request.Option) (*sfn.DescribeStateMachineOutput, error)
	DescribeStateMachineRequest(*sfn.DescribeStateMachineInput) (*request.Request, *sfn.DescribeStateMachineOutput)

	DescribeStateMachineAlias(*sfn.DescribeStateMachineAliasInput) (*sfn.DescribeStateMachineAliasOutput, error)
	DescribeStateMachineAliasWithContext(aws.Context, *sfn.DescribeStateMachineAliasInput, ...request.Option) (*sfn.DescribeStateMachineAliasOutput, error)
	DescribeStateMachineAliasRequest(*sfn.DescribeStateMachineAliasInput) (*request.Request, *sfn.DescribeStateMachineAliasOutput)

	DescribeStateMachineForExecution(*sfn.DescribeStateMachineForExecutionInput) (*sfn.DescribeStateMachineForExecutionOutput, error)
	DescribeStateMachineForExecutionWithContext(aws.Context, *sfn.DescribeStateMachineForExecutionInput, ...request.Option) (*sfn.DescribeStateMachineForExecutionOutput, error)
	DescribeStateMachineForExecutionRequest(*sfn.DescribeStateMachineForExecutionInput) (*request.Request, *sfn.DescribeStateMachineForExecutionOutput)

	GetActivityTask(*sfn.GetActivityTaskInput) (*sfn.GetActivityTaskOutput, error)
	GetActivityTaskWithContext(aws.Context, *sfn.GetActivityTaskInput, ...request.Option) (*sfn.GetActivityTaskOutput, error)
	GetActivityTaskRequest(*sfn.GetActivityTaskInput) (*request.Request, *sfn.GetActivityTaskOutput)

	GetExecutionHistory(*sfn.GetExecutionHistoryInput) (*sfn.GetExecutionHistoryOutput, error)
	GetExecutionHistoryWithContext(aws.Context, *sfn.GetExecutionHistoryInput, ...request.Option) (*sfn.GetExecutionHistoryOutput, error)
	GetExecutionHistoryRequest(*sfn.GetExecutionHistoryInput) (*request.Request, *sfn.GetExecutionHistoryOutput)

	GetExecutionHistoryPages(*sfn.GetExecutionHistoryInput, func(*sfn.GetExecutionHistoryOutput, bool) bool) error
	GetExecutionHistoryPagesWithContext(aws.Context, *sfn.GetExecutionHistoryInput, func(*sfn.GetExecutionHistoryOutput, bool) bool, ...request.Option) error

	ListActivities(*sfn.ListActivitiesInput) (*sfn.ListActivitiesOutput, error)
	ListActivitiesWithContext(aws.Context, *sfn.ListActivitiesInput, ...request.Option) (*sfn.ListActivitiesOutput, error)
	ListActivitiesRequest(*sfn.ListActivitiesInput) (*request.Request, *sfn.ListActivitiesOutput)

	ListActivitiesPages(*sfn.ListActivitiesInput, func(*sfn.ListActivitiesOutput, bool) bool) error
	ListActivitiesPagesWithContext(aws.Context, *sfn.ListActivitiesInput, func(*sfn.ListActivitiesOutput, bool) bool, ...request.Option) error

	ListExecutions(*sfn.ListExecutionsInput) (*sfn.ListExecutionsOutput, error)
	ListExecutionsWithContext(aws.Context, *sfn.ListExecutionsInput, ...request.Option) (*sfn.ListExecutionsOutput, error)
	ListExecutionsRequest(*sfn.ListExecutionsInput) (*request.Request, *sfn.ListExecutionsOutput)

	ListExecutionsPages(*sfn.ListExecutionsInput, func(*sfn.ListExecutionsOutput, bool) bool) error
	ListExecutionsPagesWithContext(aws.Context, *sfn.ListExecutionsInput, func(*sfn.ListExecutionsOutput, bool) bool, ...request.Option) error

	ListMapRuns(*sfn.ListMapRunsInput) (*sfn.ListMapRunsOutput, error)
	ListMapRunsWithContext(aws.Context, *sfn.ListMapRunsInput, ...request.Option) (*sfn.ListMapRunsOutput, error)
	ListMapRunsRequest(*sfn.ListMapRunsInput) (*request.Request, *sfn.ListMapRunsOutput)

	ListMapRunsPages(*sfn.ListMapRunsInput, func(*sfn.ListMapRunsOutput, bool) bool) error
	ListMapRunsPagesWithContext(aws.Context, *sfn.ListMapRunsInput, func(*sfn.ListMapRunsOutput, bool) bool, ...request.Option) error

	ListStateMachineAliases(*sfn.ListStateMachineAliasesInput) (*sfn.ListStateMachineAliasesOutput, error)
	ListStateMachineAliasesWithContext(aws.Context, *sfn.ListStateMachineAliasesInput, ...request.Option) (*sfn.ListStateMachineAliasesOutput, error)
	ListStateMachineAliasesRequest(*sfn.ListStateMachineAliasesInput) (*request.Request, *sfn.ListStateMachineAliasesOutput)

	ListStateMachineVersions(*sfn.ListStateMachineVersionsInput) (*sfn.ListStateMachineVersionsOutput, error)
	ListStateMachineVersionsWithContext(aws.Context, *sfn.ListStateMachineVersionsInput, ...request.Option) (*sfn.ListStateMachineVersionsOutput, error)
	ListStateMachineVersionsRequest(*sfn.ListStateMachineVersionsInput) (*request.Request, *sfn.ListStateMachineVersionsOutput)

	ListStateMachines(*sfn.ListStateMachinesInput) (*sfn.ListStateMachinesOutput, error)
	ListStateMachinesWithContext(aws.Context, *sfn.ListStateMachinesInput, ...request.Option) (*sfn.ListStateMachinesOutput, error)
	ListStateMachinesRequest(*sfn.ListStateMachinesInput) (*request.Request, *sfn.ListStateMachinesOutput)

	ListStateMachinesPages(*sfn.ListStateMachinesInput, func(*sfn.ListStateMachinesOutput, bool) bool) error
	ListStateMachinesPagesWithContext(aws.Context, *sfn.ListStateMachinesInput, func(*sfn.ListStateMachinesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*sfn.ListTagsForResourceInput) (*sfn.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *sfn.ListTagsForResourceInput, ...request.Option) (*sfn.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*sfn.ListTagsForResourceInput) (*request.Request, *sfn.ListTagsForResourceOutput)

	PublishStateMachineVersion(*sfn.PublishStateMachineVersionInput) (*sfn.PublishStateMachineVersionOutput, error)
	PublishStateMachineVersionWithContext(aws.Context, *sfn.PublishStateMachineVersionInput, ...request.Option) (*sfn.PublishStateMachineVersionOutput, error)
	PublishStateMachineVersionRequest(*sfn.PublishStateMachineVersionInput) (*request.Request, *sfn.PublishStateMachineVersionOutput)

	SendTaskFailure(*sfn.SendTaskFailureInput) (*sfn.SendTaskFailureOutput, error)
	SendTaskFailureWithContext(aws.Context, *sfn.SendTaskFailureInput, ...request.Option) (*sfn.SendTaskFailureOutput, error)
	SendTaskFailureRequest(*sfn.SendTaskFailureInput) (*request.Request, *sfn.SendTaskFailureOutput)

	SendTaskHeartbeat(*sfn.SendTaskHeartbeatInput) (*sfn.SendTaskHeartbeatOutput, error)
	SendTaskHeartbeatWithContext(aws.Context, *sfn.SendTaskHeartbeatInput, ...request.Option) (*sfn.SendTaskHeartbeatOutput, error)
	SendTaskHeartbeatRequest(*sfn.SendTaskHeartbeatInput) (*request.Request, *sfn.SendTaskHeartbeatOutput)

	SendTaskSuccess(*sfn.SendTaskSuccessInput) (*sfn.SendTaskSuccessOutput, error)
	SendTaskSuccessWithContext(aws.Context, *sfn.SendTaskSuccessInput, ...request.Option) (*sfn.SendTaskSuccessOutput, error)
	SendTaskSuccessRequest(*sfn.SendTaskSuccessInput) (*request.Request, *sfn.SendTaskSuccessOutput)

	StartExecution(*sfn.StartExecutionInput) (*sfn.StartExecutionOutput, error)
	StartExecutionWithContext(aws.Context, *sfn.StartExecutionInput, ...request.Option) (*sfn.StartExecutionOutput, error)
	StartExecutionRequest(*sfn.StartExecutionInput) (*request.Request, *sfn.StartExecutionOutput)

	StartSyncExecution(*sfn.StartSyncExecutionInput) (*sfn.StartSyncExecutionOutput, error)
	StartSyncExecutionWithContext(aws.Context, *sfn.StartSyncExecutionInput, ...request.Option) (*sfn.StartSyncExecutionOutput, error)
	StartSyncExecutionRequest(*sfn.StartSyncExecutionInput) (*request.Request, *sfn.StartSyncExecutionOutput)

	StopExecution(*sfn.StopExecutionInput) (*sfn.StopExecutionOutput, error)
	StopExecutionWithContext(aws.Context, *sfn.StopExecutionInput, ...request.Option) (*sfn.StopExecutionOutput, error)
	StopExecutionRequest(*sfn.StopExecutionInput) (*request.Request, *sfn.StopExecutionOutput)

	TagResource(*sfn.TagResourceInput) (*sfn.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *sfn.TagResourceInput, ...request.Option) (*sfn.TagResourceOutput, error)
	TagResourceRequest(*sfn.TagResourceInput) (*request.Request, *sfn.TagResourceOutput)

	UntagResource(*sfn.UntagResourceInput) (*sfn.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *sfn.UntagResourceInput, ...request.Option) (*sfn.UntagResourceOutput, error)
	UntagResourceRequest(*sfn.UntagResourceInput) (*request.Request, *sfn.UntagResourceOutput)

	UpdateMapRun(*sfn.UpdateMapRunInput) (*sfn.UpdateMapRunOutput, error)
	UpdateMapRunWithContext(aws.Context, *sfn.UpdateMapRunInput, ...request.Option) (*sfn.UpdateMapRunOutput, error)
	UpdateMapRunRequest(*sfn.UpdateMapRunInput) (*request.Request, *sfn.UpdateMapRunOutput)

	UpdateStateMachine(*sfn.UpdateStateMachineInput) (*sfn.UpdateStateMachineOutput, error)
	UpdateStateMachineWithContext(aws.Context, *sfn.UpdateStateMachineInput, ...request.Option) (*sfn.UpdateStateMachineOutput, error)
	UpdateStateMachineRequest(*sfn.UpdateStateMachineInput) (*request.Request, *sfn.UpdateStateMachineOutput)

	UpdateStateMachineAlias(*sfn.UpdateStateMachineAliasInput) (*sfn.UpdateStateMachineAliasOutput, error)
	UpdateStateMachineAliasWithContext(aws.Context, *sfn.UpdateStateMachineAliasInput, ...request.Option) (*sfn.UpdateStateMachineAliasOutput, error)
	UpdateStateMachineAliasRequest(*sfn.UpdateStateMachineAliasInput) (*request.Request, *sfn.UpdateStateMachineAliasOutput)
}

var _ SFNAPI = (*sfn.SFN)(nil)
