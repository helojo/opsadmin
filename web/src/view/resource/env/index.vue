<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="openDialog('addEnv')" type="primary">新增环境</el-button>
    </div>
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

    <el-dialog :before-close="closeDialog" :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form :model="form" :rules="rules" label-width="80px" ref="EnvForm">
        <el-form-item label="名称" prop="name">
          <el-input autocomplete="off" v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="标签" prop="env_label">
            <el-radio-group v-model="form.env_label">
                  <el-radio :label="1">开发</el-radio>
                  <el-radio :label="2">测试</el-radio>
                  <el-radio :label="3">灰度</el-radio>
                  <el-radio :label="4">生产</el-radio>
              </el-radio-group>
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
  import { envList, envCreate, envUpdate, envDelete } from '@/api/resource/env'
  import infoList from '@/components/mixins/infoList'
  export default {
    name: 'Env',
    mixins: [infoList],
    data() {
      return {
        listApi: envList,
        dialogFormVisible: false,
        dialogTitle: '新增环境',
        dialogType: '',
        form: {
          id: '',
          name: '',
          env_label: 1,
        },
        type: '',
        rules: {
          name: [
            { required: true, message: '请输入环境名称', trigger: 'blur' }
          ],
          env_label: [
            { required: true, message: '请输选择标签', trigger: 'blur' }
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
          env_label: 1
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