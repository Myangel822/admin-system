<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="标题">
        <el-input v-model="form.title" />
      </el-form-item>
      <el-form-item label="内容">
        <el-input v-model="form.content" />
      </el-form-item>

      <el-form-item label="添加视频" class="video-upload">
        <el-upload
          id="video"
          class="avatar-uploader"
          action="http://127.0.0.1:8000/api/v1/demo"
          accept=".mp4,.qlv,.qsv,.ogg,.flv,.avi,.wmv,.rmvb"
          :auto-upload="false"
          :show-file-list="false"
          :on-change="handleChangeVideo"
        >
          <video
            v-if="Video !=''"
            :src="Video"
            width="350"
            height="180"
            controls="controls"
          >您的浏览器不支持视频播放
          </video>
          <i
            v-else-if="Video == ''"
            class="el-icon-plus avatar-uploader-icon"
          />            <!--没选择视频之前显示-->
        </el-upload>
      </el-form-item>
      <el-form-item label="日期">
        <el-col :span="11">
          <el-date-picker v-model="form.date" type="date" placeholder="Pick a date" style="width: 100%;" value-format="yyyy-MM-ddTHH:mm:ss" />
        </el-col>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onSubmit">编辑</el-button>
        <el-button @click="onCancel">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import service from '@/utils/request'
import { getToken } from '@/utils/auth'
export default {
  data() {
    return {
      form: {
        id: '',
        title: '',
        content: '',
        date: ''

      },
      Video: {}
    }
  },
  methods: {
    onSubmit() {
      this.form.id = (this.$route.query.id).toString()
      this.$confirm('此操作将永久更改, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        const formData = new FormData()
        Object.keys(this.form).forEach((ele) => {
          formData.append(ele, this.form[ele])
        })

        console.log('----------------------------')
        console.log(this.form)
        service.put('http://127.0.0.1:8000/api/v1/editDemo?token=' + getToken(), formData, { withCredentials: true }).then(res => {
          console.log(res.data)
        }).then(() => {
          this.$message({
            type: 'success',
            message: '编辑成功!'
          })
          this.fetchData()
        })
      })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消'
          })
        })
    },
    onCancel() {
      this.$message({
        message: 'cancel!',
        type: 'warning'
      })
    },
    handleChangeVideo(file) {
      this.Video = URL.createObjectURL(file.raw)
    	console.log(file)
    	this.form.file = file.raw
    }
  }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

