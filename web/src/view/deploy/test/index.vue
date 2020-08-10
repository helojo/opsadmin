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
                    <span class="operate-span-warning" v-if="scope.row.status === 0 " >提测中</span>
                    <span class="operate-span-primary" v-else-if="scope.row.status === 1 " >提成成功</span>
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
      <el-form :model="form" :rules="rules" label-width="50px" ref="projectForm">
        <el-form-item label="环境" prop="resource_env_id">
                    <el-select  @change="EnvChange" filterable placeholder="请选择" style="width:100%" v-model="form.resource_env_id">
                        <el-option
                                :key="item.id"
                                :label="item.name"
                                :value="item.id"
                                v-for="item in env_List" />
                    </el-select>
        </el-form-item>
        <el-form-item  label="项目" prop="deploy_project_id">
                    <el-select @change="ProjectChange" filterable placeholder="请选择" style="width:100%" v-model="form.deploy_project_id">
                        <el-option
                                :key="item.id"
                                :label="item.name"
                                :value="item.id"
                                v-for="item in project_List" />
                    </el-select>
        </el-form-item>
        <el-form-item  label="Tag" prop="tag">
                    <el-select filterable placeholder="请选择" style="width:100%" v-model="form.tag">
                        <el-option
                                :key="item.id"
                                :label="item.name"
                                :value="item.id"
                                v-for="item in tag_List" />
                    </el-select>
        </el-form-item>   
        <el-form-item label="文件" prop="files">
            <el-transfer 
            filterable 
            tooltip-effect="dark"
            :render-content="renderFunc"
            :titles="titles" 
            v-model="value" 
            :data="data" 
            el-transfer/>          
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
  import { testingList } from '@/api/deploy/test'
  import { serverList } from '@/api/resource/server'
  import { projectTags } from '@/api/gitlab'
  import { projectList, projectCreate, projectUpdate, projectDelete } from '@/api/deploy/project'
  import infoList from '@/components/mixins/infoList'
  export default {
    name: 'Env',
    mixins: [infoList],
    data() {
      return {
        listApi: testingList,
        dialogFormVisible: false,
        dialogTitle: '项目提测',
        dialogType: '',
        env_List: [],
        project_List: [],
        tag_List: [],
        titles: ["源文件", "目标文件"],
        data: [
          {
            key: 1,
            label: `备选项1111111111111111111111111111111111111111111`,
            disabled: false
          },
          {
            key: 2,
            label: `备选项2`,
            disabled: false
          },
          {
            key: 3,
            label: `备选项3`,
            disabled: false
          },
          {
            key: 4,
            label: `备选项4`,
            disabled: false
          },
        ],
        renderFunc(h, option) {
          return <span title="新增">新增-{ option.label }</span>;
        },
        value: [],
        form: {
          id: '',
          tag: '',
          files: '',
          resource_env_id: '',
          deploy_project_id: '',       
        },
        type: '',
        rules: {
          tag: [
            { required: true, message: '请选择tag', trigger: 'blur' }
          ],  
          files: [
            { required: true, message: '请选择文件', trigger: 'blur' }
          ],           
          deploy_project_id: [
            { required: true, message: '请输入选择项目', trigger: 'blur' }
          ],  
          resource_env_id: [
            { required: true, message: '请输入选择环境', trigger: 'blur' }
          ],                    
        }
      }
    },
    methods: {
      initForm() {
        this.$refs.projectForm.resetFields()
        this.form= {
          id: '',
          tag: '',
          files: '',
          resource_env_id: '',
          deploy_project_id: '',  
        }
      },
      closeDialog() {
        this.initForm()
        this.dialogFormVisible = false
      },
      openDialog(type) {
        switch (type) {
          case 'add':
            this.dialogTitle = '项目提测'
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
        this.GetServerList(this.form.resource_env_id)

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
     async GetServerList(id){
         const ret = await serverList({"page": 1, "pageSize": 9999, "resource_env_id": id})
         if(ret.code === 0){
            this.server_List = ret.data.list
        }
     },
     async EnvChange(row){
        this.project_List = []
        const ret = await projectList({"page": 1, "pageSize": 9999, "resource_env_id": row})
        if(ret.code === 0){
            this.form.deploy_project_id = ''
            this.project_List = ret.data.list
          }
     },
     async ProjectChange(row){
        this.tag_List = []
        const ret = await projectTags({"id": row})
        if(ret.code === 0){
            this.form.tag = ''
            this.tag_List = ret.data.list
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