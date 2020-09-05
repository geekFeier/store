<template>
  <section class="hami-price-plan-area mt-50">
    <div class="container">
      <div class="row">
        <div class="col-12">
          <div class="section-heading text-center">
            <h2>sealyun</h2>
            <p>专注于kubernetes安装</p>
          </div>
        </div>
      </div>

      <div class="row justify-content-center">

        <div class="col-12 col-md-6 col-lg-4">
          <div class="single-price-plan mb-100">
            <div class="price-plan-title">
              <h4>大道至简</h4>
            </div>
            <!-- <a href="#" class="btn hami-btn w-100 mb-30">去购买</a> -->
            <div class="price-plan-desc">
              <p>一个二进制工具加一个资源包，不依赖haproxy keepalived ansible等重量级工具，一条命令就可实现kubernetes高可用集群构建，无论是单节点还是集群，单master还是多master，生产还是测试都能很好支持！简单不意味着阉割功能，照样能全量支持kubeadm所有配置。</p>
            </div>
          </div>
        </div>

        <div class="col-12 col-md-6 col-lg-4">
          <div class="single-price-plan active mb-100">
            <!-- <div class="popular-tag"> -->
            <!-- <i class="icon_star"></i> 成为会员 <i class="icon_star"></i> -->
            <!-- </div> -->
            <div class="price-plan-title">
              <h4>服务至上</h4>
            </div>
            <!-- <a href="#" class="btn hami-btn w-100 mb-30">成为会员</a> -->
            <div class="price-plan-desc">
              <p>您只需要69元年费会员就可获得跪舔式售后服务，任何安装不成功都可远程协助，即便是用的免费版本也可获得群内实时反馈问题服务,特殊情况下客服在撸代码或在和煞笔领导吵架时不能及时回复谅解</p>
            </div>
          </div>
        </div>

        <div class="col-12 col-md-6 col-lg-4">
          <div class="single-price-plan  mb-100">

            <div class="price-plan-title">
              <h4>稳定可靠</h4>
            </div>
            <!-- <a href="#" class="btn hami-btn w-100 mb-30">去购买</a> -->
            <div class="price-plan-desc">
              <p>99年证书，再也不用担心生产集群证书过期，ipvs负载多master可用性与稳定性更高，我们有上千用户在生产环境中使用sealos, 我们自身也有超过上千台服务器生产环境中使用sealos</p>
            </div>
          </div>
        </div>

      </div>

      <img src="../../../assets/img/core-img/system.jpg">
    </div>
  </section>

</template>
<script>
import '../../../assets/style.css'
export default {
  name: 'description',
}
</script> 
<style >
.hami-btn {
  position: relative;
  z-index: 1;
  min-width: 150px;
  height: 50px;
  line-height: 50px;
  font-size: 16px;
  font-weight: 600;
  display: inline-block;
  padding: 0 40px;
  text-align: center;
  text-transform: capitalize;
  background-color: #1c66de;
  color: #ffffff;
  border: none;
  border-radius: 100px;
  -webkit-transition-duration: 500ms;
  -o-transition-duration: 500ms;
  transition-duration: 500ms;
}
</style>