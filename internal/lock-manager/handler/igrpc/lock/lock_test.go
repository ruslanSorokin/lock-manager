package lock_test

import (
	"context"
	"errors"
	"testing"

	"github.com/go-logr/logr"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ruslanSorokin/lock-manager-api/gen/grpc/go"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/igrpc/lock"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/service/mock"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/testutil"
)

type (
	in struct {
		ctx context.Context
		req *pb.LockReq
	}
	out struct {
		res *pb.LockRes
		err error
	}
	mockIn struct {
		ctx   context.Context
		resID string
	}
	mockOut struct {
		tkn string
		err error
	}
)

func newIn(ctx context.Context, rID *string) in {
	return in{
		ctx: ctx,
		req: &pb.LockReq{
			ResourceId: rID,
		},
	}
}

func newOut(tkn *string, err error) out {
	return out{
		res: &pb.LockRes{Token: tkn},
		err: err,
	}
}

func newRunner() func(lock.Handler, in) out {
	return func(l lock.Handler, i in) out {
		r, e := l(i.ctx, i.req)
		return out{r, e}
	}
}

func TestLock(t *testing.T) {
	// t.Parallel()
	runner := newRunner()

	mockValidResourceID := testutil.Must(uuid.NewV4()).String()
	mockInvalidResourceID := "invalid resource id"
	mockValidToken := testutil.Must(uuid.NewV4()).String()

	ctx := context.TODO()

	tcs := []struct {
		desc    string
		args    func() in
		want    func() out
		mockIn  mockIn
		mockOut mockOut
		prepare func(*mock.LockService, mockIn, mockOut)
		run     func(lock.Handler, in) out
	}{
		{
			desc: "OK",

			args: func() in {
				rID := mockValidResourceID
				return newIn(ctx, &rID)
			},

			want: func() out {
				t := mockValidToken
				return newOut(&t, nil)
			},

			mockIn:  mockIn{ctx: ctx, resID: mockValidResourceID},
			mockOut: mockOut{tkn: mockValidToken, err: nil},

			prepare: func(m *mock.LockService, i mockIn, o mockOut) {
				m.On("Lock", i.ctx, i.resID).Return(o.tkn, o.err)
			},

			run: runner,
		},
		{
			desc: "InvalidArgument ErrInvalidResourceID",

			args: func() in {
				rID := mockInvalidResourceID
				return newIn(ctx, &rID)
			},

			want: func() out {
				err := status.Error(codes.InvalidArgument, "INVALID_RESOURCE_ID")
				return newOut(nil, err)
			},

			mockIn:  mockIn{ctx: ctx, resID: mockInvalidResourceID},
			mockOut: mockOut{tkn: "", err: service.ErrInvalidResourceID},

			prepare: func(m *mock.LockService, i mockIn, o mockOut) {
				m.On("Lock", i.ctx, i.resID).Return(o.tkn, o.err)
			},

			run: runner,
		},
		{
			desc: "AlreadyExists ErrLockAlreadyExists",

			args: func() in {
				rID := mockValidResourceID
				return newIn(ctx, &rID)
			},

			want: func() out {
				err := status.Error(codes.AlreadyExists, "RESOURCE_ALREADY_LOCKED")
				return newOut(nil, err)
			},

			mockIn:  mockIn{ctx: ctx, resID: mockValidResourceID},
			mockOut: mockOut{tkn: "", err: provider.ErrLockAlreadyExists},

			prepare: func(m *mock.LockService, i mockIn, o mockOut) {
				m.On("Lock", i.ctx, i.resID).Return(o.tkn, o.err)
			},

			run: runner,
		},
		{
			desc: "Internal",

			args: func() in {
				rID := mockValidResourceID
				return newIn(ctx, &rID)
			},

			want: func() out {
				err := status.Error(codes.Internal, "INTERNAL_ERROR")
				return newOut(nil, err)
			},

			mockIn:  mockIn{ctx: ctx, resID: mockValidResourceID},
			mockOut: mockOut{tkn: "", err: errors.New("unexpected error")},

			prepare: func(m *mock.LockService, i mockIn, o mockOut) {
				m.On("Lock", i.ctx, i.resID).Return(o.tkn, o.err)
			},

			run: runner,
		},
	}
	require := require.New(t)
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			// t.Parallel()

			svc := mock.NewLockService(t)
			h := lock.New(logr.Discard(), svc)

			tc.prepare(svc, tc.mockIn, tc.mockOut)
			got := tc.run(h, tc.args())

			require.Equal(tc.want(), got, tc.desc)
		})
	}
}
