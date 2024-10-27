package rpcserver

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	zerr "github.com/zeromicro/x/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LogInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any,
	err error) {
	resp, err = handler(ctx, req)
	if err == nil {
		return resp, nil
	}

	// 如果发生了错误，记录错误日志。
	logx.WithContext(ctx).Errorf("【RPC SRV ERR】 %v", err)

	// 获取错误的根本原因。
	causeErr := errors.Cause(err)
	// 检查根本原因是否是一个自定义的错误类型CodeMsg。
	if e, ok := causeErr.(*zerr.CodeMsg); ok {
		// 如果是，根据自定义错误类型中的Code和Msg创建一个新的gRPC错误，并用其替换原始错误。
		err = status.Error(codes.Code(e.Code), e.Msg)
	}

	return resp, err
}
