CREATE INDEX idx_students_dept ON students(department_id);
CREATE INDEX idx_teachers_dept ON teachers(department_id);
CREATE INDEX idx_courses_dept ON courses(department_id);

CREATE INDEX idx_enrollments_student ON enrollments(student_id);
CREATE INDEX idx_enrollments_course ON enrollments(course_id);

CREATE INDEX idx_exam_marks_student ON exam_marks(student_id);
CREATE INDEX idx_attendance_student ON attendance(student_id);
CREATE INDEX idx_attendance_date ON attendance(date);

CREATE INDEX idx_students_name ON students(name);
