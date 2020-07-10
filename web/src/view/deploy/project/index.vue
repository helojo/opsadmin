<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="openDialog('add')" type="primary">新增项目</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="id" min-width="60" prop="id" ></el-table-column>
      <el-table-column label="项目名" min-width="150" prop="name"></el-table-column>
      <el-table-column label="当前版本" min-width="150" prop="release_version"></el-table-column>
      <el-table-column label="Git地址" min-width="150" prop="git_url"></el-table-column>
      <el-table-column label="目录" min-width="150" prop="directory"></el-table-column>
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
                    {{ scope.row.resourceenv.name }}
                </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="editProject(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
          <el-button @click="deleteProject(scope.row)" size="small" type="danger" icon="el-icon-delete">删除</el-button>
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
      <el-form :model="form" :rules="rules" label-width="80px" ref="projectForm">
        <el-form-item label="环境" prop="resource_env_id">
                    <el-select  @change="EnvChange" filterable placeholder="请选择" style="width:100%" v-model="form.resource_env_id">
                        <el-option
                                :key="item.id"
                                :label="item.name"
                                :value="item.id"
                                v-for="item in env_List" />
                    </el-select>
        </el-form-item>
        <el-form-item label="主机" prop="resource_server_id">
                    <el-select  filterable placeholder="请选择" style="width:100%" v-model="form.resource_server_id">
                        <el-option
                                :key="item.id"
                                :label="item.name"
                                :value="item.id"
                                v-for="item in server_List" />
                    </el-select>
        </el-form-item>
        <el-form-item label="项目名" prop="name">
          <el-input autocomplete="off" v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="Git地址" prop="git_url">
          <el-input autocomplete="off" v-model="form.git_url"></el-input>
        </el-form-item>    
        <el-form-item label="目录" prop="directory">
          <el-input autocomplete="off" v-model="form.directory"></el-input>
        </el-form-item>
        <el-form-item label="忽略文件" prop="ignore_files">
          <el-input autocomplete="off" v-model="form.ignore_files"></el-input>
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
  import { serverList } from '@/api/resource/server'
  import { projectList, projectCreate, projectUpdate, projectDelete } from '@/api/deploy/project'
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
        env_List: [],
        server_List: [],
        form: {
          id: '',
          name: '',
          git_url: '',
          directory: '',
          ignore_files: '.git',
          resource_env_id: '',
          resource_server_id: '',
        },
        type: '',
        rules: {
          name: [
            { required: true, message: '请输入环境名称', trigger: 'blur' }
          ],
          git_url: [
            { required: true, message: '请输入Git地址', trigger: 'blur' }
          ],
          directory: [
            { required: true, message: '请输入项目目录', trigger: 'blur' }
          ],     
          ignore_files: [
            { required: true, message: '请输入忽略文件，多个空格区分 ', trigger: 'blur' }
          ],   
          resource_env_id: [
            { required: true, message: '请输入选择环境', trigger: 'blur' }
          ],         
          resource_server_id: [
            { required: true, message: '请输入选择主机', trigger: 'blur' }
          ],             
        }
      }
    },
    methods: {
      initForm() {
        this.$refs.projectForm.resetFields()
        this.form= {
          id: '',
          name: '',
          git_url: '',
          directory: '',
          ignore_files: '.git',
          resource_env_id: '',
          resource_server_id: '',
        }
      },
      closeDialog() {
        this.initForm()
        this.dialogFormVisible = false
      },
      openDialog(type) {
        switch (type) {
          case 'add':
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
      async editProject(row) {
        this.dialogTitle = '编辑环境'
        this.dialogType = 'edit'
        for (let key in this.form) {
          this.form[key] = row[key]
        }
        this.dialogFormVisible = true
        this.GetServerList()

      },
      async deleteProject(row) {
        this.$confirm('此操作将永久删除环境, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })
                .then(async () => {
                  const res = await projectDelete(row)
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
        this.$refs.projectForm.validate(async valid => {
          if (valid) {
            switch (this.dialogType) {
              case 'add':
              {
                const res = await projectCreate(this.form)
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
                const res = await projectUpdate(this.form)
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
      },
      async GetEnvList(){
         this.env_List = []
         const ret = await envList({"page": 1, "pageSize": 9999})
         if(ret.code === 0){
            this.env_List = ret.data.list
           }
     },
     async GetServerList(){
         this.serverList = []
         const ret = await serverList({"page": 1, "pageSize": 9999})
         if(ret.code === 0){
            this.serverList = ret.data.list
        }
     },
     async EnvChange(row){
        this.server_List = []
        const ret = await serverList({"page": 1, "pageSize": 9999, "resource_env_id": row})
        if(ret.code === 0){
            this.form.resource_server_id = ''
            this.server_List = ret.data.list
          }
     }
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