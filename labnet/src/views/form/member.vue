<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="用户名">
        <el-input v-model="form.username" />
      </el-form-item>
      <el-form-item label="密码">
        <el-input v-model="form.password" />
      </el-form-item>
      <el-form-item label="姓名">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item label="头像上传">
        <el-upload
          action="http://127.0.0.1:8000/api/v1/news"
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
      <el-form-item label="电话">
        <el-input v-model="form.phone" />
      </el-form-item>
      <el-form-item label="邮箱">
        <el-input v-model="form.mail" />
      </el-form-item>

      <el-form-item label="身份" prop="countryType">
        <el-select v-model="form.identity" placeholder="请选择成员身份">
          <el-option
            v-for="item in identityList"
            :key="item.id"
            :label="item.name"
            :value="item.name"
          />
        </el-select>
      </el-form-item>

      <el-form-item v-if="form.identity=='毕业学生'" id="showClass" label="毕业去向">
        <el-input v-model="form.graduation" />
      </el-form-item>

      <el-form-item v-if="form.identity=='毕业学生'" label="毕业日期">
        <el-col :span="11">
          <el-date-picker v-model="form.graduationtime" type="date" placeholder="选择日期" style="width: 100%;" />
        </el-col>
      </el-form-item>

      <el-form-item label="简介">
        <el-input v-model="form.introduction" />
      </el-form-item>
      <el-form-item label="研究方向">
        <el-input v-model="form.research" />
      </el-form-item>
      <el-form-item label="成就">
        <el-input v-model="form.achievement" />
      </el-form-item>
      <el-form-item label="状态" prop="countryType">
        <el-select v-model="form.state" placeholder="请选择成员状态">
          <el-option
            v-for="item in stateList"
            :key="item.id"
            :label="item.name"
            :value="item.name"
          />
        </el-select>
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
        username: '',
        password: '',
        name: '',
        phone: '',
        mail: '',
        graduation: '',
        graduationtime: '',
        identity: '',
        introduction: '',
        research: '',
        achievement: '',
        State: '',
        image: []
      },
      dialogImageUrl: '',
      dialogVisible: false,
      identityList: [
        { id: 1, name: '在校学生' },
        { id: 2, name: '毕业学生' },
        { id: 3, name: '教师' },
        { id: 4, name: '系统管理员' }
      ],
      stateList: [
        { id: 1, name: '可用' },
        { id: 2, name: '不可用' }
      ]
    }
  },
  methods: {
    onSubmit() {
      console.log(this.form.image)
      var Form = {
        username: this.form.username,
        password: this.form.password,
        name: this.form.name,
        phone: this.form.phone,
        mail: this.form.mail,
        graduation: this.form.graduation,
        graduationtime: this.form.graduationtime,
        identity: this.form.identity,
        introduction: this.form.introduction,
        research: this.form.research,
        achievement: this.form.achievement,
        State: this.form.State,
        image: this.form.image
      }
      service.post('http://127.0.0.1:8000/api/v1/member?token=' + getToken(), Form, { withCredentials: true }).then(res => {
        console.log(res.data)
        this.$message('添加成功')
        this.Form = ''
        console.log('上传成功')
      })
    },
    beforeProductUpload(file) {
      var _this = this
      return new Promise(function(resolve, reject) {
        var reader = new FileReader()
        reader.readAsDataURL(file)// 这里是最关键的一步，转换成base64
        console.log(file)
        reader.onload = function(event) {
          _this.form.image.push(event.target.result) // 定义参数获取图片路径
          console.log(_this.image)
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
