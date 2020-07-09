<template>
  <div>
    <el-row style="float: right">
      <el-button @click="openDialog('add')" type="primary">新增主机信息</el-button>
      <el-button @click="enterplatformCreateKey()" type="primary">生成平台密钥对</el-button>
      <el-button type="primary" @click="RefreshStatus">刷新状态</el-button>

    </el-row>
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
      <el-table-column
                    label="状态"
                    prop="status"
                    type="scope">
                <template slot-scope="scope">
                    <span class="operate-span-warning" v-if="scope.row.status === 0 " >未测试连接</span>
                    <span class="operate-span-warning" v-else-if="scope.row.status === 1 " >未测试连接</span>
                    <span class="operate-span-warning" v-else-if="scope.row.status === 2 " >连接中</span>
                    <span class="operate-span-primary" v-else-if="scope.row.status === 3 " >连接成功</span>       
                    <span class="operate-span-danger" v-else-if="scope.row.status === 4 " >连接失败</span>
                    <span class="operate-span-primary" v-else-if="scope.row.status === 5 " >公钥推送成功</span>
                    <span class="operate-span-danger" v-else-if="scope.row.status === 5 " >公钥推送失败</span>


                </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <span class="operate-span" @click="editServer(scope.row)" >编辑</span>         
          <span class="operate-span" @click="connectserver(scope.row)" >测试连接</span>
          <span class="operate-span" @click="pushpubkey(scope.row)" >推送公钥</span>
          <span class="operate-span-danger" @click="deleteServer(scope.row)" >删除</span>
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
        <el-form-item label="环境:" prop="resource_env_id">
          <el-select filterable placeholder="请选择" style="width:100%" v-model="form.resource_env_id">
                <el-option
                      :key="item.id"
                      :label="item.name"
                      :value="item.id"
                      v-for="item in env_list" />
          </el-select>
        </el-form-item>
        <el-form-item label="名称" prop="name">
          <el-input autocomplete="off" v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="主机地址" prop="host">
          <el-input autocomplete="off" v-model="form.host"></el-input>
        </el-form-item>
        <el-form-item label="SSH端口" prop="port">
          <el-input autocomplete="off" pattern="[0-9]*" v-model="form.port" type="number" ></el-input>
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
  import { envList } from '@/api/resource/env'
  import { serverList, serverCreate, serverUpdate, serverDelete, platformCreateKey, serverConnect } from '@/api/resource/server'

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
        env_list: [],
        form: {
          id: '',
          name: '',
          host: '',
          port: 22,
          user: '',
          pwd: '',
          resource_env_id: '',
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
          resource_env_id: [
            { required: true, message: '请选择环境', trigger: 'blur' }
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
          port: 22,
          user: '',
          pwd: '',
          resource_env_id: '',
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
        this.$confirm('此操作将永久删除主机信息, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(async () => {
                const res = await serverDelete(row)
                if (res.code == 0) {
                  this.$message({
                      type: 'success',
                      message: res.msg
                    })
                    this.getTableData()
                  }
                }).catch(() => {
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
                this.form.port = parseInt(this.form.port)
                const res = await serverCreate(this.form)
                if (res.code === 0) {
                  this.$message({
                    type: 'success',
                    message: res.msg
                  })
                this.RefreshStatus()
                this.closeDialog()
                }
              }
                break
              case 'edit':
              {
                this.form.port = parseInt(this.form.port)
                const res = await serverUpdate(this.form)
                if (res.code == 0) {
                  this.$message({
                    type: 'success',
                    message: res.msg
                  })
                  this.RefreshStatus()
                  this.closeDialog()
                }
              }
                break
              default:
              {
                this.$message({
                  type: 'error',
                  message: '未知操作'
                })
              }
                break
            }
          }
        })
      },
      async GetEnvList(){
                this.env_list = []
                const ret = await envList({"page": 1, "pageSize": 9999})
                if(ret.code === 0){
                    this.env_list = ret.data.list
                }
            },
      async enterplatformCreateKey() {
        this.$confirm('您将，生成平台密钥对, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(async () => {
                const res = await platformCreateKey()
                if (res.code == 0) {
                  this.$message({
                      type: 'success',
                      message: res.msg
                    })
                  }
                  this.getTableData()
                }).catch(() => {
                  this.$message({
                    type: 'info',
                    message: '已取消平台密钥生成'
                  })
                })
       },
       async pushpubkey(row) {
         console.log(row)

       },
       async connectserver(row) {
        this.$confirm('您将, 测试平台与主机的连通性, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(async () => {
                const res = await serverConnect(row)
                if (res.code == 0) {
                  this.$message({
                      type: 'success',
                      message: res.msg
                    })
                  }
                  this.getTableData()
                }).catch(() => {
                  this.$message({
                    type: 'info',
                    message: '已经，取消连接测试'
                  })
                })
       },
       // 刷新状态
      async RefreshStatus() {
          await this.getTableData()
      },
    },
    created(){
      this.GetEnvList()
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