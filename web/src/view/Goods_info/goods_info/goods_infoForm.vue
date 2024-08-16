<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="商品名字:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入商品名字" />
       </el-form-item>
        <el-form-item label="价格:" prop="price">
          <el-input v-model="formData.price" :clearable="true"  placeholder="请输入价格" />
       </el-form-item>
        <el-form-item label="商品介绍:" prop="bio">
          <el-input v-model="formData.bio" :clearable="true"  placeholder="请输入商品介绍" />
       </el-form-item>
        <el-form-item label="商品id:" prop="goods_id">
          <el-input v-model="formData.goods_id" :clearable="true"  placeholder="请输入商品id" />
       </el-form-item>
        <el-form-item label="商品图:" prop="img">
          <SelectImage v-model="formData.img" file-type="image"/>
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
  createGoods_info,
  updateGoods_info,
  findGoods_info
} from '@/api/Goods_info/goods_info'

defineOptions({
    name: 'Goods_infoForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import SelectImage from '@/components/selectImage/selectImage.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            name: '',
            price: '',
            bio: '',
            goods_id: '',
            img: "",
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findGoods_info({ ID: route.query.id })
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
               res = await createGoods_info(formData.value)
               break
             case 'update':
               res = await updateGoods_info(formData.value)
               break
             default:
               res = await createGoods_info(formData.value)
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
