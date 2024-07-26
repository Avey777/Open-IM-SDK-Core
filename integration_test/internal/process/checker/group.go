package checker

import (
	"context"
	"github.com/openimsdk/openim-sdk-core/v3/integration_test/internal/sdk"
	"github.com/openimsdk/openim-sdk-core/v3/integration_test/internal/vars"
	"github.com/openimsdk/tools/log"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

func CheckGroupNum(ctx context.Context) (map[string]*CountChecker, error) {
	tm := time.Now()
	log.ZDebug(ctx, "checkGroupNum begin")
	defer func() {
		log.ZDebug(ctx, "checkGroupNum end", "time consuming", time.Since(tm))
	}()

	var (
		gr, _    = errgroup.WithContext(ctx)
		checkers = make(map[string]*CountChecker, len(sdk.TestSDKs))
		mapLock  = sync.RWMutex{}
	)

	gr.SetLimit(vars.ErrGroupCommonLimit)
	correctGroupNum := calCorrectGroupNum()
	for _, core := range sdk.TestSDKs {
		core := core
		gr.Go(func() error {
			_, groupNum, err := core.GetAllJoinedGroup(ctx)
			if err != nil {
				return err
			}
			isEqual := groupNum == correctGroupNum
			if !isEqual {
				mapLock.Lock()
				checkers[core.UserID] = NewCountChecker(groupNum, correctGroupNum, isEqual)
				mapLock.Unlock()
			}
			return nil
		})
	}
	if err := gr.Wait(); err != nil {
		return nil, err
	}

	return checkers, nil
}

func calCorrectGroupNum() int {
	largeNum := vars.LargeGroupNum
	commonNum := vars.CommonGroupNum * vars.CommonGroupMemberNum
	return largeNum + commonNum
}