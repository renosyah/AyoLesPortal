package router

import (
	template "html/template"

	"github.com/renosyah/AyoLesPortal/api"
	"github.com/renosyah/AyoLesPortal/util"
)

var (
	req                          *util.PostData
	studentModule                *api.StudentModule
	categoryModule               *api.CategoryModule
	bannerModule                 *api.BannerModule
	teacherModule                *api.TeacherModule
	courseModule                 *api.CourseModule
	courseDetailModule           *api.CourseDetailModule
	classRoomModule              *api.ClassRoomModule
	courseMaterialModule         *api.CourseMaterialModule
	courseMaterialDetailModule   *api.CourseMaterialDetailModule
	classRoomProgressModule      *api.ClassRoomProgressModule
	courseExamModule             *api.CourseExamModule
	courseExamAnswerModule       *api.CourseExamAnswerModule
	classRoomExamProgressModule  *api.ClassRoomExamProgressModule
	classRoomExamResultModule    *api.ClassRoomExamResultModule
	classRoomCertificateModule   *api.ClassRoomCertificateModule
	courseQualificationModule    *api.CourseQualificationModule
	classRoomQualificationModule *api.ClassRoomQualificationModule
	courseExamSolutionModule     *api.CourseExamSolutionModule
	temp                         *template.Template
)

func Init(r *util.PostData, temppath string) {
	req = r
	temp = template.Must(template.ParseGlob(temppath))
	studentModule = api.NewStudentModule(req)
	categoryModule = api.NewCategoryModule(req)
	bannerModule = api.NewBannerModule(req)
	teacherModule = api.NewTeacherModule(req)
	courseModule = api.NewCourseModule(req)
	courseDetailModule = api.NewCourseDetailModule(req)
	classRoomModule = api.NewClassRoomModule(req)
	courseMaterialModule = api.NewCourseMaterialModule(req)
	courseMaterialDetailModule = api.NewCourseMaterialDetailModule(req)
	classRoomProgressModule = api.NewClassRoomProgressModule(req)
	courseExamModule = api.NewCourseExamModule(req)
	courseExamAnswerModule = api.NewCourseExamAnswerModule(req)
	classRoomExamProgressModule = api.NewClassRoomExamProgressModule(req)
	classRoomExamResultModule = api.NewClassRoomExamResultModule(req)
	classRoomCertificateModule = api.NewClassRoomCertificateModule(req)
	courseQualificationModule = api.NewCourseQualificationModule(req)
	classRoomQualificationModule = api.NewClassRoomQualificationModule(req)
	courseExamSolutionModule = api.NewCourseExamSolutionModule(req)
}
