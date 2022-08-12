package oauth

import "strings"

const (
	// 1.用户相关

	// 授权登录与用户基础信息
	ScopeUserInfo = "user_info" // 获取用户公开信息

	// 关注和粉丝列表 搜索用户并获取用户公开信息和关注列表
	ScopeFansList      = "fans.list"      // ScopeFansList 粉丝列表
	ScopeFollowingList = "following.list" // 关注列表
	ScopeFansCheck     = "fans.check"     // 检查关注

	// 2.视频相关

	// 视频查询
	//可通过接口进行视频数据的查询
	ScopeVideoList = "video.list"
	ScopeVideoData = "video.data" // 查询指定视频数据

	//发布内容至抖音：APP场景
	ScopeAwemeshare = "aweme.share" // 抖音分享id机制

	//分享给抖音好友/群
	ScopeImShare = "im.share" // 支持从第三方APP分享单图片或链接给抖音好友/群

	// 3.数据权限
	//视频数据
	ScopeDataExternaltem = "data.external.item" //用户授权后，该接口可用于查询作品的获赞，评论，分享等相关数据
	ScopeHotsearch       = "hotsearch"          // 获取实时热点词 --获取热点词聚合的视频

	//测试应用白名单权限
	ScopeTrialWhitelist = "trial.whitelist" //允许测试应用将用户添加进白名单
)

// GetUserScope 获取用户相关Scope.
func GetUserScope() string {
	scopes := []string{ScopeUserInfo, ScopeFansList, ScopeFollowingList, ScopeFansCheck}
	return strings.Join(scopes, ",")
}

// GetVideoScope 获取视频相关Scope.
func GetVideoScope() string {
	scopes := []string{ScopeVideoList, ScopeVideoData, ScopeAwemeshare, ScopeHotsearch, ScopeDataExternaltem}
	return strings.Join(scopes, ",")
}

// GetInteractScope 获取互动相关Scope.
func GetInteractScope() string {
	scopes := []string{ScopeImShare}
	return strings.Join(scopes, ",")
}

// GetAllScope 获取所有Scope.
func GetAllScope() string {
	scopes := []string{GetInteractScope(), GetVideoScope(), GetUserScope(), ScopeTrialWhitelist}
	return strings.Join(scopes, ",")
}
