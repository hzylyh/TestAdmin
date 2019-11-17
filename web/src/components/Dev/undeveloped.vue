<!--
 * @Description:首页模块
 * @Author: 吴文周
 * @Github: https://github.com/fodelf
 * @Date: 2019-06-06 18:05:02
 * @LastEditors: 吴文周
 * @LastEditTime: 2019-10-11 11:51:42
 -->
<template>
  <div class="content">
    <div class="mainContent">
      <img src="../../assets/img/undeveloped.png" />
      <span class="undev_word">建设中，敬请期待！</span>
    </div>
  </div>
</template>

<script>
export default {

}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.content {
  height: 100%;
  width: 100%;
  padding-left: 10px;
  padding-right: 10px;
  background-color: #f0f2f5;
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
}
.mainContent {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: center;
  background: white;
}
.undev_word {
  margin-top: 46px;
  font-weight: bold;
  color: #363636;
  font-size: 18px;
  font-family: Source Han Sans CN;
}
</style>
