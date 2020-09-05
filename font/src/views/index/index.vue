<template>
  <div id="app">

    <annotation :hasAnnotaion="hasAnnotaion" />

    <head-common :hasAnnotaion="hasAnnotaion" />

    <banner />

    <description />

    <goods-list /> <!-- TODO: -->

    <buyer /> <!-- TODO: -->

    <foot-common />
  </div>
</template>

<script>
import goodsList from './component/goodsList.vue'
import banner from './component/banner.vue'
import annotation from './component/annotation.vue'
import buyer from './component/buyer.vue'
import description from './component/description.vue'
import footCommon from '@/components/layout/footCommon.vue'
import headCommon from '@/components/layout/headCommon.vue'
import VueCookies from 'vue-cookies'
import '../../assets/style.css'

export default {
  mounted() {},

  data() {
    return {
      hasAnnotaion: true, //是否有公告
    }
  },
  name: 'app',
  components: {
    banner,
    annotation,
    'head-common': headCommon,
    'foot-common': footCommon,
    goodsList,
    buyer,
    description,
  },
  created() {},
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
