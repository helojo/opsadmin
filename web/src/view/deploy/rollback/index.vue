<template>
  <div>
    <div class="button-box clearflex">
      <el-form :inline="true" class="demo-form-inline">
        <el-form-item style="float: right" >
          <el-button @click="Refresh" type="primary">刷新列表 </el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="创建时间" min-width="150" prop="CreatedAt"></el-table-column>
      <el-table-column
              label="环境"
              prop="env"
              type="scope">
        <template slot-scope="scope">
          <span class="operate-span"> {{ scope.row.deployproject.environment.name }}</span>
        </template>
      </el-table-column>
      <el-table-column
              label="主机"
              prop="server"
              type="scope">
        <template slot-scope="scope">
          <div :key="item.id" v-for="item in scope.row.deployproject.Server">
            <span style="float: left">{{ item.name }} , </span>
          </div>
        </template>
      </el-table-column>
      <el-table-column
              label="项目"
              prop="name"
              type="scope">
        <template slot-scope="scope">
          <span class="operate-span"> {{ scope.row.deployproject.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="回退前版本" min-width="150" prop="release_version"></el-table-column>
      <el-table-column label="回退后版本" min-width="150" prop="after_version"></el-table-column>
      <el-table-column
              label="状态"
              prop="status"
              type="scope">
        <template slot-scope="scope">
          <span class="operate-span-danger" v-if="scope.row.status === 0 " >回滚中</span>
          <span class="operate-span-primary" v-else-if="scope.row.status === 1 " >成功</span>
          <span class="operate-span-danger" v-else-if="scope.row.status === 2 " >失败</span>
        </template>
      </el-table-column>
      <el-table-column label="操作人" min-width="150" prop="aperator"></el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="RollbackLogs(scope.row)" size="small" type="primary" >日志</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-row justify="center" style="padding-top:20px;" type="flex">
      <span class="demonstration" />
      <el-pagination
              :current-page="page"
              :page-size="pageSize"
              :page-sizes="[10, 30, 50, 100]"
              :style="{float:'right',padding:'20px'}"
              :total="total"
              @current-change="handleCurrentChange"
              @size-change="handleSizeChange"
              layout="total, sizes, prev, pager, next, jumper"
      ></el-pagination>
    </el-row>
    <el-dialog :before-close="closeresultDialog" :title="dialogTitle" :visible.sync="dialogFormVisible">
      <div style="background: #000000">
        <pre style="color: #e8e8e8"> {{ result }} </pre>
      </div>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeresultDialog">取 消</el-button>
      </div>
    </el-dialog>

  </div>

</template>


<script>
  // 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索
  import { rollbackList } from '@/api/deploy/rollback'
  import infoList from '@/components/mixins/infoList'
  export default {
    name: 'Env',
    mixins: [infoList],
    data() {
      return {
        listApi: rollbackList,
        dialogTitle: '日志详情',
        result: '',
        dialogFormVisible: false
      }
    },
    methods: {
      async RollbackLogs(row) {
        this.dialogTitle = '日志详情'
        this.result = row.result
        this.dialogFormVisible = true
      },
      closeresultDialog(){
        this.dialogFormVisible = false
      },
      async Refresh() {
        this.getTableData()
      }
    },
    created(){
      this.getTableData()
    }
  }
</script>
<style scoped lang="scss">
  .button-box {
    padding: 10px 20px;
    .el-button {
      float: right;
    }
  }
  .el-tag--mini {
    margin-left: 5px;
  }
  .warning {
    color: #dc143c;
  }
</style>