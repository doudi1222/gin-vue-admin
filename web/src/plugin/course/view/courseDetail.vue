<template>
  <div class="course-detail">
    <el-card shadow="never" class="header-card">
      <div class="course-header">
        <div class="cover-image">
          <img :src="courseDetail.cover_image" alt="课程封面" v-if="courseDetail.cover_image" />
          <div class="placeholder" v-else>暂无封面</div>
        </div>
        <div class="course-info">
          <h2 class="course-title">{{ courseDetail.title }}</h2>
          <div class="course-meta">
            <el-tag>{{ getCategoryText(courseDetail.category) }}</el-tag>
            <el-tag :type="getLevelType(courseDetail.level)">{{ getLevelText(courseDetail.level) }}</el-tag>
            <span class="student-count">{{ courseDetail.student_count }} 名学生</span>
          </div>
          <div class="course-price">
            <span class="price">¥{{ courseDetail.price }}</span>
            <el-button type="primary" @click="handleEnroll" v-if="!isEnrolled">立即报名</el-button>
            <el-button type="success" @click="handleLearn" v-else>开始学习</el-button>
          </div>
          <div class="course-description">
            <p>{{ courseDetail.description }}</p>
          </div>
        </div>
      </div>
    </el-card>

    <el-card shadow="never" class="content-card">
      <template #header>
        <div class="card-header">
          <el-tabs v-model="activeTab">
            <el-tab-pane label="课程内容" name="content" />
            <el-tab-pane label="作业管理" name="homework" />
            <el-tab-pane label="学员评论" name="comment" />
          </el-tabs>
        </div>
      </template>

      <!-- 课程内容 -->
      <div v-if="activeTab === 'content'" class="course-content">
        <div v-for="chapter in courseDetail.chapters" :key="chapter.ID" class="chapter">
          <div class="chapter-header" @click="toggleChapter(chapter.ID)">
            <el-icon :class="{ 'rotated': expandedChapters.includes(chapter.ID) }">
              <ArrowRight />
            </el-icon>
            <span class="chapter-title">{{ chapter.title }}</span>
            <span class="chapter-lesson-count">({{ chapter.lessons.length }} 课时)</span>
          </div>
          <div class="chapter-lessons" v-show="expandedChapters.includes(chapter.ID)">
            <div v-for="lesson in chapter.lessons" :key="lesson.ID" class="lesson-item">
              <el-icon class="lesson-icon"><VideoPlay /></el-icon>
              <span class="lesson-title">{{ lesson.title }}</span>
              <span class="lesson-duration">{{ formatDuration(lesson.duration) }}</span>
              <el-button type="primary" size="small" @click="playLesson(lesson)">播放</el-button>
            </div>
          </div>
        </div>
      </div>

      <!-- 作业管理 -->
      <div v-if="activeTab === 'homework'" class="homework-section">
        <el-button type="primary" @click="handleCreateHomework" plain>
          <el-icon><Plus /></el-icon> 新建作业
        </el-button>
        <el-table
          v-loading="homeworkLoading"
          :data="homeworkList"
          style="width: 100%"
          border
          class="homework-table"
        >
          <el-table-column prop="title" label="作业标题" min-width="200" />
          <el-table-column prop="description" label="作业描述" min-width="300" />
          <el-table-column prop="deadline" label="截止时间" width="180" />
          <el-table-column label="操作" width="200" fixed="right">
            <template #default="scope">
              <el-button type="primary" size="small" @click="handleEditHomework(scope.row)" plain>
                <el-icon><Edit /></el-icon> 编辑
              </el-button>
              <el-button type="danger" size="small" @click="handleDeleteHomework(scope.row.ID)" plain>
                <el-icon><Delete /></el-icon> 删除
              </el-button>
              <el-button type="info" size="small" @click="handleViewSubmissions(scope.row.ID)" plain>
                <el-icon><View /></el-icon> 查看提交
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>

      <!-- 学员评论 -->
      <div v-if="activeTab === 'comment'" class="comment-section">
        <el-card shadow="never" class="comment-form-card">
          <el-form :model="commentForm" class="comment-form">
            <el-form-item label="评分">
              <el-rate v-model="commentForm.rating" show-score />
            </el-form-item>
            <el-form-item label="评论内容">
              <el-input v-model="commentForm.content" type="textarea" rows="4" placeholder="请输入评论内容" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="handleSubmitComment">提交评论</el-button>
            </el-form-item>
          </el-form>
        </el-card>

        <el-table
          v-loading="commentLoading"
          :data="commentList"
          style="width: 100%"
          border
          class="comment-table"
        >
          <el-table-column prop="user_name" label="用户" width="120" />
          <el-table-column prop="content" label="评论内容" min-width="300" />
          <el-table-column prop="rating" label="评分" width="100">
            <template #default="scope">
              <el-rate v-model="scope.row.rating" disabled show-score />
            </template>
          </el-table-column>
          <el-table-column prop="created_at" label="评论时间" width="180" />
          <el-table-column label="操作" width="100" fixed="right">
            <template #default="scope">
              <el-button type="danger" size="small" @click="handleDeleteComment(scope.row.ID)" plain>
                <el-icon><Delete /></el-icon> 删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-card>

    <!-- 播放视频弹窗 -->
    <el-dialog
      v-model="videoDialogVisible"
      :title="currentLesson?.title"
      width="80%"
    >
      <div class="video-container">
        <video
          :src="currentLesson?.video_url"
          controls
          autoplay
          width="100%"
          height="500px"
        />
        <div class="lesson-content" v-if="currentLesson?.content">
          <h3>课时内容</h3>
          <div v-html="currentLesson.content"></div>
        </div>
      </div>
    </el-dialog>

    <!-- 新建/编辑作业抽屉 -->
    <el-drawer
      v-model="homeworkDrawerVisible"
      :title="homeworkDrawerTitle"
      size="60%"
      destroy-on-close
    >
      <el-form
        :model="homeworkForm"
        label-width="100px"
        class="homework-form"
      >
        <el-form-item label="作业标题" required>
          <el-input v-model="homeworkForm.title" placeholder="请输入作业标题" />
        </el-form-item>
        <el-form-item label="作业描述" required>
          <el-input v-model="homeworkForm.description" type="textarea" rows="4" placeholder="请输入作业描述" />
        </el-form-item>
        <el-form-item label="截止时间" required>
          <el-date-picker
            v-model="homeworkForm.deadline"
            type="datetime"
            placeholder="选择截止时间"
            format="YYYY-MM-DD HH:mm:ss"
            value-format="YYYY-MM-DD HH:mm:ss"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSubmitHomework">提交</el-button>
          <el-button @click="homeworkDrawerVisible = false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive, computed } from 'vue'
