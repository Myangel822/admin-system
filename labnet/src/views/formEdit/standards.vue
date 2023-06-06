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
        <el-select v-model="form.kind" placeholder="请选择技术标准类别">
          <el-option
            v-for="item in kindList"
            :key="item.id"
            :label="item.name"
            :value="item.name"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="状态" prop="countryType">
        <el-select v-model="form.state" placeholder="请选择技术标准状态">
          <el-option
            v-for="item in stateList"
            :key="item.id"
            :label="item.name"
            :value="item.name"
          />
        </el-select>
      </el-form-item>

      <el-form-item label="日期">
        <el-col :span="11">
          <el-date-picker v-model="form.date" type="date" placeholder="Pick a date" style="width: 100%;" />
        </el-col>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onSubmit">编辑</el-button>
        <el-button @click="onCancel">提交</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { editStandards } from '@/api/table'
export default {
  data() {
    return {
      form: {
        title: '',
        content: '',
        kind: '',
        state: '',
        date: ''

      },
      kindList: [
        { id: 1, name: '国内' },
        { id: 2, name: '国外' }
      ],
      stateList: [
        { id: 1, name: '申请中' },
        { id: 2, name: '已申请' }
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
        editStandards(this.id, this.form).then(() => {
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

