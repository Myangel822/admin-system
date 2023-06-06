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
          {{ scope.$index+1 }}
        </template>
      </el-table-column>
      <el-table-column label="名称" width="110">
        <template slot-scope="scope">
          {{ scope.row.title }}
        </template>
      </el-table-column>
      <el-table-column label="链接" width="110">
        <template slot-scope="scope">
          {{ scope.row.content }}
        </template>
      </el-table-column>
      <el-table-column label="种类" width="110">
        <template slot-scope="scope">
          {{ scope.row.state }}
        </template>
      </el-table-column>
      <el-table-column label="年份" width="110">
        <template slot-scope="scope">
          {{ scope.row.date }}
        </template>
      </el-table-column>

      <el-table-column label="操作" align="center" min-width="100">
        <template slot-scope="scope">
          <el-button type="primary" @click="editRows(scope.row.id)">编辑</el-button>
          <el-button type="primary" @click="deleteRows(scope.row.id)">删除</el-button>
          <el-button type="primary" @click="handleEdit(scope.$index,scope.row)">状态</el-button>

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
                <el-select v-model="form.state" placeholder="请选择技术标准状态">
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
import { getProject, deleteProject, getProjectByName, editProjectState } from '@/api/table'

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
      stateList: [
        { id: 1, name: '在研' },
        { id: 2, name: '结题' }
      ]
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getProject().then(response => {
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
          deleteProject(id).then(() => {
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
        getProjectByName(this.inputVal).then(res => {
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
    editRows(id) {
      this.$router.push({ name: 'EditProject', query: {
        id: id
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
      editProjectState(params).then(
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
