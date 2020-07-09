<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="openDialog('addEnv')" type="primary">新增环境</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="id" min-width="60" prop="id" ></el-table-column>
      <el-table-column label="项目" min-width="150" prop="name"></el-table-column>
      <el-table-column label="当前版本" min-width="150" prop="release_version"></el-table-column>
      <el-table-column label="项目地址" min-width="150" prop="git_url"></el-table-column>
      <el-table-column label="项目目录" min-width="150" prop="directory"></el-table-column>
      <el-table-column label="忽略文件" min-width="150" prop="ignore_files"></el-table-column>
      <el-table-column
                    label="主机"
                    prop="resourceserver"
                    type="scope">
                <template slot-scope="scope">
                    {{ scope.row.resourceserver.name }}
                </template>
      </el-table-column>
            <el-table-column
                    label="环境"
                    prop="resourceenv"
                    type="scope">
                <template slot-scope="scope">
                    {{ scope.row.resourceserver.resourceenv.name }}
                </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="editEnv(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
          <el-button @click="deleteEnv(scope.row)" size="small" type="danger" icon="el-icon-delete">删除</el-button>
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
      <el-form :inline="true" :model="form" :rules="rules" label-width="80px" ref="EnvForm">
        <el-form-item label="名称" prop="name">
          <el-input autocomplete="off" v-model="form.name"></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
  // 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索
  import { envCreate, envUpdate, envDelete } from '@/api/resource/env'
  import { projectList } from '@/api/deploy/project'
  import infoList from '@/components/mixins/infoList'
  export default {
    name: 'Env',
    mixins: [infoList],
    data() {
      return {
        listApi: projectList,
        dialogFormVisible: false,
        dialogTitle: '新增环境',
        dialogType: '',
        form: {
          id: '',
          name: '',
          git_url: '',
          directory: '',
          ignore_files: '',
        },
        type: '',
        rules: {
          name: [
            { required: true, message: '请输入环境名称', trigger: 'blur' }
          ]
        }
      }
    },
    methods: {
      initForm() {
        this.$refs.EnvForm.resetFields()
        this.form= {
          id: '',
          name: '',
          git_url: '',
          directory: '',
          ignore_files: '',
        }
      },
      closeDialog() {
        this.initForm()
        this.dialogFormVisible = false
      },
      openDialog(type) {
        switch (type) {
          case 'addEnv':
            this.dialogTitle = '新增环境'
            break
          case 'edit':
            this.dialogTitle = '编辑环境'
            break
          default:
            break
        }
        this.dialogType = type
        this.dialogFormVisible = true
      },
      async editEnv(row) {
        this.dialogTitle = '编辑环境'
        this.dialogType = 'edit'
        for (let key in this.form) {
          this.form[key] = row[key]
        }
        this.dialogFormVisible = true
      },
      async deleteEnv(row) {
        this.$confirm('此操作将永久删除环境, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })
                .then(async () => {
                  const res = await envDelete(row)
                  if (res.code == 0) {
                    this.$message({
                      type: 'success',
                      message: '删除成功!'
                    })
                    this.getTableData()
                  }
                })
                .catch(() => {
                  this.$message({
                    type: 'info',
                    message: '已取消删除'
                  })
                })
      },
      async enterDialog() {
        this.$refs.EnvForm.validate(async valid => {
          if (valid) {
            switch (this.dialogType) {
              case 'addEnv':
              {
                const res = await envCreate(this.form)
                if (res.code === 0) {
                  this.$message({
                    type: 'success',
                    message: '添加成功',
                    showClose: true
                  })
                }
                this.getTableData()
                this.closeDialog()
              }
                break
              case 'edit':
              {
                const res = await envUpdate(this.form)
                if (res.code == 0) {
                  this.$message({
                    type: 'success',
                    message: '编辑成功',
                    showClose: true
                  })
                }
                this.getTableData()
                this.closeDialog()
              }
                break
              default:
              {
                this.$message({
                  type: 'error',
                  message: '未知操作',
                  showClose: true
                })
              }
                break
            }
          }
        })
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