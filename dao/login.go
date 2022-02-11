package dao

//用户已删除，用户不存在，密码错误

//func Login(member *model.Member) (id int, err error) {
//	find := model.DB.Where("username=? AND password=?", member.Username, member.Password).Find(&member)
//	err = find.Error
//	if err != nil {
//		return
//	}
//	affected := find.RowsAffected
//	if affected > 0 {
//		//如果用户已删除
//		if member.State {
//			err = errors.New("用户已删除")
//			return
//		}
//
//	} else {
//		err = errors.New("不存在此用户")
//	}
//	return
//}
