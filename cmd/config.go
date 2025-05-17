package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/graingo/maltose/os/mcfg"
	"github.com/graingo/maltose/os/mlog"
)

// 初始化函数，注册配置钩子
func init() {
	// 注册钩子处理 MSE 数据库配置
	mcfg.RegisterAfterLoadHook(processMSEDatabaseConfig)
}

// processMSEDatabaseConfig 处理 MSE 数据库配置的钩子函数
func processMSEDatabaseConfig(ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	// 遍历所有配置项，寻找符合条件的数据库配置
	for key, value := range data {
		configMap, ok := value.(map[string]interface{})
		if !ok {
			continue
		}

		// 检查是否是带有 mseDataId 和 mseGroup 的数据库配置
		mseDataId, hasMseDataId := configMap["mseDataId"].(string)
		mseGroup, hasMseGroup := configMap["mseGroup"].(string)

		if hasMseDataId && hasMseGroup && mseGroup == "db" {
			// 记录日志
			mlog.Info(ctx, "Processing MSE database config",
				"dataId", mseDataId,
				"group", mseGroup,
				"key", key)

			// 获取 MSE 中的数据库账号密码
			username, password, err := getMSECredentials(ctx, mseDataId, mseGroup)
			if err != nil {
				mlog.Error(ctx, "Failed to get MSE credentials",
					"error", err,
					"dataId", mseDataId,
					"group", mseGroup)
				return nil, fmt.Errorf("failed to get MSE credentials: %w", err)
			}

			// 更新配置
			configMap["username"] = username
			configMap["password"] = password
		}
	}

	return data, nil
}

// getMSECredentials 从 MSE 服务获取数据库账号密码
func getMSECredentials(ctx context.Context, dataId, group string) (string, string, error) {
	// 模拟实现，实际需要调用 MSE 服务获取凭证
	// TODO: 实现实际的 MSE 服务调用

	// 这里是示例实现，真实场景需要替换
	mlog.Debug(ctx, "Getting credentials from MSE service",
		"dataId", dataId,
		"group", group)

	// 模拟网络延迟
	time.Sleep(100 * time.Millisecond)

	// 返回模拟的结果
	return "mse_username", "mse_password", nil
}
