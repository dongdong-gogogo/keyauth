package all

import (
	// 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载

	_ "gitee.com/dongdong-0421/keyauth/apps/user/api"

	_ "gitee.com/dongdong-0421/keyauth/apps/token/api"
)
