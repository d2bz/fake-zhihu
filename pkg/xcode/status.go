package xcode

import (
	"context"
	"fmt"
	"strconv"
	"zhihu/pkg/xcode/types"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type Status struct {
	sts *types.Status
}

func ErrToStatus(err Err) *Status {
	return &Status{
		sts: &types.Status{
			Code:    int32(err.Code()),
			Message: err.Message(),
		},
	}
}

// ...表示可变参数，即可以传入多个参数
func Errorf(err Err, format string, args ...interface{}) *Status {
	err.msg = fmt.Sprintf(format, args...)
	return ErrToStatus(err)
}

func (s *Status) Error() string {
	return s.Message()
}

func (s *Status) Code() int {
	return int(s.sts.Code)
}

func (s *Status) Message() string {
	if s.sts.Message == "" {
		return strconv.Itoa(int(s.sts.Code))
	}

	return s.sts.Message
}

func (s *Status) Detail() []interface{} {
	if s == nil || s.sts == nil {
		return nil
	}

	details := make([]interface{}, 0, len(s.sts.Details))
	for _, d := range s.sts.Details {
		msg, err := d.UnmarshalNew()
		if err != nil {
			details = append(details, err)
			continue
		}
		details = append(details, msg)
	}

	return details
}

func (s *Status) WithDetails(msgs ...proto.Message) (*Status, error) {
	for _, msg := range msgs {
		anyMsg, err := anypb.New(msg)
		if err != nil {
			return s, err
		}
		s.sts.Details = append(s.sts.Details, anyMsg)
	}

	return s, nil
}

func (s *Status) Proto() *types.Status {
	return s.sts
}

func FromErr(err Err) *Status {
	return &Status{
		sts: &types.Status{
			Code:    int32(err.Code()),
			Message: err.Message(),
		},
	}
}

func FromProto(pbMsg proto.Message) Xcode {
	msg, ok := pbMsg.(*types.Status)
	if ok {
		if len(msg.Message) == 0 || msg.Message == strconv.FormatInt(int64(msg.Code), 10) {
			return Err{code: int(msg.Code)}
		}
		return &Status{sts: msg}
	}

	return Errorf(ServerErr, "invalid proto message get %v", pbMsg)
}

func ToXcode(grpcStatus *status.Status) Err {
	grpcCode := grpcStatus.Code()
	switch grpcCode {
	case codes.OK:
		return OK
	case codes.InvalidArgument:
		return RequestErr
	case codes.NotFound:
		return NotFound
	case codes.PermissionDenied:
		return AccessDenied
	case codes.Unauthenticated:
		return Unauthorized
	case codes.ResourceExhausted:
		return LimitExceeded
	case codes.Unimplemented:
		return MethodNotAllowd
	case codes.DeadlineExceeded:
		return DeadlineExceeded
	case codes.Unavailable:
		return SeviceUnavaliable
	case codes.Unknown:
		return StringCodeToErr(grpcStatus.Message())
	}

	return ServerErr
}

func CodeFromError(err error) Xcode {
	// errors.Cause()返回原始错误
	err = errors.Cause(err)
	if code, ok := err.(Xcode); ok {
		return code
	}

	switch err {
	case context.Canceled:
		return Canceled
	case context.DeadlineExceeded:
		return DeadlineExceeded
	}

	return ServerErr
}

func gRPCStatusFromXcode(code Xcode) (*status.Status, error) {
	var sts *Status
	switch v := code.(type) {
	case *Status:
		sts = v
	case Err:
		sts = FromErr(v)
	default:
		sts = ErrToStatus(Err{code.Code(), code.Message()})
		for _, detail := range code.Detail() {
			if msg, ok := detail.(proto.Message); ok {
				_, _ = sts.WithDetails(msg)
			}
		}
	}

	stas := status.New(codes.Unknown, strconv.Itoa(sts.Code()))
	return stas.WithDetails(sts.Proto())
}

func FromError(err error) *status.Status {
	err = errors.Cause(err)
	if code, ok := err.(Xcode); ok {
		grpcStatus, e := gRPCStatusFromXcode(code)
		if e == nil {
			return grpcStatus
		}
	}

	var grpcStatus *status.Status
	switch err {
	case context.Canceled:
		grpcStatus, _ = gRPCStatusFromXcode(Canceled)
	case context.DeadlineExceeded:
		grpcStatus, _ = gRPCStatusFromXcode(DeadlineExceeded)
	default:
		grpcStatus, _ = status.FromError(err)
	}

	return grpcStatus
}

func GrpcStatusToXcode(gstatus *status.Status) Xcode {
	details := gstatus.Details()
	for i := len(details) - 1; i >= 0; i-- {
		detail := details[i]
		if pb, ok := detail.(proto.Message); ok {
			return FromProto(pb)
		}
	}

	return ToXcode(gstatus)
}
