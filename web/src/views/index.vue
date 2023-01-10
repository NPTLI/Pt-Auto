<template>
  <div class="min-h-screen bg-gray-50 dark:text-slate-400 dark:bg-black">
    <Header />
    <div class="mt-5 px-4 mx-auto pb-22 md:px-8 xl:px-12">
      <!-- <div class="pt-8 pb-7 md:pb-8 md:text-center">
        <h1 class="relative mb-4 text-4xl tracking-tight font-blimone md:text-5xl lg:text-6xl text-slate-900 dark:text-slate-200">
          {{ appStore.h1 }}
        </h1>
      </div> -->
      <img class="absolute top-0 bottom-0 right-0 left-0 m-auto" v-if="loading" src="../assets/images/loading.gif" alt="">
      <div class="grid gap-5 md:gap-10 grid-cols-1 md:grid-cols-2 xl:grid-cols-3">
        <div class="opacity-0 relative drop-shadow-xl min-h-xs bg-white flex items-center flex-col justify-between w-full p-4 rounded-2 dark:bg-[#1d1d1d]" :class="[item.class]" v-for="(item, i) in data" :key="i">
          <a-switch class="ml-3 absolute right-4 top-4" checked-color="#3491fb" unchecked-color="#ccc" />
          <div class="flex flex-col items-center justify-center">
            <img class="w-16 h-16 md:w-20 md:h-20 xl:w-24 xl:h-24 rounded-full mr-2 shrink-0 object-cover" :src="item.avatar" alt="">
            <div class="text-xl font-bold text-slate-900 mt-3 dark:text-white">{{ item.user_name }}</div>
            <!-- 等级 -->
            <a-tooltip class="hidden lg:block" content="等级">
              <div class="flex items-center mb-4 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300">
                <div class="text-sm text font-bold text-[#999]">{{ item.level }}</div>
              </div>
            </a-tooltip>
          </div>

          <div class="grid grid-cols-4 w-full">
            <!-- 上传 -->
            <a-tooltip class="hidden lg:block" content="上传">
              <div class="flex items-center ml-3 mb-2 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300">
                <SvgIcon name="svg-upload" size="small" />
                <div class="text-base text font-bold text-[#999]">{{ item.upload }}</div>
              </div>
            </a-tooltip>

            <!-- 下载 -->
            <a-tooltip class="hidden lg:block" content="下载">
              <div class="flex items-center ml-3 mb-2 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300">
                <SvgIcon name="svg-download" size="small" />
                <div class="text-base text font-bold text-[#999]">{{ item.download }}</div>
              </div>
            </a-tooltip>

            <!-- 做种数 -->
            <a-tooltip class="hidden lg:block" content="做种数">
              <div class="flex items-center ml-3 mb-2 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300">
                <SvgIcon name="svg-seed-num" size="small" />
                <div class="text-base text font-bold text-[#999]">{{ item.seed_num }}</div>
              </div>
            </a-tooltip>

            <!-- 做种体积 -->
            <a-tooltip class="hidden lg:block" content="做种体积">
              <div class="flex items-center ml-3 mb-2 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300">
                <SvgIcon name="svg-seed" size="small" />
                <div class="text-base text font-bold text-[#999]">{{ item.seed }}</div>
              </div>
            </a-tooltip>
          </div>
          <div class="grid grid-cols-4 w-full">
            <!-- 分享率 -->
            <a-tooltip class="hidden lg:block" content="分享率">
              <div class="flex items-center ml-3 mb-2 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300">
                <SvgIcon name="svg-share" size="small" />
                <div class="text-base text font-bold text-[#999]">{{ item.share }}</div>
              </div>
            </a-tooltip>

            <!-- 时魔 -->
            <a-tooltip class="hidden lg:block" content="时魔">
              <div class="flex items-center ml-3 mb-2 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300">
                <SvgIcon name="svg-shimo" size="small" />
                <div class="text-base text font-bold text-[#999]">{{ item.shimo }}</div>
              </div>
            </a-tooltip>

            <!-- 积分 -->
            <div class="col-span-2">
              <a-tooltip class="hidden lg:block" content="积分">
                <div class="flex items-center ml-3 mb-2 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300">
                  <SvgIcon name="svg-integral" size="small" />
                  <div class="text-base text font-bold text-[#999]">{{ item.integral }}</div>
                </div>
              </a-tooltip>
            </div>
          </div>
          <div class="w-full">
            <!-- 入站时间 -->
            <a-tooltip class="hidden lg:block" content="入站时间">
              <div class="flex items-center ml-3 mb-2 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300">
                <SvgIcon name="svg-enter" size="small" />
                <div class="text-base text font-bold text-[#999]">{{ item.create_time }}</div>
              </div>
            </a-tooltip>

            <!-- 更新时间 -->
            <a-tooltip class="hidden lg:block" content="更新时间">
              <div class="flex items-center ml-3 mb-2 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300">
                <SvgIcon name="svg-update" size="small" />
                <div class="text-base text font-bold text-[#999]">{{ item.update_time }}</div>
              </div>
            </a-tooltip>
          </div>
          <div class="flex items-center justify-end w-full">
            <SvgIcon class="ml-3 shrink-0 cursor-pointer" name="svg-setting" size="small" />
          </div>
        </div>
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
const loading = ref(true)
const getList = () => {
  let ary: any = []
  for (let i = 0; i < 10; i++) {
    let obj = {
      avatar: Mock.mock('@image'),
      user_name: Mock.mock('@name()'),
      level: Mock.mock('@cword(5)'),
      upload: Mock.mock('@integer(0, 10000)'),
      download: Mock.mock('@integer(0, 10000)'),
      share: Mock.mock('@integer(0, 100)') + '%',
      seed_num: Mock.mock('@integer(0, 10000)'),
      seed: Mock.mock('@integer(0, 10000)'),
      integral: Mock.mock('@integer(10000, 1000000)'),
      shimo: Mock.mock('@integer(0, 10000)'),
      create_time: Mock.mock('@datetime()'),
      update_time: Mock.mock('@datetime()'),
      status: '在线',
      open: false,
      class: `animation${i + 1}`,
    }
    ary.push(obj)
  }
  setTimeout(() => {
    loading.value = false
    data.value = ary
  }, 2000);

}
getList()
</script>

<style lang="scss" scoped>
@for $i from 1 to 16 {
  .animation#{$i} {
    animation: fade-in-bottom .8s forwards $i*0.2s;
  }
}
</style>
