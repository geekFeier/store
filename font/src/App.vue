<template>
  <div id="app">
    <header class="header-area" :style="hasAnnotaion ? 'margin-top:60px':'margin-top:0;'">
      <div class="main-header-area">
        <div class="classy-nav-container breakpoint-off">
          <div class="container">
            <nav class="classy-navbar justify-content-between" id="hamiNav">

              <a class="nav-brand" href="http://store.lameleg.com"><img width="100px" src="./assets/logo.png" alt=""></a>

              <div class="classy-navbar-toggler">
                <span class="navbarToggler"><span></span><span></span><span></span></span>
              </div>

              <div class="classy-menu">
                <div class="classycloseIcon">
                  <div class="cross-wrap"><span class="top"></span><span class="bottom"></span></div>
                </div>
                <div class="classynav">
                  <ul id="nav">
                    <li class="active"><a target="_blank" href="http://store.lameleg.com">SealYun</a></li>
                    <li><a target="_blank" href="https://sealyun.com">使用说明</a></li>

                  </ul>

                  <div class="live-chat-btn ml-5 mt-4 mt-lg-0 ml-md-4" style="margin-bottom:12px;">
                    <Avatar v-if="avata" :src="avata_url" />
                    <Tooltip v-else content="请使用github账户登录">
                      <a :href="loginurl" target="_blank" rel="noopener">
                        <li>
                          <i class="ivu-icon ivu-icon-logo-github"></i>
                          登录
                        </li>
                      </a>
                    </Tooltip>
                  </div>
                </div>
              </div>
            </nav>
          </div>
        </div>
      </div>
    </header>
    <banner />

    <annotation :hasAnnotaion="hasAnnotaion" />

    <goods-list />

    <section class="hami-cta-area">
      <div class="container">
        <div class="cta-text">
          <h2>已购用户</h2>
          <ul class="grid">
            <li v-for="u in payedUser">
              <a :href="'https://github.com/'+u.login" target="_blank">
                <img :src="u.avatar_url" alt="" />
              </a>
            </li>
          </ul>
        </div>
      </div>
    </section>

    <!-- 留言 -->
    <div id="gitalk-container"></div>
    <!-- 放底部 -->
    <!-- 加入组织: 钉钉群(35371178), Telegram -->
    <a target="_blank" href="https://mp.weixin.qq.com/s/Ra722VtdLitDbM0GExom6A">售后支持</a>
    <a href="javascript:void(0)">
      <Icon type="logo-github" />
      开源项目</a>
    <ul class="dropdown">
      <!-- <li><a>- 简单而不失强大</a></li> -->
      <li><a target="_blank" href="https://github.com/fanux/sealos">- kubernetes一键HA</a></li>
      <li><a target="_blank" href="https://github.com/fanux/fist">- 轻量级kubernetes管理工具</a></li>
      <li><a target="_blank" href="https://github.com/fanux/lhttp">- 好用的websocket框架</a></li>
    </ul>

    <a href="javascript:void(0)">
      友情链接</a>
    <ul class="dropdown">
      <li><a target="_blank" href="https://www.yangcs.net#sealyun">- 骚客米开朗琪杨</a></li>
      <li><a target="_blank" href="https://www.qikqiak.com/?utm_source=sealyun.com">- 阳明的博客</a></li>
      <li><a target="_blank" href="https://zhangguanzhang.github.io/#sealyun">- 张馆长</a></li>
    </ul>
    <!-- 
    <Menu mode="horizontal" :theme="theme1" active-name="1">
      <Row> 
        <Col span="3">
        <Tooltip content="年费会员，任意下载所有版本软件包">
          <a href="http://store.lameleg.com:8080/user/vip/charge" target="_blank" rel="noopener">
            <Button id="buy" type="success">会员69元/年</Button>
          </a>
        </Tooltip>
        </Col>
        <Col span="3">
        <MenuItem name="4">
        <Tooltip :content="'分享收入可提现金额,您可提现'+amount+'元'">
          <li @click="payeeFormCheck">
            <Modal v-model="payeeForm" @on-ok="ok" title="设置收款支付宝账号与提现密码" @on-cancel="cancel">
              <div class="payee">
                <label>提现账号</label>
                <input v-model="account" placeholder="收款支付宝账号">
                <br>
                <label>提现密码</label>
                <input v-model="passwd" placeholder="提现安全密码">
                <br>
                <label>确认密码</label>
                <input v-model="passwdCheck" placeholder="切勿使用支付宝密码">
                <br>
              </div>
            </Modal>
            <i class="ivu-icon ivu-icon-logo-yen">
              <Badge :count="amount">&nbsp;&nbsp;&nbsp;&nbsp;</Badge>
            </i>
          </li>
        </Tooltip>
        </MenuItem>

        <MenuItem name="5">
        <Avatar v-if="avata" :src="avata_url" />
        <Tooltip v-else content="请使用github账户登录">
          <a :href="loginurl" target="_blank" rel="noopener">
            <li>
              <i class="ivu-icon ivu-icon-logo-github"></i>
              Login
            </li>
          </a>
        </Tooltip>
        </MenuItem>
        </Col>
      </Row>
    </Menu>
    <Row> -->
    <!-- <index msg="Welcome to SealYun" /> -->
    <help />

  </div>
