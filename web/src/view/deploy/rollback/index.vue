<template>
  <div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="id" min-width="60" prop="id" ></el-table-column>
      <el-table-column label="名称" min-width="150" prop="name"></el-table-column>
      <el-table-column
              label="标签">
        <template slot-scope="scope">
          <span v-if="scope.row.env_label === 0 " class="operate-span-primary" >无标签</span>
          <span v-else-if="scope.row.env_label === 1 " class="operate-span-primary" >开发</span>
          <span v-else-if="scope.row.env_label === 2 " class="operate-span-primary" >测试</span>
          <span v-else-if="scope.row.env_label === 3 " class="operate-span-primary" >灰度</span>
          <span v-else-if="scope.row.env_label === 4 " class="operate-span-primary" >生产</span>
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <span @click="editEnv(scope.row)" class="operate-span">编辑</span>
          <span @click="deleteEnv(scope.row)" class="operate-span-danger">删除</span>
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
      }
    },
    methods: {
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