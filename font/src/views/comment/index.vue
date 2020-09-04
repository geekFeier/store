<template>
  <div class="comment">
    <!-- 留言 -->
    <div id="gitalk-container"></div>
  </div>
</template>

<script type="text/javascript">
import VueCookies from 'vue-cookies'
import 'gitalk/dist/gitalk.css'
import Gitalk from 'gitalk'
import '../../assets/style.css'
export default {
  name: 'comment',
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
    return {}
  },
}
</script> 