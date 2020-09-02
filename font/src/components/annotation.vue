<template>
  <!-- <section class="find-domain-area section-padding-100-0">
    <div class="container">
      <div class="row align-items-center">
        <div class="col-12 col-md-4">
          <div class="domain-text mb-100">
            <h2>公告111</h2>
         <h6>公告描述。。。。。。</h6>  
          </div>
        </div>
        <div class="col-12 col-md-8">
          <div class="domain-search-form mb-100" style="font-size:15px">
            kubernetes集群离线安装包, 一条命令高可用<br />
            使用kubeadm安装，包含安装所需的所有二进制文件，镜像文件，systemd配置，yaml配置与一些简单的启动脚本，五分钟搞定kubernetes安装~ 遇到任何问题可进群咨询，安装出问题提供远程协助,内核负载高可用，不依赖haproxy keepalived ansible,99年证书
            <div class="domain-price-help mt-30 d-flex align-items-center justify-content-between">
                <p>$公告一</p>
                <p>公告二</p>
                <p>公告二</p>
                <p>公告二</p>
                <p>公告二</p>
                <p>公告二</p>
              </div>
          </div>
        </div>
      </div>
    </div> 
  </section> -->

  <div>
    <div class="annotation" v-if="hasAnnotaion">
      <div class="annotation-left">
        <Icon type="ivu-icon ios-bell-outline"></Icon> 公告
      </div>
      <div class="annotation-right">

        <marquee direction="left" align="bottom" width="100%" onmouseout="this.start()" onmouseover="this.stop()" scrollamount="7" scrolldelay="1">使用kubeadm安装，包含安装所需的所有二进制文件，镜像文件，systemd配置，yaml配置与一些简单的启动脚本，五分钟搞定kubernetes安装~ 遇到任何问题可进群咨询，安装出问题提供远程协助,内核负载高可用，不依赖haproxy keepalived ansible,99年证书</marquee>

      </div>
    </div>
  </div>
</template>
 
<script type="text/javascript">
export default {
  name: 'annotation',
  props: ['hasAnnotaion'],
  data() {},
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