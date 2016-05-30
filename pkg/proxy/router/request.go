// Copyright 2014 Wandoujia Inc. All Rights Reserved.
// Licensed under the MIT (MIT-LICENSE.txt) license.

package router

import (
	"sync"

	"github.com/chenyf/codis/pkg/proxy/redis"
	"github.com/chenyf/codis/pkg/utils/atomic2"
)

type Dispatcher interface {
	Dispatch(r *Request) error
}

type Request struct {
	OpStr string
	Start int64
	// Add by WangChunyan
	RedisAddr string
	Resp      *redis.Resp

	Coalesce func() error
	Response struct {
		Resp *redis.Resp
		Err  error
	}

	Wait *sync.WaitGroup
	slot *sync.WaitGroup

	Failed *atomic2.Bool
}
