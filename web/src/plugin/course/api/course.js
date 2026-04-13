import service from '@/utils/request'

/**
 * 获取课程列表
 * @param {Object} data 查询参数
 * @param {number} data.page 页码
 * @param {number} data.pageSize 每页数量
 * @param {string} data.keyword 关键词
 * @param {string} data.category 分类
 * @param {string} data.level 难度级别
 * @returns {Promise} 课程列表数据
 */
export const getCourseList = (data) => {
  return service({
    url: '/course/getCourseList',
    method: 'post',
    data: data
  })
}

/**
 * 获取课程详情
 * @param {number} id 课程ID
 * @returns {Promise} 课程详情数据
 */
export const getCourseDetail = (id) => {
  return service({
    url: `/course/getCourseDetail/${id}`,
    method: 'get'
  })
}

/**
 * 创建课程
 * @param {Object} data 课程数据
 * @param {string} data.title 课程标题
 * @param {string} data.description 课程描述
 * @param {number} data.price 课程价格
 * @param {string} data.level 难度级别
 * @param {string} data.category 课程分类
 * @param {string} data.cover_image 封面图片
 * @param {string} data.intro_video 介绍视频
 * @returns {Promise} 创建结果
 */
export const createCourse = (data) => {
  return service({
    url: '/course/createCourse',
    method: 'post',
    data: data
  })
}

/**
 * 更新课程
 * @param {Object} data 课程数据
 * @param {number} data.ID 课程ID
 * @param {string} data.title 课程标题
 * @param {string} data.description 课程描述
 * @param {number} data.price 课程价格
 * @param {string} data.level 难度级别
 * @param {string} data.category 课程分类
 * @param {string} data.cover_image 封面图片
 * @param {string} data.intro_video 介绍视频
 * @returns {Promise} 更新结果
 */
export const updateCourse = (data) => {
  return service({
    url: '/course/updateCourse',
    method: 'put',
    data: data
  })
}

/**
 * 删除课程
 * @param {number} id 课程ID
 * @returns {Promise} 删除结果
 */
export const deleteCourse = (id) => {
  return service({
    url: `/course/deleteCourse/${id}`,
    method: 'delete'
  })
}

/**
 * 创建章节
 * @param {Object} data 章节数据
 * @param {number} data.course_id 课程ID
 * @param {string} data.title 章节标题
 * @param {number} data.order 排序
 * @returns {Promise} 创建结果
 */
export const createChapter = (data) => {
  return service({
    url: '/course/createChapter',
    method: 'post',
    data: data
  })
}

/**
 * 更新章节
 * @param {Object} data 章节数据
 * @param {number} data.ID 章节ID
 * @param {string} data.title 章节标题
 * @param {number} data.order 排序
 * @returns {Promise} 更新结果
 */
export const updateChapter = (data) => {
  return service({
    url: '/course/updateChapter',
    method: 'put',
    data: data
  })
}

/**
 * 删除章节
 * @param {number} id 章节ID
 * @returns {Promise} 删除结果
 */
export const deleteChapter = (id) => {
  return service({
    url: `/course/deleteChapter/${id}`,
    method: 'delete'
  })
}

/**
 * 创建课时
 * @param {Object} data 课时数据
 * @param {number} data.chapter_id 章节ID
 * @param {string} data.title 课时标题
 * @param {string} data.video_url 视频URL
 * @param {number} data.duration 时长
 * @param {string} data.content 内容
 * @param {number} data.order 排序
 * @returns {Promise} 创建结果
 */
export const createLesson = (data) => {
  return service({
    url: '/course/createLesson',
    method: 'post',
    data: data
  })
}

/**
 * 更新课时
 * @param {Object} data 课时数据
 * @param {number} data.ID 课时ID
 * @param {string} data.title 课时标题
 * @param {string} data.video_url 视频URL
 * @param {number} data.duration 时长
 * @param {string} data.content 内容
 * @param {number} data.order 排序
 * @returns {Promise} 更新结果
 */
export const updateLesson = (data) => {
  return service({
    url: '/course/updateLesson',
    method: 'put',
    data: data
  })
}

/**
 * 删除课时
 * @param {number} id 课时ID
 * @returns {Promise} 删除结果
 */
export const deleteLesson = (id) => {
  return service({
    url: `/course/deleteLesson/${id}`,
    method: 'delete'
  })
}

/**
 * 报名课程
 * @param {Object} data 报名数据
 * @param {number} data.course_id 课程ID
 * @returns {Promise} 报名结果
 */
export const enrollCourse = (data) => {
  return service({
    url: '/course/enrollCourse',
    method: 'post',
    data: data
  })
}

/**
 * 获取用户已报名课程
 * @param {Object} data 查询参数
 * @param {number} data.page 页码
 * @param {number} data.pageSize 每页数量
 * @returns {Promise} 已报名课程列表
 */
export const getEnrolledCourses = (data) => {
  return service({
    url: '/course/getEnrolledCourses',
    method: 'post',
    data: data
  })
}

