<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="订单号:" prop="out_trade_no">
          <el-input v-model="formData.out_trade_no" :clearable="true"  placeholder="请输入订单号" />
       </el-form-item>
        <el-form-item label="订单状态:" prop="status">
          <el-input v-model="formData.status" :clearable="true"  placeholder="请输入订单状态" />
       </el-form-item>
        <el-form-item label="卡密:" prop="key">
          <el-input v-model="formData.key" :clearable="true"  placeholder="请输入卡密" />
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
  createGet_pay,
  updateGet_pay,
  findGet_pay
} from '@/api/GetPay/get_pay'

defineOptions({
    name: 'Get_payForm'
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
            out_trade_no: '',
            status: '',
            key: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findGet_pay({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
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
               res = await createGet_pay(formData.value)
               break
             case 'update':
               res = await updateGet_pay(formData.value)
               break
             default:
               res = await createGet_pay(formData.value)
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
