<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="名称">
        <el-input v-model="form.title" />
      </el-form-item>
      <el-form-item label="目录">
        <el-input v-model="form.kind" />
      </el-form-item>
      <el-form-item label="地址">
        <el-input v-model="form.content" />
      </el-form-item>
      <el-form-item label="Date">
        <el-col :span="11">
          <el-date-picker v-model="form.date" type="date" placeholder="选择日期" style="width: 100%;" />
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
import { editAchievement } from '@/api/table'
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
      this.id = this.$route.query.id
      this.$confirm('此操作将永久更改, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        editAchievement(this.id, this.form).then(() => {
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
    }
  }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

