<template>
  <div class="course-learn">
    <el-card shadow="never" class="course-info-card">
      <div class="course-info">
        <h2 class="course-title">{{ courseDetail.title }}</h2>
        <div class="course-progress">
          <span>学习进度：{{ progressPercentage }}%</span>
          <el-progress :percentage="progressPercentage" :color="progressColor" />
        </div>
      </div>
    </el-card>

    <div class="learn-content">
      <!-- 左侧章节列表 -->
      <div class="chapter-sidebar">
        <el-card shadow="never" class="sidebar-card">
          <template #header>
            <div class="sidebar-header">
              <span>课程章节</span>
            </div>
          </template>
          <div class="chapter-list">
            <div v-for="chapter in courseDetail.chapters" :key="chapter.ID" class="chapter-item">
              <div class="chapter-title" @click="toggleChapter(chapter.ID)">
                <el-icon :class="{ 'rotated': expandedChapters.includes(chapter.ID) }">
                  <ArrowRight />
                </el-icon>
                <span>{{ chapter.title }}</span>
              </div>
              <div class="lesson-list" v-show="expandedChapters.includes(chapter.ID)">
                <div
                  v-for="lesson in chapter.lessons"
                  :key="lesson.ID"
                  class="lesson-item"
                  :class="{ 'active': currentLesson && currentLesson.ID === lesson.ID, 'completed': lesson.is_completed }"
                  @click="selectLesson(lesson)"
                >
                  <el-icon class="lesson-icon">
                    <VideoPlay v-if="!lesson.is_completed" />
                    <Check v-else />
                  </el-icon>
                  <span class="lesson-title">{{ lesson.title }}</span>
                  <span class="lesson-duration">{{ formatDuration(lesson.duration) }}</span>
                </div>
              </div>
            </div>
          </div>
        </el-card>
      </div>

      <!-- 右侧视频内容 -->
      <div class="video-content">
        <el-card shadow="never" class="video-card">
          <template #header>
            <div class="video-header">
              <h3>{{ currentLesson?.title }}</h3>
            </div>
          </template>
          <div class="video-container" v-if="currentLesson">
            <video
              ref="videoRef"
              :src="currentLesson.video_url"
              controls
              @timeupdate="updateProgress"
              @ended="handleVideoEnd"
              width="100%"
              height="500px"
            />
            <div class="lesson-content" v-if="currentLesson.content">
              <h4>课时内容</h4>
              <div v-html="currentLesson.content"></div>
            </div>
          </div>
          <div class="empty-state" v-else>
            <el-empty description="请选择一个课时开始学习" />
          </div>
        </el-card>

        <!-- 作业列表 -->
        <el-card shadow="never" class="homework-card" v-if="homeworkList.length > 0">
          <template #header>
            <div class="homework-header">
              <span>课程作业</span>
            </div>
          </template>
          <div class="homework-list">
            <div v-for="homework in homeworkList" :key="homework.ID" class="homework-item">
              <div class="homework-info">
                <h4>{{ homework.title }}</h4>
                <p class="homework-description">{{ homework.description }}</p>
                <div class="homework-meta">
                  <span class="deadline">截止时间：{{ homework.deadline }}</span>
                  <el-tag :type="getHomeworkStatusType(homework.status)">{{ homework.status }}</el-tag>
                </div>
              </div>
              <div class="homework-actions">
                <el-button type="primary" size="small" @click="handleSubmitHomework(homework)">
                  提交作业
                </el-button>
              </div>
            </div>
          </div>
        </el-card>
      </div>
    </div>

    <!-- 提交作业抽屉 -->
    <el-drawer
      v-model="homeworkDrawerVisible"
      :title="`提交作业：${currentHomework?.title}`"
      size="60%"
      destroy-on-close
    >
      <el-form
        :model="homeworkSubmissionForm"
        label-width="100px"
        class="homework-submission-form"
      >
        <el-form-item label="作业内容" required>
          <el-input v-model="homeworkSubmissionForm.content" type="textarea" rows="6" placeholder="请输入作业内容" />
        </el-form-item>
        <el-form-item label="附件上传">
          <el-upload
            class="upload-demo"
            :action="uploadUrl"
            :on-success="handleFileUploadSuccess"
            v-model:file-list="fileList"
            list-type="text"
          >
            <el-button type="primary">
              <el-icon><Upload /></el-icon> 上传附件
            </el-button>
          </el-upload>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitHomework">提交</el-button>
          <el-button @click="homeworkDrawerVisible = false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive, computed, watch } from 'vue'
