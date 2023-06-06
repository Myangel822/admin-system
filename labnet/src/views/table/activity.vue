<template>
  <div class="app-container">
    <div class="filter-container" style="margin-bottom: 20px">
      <el-input
        v-model="inputVal"
        clearable

        placeholder="请输入标题"
        style="width:300px"
      />
      <el-select

        placeholder="活动类型"
        clearable
        style="width: 150px; margin-right: 30px"
        class="filter-item"
        @change="changeHandler"
      >
        <el-option
          v-for="item in kind"
          :key="item.id"
          :label="item.name"
          :value="item.name"
        />

      </el-select>
      <el-button type="primary" @click="searchData(true)">查询</el-button>
    </div>
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
      <el-table-column label="标题" width="110">
        <template slot-scope="scope">
          {{ scope.row.title }}
        </template>
      </el-table-column>
      <el-table-column label="内容" width="110">
        <template slot-scope="scope">
          {{ scope.row.content }}
        </template>
      </el-table-column>
      <el-table-column label="种类" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.kind }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="img" label="图片" width="200">
        <template v-slot:default="scope">
          <el-image :src="('/static/dist/' + scope.row.imageUrl)" />
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="日期" width="200">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.date }}</span>
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
import { getActivity, deleteActivity, getActivityByName, getActivityByKind } from '@/api/table'

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
      pvData: [],
      kind: [
        { id: 1, name: '会议' },
        { id: 2, name: '竞赛' },
        { id: 3, name: '开题答辩' },
        { id: 4, name: '毕业答辩' },
        { id: 5, name: '活动' }
      ]

    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getActivity().then(response => {
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
          deleteActivity(id).then(() => {
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
        getActivityByName(this.inputVal).then(res => {
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
    changeHandler(value) {
      this.currentPage = 1
      getActivityByKind(value).then(res => {
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
    },

    editRows(id) {
      this.$router.push({ name: 'EditActivity', query: {
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
