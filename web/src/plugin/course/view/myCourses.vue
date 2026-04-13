<template>
  <div class="my-courses">
    <el-card shadow="never" class="search-card">
      <el-form :inline="true" :model="searchForm" class="search-form">
        <el-form-item label="关键词">
          <el-input v-model="searchForm.keyword" placeholder="请输入课程标题" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="进行中" value="active" />
            <el-option label="已完成" value="completed" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">搜索</el-button>
          <el-button @click="resetForm">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card shadow="never" class="table-card">
      <el-table
        v-loading="loading"
        :data="courseList"
        style="width: 100%"
        border
      >
        <el-table-column prop="ID" label="课程ID" width="80" />
        <el-table-column prop="title" label="课程标题" min-width="200">
          <template #default="scope">
            <el-link type="primary" @click="handleView(scope.row.course_id)">{{ scope.row.title }}</el-link>
          </template>
        </el-table-column>
        <el-table-column prop="category" label="分类" width="120">
          <template #default="scope">
            <el-tag>{{ getCategoryText(scope.row.category) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="level" label="难度" width="100">
          <template #default="scope">
            <el-tag :type="getLevelType(scope.row.level)">{{ getLevelText(scope.row.level) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="progress" label="学习进度" width="120">
          <template #default="scope">
            <el-progress :percentage="scope.row.progress" :color="getProgressColor(scope.row.progress)" />
          </template>
        </el-table-column>
        <el-table-column prop="enrolled_at" label="报名时间" width="180" />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="scope">
            <el-button type="primary" size="small" @click="handleLearn(scope.row.course_id)" plain>
              开始学习
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="pagination-container">
        <el-pagination
          v-model:current-page="searchForm.page"
          v-model:page-size="searchForm.pageSize"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive, computed } from 'vue'
import * as courseApi from '../api/course'
import { ElMessage } from 'element-plus'

// 搜索表单
const searchForm = reactive({
  page: 1,
  pageSize: 10,
  keyword: '',
  status: ''
})

// 课程列表
const courseList = ref([])
const total = ref(0)
const loading = ref(false)

// 搜索课程
const handleSearch = () => {
  getEnrolledCourses()
}

// 重置表单
const resetForm = () => {
  searchForm.keyword = ''
  searchForm.status = ''
  searchForm.page = 1
  getEnrolledCourses()
}

// 获取已报名课程
const getEnrolledCourses = async () => {
  loading.value = true
  try {
    const res = await courseApi.getEnrolledCourses(searchForm)
    if (res.code === 0) {
      courseList.value = res.data.list
      total.value = res.data.total
    } else {
      ElMessage.error(res.msg)
    }
  } catch (error) {
    ElMessage.error('获取课程列表失败')
    console.error(error)
  } finally {
    loading.value = false
  }
}

// 分页处理
const handleSizeChange = (size) => {
  searchForm.pageSize = size
  getEnrolledCourses()
}

const handleCurrentChange = (current) => {
  searchForm.page = current
  getEnrolledCourses()
}

// 查看课程详情
const handleView = (courseId) => {
  window.open(`/courseDetail/${courseId}`, '_blank')
}

// 开始学习
const handleLearn = (courseId) => {
  window.open(`/courseLearn/${courseId}`, '_blank')
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

// 获取进度颜色
const getProgressColor = (progress) => {
  if (progress === 100) {
    return '#52c41a'
  } else if (progress >= 50) {
    return '#faad14'
  } else {
    return '#1890ff'
  }
}

// 初始化
onMounted(() => {
  getEnrolledCourses()
})
</script>

<style scoped>
.my-courses {
  padding: 20px;
}

.search-card {
  margin-bottom: 20px;
}

.search-form {
  margin-bottom: 0;
}

.table-card {
  margin-bottom: 20px;
}

.pagination-container {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>
