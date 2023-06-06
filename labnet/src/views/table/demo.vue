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
      <el-table-column label="标题" width="200">
        <template slot-scope="scope">
          {{ scope.row.title }}
        </template>
      </el-table-column>
      <el-table-column label="内容" width="110">
        <template slot-scope="scope">
          {{ scope.row.content }}
        </template>
      </el-table-column>

      <el-table-column prop="hotVideoPath" label="视频" align="center" show-overflow-tooltip>
        <template slot-scope="scope">
          <div class="video">
            <el-dialog
              title="播放视频"
              width="72%"
              append-to-body
              top="20px"
              :visible.sync="dialogVisible"
            >
              <h3>{{ '/static/dist/' + playvideo }}</h3>
              <video
                id="video"
                width="100%"
                autoplay="autoplay"
                :src="('/static/dist/' + playvideo)"
                :poster="'/static/dist/' + playvideo"
                controls="controls"
                preload
              />
            </el-dialog>
            <!-- 页面table显示的video标签 -->
            <video
              :src="('/static/dist/' + scope.row.Video)"
              width="100"
              preload
            />
            <i
              class="el-icon-video-play playIcon"
              @click="handleCheck(scope.row)"
            />
          </div>
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
import { getDemo, deleteDemo, getDemoByName } from '@/api/table'

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
      dialogVisible: false, // 视频播放弹窗
      playvideo: '2023-04-20_135609.mp4', // 存储用户点击的视频播放链接
      playvideoName: null // 存储正在播放视频的名称
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getDemo().then(response => {
        this.pvData = response.data.lists
        this.listLoading = false
        this.totalItems = response.data.lists.length
        console.log('-----------------------------')
        console.log(this.pvData)
      })
    },
    deleteRows(id) {
      this.$confirm('此操作将永久删除, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          deleteDemo(id).then(() => {
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
        getDemoByName(this.inputVal).then(res => {
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
      this.$router.push({ name: 'EditDemo', query: {
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
    handleCheck(row) {
      this.playvideo = row.Video // 存储用户点击的视频播放链接
      console.log(this.playvideo)
      console.log('-------------------------------')
      console.log(row.id)
      this.playvideoName = row.Video // 存储用户点击的视频播放链接
      this.dialogVisible = true
    }

  }
}
</script>
