<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="标题">
        <el-input v-model="form.title" />
      </el-form-item>
      <el-form-item label="内容">
        <el-input v-model="form.content" />
      </el-form-item>
      <el-form-item label="项目状态" prop="countryType">
        <el-select v-model="form.state" placeholder="请选择项目状态">
          <el-option
            v-for="item in stateList"
            :key="item.id"
            :label="item.name"
            :value="item.name"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="年份">
        <el-input v-model="form.theyear" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">提交</el-button>
        <el-button @click="onCancel">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { addProject } from '@/api/table'
export default {
  data() {
    return {
      form: {
        title: '',
        content: '',
        State: '',
        theyear: ''
      },
      stateList: [
        { id: 1, name: '在研' },
        { id: 2, name: '结题' }
      ]
    }
  },
  methods: {
    onSubmit() {
      this.$message('submit!')
      addProject(this.form).then().catch(error => {
        reject(error)
      })
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

