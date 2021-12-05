<template>
  <!--这里边写模板代码-->
  <div>bala bala</div>
</template>

<script>
import { defineComponent } from "vue";

// 使用defineComponent()定义一个组件并暴露出去，这个组件就是整个文件定义的组件，即使template在上面
export default defineComponent({
  setup() {
    // 这里定义变量，方法等
    return {
      // 这里返回定义的变量，方法等，暴露给组件
    };
  },
});
</script>

<style>
<!--这里编写样式表-- > .c1 {
}
</style>
