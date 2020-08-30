<template>
  <div>
    <header class="header-area">
      <!-- Top Header Area Start -->
      <div class="top-header-area">
        <div class="container">
          <div class="row">

            <div class="col-6">
              <div class="top-header-content">
                <a href="#"><i class="fa fa-phone" aria-hidden="true"></i> <span>Call Us: 001-1234-88888</span></a>
                <a href="#"><i class="fa fa-envelope" aria-hidden="true"></i> <span>Email: info.cololib@gmail.com</span></a>
              </div>
            </div>

            <div class="col-6">
              <div class="top-header-content">
                <!-- Login -->
                <a href="#"><i class="fa fa-lock" aria-hidden="true"></i> <span>Login / Register</span></a>
                <!-- Language -->
                <div class="dropdown">
                  <a class="btn pr-0 dropdown-toggle" href="#" role="button" id="langdropdown" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"><img src="img/core-img/eng.png" alt=""> English</a>
                  <div class="dropdown-menu dropdown-menu-right" aria-labelledby="langdropdown">
                    <a class="dropdown-item" href="#">- Latvian</a>
                    <a class="dropdown-item" href="#">- Hindi</a>
                    <a class="dropdown-item" href="#">- Bangla</a>
                  </div>
                </div>
              </div>
            </div>

          </div>
        </div>
      </div>
      <!-- Top Header Area End -->

      <!-- Main Header Start -->
      <div class="main-header-area">
        <div class="classy-nav-container breakpoint-off">
          <div class="container">
            <!-- Classy Menu -->
            <nav class="classy-navbar justify-content-between" id="hamiNav">

              <!-- Logo -->
              <a class="nav-brand" href="index.html"><img src="./img/core-img/logo.png" alt=""></a>

              <!-- Navbar Toggler -->
              <div class="classy-navbar-toggler">
                <span class="navbarToggler"><span></span><span></span><span></span></span>
              </div>

              <!-- Menu -->
              <div class="classy-menu">
                <!-- Menu Close Button -->
                <div class="classycloseIcon">
                  <div class="cross-wrap"><span class="top"></span><span class="bottom"></span></div>
                </div>
                <!-- Nav Start -->
                <div class="classynav">
                  <ul id="nav">
                    <li class="active"><a href="./index.html">Home</a></li>
                    <li><a href="./hosting.html">Hosting</a></li>
                    <li><a href="#">Pages</a>
                      <ul class="dropdown">
                        <li><a href="./index.html">- Home</a></li>
                        <li><a href="./hosting.html">- Hosting</a></li>
                        <li><a href="./about.html">- About</a></li>
                        <li><a href="./blog.html">- Blog</a></li>
                        <li><a href="./single-blog.html">- Blog Details</a></li>
                        <li><a href="./404.html">- 404</a></li>
                        <li><a href="./coming-soon.html">- Coming Soon</a></li>
                        <li><a href="#">- Dropdown</a>
                          <ul class="dropdown">
                            <li><a href="#">- Dropdown Item</a></li>
                            <li><a href="#">- Dropdown Item</a></li>
                            <li><a href="#">- Dropdown Item</a></li>
                            <li><a href="#">- Dropdown Item</a></li>
                          </ul>
                        </li>
                      </ul>
                    </li>
                    <li><a href="./about.html">About</a></li>
                    <li><a href="./blog.html">Blog</a></li>
                    <li><a href="#">Contact</a></li>
                  </ul>

                  <!-- Live Chat -->
                  <div class="live-chat-btn ml-5 mt-4 mt-lg-0 ml-md-4">
                    <a href="#" class="btn hami-btn live--chat--btn"><i class="fa fa-comments" aria-hidden="true"></i> Live Chat</a>
                  </div>
                </div>
                <!-- Nav End -->
              </div>
            </nav>
          </div>
        </div>
      </div>
    </header>
  </div>
  <!-- <div class="hello">
    <img src="https://s2.ax1x.com/2019/09/12/nBpyZt.jpg" style="width:300px" />
    <br>kubernetes集群离线安装包, 一条命令高可用
    <p>
      使用kubeadm安装，包含安装所需的所有二进制文件，镜像文件，systemd配置，yaml配置与一些简单的启动脚本，五分钟搞定kubernetes安装~ 遇到任何问题可进群咨询，安装出问题提供远程协助,内核负载高可用，不依赖haproxy keepalived ansible,99年证书
    </p>
    <br />
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
import '@/assets/css/style.css'
export default {
  name: 'HelloWorld',
  props: {
    msg: String,
  },
  methods: {
    ok() {
      this.$Message.info('Clicked ok')
    },
    cancel() {
      this.$Message.info('Clicked cancel')
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
      console.log('cookie', VueCookies.get('referrer'))
      d.time += '&referrer=' + VueCookies.get('referrer')
    }
    console.log('url path query: ', d.time)

    this.$http
      .get('http://store.lameleg.com:8080/loginless/info/user', {
        credentials: true,
      })
      .then(
        function (res) {
          if (typeof res.data.login != 'undefined') {
            d.shareLink = 'http://store.lameleg.com?referrer=' + res.data.login
          } else {
            d.shareLink = '登录后才能看到您的推广连接'
          }
          console.log(res.data, 'info:', d)
        },
        function (res) {
          console.log(res.data)
        }
      )

    this.$http.get('http://store.lameleg.com:8080/loginless/pro', { credentials: true }).then(
      function (res) {
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
        console.log('products: ', res.data, 'length: ', res.data.length)
        console.log('products list: ', d.products, '')
      },
      function (res) {
        console.log(res.data)
      }
    )

    return d
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
