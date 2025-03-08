package xcode

import (
	"fmt"
	"strconv"
	"zhihu/pkg/xcode/types"

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
