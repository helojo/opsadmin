<template>
  <div>
      <el-form :model="form" :rules="rules" label-width="100px" ref="projectForm">
        <el-form-item label="环境" prop="environment_id">
                    <el-select  @change="EnvChange" filterable placeholder="请选择" style="width:32.4%" v-model="form.environment_id">
                        <el-option
                                :key="item.id"
                                :label="item.name"
                                :value="item.id"
                                v-for="item in env_List" />
                    </el-select>
        </el-form-item>
        <el-form-item  label="项目" prop="deploy_project_id">
                    <el-select @change="ProjectChange" filterable placeholder="请选择" style="width:32.4%" v-model="form.deploy_project_id">
                        <el-option
                                :key="item.id"
                                :label="item.name"
                                :value="item.id"
                                v-for="item in project_List" />
                    </el-select>
        </el-form-item>
        <el-form-item  label="Tag" prop="tag">
                    <el-select filterable placeholder="请选择" style="width:32.4%" v-model="form.tag">
                        <el-option
                                :key="item.id"
                                :label="item.name"
                                :value="item.id"
                                v-for="item in tag_List" />
                    </el-select>
        </el-form-item>
          <el-form-item label="描述" prop="describe"  style="width:32.4%">
              <el-input autocomplete="off" type="textarea" v-model="form.describe"></el-input>
          </el-form-item>
          <el-form-item label="文件对比" prop="files">
            <el-table
                    :data="files_list"
                    style="width: 100%">
                <el-table-column
                        prop="key"
                        label="待同步文件">
                    <template slot-scope="scope">
                        <span class="operate-span-danger" v-if="scope.row.key  && scope.row.key.indexOf('删除') != -1 " > {{ scope.row.key }}</span>
                        <span class="operate-span-primary" v-else >{{ scope.row.key }}</span>
                    </template>
                </el-table-column>
            </el-table>
        </el-form-item>   

      </el-form>
      <div class="dialog-footer" slot="footer" style="float: right" >
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="Contrast" type="primary">比较</el-button>
        <el-button :disabled="CommitButton"  @click="enterDialog" type="primary">提交</el-button>
      </div>
  </div>
</template>


<script>
  // 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索
  import { envList } from '@/api/resource/env'
  import { testingList, testingContrast, testingRelease} from '@/api/deploy/test'
  import { serverList } from '@/api/resource/server'
  import { projectTags } from '@/api/gitlab'
  import { projectList } from '@/api/deploy/project'
  import infoList from '@/components/mixins/infoList'
  export default {
    name: 'Testing',
    mixins: [infoList],
    data() {
      return {
        listApi: testingList,
        dialogTitle: '项目提测',
        dialogType: '',
        env_List: [],
        project_List: [],
        tag_List: [],
        titles: ["源文件", "目标文件"],
        files_list: [],
        taget_file_list: [],
        path: "",
        CommitButton: true,
        value: [],
        form: {
          id: '',
          tag: '',
          files: '',
          environment_id: '',
          deploy_project_id: '',
          describe: '',
        },
        type: '',
        rules: {
          tag: [
            { required: true, message: '请选择tag', trigger: 'blur' }
          ],
          deploy_project_id: [
            { required: true, message: '请输入选择项目', trigger: 'blur' }
          ],
          environment_id: [
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
          environment_id: '',
          deploy_project_id: '',  
        },
        this.files_list = []
      },
      closeDialog() {
        this.initForm()
        this.taget_file_list = []
        this.files_list = []
      },
      openDialog(type) {
        switch (type) {
          case 'add':
            this.dialogTitle = '项目提测'
            break
        }
      },
      async Contrast(){
       this.$refs.projectForm.validate(async valid => {
          if (valid) {
              {
                  this.files_list = []
                  const res = await testingContrast(this.form)
                  if (res.code === 0) {
                      this.form.files = ""
                     this.files_list = res.data.list
                     this.path = res.data.path
                     this.CommitButton = false
                     this.taget_file_list = []
                 }
              }
          }
        })
      },
      async enterDialog() {
        this.$refs.projectForm.validate(async valid => {
          if (valid) {
              if (this.taget_file_list.length === 0) {
                  this.form.files = this.taget_file_list
                  this.form.path = this.path
                  console.log(this.form)
                  const res = await testingRelease(this.form)
                  if (res.code === 0) {
                      this.$message({
                          type: 'success',
                          message: '提侧成功',
                          showClose: true
                      })
                      this.initForm()
                  }
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
         const ret = await serverList({"page": 1, "pageSize": 9999, "environment_id": id})
         if(ret.code === 0){
            this.server_List = ret.data.list
        }
     },
     async EnvChange(row){
        this.project_List = []
        const ret = await projectList({"page": 1, "pageSize": 9999, "environment_id": row})
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