package unlock_test

import (
	"context"
	"errors"
	"testing"

	"github.com/go-logr/logr"
	"github.com/gofrs/uuid/v5"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ruslanSorokin/lock-manager-api/gen/grpc/go"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/igrpc/unlock"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
	servicemock "github.com/ruslanSorokin/lock-manager/internal/lock-manager/service/mock"
	util "github.com/ruslanSorokin/lock-manager/internal/pkg/util"
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

func newOut(err error) out {
	return out{
		res: &pb.UnlockRes{},
		err: err,
	}
}

func newRunner() func(unlock.Handler, in) out {
	return func(l unlock.Handler, i in) out {
		r, e := l(i.ctx, i.req)
		return out{r, e}
	}
}

func TestUnlock(t *testing.T) {
	// t.Parallel()
	runner := newRunner()

	mockValidResourceID := util.Must(uuid.NewV4()).String()
	mockInvalidResourceID := "invalid resource id"
	mockValidToken := util.Must(uuid.NewV4()).String()
	mockInvalidToken := "invalid token"

	ctx := context.TODO()

	tcs := []struct {
		desc    string
		args    func() in
		want    func() out
		mockIn  mockIn
		mockOut mockOut
		prepare func(*servicemock.LockService, mockIn, mockOut)
		run     func(unlock.Handler, in) out
	}{
		{
			desc: "OK",

			args: func() in {
				rID := mockValidResourceID
				t := mockValidToken
				return newIn(ctx, &rID, &t)
			},

			want: func() out {
				return newOut(nil)
			},

			mockIn: mockIn{
				ctx:   ctx,
				resID: mockValidResourceID,
				tkn:   mockValidToken,
			},
			mockOut: mockOut{err: nil},

			prepare: func(m *servicemock.LockService, i mockIn, o mockOut) {
				m.On("Unlock", i.ctx, i.resID, i.tkn).Return(o.err)
			},

			run: runner,
		},
		{
			desc: "InvalidArgument ErrInvalidResourceID",

			args: func() in {
				rID := mockInvalidResourceID
				t := mockValidToken
				return newIn(ctx, &rID, &t)
			},

			want: func() out {
				err := status.Error(codes.InvalidArgument, "INVALID_RESOURCE_ID")
				return newOut(err)
			},

			mockIn: mockIn{
				ctx:   ctx,
				resID: mockInvalidResourceID,
				tkn:   mockValidToken,
			},
			mockOut: mockOut{err: service.ErrInvalidResourceID},

			prepare: func(m *servicemock.LockService, i mockIn, o mockOut) {
				m.On("Unlock", i.ctx, i.resID, i.tkn).Return(o.err)
			},

			run: runner,
		},
		{
			desc: "InvalidArgument ErrInvalidToken",

			args: func() in {
				rID := mockValidResourceID
				t := mockInvalidToken
				return newIn(ctx, &rID, &t)
			},

			want: func() out {
				err := status.Error(codes.InvalidArgument, "INVALID_TOKEN")
				return newOut(err)
			},

			mockIn: mockIn{
				ctx:   ctx,
				resID: mockValidResourceID,
				tkn:   mockInvalidToken,
			},
			mockOut: mockOut{err: service.ErrInvalidToken},

			prepare: func(m *servicemock.LockService, i mockIn, o mockOut) {
				m.On("Unlock", i.ctx, i.resID, i.tkn).Return(o.err)
			},

			run: runner,
		},
		{
			desc: "InvalidArgument ErrInvalidResourceID ErrInvalidToken",

			args: func() in {
				rID := mockInvalidResourceID
				t := mockInvalidToken
				return newIn(ctx, &rID, &t)
			},

			want: func() out {
				err := status.Error(codes.InvalidArgument, "INVALID_RESOURCE_ID")
				return newOut(err)
			},

			mockIn: mockIn{
				ctx:   ctx,
				resID: mockInvalidResourceID,
				tkn:   mockInvalidToken,
			},
			mockOut: mockOut{
				err: errors.Join(
					service.ErrInvalidResourceID, service.ErrInvalidToken),
			},

			prepare: func(m *servicemock.LockService, i mockIn, o mockOut) {
				m.On("Unlock", i.ctx, i.resID, i.tkn).Return(o.err)
			},

			run: runner,
		},
		{
			desc: "InvalidArgument ErrWrongToken",

			args: func() in {
				rID := mockValidResourceID
				t := mockValidToken
				return newIn(ctx, &rID, &t)
			},

			want: func() out {
				err := status.Error(codes.InvalidArgument, "TOKEN_DOES_NOT_FIT")
				return newOut(err)
			},

			mockIn: mockIn{
				ctx:   ctx,
				resID: mockValidResourceID,
				tkn:   mockValidToken,
			},
			mockOut: mockOut{err: provider.ErrWrongToken},

			prepare: func(m *servicemock.LockService, i mockIn, o mockOut) {
				m.On("Unlock", i.ctx, i.resID, i.tkn).Return(o.err)
			},

			run: runner,
		},
		{
			desc: "NotFound ErrLockNotFound",

			args: func() in {
				rID := mockValidResourceID
				t := mockValidToken
				return newIn(ctx, &rID, &t)
			},

			want: func() out {
				err := status.Error(codes.NotFound, "LOCK_NOT_FOUND")
				return newOut(err)
			},

			mockIn: mockIn{
				ctx:   ctx,
				resID: mockValidResourceID,
				tkn:   mockValidToken,
			},
			mockOut: mockOut{err: provider.ErrLockNotFound},

			prepare: func(m *servicemock.LockService, i mockIn, o mockOut) {
				m.On("Unlock", i.ctx, i.resID, i.tkn).Return(o.err)
			},

			run: runner,
		},
		{
			desc: "Internal",

			args: func() in {
				rID := mockValidResourceID
				t := mockValidToken
				return newIn(ctx, &rID, &t)
			},

			want: func() out {
				err := status.Error(codes.Internal, "INTERNAL_ERROR")
				return newOut(err)
			},

			mockIn: mockIn{
				ctx:   ctx,
				resID: mockValidResourceID,
				tkn:   mockValidToken,
			},
			mockOut: mockOut{err: errors.New("unexpected error")},

			prepare: func(m *servicemock.LockService, i mockIn, o mockOut) {
				m.On("Unlock", i.ctx, i.resID, i.tkn).Return(o.err)
			},

			run: runner,
		},
	}
	require := require.New(t)

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			// t.Parallel()

			svc := servicemock.NewLockService(t)
			h := unlock.New(logr.Discard(), svc)

			tc.prepare(svc, tc.mockIn, tc.mockOut)
			got := tc.run(h, tc.args())

			require.Equal(tc.want(), got, tc.desc)
		})
	}
}
