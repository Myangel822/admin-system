<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="标题">
        <el-input v-model="form.title" />
      </el-form-item>
      <el-form-item label="内容">
        <el-input v-model="form.content" />
      </el-form-item>

      <el-form-item label="种类" prop="countryType">
        <el-select v-model="form.kind" placeholder="请选择活动类别">
          <el-option
            v-for="item in kindList"
            :key="item.id"
            :label="item.name"
            :value="item.name"
          />
        </el-select>
      </el-form-item>

      <el-form-item label="图片上传">
        <el-upload
          action="http://127.0.0.1:7070/api/v1/news"
          list-type="picture-card"
          :on-preview="handlePictureCardPreview"
          :on-remove="handleRemove"
          :before-upload="beforeProductUpload"
        >
          <i class="el-icon-plus" />
        </el-upload>
        <el-dialog :visible.sync="dialogVisible">
          <img width="100%" :src="dialogImageUrl" alt="">
        </el-dialog>
      </el-form-item>

      <el-form-item label="Date">
        <el-col :span="11">
          <el-date-picker v-model="form.date" type="date" placeholder="选择日期" style="width: 100%;" />
        </el-col>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onSubmit">提交</el-button>
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
        title: '',
        content: '',
        kind: '',
        date: '',
        imageUrl: []

      },
      dialogImageUrl: '',
      dialogVisible: false,
      kindList: [
        { id: 1, name: '会议' },
        { id: 2, name: '竞赛' },
        { id: 3, name: '开题答辩' },
        { id: 4, name: '毕业答辩' },
        { id: 5, name: '活动' }
      ]
    }
  },
  methods: {
    onSubmit() {
      console.log(this.form.imageurl)
      var Form = {
        title: this.form.title,
        content: this.form.content,
        kind: this.form.kind,
        date: this.form.date,
        imageUrl: this.form.imageUrl
      }
      service.post('http://127.0.0.1:8000/api/v1/activity?token=' + getToken(), Form, { withCredentials: true }).then(res => {
        console.log(res.data)
        this.$message('发布成功')
        this.Form = ''
        console.log('上传成功')
      })
      // addNews(this.form.title, this.form.content, this.form.imageurl).then(res => {
      //   console.log('--------------------------')
      //   console.log(res.data)
      //   this.$message('发布成功')
      //   this.form = ''
      //   console.log('上传成功')
      // })
    },
    beforeProductUpload(file) {
      var _this = this
      return new Promise(function(resolve, reject) {
        var reader = new FileReader()
        reader.readAsDataURL(file)// 这里是最关键的一步，转换成base64
        console.log(file)
        reader.onload = function(event) {
          _this.form.imageUrl.push(event.target.result) // 定义参数获取图片路径
          console.log(_this.imageurl)
        }
      })
    },
    handleRemove(file, fileList) {
      // 移除图片
      console.log(file, fileList)
    },
    handlePictureCardPreview(file) {
      // 图片预览
      this.dialogImageUrl = file.url
      this.dialogVisible = true
    },
    onCancel() {
      this.$message({
        message: 'cancel!',
        type: 'warning'
      })
    }
  }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

