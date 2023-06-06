<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="名称">
        <el-input v-model="form.title" />
      </el-form-item>
      <el-form-item label="类别">
        <el-input v-model="form.kind" />
      </el-form-item>
      <el-form-item label="内容">
        <el-input v-model="form.content" />
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
import { addAchievement } from '@/api/table'
export default {
  data() {
    return {
      form: {
        title: '',
        kind: '',
        content: '',
        date: ''
      }

    }
  },
  methods: {
    onSubmit() {
      this.$message('submit!')
      addAchievement(this.form).then().catch(error => {
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

