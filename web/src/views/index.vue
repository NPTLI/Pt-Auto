<template>
  <div class="min-h-screen dark:text-slate-400 dark:bg-slate-900">
    <Header />
    <div class="max-w-5xl px-4 mx-auto pb-22 sm:px-6 md:px-8 xl:px-12 xl:max-w-6xl">
      <!-- <div class="pt-8 pb-7 sm:pb-8 sm:text-center">
        <h1 class="relative mb-4 text-4xl tracking-tight font-blimone sm:text-5xl lg:text-6xl text-slate-900 dark:text-slate-200">
          {{ appStore.h1 }}
        </h1>
      </div> -->
      <div class="shadow-lg bg-white flex items-center w-full min-h-120px p-4 rounded-2 mb-10 dark:bg-black" v-for="(item, i) in data" :key="i">
        <img class="w-12 h-12 rounded-full mr-2 shrink-0 object-cover" :src="item.avatar" alt="">
        <div class="flex items-center flex-wrap mr-3">
          <a-tooltip class="hidden lg:block" :content="items.label" v-for="(items, index) in item.data" :key="index">
            <div class="flex items-center ml-3 mb-2 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300" @click="toggleTheme()">
              <SvgIcon :name="items.svg" size="small" />
              <div class="text-base text font-bold text-[#999]">{{ items.value }}</div>
            </div>
          </a-tooltip>
        </div>
        <a-switch class="ml-3" checked-color="#3491fb" unchecked-color="#ccc" size="small" />
        <SvgIcon class="ml-3 shrink-0" name="svg-setting" size="small" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useAppStore } from '/@/store/modules/app';
import Header from '/@/components/Header/index.vue';
import Mock from 'mockjs'
const appStore = useAppStore();

// 用户名 等级 数据量 分享率 做种数 做种体积 积分 时魔 入站时间 更新时间 状态 是否开启 配置
const data = ref();
const getList = () => {
  let ary: any = []
  for (let i = 0; i < 10; i++) {
    let obj = {
      avatar: Mock.mock('@image'),
      data: [{
        label: '用户名',
        svg: 'svg-user',
        value: Mock.mock('@name()')
      }, {
        label: '等级',
        svg: 'svg-level',
        value: Mock.mock('@integer(1, 6)')
      }, {
        label: '上传',
        svg: 'svg-upload',
        value: Mock.mock('@integer(0, 1000)')
      }, {
        label: '下载',
        svg: 'svg-download',
        value: Mock.mock('@integer(0, 1000)')
      }, {
        label: '分享率',
        svg: 'svg-share',
        value: Mock.mock('@integer(0, 100)')+'%'
      }, {
        label: '做种数',
        svg: 'svg-seed-num',
        value: Mock.mock('@integer(0, 1000)')
      }, {
        label: '做种体积',
        svg: 'svg-seed',
        value: Mock.mock('@integer(0, 1000)')
      }, {
        label: '积分',
        svg: 'svg-integral',
        value: Mock.mock('@integer(0, 1000)')
      }, {
        label: '时魔',
        svg: 'svg-shimo',
        value: Mock.mock('@integer(0, 1000)')
      }, {
        label: '入站时间',
        svg: 'svg-enter',
        value: Mock.mock('@datetime()')
      }, {
        label: '更新时间',
        svg: 'svg-update',
        value: Mock.mock('@datetime()')
      }, {
        label: '状态',
        svg: 'svg-status',
        value: '在线'
      }]
    }
    ary.push(obj)
  }
  data.value = ary

}
getList()
</script>

<style lang="less" scoped>

</style>
