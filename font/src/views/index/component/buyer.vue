<template>
  <div class="buyer">
    <section class="hami-cta-area">
      <div class="container">
        <div class="cta-text">
          <h2>超过 <span class="counter">2,000</span> 已购用户</h2>
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
  </div>
</template>

<script type="text/javascript">
import '../../../assets/style.css'
export default {
  name: 'buyer',

  data() {
    return {
      payedUser: [],
    }
  },
  created() {
    this.getPayedUser()
  },
  methods: {
    getPayedUser() {
      this.$http
        .get('http://store.lameleg.com:8080/loginless/pro/kubernetes1.13.1/payed', {
          credentials: true,
        })
        .then((res) => {
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

          this.payedUser = res.data
        })
    },
  },
}
</script> 