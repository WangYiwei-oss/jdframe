package casbin

import (
	"github.com/WangYiwei-oss/jdframe/src/configs"
	"github.com/WangYiwei-oss/jdframe/src/models"
)

type RoleRel struct {
	PRole string //父权限对象
	Role  string //权限对象
}

func (r *RoleRel) String() string {
	return r.PRole + ":" + r.Role
}

func GetRoles(pid int, m *[]*RoleRel, pname string) {
	proles := make([]*models.Role, 0)
	configs.NewGormAdapter().Where("role_pid=?", pid).Find(&proles)
	if len(proles) == 0 {
		return
	}
	for _, item := range proles {
		if pname != "" {
			*m = append(*m, &RoleRel{pname, item.RoleName})
		}
		GetRoles(item.RoleId, m, item.RoleName)
	}
}

// GetUserRoles 获取用户和角色的对应关系
func GetUserRoles() (users []models.User) {
	configs.NewGormAdapter().Select("a.user_name,c.role_name").Table("users a,user_roles b,roles c").
		Where("a.user_id=b.user_id and b.role_id=c.role_id").Order("a.user_id desc").
		Find(&users)
	return
}

func GetRouterRoles() (routers []*models.Router) {
	configs.NewGormAdapter().Select("a.r_uri,a.r_method,c.role_name").Table("routers a,router_roles b,roles c").
		Where("a.r_id=b.router_id and b.role_id=c.role_id").
		Order("role_name").Find(&routers)
	return
}
