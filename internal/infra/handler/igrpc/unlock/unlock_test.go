package unlock_test

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
	"github.com/ruslanSorokin/lock-manager/internal/infra/handler/igrpc/unlock"
	"github.com/ruslanSorokin/lock-manager/internal/infra/provider"
	"github.com/ruslanSorokin/lock-manager/internal/service"
	"github.com/ruslanSorokin/lock-manager/internal/service/mock"
)

type (
	in struct {
		ctx context.Context
		req *pb.UnlockReq
	}
	out struct {
		res *pb.UnlockRes
		err error
	}
	mockIn struct {
		ctx   context.Context
		resID string
		tkn   string
	}
	mockOut struct {
		err error
	}
)

func newIn(ctx context.Context, rID, tkn *string) in {
	return in{
		ctx: ctx,
		req: &pb.UnlockReq{
			ResourceId: rID,
			Token:      tkn,
		},
	}
}

func newOut(pbErrs []pb.UnlockRes_Error, err error) out {
	return out{
		res: &pb.UnlockRes{Errors: pbErrs},
		err: err,
	}
}

func newRunner() func(unlock.Handler, in) out {
	return func(l unlock.Handler, i in) out {
		r, e := l(i.ctx, i.req)
		return out{r, e}
	}
}

func newPbErrs(errs ...pb.UnlockRes_Error) []pb.UnlockRes_Error {
	return errs
}

func TestUnlock(t *testing.T) {
	t.Parallel()
	runner := newRunner()

	validResourceID := uuid.NewString()
	invalidResourceID := "invalid resource id"
	validToken := uuid.NewString()
	invalidToken := "invalid token"

	ctx := context.TODO()

	tcs := []struct {
		desc    string
		args    func() in
		want    func() out
		mockIn  mockIn
		mockOut mockOut
		prepare func(*mock.LockService, mockIn, mockOut)
		run     func(unlock.Handler, in) out
	}{
		{
			desc: "OK",

			args: func() in {
				rID := validResourceID
				t := validToken
				return newIn(ctx, &rID, &t)
			},

			want: func() out {
				return newOut(nil, nil)
			},

			mockIn:  mockIn{ctx: ctx, resID: validResourceID, tkn: validToken},
			mockOut: mockOut{err: nil},

			prepare: func(m *mock.LockService, i mockIn, o mockOut) {
				m.On("Unlock", i.ctx, i.resID, i.tkn).Return(o.err)
			},

			run: runner,
		},
		{
			desc: "InvalidArgument ErrInvalidResourceID",

			args: func() in {
				rID := invalidResourceID
				t := validToken
				return newIn(ctx, &rID, &t)
			},

			want: func() out {
				pbErr := pb.UnlockRes_ERR_INVALID_RESOURCE_ID
				pbErrs := newPbErrs(pbErr)
				err := status.Error(codes.InvalidArgument, pbErr.String())
				return newOut(pbErrs, err)
			},

			mockIn:  mockIn{ctx: ctx, resID: invalidResourceID, tkn: validToken},
			mockOut: mockOut{err: service.ErrInvalidResourceID},

			prepare: func(m *mock.LockService, i mockIn, o mockOut) {
				m.On("Unlock", i.ctx, i.resID, i.tkn).Return(o.err)
			},

			run: runner,
		},
		{
			desc: "InvalidArgument ErrInvalidToken",

			args: func() in {
				rID := validResourceID
				t := invalidToken
				return newIn(ctx, &rID, &t)
			},

			want: func() out {
				pbErr := pb.UnlockRes_ERR_INVALID_TOKEN
				pbErrs := newPbErrs(pbErr)
				err := status.Error(codes.InvalidArgument, pbErr.String())
				return newOut(pbErrs, err)
			},

			mockIn:  mockIn{ctx: ctx, resID: validResourceID, tkn: invalidToken},
			mockOut: mockOut{err: service.ErrInvalidToken},

			prepare: func(m *mock.LockService, i mockIn, o mockOut) {
				m.On("Unlock", i.ctx, i.resID, i.tkn).Return(o.err)
			},

			run: runner,
		},
		{
			desc: "InvalidArgument ErrInvalidResourceID ErrInvalidToken",

			args: func() in {
				rID := invalidResourceID
				t := invalidToken
				return newIn(ctx, &rID, &t)
			},

			want: func() out {
				pbErrs := newPbErrs(
					pb.UnlockRes_ERR_INVALID_RESOURCE_ID, pb.UnlockRes_ERR_INVALID_TOKEN)
				err := status.Error(codes.InvalidArgument, "too many errors")
				return newOut(pbErrs, err)
			},

			mockIn: mockIn{ctx: ctx, resID: invalidResourceID, tkn: invalidToken},
			mockOut: mockOut{
				err: errors.Join(
					service.ErrInvalidResourceID, service.ErrInvalidToken),
			},

			prepare: func(m *mock.LockService, i mockIn, o mockOut) {
				m.On("Unlock", i.ctx, i.resID, i.tkn).Return(o.err)
			},

			run: runner,
		},
		{
			desc: "InvalidArgument ErrWrongToken",

			args: func() in {
				rID := validResourceID
				t := validToken
				return newIn(ctx, &rID, &t)
			},

			want: func() out {
				pbErr := pb.UnlockRes_ERR_WRONG_TOKEN
				pbErrs := newPbErrs(pbErr)
				err := status.Error(codes.InvalidArgument, pbErr.String())
				return newOut(pbErrs, err)
			},

			mockIn:  mockIn{ctx: ctx, resID: validResourceID, tkn: validToken},
			mockOut: mockOut{err: provider.ErrWrongToken},

			prepare: func(m *mock.LockService, i mockIn, o mockOut) {
				m.On("Unlock", i.ctx, i.resID, i.tkn).Return(o.err)
			},

			run: runner,
		},
		{
			desc: "NotFound ErrLockNotFound",

			args: func() in {
				rID := validResourceID
				t := validToken
				return newIn(ctx, &rID, &t)
			},

			want: func() out {
				pbErr := pb.UnlockRes_ERR_LOCK_NOT_FOUND
				pbErrs := newPbErrs(pbErr)
				err := status.Error(codes.NotFound, pbErr.String())
				return newOut(pbErrs, err)
			},

			mockIn:  mockIn{ctx: ctx, resID: validResourceID, tkn: validToken},
			mockOut: mockOut{err: provider.ErrLockNotFound},

			prepare: func(m *mock.LockService, i mockIn, o mockOut) {
				m.On("Unlock", i.ctx, i.resID, i.tkn).Return(o.err)
			},

			run: runner,
		},
		{
			desc: "Internal",

			args: func() in {
				rID := validResourceID
				t := validToken
				return newIn(ctx, &rID, &t)
			},

			want: func() out {
				err := status.Error(codes.Internal, "internal error")
				return newOut(nil, err)
			},

			mockIn:  mockIn{ctx: ctx, resID: validResourceID, tkn: validToken},
			mockOut: mockOut{err: errors.New("unexpected error")},

			prepare: func(m *mock.LockService, i mockIn, o mockOut) {
				m.On("Unlock", i.ctx, i.resID, i.tkn).Return(o.err)
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
			h := unlock.New(logr.Discard(), svc)

			tc.prepare(svc, tc.mockIn, tc.mockOut)
			got := tc.run(h, tc.args())

			require.Equal(tc.want(), got, tc.desc)
		})
	}
}
