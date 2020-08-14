<template>
  <div>
      <div class="button-box clearflex">
          <el-form :inline="true" class="demo-form-inline">
              <el-form-item style="float: right" >
                  <el-button @click="Refresh" type="primary">刷新列表 </el-button>
              </el-form-item>
              <el-form-item style="float: right" >
                  <el-button @click="openTestRollbackDialog" type="primary">项目回滚</el-button>
              </el-form-item>
              <el-form-item style="float: right" >
                  <el-button @click="openTesttingDialog" type="primary">项目发布</el-button>
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
        <el-table-column label="提测版本" min-width="150" prop="version"></el-table-column>
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
                label="备份是否删除"
                prop="isdelete"
                type="scope">
            <template slot-scope="scope">
                <span class="operate-span-primary" v-if="scope.row.isdelete === 1 " >否</span>
                <span class="operate-span-danger" v-else-if="scope.row.isdelete === 2 " >是</span>
            </template>
        </el-table-column>
        <el-table-column
                    label="状态"
                    prop="status"
                    type="scope">
                <template slot-scope="scope">
                    <span class="operate-span-primary" v-if="scope.row.status === 0 " >开发审核</span>
                    <span class="operate-span-primary" v-else-if="scope.row.status === 1 " >测试审核</span>
                    <span class="operate-span-primary" v-else-if="scope.row.status === 2 " >运维审核</span>
                    <span class="operate-span-primary" v-else-if="scope.row.status === 3 " >已上线</span>
                    <span class="operate-span-danger"  v-else-if="scope.row.status === 4"  >审核中</span>
                    <span class="operate-span-danger"  v-else-if="scope.row.status === 5"  >已关闭</span>
                </template>
      </el-table-column>            
      <el-table-column label="申请人" min-width="150" prop="applicant"></el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
            <el-button @click="handleAuditor(scope.row)" size="mini" type="primary" v-if="scope.row.status === 0 ">开发审核</el-button>
            <el-button @click="handleAuditor(scope.row)" size="mini" type="primary"  v-else-if="scope.row.status === 1 " >测试审核</el-button>
            <el-button @click="handleAuditor(scope.row)" size="mini" type="primary" v-else-if="scope.row.status === 2 " >运维审核</el-button>
            <el-button @click="handleAuditor(scope.row)" size="mini" type="primary" v-else-if="scope.row.status >= 3">详情</el-button>
            <el-button @click="handleClose(scope.row)" size="mini" type="danger" v-if="scope.row.status !== 5 && scope.row.status !== 4 && scope.row.status !== 3">关闭</el-button>
          <el-button @click="TestingLogs(scope.row)" size="small" type="primary" >日志</el-button>
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

      <!-- 项目提测模态框 -->
      <el-dialog :before-close="closeDialog" :title="dialogTitle" :visible.sync="dialogFormTesttingVisible">
          <el-form :model="form" :rules="rules" label-width="100px" ref="OnlinesForm">
              <el-form-item label="环境" prop="environment_id">
                  <el-select  @change="EnvChange" filterable placeholder="请选择" v-model="form.environment_id">
                      <el-option
                              :key="item.id"
                              :label="item.name"
                              :value="item.id"
                              v-for="item in env_List" />
                  </el-select>
              </el-form-item>
              <el-form-item  label="项目" prop="deploy_project_id">
                  <el-select @change="ProjectChange" filterable placeholder="请选择" v-model="form.deploy_project_id">
                      <el-option
                              :key="item.id"
                              :label="item.name"
                              :value="item.id"
                              v-for="item in project_List" />
                  </el-select>
              </el-form-item>
              <el-form-item  label="Tag" prop="tag">
                  <el-select filterable placeholder="请选择" v-model="form.tag">
                      <el-option
                              :key="item.id"
                              :label="item.name"
                              :value="item.id"
                              v-for="item in tag_List" />
                  </el-select>
              </el-form-item>
              <el-form-item label="描述" prop="describe">
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
          <div class="dialog-footer" slot="footer">
              <el-button @click="closeDialog">取 消</el-button>
              <el-button @click="Contrast" type="primary">比较</el-button>
              <el-button :disabled="CommitButton"  @click="enterDialog" type="primary">提交</el-button>
          </div>
      </el-dialog>  <!-- 项目提测模态框 -->

      <!-- 项目回滚模态框 -->
      <el-dialog :before-close="closeDialog" :title="dialogTitle" :visible.sync="dialogFormTestRollbackVisible">
          <el-form :model="form" :rules="rules" label-width="100px" ref="OnlinesForm">
              <el-form-item label="环境" prop="environment_id">
                  <el-select  @change="EnvChange" filterable placeholder="请选择" v-model="form.environment_id">
                      <el-option
                              :key="item.id"
                              :label="item.name"
                              :value="item.id"
                              v-for="item in env_List" />
                  </el-select>
              </el-form-item>
              <el-form-item  label="项目" prop="deploy_project_id">
                  <el-select @change="ProjectRollbackChange" filterable placeholder="请选择" v-model="form.deploy_project_id">
                      <el-option
                              :key="item.id"
                              :label="item.name"
                              :value="item.id"
                              v-for="item in project_List" />
                  </el-select>
              </el-form-item>
              <el-form-item  label="回滚版本" prop="version">
                  <el-select filterable placeholder="请选择" v-model="form.version">
                      <el-option
                              :key="item.ID"
                              :label="item.version"
                              :value="item.ID"
                              v-for="item in version_List" />
                  </el-select>
              </el-form-item>
              <el-form-item label="描述" prop="describe">
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
          <div class="dialog-footer" slot="footer">
              <el-button @click="closeDialog">取 消</el-button>
              <el-button @click="RollbackContrast" type="primary">比较</el-button>
              <el-button :disabled="CommitButton"  @click="enterRollback" type="primary">提交</el-button>
          </div>
      </el-dialog>  <!-- 项目回滚模态框 -->

      <!-- 审核工单模态框 -->
      <el-dialog
              :visible.sync="dialogVisibleForaudit"
              @close="handleCancelExamine"
              title="审核工单"
              width="50%">
          <el-steps :active="active" finish-status="success" simple style="margin-top: 20px">
              <el-step :key="item"
                       :title="item"
                       v-for="item in online_status"/>
          </el-steps>
          <el-form :model="form" label-width="100px" ref="OnlinesForm">
              <el-form-item label="项目名称" prop="deployproject">
                  <el-input placeholder="" readonly v-model="form.deployproject.name"/>
              </el-form-item>
              <el-form-item label="上线版本" prop="version">
                  <el-input placeholder="" readonly v-model="form.version"/>
              </el-form-item>
              <el-form-item label="描述" prop="describe">
                  <el-input placeholder="" readonly v-model="form.describe"/>
              </el-form-item>
              <el-form-item label="开发审核" prop="dev_auditor" v-if="form.dev_auditor !== ''">
                  <el-input placeholder="" readonly v-model="form.dev_auditor"/>
              </el-form-item>
              <el-form-item label="测试审核" prop="test_auditor" v-if="form.test_auditor !== ''">
                  <el-input placeholder="" readonly v-model="form.test_auditor"/>
              </el-form-item>
              <el-form-item label="运维审核" prop="op_auditor" v-if="form.op_auditor !== ''">
                  <el-input placeholder="" readonly v-model="form.op_auditor"/>
              </el-form-item>
          </el-form>
          <div class="dialog-footer" slot="footer">
              <el-button @click="handleauditDialog" type="primary" v-if="form.status !== 3 && form.status !== 5 && form.status !== 4 && form.status !== 6">通过</el-button>
              <el-button @click="handleCancelExamine">取消</el-button>
          </div>
      </el-dialog>

  </div>
