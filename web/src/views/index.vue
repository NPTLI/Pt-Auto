<template>
  <div class="min-h-screen dark:text-slate-400 dark:bg-slate-900">
    <Header />
    <div class="mt-5 px-4 mx-auto pb-22 sm:px-6 md:px-8 xl:px-12">
      <!-- <div class="pt-8 pb-7 sm:pb-8 sm:text-center">
        <h1 class="relative mb-4 text-4xl tracking-tight font-blimone sm:text-5xl lg:text-6xl text-slate-900 dark:text-slate-200">
          {{ appStore.h1 }}
        </h1>
      </div> -->
      <div class="grid gap-10 grid-cols-1 sm:grid-cols-2 xl:grid-cols-3">
        <div class="shadow-lg min-h-xs bg-white flex items-center flex-col justify-between w-full p-4 rounded-2 mb-10 dark:bg-black" v-for="(item, i) in data" :key="i">
          <div class="flex flex-col items-center justify-center">
            <img class="w-16 h-16 sm:w-20 sm:h-20 xl:w-24 xl:h-24 rounded-full mr-2 shrink-0 object-cover" :src="item.avatar" alt="">
            <div class="text-xl font-bold text-slate-900 mt-3">{{ item.user_name }}</div>
          </div>
          <div class="flex items-center flex-wrap mr-3">
            <!-- 等级 -->
            <a-tooltip class="hidden lg:block" content="等级">
              <div class="flex items-center ml-3 mb-2 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300" @click="toggleTheme()">
                <SvgIcon class="block lg:hidden" name="svg-level" size="small" />
                <SvgIcon class="hidden lg:block" name="svg-level" size="default" />
                <div class="text-base text font-bold text-[#999]">{{ item.level }}</div>
              </div>
            </a-tooltip>
            <!-- 上传 -->
            <a-tooltip class="hidden lg:block" content="上传">
              <div class="flex items-center ml-3 mb-2 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300" @click="toggleTheme()">
                <SvgIcon class="block lg:hidden" name="svg-upload" size="small" />
                <SvgIcon class="hidden lg:block" name="svg-upload" size="default" />
                <div class="text-base text font-bold text-[#999]">{{ item.upload }}</div>
              </div>
            </a-tooltip>

            <!-- 下载 -->
            <a-tooltip class="hidden lg:block" content="下载">
              <div class="flex items-center ml-3 mb-2 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300" @click="toggleTheme()">
                <SvgIcon class="block lg:hidden" name="svg-download" size="small" />
                <SvgIcon class="hidden lg:block" name="svg-download" size="default" />
                <div class="text-base text font-bold text-[#999]">{{ item.download }}</div>
              </div>
            </a-tooltip>

            <!-- 分享率 -->
            <a-tooltip class="hidden lg:block" content="分享率">
              <div class="flex items-center ml-3 mb-2 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300" @click="toggleTheme()">
                <SvgIcon class="block lg:hidden" name="svg-share" size="small" />
                <SvgIcon class="hidden lg:block" name="svg-share" size="default" />
                <div class="text-base text font-bold text-[#999]">{{ item.share }}</div>
              </div>
            </a-tooltip>

            <!-- 做种数 -->
            <a-tooltip class="hidden lg:block" content="做种数">
              <div class="flex items-center ml-3 mb-2 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300" @click="toggleTheme()">
                <SvgIcon class="block lg:hidden" name="svg-seed-num" size="small" />
                <SvgIcon class="hidden lg:block" name="svg-seed-num" size="default" />
                <div class="text-base text font-bold text-[#999]">{{ item.seed_num }}</div>
              </div>
            </a-tooltip>

            <!-- 做种体积 -->
            <a-tooltip class="hidden lg:block" content="做种体积">
              <div class="flex items-center ml-3 mb-2 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300" @click="toggleTheme()">
                <SvgIcon class="block lg:hidden" name="svg-seed" size="small" />
                <SvgIcon class="hidden lg:block" name="svg-seed" size="default" />
                <div class="text-base text font-bold text-[#999]">{{ item.seed }}</div>
              </div>
            </a-tooltip>

            <!-- 积分 -->
            <a-tooltip class="hidden lg:block" content="积分">
              <div class="flex items-center ml-3 mb-2 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300" @click="toggleTheme()">
                <SvgIcon class="block lg:hidden" name="svg-integral" size="small" />
                <SvgIcon class="hidden lg:block" name="svg-integral" size="default" />
                <div class="text-base text font-bold text-[#999]">{{ item.integral }}</div>
              </div>
            </a-tooltip>

            <!-- 时魔 -->
            <a-tooltip class="hidden lg:block" content="时魔">
              <div class="flex items-center ml-3 mb-2 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300" @click="toggleTheme()">
                <SvgIcon class="block lg:hidden" name="svg-shimo" size="small" />
                <SvgIcon class="hidden lg:block" name="svg-shimo" size="default" />
                <div class="text-base text font-bold text-[#999]">{{ item.shimo }}</div>
              </div>
            </a-tooltip>

            <!-- 入站时间 -->
            <a-tooltip class="hidden lg:block" content="入站时间">
              <div class="flex items-center ml-3 mb-2 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300" @click="toggleTheme()">
                <SvgIcon class="block lg:hidden" name="svg-enter" size="small" />
                <SvgIcon class="hidden lg:block" name="svg-enter" size="default" />
                <div class="text-base text font-bold text-[#999]">{{ item.create_time }}</div>
              </div>
            </a-tooltip>

            <!-- 更新时间 -->
            <a-tooltip class="hidden lg:block" content="更新时间">
              <div class="flex items-center ml-3 mb-2 text-slate-400 cursor-pointer hover:text-slate-500 dark:hover:text-slate-300" @click="toggleTheme()">
                <SvgIcon class="block lg:hidden" name="svg-update" size="small" />
                <SvgIcon class="hidden lg:block" name="svg-update" size="default" />
                <div class="text-base text font-bold text-[#999]">{{ item.update_time }}</div>
              </div>
            </a-tooltip>


          </div>
          <div class="flex items-center justify-between w-full">
            <a-switch class="ml-3" checked-color="#3491fb" unchecked-color="#ccc" size="small" />
            <SvgIcon class="ml-3 shrink-0" name="svg-setting" size="small" />
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
const getList = () => {
  let ary: any = []
  for (let i = 0; i < 10; i++) {
    let obj = {
      avatar: Mock.mock('@image'),
      user_name: Mock.mock('@name()'),
      level: Mock.mock('@integer(1, 6)'),
      upload: Mock.mock('@integer(0, 1000)'),
      download: Mock.mock('@integer(0, 1000)'),
      share: Mock.mock('@integer(0, 100)') + '%',
      seed_num: Mock.mock('@integer(0, 1000)'),
      seed: Mock.mock('@integer(0, 1000)'),
      integral: Mock.mock('@integer(10000, 1000000)'),
      shimo: Mock.mock('@integer(0, 1000)'),
      create_time: Mock.mock('@datetime()'),
      update_time: Mock.mock('@datetime()'),
      status: '在线',
      open: false,
    }
    ary.push(obj)
  }
  data.value = ary

}
getList()
</script>

<style lang="less" scoped>

</style>