</template>

<script>
import index from './components/index.vue'
import goodsList from './components/goodsList.vue'
import banner from './components/banner.vue'
import help from './components/help.vue'
import annotation from './components/annotation.vue'
import VueCookies from 'vue-cookies'
import 'gitalk/dist/gitalk.css'
import Gitalk from 'gitalk'
import './assets/style.css'

/*
const gitalk = new Gitalk({
  clientID: '98478b0f6bfacff7cdf0',
  clientSecret: 'a882c5ed737d7453392b83c6b25e232a9d859d03',
  repo: 'https://github.com/fanux/store',
  owner: 'fanux',
  admin: ['fanux'],
  id: "gitalk.store.lameleg.com",      // Ensure uniqueness and length less than 50
  distractionFreeMode: false  // Facebook-like distraction free mode
})

gitalk.render('gitalk-container')
*/

export default {
  mounted() {
    const gitalk = new Gitalk({
      clientID: '98478b0f6bfacff7cdf0',
      clientSecret: 'a882c5ed737d7453392b83c6b25e232a9d859d03',
      repo: 'gitalk',
      owner: 'fanux',
      admin: ['fanux'],
      id: window.location.pathname, // Ensure uniqueness and length less than 50
      distractionFreeMode: false, // Facebook-like distraction free mode
    })

    gitalk.render('gitalk-container')
  },

  data() {
    let a = {
      hasAnnotaion: true, //是否有公告
      account: '',
      passwd: '',
      passwdCheck: '',
      amount: 1,
      theme1: 'light',
      avata: false,
      avata_url: '',
      payeeForm: false,
      payedUser: [],
      loginurl:
        'https://github.com/login/oauth/authorize?client_id=89c1b05d77fb1c92a1ef&scope=user:email',
    }

    this.$http
      .get('http://store.lameleg.com:8080/loginless/pro/kubernetes1.13.1/payed', {
        credentials: true,
      })
      .then(
        function (res) {
          for (let i = 0; i < res.data.length; i++) {
            if (res.data[i] == null) {
              continue
            }
            for (let j = i + 1; j < res.data.length; j++) {
              if (res.data[j] == null) {
                continue
              }
              if (res.data[i].login == res.data[j].login) {
                res.data[j].avata_url = 'https://avatars2.githubusercontent.com/u/8912557?v=4'
              }
            }
          }

          a.payedUser = res.data
        },
        function (res) {}
      )

    if (typeof this.$route.query.referrer != 'undefined') {
      VueCookies.set('referrer', this.$route.query.referrer)
    }
    // if (VueCookies.get('referrer') != null) {
    //   console.log('cookie', VueCookies.get('referrer'))
    // }
    this.$http
      .get('http://store.lameleg.com:8080/loginless/user/payee', {
        credentials: true,
      })
      .then(function (res) {
        a.amount = res.data.Amount
        if (typeof a.amount == 'undefined') {
          a.amount = 0
        }
      })

    this.$http
      .get('http://store.lameleg.com:8080/loginless/info/user', {
        credentials: true,
      })
      .then(function (res) {
        a.avata_url = res.data.avatar_url
        if (typeof a.avata_url != 'undefined') {
          a.avata = true
        }
      })

    return a
  },
  name: 'app',
  components: {
    index,
    banner,
    help,
    annotation,
    goodsList,
  },
  methods: {
    ok() {
      if (this.passwd != this.passwdCheck) {
        this.$Message.info('两次密码设置不相同')
        return
      }

      let para = {
        PayeeAccount: this.account,
        Passwd: this.passwd,
      }

      this.$http
        .put('http://store.lameleg.com:8080/user/info/payee', para, {
          credentials: true,
          responseType: 'json',
        })
        .then(
          function (res) {
            this.$Message.info(res.data.Reason)
          },
          function (res) {
            this.$Message.info(res.data.Reason)
          }
        )
    },
    cancel() {
      this.$Message.info('取消设置')
    },
    payeeFormCheck: function (event) {
      this.$http
        .post('http://store.lameleg.com:8080/user/info/withdraw', {}, { credentials: true })
        .then(
          function (res) {
            this.payeeForm = true
            if (res.data.Amount == 0) {
              console.log('withdraw is 0')
            }
            if (res.data.PayeeAccount == '' || res.data.Passwd == '') {
              // if not set, else withdraw
              this.payeeForm = true
            }
            this.account = res.data.PayeeAccount
            console.log(res.data)
          },
          function (res) {
            this.payeeForm = true
          }
        )
    },
  },
}
</script>

 <style lang="css">
.btn2 {
  text-decoration: none;
  transition: color 0.15s ease-in-out, background-color 0.15s ease-in-out,
    border-color 0.15s ease-in-out, box-shadow 0.15s ease-in-out;
}
</style>