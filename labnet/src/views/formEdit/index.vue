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
      <el-form-item label="种类">
        <el-input v-model="form.papercode" />
      </el-form-item>
      <el-form-item label="年份">
        <el-input v-model="form.theyear" />
      </el-form-item>
      <el-form-item label="状态" prop="countryType">
        <el-select v-model="form.State">
          <el-option
            v-for="item in stateList"
            :key="item.id"
            :label="item.name"
            :value="item.name"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="隐藏" prop="countryType">
        <el-select v-model="form.IsHide">
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
          <el-date-picker v-model="form.date" type="date" placeholder="Pick a date" style="width: 100%;" />
        </el-col>
      </el-form-item>

      <el-form-item label="摘要">
        <el-input v-model="form.abstract" type="textarea" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">编辑</el-button>
        <el-button @click="onCancel">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { editArticle } from '@/api/table'
export default {
  data() {
    return {
      form: {
        id: 0,
        title: '',
        author: '',
        authors: '',
        date: '',
        abstract: '',
        theyear: '',
        papercode: '',
        link: '',
        journal: '',
        State: '',
        IsHide: ''

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
  created() {
    this.form.id = this.$route.query.row.id
    this.form.title = this.$route.query.row.title
    this.form.author = this.$route.query.row.author
    this.form.authors = this.$route.query.row.authors
    this.form.date = this.$route.query.row.date
    this.form.abstract = this.$route.query.row.abstract
    this.form.theyear = this.$route.query.row.theyear
    this.form.papercode = this.$route.query.row.papercode
    this.form.link = this.$route.query.row.link
    this.form.journal = this.$route.query.row.journal
    this.form.State = this.$route.query.row.state
    this.form.IsHide = this.$route.query.row.hide
    console.log('---------------------------')
    console.log('this.$route.query.row', this.$route.query.row)
    console.log('form', this.form)
  },
  methods: {
    onSubmit() {
      this.$confirm('此操作将永久更改, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        editArticle(this.form).then(() => {
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