import { ArrowRight, VideoPlay, Plus, Edit, Delete, View } from '@element-plus/icons-vue'
import * as courseApi from '../api/course'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useRoute } from 'vue-router'

const route = useRoute()
const courseId = computed(() => route.params.id)

// 课程详情
const courseDetail = ref({
  ID: 0,
  title: '',
  description: '',
  price: 0,
  category: '',
  level: '',
  cover_image: '',
  intro_video: '',
  student_count: 0,
  chapters: []
})
const loading = ref(false)

// 展开的章节
const expandedChapters = ref([])

// 标签页
const activeTab = ref('content')

// 作业相关
const homeworkList = ref([])
const homeworkLoading = ref(false)
const homeworkDrawerVisible = ref(false)
const homeworkDrawerTitle = ref('新建作业')
const homeworkForm = reactive({
  ID: 0,
  course_id: 0,
  title: '',
  description: '',
  deadline: ''
})

// 评论相关
const commentList = ref([])
const commentLoading = ref(false)
const commentForm = reactive({
  course_id: 0,
  content: '',
  rating: 5,
  parent_id: 0
})

// 视频播放
const videoDialogVisible = ref(false)
const currentLesson = ref(null)

// 是否已报名
const isEnrolled = ref(false)

// 获取课程详情
const getCourseDetail = async () => {
  loading.value = true
  try {
    const res = await courseApi.getCourseDetail(courseId.value)
    if (res.code === 0) {
      courseDetail.value = res.data
      // 默认展开第一个章节
      if (res.data.chapters && res.data.chapters.length > 0) {
        expandedChapters.value = [res.data.chapters[0].ID]
      }
      // 检查是否已报名
      checkEnrollment()
    } else {
      ElMessage.error(res.msg)
    }
  } catch (error) {
    ElMessage.error('获取课程详情失败')
    console.error(error)
  } finally {
    loading.value = false
  }
}

// 检查是否已报名
const checkEnrollment = async () => {
  try {
    const res = await courseApi.getEnrolledCourses({ page: 1, pageSize: 100 })
    if (res.code === 0) {
      isEnrolled.value = res.data.list.some(item => item.course_id === Number(courseId.value))
    }
  } catch (error) {
    console.error('检查报名状态失败', error)
  }
}

