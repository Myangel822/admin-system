<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="标题">
        <el-input v-model="form.title" />
      </el-form-item>
      <el-form-item label="内容">
        <el-input v-model="form.content" />
      </el-form-item>
      <el-form-item label="类别" prop="countryType">
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
import { addStandards } from '@/api/table'
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
      this.$message('submit!')
      addStandards(this.form).then().catch(error => {
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

