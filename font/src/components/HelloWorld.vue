<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <p>
      为了能记录您已经购买了哪些产品，点击连接时会跳转到github进行第三方登录<br>
      网站丑的原因： 相信客户不会为网站漂亮买单<br>
      通过用户专有的分享连接交易完成的商品，用户分成60%<br>
      支付成功后会自动跳转到下载安装包<br>
      <!--a href="https://github.com/login/oauth/authorize?client_id=89c1b05d77fb1c92a1ef&scope=user:email" target="_blank" rel="noopener">login github</a-->.
    </p>
    <h3>商品列表</h3>
    <div id="pro-link">
    <ul>
      <li><a :href="time" rel="noopener">kubernetes1.13.1</a> 分享链接（开发中）</li>
    </ul>
    </div>
    <h3>LINKS</h3>
    <ul>
      <li><a href="https://sealyun.com/pro/products/" target="_blank" rel="noopener">安装文档与常见问题</a></li><br/>
      <li><a href="https://sealyun.com" target="_blank" rel="noopener">博客</a></li><br/>
      <li><img src="https://sealyun.com/kubernetes-qrcode.jpg" /></li><br/>公众号<br/>
      <li>联系方式：fhtjob@hotmail.com</li>
    </ul>
  </div>
</template>

<script  type="text/javascript">

export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  data: function(){
    return {time: 'http://store.lameleg.com:8080/pro/kubernetes1.13.1?time=' + (new Date()).getTime()}
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
