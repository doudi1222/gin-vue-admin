<template>
  <div class="course-list">
    <el-card shadow="never" class="search-card">
      <el-form :inline="true" :model="searchForm" class="search-form">
        <el-form-item label="关键词">
          <el-input v-model="searchForm.keyword" placeholder="请输入课程标题或描述" clearable />
        </el-form-item>
        <el-form-item label="分类">
          <el-select v-model="searchForm.category" placeholder="请选择分类" clearable>
            <el-option label="前端开发" value="frontend" />
            <el-option label="后端开发" value="backend" />
            <el-option label="移动开发" value="mobile" />
            <el-option label="人工智能" value="ai" />
            <el-option label="云计算" value="cloud" />
          </el-select>
        </el-form-item>
        <el-form-item label="难度">
          <el-select v-model="searchForm.level" placeholder="请选择难度" clearable>
            <el-option label="初级" value="beginner" />
            <el-option label="中级" value="intermediate" />
            <el-option label="高级" value="advanced" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">搜索</el-button>
          <el-button @click="resetForm">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card shadow="never" class="action-card">
      <el-button type="primary" @click="handleCreate" plain>
        <el-icon><Plus /></el-icon> 新建课程
      </el-button>
    </el-card>

    <el-card shadow="never" class="table-card">
      <el-table
        v-loading="loading"
        :data="courseList"
        style="width: 100%"
        border
      >
        <el-table-column prop="ID" label="课程ID" width="80" />
        <el-table-column prop="title" label="课程标题" min-width="200" />
        <el-table-column prop="category" label="分类" width="120">
          <template #default="scope">
            <el-tag v-if="scope.row.category === 'frontend'">前端开发</el-tag>
            <el-tag v-else-if="scope.row.category === 'backend'">后端开发</el-tag>
            <el-tag v-else-if="scope.row.category === 'mobile'">移动开发</el-tag>
            <el-tag v-else-if="scope.row.category === 'ai'">人工智能</el-tag>
            <el-tag v-else-if="scope.row.category === 'cloud'">云计算</el-tag>
            <el-tag v-else>{{ scope.row.category }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="level" label="难度" width="100">
          <template #default="scope">
            <el-tag v-if="scope.row.level === 'beginner'" size="small">初级</el-tag>
            <el-tag v-else-if="scope.row.level === 'intermediate'" size="small" type="warning">中级</el-tag>
            <el-tag v-else-if="scope.row.level === 'advanced'" size="small" type="danger">高级</el-tag>
            <el-tag v-else size="small">{{ scope.row.level }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="price" label="价格" width="100" />
        <el-table-column prop="student_count" label="学生数" width="100" />
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="scope">
            <el-button type="primary" size="small" @click="handleEdit(scope.row)" plain>
              <el-icon><Edit /></el-icon> 编辑
            </el-button>
            <el-button type="danger" size="small" @click="handleDelete(scope.row.ID)" plain>
              <el-icon><Delete /></el-icon> 删除
            </el-button>
            <el-button type="info" size="small" @click="handleView(scope.row.ID)" plain>
              <el-icon><View /></el-icon> 详情
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

    <!-- 新建/编辑课程抽屉 -->
    <el-drawer
      v-model="drawerVisible"
      :title="drawerTitle"
      size="80%"
      destroy-on-close
    >
      <el-form
        :model="courseForm"
        ref="courseFormRef"
        label-width="100px"
        class="course-form"
      >
        <el-form-item label="课程标题" required>
          <el-input v-model="courseForm.title" placeholder="请输入课程标题" />
        </el-form-item>
        <el-form-item label="课程描述" required>
          <el-input v-model="courseForm.description" type="textarea" rows="4" placeholder="请输入课程描述" />
        </el-form-item>
        <el-form-item label="课程价格" required>
          <el-input-number v-model="courseForm.price" :min="0" :step="0.01" placeholder="请输入课程价格" />
        </el-form-item>
        <el-form-item label="课程分类" required>
          <el-select v-model="courseForm.category" placeholder="请选择课程分类">
            <el-option label="前端开发" value="frontend" />
            <el-option label="后端开发" value="backend" />
            <el-option label="移动开发" value="mobile" />
            <el-option label="人工智能" value="ai" />
            <el-option label="云计算" value="cloud" />
          </el-select>
        </el-form-item>
        <el-form-item label="难度级别" required>
          <el-select v-model="courseForm.level" placeholder="请选择难度级别">
            <el-option label="初级" value="beginner" />
            <el-option label="中级" value="intermediate" />
            <el-option label="高级" value="advanced" />
          </el-select>
        </el-form-item>
        <el-form-item label="封面图片">
          <el-upload
            class="upload-demo"
            :action="uploadUrl"
            :on-success="handleUploadSuccess"
            :before-upload="beforeUpload"
            v-model:file-list="coverFileList"
            list-type="picture"
          >
            <el-button type="primary">
              <el-icon><Upload /></el-icon> 上传封面
            </el-button>
          </el-upload>
        </el-form-item>
        <el-form-item label="介绍视频">
          <el-upload
            class="upload-demo"
            :action="uploadUrl"
            :on-success="handleVideoUploadSuccess"
            :before-upload="beforeVideoUpload"
            v-model:file-list="videoFileList"
            list-type="text"
          >
            <el-button type="primary">
              <el-icon><Upload /></el-icon> 上传视频
            </el-button>
          </el-upload>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSubmit">提交</el-button>
          <el-button @click="drawerVisible = false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive, computed } from 'vue'
import { Plus, Edit, Delete, View, Upload } from '@element-plus/icons-vue'
import * as courseApi from '../api/course'
import { ElMessage } from 'element-plus'
import { useBtnAuth } from '@/utils/btnAuth'

// 按钮权限
const btnAuth = useBtnAuth()

// 搜索表单
const searchForm = reactive({
  page: 1,
  pageSize: 10,
  keyword: '',
  category: '',
  level: ''
})

// 课程列表
const courseList = ref([])
const total = ref(0)
const loading = ref(false)

// 抽屉相关
const drawerVisible = ref(false)
const drawerTitle = ref('新建课程')
const courseForm = reactive({
  ID: 0,
  title: '',
  description: '',
  price: 0,
  category: '',
  level: '',
  cover_image: '',
  intro_video: ''
})
const courseFormRef = ref(null)

// 上传相关
const uploadUrl = '/api/v1/fileUploadAndDownload/upload'
const coverFileList = ref([])
const videoFileList = ref([])

// 搜索课程
const handleSearch = () => {
  getCourseList()
}

// 重置表单
const resetForm = () => {
  searchForm.keyword = ''
  searchForm.category = ''
  searchForm.level = ''
  searchForm.page = 1
  getCourseList()
}

// 获取课程列表
const getCourseList = async () => {
  loading.value = true
  try {
    const res = await courseApi.getCourseList(searchForm)
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
  getCourseList()
}

const handleCurrentChange = (current) => {
  searchForm.page = current
  getCourseList()
}

// 新建课程
const handleCreate = () => {
  drawerTitle.value = '新建课程'
  Object.assign(courseForm, {
    ID: 0,
    title: '',
    description: '',
    price: 0,
    category: '',
    level: '',
    cover_image: '',
    intro_video: ''
  })
  coverFileList.value = []
  videoFileList.value = []
  drawerVisible.value = true
}

// 编辑课程
const handleEdit = (row) => {
  drawerTitle.value = '编辑课程'
  Object.assign(courseForm, row)
  // 处理文件列表
  coverFileList.value = row.cover_image ? [{ url: row.cover_image }] : []
  videoFileList.value = row.intro_video ? [{ url: row.intro_video }] : []
  drawerVisible.value = true
}

// 查看课程详情
const handleView = (id) => {
  // 跳转到课程详情页面
  window.open(`/courseDetail/${id}`, '_blank')
}

// 删除课程
const handleDelete = (id) => {
  ElMessageBox.confirm('确定要删除该课程吗？', '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      const res = await courseApi.deleteCourse(id)
      if (res.code === 0) {
        ElMessage.success('删除成功')
        getCourseList()
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

// 提交表单
const handleSubmit = async () => {
  if (!courseForm.title) {
    ElMessage.warning('请输入课程标题')
    return
  }
  
  try {
    let res
    if (courseForm.ID === 0) {
      res = await courseApi.createCourse(courseForm)
    } else {
      res = await courseApi.updateCourse(courseForm)
    }
    
    if (res.code === 0) {
      ElMessage.success(courseForm.ID === 0 ? '创建成功' : '更新成功')
      drawerVisible.value = false
      getCourseList()
    } else {
      ElMessage.error(res.msg)
    }
  } catch (error) {
    ElMessage.error(courseForm.ID === 0 ? '创建失败' : '更新失败')
    console.error(error)
  }
}

// 上传封面图片前处理
const beforeUpload = (file) => {
  const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png'
  const isLt2M = file.size / 1024 / 1024 < 2
  
  if (!isJpgOrPng) {
    ElMessage.error('只能上传 JPG/PNG 图片')
    return false
  }
  if (!isLt2M) {
    ElMessage.error('图片大小不能超过 2MB')
    return false
  }
  return true
}

// 上传视频前处理
const beforeVideoUpload = (file) => {
  const isVideo = file.type.startsWith('video/')
  const isLt100M = file.size / 1024 / 1024 < 100
  
  if (!isVideo) {
    ElMessage.error('只能上传视频文件')
    return false
  }
  if (!isLt100M) {
    ElMessage.error('视频大小不能超过 100MB')
    return false
  }
  return true
}

// 上传封面成功处理
const handleUploadSuccess = (response, uploadFile) => {
  if (response.code === 0) {
    courseForm.cover_image = response.data
  } else {
    ElMessage.error('上传失败')
  }
}

// 上传视频成功处理
const handleVideoUploadSuccess = (response, uploadFile) => {
  if (response.code === 0) {
    courseForm.intro_video = response.data
  } else {
    ElMessage.error('上传失败')
  }
}

// 初始化
onMounted(() => {
  getCourseList()
})
</script>

<style scoped>
.course-list {
  padding: 20px;
}

.search-card {
  margin-bottom: 20px;
}

.search-form {
  margin-bottom: 0;
}

.action-card {
  margin-bottom: 20px;
  padding: 10px 20px;
}

.table-card {
  margin-bottom: 20px;
}

.pagination-container {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}

.course-form {
  max-width: 800px;
}
</style>
