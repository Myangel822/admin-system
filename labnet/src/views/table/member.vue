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
      <el-table-column label="用户名" width="110">
        <template slot-scope="scope">
          {{ scope.row.username }}
        </template>
      </el-table-column><el-table-column label="密码" width="110">
        <template slot-scope="scope">
          {{ scope.row.password }}
        </template>
      </el-table-column>
      <el-table-column label="姓名" width="110">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="电话" width="110">
        <template slot-scope="scope">
          {{ scope.row.phone }}
        </template>
      </el-table-column>
      <el-table-column label="邮箱" width="110">
        <template slot-scope="scope">
          {{ scope.row.mail }}
        </template>
      </el-table-column>
      <el-table-column label="简介" width="400">
        <template slot-scope="scope">
          {{ scope.row.introduction }}
        </template>
      </el-table-column>
      <el-table-column label="身份" width="110">
        <template slot-scope="scope">
          {{ scope.row.identity }}
        </template>
      </el-table-column>
      <el-table-column label="研究方向" width="110">
        <template slot-scope="scope">
          {{ scope.row.research }}
        </template>
      </el-table-column>
      <el-table-column label="成就" width="400">
        <template slot-scope="scope">
          {{ scope.row.achievement }}
        </template>
      </el-table-column>

      </el-table-column>
      <el-table-column label="操作" align="center" min-width="400">
        <template slot-scope="scope">
          <el-button type="primary" @click="editRows(scope.row.id)">编辑</el-button>
          <el-button type="primary" @click="deleteRows(scope.row.id)">删除</el-button>
          <el-button type="primary" @click="reset(scope.row.id)">重置</el-button>
          <el-button v-if="scope.row.identity=='毕业学生'" type="primary" @click="graduationTo(scope.row)">毕业去向</el-button>
          <el-dialog
            title="毕业去向"
            :visible.sync="dialogFormVisible"
            width="500px"
            top="200px"
          >
            <p>毕业去向：{{ graduation }}</p>
            <p>毕业时间：{{ graduationtime }}</p>
            <div
              slot="footer"
              class="dialog-footer"
            >
              <el-button @click="dialogFormVisible = false">取 消</el-button>
            </div>
          </el-dialog>
        </template></el-table-column>

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
import { getMember, deleteMember, resetPassword, getMemberByName } from '@/api/table'

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

      graduation: '..',
      graduationtime: '..',

      dialogFormVisible: false,
      list: null,
      listLoading: true,
      inputVal: '',
      currentPage: 1, // 初始页
      pagesize: 2, //    每页的数据
      total: 0,
      pvData: []
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getMember().then(response => {
        this.pvData = response.data.lists
        console.log('---------------------------------------')
        console.log(this.pvData)
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
          deleteMember(id).then(() => {
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
        getMemberByName(this.inputVal).then(res => {
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
    reset(id) {
      this.$confirm('此操作将永久重置, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          resetPassword(id).then(() => {
            this.$message({
              type: 'success',
              message: '重置成功!'
            })
            this.fetchData()
          })
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消重置'
          })
        })
    },
    editRows(id) {
      this.$router.push({ name: 'EditMember', query: {
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
    graduationTo(row) {
      var _this = this
      this.dialogFormVisible = true
      _this.graduation = row.Graduation
      _this.graduationtime = row.Graduationtime

      console.log('毕业去向')
      console.log(row)
    }
  }
}
</script>
