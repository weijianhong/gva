<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户名:" prop="username">
          <el-input v-model="formData.username" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="密码:" prop="password">
          <el-input v-model="formData.password" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="指纹ID:" prop="fingerprintId">
          <el-input v-model.number="formData.fingerprintId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="指纹名称:" prop="fingerprintName">
          <el-input v-model="formData.fingerprintName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="常用端口:" prop="port">
          <el-input v-model="formData.port" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="协议:" prop="protocol">
          <el-input v-model="formData.protocol" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="厂商/公司/组织:" prop="corpId">
          <el-input v-model.number="formData.corpId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态 1有效，0无效，默认1有效:" prop="status">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="creator字段:" prop="creator">
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
  createWeakpass,
  updateWeakpass,
  findWeakpass
} from '@/api/weakpass'

defineOptions({
    name: 'WeakpassForm'
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
            username: '',
            password: '',
            fingerprintId: 0,
            fingerprintName: '',
            port: '',
            protocol: '',
            corpId: 0,
            status: false,
            creator: '',
            creatorId: 0,
            updater: '',
            updaterId: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWeakpass({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reweakpass
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
               res = await createWeakpass(formData.value)
               break
             case 'update':
               res = await updateWeakpass(formData.value)
               break
             default:
               res = await createWeakpass(formData.value)
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
