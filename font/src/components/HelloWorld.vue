<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <p>
      为了能记录您已经购买了哪些产品，和未来的分销功能，请登录:<br>
      <a href="https://github.com/login/oauth/authorize?client_id=89c1b05d77fb1c92a1ef&scope=user:email" target="_blank" rel="noopener">login github</a>.
    </p>
    <h3>商品列表</h3>
    <ul>
      <li><a href="http://store.lameleg.com:8080/pro/kubernetes1.13.1" rel="noopener">kubernetes1.13.1</a> 分销链接（开发中）</li>
    </ul>
    <h3>LINKS</h3>
    <ul>
      <li><a href="https://sealyun.com/pro/products/" target="_blank" rel="noopener">安装文档与常见问题</a></li><br/>
      <li><a href="https://sealyun.com" target="_blank" rel="noopener">博客</a></li><br/>
      <li><img src="https://sealyun.com/kubernetes-qrcode.jpg"></img></li><br/>公众号<br/>
      <li>联系方式：fhtjob@hotmail.com</li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
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