/**
 * 取消报名
 * @param {number} id 报名ID
 * @returns {Promise} 取消报名结果
 */
export const cancelEnrollment = (id) => {
  return service({
    url: `/course/cancelEnrollment/${id}`,
    method: 'delete'
  })
}

/**
 * 更新学习进度
 * @param {Object} data 进度数据
 * @param {number} data.user_id 用户ID
 * @param {number} data.course_id 课程ID
 * @param {number} data.lesson_id 课时ID
 * @param {number} data.watched_duration 已观看时长
 * @param {boolean} data.is_completed 是否完成
 * @returns {Promise} 更新结果
 */
export const updateProgress = (data) => {
  return service({
    url: '/course/updateProgress',
    method: 'put',
    data: data
  })
}

/**
 * 获取学习进度
 * @param {number} courseId 课程ID
 * @param {number} lessonId 课时ID
 * @returns {Promise} 学习进度数据
 */
export const getProgress = (courseId, lessonId) => {
  return service({
    url: `/course/getProgress/${courseId}/${lessonId}`,
    method: 'get'
  })
}

/**
 * 创建作业
 * @param {Object} data 作业数据
 * @param {number} data.course_id 课程ID
 * @param {string} data.title 作业标题
 * @param {string} data.description 作业描述
 * @param {string} data.deadline 截止时间
 * @returns {Promise} 创建结果
 */
export const createHomework = (data) => {
  return service({
    url: '/course/createHomework',
    method: 'post',
    data: data
  })
}

/**
 * 更新作业
 * @param {Object} data 作业数据
 * @param {number} data.ID 作业ID
 * @param {string} data.title 作业标题
 * @param {string} data.description 作业描述
 * @param {string} data.deadline 截止时间
 * @returns {Promise} 更新结果
 */
export const updateHomework = (data) => {
  return service({
    url: '/course/updateHomework',
    method: 'put',
    data: data
  })
}

/**
 * 删除作业
 * @param {number} id 作业ID
 * @returns {Promise} 删除结果
 */
export const deleteHomework = (id) => {
  return service({
    url: `/course/deleteHomework/${id}`,
    method: 'delete'
  })
}

/**
 * 提交作业
 * @param {Object} data 提交数据
 * @param {number} data.homework_id 作业ID
 * @param {string} data.content 作业内容
 * @param {string} data.file_url 附件URL
 * @returns {Promise} 提交结果
 */
export const submitHomework = (data) => {
  return service({
    url: '/course/submitHomework',
    method: 'post',
    data: data
  })
}

/**
 * 批改作业
 * @param {Object} data 批改数据
 * @param {number} data.ID 提交ID
 * @param {number} data.grade 分数
 * @param {string} data.feedback 反馈
 * @param {string} data.status 状态
 * @returns {Promise} 批改结果
 */
export const gradeHomework = (data) => {
  return service({
    url: '/course/gradeHomework',
    method: 'put',
    data: data
  })
}

/**
 * 获取作业列表
 * @param {number} courseId 课程ID
 * @param {Object} data 查询参数
 * @param {number} data.page 页码
 * @param {number} data.pageSize 每页数量
 * @returns {Promise} 作业列表
 */
export const getHomeworkList = (courseId, data) => {
  return service({
    url: `/course/getHomeworkList/${courseId}`,
    method: 'post',
    data: data
  })
}

/**
 * 生成证书
 * @param {number} courseId 课程ID
 * @returns {Promise} 证书生成结果
 */
export const generateCertificate = (courseId) => {
  return service({
    url: `/course/generateCertificate/${courseId}`,
    method: 'post'
  })
}

/**
 * 获取证书列表
 * @param {Object} data 查询参数
 * @param {number} data.page 页码
 * @param {number} data.pageSize 每页数量
 * @returns {Promise} 证书列表
 */
export const getCertificateList = (data) => {
  return service({
    url: '/course/getCertificateList',
    method: 'post',
    data: data
  })
}

/**
 * 创建评论
 * @param {Object} data 评论数据
 * @param {number} data.course_id 课程ID
 * @param {string} data.content 评论内容
 * @param {number} data.rating 评分
 * @param {number} data.parent_id 父评论ID
 * @returns {Promise} 创建结果
 */
export const createComment = (data) => {
  return service({
    url: '/course/createComment',
    method: 'post',
    data: data
  })
}

/**
 * 获取评论列表
 * @param {number} courseId 课程ID
 * @param {Object} data 查询参数
 * @param {number} data.page 页码
 * @param {number} data.pageSize 每页数量
 * @returns {Promise} 评论列表
 */
export const getCommentList = (courseId, data) => {
  return service({
    url: `/course/getCommentList/${courseId}`,
    method: 'post',
    data: data
  })
}

/**
 * 删除评论
 * @param {number} id 评论ID
 * @returns {Promise} 删除结果
 */
export const deleteComment = (id) => {
  return service({
    url: `/course/deleteComment/${id}`,
    method: 'delete'
  })
}