</template>


<script>
  // 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索
  import { onlineList, onlineContrast, onlineCreate, onlineRversion, devAudit, testAudit, opsAudit} from '@/api/deploy/online'
  import { rollbackContrast, rollbackRelease } from '@/api/deploy/rollback'
  import { envList } from '@/api/resource/env'
  import { projectTags } from '@/api/gitlab'
  import { projectList } from '@/api/deploy/project'
  import { serverList } from '@/api/resource/server'
  import infoList from '@/components/mixins/infoList'
  export default {
    name: 'Env',
    mixins: [infoList],
    data() {
      return {
        listApi: onlineList,
        dialogFormVisible: false,
        dialogFormTesttingVisible: false,
        dialogFormTestRollbackVisible: false,
        dialogVisibleForaudit: false,
        dialogTitle: '项目提测',
        dialogType: '',
        result: '',
        env_List: [],
        project_List: [],
        tag_List: [],
        files_list: [],
        version_List: [],
        online_status: [],
        CommitButton: true,
        active: 1,
        form: {
              ID: '',
              tag: '',
              files: '',
              version: '',
              status: '',
              environment_id: '',
              deploy_project_id: '',
              describe: '',
              dev_auditor: '',
              test_auditor: '',
              op_auditor: '',
              deployproject: [],
          },
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
          },
      }

    },
    methods: {
      initForm() {
          this.$refs.OnlinesForm.resetFields()
            this.form= {
                ID: '',
                tag: '',
                files: '',
                version: '',
                status: '',
                environment_id: '',
                deploy_project_id: '',
                describe: '',
                dev_auditor: '',
                test_auditor: '',
                op_auditor: '',
                deployproject: [],
            },
                this.files_list = []
      },
      closeresultDialog(){
        this.dialogFormVisible = false
      },
      closeDialog() {
        this.initForm()
        this.dialogFormTesttingVisible = false
        this.dialogFormTestRollbackVisible = false
        this.taget_file_list = []
        this.files_list = []
      },
      async TestingLogs(row) {
        this.dialogTitle = '日志详情'
        this.result = row.result
        this.dialogFormVisible = true

      },
      async openTesttingDialog(){
          this.dialogTitle = '项目提侧'
          this.GetEnvList()
          this.dialogFormTesttingVisible = true
      },
        async enterDialog() {
            this.$refs.OnlinesForm.validate(async valid => {
                if (valid) {
                    this.form.files = this.taget_file_list
                    this.form.path = this.path
                    const res = await onlineCreate(this.form)
                    if (res.code === 0) {
                        this.$message({
                            type: 'success',
                            message: '提交成功',
                        })
                     this.dialogFormTesttingVisible = false
                     this.getTableData()
                     this.closeDialog()
                    }

                }
            })
        },
     async GetEnvList(){
            this.env_List = []
            const ret = await envList({"page": 1, "status": 2, "pageSize": 9999})
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
        },
        async Contrast(){
            this.$refs.OnlinesForm.validate(async valid => {
                if (valid) {
                    {
                        this.files_list = []
                        this.form.path = this.path
                        const res = await onlineContrast(this.form)
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
        async openTestRollbackDialog(){
            this.dialogTitle = '项目回滚'
            this.GetEnvList()
            this.dialogFormTestRollbackVisible = true
        },
        async ProjectRollbackChange(row) {
          console.log(row)
            const res = await onlineRversion({"id": row})
            if (res.code === 0) {
                this.version_List = res.data.list
            }
        },
        async RollbackContrast(){
            this.form.tag = "rollback"
            this.$refs.OnlinesForm.validate(async valid => {
                if (valid) {
                    {
                        const res = await rollbackContrast(this.form)
                        if (res.code === 0) {
                            console.log(res)
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
        async enterRollback(){
          console.log(this.form)
            const res = await rollbackRelease(this.form)
            if (res.code === 0) {
                this.$message({
                    type: 'success',
                    message: '回滚操作成功',
                })
                this.dialogFormTestRollbackVisible = false
                this.getTableData()
                this.closeDialog()
            }
        },
        handleCancelExamine() {
            this.dialogVisibleForaudit = false
            this.getTableData()
        },
        // 审核工单
        async handleAuditor(row) {
            for(let key in this.form){
                this.form[key] = row[key]
            }
            this.dialogVisibleForaudit  = true
            this.online_status = ['开发审核', '测试审核', '运维审核']
            this.active = row.status
        },
        // 通过审核
        async handleauditDialog(){
            switch (this.form.status) {
                // 开发审核
                case 0:
                {
                    const res = await devAudit({"id": this.form.ID})
                    if (res.code === 0) {
                        this.$message({
                            type: 'success',
                            message: '开发审核成功!'
                        })
                        this.dialogVisibleForaudit = false
                        this.getTableData()
                        this.closeDialog()
                    }
                }
                    break;
                case 1 :
                {
                    const res = await testAudit({"id": this.form.ID})
                    if (res.code === 0) {
                        this.$message({
                            type: 'success',
                            message: '测试审核成功!'
                        })
                        this.dialogVisibleForaudit = false
                        this.getTableData()
                        this.closeDialog()
                    }
                }
                    break;
                case 2:
                {
                    const res = await opsAudit({"id": this.form.ID})
                    if (res.code === 0) {
                        this.$message({
                            type: 'success',
                            message: '运维审核成功!'
                        })
                        this.dialogVisibleForaudit = false
                        this.getTableData()
                        this.closeDialog()
                    }
                }
                    break;
            }
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