// 报名课程
const handleEnroll = async () => {
  try {
    const res = await courseApi.enrollCourse({ course_id: Number(courseId.value) })
    if (res.code === 0) {
      ElMessage.success('报名成功')
      isEnrolled.value = true
    } else {
      ElMessage.error(res.msg)
    }
  } catch (error) {
    ElMessage.error('报名失败')
    console.error(error)
  }
}

// 开始学习
const handleLearn = () => {
  // 跳转到学习页面
  window.open(`/courseLearn/${courseId.value}`, '_blank')
}

// 切换章节展开/收起
const toggleChapter = (chapterId) => {
  const index = expandedChapters.value.indexOf(chapterId)
  if (index > -1) {
    expandedChapters.value.splice(index, 1)
  } else {
    expandedChapters.value.push(chapterId)
  }
}

// 播放课时
const playLesson = (lesson) => {
  currentLesson.value = lesson
  videoDialogVisible.value = true
}

// 格式化时长
const formatDuration = (seconds) => {
  if (!seconds) return '00:00'
  const minutes = Math.floor(seconds / 60)
  const secs = seconds % 60
  return `${minutes.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`
}

// 获取分类文本
const getCategoryText = (category) => {
  const map = {
    frontend: '前端开发',
    backend: '后端开发',
    mobile: '移动开发',
    ai: '人工智能',
    cloud: '云计算'
  }
  return map[category] || category
}

// 获取难度文本
const getLevelText = (level) => {
  const map = {
    beginner: '初级',
    intermediate: '中级',
    advanced: '高级'
  }
  return map[level] || level
}

// 获取难度类型
const getLevelType = (level) => {
  const map = {
    beginner: '',
    intermediate: 'warning',
    advanced: 'danger'
  }
  return map[level] || ''
}

// 获取作业列表
const getHomeworkList = async () => {
  homeworkLoading.value = true
  try {
    const res = await courseApi.getHomeworkList(courseId.value, { page: 1, pageSize: 100 })
    if (res.code === 0) {
      homeworkList.value = res.data.list
    } else {
      ElMessage.error(res.msg)
    }
  } catch (error) {
    ElMessage.error('获取作业列表失败')
    console.error(error)
  } finally {
    homeworkLoading.value = false
  }
}

// 新建作业
const handleCreateHomework = () => {
  homeworkDrawerTitle.value = '新建作业'
  Object.assign(homeworkForm, {
    ID: 0,
    course_id: Number(courseId.value),
    title: '',
    description: '',
    deadline: ''
  })
  homeworkDrawerVisible.value = true
}

// 编辑作业
const handleEditHomework = (row) => {
  homeworkDrawerTitle.value = '编辑作业'
  Object.assign(homeworkForm, row)
  homeworkDrawerVisible.value = true
}

// 删除作业
const handleDeleteHomework = (id) => {
  ElMessageBox.confirm('确定要删除该作业吗？', '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      const res = await courseApi.deleteHomework(id)
      if (res.code === 0) {
        ElMessage.success('删除成功')
        getHomeworkList()
      } else {
        ElMessage.error(res.msg)
      }
    } catch (error) {
      ElMessage.error('删除失败')
      console.error(error)
    }
  }).catch(() => {
    // 取消删除
  })
}

// 查看作业提交
const handleViewSubmissions = (homeworkId) => {
  // 跳转到作业提交列表页面
  window.open(`/homeworkSubmissions/${homeworkId}`, '_blank')
}

// 提交作业
const handleSubmitHomework = async () => {
  if (!homeworkForm.title) {
    ElMessage.warning('请输入作业标题')
    return
  }
  
  try {
    let res
    if (homeworkForm.ID === 0) {
      res = await courseApi.createHomework(homeworkForm)
    } else {
      res = await courseApi.updateHomework(homeworkForm)
    }
    
    if (res.code === 0) {
      ElMessage.success(homeworkForm.ID === 0 ? '创建成功' : '更新成功')
      homeworkDrawerVisible.value = false
      getHomeworkList()
    } else {
      ElMessage.error(res.msg)
    }
  } catch (error) {
    ElMessage.error(homeworkForm.ID === 0 ? '创建失败' : '更新失败')
    console.error(error)
  }
}