import { ArrowRight, VideoPlay, Check, Upload } from '@element-plus/icons-vue'
import * as courseApi from '../api/course'
import { ElMessage } from 'element-plus'
import { useRoute } from 'vue-router'

const route = useRoute()
const courseId = computed(() => route.params.id)

// 课程详情
const courseDetail = ref({
  ID: 0,
  title: '',
  chapters: []
})
const loading = ref(false)

// 展开的章节
const expandedChapters = ref([])

// 当前选中的课时
const currentLesson = ref(null)
const videoRef = ref(null)

// 学习进度
const progressPercentage = ref(0)
const progressColor = ref('#1890ff')

// 作业相关
const homeworkList = ref([])
const homeworkDrawerVisible = ref(false)
const currentHomework = ref(null)
const homeworkSubmissionForm = reactive({
  homework_id: 0,
  content: '',
  file_url: ''
})

// 上传相关
const uploadUrl = '/api/v1/fileUploadAndDownload/upload'
const fileList = ref([])

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
        // 默认选择第一个课时
        if (res.data.chapters[0].lessons && res.data.chapters[0].lessons.length > 0) {
          selectLesson(res.data.chapters[0].lessons[0])
        }
      }
      // 计算学习进度
      calculateProgress()
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

// 获取作业列表
const getHomeworkList = async () => {
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
  }
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

// 选择课时
const selectLesson = async (lesson) => {
  currentLesson.value = lesson
  // 获取学习进度
  try {
    const res = await courseApi.getProgress(courseId.value, lesson.ID)
    if (res.code === 0 && res.data) {
      // 设置视频播放位置
      if (videoRef.value && res.data.watched_duration) {
        videoRef.value.currentTime = res.data.watched_duration
      }
    }
  } catch (error) {
    console.error('获取学习进度失败', error)
  }
}

// 更新学习进度
const updateProgress = async () => {
  if (!videoRef.value || !currentLesson.value) return
  
  const currentTime = videoRef.value.currentTime
  const totalTime = videoRef.value.duration
  
  // 每30秒更新一次进度
  if (Math.floor(currentTime) % 30 === 0) {
    try {
      await courseApi.updateProgress({
        user_id: 0, // 实际应该从登录状态获取
        course_id: Number(courseId.value),
        lesson_id: currentLesson.value.ID,
        watched_duration: Math.floor(currentTime),
        is_completed: currentTime >= totalTime * 0.95 // 观看95%以上算完成
      })
    } catch (error) {
      console.error('更新学习进度失败', error)
    }
  }
}

// 视频结束处理
const handleVideoEnd = async () => {
  if (!currentLesson.value) return
  
  try {
    await courseApi.updateProgress({
      user_id: 0, // 实际应该从登录状态获取
      course_id: Number(courseId.value),
      lesson_id: currentLesson.value.ID,
      watched_duration: currentLesson.value.duration,
      is_completed: true
    })
    // 更新课时状态
    currentLesson.value.is_completed = true
    // 重新计算进度
    calculateProgress()
    ElMessage.success('课时学习完成')
  } catch (error) {
    console.error('更新学习进度失败', error)
  }
}

// 计算学习进度
const calculateProgress = () => {
  let totalLessons = 0
  let completedLessons = 0
  
  courseDetail.value.chapters.forEach(chapter => {
    chapter.lessons.forEach(lesson => {
      totalLessons++
      if (lesson.is_completed) {
        completedLessons++
      }
    })
  })
  
  if (totalLessons > 0) {
    progressPercentage.value = Math.round((completedLessons / totalLessons) * 100)
    // 根据进度设置颜色
    if (progressPercentage.value === 100) {
      progressColor.value = '#52c41a'
    } else if (progressPercentage.value >= 50) {
      progressColor.value = '#faad14'
    } else {
      progressColor.value = '#1890ff'
    }
  }
}

// 格式化时长
const formatDuration = (seconds) => {
  if (!seconds) return '00:00'
  const minutes = Math.floor(seconds / 60)
  const secs = seconds % 60
  return `${minutes.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`
}

// 获取作业状态类型
const getHomeworkStatusType = (status) => {
  const map = {
    'pending': 'info',
    'submitted': 'primary',
    'graded': 'success',
    'late': 'warning'
  }
  return map[status] || ''
}

