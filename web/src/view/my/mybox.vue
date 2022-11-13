<!--
 * @Author: Orinten 1595583713@qq.com
 * @Date: 2022-11-13 14:05:25
 * @LastEditors: Orinten 1595583713@qq.com
 * @LastEditTime: 2022-11-13 14:06:01
 * @FilePath: \serverd:\OnGoingDir\221113_Gin_Vue_Admin\gin-vue-admin\web\src\view\my\mybox.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<!-- 此路由可作为父类路由通用路由页面使用 如需自定义父类路由页面 请参考 @/view/superAdmin/index.vue -->
<template>
    <div>
      <router-view v-slot="{ Component }">
        <transition mode="out-in" name="el-fade-in-linear">
          <keep-alive :include="routerStore.keepAliveRouters">
            <component :is="Component" />
          </keep-alive>
        </transition>
      </router-view>
    </div>
  </template>
  
  <script>
  export default {
    name: 'MyBox'
  }
  </script>
  
  <script setup>
  import { useRouterStore } from '@/pinia/modules/router'
  const routerStore = useRouterStore()
  </script>
  