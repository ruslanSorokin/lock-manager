package lock_test

import (
	"context"
	"errors"
	"testing"

	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ruslanSorokin/lock-manager-api/gen/grpc/go"
	"github.com/ruslanSorokin/lock-manager/internal/infra/handler/igrpc/lock"
	"github.com/ruslanSorokin/lock-manager/internal/infra/provider"
	"github.com/ruslanSorokin/lock-manager/internal/service"
	"github.com/ruslanSorokin/lock-manager/internal/service/mock"
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

func newOut(pbErrs []pb.LockRes_Error, tkn *string, err error) out {
	return out{
		res: &pb.LockRes{Errors: pbErrs, Token: tkn},
		err: err,
	}
}

func newRunner() func(lock.Handler, in) out {
	return func(l lock.Handler, i in) out {
		r, e := l(i.ctx, i.req)
		return out{r, e}
	}
}

func newPbErrs(errs ...pb.LockRes_Error) []pb.LockRes_Error {
	return errs
}

func TestLock(t *testing.T) {
	t.Parallel()
	runner := newRunner()

	validResourceID := uuid.NewString()
	invalidResourceID := "invalid resource id"
	validToken := uuid.NewString()
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
				rID := validResourceID
				return newIn(ctx, &rID)
			},

			want: func() out {
				t := validToken
				return newOut(nil, &t, nil)
			},

			mockIn:  mockIn{ctx: ctx, resID: validResourceID},
			mockOut: mockOut{tkn: validToken, err: nil},

			prepare: func(m *mock.LockService, i mockIn, o mockOut) {
				m.On("Lock", i.ctx, i.resID).Return(o.tkn, o.err)
			},

			run: runner,
		},
		{
			desc: "InvalidArgument ErrInvalidResourceID",

			args: func() in {
				rID := invalidResourceID
				return newIn(ctx, &rID)
			},

			want: func() out {
				pbErr := pb.LockRes_ERR_INVALID_RESOURCE_ID
				pbErrs := newPbErrs(pbErr)
				err := status.Error(codes.InvalidArgument, pbErr.String())
				return newOut(pbErrs, nil, err)
			},

			mockIn:  mockIn{ctx: ctx, resID: invalidResourceID},
			mockOut: mockOut{tkn: "", err: service.ErrInvalidResourceID},

			prepare: func(m *mock.LockService, i mockIn, o mockOut) {
				m.On("Lock", i.ctx, i.resID).Return(o.tkn, o.err)
			},

			run: runner,
		},
		{
			desc: "AlreadyExists ErrLockAlreadyExists",

			args: func() in {
				rID := validResourceID
				return newIn(ctx, &rID)
			},

			want: func() out {
				pbErr := pb.LockRes_ERR_LOCK_ALREADY_EXISTS
				pbErrs := newPbErrs(pbErr)
				err := status.Error(codes.AlreadyExists, pbErr.String())
				return newOut(pbErrs, nil, err)
			},

			mockIn:  mockIn{ctx: ctx, resID: validResourceID},
			mockOut: mockOut{tkn: "", err: provider.ErrLockAlreadyExists},

			prepare: func(m *mock.LockService, i mockIn, o mockOut) {
				m.On("Lock", i.ctx, i.resID).Return(o.tkn, o.err)
			},

			run: runner,
		},
		{
			desc: "Internal",

			args: func() in {
				rID := validResourceID
				return newIn(ctx, &rID)
			},

			want: func() out {
				err := status.Error(codes.Internal, "internal error")
				return newOut(nil, nil, err)
			},

			mockIn:  mockIn{ctx: ctx, resID: validResourceID},
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
			t.Parallel()

			svc := mock.NewLockService(t)
			h := lock.New(logr.Discard(), svc)

			tc.prepare(svc, tc.mockIn, tc.mockOut)
			got := tc.run(h, tc.args())

			require.Equal(tc.want(), got, tc.desc)
		})
	}
}
