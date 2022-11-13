<!--
 * @Author: Orinten 1595583713@qq.com
 * @Date: 2022-11-13 13:51:19
 * @LastEditors: Orinten 1595583713@qq.com
 * @LastEditTime: 2022-11-13 15:43:55
 * @FilePath: \serverd:\OnGoingDir\221113_Gin_Vue_Admin\gin-vue-admin\web\src\view\my\index.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
    <div>
{{myName}}
    </div>
</template>

<script>
export default {
        name:'My',
}
</script>

<script setup>
import{ ref } from 'vue'
import{ useRoute } from 'vue-router'
const myName = ref('name')
const route = useRoute()
console.log(route.params.id)
</script>
