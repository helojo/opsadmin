<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="openDialog('add')" type="primary">新增主机信息</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="id" min-width="60" prop="id" ></el-table-column>
      <el-table-column label="名称" min-width="150" prop="name"></el-table-column>
      <el-table-column label="主机地址" min-width="150" prop="host"></el-table-column>
      <el-table-column label="SSH端口" min-width="150" prop="port"></el-table-column>
      <el-table-column label="用户名" min-width="150" prop="user"></el-table-column>
      <el-table-column
                    label="环境"
                    prop="resourceenv"
                    type="scope">
                <template slot-scope="scope">
                    {{ scope.row.resourceenv.name }}
                </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="editServer(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
          <el-button @click="deleteServer(scope.row)" size="small" type="danger" icon="el-icon-delete">删除</el-button>
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
      <el-form :model="form" :rules="rules" label-width="100px" ref="ServerForm">
        <el-form-item label="名称" prop="name">
          <el-input autocomplete="off" v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="主机地址" prop="host">
          <el-input autocomplete="off" v-model="form.host"></el-input>
        </el-form-item>
        <el-form-item label="SSH端口" prop="port">
          <el-input autocomplete="off"  v-model="form.port" type="number" ></el-input>
        </el-form-item>
        <el-form-item label="用户名" prop="user">
          <el-input autocomplete="off"  v-model="form.user"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="pwd">
          <el-input autocomplete="off"  v-model="form.pwd" type="password"></el-input>
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
    import { serverList } from '@/api/resource/server'

  import infoList from '@/components/mixins/infoList'
  export default {
    name: 'Server',
    mixins: [infoList],
    data() {
      return {
        listApi: serverList,
        dialogFormVisible: false,
        dialogTitle: '新增主机信息',
        dialogType: '',
        form: {
          id: '',
          name: '',
          host: '',
          port: '',
          user: '',
          pwd: '',
        },
        type: '',
        rules: {
          name: [
            { required: true, message: '请输入名称', trigger: 'blur' }
          ],
          host: [
            { required: true, message: '请输入主机地址', trigger: 'blur' }
          ],
          port: [
            { required: true, message: '请输入SSh端口', trigger: 'blur' }
          ],
          user: [
            { required: true, message: '请输入用户名', trigger: 'blur' }
          ],   
          pwd: [
            { required: true, message: '请输入密码', trigger: 'blur' }
          ],                               
        }
      }
    },
    methods: {
      initForm() {
        this.$refs.ServerForm.resetFields()
        this.form= {
          id: '',
          name: '',
          host: '',
          port: '',
          user: '',
          pwd: '',
        }
      },
      closeDialog() {
        this.initForm()
        this.dialogFormVisible = false
      },
      openDialog(type) {
        switch (type) {
          case 'add':
            this.dialogTitle = '新增主机信息'
            break
          case 'edit':
            this.dialogTitle = '编辑主机信息'
            break
          default:
            break
        }
        this.dialogType = type
        this.dialogFormVisible = true
      },
      async editServer(row) {
        this.dialogTitle = '编辑主机信息'
        this.dialogType = 'edit'
        for (let key in this.form) {
          this.form[key] = row[key]
        }
        this.dialogFormVisible = true
      },
      async deleteServer(row) {
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
        this.$refs.ServerForm.validate(async valid => {
          if (valid) {
            switch (this.dialogType) {
              case 'add':
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