// 获取评论列表
const getCommentList = async () => {
  commentLoading.value = true
  try {
    const res = await courseApi.getCommentList(courseId.value, { page: 1, pageSize: 100 })
    if (res.code === 0) {
      commentList.value = res.data.list
    } else {
      ElMessage.error(res.msg)
    }
  } catch (error) {
    ElMessage.error('获取评论列表失败')
    console.error(error)
  } finally {
    commentLoading.value = false
  }
}

// 提交评论
const handleSubmitComment = async () => {
  if (!commentForm.content) {
    ElMessage.warning('请输入评论内容')
    return
  }
  
  commentForm.course_id = Number(courseId.value)
  
  try {
    const res = await courseApi.createComment(commentForm)
    if (res.code === 0) {
      ElMessage.success('评论成功')
      commentForm.content = ''
      commentForm.rating = 5
      getCommentList()
    } else {
      ElMessage.error(res.msg)
    }
  } catch (error) {
    ElMessage.error('评论失败')
    console.error(error)
  }
}

// 删除评论
const handleDeleteComment = (id) => {
  ElMessageBox.confirm('确定要删除该评论吗？', '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      const res = await courseApi.deleteComment(id)
      if (res.code === 0) {
        ElMessage.success('删除成功')
        getCommentList()
      } else {
        ElMessage.error(res.msg)
      }
    } catch (error) {
      ElMessage.error('删除失败')
      console.error(error)
    }
  }).catch(() => {
    // 取消删除
  })
}

// 初始化
onMounted(() => {
  getCourseDetail()
  getHomeworkList()
  getCommentList()
})
</script>

<style scoped>
.course-detail {
  padding: 20px;
}

.header-card {
  margin-bottom: 20px;
}

.course-header {
  display: flex;
  gap: 20px;
}

.cover-image {
  width: 300px;
  height: 200px;
  border-radius: 8px;
  overflow: hidden;
  flex-shrink: 0;
}

.cover-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.cover-image .placeholder {
  width: 100%;
  height: 100%;
  background-color: #f0f0f0;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #999;
}

.course-info {
  flex: 1;
}

.course-title {
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 10px;
}

.course-meta {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 15px;
}

.student-count {
  color: #666;
  font-size: 14px;
}

.course-price {
  display: flex;
  align-items: center;
  gap: 20px;
  margin-bottom: 15px;
}

.price {
  font-size: 28px;
  font-weight: bold;
  color: #ff4d4f;
}

.course-description {
  margin-top: 15px;
  line-height: 1.6;
  color: #333;
}

.content-card {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

/* 课程内容样式 */
.course-content {
  padding: 20px 0;
}

.chapter {
  margin-bottom: 15px;
  border: 1px solid #e8e8e8;
  border-radius: 4px;
  overflow: hidden;
}

.chapter-header {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 15px 20px;
  background-color: #f5f5f5;
  cursor: pointer;
  transition: all 0.3s;
}

.chapter-header:hover {
  background-color: #e6f7ff;
}

.chapter-header .el-icon {
  transition: transform 0.3s;
}

.chapter-header .rotated {
  transform: rotate(90deg);
}

.chapter-title {
  font-weight: bold;
  flex: 1;
}

.chapter-lesson-count {
  color: #666;
  font-size: 14px;
}

.chapter-lessons {
  padding: 10px 0;
}

.lesson-item {
  display: flex;
  align-items: center;
  gap: 15px;
  padding: 10px 60px;
  border-bottom: 1px solid #f0f0f0;
  transition: all 0.3s;
}

.lesson-item:hover {
  background-color: #f9f9f9;
}

.lesson-icon {
  color: #1890ff;
}

.lesson-title {
  flex: 1;
}

.lesson-duration {
  color: #666;
  font-size: 14px;
}

/* 作业样式 */
.homework-section {
  padding: 20px 0;
}

.homework-table {
  margin-top: 20px;
}

/* 评论样式 */
.comment-section {
  padding: 20px 0;
}

.comment-form-card {
  margin-bottom: 20px;
}

.comment-form {
  max-width: 800px;
}

.comment-table {
  margin-top: 20px;
}

/* 视频弹窗样式 */
.video-container {
  width: 100%;
}

.lesson-content {
  margin-top: 20px;
  padding: 20px;
  border-top: 1px solid #e8e8e8;
}

.lesson-content h3 {
  margin-bottom: 10px;
}

/* 作业表单样式 */
.homework-form {
  max-width: 600px;
}

@media (max-width: 768px) {
  .course-header {
    flex-direction: column;
  }
  
  .cover-image {
    width: 100%;
    height: 200px;
  }
}
</style>
