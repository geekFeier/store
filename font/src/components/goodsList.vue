<template>
  <div>
    <section class="hami-price-plan-area mt-50">
      <div class="container">
        <div class="row">
          <div class="col-12">
            <div class="section-heading text-center">
              <h2>商品列表</h2>
              <p>
                您还可获取专有分享链接，享受60%交易提成

                <Poptip title="专有分享链接 - 通过sealyun赚钱" placement="right-end">
                  <div slot="content" class="poptip">
                    <p>{{ shareLink }}</p>
                    <p>任何用户通过上面链接访问网站并成功交易您将获得交易的60%提成</p>
                    <p>如嵌入自己的markdown文档中, 发到群里，或者把链接直接发给有需要的朋友</p>
                  </div>
                  <span id="sharelink">推广链接</span>
                </Poptip>
              </p>
            </div>
          </div>
        </div>
        <ul>
          <li v-for="p in products" class="li">
            <div>
              <a :href="p.url" target="_blank" rel="noopener">{{ p.name }}离线安装包</a>
              <Tooltip content="购买完自动跳转下载，已付款点击自动下载">
                <a :href="p.url" target="_blank" rel="noopener">
                  <Button id="buy" type="success">点击购买 {{ p.price }}元</Button>
                </a>
              </Tooltip>
            </div>
          </li>
        </ul>
      </div>
    </section>

    <Back-top></Back-top>
  </div>
  <!-- <div class="hello">
    <img src="https://s2.ax1x.com/2019/09/12/nBpyZt.jpg" style="width:300px" />
    <br>kubernetes集群离线安装包, 一条命令高可用
    <p>
      使用kubeadm安装，包含安装所需的所有二进制文件，镜像文件，systemd配置，yaml配置与一些简单的启动脚本，五分钟搞定kubernetes安装~ 遇到任何问题可进群咨询，安装出问题提供远程协助,内核负载高可用，不依赖haproxy keepalived ansible,99年证书
    </p>
    <h3>商品列表</h3>
    <div id="pro-link">
      <ul>
        <li v-for="p in products" class="li">
          <div>
            <a :href="p.url" target="_blank" rel="noopener">{{ p.name }}离线安装包</a>
            <Tooltip content="购买完自动跳转下载，已付款点击自动下载">
              <a :href="p.url" target="_blank" rel="noopener">
                <Button id="buy" type="success">点击购买 {{ p.price }}元</Button>
              </a>
            </Tooltip>
            <Tooltip content="获取专有分享链接，享受60%交易提成">
              <Button @click="share = true" id="sharelink" type="success">推广链接</Button>
              <Modal v-model="share" title="专有分享链接 - 通过sealyun赚钱" @on-ok="ok" @on-cancel="cancel">
                <p>{{ shareLink }}</p>
                <p>任何用户通过上面链接访问网站并成功交易您将获得交易的60%提成</p>
                <p>如嵌入自己的markdown文档中, 发到群里，或者把链接直接发给有需要的朋友</p>
              </Modal>
            </Tooltip>
          </div>
        </li>
      </ul>
    </div>
    <h3>LINKS</h3>
    <ul>
      <li>
        <img src="https://sealyun.com/kubernetes-qrcode.jpg">
      </li>
      <br>公众号
      <br>
      <li>联系方式：fhtjob@hotmail.com</li>
    </ul>
  </div> -->
</template>

<script  type="text/javascript">
import VueCookies from 'vue-cookies'
export default {
  props: {
    msg: String,
  },
  methods: {
    ok() {
      // this.$Message.info('Clicked ok')
    },
  },
  data: function () {
    var d = {
      time: 'http://store.lameleg.com:8080/pro/kubernetes1.13.1?time=' + new Date().getTime(),
      shareLink: '',
      share: false,
      products: [],
    }
    if (typeof this.$route.query.referrer != 'undefined') {
      VueCookies.set('referrer', this.$route.query.referrer)
    }
    if (VueCookies.get('referrer') != null) {
      d.time += '&referrer=' + VueCookies.get('referrer')
    }

    this.$http
      .get('http://store.lameleg.com:8080/loginless/info/user', {
        credentials: true,
      })
      .then(function (res) {
        if (typeof res.data.login != 'undefined') {
          d.shareLink = 'http://store.lameleg.com?referrer=' + res.data.login
        } else {
          d.shareLink = '登录后才能看到您的推广连接'
        }
      })

    this.$http
      .get('http://store.lameleg.com:8080/loginless/pro', { credentials: true })
      .then(function (res) {
        for (var i = 0; i < res.data.length; i++) {
          var p = {}
          p.name = res.data[i].ProductName
          p.url =
            'http://store.lameleg.com:8080/pro/' +
            p.name +
            '?time=' +
            new Date().getTime() +
            '&referrer=' +
            VueCookies.get('referrer')
          p.price = res.data[i].ProductPrice
          d.products.push(p)
        }
        d.products.reverse()
      })
    return d
  },
}
</script>
 <style scoped>
.poptip p {
  font-size: 13px;
  text-align: left;
}
#sharelink {
  cursor: pointer;
  color: #2c66de;
  font-size: 13px;
}
</style>