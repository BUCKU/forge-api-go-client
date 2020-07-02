package dm

import (
	"forge-api-go-client/oauth"
	"reflect"
	"testing"
)

func TestIssuesAPI_GetIssues(t *testing.T) {
	type fields struct {
		TwoLeggedAuth oauth.TwoLeggedAuth
		IssuesAPIPath string
	}
	type args struct {
		issueContainerId string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult *IssuesContainerData
		wantErr    bool
	}{
		{
			name:       "Test get issues",
			fields:     fields{
				TwoLeggedAuth: oauth.NewTwoLeggedClient("eq3LCJgGxVwByt9lGKMF0A13RTnJqzgT", "p5mzDfgr1xMmNu9V"),
				IssuesAPIPath: "/issues/v1/containers",
			},
			args:       args{
				issueContainerId: "eccc0a0a-de91-4eef-bf5e-ae22d5494534",
			},
			wantResult: nil,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &IssuesAPI{
				TwoLeggedAuth: tt.fields.TwoLeggedAuth,
				IssuesAPIPath: tt.fields.IssuesAPIPath,
			}
			gotResult, err := api.GetIssues(tt.args.issueContainerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIssues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("GetIssues() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestNewIssuesAPIWithCredentials(t *testing.T) {
	type args struct {
		clientID     string
		clientSecret string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test api constructor",
			args: args{clientID: "", clientSecret: ""},
			want: "/issues/v1/containers",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIssuesAPIWithCredentials(tt.args.clientID, tt.args.clientSecret).IssuesAPIPath; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIssuesAPIWithCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIssues(t *testing.T) {
	type args struct {
		path  string
		token string
	}
	tests := []struct {
		name       string
		args       args
		wantResult *IssuesContainerData
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := getIssues(tt.args.path, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIssues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("getIssues() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}