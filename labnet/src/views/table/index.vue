<template>
  <div class="app-container">
    <el-input
      v-model="inputVal"
      clearable

      placeholder="请输入标题"
      style="width:300px"
    />
    <el-button type="primary" @click="searchData(true)">查询</el-button>
    <el-table
      v-loading="listLoading"
      :data="pvData.slice((currentPage - 1) * pagesize, currentPage * pagesize)"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="编号" width="95">
        <template slot-scope="scope">
          {{ scope.$index +1 }}
        </template>
      </el-table-column>
      <el-table-column label="标题" width="100">
        <template slot-scope="scope">
          {{ scope.row.title }}
        </template>
      </el-table-column>
      <el-table-column label="期刊" width="100">
        <template slot-scope="scope">
          {{ scope.row.journal }}
        </template>
      </el-table-column>
      <el-table-column label="作者" width="100" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.author }}</span>
        </template>
      </el-table-column>
      <el-table-column label="其他作者" width="300" align="center">
        <template slot-scope="scope">
          {{ scope.row.authors }}
        </template>
      </el-table-column>

      <el-table-column align="center" prop="created_at" label="日期" width="110">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.date }}</span>
        </template>
      </el-table-column>
      <el-table-column label="状态" width="100" align="center">
        <template slot-scope="scope">
          {{ scope.row.state }}
        </template>
      </el-table-column>
      <el-table-column label="隐藏" width="100" align="center">
        <template slot-scope="scope">
          {{ scope.row.hide }}
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" min-width="400">
        <template slot-scope="scope">
          <el-button type="primary" @click="editRows(scope.row)">编辑</el-button>
          <el-button type="primary" @click="deleteRows(scope.row.id)">删除</el-button>
          <el-button type="primary" @click="handleEdit(scope.$index,scope.row)">状态</el-button>
          <el-button type="primary" @click="handleEdit1(scope.$index,scope.row)">隐藏</el-button>

          <el-dialog

            title="编辑"
            :visible.sync="dialogFormVisible"
            width="500px"
            top="200px"
          >
            <el-form :model="form">
              <el-form-item
                label="状态"
                :label-width="formLabelWidth"
              >
                <el-select v-model="form.state" placeholder="请选择论文状态">
                  <el-option
                    v-for="item in stateList"
                    :key="item.id"
                    :label="item.name"
                    :value="item.name"
                  />
                </el-select>
              </el-form-item>
            </el-form>
            <div
              slot="footer"
              class="dialog-footer"
            >

              <el-button @click="dialogFormVisible = false">取 消</el-button>
              <el-button
                type="primary"
                @click="handleSet(scope.$index)"
              >确 定</el-button>
            </div>
          </el-dialog>
          <el-dialog
            title="编辑"
            :visible.sync="dialogFormVisible1"
            width="500px"
            top="200px"
          >
            <el-form :model="form">
              <el-form-item
                label="隐藏"
                :label-width="formLabelWidth"
              >
                <el-select v-model="Form.hide" placeholder="请选择论文状态">
                  <el-option
                    v-for="item in isHideList"
                    :key="item.id"
                    :label="item.name"
                    :value="item.name"
                  />
                </el-select>
              </el-form-item>
            </el-form>
            <div
              slot="footer"
              class="dialog-footer"
            >

              <el-button @click="dialogFormVisible1 = false">取 消</el-button>
              <el-button
                type="primary"
                @click="handleSet1(scope.$index)"
              >确 定</el-button>
            </div>
          </el-dialog>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      align="center"
      style="margin: 12px 0px"
      background
      :current-page="currentPage"
      :page-sizes="[1, 2, 4, 8]"
      :page-size="pagesize"
      layout="total, sizes, prev, pager, next, jumper"
      :total="pvData.length"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
    />

  </div>
</template>

<script>
import { getArticle, deleteArticle, getArticleByName, editArticleState, editArticleHide } from '@/api/table'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'gray',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      dialogFormVisible1: false,
      dialogFormVisible: false,
      list: null,
      listLoading: true,
      inputVal: '',
      formLabelWidth: '50px',
      currentPage: 1, // 初始页
      pagesize: 2, //    每页的数据
      total: 0,
      pvData: [],
      form: {
        id: 0,
        state: ''
      },
      Form: {
        id: 0,
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
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getArticle().then(response => {
        this.pvData = response.data.lists
        this.listLoading = false
        this.totalItems = response.data.lists.length
      })
    },
    deleteRows(id) {
      this.$confirm('此操作将永久删除, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          deleteArticle(id).then(() => {
            this.$message({
              type: 'success',
              message: '删除成功!'
            })
            this.fetchData()
          })
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
    },
    searchData(bool) {
      this.currentPage = 1
      if (bool) {
        getArticleByName(this.inputVal).then(res => {
          this.pvData = res.data.lists
          console.log('-----------------------------------------')
          console.log(this.Data)
          this.total = res.data.lists.length
          // this.reloadPart()

          this.$message({
            type: 'success',
            message: '查询成功!'
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '查询失败'
          })
        })
      } else {
        this.refreshData() // 更新表
      }
    },
    editRows(row) {
      this.$router.push({ name: 'EditArticle', query: {
        row: row
      }})
    },
    handleSizeChange: function(size) {
      this.pagesize = size
      console.log(this.pagesize) // 每页下拉显示数据
    },
    handleCurrentChange: function(currentPage) {
      this.currentPage = currentPage
      console.log(this.currentPage) // 点击第几页
    },
    handleEdit(index, row) {
      console.log(row.id)
      this.dialogFormVisible = true // 开启弹出层
      this.form.state = row.state
      this.form.id = row.id
    },
    handleSet(index) {
      console.log(this.form.id)
      var params = {
        state: this.form.state,
        id: this.form.id
      }
      editArticleState(params).then(
        res => {
          if (res.code === 200) {
            this.$message('修改成功！')
            this.fetchData()
          } else {
            this.$message(res.data)
          }
        }
      )
    },
    handleEdit1(index, row) {
      console.log(row.id)
      this.dialogFormVisible1 = true // 开启弹出层
      this.Form.hide = row.hide
      this.Form.id = row.id
    },
    handleSet1(index) {
      console.log(this.Form.id)
      var params = {
        hide: this.Form.hide,
        id: this.Form.id
      }
      editArticleHide(params).then(
        res => {
          if (res.code === 200) {
            this.$message('修改成功！')
            this.fetchData()
          } else {
            this.$message(res.data)
          }
        }
      )
    }
  }
}
</script>
