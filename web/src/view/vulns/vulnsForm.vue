<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="编号:" prop="tfvdId">
          <el-input v-model="formData.tfvdId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="CVE编号:" prop="cveId">
          <el-input v-model="formData.cveId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="cnv_id:" prop="cnvdId">
          <el-input v-model="formData.cnvdId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="CNNVD编号:" prop="cnnvdId">
          <el-input v-model="formData.cnnvdId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="漏洞名称:" prop="vulnName">
          <el-input v-model="formData.vulnName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="CVSS值:" prop="cvss">
          <el-input v-model="formData.cvss" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否研判:" prop="verified">
          <el-switch v-model="formData.verified" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="简介:---" prop="summary">
          <el-input v-model="formData.summary" :clearable="true" placeholder="请输入"   type="textarea" />
       </el-form-item>
        <el-form-item label="漏洞类型 id逗号分隔:" prop="vulnCategory">
          <el-input v-model="formData.vulnCategory" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="修复建议:" prop="vulnPatch">
          <el-input v-model="formData.vulnPatch" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="风险等级" prop="riskLevel">
          <el-switch v-model="formData.riskLevel" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="报警类别 alarmtypeset表的alarmType:" prop="alarmType">
          <el-input v-model.number="formData.alarmType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="可修复性：" prop="repairability">
          <el-input v-model.number="formData.repairability" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="难易度：" prop="difficultyLevel">
          <el-input v-model.number="formData.difficultyLevel" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建人:" prop="creator">
          <el-input v-model="formData.creator" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建人id:" prop="creatorId">
          <el-input v-model.number="formData.creatorId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="更新人:" prop="updater">
          <el-input v-model="formData.updater" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="更新人id:" prop="updaterId">
          <el-input v-model.number="formData.updaterId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="受影响实体 逗号隔开:" prop="affectedEntity">
          <el-input v-model="formData.affectedEntity" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createVulns,
  updateVulns,
  findVulns
} from '@/api/vulns'

defineOptions({
    name: 'VulnsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            tfvdId: '',
            cveId: '',
            cnvdId: '',
            cnnvdId: '',
            vulnName: '',
            cvss: '',
            verified: false,
            summary: '',
            vulnCategory: '',
            vulnPatch: '',
            riskLevel: false,
            alarmType: 0,
            repairability: 0,
            difficultyLevel: 0,
            creator: '',
            creatorId: 0,
            updater: '',
            updaterId: 0,
            affectedEntity: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findVulns({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.revulns
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createVulns(formData.value)
               break
             case 'update':
               res = await updateVulns(formData.value)
               break
             default:
               res = await createVulns(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
