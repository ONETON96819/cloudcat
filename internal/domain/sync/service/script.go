package service

import (
	"github.com/scriptscat/cloudcat/internal/domain/sync/dto"
	"github.com/scriptscat/cloudcat/internal/domain/sync/errs"
	"github.com/scriptscat/cloudcat/internal/domain/sync/repository"
	"github.com/scriptscat/cloudcat/internal/pkg/cnt"
	"github.com/sirupsen/logrus"
)

/*
	脚本同步使用一个版本号标记脚本的变动,使用redis zset存储,另外使用一个key,每次incr来生成版本号.
	数据库中只存储最新的脚本数据,zset中只存储脚本删改的记录.
	前端push之前先使用pull命令查询当前数据是否为最新的,处理pull数据合并之后再进行push操作.
	如果pull与push的脚本有冲突,以脚本的最新时间为合并依据.pull>push不进行推送,并覆盖本地记录.
	调用pull接口时,后端使用前端提交的version,使用zrangebyscore从zset中取出version到最新之间的记录,合并处理脚本的一个最终状态返回给前端
	调用push接口时,先对比前端传递的version是否为最新,不是最新禁止前端push.push成功返回一个version给前端记录起来.
*/

type Sync interface {
	PushScript(user, device, version int64, scripts []*dto.SyncScript) ([]*dto.SyncScript, int64, error)
	PullScript(user, device, version int64) ([]*dto.SyncScript, int64, error)

	PushSubscribe(user, device int64, version int64, sub []*dto.SyncSubscribe) ([]*dto.SyncSubscribe, int64, error)
	PullSubscribe(user, device int64, version int64) ([]*dto.SyncSubscribe, int64, error)

	SyncSetting(user, device int64) error
}

type sync struct {
	script    repository.Script
	subscribe repository.Subscribe
}

func NewSync(script repository.Script, subscribe repository.Subscribe) Sync {
	return &sync{
		script:    script,
		subscribe: subscribe,
	}
}

func (s *sync) PushScript(user, device, version int64, scripts []*dto.SyncScript) ([]*dto.SyncScript, int64, error) {
	if len(scripts) == 0 {
		return nil, 0, errs.ErrSyncIsNil
	}
	if v, err := s.script.LatestVersion(user, device); err != nil {
		return nil, 0, err
	} else if v != version {
		return nil, 0, errs.ErrSyncVersionError
	}
	data := make([]*dto.SyncScript, 0)
	ret := make([]*dto.SyncScript, len(scripts))
	for i, v := range scripts {
		script, err := s.script.FindByUUID(user, device, v.UUID)
		if err != nil {
			logrus.Warnf("push script find script error: %v", err)
			ret[i] = &dto.SyncScript{
				Action: "error",
				Msg:    "同步失败,系统错误",
			}
			continue
		}
		// 时间小或者为空,更新脚本
		v.Script.UserID = user
		v.Script.DeviceID = device
		v.Script.Updatetime = v.Actiontime
		if v.Action == "uninstall" {
			v.Script.Status = cnt.DELETE
		} else {
			v.Script.Status = cnt.ACTIVE
		}
		if script != nil {
			v.Script.ID = script.ID
			v.Script.Createtime = script.Createtime
		}
		if err := s.script.Save(v.Script); err != nil {
			logrus.Warnf("sync script save script error: %v", err)
			ret[i] = &dto.SyncScript{
				Action: "error",
				Msg:    "同步失败,系统错误",
			}
			continue
		}
		ret[i] = &dto.SyncScript{
			Action: "ok",
			Script: v.Script,
		}
		data = append(data, &dto.SyncScript{
			Action:     v.Action,
			Actiontime: v.Actiontime,
			UUID:       v.UUID,
		})
	}
	version, err := s.script.PushVersion(user, device, data)
	if err != nil {
		return nil, 0, err
	}
	return ret, version, nil
}

