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
      <el-table-column label="名称" width="200">
        <template slot-scope="scope">
          {{ scope.row.title }}
        </template>
      </el-table-column>
      <el-table-column label="分类" width="200">
        <template slot-scope="scope">
          {{ scope.row.kind }}
        </template>
      </el-table-column>
      <el-table-column label="地址" width="200">
        <template slot-scope="scope">
          {{ scope.row.content }}
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" min-width="100">
        　　　　<template slot-scope="scope">
          　　　　　　<el-button type="primary" @click="editRows(scope.row.id)">编辑</el-button>
          　　　　　　<el-button type="primary" @click="deleteRows(scope.row.id)">删除</el-button>
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
import { getAchievement, deleteAchievement, getAchievementByName } from '@/api/table'

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
      getAchievement().then(response => {
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
          deleteAchievement(id).then(() => {
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
        getAchievementByName(this.inputVal).then(res => {
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
      this.$router.push({ name: 'EditAchievement', query: {
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
    }
  }
}
</script>
