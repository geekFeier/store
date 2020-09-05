<template>
  <header class="header-area" :style="hasAnnotaion ? 'margin-top:60px':'margin-top:0;'">
    <div class="main-header-area">
      <div class="classy-nav-container breakpoint-off">
        <div class="container">
          <nav class="classy-navbar justify-content-between" id="hamiNav">
            <a class="nav-brand" href="http://store.lameleg.com"><img width="100px" src="../../assets/logo.png" alt=""></a>
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
                  <li><a target="_blank" href="./comment">评论区</a></li>
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
</template>


<script type="text/javascript">
import VueCookies from 'vue-cookies'
import '../../assets/style.css'
export default {
  name: 'headCommon',
  mounted() {},
  props: {
    hasAnnotaion: {
      type: Boolean,
      default: true,
    },
  },
  data() {
    let a = {
      account: '',
      passwd: '',
      passwdCheck: '',
      amount: 1,
      theme1: 'light',
      avata: false,
      avata_url: '',
      payeeForm: false,
      loginurl:
        'https://github.com/login/oauth/authorize?client_id=89c1b05d77fb1c92a1ef&scope=user:email',
    }

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
}
</script> 