func (s *sync) PullScript(user, device, version int64) ([]*dto.SyncScript, int64, error) {
	latest, err := s.script.LatestVersion(user, device)
	if err != nil {
		return nil, 0, err
	}
	if latest == version {
		return nil, latest, nil
	}
	list, err := s.script.ActionList(user, device, version)
	if err != nil {
		return nil, 0, err
	}
	ret := make([]*dto.SyncScript, 0)
	unique := make(map[string]*dto.SyncScript, 0)
	for _, v := range list {
		for _, v := range v {
			unique[v.UUID] = v
		}
	}
	for _, v := range unique {
		switch v.Action {
		case "install", "reinstall":
			script, err := s.script.FindByUUID(user, device, v.UUID)
			if err != nil {
				logrus.Warnf("pull script find script error: %v", err)
				v.Action = "error"
				v.Msg = "同步失败,系统错误"
			} else {
				v.Script = script
			}
		}
		ret = append(ret, v)
	}
	return ret, latest, nil
}

func (s *sync) PushSubscribe(user, device, version int64, sub []*dto.SyncSubscribe) ([]*dto.SyncSubscribe, int64, error) {
	if len(sub) == 0 {
		return nil, 0, errs.ErrSyncIsNil
	}
	if v, err := s.subscribe.LatestVersion(user, device); err != nil {
		return nil, 0, err
	} else if v != version {
		return nil, 0, errs.ErrSyncVersionError
	}
	data := make([]*dto.SyncSubscribe, 0)
	ret := make([]*dto.SyncSubscribe, len(sub))
	for i, v := range sub {
		script, err := s.subscribe.FindByUrl(user, device, v.URL)
		if err != nil {
			logrus.Warnf("push subscribe find subscribe error: %v", err)
			ret[i] = &dto.SyncSubscribe{
				Action: "error",
				Msg:    "同步失败,系统错误",
			}
			continue
		}
		// 时间小或者为空,更新脚本
		v.Subscribe.UserID = user
		v.Subscribe.DeviceID = device
		v.Subscribe.Updatetime = v.Actiontime
		if v.Action == "unsubscribe" {
			v.Subscribe.Status = cnt.DELETE
		} else {
			v.Subscribe.Status = cnt.ACTIVE
		}
		if script != nil {
			v.Subscribe.ID = script.ID
			v.Subscribe.Createtime = script.Createtime
		}
		if err := s.subscribe.Save(v.Subscribe); err != nil {
			logrus.Warnf("sync subscribe save subscribe error: %v", err)
			ret[i] = &dto.SyncSubscribe{
				Action: "error",
				Msg:    "同步失败,系统错误",
			}
			continue
		}
		ret[i] = &dto.SyncSubscribe{
			Action:    "ok",
			Subscribe: v.Subscribe,
		}
		data = append(data, &dto.SyncSubscribe{
			Action:     v.Action,
			Actiontime: v.Actiontime,
			URL:        v.URL,
		})
	}
	version, err := s.subscribe.PushVersion(user, device, data)
	if err != nil {
		return nil, 0, err
	}
	return ret, version, nil
}

func (s *sync) PullSubscribe(user, device int64, version int64) ([]*dto.SyncSubscribe, int64, error) {
	latest, err := s.subscribe.LatestVersion(user, device)
	if err != nil {
		return nil, 0, err
	}
	if latest == version {
		return nil, latest, nil
	}
	list, err := s.subscribe.ActionList(user, device, version)
	if err != nil {
		return nil, 0, err
	}
	ret := make([]*dto.SyncSubscribe, 0)
	unique := make(map[string]*dto.SyncSubscribe, 0)
	for _, v := range list {
		for _, v := range v {
			unique[v.URL] = v
		}
	}
	for _, v := range unique {
		switch v.Action {
		case "subscribe", "resubscribe":
			subscribe, err := s.subscribe.FindByUrl(user, device, v.URL)
			if err != nil {
				logrus.Warnf("pull subscribe find subscribe error: %v", err)
				v.Action = "error"
				v.Msg = "同步失败,系统错误"
			} else {
				v.Subscribe = subscribe
			}
		}
		ret = append(ret, v)
	}
	return ret, latest, nil
}

func (s *sync) SyncSetting(user, device int64) error {
	panic("implement me")
}