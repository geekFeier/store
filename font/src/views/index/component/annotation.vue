<template>
  <div>
    <div class="annotation" v-if="hasAnnotaion">
      <div class="annotation-left">
        <Icon type="ivu-icon ios-bell-outline"></Icon> 公告
      </div>
      <div class="annotation-right">
        <marquee direction="left" align="bottom" width="100%" onmouseout="this.start()" onmouseover="this.stop()" scrollamount="7" scrolldelay="1">只能用丝滑一词形容的kubernetes高可用安装（kubernetes install）工具，一条命令，离线安装，包含所有依赖，内核负载不依赖haproxy keepalived,纯golang开发,99年证书,支持v1.16 v1.15 v1.17 v1.18 v1.19! <a href="https://github.com/fanux/sealos">github地址</a></marquee>
      </div>
    </div>
  </div>
</template>
 
<script type="text/javascript">
export default {
  name: 'annotation',
  props: ['hasAnnotaion'],
  data() {
    return {}
  },
}
</script> 
<style lang="css">
.annotation {
  position: fixed !important;
  top: 0;
  left: 0;
  padding: 0 40px;
  width: 100%;
  height: 60px;
  line-height: 60px;
  background: #161616;
  color: #fff;
  z-index: 10;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}
.annotation-left {
  float: left;
  position: relative;
  margin: 0;
  font-size: 24px;
  line-height: 60px;
  font-weight: 300;
  letter-spacing: -0.5px;
}
.annotation-right {
  padding-left: 100px;
  font-size: 15px;
}
</style>