// 提交作业
const handleSubmitHomework = (homework) => {
  currentHomework.value = homework
  Object.assign(homeworkSubmissionForm, {
    homework_id: homework.ID,
    content: '',
    file_url: ''
  })
  fileList.value = []
  homeworkDrawerVisible.value = true
}

// 上传文件成功处理
const handleFileUploadSuccess = (response, uploadFile) => {
  if (response.code === 0) {
    homeworkSubmissionForm.file_url = response.data
  } else {
    ElMessage.error('上传失败')
  }
}

// 提交作业
const submitHomework = async () => {
  if (!homeworkSubmissionForm.content) {
    ElMessage.warning('请输入作业内容')
    return
  }
  
  try {
    const res = await courseApi.submitHomework(homeworkSubmissionForm)
    if (res.code === 0) {
      ElMessage.success('提交成功')
      homeworkDrawerVisible.value = false
      getHomeworkList()
    } else {
      ElMessage.error(res.msg)
    }
  } catch (error) {
    ElMessage.error('提交失败')
    console.error(error)
  }
}

// 初始化
onMounted(() => {
  getCourseDetail()
  getHomeworkList()
})
</script>

<style scoped>
.course-learn {
  padding: 20px;
}

.course-info-card {
  margin-bottom: 20px;
}

.course-info {
  padding: 20px;
}

.course-title {
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 20px;
}

.course-progress {
  margin-top: 10px;
}

.course-progress span {
  display: block;
  margin-bottom: 10px;
  font-size: 14px;
  color: #666;
}

.learn-content {
  display: flex;
  gap: 20px;
}

.chapter-sidebar {
  width: 300px;
  flex-shrink: 0;
}

.sidebar-card {
  height: calc(100vh - 200px);
  overflow-y: auto;
}

.sidebar-header {
  font-weight: bold;
  font-size: 16px;
}

.chapter-list {
  margin-top: 10px;
}

.chapter-item {
  margin-bottom: 10px;
}

.chapter-title {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px 15px;
  cursor: pointer;
  border-radius: 4px;
  transition: all 0.3s;
}

.chapter-title:hover {
  background-color: #f0f0f0;
}

.chapter-title .el-icon {
  transition: transform 0.3s;
}

.chapter-title .rotated {
  transform: rotate(90deg);
}

.lesson-list {
  margin-left: 25px;
  margin-top: 5px;
}

.lesson-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 8px 15px;
  cursor: pointer;
  border-radius: 4px;
  transition: all 0.3s;
  margin-bottom: 5px;
}

.lesson-item:hover {
  background-color: #f0f0f0;
}

.lesson-item.active {
  background-color: #e6f7ff;
  border-left: 3px solid #1890ff;
}

.lesson-item.completed .lesson-icon {
  color: #52c41a;
}

.lesson-title {
  flex: 1;
  font-size: 14px;
}

.lesson-duration {
  font-size: 12px;
  color: #999;
}

.video-content {
  flex: 1;
}

.video-card {
  margin-bottom: 20px;
}

.video-header h3 {
  margin: 0;
  font-size: 18px;
  font-weight: bold;
}

.video-container {
  margin-top: 20px;
}

.lesson-content {
  margin-top: 20px;
  padding: 20px;
  border-top: 1px solid #e8e8e8;
}

.lesson-content h4 {
  margin-bottom: 10px;
  font-size: 16px;
}

.empty-state {
  padding: 100px 0;
  text-align: center;
}

.homework-card {
  margin-top: 20px;
}

.homework-header {
  font-weight: bold;
  font-size: 16px;
}

.homework-list {
  margin-top: 10px;
}

.homework-item {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  padding: 15px;
  border: 1px solid #e8e8e8;
  border-radius: 4px;
  margin-bottom: 10px;
  transition: all 0.3s;
}

.homework-item:hover {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.homework-info {
  flex: 1;
}

.homework-info h4 {
  margin: 0 0 10px 0;
  font-size: 16px;
}

.homework-description {
  margin: 0 0 10px 0;
  font-size: 14px;
  color: #666;
  line-height: 1.4;
}

.homework-meta {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 14px;
  color: #999;
}

.homework-actions {
  margin-left: 20px;
}

.homework-submission-form {
  max-width: 600px;
}

@media (max-width: 768px) {
  .learn-content {
    flex-direction: column;
  }
  
  .chapter-sidebar {
    width: 100%;
  }
  
  .sidebar-card {
    height: auto;
    max-height: 400px;
  }
}
</style>
