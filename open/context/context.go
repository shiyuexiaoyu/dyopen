package context

import (
	"github.com/shiyuexiaoyu/dyopen/open/config"
	"github.com/shiyuexiaoyu/dyopen/open/credential"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
