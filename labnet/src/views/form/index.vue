<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="标题">
        <el-input v-model="form.title" />
      </el-form-item>
      <el-form-item label="期刊">
        <el-input v-model="form.journal" />
      </el-form-item>
      <el-form-item label="作者">
        <el-input v-model="form.author" />
      </el-form-item>
      <el-form-item label="其他作者">
        <el-input v-model="form.authors" />
      </el-form-item>

      <el-form-item label="链接">
        <el-input v-model="form.link" />
      </el-form-item>
      <el-form-item label="分类">
        <el-input v-model="form.papercode" />
      </el-form-item>
      <el-form-item label="年份">
        <el-input v-model="form.theyear" />
      </el-form-item>
      <el-form-item label="状态" prop="countryType">
        <el-select v-model="form.state" placeholder="请选择论文状态">
          <el-option
            v-for="item in stateList"
            :key="item.id"
            :label="item.name"
            :value="item.name"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="隐藏" prop="countryType">
        <el-select v-model="form.hide" placeholder="请选择论文是否隐藏">
          <el-option
            v-for="item in isHideList"
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

      <el-form-item label="摘要">
        <el-input v-model="form.abstract" type="textarea" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">提交</el-button>
        <el-button @click="onCancel">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { addArticle } from '@/api/table'
export default {
  data() {
    return {
      form: {
        title: '',
        author: '',
        authors: '',
        date: '',
        abstract: '',
        theyear: '',
        papercode: '',
        link: '',
        journal: '',
        state: '',
        hide: ''

      },
      stateList: [
        { id: 1, name: '草稿' },
        { id: 2, name: '已发表' }
      ],
      isHideList: [
        { id: 1, name: '隐藏' },
        { id: 2, name: '显示' }
      ]
    }
  },
  methods: {
    onSubmit() {
      this.$message('submit!')
      addArticle(this.form).then().catch(error => {
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

