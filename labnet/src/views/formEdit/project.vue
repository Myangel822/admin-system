<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="名称">
        <el-input v-model="form.title" />
      </el-form-item>
      <el-form-item label="链接">
        <el-input v-model="form.content" />
      </el-form-item>
      <el-form-item label="状态" prop="countryType">
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
        <el-button type="primary" @click="onSubmit">编辑</el-button>
        <el-button @click="onCancel">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { editProject } from '@/api/table'
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
      this.id = this.$route.query.id
      this.$confirm('此操作将永久更改, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        editProject(this.id, this.form).then(() => {
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

