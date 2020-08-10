<template>
  <div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="创建时间" min-width="150" prop="CreatedAt"></el-table-column>
      <el-table-column
                    label="项目"
                    prop="resourceenv"
                    type="scope">
                <template slot-scope="scope">
                    <span class="operate-span"> {{ scope.row.deployproject.name }}</span>         
                </template>
      </el-table-column>
      <el-table-column label="提测Tag" min-width="150" prop="tag"></el-table-column>
            <el-table-column
                    type="expand"
                    prop="describe"
                    label="描述">
                <template slot-scope="props">
                    <span><pre>{{ props.row.describe }}</pre></span>
                </template>
            </el-table-column>
      <el-table-column
                    label="状态"
                    prop="status"
                    type="scope">
                <template slot-scope="scope">
                    <span class="operate-span-primary" v-if="scope.row.status === 1 " >提成成功</span>
                    <span class="operate-span-danger" v-else-if="scope.row.status === 2 " >提测失败</span>
                </template>
      </el-table-column>            
      <el-table-column label="申请人" min-width="150" prop="applicant"></el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="editProject(scope.row)" size="small" type="primary" >日志</el-button>
        </template>
      </el-table-column>
    </el-table>
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

    <el-dialog :before-close="closeDialog" :title="dialogTitle" :visible.sync="dialogFormVisible">
      <div>
          <pre> {{ result }} </pre>
      </div>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
  // 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索
  import { testingList } from '@/api/deploy/test'
  import infoList from '@/components/mixins/infoList'
  export default {
    name: 'Env',
    mixins: [infoList],
    data() {
      return {
        listApi: testingList,
        dialogFormVisible: false,
        dialogTitle: '日志提测',
        dialogType: '',
        result: '',
      }
    },
    methods: {
      closeDialog() {
        this.dialogFormVisible = false
      },
      async editProject(row) {
        this.dialogTitle = '日志详情'
        this.result = row.result
        this.dialogFormVisible = true

      },
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