<template>
  <n-flex vertical style="width: 100%">
      <div>
        <div class="position">图片展示</div>
        <div class="images">
          <div v-for="(item, index) in props" :key="index" class="image-middle">
            <el-card shadow="hover" :body-style="{ padding: '10px' }">
              <el-popover>
                <img :src="item.image_url" class="image"/>
                <img :src="item.image_url" class="imagePreview"/>
              </el-popover>
              <div style="text-align:center;padding-top:12px">
                <span>{{item.qualifier_name}}</span>
              </div>
            </el-card>
          </div>
        </div>
      </div>
  </n-flex>
  <n-button @click="refresh">刷新</n-button>
</template>

<style scoped>

</style>


<script setup>
// import {CertificateImage} from "./certificateImage.vue"
import {useWalletQuery} from "suiue";
import {ElCard, ElPopover,} from "element-plus";
import { ref} from "vue";
import {NFlex} from "naive-ui";
const {objects,loadAllObjects} = useWalletQuery();
loadAllObjects();
console.log("nft objects",objects)
for (var obj in objects){
  console.log(obj)
}

const props = ref([])

const refresh = ()=>{
  loadAllObjects();
  props.value = filteredCerficates()
}


const filteredCerficates = () => {
  const result = []
  for (const key in objects) {
    for (const obj of objects[key]) {
      // if (!obj.display) {
      //   continue
      // }

      let newObj = {}
      Object.assign(newObj, obj.display)
      result.push(newObj)
    }
  }
  console.log("result",result)
  return result
}


</script>
