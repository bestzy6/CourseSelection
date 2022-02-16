package model

type Course struct {
	CourseID  int    `gorm:"primaryKey;column:courseid"` //课程id
	Name      string `gorm:"column:name"`                //课程名称
	TeacherID int    `gorm:"column:teacherid"`           //授课教师id
	CapTotal  int    `gorm:"column:cap_total"`           //课程容量
	CapUsed   int    `gorm:"column:cap_used"`            //课程已选人数
}

func (Course) TableName() string {
	return "course"
}

// GetTeacherCourse 获取老师绑定的所有课程
func (course *Course) GetTeacherCourse() []Course {
	var allCourse []Course
	db.Where("teacherid=?", course.TeacherID).Select("courseid", "name").Find(&allCourse)
	return allCourse
}

// UnBindCourse 将课程与教师解绑
func (course *Course) UnBindCourse() error {
	err := db.Model(course).Update("teacherid", nil).Error
	return err
}

// BindCourse 将课程与教师绑定
func (course *Course) BindCourse() error {
	err := db.Model(course).Update("teacherid", course.TeacherID).Error
	return err
}

// GetCourseBindState 获取课程的绑定信息
func (course *Course) GetCourseBindState() error {
	err := db.Select("teacherid").Find(course).Error
	return err
}

// CreateCourse 创建课程
func (course *Course) CreateCourse() error {
	//插入时默认外键为空
	err := db.Debug().Omit("teacherid").Create(course).Error
	return err
}

// GetCourse 根据课程ID查找课程
func (course *Course) GetCourse() (int, error) {
	result := db.Select("courseid", "name", "teacherid").Find(course)
	return int(result.RowsAffected), result.Error
}
